# Production

A shape to run agents on for real, on GKE.
- **Dedicated nodes:** the control plane runs on nodes labelled and tainted `ate.dev/workloadType=ate-control-plane`, so a surge of agents never starves it. Upstream's production posture.
- **Authorization:** every API call is checked against Substrate's access policies. Your orchestrator's service account is the first owner and grants the rest.
- **Scale:** three API servers, and two router and egress pods.
- **The node agent** runs only on the nodes your worker pools use.
- **Telemetry** goes to your OpenTelemetry collector.

Change first:
- Label and taint the control-plane nodes before you deploy.
- `bootstrapOwners`: your product's service account as its API token names it, `system:serviceaccount:<namespace>:<name>`.
- `allowedClientSpiffeIds`: the same service account as its pod certificate names it, `spiffe://cluster.local/ns/<namespace>/sa/<name>`.
- `nodeAgent.nodeSelector`: the label your worker node pool carries.
- `otlpEndpoint`: your collector, or remove it to keep GKE's managed one.
- `clusterSize: size10` for clusters running thousands of agents; it raises the pod-certificate controller's API rate.

See [04-production.yaml](./04-production.yaml) for the manifest.
