# Generated from pkg/kubernetes/kindscope/cluster_scoped.txt by
# pkg/iac/tofu/generators (KindScopeTF) -- never edit by hand. After a table
# refresh, regenerate with:
#   PLANTON_REGEN_KINDSCOPE_TF=1 go test ./pkg/iac/tofu/generators/ -run TestKindScopeTFDrift
#
# The built-in Kubernetes kinds that are cluster-scoped, keyed "<group>/<Kind>"
# (the bare kind for the core group), and the API groups Kubernetes itself
# serves ("" is the core group), whose other kinds are namespaced: the same
# tables the Pulumi module reads.
locals {
  kindscope_builtin_cluster_scoped = toset([
    "ComponentStatus",
    "Namespace",
    "Node",
    "PersistentVolume",
    "admissionregistration.k8s.io/MutatingAdmissionPolicy",
    "admissionregistration.k8s.io/MutatingAdmissionPolicyBinding",
    "admissionregistration.k8s.io/MutatingWebhookConfiguration",
    "admissionregistration.k8s.io/ValidatingAdmissionPolicy",
    "admissionregistration.k8s.io/ValidatingAdmissionPolicyBinding",
    "admissionregistration.k8s.io/ValidatingWebhookConfiguration",
    "apiextensions.k8s.io/CustomResourceDefinition",
    "apiregistration.k8s.io/APIService",
    "authentication.k8s.io/SelfSubjectReview",
    "authentication.k8s.io/TokenReview",
    "authorization.k8s.io/SelfSubjectAccessReview",
    "authorization.k8s.io/SelfSubjectRulesReview",
    "authorization.k8s.io/SubjectAccessReview",
    "certificates.k8s.io/CertificateSigningRequest",
    "certificates.k8s.io/ClusterTrustBundle",
    "flowcontrol.apiserver.k8s.io/FlowSchema",
    "flowcontrol.apiserver.k8s.io/PriorityLevelConfiguration",
    "networking.k8s.io/IPAddress",
    "networking.k8s.io/IngressClass",
    "networking.k8s.io/ServiceCIDR",
    "node.k8s.io/RuntimeClass",
    "rbac.authorization.k8s.io/ClusterRole",
    "rbac.authorization.k8s.io/ClusterRoleBinding",
    "resource.k8s.io/DeviceClass",
    "resource.k8s.io/DeviceTaintRule",
    "resource.k8s.io/ResourceSlice",
    "scheduling.k8s.io/PriorityClass",
    "storage.k8s.io/CSIDriver",
    "storage.k8s.io/CSINode",
    "storage.k8s.io/StorageClass",
    "storage.k8s.io/VolumeAttachment",
    "storage.k8s.io/VolumeAttributesClass",
    "storagemigration.k8s.io/StorageVersionMigration",
  ])
  kindscope_builtin_groups = toset([
    "",
    "admissionregistration.k8s.io",
    "apiextensions.k8s.io",
    "apiregistration.k8s.io",
    "apps",
    "authentication.k8s.io",
    "authorization.k8s.io",
    "autoscaling",
    "batch",
    "certificates.k8s.io",
    "coordination.k8s.io",
    "discovery.k8s.io",
    "events.k8s.io",
    "flowcontrol.apiserver.k8s.io",
    "networking.k8s.io",
    "node.k8s.io",
    "policy",
    "rbac.authorization.k8s.io",
    "resource.k8s.io",
    "scheduling.k8s.io",
    "storage.k8s.io",
    "storagemigration.k8s.io",
  ])
}
