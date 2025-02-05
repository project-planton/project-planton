variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(object({
      name = optional(string),
      id = optional(string),
    })),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # The GCP project ID in which the GKE cluster will be created.
    cluster_project_id = string

    istio = optional(object({
      enabled = optional(bool, false)
      cluster_region = optional(string)
      sub_network_self_link = optional(string)
    }))

    # A flag to control the installation of the PostgreSQL operator.
    install_postgres_operator = optional(bool, false)

    # A flag to control the installation of the Kafka operator.
    install_kafka_operator = optional(bool, false)

    # A flag to control the installation of the Solr operator.
    install_solr_operator = optional(bool, false)

    # A flag to control the installation of Kubecost.
    install_kubecost = optional(bool, false)

    # A flag to control the installation of Ingress NGINX.
    install_ingress_nginx = optional(bool, false)

    # A flag to control the installation of Cert-Manager.
    install_cert_manager = optional(bool, false)

    # A flag to control the installation of External DNS.
    install_external_dns = optional(bool, false)

    # A flag to control the installation of External Secrets.
    install_external_secrets = optional(bool, false)

    # A flag to control the installation of the Elastic operator.
    install_elastic_operator = optional(bool, false)

    # A flag to control the installation of the Keycloak operator.
    install_keycloak_operator = optional(bool, false)

  })
}
