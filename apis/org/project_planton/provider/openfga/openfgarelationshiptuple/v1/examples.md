# OpenFGA Relationship Tuple Examples

## Basic Document Access

Grant a user viewer access to a specific document:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: anne-views-budget
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:anne"
  relation: "viewer"
  object: "document:budget-2024"
```

## Role-Based Access

Grant different roles to different users:

```yaml
# Owner has full control
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: bob-owns-project
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:bob"
  relation: "owner"
  object: "project:acme-corp"
---
# Editor can modify
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: carol-edits-project
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:carol"
  relation: "editor"
  object: "project:acme-corp"
---
# Viewer can only read
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: dave-views-project
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:dave"
  relation: "viewer"
  object: "project:acme-corp"
```

## Group Membership

Add a user to a group:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: anne-in-engineering
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:anne"
  relation: "member"
  object: "group:engineering"
```

## Userset Access

Grant access to all members of a group (userset):

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: engineering-views-docs
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "group:engineering#member"
  relation: "viewer"
  object: "folder:engineering-docs"
```

## Hierarchical Relationships

Create folder-document hierarchy (document in folder):

```yaml
# Document is in the folder
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: budget-in-reports
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "folder:reports"
  relation: "parent"
  object: "document:budget-2024"
---
# User has access to folder (inherited by documents via model)
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: anne-views-reports
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:anne"
  relation: "viewer"
  object: "folder:reports"
```

## Public Access (Wildcard)

Make a resource publicly accessible:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: public-announcement
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:*"
  relation: "viewer"
  object: "document:company-announcement"
```

## Conditional Access

Grant access with a condition (requires condition defined in model):

```yaml
# Access only from allowed IP ranges
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: anne-views-sensitive
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:anne"
  relation: "viewer"
  object: "document:sensitive-data"
  condition:
    name: "in_allowed_ip_range"
    contextJson: |
      {
        "allowed_ips": ["192.168.1.0/24", "10.0.0.0/8"]
      }
```

## Multi-Tenant Organization

Set up organization-level access:

```yaml
# User is admin of organization
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: alice-admin-acme
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:alice"
  relation: "admin"
  object: "organization:acme-corp"
---
# User is member of organization
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: bob-member-acme
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "user:bob"
  relation: "member"
  object: "organization:acme-corp"
---
# Project belongs to organization
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: project-in-acme
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  user: "organization:acme-corp"
  relation: "organization"
  object: "project:internal-tools"
```

## Specifying Authorization Model

Pin to a specific authorization model version:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaRelationshipTuple
metadata:
  name: anne-views-budget-v2
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  authorizationModelId: "01HABC_MODEL_V2_ID"  # Specific model version
  user: "user:anne"
  relation: "viewer"
  object: "document:budget-2024"
```

## Deployment

All examples require Terraform/Tofu as the provisioner:

```bash
# Create OpenFGA credentials
cat > openfga-creds.yaml << EOF
apiUrl: https://api.fga.example.com
apiToken: your-api-token
EOF

# Deploy the relationship tuple
project-planton apply --manifest tuple.yaml \
  --openfga-provider-config openfga-creds.yaml \
  --provisioner tofu
```

## Bulk Deployment

Deploy multiple tuples from a single file:

```bash
# tuples.yaml contains multiple YAML documents separated by ---
project-planton apply --manifest tuples.yaml \
  --openfga-provider-config openfga-creds.yaml \
  --provisioner tofu
```

## Verification

After deploying tuples, verify access using the OpenFGA CLI:

```bash
# Check if user:anne can view document:budget-2024
fga query check user:anne viewer document:budget-2024 \
  --store-id 01HXYZ... \
  --api-url http://localhost:8080
```
