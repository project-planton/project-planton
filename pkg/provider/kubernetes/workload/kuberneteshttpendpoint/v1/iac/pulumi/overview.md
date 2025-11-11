# Overview

The **Kubernetes HTTP Endpoint Pulumi Module** is designed to create and manage HTTP and HTTPS endpoints on Kubernetes clusters. This module takes a Kubernetes-style API resource (`KubernetesHttpEndpoint`) as input, which allows for the flexible configuration of routing rules, TLS certificates, and backend services. The module automates the creation of ingress resources, gateway listeners, and HTTP routes, making it easier to manage HTTP endpoints that can route traffic to different microservices based on URL paths.

The module supports both HTTP and HTTPS configurations, enabling TLS termination with the option to integrate with Kubernetes-managed certificates via cluster issuers. It also supports routing rules that can direct traffic to different backends based on URL prefixes. Developers can use this module to simplify the process of managing external and internal HTTP endpoints on Kubernetes, ensuring that requests are routed securely and efficiently to the appropriate services.

