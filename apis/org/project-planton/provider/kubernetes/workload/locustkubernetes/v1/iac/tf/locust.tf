# Main Python script (main.py)
resource "kubernetes_config_map" "main_py" {
  metadata {
    name      = "main-py"
    namespace = kubernetes_namespace.this.metadata[0].name
    labels    = local.final_labels
  }

  data = {
    "main.py" = var.spec.load_test.main_py_content
  }

  depends_on = [
    kubernetes_namespace.this
  ]
}

# Additional library files (lib_files_content)
resource "kubernetes_config_map" "lib_files" {
  metadata {
    name      = "lib-files"
    namespace = kubernetes_namespace.this.metadata[0].name
    labels    = local.final_labels
  }

  # Because lib_files_content is a map of filename->content,
  # we can directly assign it to the data block.
  data = var.spec.load_test.lib_files_content

  depends_on = [
    kubernetes_namespace.this
  ]
}

# Merge base helm values with user-provided overrides
locals {
  base_helm_values = {
    fullnameOverride = local.kube_service_name

    master = {
      replicas = var.spec.master_container.replicas
      resources = {
        requests = {
          cpu = try(var.spec.master_container.resources.requests.cpu, null)
          memory = try(var.spec.master_container.resources.requests.memory, null)
        }
        limits = {
          cpu = try(var.spec.master_container.resources.limits.cpu, null)
          memory = try(var.spec.master_container.resources.limits.memory, null)
        }
      }
    }

    worker = {
      replicas = var.spec.worker_container.replicas
      resources = {
        requests = {
          cpu = try(var.spec.worker_container.resources.requests.cpu, null)
          memory = try(var.spec.worker_container.resources.requests.memory, null)
        }
        limits = {
          cpu = try(var.spec.worker_container.resources.limits.cpu, null)
          memory = try(var.spec.worker_container.resources.limits.memory, null)
        }
      }
    }

    loadtest = {
      name                        = var.spec.load_test.name
      locust_locustfile_configmap = "main-py"
      locust_lib_configmap        = "lib-files"
    }
  }

  # Merge user-provided helm_values (if any) with base_helm_values
  merged_helm_values = merge(
    local.base_helm_values,
    try(var.spec.helm_values, {})
  )
}

resource "helm_release" "this" {
  name             = local.resource_id
  repository       = "https://charts.deliveryhero.io"
  chart = "locust"
  # set the chart version you want to deploy
  version          = "0.31.5"
  namespace        = kubernetes_namespace.this.metadata[0].name
  create_namespace = false

  values = [
    yamlencode(local.merged_helm_values)
  ]

  depends_on = [
    kubernetes_namespace.this,
    kubernetes_config_map.main_py,
    kubernetes_config_map.lib_files
  ]
}
