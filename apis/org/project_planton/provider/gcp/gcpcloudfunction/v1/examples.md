# GcpCloudFunction API - Examples

Here are examples of how to create and configure a **GcpCloudFunction** API resource using the Planton Cloud CLI. The examples cover basic HTTP functions, event-driven functions, and production configurations with VPC, secrets, and scaling.

## Create using CLI

First, create a YAML file using the examples provided below. After the YAML file is created, you can apply the configuration using the following command:

```shell
planton apply -f <yaml-path>
```

## Basic HTTP Function (Development)

This basic example demonstrates a simple HTTP-triggered Cloud Function for development environments. It uses default settings for memory (256MB), timeout (60s), and scaling.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: hello-http-dev
spec:
  project_id: my-gcp-project
  region: us-central1
  build_config:
    runtime: python311
    entry_point: hello_http
    source:
      bucket: my-code-bucket
      object: functions/hello-http-v1.0.0.zip
```

## HTTP Function with Environment Variables

This example shows an HTTP function with custom configuration: increased memory, custom timeout, and environment variables for non-sensitive configuration.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: api-gateway-dev
spec:
  project_id: my-gcp-project
  region: us-central1
  build_config:
    runtime: nodejs20
    entry_point: handleRequest
    source:
      bucket: my-code-bucket
      object: functions/api-gateway-v2.1.3.zip
  service_config:
    available_memory_mb: 512
    timeout_seconds: 120
    environment_variables:
      LOG_LEVEL: "debug"
      API_VERSION: "v2"
      ENVIRONMENT: "development"
```

## Production HTTP Function with Secrets and VPC

This example demonstrates a production-ready HTTP function with:
- Secrets from Google Secret Manager
- VPC connectivity for private database access
- Custom service account for least-privilege access
- Min instances for warm starts (eliminates cold starts)
- Public access enabled

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: api-gateway-prod
  env: prod
  org: my-org
spec:
  project_id: my-gcp-project-prod
  region: us-east1
  build_config:
    runtime: nodejs22
    entry_point: handleRequest
    source:
      bucket: my-prod-code-bucket
      object: functions/api-gateway-v3.4.1.zip
  service_config:
    service_account_email: api-gateway@my-gcp-project-prod.iam.gserviceaccount.com
    available_memory_mb: 1024
    timeout_seconds: 60
    max_instance_request_concurrency: 80
    environment_variables:
      LOG_LEVEL: "info"
      API_VERSION: "v3"
      ENVIRONMENT: "production"
    secret_environment_variables:
      JWT_SECRET: jwt-secret
      DB_CONNECTION_STRING: db-connection
      API_KEY: external-api-key
    vpc_connector: projects/my-gcp-project-prod/locations/us-east1/connectors/vpc-connector
    vpc_connector_egress_settings: 0  # PRIVATE_RANGES_ONLY
    ingress_settings: 0  # ALLOW_ALL (public)
    scaling:
      min_instance_count: 2  # 2 warm instances for HA
      max_instance_count: 100
    allow_unauthenticated: true
```

## Private HTTP Function (Internal Only)

This example shows a private HTTP function accessible only from within the VPC or same GCP project. Public access is disabled.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: internal-api
spec:
  project_id: my-gcp-project
  region: us-central1
  build_config:
    runtime: python312
    entry_point: process_request
    source:
      bucket: my-code-bucket
      object: functions/internal-api-v1.2.0.zip
  service_config:
    service_account_email: internal-api@my-gcp-project.iam.gserviceaccount.com
    available_memory_mb: 512
    timeout_seconds: 180
    vpc_connector: projects/my-gcp-project/locations/us-central1/connectors/vpc-connector
    ingress_settings: 1  # ALLOW_INTERNAL_ONLY (private)
    allow_unauthenticated: false
```

## Event-Driven Function (Pub/Sub Trigger)

This example demonstrates an event-driven function triggered by Pub/Sub messages. It processes asynchronous jobs from a Pub/Sub topic.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: pubsub-worker
spec:
  project_id: my-gcp-project
  region: us-central1
  build_config:
    runtime: go122
    entry_point: ProcessMessage
    source:
      bucket: my-code-bucket
      object: functions/pubsub-worker-v1.0.0.zip
  service_config:
    service_account_email: pubsub-worker@my-gcp-project.iam.gserviceaccount.com
    available_memory_mb: 512
    timeout_seconds: 300
    secret_environment_variables:
      DB_PASSWORD: database-password
    vpc_connector: projects/my-gcp-project/locations/us-central1/connectors/vpc-connector
    scaling:
      min_instance_count: 0
      max_instance_count: 50
  trigger:
    trigger_type: 1  # EVENT_TRIGGER
    event_trigger:
      event_type: google.cloud.pubsub.topic.v1.messagePublished
      pubsub_topic: projects/my-gcp-project/topics/job-queue
      trigger_region: us-central1
      retry_policy: 1  # RETRY_POLICY_RETRY (automatic retries)
```

## Event-Driven Function (Cloud Storage Trigger)

This example shows a function triggered by Cloud Storage events (object creation). It processes new files uploaded to a GCS bucket.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: image-processor
spec:
  project_id: my-gcp-project
  region: us-central1
  build_config:
    runtime: python311
    entry_point: process_image
    source:
      bucket: my-code-bucket
      object: functions/image-processor-v2.0.0.zip
  service_config:
    service_account_email: image-processor@my-gcp-project.iam.gserviceaccount.com
    available_memory_mb: 2048
    timeout_seconds: 540  # 9 minutes for image processing
    environment_variables:
      OUTPUT_BUCKET: processed-images
      IMAGE_FORMAT: "webp"
    scaling:
      min_instance_count: 0
      max_instance_count: 10
  trigger:
    trigger_type: 1  # EVENT_TRIGGER
    event_trigger:
      event_type: google.cloud.storage.object.v1.finalized
      event_filters:
        - attribute: bucket
          value: upload-bucket
      trigger_region: us-central1
      retry_policy: 0  # RETRY_POLICY_DO_NOT_RETRY (at-most-once)
```

## Notes

- **Runtimes**: Use current, non-deprecated runtimes (python311+, nodejs20+, go121+, etc.)
- **Source Code**: Must be uploaded to a GCS bucket as a ZIP file
- **Service Account**: Always use a dedicated service account with least-privilege permissions for production
- **Secrets**: Use `secret_environment_variables` for sensitive data like API keys, passwords, connection strings
- **VPC**: Use `vpc_connector` to access private resources (Cloud SQL, Memorystore, internal APIs)
- **Scaling**: Set `min_instance_count > 0` to eliminate cold starts for latency-sensitive workloads
- **Triggers**: HTTP (default, trigger_type=0) or Event-driven (trigger_type=1)
- **Authentication**: Set `allow_unauthenticated: true` only for public APIs; keep false for private functions

These examples demonstrate various configurations for deploying Cloud Functions Gen 2 on GCP, from simple HTTP endpoints to complex event-driven workflows.
