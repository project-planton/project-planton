##############################################
# main.tf
#
# Creates NATS resources on Kubernetes using
# the official NATS Helm chart.
##############################################

# Create namespace for NATS deployment (conditionally based on create_namespace flag)
resource "kubernetes_namespace" "nats_namespace" {
  count = var.spec.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

# Generate random password for bearer token authentication
resource "random_password" "nats_bearer_token" {
  count = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "bearer_token" ? 1 : 0

  length  = 32
  special = false
}

# Generate random password for basic auth
resource "random_password" "nats_admin_password" {
  count = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "basic_auth" ? 1 : 0

  length  = 32
  special = false
}

# Create secret for bearer token authentication
resource "kubernetes_secret_v1" "nats_bearer_token_secret" {
  count = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "bearer_token" ? 1 : 0

  metadata {
    name      = "auth-nats"
    namespace = local.namespace
    labels    = local.labels
  }

  data = {
    "token" = base64encode(random_password.nats_bearer_token[0].result)
  }

  type = "Opaque"
}

# Create secret for basic auth admin credentials
resource "kubernetes_secret_v1" "nats_admin_secret" {
  count = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "basic_auth" ? 1 : 0

  metadata {
    name      = "auth-nats"
    namespace = local.namespace
    labels    = local.labels
  }

  data = {
    "user"     = base64encode("nats")
    "password" = base64encode(random_password.nats_admin_password[0].result)
  }

  type = "Opaque"
}

# Create secret for no-auth user (if enabled with basic auth)
resource "kubernetes_secret_v1" "nats_noauth_secret" {
  count = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "basic_auth" && try(var.spec.auth.no_auth_user.enabled, false) ? 1 : 0

  metadata {
    name      = "no-auth-user"
    namespace = local.namespace
    labels    = local.labels
  }

  data = {
    "user"     = base64encode("noauth")
    "password" = base64encode("nopassword")
  }

  type = "Opaque"
}

# Generate self-signed TLS certificate if TLS is enabled
resource "tls_private_key" "nats_tls" {
  count = try(var.spec.tls_enabled, false) ? 1 : 0

  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "nats_tls" {
  count = try(var.spec.tls_enabled, false) ? 1 : 0

  private_key_pem = tls_private_key.nats_tls[0].private_key_pem

  subject {
    common_name  = "${var.metadata.name}.${local.namespace}.svc.cluster.local"
    organization = "Project Planton"
  }

  validity_period_hours = 8760 # 1 year

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
    "client_auth",
  ]

  dns_names = [
    "${var.metadata.name}-nats",
    "${var.metadata.name}-nats.${local.namespace}",
    "${var.metadata.name}-nats.${local.namespace}.svc",
    "${var.metadata.name}-nats.${local.namespace}.svc.cluster.local",
  ]
}

# Create TLS secret
resource "kubernetes_secret_v1" "nats_tls_secret" {
  count = try(var.spec.tls_enabled, false) ? 1 : 0

  metadata {
    name      = "tls-${var.metadata.name}"
    namespace = local.namespace
    labels    = local.labels
  }

  data = {
    "tls.crt" = base64encode(tls_self_signed_cert.nats_tls[0].cert_pem)
    "tls.key" = base64encode(tls_private_key.nats_tls[0].private_key_pem)
  }

  type = "kubernetes.io/tls"
}

# Deploy NATS Helm chart
resource "helm_release" "nats" {
  name       = var.metadata.name
  repository = "https://nats-io.github.io/k8s/helm/charts"
  chart      = "nats"
  version    = "1.3.6"
  namespace  = local.namespace

  values = [
    yamlencode({
      # Container resources
      container = {
        merge = {
          resources = {
            limits = {
              cpu    = var.spec.server_container.resources.limits.cpu
              memory = var.spec.server_container.resources.limits.memory
            }
            requests = {
              cpu    = var.spec.server_container.resources.requests.cpu
              memory = var.spec.server_container.resources.requests.memory
            }
          }
        }
      }

      # NATS configuration
      config = merge(
        {
          # Clustering configuration
          cluster = {
            enabled  = var.spec.server_container.replicas > 1
            replicas = var.spec.server_container.replicas
          }

          # JetStream configuration
          jetstream = var.spec.disable_jet_stream ? {
            enabled = false
            } : {
            enabled = true
            fileStore = {
              enabled = true
              pvc = {
                size = var.spec.server_container.disk_size
              }
            }
          }
        },
        # Conditionally add auth configuration for basic auth
        try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "basic_auth" ? {
          patch = concat(
            [
              {
                op   = "add"
                path = "/authorization"
                value = {
                  users = concat(
                    [
                      {
                        username = "nats"
                        password = random_password.nats_admin_password[0].result
                      }
                    ],
                    try(var.spec.auth.no_auth_user.enabled, false) ? [
                      {
                        username = "noauth"
                        password = "nopassword"
                        permissions = {
                          publish   = try(var.spec.auth.no_auth_user.publish_subjects, [])
                          subscribe = []
                        }
                      }
                    ] : []
                  )
                }
              }
            ],
            try(var.spec.auth.no_auth_user.enabled, false) ? [
              {
                op    = "add"
                path  = "/no_auth_user"
                value = "noauth"
              }
            ] : []
          )
        } : {}
      )

      # Authentication configuration (bearer token)
      auth = try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "bearer_token" ? {
        enabled = true
        token = {
          users = [
            {
              existingSecret = {
                name = kubernetes_secret_v1.nats_bearer_token_secret[0].metadata[0].name
                key  = "token"
              }
            }
          ]
        }
        } : try(var.spec.auth.enabled, false) && try(var.spec.auth.scheme, "") == "basic_auth" ? {
        enabled = true
        basic   = {}
      } : {}

      # TLS configuration
      tls = try(var.spec.tls_enabled, false) ? {
        enabled = true
        secret = {
          name = kubernetes_secret_v1.nats_tls_secret[0].metadata[0].name
        }
      } : {}

      # NATS box configuration
      natsbox = {
        enabled = !try(var.spec.disable_nats_box, false)
      }
    })
  ]

  depends_on = [
    kubernetes_secret_v1.nats_bearer_token_secret,
    kubernetes_secret_v1.nats_admin_secret,
    kubernetes_secret_v1.nats_noauth_secret,
    kubernetes_secret_v1.nats_tls_secret
  ]
}

# Create LoadBalancer service for external access if ingress is enabled
resource "kubernetes_service_v1" "nats_external_lb" {
  count = try(var.spec.ingress.enabled, false) ? 1 : 0

  metadata {
    name      = "nats-external-lb"
    namespace = local.namespace
    labels    = local.labels

    annotations = try(var.spec.ingress.hostname, "") != "" ? {
      "external-dns.alpha.kubernetes.io/hostname" = var.spec.ingress.hostname
    } : {}
  }

  spec {
    type = "LoadBalancer"

    port {
      name        = "client"
      port        = 4222
      protocol    = "TCP"
      target_port = 4222
    }

    selector = {
      "app.kubernetes.io/name"      = "nats"
      "app.kubernetes.io/component" = "nats"
      "app.kubernetes.io/instance"  = var.metadata.name
    }
  }

  depends_on = [
    helm_release.nats
  ]
}
