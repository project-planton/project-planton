# Overview

The **Kubernetes HTTP Endpoint API resource** provides a consistent and streamlined interface for creating and managing HTTP endpoints in Kubernetes environments using Istio for ingress management. This resource enables the configuration of routes and backend services to facilitate HTTP traffic routing within the Kubernetes cluster. It supports various features such as TLS, gRPC-web compatibility, and path-based routing, ensuring flexibility and compliance with organizational needs.

## Why We Created This API Resource

Managing HTTP endpoints in Kubernetes can be complex, especially when dealing with advanced routing configurations, TLS management, and gRPC-web compatibility. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:

- **Simplify Deployment**: Easily create and manage HTTP endpoints with Istio ingress for Kubernetes applications.
- **Ensure Consistency**: Maintain uniform endpoint configurations across different environments and clusters.
- **Enhance Flexibility**: Configure path-based routing, enabling traffic to be routed to different services based on URL paths.
- **Optimize Security**: Optionally enable TLS for secure communications and manage certificates via the cluster issuer.

## Key Features

### Environment Integration

- **Environment Info**: Integrates seamlessly with our environment management system to deploy HTTP endpoints within specific environments.
- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments, ensuring consistent and repeatable provisioning processes.

### Kubernetes Credential Management

- **Kubernetes Cluster Credential ID**: Utilizes the specified Kubernetes credentials (`kubernetes_cluster_credential_id`) to ensure secure and authorized operations within Kubernetes clusters.

### Flexible HTTP Endpoint Configuration

#### TLS Enablement

- **TLS Support**: Toggle TLS on or off (`is_tls_enabled`) for the endpoint. When TLS is enabled, the API manages the certificate provisioning using the specified cluster issuer.
- **Cluster Issuer**: Specify the `cert_cluster_issuer_name` for issuing TLS certificates. This field is optional and only required if TLS is enabled.

#### gRPC-Web Compatibility

- **gRPC-Web Support**: Enable gRPC-web compatibility (`is_grpc_web_compatible`) to configure the endpoint for gRPC-web clients. Envoy proxy adds the necessary headers to support gRPC-web traffic.

#### Routing Rules

- **Path-Based Routing**: Configure multiple `routing_rules` that define how HTTP traffic should be routed based on URL path prefixes. This allows for fine-grained control of traffic flow within the Kubernetes cluster.
    - **URL Path Prefix**: For example, requests to `console.example.com/api/*` can be routed to a specific service, while `console.example.com/private/api/*` can be routed to another.
    - **Backend Service**: Define the backend Kubernetes service to handle traffic for each URL path prefix. The backend service is identified by its:
        - **Service Name**: The name of the Kubernetes service.
        - **Namespace**: The namespace in which the service is deployed.
        - **Service Port**: The port on which the service listens.

## Benefits

- **Simplified Deployment**: Abstracts the complexities of managing HTTP endpoints in Kubernetes into an easy-to-use API resource.
- **Consistency**: Ensures all HTTP endpoint configurations adhere to organizational standards and best practices.
- **Flexible Routing**: Supports path-based routing, allowing traffic to be directed to different services based on URL patterns.
- **Security**: Optional TLS support provides secure communication channels, and certificate management is simplified via cluster issuers.
- **gRPC-Web Support**: Seamlessly integrates gRPC-web traffic handling, making it easier to build services that rely on gRPC over HTTP.
