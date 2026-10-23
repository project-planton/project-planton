# Workloads Mount the Agent Substrate's Identity, and Manifests Need No Namespace

## Why

A product that runs agents on a KubernetesAgentSubstrate dials its API with TLS, verifying the server against a ClusterTrustBundle, and pushes to actors through a router that demands a client certificate. KubernetesDeployment could mount the token the API wants but neither the trust bundle nor the certificate, so the only way through was copying the CA's PEM into a ConfigMap by hand — broken the first time the substrate rotated it (#631). Separately, KubernetesManifest required a namespace even for a manifest of only cluster-scoped objects, where the value changed nothing, and with `create_namespace` it created and deleted an empty namespace nobody asked for (#630).

## What changed

- **The shared VolumeMount gains two projected sources**, written by both engines' shared code for every kind that mounts volumes (the five workload kinds' app, sidecar and init containers, KubernetesOtelCollector and KubernetesLangfuse):
  - `cluster_trust_bundle` — one bundle by name (a reference to the substrate's `api_trust_bundle_name` by default) or every bundle of a signer with a label selector; the kubelet keeps the file current as the bundle rotates.
  - `pod_certificate` — a kubelet-generated key and the certificate a signer issues for the pod, rotated before expiry.
  - A mount takes at most one source (a rule the spec now enforces; a mount with none names a volume defined elsewhere, and the OpenTofu modules no longer turn such a mount into an empty directory).
- **The OpenTofu workload modules refuse the two new sources at plan**, naming the mounts and asking for the Pulumi engine: the hashicorp/kubernetes provider models neither projection. The refusal is proven live, beside the Pulumi lane that mounts all three identity files from a real substrate.
- **KubernetesAgentSubstrate exports `router_client_signer_name`**, the signer of the client certificates its router's TLS endpoint accepts.
- **KubernetesManifest's namespace is optional.** Without one, a namespaced document that names no namespace of its own is refused before anything applies, so nothing lands in `default` by accident; `create_namespace` needs a namespace. Both engines decide a document's scope from one table, `pkg/kubernetes/kindscope` — the built-in cluster-scoped kinds and API groups read from real clusters' discovery, the manifest's own CRDs, and the cluster's CRDs only when a kind is neither. The OpenTofu copy is generated (`kindscope.tf`, held by `TestKindScopeTFDrift`), so an anchored plan never contacts the cluster.
- The velero charts' snapshot classes and the CRD-only test fixtures drop the namespace they never used.
- Presets: KubernetesDeployment **Agent Substrate Client** and KubernetesNamespace **Agent Substrate Client Namespace**.
