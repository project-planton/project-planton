# In-Cluster Plaintext Driver

A CSI driver registered with Substrate over plaintext gRPC at its Service's cluster DNS name, with a node socket in the driver's own plugin directory: the shape upstream registers its own NFS driver with (`tcp://csi-nfs-controller.kube-system.svc.cluster.local:50052`).

## When to Use

- The driver's Controller service is reachable only from inside the cluster (a NetworkPolicy admits only the control plane's namespace).
- The driver registers its node plugin under a directory that is not its CSI name.

## How It Works

When an actor's external volume names a StorageClass provisioned by `csi.example.com`, the API server finds this config by that name and dials the Controller through cluster DNS. Node agents reach the driver's Node service over the socket in `nodeSocketOverride` instead of `unix:///var/lib/kubelet/plugins/csi.example.com/csi.sock`.

## Key Configuration Choices

- **No TLS.** Restrict the Controller Service to the control plane's namespace with a NetworkPolicy; switch to the **TLS With Pod Identity** preset otherwise.
- **The node socket override** is for drivers whose plugin directory differs from their name (`csi-nfsplugin` for `nfs.csi.k8s.io`); drop it when the directory is the driver's name.

## Prerequisites

- **Agent Substrate** on the cluster (`KubernetesAgentSubstrate`).
- The CSI driver running, with a TCP bridge in front of its Controller socket, and its node plugin bind-mounting `/var/lib/ate` with bidirectional propagation.
- A StorageClass whose provisioner is the driver's name (`KubernetesStorageClass`).

## Placeholders to Replace

| Placeholder | Description |
|-------------|-------------|
| `<config-resource-name>` | Name of this Planton resource. |
| `csi.example.com` (name and driver name) | The driver's CSI name, in both fields. |
| `<controller-service>` | The Service in front of the driver's Controller. |
| `<controller-namespace>` | The namespace of that Service. |
| `<plugin-directory>` | The directory the driver's node plugin registers its socket under. |
