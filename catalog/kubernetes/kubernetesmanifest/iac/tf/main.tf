# KubernetesManifest Terraform module.
#
# Applies raw multi-document YAML through kubectl_manifest (one resource per
# document), with an optional anchor namespace created first.
#
# NAMESPACE SEMANTICS (parity contract with the Pulumi module): documents
# that declare no metadata.namespace get the anchor namespace via
# override_namespace; documents with an explicit namespace pass through
# untouched. Cluster-scoped documents without a namespace also receive the
# override — the API server ignores metadata.namespace on cluster-scoped
# objects, so the outcome matches the Pulumi provider's scope-aware
# defaulting (which skips them client-side instead). Without an anchor,
# scope_check refuses namespaced documents that name no namespace, deciding
# scope from kindscope.tf exactly as the Pulumi module does.

# The optional anchor namespace. Created before any document applies;
# documents and namespace are deleted together on destroy.
resource "kubernetes_namespace_v1" "manifest_namespace" {
  count = local.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace_name
    labels = local.labels
  }
}

# One kubectl_manifest per document. kubectl_manifest applies server-side
# (same apply mechanism as the Pulumi provider, whose helper enables
# EnableServerSideApply) and needs no cluster connection at plan time, so
# offline plan proofs work and CRD+CR manifests apply in one pass.
resource "kubectl_manifest" "documents" {
  for_each = local.manifest_docs_by_identity

  yaml_body = each.value

  # Anchor-namespace defaulting: only for documents that declare none, and
  # only when the manifest has an anchor (without one, scope_check has
  # already refused every namespaced document that names no namespace).
  override_namespace = try(yamldecode(each.value).metadata.namespace, "") == "" && local.namespace_name != "" ? local.namespace_name : null

  server_side_apply = true
  force_conflicts   = true

  # skip_await parity: wait/wait_for_rollout here, SkipAwait on the Pulumi
  # ConfigGroup. Both engines default to awaiting: wait_for_rollout blocks
  # apply until workload rollouts complete, and wait blocks destroy until
  # each document is actually gone (foreground propagation) — matching the
  # Pulumi engine's apply/deletion awaits, so verify-clean phases never
  # race. (Await BREADTH differs benignly: the Pulumi engine also readiness-
  # checks non-workload kinds like Services; kubectl awaits workload
  # rollouts only. Behavioral difference documented, not a parity
  # exception — the applied objects are identical.)
  wait             = !try(var.spec.skip_await, false)
  wait_for_rollout = !try(var.spec.skip_await, false)

  depends_on = [kubernetes_namespace_v1.manifest_namespace, terraform_data.scope_check]
}

# The cluster's CRDs, read only when a document without a namespace has a kind
# neither built in nor defined by the manifest (see locals.scope_needs_cluster),
# so an anchored plan never contacts the cluster.
data "kubernetes_resources" "cluster_crds" {
  count = local.scope_needs_cluster ? 1 : 0

  api_version = "apiextensions.k8s.io/v1"
  kind        = "CustomResourceDefinition"
}

# Refuses, before any document applies, a manifest without an anchor whose
# namespaced documents name no namespace of their own.
resource "terraform_data" "scope_check" {
  count = local.namespace_name == "" ? 1 : 0

  lifecycle {
    precondition {
      condition     = length(local.unplaced_documents) == 0
      error_message = "spec.namespace is not set, and these namespaced documents name no metadata.namespace of their own: ${join(", ", local.unplaced_documents)} -- give each its namespace, or set spec.namespace as their anchor."
    }
  }
}
