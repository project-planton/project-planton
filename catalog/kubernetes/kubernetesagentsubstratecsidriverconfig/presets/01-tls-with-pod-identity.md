# TLS With Pod Identity

A CSI driver registered with Substrate over an encrypted, mutually authenticated connection: the API server presents its Substrate pod identity to the driver's Controller service and trusts only a serving certificate from Substrate's own service-DNS signer. The node socket is the kubelet convention.

## When to Use

- The driver's Controller service is reachable from more than the API server, or crosses a network you do not fully trust.
- You want volume operations accepted only from Substrate's API server.

## How It Works

When an actor's external volume names a StorageClass provisioned by `csi.example.com`, the API server finds this config by that name and dials the Controller at the endpoint over TLS, verifying the server name and presenting its pod-identity certificate. The driver's TCP bridge must serve a certificate from Substrate's service-DNS signer: give it a pod-certificate projected volume with signer `servicedns.podcert.ate.dev/identity`. Node agents reach the driver's Node service over `unix:///var/lib/kubelet/plugins/csi.example.com/csi.sock`.

## Key Configuration Choices

- **Name equals driver name.** Substrate finds the config by its object name.
- **Pod identity is the only TLS mode** at this release; certificates from Secrets are not supported.
- **The server name** is the name the bridge's serving certificate carries, `<service>.<namespace>.svc`.

## Prerequisites

- **Agent Substrate** on the cluster (`KubernetesAgentSubstrate`).
- The CSI driver running, with a TCP bridge in front of its Controller socket serving a certificate from Substrate's service-DNS signer, and its node plugin bind-mounting `/var/lib/ate` with bidirectional propagation.
- A StorageClass whose provisioner is the driver's name (`KubernetesStorageClass`).

## Placeholders to Replace

| Placeholder | Description |
|-------------|-------------|
| `<config-resource-name>` | Name of this Planton resource. |
| `csi.example.com` (name and driver name) | The driver's CSI name, in both fields; it is what its StorageClasses name as provisioner. |
| `<controller-service>` | The Service in front of the driver's Controller (its TCP bridge). |
| `<controller-namespace>` | The namespace of that Service. |
