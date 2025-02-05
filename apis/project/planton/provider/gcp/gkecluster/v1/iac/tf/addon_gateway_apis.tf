###############################################################################
# Gateway API CRDs
#
# This file applies the Gateway API CRDs from the official upstream repository:
# https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.1.0/config/crd/standard/
#
# It fetches each CRD YAML via the HTTP data source, then applies them to the
# Kubernetes cluster using the kubernetes_manifest resource.
#
# Note: This requires:
#   - The "kubernetes_manifest" resource (introduced in the official
#     hashicorp/kubernetes provider v2.24.0+).
#   - The "http" data source (from the hashicorp/http provider).
# Make sure you have these providers available in your Terraform configuration.
#
###############################################################################

locals {
  gateway_api_crd_filenames = [
    "gateway.networking.k8s.io_gatewayclasses.yaml",
    "gateway.networking.k8s.io_gateways.yaml",
    "gateway.networking.k8s.io_grpcroutes.yaml",
    "gateway.networking.k8s.io_httproutes.yaml",
    "gateway.networking.k8s.io_referencegrants.yaml"
  ]
}

data "http" "gateway_api_crds" {
  for_each = toset(local.gateway_api_crd_filenames)
  url      = "https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.1.0/config/crd/standard/${each.value}"
}

# issue: https://github.com/hashicorp/terraform-provider-kubernetes/issues/1428
# resource "kubernetes_manifest" "gateway_api_crds" {
#   for_each = data.http.gateway_api_crds
#   manifest = yamldecode(each.value.response_body)
# }
