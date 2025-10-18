# Temporal Kubernetes Ingress Hostname Field

**Date**: October 18, 2025  
**Type**: Breaking Change, Enhancement  
**Component**: TemporalKubernetes

## Summary

Refactored the Temporal Kubernetes ingress configuration from a shared `IngressSpec` (with `enabled` and `dns_domain` fields) to a hierarchical structure with custom `TemporalKubernetesIngress` and `TemporalKubernetesIngressEndpoint` messages featuring `enabled` and `hostname` fields. This change gives users full control over both frontend (gRPC) and Web UI ingress hostnames independently, eliminates hostname auto-construction logic, and provides clear separation between the two Temporal access points.

## Motivation

### The Problem

The previous implementation used the shared `IngressSpec` from `project.planton.shared.kubernetes`:

```yaml
ingress:
  enabled: true
  dnsDomain: "planton.live"
```

This approach had several limitations:

1. **Hostname Auto-Construction**: The system automatically constructed hostnames from resource ID and DNS domain, giving users no control over the exact hostname pattern:
   - Frontend: `{namespace}-frontend.{dns-domain}`
   - Web UI: `{namespace}-ui.{dns-domain}`

2. **Single Enable Toggle**: Both frontend and Web UI ingress were controlled by a single `enabled` flag, preventing independent control. Users couldn't expose only the frontend or only the Web UI.

3. **Inflexibility**: Users couldn't specify custom hostnames like:
   - `temporal-grpc.example.com` for frontend
   - `temporal.example.com` for Web UI
   - Different domains for each endpoint

4. **Module Complexity**: The Pulumi module contained hostname construction logic that could be eliminated if users specified full hostnames directly.

5. **Shared Spec Limitations**: The generic `IngressSpec` was designed for multiple resource types, forcing Temporal to inherit patterns that didn't match its dual-endpoint architecture.

### The Solution

Replace flat ingress configuration with hierarchical structure:

```yaml
ingress:
  frontend:
    enabled: true
    hostname: "temporal-frontend.example.com"
  webUi:
    enabled: true
    hostname: "temporal-ui.example.com"
```

This approach:
- ✅ Provides clear hierarchical organization
- ✅ Gives users complete control over both hostnames independently
- ✅ Enables independent enable/disable for each endpoint
- ✅ Simplifies Pulumi module (removed hostname construction logic)
- ✅ Provides clearer, more intuitive API
- ✅ Enables any hostname pattern users need
- ✅ Maintains validation (hostname required when enabled)

## What's New

### 1. Hierarchical Ingress Structure

**Before (Flat Structure with Single Field)**:
```protobuf
message TemporalKubernetesSpec {
  project.planton.shared.kubernetes.IngressSpec ingress = 6;
}
```

**After (Hierarchical Structure)**:
```protobuf
message TemporalKubernetesSpec {
  TemporalKubernetesIngress ingress = 6;
}

message TemporalKubernetesIngress {
  TemporalKubernetesIngressEndpoint frontend = 1;
  TemporalKubernetesIngressEndpoint web_ui = 2;
}

message TemporalKubernetesIngressEndpoint {
  bool enabled = 1;
  string hostname = 2;
  
  option (buf.validate.message).cel = {
    id: "spec.ingress.hostname.required"
    expression: "!this.enabled || size(this.hostname) > 0"
    message: "hostname is required when ingress is enabled"
  };
}
```

### 2. Updated YAML Syntax

**Before**:
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-prod
spec:
  ingress:
    enabled: true
    dnsDomain: example.com
  # System creates:
  # - Frontend: temporal-prod-frontend.example.com
  # - Web UI: temporal-prod-ui.example.com
```

**After**:
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-prod
spec:
  ingress:
    frontend:
      enabled: true
      hostname: temporal-frontend.example.com
    webUi:
      enabled: true
      hostname: temporal-ui.example.com
  # Hostnames: exactly as specified by user
```

### 3. Simplified Pulumi Module

**Before** (Hostname Construction):
```go
// locals.go
func initializeLocals(ctx *pulumi.Context, stackInput *temporalkubernetesv1.TemporalKubernetesStackInput) *Locals {
    // ... other initialization

    if target.Spec.Ingress != nil &&
        target.Spec.Ingress.Enabled &&
        target.Spec.Ingress.DnsDomain != "" {

        locals.IngressFrontendHostname = fmt.Sprintf("%s-frontend.%s",
            locals.Namespace, target.Spec.Ingress.DnsDomain)
        locals.IngressUIHostname = fmt.Sprintf("%s-ui.%s",
            locals.Namespace, target.Spec.Ingress.DnsDomain)

        ctx.Export(OpExternalFrontendHostname, pulumi.String(locals.IngressFrontendHostname))
        ctx.Export(OpExternalUIHostname, pulumi.String(locals.IngressUIHostname))
    }
}
```

**After** (Direct Usage):
```go
// locals.go
func initializeLocals(ctx *pulumi.Context, stackInput *temporalkubernetesv1.TemporalKubernetesStackInput) *Locals {
    // ... other initialization

    // Frontend ingress
    if target.Spec.Ingress != nil &&
        target.Spec.Ingress.Frontend != nil &&
        target.Spec.Ingress.Frontend.Enabled &&
        target.Spec.Ingress.Frontend.Hostname != "" {

        locals.IngressFrontendHostname = target.Spec.Ingress.Frontend.Hostname
        ctx.Export(OpExternalFrontendHostname, pulumi.String(locals.IngressFrontendHostname))
    }

    // Web UI ingress
    if target.Spec.Ingress != nil &&
        target.Spec.Ingress.WebUi != nil &&
        target.Spec.Ingress.WebUi.Enabled &&
        target.Spec.Ingress.WebUi.Hostname != "" {

        locals.IngressUIHostname = target.Spec.Ingress.WebUi.Hostname
        ctx.Export(OpExternalUIHostname, pulumi.String(locals.IngressUIHostname))
    }
}
```

**Removed**:
- 8+ lines of hostname construction logic
- DNS domain parsing and validation

### 4. Updated Frontend Ingress Logic

**File**: `iac/pulumi/module/frontend_ingress.go`

**Before**:
```go
func frontendIngress(ctx *pulumi.Context, locals *Locals,
    createdNamespace *kubernetescorev1.Namespace) error {

    ingress := locals.TemporalKubernetes.Spec.Ingress
    if ingress == nil || !ingress.Enabled || ingress.DnsDomain == "" {
        return nil
    }
```

**After**:
```go
func frontendIngress(ctx *pulumi.Context, locals *Locals,
    createdNamespace *kubernetescorev1.Namespace) error {

    if locals.TemporalKubernetes.Spec.Ingress == nil ||
        locals.TemporalKubernetes.Spec.Ingress.Frontend == nil ||
        !locals.TemporalKubernetes.Spec.Ingress.Frontend.Enabled ||
        locals.TemporalKubernetes.Spec.Ingress.Frontend.Hostname == "" {
        return nil
    }
```

### 5. Updated Web UI Ingress Logic

**File**: `iac/pulumi/module/web_ui_ingress.go`

**Before**:
```go
func webUiIngress(ctx *pulumi.Context, locals *Locals,
    kubernetesProvider *kubernetes.Provider,
    createdNamespace *kubernetescorev1.Namespace) error {

    if locals.TemporalKubernetes.Spec.Ingress == nil ||
        !locals.TemporalKubernetes.Spec.Ingress.Enabled ||
        locals.TemporalKubernetes.Spec.Ingress.DnsDomain == "" ||
        locals.TemporalKubernetes.Spec.DisableWebUi {
        return nil
    }

    // ClusterIssuer from DNS domain
    IssuerRef: certmanagerv1.CertificateSpecIssuerRefArgs{
        Kind: pulumi.String("ClusterIssuer"),
        Name: pulumi.String(locals.TemporalKubernetes.Spec.Ingress.DnsDomain),
    },
```

**After**:
```go
func webUiIngress(ctx *pulumi.Context, locals *Locals,
    kubernetesProvider *kubernetes.Provider,
    createdNamespace *kubernetescorev1.Namespace) error {

    if locals.TemporalKubernetes.Spec.Ingress == nil ||
        locals.TemporalKubernetes.Spec.Ingress.WebUi == nil ||
        !locals.TemporalKubernetes.Spec.Ingress.WebUi.Enabled ||
        locals.TemporalKubernetes.Spec.Ingress.WebUi.Hostname == "" ||
        locals.TemporalKubernetes.Spec.DisableWebUi {
        return nil
    }

    // Extract domain from hostname for ClusterIssuer name
    hostnameParts := strings.Split(uiHostname, ".")
    var clusterIssuerName string
    if len(hostnameParts) > 1 {
        clusterIssuerName = strings.Join(hostnameParts[1:], ".")
    }

    IssuerRef: certmanagerv1.CertificateSpecIssuerRefArgs{
        Kind: pulumi.String("ClusterIssuer"),
        Name: pulumi.String(clusterIssuerName),
    },
```

## Implementation Details

### Protobuf Changes

**File**: `apis/project/planton/provider/kubernetes/workload/temporalkubernetes/v1/spec.proto`

**Changes Made**:
1. **Removed Import**: No longer imports `project/planton/shared/kubernetes/kubernetes.proto`
2. **Updated Field Type**: Line 42 changed from `project.planton.shared.kubernetes.IngressSpec` to `TemporalKubernetesIngress`
3. **Added Messages**: New `TemporalKubernetesIngress` and `TemporalKubernetesIngressEndpoint` messages with CEL validation

**Validation Strategy**: Uses CEL (Common Expression Language) to validate that `hostname` is required when `enabled` is true for each endpoint, providing clear error messages and type safety.

### Pulumi Module Updates

**Files Modified**:

1. **`iac/pulumi/module/locals.go`**:
   - Replaced single ingress check with separate frontend and WebUI checks
   - Uses hostnames directly from spec instead of constructing them
   - Removed DNS domain concatenation logic

2. **`iac/pulumi/module/frontend_ingress.go`**:
   - Updated condition to check `Ingress.Frontend.Enabled` and `Ingress.Frontend.Hostname`
   - Uses hostname directly from spec

3. **`iac/pulumi/module/web_ui_ingress.go`**:
   - Added `strings` import for domain extraction
   - Updated condition to check `Ingress.WebUi.Enabled` and `Ingress.WebUi.Hostname`
   - Extracts domain from hostname for ClusterIssuer name
   - Uses hostname directly in Gateway listeners

### Documentation Updates

All documentation files updated with correct syntax:

1. **`v1/README.md`**:
   - Updated main example with hierarchical ingress structure
   - Removed search attributes example (deprecated feature)

2. **`v1/examples.md`**:
   - Updated Examples 1-4 with new ingress format
   - Separate frontend and webUi configurations

3. **`v1/overview.md`**:
   - Updated example with new hierarchical structure

4. **`v1/hack/manifest.yaml`**:
   - Simplified to basic configuration (no ingress, no search attributes)

5. **`iac/pulumi/README.md`**:
   - Updated input variables table with separate frontend and webUi fields

6. **`v1/api_test.go`**:
   - Updated test input to use new ingress structure
   - Removed unused kubernetes package import

## Migration Guide

### Breaking Change Impact

This is a **breaking change** for all existing TemporalKubernetes resources with ingress enabled.

**Affected Users**: Users who have deployed Temporal with ingress enabled (likely a small subset given the resource's usage patterns).

### Migration Steps

#### Step 1: Identify Affected Resources

Find all TemporalKubernetes manifests with ingress configuration:

```bash
# Search for manifests with ingress enabled
grep -r "kind: TemporalKubernetes" -A 30 *.yaml | grep -E "(ingress|dnsDomain)"
```

#### Step 2: Update Manifest Structure

**Before Migration**:
```yaml
spec:
  database:
    backend: postgresql
  ingress:
    enabled: true
    dnsDomain: "example.com"
  # System creates:
  # - Frontend: {namespace}-frontend.example.com
  # - Web UI: {namespace}-ui.example.com
```

**After Migration**:
```yaml
spec:
  database:
    backend: postgresql
  ingress:
    frontend:
      enabled: true
      hostname: "temporal-frontend.example.com"
    webUi:
      enabled: true
      hostname: "temporal-ui.example.com"
  # User controls exact hostnames
```

**Field Path Changes**:
| Old Field Path           | New Field Path                  | Notes |
|--------------------------|---------------------------------|-------|
| `spec.ingress.enabled`   | `spec.ingress.frontend.enabled` + `spec.ingress.webUi.enabled` | ✅ Independent control |
| `spec.ingress.dnsDomain` | `spec.ingress.frontend.hostname` + `spec.ingress.webUi.hostname` | ⚠️ Changed: full hostnames |

#### Step 3: Determine Your Hostnames

The old system constructed hostnames as:
- Frontend: `{namespace}-frontend.{dns-domain}`
- Web UI: `{namespace}-ui.{dns-domain}`

You need to replicate this or choose new hostnames:

**Option A - Keep Existing Hostnames** (recommended for minimal disruption):
```yaml
# If your manifest had:
metadata:
  name: prod-temporal
spec:
  ingress:
    dnsDomain: "example.com"

# The old system created:
# - Frontend: prod-temporal-frontend.example.com
# - Web UI: prod-temporal-ui.example.com

# So use:
spec:
  ingress:
    frontend:
      enabled: true
      hostname: "prod-temporal-frontend.example.com"
    webUi:
      enabled: true
      hostname: "prod-temporal-ui.example.com"
```

**Option B - Choose New Hostnames** (take advantage of flexibility):
```yaml
spec:
  ingress:
    frontend:
      enabled: true
      hostname: "temporal-grpc.example.com"
    webUi:
      enabled: true
      hostname: "temporal.example.com"
```

#### Step 4: Update CLI and Regenerate Code

```bash
# Update CLI
brew update && brew upgrade project-planton

# Or fresh install
brew install project-planton/tap/project-planton

# Verify version
project-planton version

# For developers: regenerate protobuf stubs
cd apis
make protos
```

#### Step 5: Update DNS Records (if changing hostnames)

If you chose different hostnames than the auto-constructed ones:

1. **Before applying**: Note the current hostnames
2. **Update manifest**: Apply new configuration
3. **Verify Services/Gateways**: Resources will get new hostname configurations
4. **DNS propagation**: Update external-dns annotations or Gateway resources
5. **Verify**: Test the new hostnames work before removing old DNS entries

**Note**: If you keep the same hostnames, DNS records won't change.

#### Step 6: Apply Changes

```bash
# Preview changes
project-planton pulumi preview --manifest temporal.yaml

# Apply
project-planton pulumi up --manifest temporal.yaml
```

### Automated Migration Script

For users with many manifests:

```bash
#!/bin/bash
# migrate-temporal-ingress.sh

get_namespace() {
    local file=$1
    local namespace=$(yq eval '.metadata.name' "$file")
    echo "$namespace"
}

migrate_file() {
    local file=$1
    echo "Processing $file..."
    
    if ! grep -q "kind: TemporalKubernetes" "$file"; then
        echo "  Not a TemporalKubernetes resource, skipping"
        return
    fi
    
    # Check if ingress is configured
    local ingress_enabled=$(yq eval '.spec.ingress.enabled' "$file")
    local dns_domain=$(yq eval '.spec.ingress.dnsDomain // .spec.ingress.dns_domain' "$file")
    
    if [[ "$ingress_enabled" != "true" && "$dns_domain" == "null" ]]; then
        echo "  No ingress configuration, skipping"
        return
    fi
    
    # Extract namespace
    local namespace=$(get_namespace "$file")
    
    if [[ "$dns_domain" != "null" ]]; then
        # Construct hostnames matching old pattern
        local frontend_hostname="${namespace}-frontend.${dns_domain}"
        local webui_hostname="${namespace}-ui.${dns_domain}"
        
        echo "  Namespace: $namespace"
        echo "  DNS Domain: $dns_domain"
        echo "  Frontend Hostname: $frontend_hostname"
        echo "  Web UI Hostname: $webui_hostname"
        
        # Restructure ingress configuration
        yq eval -i "
          .spec.ingress.frontend.enabled = .spec.ingress.enabled |
          .spec.ingress.frontend.hostname = \"$frontend_hostname\" |
          .spec.ingress.webUi.enabled = .spec.ingress.enabled |
          .spec.ingress.webUi.hostname = \"$webui_hostname\" |
          del(.spec.ingress.enabled) |
          del(.spec.ingress.dnsDomain) |
          del(.spec.ingress.dns_domain) |
          del(.spec.ingress.isEnabled)
        " "$file"
    else
        # Just fix field naming (enabled vs isEnabled)
        yq eval -i "
          .spec.ingress.frontend.enabled = .spec.ingress.enabled |
          .spec.ingress.webUi.enabled = .spec.ingress.enabled |
          del(.spec.ingress.enabled) |
          del(.spec.ingress.isEnabled)
        " "$file"
    fi
    
    echo "  ✅ Migrated successfully"
}

# Find and migrate all TemporalKubernetes manifests
find . -name "*.yaml" -type f | while read file; do
    if grep -q "kind: TemporalKubernetes" "$file"; then
        migrate_file "$file"
    fi
done

echo ""
echo "✅ Migration complete!"
echo ""
echo "Next steps:"
echo "1. Review the changes with: git diff"
echo "2. Test with: project-planton pulumi preview --manifest <file>"
echo "3. Apply with: project-planton pulumi up --manifest <file>"
```

**Usage**:
```bash
chmod +x migrate-temporal-ingress.sh
./migrate-temporal-ingress.sh
```

## Examples

### Basic Configuration with Both Endpoints

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-basic
spec:
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      hostname: temporal-frontend.example.com
    webUi:
      enabled: true
      hostname: temporal-ui.example.com
```

### Production with External PostgreSQL

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-production
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: "postgres.prod.example.com"
      port: 5432
      username: "temporal_user"
      password: "secure_password"
  externalElasticsearch:
    host: "elasticsearch.prod.example.com"
    port: 9200
    user: "elastic_user"
    password: "elastic_password"
  ingress:
    frontend:
      enabled: true
      hostname: temporal-frontend-prod.company.com
    webUi:
      enabled: true
      hostname: temporal-ui-prod.company.com
```

### Frontend Only (No Web UI Access)

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-frontend-only
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: "postgres.example.com"
      port: 5432
      username: "temporal_user"
      password: "secure_password"
  disableWebUi: true  # Disable Web UI deployment
  ingress:
    frontend:
      enabled: true
      hostname: temporal-grpc.example.com
```

### Web UI Only (Internal Frontend)

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-ui-only
spec:
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: false  # Frontend accessible only within cluster
    webUi:
      enabled: true
      hostname: temporal.example.com
```

### Using with External-DNS (Frontend)

The `hostname` field works seamlessly with external-dns annotation:

```yaml
# Temporal manifest
spec:
  ingress:
    frontend:
      enabled: true
      hostname: temporal-grpc.example.com
```

This creates a LoadBalancer service with:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend-external-lb
  annotations:
    external-dns.alpha.kubernetes.io/hostname: temporal-grpc.example.com
spec:
  type: LoadBalancer
  # ... service configuration
```

External-DNS then automatically:
1. Detects the annotation
2. Waits for LoadBalancer IP assignment
3. Creates DNS A record: `temporal-grpc.example.com` → `<LoadBalancer-IP>`
4. Creates TXT record for ownership tracking

### Using with Gateway API (Web UI)

The Web UI hostname works with Gateway API and cert-manager:

```yaml
# Temporal manifest
spec:
  ingress:
    webUi:
      enabled: true
      hostname: temporal.example.com
```

This creates Gateway resources with appropriate listeners and HTTPRoutes, with TLS certificates from cert-manager.

## Benefits

### 1. Hierarchical Organization

**Before**: Flat structure with single enable toggle
```yaml
spec:
  ingress:
    enabled: true
    dnsDomain: "example.com"
```

**After**: Clear grouping of frontend and Web UI settings
```yaml
spec:
  ingress:
    frontend:
      enabled: true
      hostname: "temporal-grpc.example.com"
    webUi:
      enabled: true
      hostname: "temporal.example.com"
```

**Benefits**:
- Clear separation of frontend (gRPC) and Web UI concerns
- Easier to understand that they are independent endpoints
- Logical grouping of related fields
- More intuitive API structure

### 2. Independent Endpoint Control

**Before**: System decides hostname pattern for both endpoints, single toggle
```yaml
ingress:
  enabled: true
  dnsDomain: "example.com"
# System creates:
# - temporal-prod-frontend.example.com
# - temporal-prod-ui.example.com
# Both enabled or both disabled
```

**After**: User decides exact hostname for each endpoint independently
```yaml
ingress:
  frontend:
    enabled: true
    hostname: "temporal-grpc.example.com"
  webUi:
    enabled: false  # Can disable UI while keeping frontend
# System uses exact hostnames provided
```

### 3. Simplified Module Implementation

**Code Reduction**:
- Pulumi: 10+ lines removed (hostname construction logic)
- Total: ~10 lines of code eliminated

**Maintenance Benefits**:
- Fewer edge cases to handle
- No string manipulation or formatting logic
- Direct pass-through from manifest to Kubernetes
- Clearer code flow

### 4. Clearer API

**Before** (Multi-step mental model):
1. User provides DNS domain
2. System constructs frontend hostname from namespace + "-frontend" + DNS domain
3. System constructs Web UI hostname from namespace + "-ui" + DNS domain
4. Both endpoints enabled/disabled together

**After** (Direct mental model):
1. User provides exact frontend hostname (if needed)
2. User provides exact Web UI hostname (if needed)
3. Each endpoint independently controlled

### 5. Flexibility

Users can now use any hostname patterns:
- Different domains: `temporal-grpc.company.com`, `temporal-ui.internal.com`
- Subdomain patterns: `grpc.temporal.example.com`, `ui.temporal.example.com`
- Environment-specific: `temporal-frontend-prod.example.com`, `temporal-ui-prod.example.com`
- Descriptive: `workflow-engine.example.com`, `temporal-dashboard.example.com`
- Independent deployment: Enable frontend only or Web UI only

## Validation

### CEL Validation Rules

The new `TemporalKubernetesIngressEndpoint` message includes built-in validation:

```protobuf
option (buf.validate.message).cel = {
  id: "spec.ingress.hostname.required"
  expression: "!this.enabled || size(this.hostname) > 0"
  message: "hostname is required when ingress is enabled"
};
```

**Validation Behavior**:

✅ **Valid** - Both endpoints disabled:
```yaml
ingress:
  frontend:
    enabled: false
  webUi:
    enabled: false
```

✅ **Valid** - One endpoint enabled with hostname:
```yaml
ingress:
  frontend:
    enabled: true
    hostname: "temporal-frontend.example.com"
  webUi:
    enabled: false
```

✅ **Valid** - Both endpoints enabled with hostnames:
```yaml
ingress:
  frontend:
    enabled: true
    hostname: "temporal-frontend.example.com"
  webUi:
    enabled: true
    hostname: "temporal-ui.example.com"
```

❌ **Invalid** - Frontend ingress enabled without hostname:
```yaml
ingress:
  frontend:
    enabled: true
    # Error: hostname is required when ingress is enabled
```

❌ **Invalid** - Empty hostname with ingress enabled:
```yaml
ingress:
  webUi:
    enabled: true
    hostname: ""
    # Error: hostname is required when ingress is enabled
```

## Testing

### Test Scenarios

**Scenario 1: New Deployment with Both Endpoints**
```bash
# Create manifest with new syntax
cat > temporal-test.yaml <<EOF
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: test-temporal
spec:
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      hostname: test-temporal-frontend.example.com
    webUi:
      enabled: true
      hostname: test-temporal-ui.example.com
EOF

# Deploy
project-planton pulumi up --manifest temporal-test.yaml

# Verify LoadBalancer service created for frontend
kubectl get svc -n test-temporal frontend-external-lb -o yaml | \
  grep "external-dns.alpha.kubernetes.io/hostname"
# Should show: test-temporal-frontend.example.com

# Verify Gateway created for Web UI
kubectl get gateway -n istio-ingress | grep test-temporal
kubectl get httproute -n test-temporal
```

**Scenario 2: Frontend Only Deployment**
```bash
# Frontend ingress enabled, Web UI disabled
cat > temporal-frontend-only.yaml <<EOF
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: frontend-only
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: postgres.example.com
      port: 5432
      username: temporal
      password: pass
  disableWebUi: true
  ingress:
    frontend:
      enabled: true
      hostname: temporal-grpc.example.com
EOF

# Deploy
project-planton pulumi up --manifest temporal-frontend-only.yaml

# Verify only frontend LoadBalancer created
kubectl get svc -n frontend-only
# Should show frontend-external-lb only
```

**Scenario 3: Validation Error**
```bash
# Try invalid configuration
cat > temporal-invalid.yaml <<EOF
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: invalid-temporal
spec:
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      # Missing hostname - should fail validation
EOF

# Attempt deploy
project-planton pulumi up --manifest temporal-invalid.yaml
# Expected error: hostname is required when ingress is enabled
```

## Performance Impact

**No runtime performance impact**:
- Ingress configuration applied once at deployment time
- LoadBalancer and Gateway resources created/updated during Pulumi apply
- No ongoing hostname construction or manipulation
- Module simplification reduces deployment time slightly (fewer operations)

## Security Considerations

**No security impact**:
- Hostnames are public information (used in DNS records and Gateway listeners)
- No changes to authentication, authorization, or encryption
- LoadBalancer, Gateway API, and cert-manager security model unchanged
- Validation ensures hostname cannot be empty when ingress is enabled

**Operational Security**:
- Users must ensure hostname ownership before use
- DNS domains should be under organization's control
- Cert-manager requires proper ClusterIssuer configuration
- Gateway API requires proper RBAC permissions

## Related Documentation

- **Temporal Kubernetes API**: `apis/project/planton/provider/kubernetes/workload/temporalkubernetes/v1/`
- **Temporal Documentation**: https://docs.temporal.io/
- **Gateway API**: https://gateway-api.sigs.k8s.io/
- **CEL Validation**: https://github.com/bufbuild/protovalidate

## Breaking Change Checklist

- [x] Migration guide provided with step-by-step instructions
- [x] Automated migration script included
- [x] Documentation updated (examples, README, API docs)
- [x] Validation added to catch invalid configurations
- [x] Clear error messages for validation failures
- [x] Before/after comparison provided
- [x] Testing instructions included
- [x] Pulumi module updated
- [x] All examples updated to new syntax

## Deployment Status

✅ **Protobuf Contract**: Updated with hierarchical ingress structure and CEL validation  
✅ **Pulumi Module**: Hostname construction removed, direct usage implemented  
✅ **Documentation**: All examples and READMEs updated  
✅ **Migration Script**: Automated migration script provided  
✅ **Validation**: CEL validation ensures hostname required when enabled  
✅ **Build Verification**: All code compiled successfully with no errors

**Ready for**: User migration and deployment

## Future Enhancements

1. **Multiple Hostnames**: Support array of hostnames for multi-domain access
   ```yaml
   ingress:
     frontend:
       enabled: true
       hostnames:
         - temporal-frontend.example.com
         - temporal-grpc.example.com
   ```

2. **Hostname Validation**: Add regex validation for DNS compliance
   ```protobuf
   string hostname = 2 [
     (buf.validate.field).string.pattern = "^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
   ];
   ```

3. **TLS Configuration**: Add optional TLS cert configuration for frontend
   ```yaml
   ingress:
     frontend:
       enabled: true
       hostname: temporal-frontend.example.com
       tls:
         enabled: true
         secretName: temporal-frontend-tls
   ```

4. **Gateway Class Selection**: Support different gateway classes for Web UI
   ```yaml
   ingress:
     webUi:
       enabled: true
       hostname: temporal.example.com
       gatewayClassName: internal-gateway
   ```

## Support

For questions or issues with migration:
1. Review the [migration guide](#migration-guide) above
2. Use the [automated migration script](#automated-migration-script)
3. Check [examples](#examples) for reference configurations
4. Verify [validation rules](#validation) are met
5. Contact Project Planton support if issues persist

---

**Impact**: This change improves the TemporalKubernetes API by providing hierarchical organization, independent hostname control for frontend and Web UI endpoints, simplified implementation, and enabling flexible deployment patterns. The migration path is straightforward with clear documentation and automation tools.


