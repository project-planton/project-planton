# SigNoz Kubernetes Pulumi Module

## Key Features

- **Standardized API Resource Model**  
  Provides a unified way to define and deploy SigNoz observability platform on Kubernetes. By describing resource allocations, database configuration, and component settings in a simple API resource, you ensure consistency across environments.

- **Automated Kubernetes Resource Creation**  
  Automatically creates Namespaces, Helm chart deployments, Services, and optional Load Balancer resources based on the provided specifications. Eliminates the need for hand-maintained YAML files.

- **Dual Database Modes**  
  Supports both self-managed ClickHouse deployment within the cluster and external ClickHouse integration, providing flexibility for different operational requirements.

- **OpenTelemetry Native**  
  Built on OpenTelemetry standards, SigNoz unifies logs, metrics, and traces in a single platform with seamless correlation and vendor-neutral instrumentation.

- **High Availability Support**  
  Configure distributed ClickHouse clusters with sharding and replication, coordinated by Zookeeper for production-grade deployments.

- **Scalable Data Ingestion**  
  Deploy multiple OpenTelemetry Collector replicas to handle high-volume telemetry data ingestion with horizontal scaling.

- **Component Independence**  
  Scale SigNoz UI/API, OpenTelemetry Collector, and ClickHouse independently based on workload requirements.

- **Persistence Management**  
  Configure data persistence with customizable disk sizes for ClickHouse. When enabled, telemetry data survives pod restarts.

- **Ingress Integration**  
  When enabled, the module sets up Load Balancer services with external DNS annotations for both SigNoz UI and OpenTelemetry Collector endpoints.

- **Output Exports**  
  Exports useful values such as namespace, service names, internal service FQDNs, port-forward commands, and ClickHouse credentials. These can be leveraged for further automation or debugging.

## Usage

See [examples.md](examples.md) for usage details and step-by-step examples. In general:

1. Define a YAML resource describing your SigNoz cluster using the **SignozKubernetes** API.
2. Run:
   ```bash
   planton pulumi up --stack-input <your-signoz-file.yaml>
   ```

to apply the resource on your cluster.

## Important: Docker Image Registry

**⚠️ Bitnami Registry Changes (September 2025)**

Due to Bitnami's transition to a paid model, this module now uses `docker.io/bitnamilegacy` registry for ClickHouse and ZooKeeper images (when using self-managed mode). The legacy images receive no updates or security patches but provide a temporary migration solution.

**Long-term Alternatives:**
- Subscribe to Bitnami Secure Images ($50k-$72k/year)
- Use official ClickHouse and ZooKeeper images
- Configure external ClickHouse database
- Build custom images from open-source Bitnami code (Apache 2.0)

For more details, see: [MIGRATION.md](MIGRATION.md) or https://github.com/bitnami/containers/issues/83267

## Getting Started

1. **Craft Your Specification**  
   Include SigNoz container settings, OpenTelemetry Collector configuration, database mode (self-managed or external), and optionally ingress preferences. If you need external access, enable ingress for both UI and data ingestion.

2. **Apply via CLI**  
   Execute `planton pulumi up --stack-input <signoz-spec.yaml>` (or your organization's standard CLI command). The Pulumi module automatically compiles your specification into Kubernetes resources.

3. **Validate & Observe**  
   Check the logs of your SigNoz deployment, confirm the Namespace, Deployments, and Services are created, and if ingress is enabled, verify external access.

## Module Structure

1. **Initialization**  
   Reads your `SignozKubernetesStackInput` (containing cluster creds, resource definitions), sets up local variables, and merges labels.

2. **Provider Setup**  
   Establishes a Pulumi Kubernetes Provider for your target cluster.

3. **Namespace Creation**  
   Creates (or identifies) a namespace to house all your SigNoz resources.

4. **Helm Chart Deployment**  
   Deploys the SigNoz Helm chart with configured values for:
   - SigNoz binary (UI, API, Ruler, Alertmanager)
   - OpenTelemetry Collector
   - ClickHouse (self-managed mode) or external connection
   - Zookeeper (for distributed ClickHouse clusters)

5. **Ingress Resources** (Optional)  
   If ingress is enabled, creates Kubernetes Gateway API resources for external access:
   - **SigNoz UI Ingress** (`ingress_signoz.go`): Certificate, Gateway, and HTTPRoutes for web UI access
   - **OTEL Collector Ingress** (`ingress_otel.go`): Certificate, Gateway, and HTTPRoutes for telemetry data ingestion

6. **Output Exports**  
   Publishes final references (e.g., namespace, service names, cluster endpoints, ClickHouse credentials), which can aid in post-deployment automation.

## Ingress Architecture

The module supports two types of external ingress, each implemented in a dedicated file:

### SigNoz UI Ingress

**Implementation**: `module/ingress_signoz.go`

Creates external access to the SigNoz web interface for viewing traces, metrics, and logs.

**Resources Created** (in `istio-ingress` namespace):
- **Certificate**: TLS certificate for the UI hostname (managed by cert-manager)
- **Gateway**: Kubernetes Gateway with HTTPS listener (port 443) and HTTP listener (port 80)
- **HTTPRoutes**: 
  - HTTP to HTTPS redirect (301)
  - HTTPS traffic routing to SigNoz frontend service (port 3301)

**Endpoints**:
- External: `{namespace}.{dns-domain}` (e.g., `signoz-app-prod-main.planton.live`)
- Internal: `{namespace}-internal.{dns-domain}` (e.g., `signoz-app-prod-main-internal.planton.live`)

**Configuration**:
```yaml
signozIngress:
  enabled: true
  dnsDomain: planton.live
```

### OTEL Collector Ingress

**Implementation**: `module/ingress_otel.go`

Creates external access to OpenTelemetry Collector for telemetry data ingestion from sources outside the Kubernetes cluster (CLI tools, services on developer laptops, external applications).

**Resources Created** (in `istio-ingress` namespace):
- **Certificate**: TLS certificate covering both gRPC and HTTP hostnames
- **Gateway**: Kubernetes Gateway with two HTTPS listeners:
  - `https-otel-grpc`: For gRPC traffic (OTLP/gRPC protocol)
  - `https-otel-http`: For HTTP traffic (OTLP/HTTP protocol)
- **HTTPRoutes**:
  - gRPC route to OTEL Collector service port 4317
  - HTTP route to OTEL Collector service port 4318

**Endpoints**:
- gRPC: `{namespace}-ingest-grpc.{dns-domain}` (e.g., `signoz-app-prod-main-ingest-grpc.planton.live`)
- HTTP: `{namespace}-ingest-http.{dns-domain}` (e.g., `signoz-app-prod-main-ingest-http.planton.live`)

**Configuration**:
```yaml
otelCollectorIngress:
  enabled: true
  dnsDomain: planton.live
```

**Usage Example (Java/Spring Boot)**:
```yaml
# application.yml
otel:
  exporter:
    otlp:
      endpoint: https://signoz-app-prod-main-ingest-grpc.planton.live:443
  traces:
    exporter: otlp
  metrics:
    exporter: otlp
```

**Usage Example (CLI/Environment Variables)**:
```bash
export OTEL_EXPORTER_OTLP_ENDPOINT=https://signoz-app-prod-main-ingest-grpc.planton.live:443
export OTEL_TRACES_EXPORTER=otlp
```

### Why HTTPRoute for gRPC?

gRPC uses HTTP/2 as its transport protocol. In Gateway API:
- HTTPRoute can handle gRPC traffic (gRPC is essentially HTTP/2)
- The Gateway listener uses HTTPS protocol with TLS termination
- The HTTPRoute routes traffic to the OTEL Collector gRPC port (4317)
- This is the standard, widely-supported approach in Gateway API

### TLS Configuration

- **External (Client to Gateway)**: HTTPS/TLS on port 443
- **Internal (Gateway to Service)**: Plain HTTP/gRPC (no TLS)
- TLS termination happens at the Gateway (Istio Ingress)
- Certificates automatically managed by cert-manager using Let's Encrypt

### Traffic Flow Example

**OTEL Collector gRPC Traffic**:
```
Java App on Laptop
  ↓ gRPC over HTTPS (port 443)
DNS: signoz-app-prod-main-ingest-grpc.planton.live
  ↓
Istio Ingress Gateway
  ↓ TLS Termination
HTTPRoute: https-otel-grpc
  ↓ gRPC (port 4317)
Service: main-otel-collector
  ↓
OTEL Collector Pod
  ↓ Processes & Batches
ClickHouse Database
```

## Benefits

- **Simplified Deployment**  
  Focus on high-level configuration rather than writing raw Kubernetes manifests. Consistent patterns reduce the risk of misconfiguration.

- **Unified Observability**  
  Single platform for logs, metrics, and traces eliminates tool sprawl and reduces operational complexity.

- **Cost Effective**  
  Self-hosted deployment with predictable infrastructure costs vs. per-GB SaaS pricing from proprietary solutions.

- **OpenTelemetry Native**  
  Leverage open standards and avoid vendor lock-in. Instrument once with OpenTelemetry, use with any backend.

- **Scalability**  
  Easily configure standalone or clustered deployments with independent component scaling to handle varying workloads.

- **Data Control**  
  Full control over telemetry data location and retention for compliance requirements.

- **Extensibility**  
  The module is built on Pulumi's Kubernetes provider. You can augment or override resources if your team needs advanced configurations through helm_values.

## Contributing

Contributions are always welcome! Please open an issue or submit a pull request in the main repository if you want to add features, fix bugs, or improve documentation.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to adapt it for your internal workflows.

