resource "helm_release" "helm_chart" {
  # Set the release name to your resource_id from locals
  name = local.resource_id

  # The Helm repo, chart, and version
  repository = local.helm_repo
  chart      = local.helm_chart
  version = local.helm_version

  # Install the chart into the namespace we created
  namespace = kubernetes_namespace.helm_release_namespace.metadata[0].name

  # # "values" must be a list of YAML strings
  # values = [
  #   yamlencode(local.helm_values)
  # ]
}
