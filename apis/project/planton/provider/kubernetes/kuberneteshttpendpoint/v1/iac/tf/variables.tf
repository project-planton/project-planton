variable "metadata" {
  description = "metadata"
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = object({

      # name of the environment
      name = string

      # id of the environment
      id = string
    })

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # tags for the resource
    tags = list(string)
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # (optional for create) flag to toggle tls for istio-http-endpoint.
    # defaults to "false".
    # (important note) certificates are not created for endpoints that do not need tls.
    # (important note) endpoint domains with out tls enabled are not eligible to be used for
    # creating endpoints for postgres-clusters or kafka-clusters.
    is_tls_enabled = bool

    # name of the cluster-issuer to be used for provisioning the certificate.
    # this is only required when tls is enabled and can be left empty otherwise.
    # this will be one of the domains from gke-cluster.spec.ingress-domains for
    # gke-clusters deployed using planton-cloud provided pulumi module since a cluster-issuer for
    # each entry in the ingress-domains will be created with the name of the cluster-issuer set to the domain name.
    # if the cluster is not managed by planton-cloud, specify the cluster-issuer configured to issue
    # certificate for the http-endpoint domain name.
    # for all other scenarios, the devops team can deploy a customized pulumi module.
    cert_cluster_issuer_name = string

    # A flag to enable virtual service configuration compatible with gRPC-Web clients.
    # gRPC-Web clients rely on extra headers added by the Envoy proxy.
    is_grpc_web_compatible = bool

    # A list of routing rules to configure backends for URL paths of the domain.
    # Routes allow configuring requests to be routed to different microservices based on the URL path.
    # For example, if the endpoint domain name is `console.example.com`, then `console.example.com/public/api/*` can be
    # routed to the `public-api` microservice, and `console.example.com/private/api/*` can be routed to the `private-api` microservice.
    routing_rules = list(object({

      # The URL path prefix to match.
      # For example, if the endpoint domain name is `console.example.com`, and the `url_path_prefix` is `/api`,
      # then all requests matching `console.example.com/api/*` are forwarded to the configured Kubernetes service.
      url_path_prefix = string

      # The backend service to which the requests should be forwarded.
      backend_service = object({

        # The name of the Kubernetes service.
        name = string

        # The namespace of the Kubernetes service.
        namespace = string

        # The port on which the Kubernetes service is exposed.
        port = number
      })
    }))
  })
}