# Local values for Cert-Manager deployment

locals {
  # Namespace (use from spec or default)
  namespace = var.spec.namespace

  # Helm chart constants
  helm_chart_name    = "cert-manager"
  helm_chart_repo    = "https://charts.jetstack.io"
  helm_chart_version = var.spec.helm_chart_version

  # ServiceAccount name
  ksa_name = "cert-manager"

  # Build ServiceAccount annotations for workload identity
  # Multiple providers may configure different annotations
  sa_annotations = merge(
    # GCP Workload Identity
    [for provider in var.spec.dns_providers :
      provider.gcp_cloud_dns != null ? {
        "iam.gke.io/gcp-service-account" = provider.gcp_cloud_dns.service_account_email
      } : {}
    ]...,
    # AWS IRSA
    [for provider in var.spec.dns_providers :
      provider.aws_route53 != null ? {
        "eks.amazonaws.com/role-arn" = provider.aws_route53.role_arn
      } : {}
    ]...,
    # Azure Managed Identity
    [for provider in var.spec.dns_providers :
      provider.azure_dns != null ? {
        "azure.workload.identity/client-id" = provider.azure_dns.client_id
      } : {}
    ]...
  )

  # Extract Cloudflare providers for secret creation
  cloudflare_providers = [
    for provider in var.spec.dns_providers :
    provider if provider.cloudflare != null
  ]

  # Build ClusterIssuer list (one per domain across all providers)
  cluster_issuers = flatten([
    for provider in var.spec.dns_providers : [
      for zone in provider.dns_zones : {
        domain       = zone
        provider_name = provider.name
        acme_email   = var.spec.acme.email
        acme_server  = var.spec.acme.server

        # Provider-specific configuration
        gcp_cloud_dns = provider.gcp_cloud_dns
        aws_route53   = provider.aws_route53
        azure_dns     = provider.azure_dns
        cloudflare    = provider.cloudflare
      }
    ]
  ])

  # Outputs for stack_outputs
  release_name = local.helm_chart_name

  # Solver identity (first non-null identity found)
  solver_identity = coalesce(
    [for provider in var.spec.dns_providers :
      provider.gcp_cloud_dns != null ? provider.gcp_cloud_dns.service_account_email : null
    ]...,
    [for provider in var.spec.dns_providers :
      provider.aws_route53 != null ? provider.aws_route53.role_arn : null
    ]...,
    [for provider in var.spec.dns_providers :
      provider.azure_dns != null ? provider.azure_dns.client_id : null
    ]...,
    ""
  )

  # Cloudflare secret name (first cloudflare provider's secret)
  cloudflare_secret_name = length(local.cloudflare_providers) > 0 ? "cert-manager-${local.cloudflare_providers[0].name}-credentials" : ""
}

