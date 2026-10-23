# Computed values shared across the module.

locals {
  # Resource-identity labels: the kuberneteslabelkeys set, identical to what
  # the Pulumi module stamps for the same manifest. Stamped on the namespace
  # this module creates — NEVER injected into the manifest's own documents
  # (the manifest is applied exactly as the user wrote it).
  identity_labels_base = {
    "planton.ai/resource" = "true"
    "planton.ai/name"     = var.metadata.name
    "planton.ai/kind"     = "KubernetesManifest"
  }

  id_label = (
    var.metadata.id != null && try(var.metadata.id, "") != ""
  ) ? { "planton.ai/id" = var.metadata.id } : {}

  org_label = (
    var.metadata.org != null && try(var.metadata.org, "") != ""
  ) ? { "planton.ai/organization" = var.metadata.org } : {}

  env_label = (
    var.metadata.env != null && try(var.metadata.env, "") != ""
  ) ? { "planton.ai/environment" = var.metadata.env } : {}

  labels = merge(
    local.identity_labels_base,
    local.id_label,
    local.org_label,
    local.env_label,
  )

  # The anchor namespace (the generator flattens the spec's value-or-ref to a
  # plain string; "" when the manifest has none) and whether this module
  # creates it.
  namespace_name   = try(var.spec.namespace, "")
  create_namespace = try(var.spec.create_namespace, false)

  # Document split rule (identical in the Pulumi module's inventory parser):
  # separators are lines starting with `---`; the prepended newline makes a
  # manifest that STARTS with `---` yield an empty first chunk instead of a
  # missed document. Blank chunks are dropped here; comment-only chunks
  # decode to null and are dropped below; an INVALID chunk fails the plan
  # through yamldecode — never silently skipped (the Pulumi engine rejects
  # the same document at apply).
  manifest_document_chunks = [
    for doc in split("\n---", "\n${var.spec.manifest_yaml}") : doc
    if trimspace(doc) != ""
  ]

  # Decoded documents in manifest order (nulls from comment-only chunks
  # dropped). List comprehension keeps the original document order — never
  # rebuild this from map keys, which sort lexically.
  manifest_documents = [
    for doc in local.manifest_document_chunks : doc
    if yamldecode(doc) != null
  ]

  # for_each keys carry the document's full identity
  # (apiVersion/Kind/namespace/name) so reordering documents never re-keys
  # Terraform state addresses (index keys would churn every later resource).
  # Two documents with the same identity collide loudly at plan time — a
  # duplicate document is a manifest bug, not something to apply twice.
  manifest_docs_by_identity = {
    for doc in local.manifest_documents :
    format("%s/%s/%s/%s",
      try(yamldecode(doc).apiVersion, ""),
      try(yamldecode(doc).kind, ""),
      try(yamldecode(doc).metadata.namespace, ""),
      try(yamldecode(doc).metadata.name, ""),
    ) => doc
  }

  # The applied-resource inventory ("apiVersion/Kind/name" per document,
  # manifest order) — derived from the input YAML so both engines export an
  # identical list regardless of how each engine tracks child resources.
  applied_resources = [
    for doc in local.manifest_documents :
    format("%s/%s/%s",
      try(yamldecode(doc).apiVersion, ""),
      try(yamldecode(doc).kind, ""),
      try(yamldecode(doc).metadata.name, ""),
    )
  ]

  # ---- No-anchor scope check (twin of the Pulumi module's
  # refuseUnplacedDocuments, deciding from the same tables in kindscope.tf) --
  # Without an anchor, a namespaced document that names no namespace of its
  # own has nowhere it was asked to go, so the plan refuses it instead of
  # letting the API place it in "default".
  unanchored_documents = local.namespace_name != "" ? [] : [
    for doc in local.manifest_documents : {
      label = trimspace(format("%s %s %s", try(yamldecode(doc).apiVersion, ""), try(yamldecode(doc).kind, ""), try(yamldecode(doc).metadata.name, "")))
      group = length(split("/", try(yamldecode(doc).apiVersion, ""))) == 2 ? split("/", yamldecode(doc).apiVersion)[0] : ""
      key = (
        length(split("/", try(yamldecode(doc).apiVersion, ""))) == 2
        ? format("%s/%s", split("/", yamldecode(doc).apiVersion)[0], try(yamldecode(doc).kind, ""))
        : try(yamldecode(doc).kind, "")
      )
    }
    if try(yamldecode(doc).metadata.namespace, "") == ""
  ]

  # Kinds the manifest's own CRDs define, and whether each is namespaced.
  manifest_defined_kinds = {
    for doc in local.manifest_documents :
    format("%s/%s", yamldecode(doc).spec.group, yamldecode(doc).spec.names.kind) => try(yamldecode(doc).spec.scope, "Namespaced") != "Cluster"
    if try(yamldecode(doc).kind, "") == "CustomResourceDefinition" && startswith(try(yamldecode(doc).apiVersion, ""), "apiextensions.k8s.io/")
  }

  # The cluster is read only for a kind neither built in nor defined above.
  scope_needs_cluster = length([
    for d in local.unanchored_documents : d
    if !contains(local.kindscope_builtin_cluster_scoped, d.key) && !contains(local.kindscope_builtin_groups, d.group) && !contains(keys(local.manifest_defined_kinds), d.key)
  ]) > 0

  cluster_crd_kinds = local.scope_needs_cluster ? {
    for crd in data.kubernetes_resources.cluster_crds[0].objects :
    format("%s/%s", crd.spec.group, crd.spec.names.kind) => crd.spec.scope != "Cluster"
  } : {}

  # Namespaced (or of unknown scope) and naming no namespace: refused.
  unplaced_documents = [
    for d in local.unanchored_documents : d.label
    if !contains(local.kindscope_builtin_cluster_scoped, d.key) && (
      contains(local.kindscope_builtin_groups, d.group) ? true :
      lookup(local.manifest_defined_kinds, d.key, lookup(local.cluster_crd_kinds, d.key, true))
    )
  ]
}
