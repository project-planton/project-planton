# Kubernetes HTTP Endpoint Pulumi Module

## Key Features

- **API Resource-Driven Management:** The module uses a Kubernetes-style API resource (`KubernetesHttpEndpoint`) to define and manage HTTP endpoints. This declarative approach aligns with Kubernetes best practices and makes it easy to integrate into existing workflows.
  
- **HTTP and HTTPS Support:** The module supports both HTTP and HTTPS configurations for external endpoints. It allows users to enable TLS and manage certificates via Kubernetes' native certificate management capabilities. For HTTPS endpoints, the module can create and manage TLS certificates, including integration with Kubernetes cluster issuers.

- **Dynamic Routing Rules:** The module allows for flexible routing of HTTP requests based on URL path prefixes. Users can define routing rules that direct traffic to different backend Kubernetes services, enabling a microservices architecture where different URL paths can route traffic to separate services.

- **TLS Certificate Management:** For TLS-enabled endpoints, the module integrates with Kubernetes' `cert-manager` to automatically provision and renew certificates. You can specify the `cert_cluster_issuer_name` to leverage Kubernetes' native certificate management for securing your HTTPS endpoints.

- **gRPC-Web Compatibility:** The module includes support for gRPC-Web clients by enabling specific routing rules that allow gRPC-Web clients to function with Kubernetes ingress controllers. This ensures compatibility for applications using the gRPC-Web protocol.

- **Port and Backend Service Management:** Each route can be configured to forward requests to specific Kubernetes services and ports, enabling precise control over how traffic flows within your Kubernetes cluster. The module can manage the entire lifecycle of these backend services, including routing updates and service management.

- **Ingress and Gateway Automation:** The module automates the creation of Kubernetes gateway listeners and HTTP routes based on the endpoint specifications. Users can define multiple routes for a single domain, ensuring efficient traffic management and control over how requests are handled.

## Inputs

The module accepts the following key input parameters from the `KubernetesHttpEndpoint` API resource:

- **kubernetes_credential_id**: (Required) The ID of the Kubernetes cluster credentials to authenticate and deploy resources on the target cluster.
  
- **is_tls_enabled**: (Optional) A flag to enable TLS for the HTTP endpoint. When enabled, the module provisions and manages TLS certificates using Kubernetes-native capabilities.

- **cert_cluster_issuer_name**: (Optional) The name of the cluster issuer used for certificate provisioning. This is required when TLS is enabled and can be left empty if TLS is not enabled.

- **is_grpc_web_compatible**: (Optional) A flag to enable gRPC-Web compatibility. This configures the endpoint to add the necessary headers and routing rules for gRPC-Web clients.

- **routing_rules**: (Optional) A list of routing rules that define how traffic is routed based on URL path prefixes. Each rule specifies a `url_path_prefix` and a backend Kubernetes service to handle the requests that match the prefix.

## Outputs

The module provides several key outputs, which are critical for managing and accessing the deployed HTTP endpoints:

- **namespace**: The Kubernetes namespace where the HTTP endpoint is deployed.

- **service**: The name of the Kubernetes service associated with the backend of the HTTP endpoint.

- **port_forward_command**: (Optional) A command for setting up port forwarding to access the HTTP endpoint locally when ingress is not enabled.

- **tls_enabled_status**: Whether TLS is enabled for the endpoint.

- **endpoint_domain_name**: The domain name for the HTTP/HTTPS endpoint, used to access the service externally.

## Benefits

This module abstracts the complexity of managing ingress and HTTP endpoints in Kubernetes, making it easier to handle secure, scalable, and well-structured traffic routing. With built-in support for TLS, routing, gRPC-Web, and Kubernetes-native management of services, the module enables developers and DevOps teams to efficiently manage HTTP endpoints for microservices and applications. The declarative approach ensures consistent and repeatable deployments while reducing manual intervention and configuration errors.

By utilizing Kubernetes' native resources and integrating with services like `cert-manager` for certificate management, the module provides a production-ready solution for handling HTTP and HTTPS traffic. It also ensures flexibility in routing requests to various services based on URL paths, enabling efficient microservice architectures.

## Usage

Refer to the **example section** for detailed usage instructions on how to configure the API resource and use this Pulumi module.
