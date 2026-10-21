# The agent-sandbox controller is proven live, its router reaches the API on every cluster, and a reference reads what its target declares

## Why

`KubernetesAgentSandboxController` was built and checked offline, but no live cluster had run its lanes. Running all of them on kind, on both engines, found three defects a user would meet and one verifier that proved less than its scenario claimed.

## What changed

### The router reaches the cluster API behind its own NetworkPolicy

- **The defect:** upstream's router NetworkPolicy (on by default) allows the router's API traffic, its pod cache and TokenReview, only toward pods in the `default` namespace. On Cilium (GKE Dataplane V2, most EKS clusters), and wherever the API server sits on a node or behind a load balancer, nothing matches. The router's initial pod-cache sync never completes, and the router crash-loops.
- **The fix:** both engines now keep upstream's ports (443, 6443) on that rule and drop its peer. A NetworkPolicy cannot name the API server, and the router reaches its upstreams over plain HTTP only, so a TLS endpoint on those ports refuses it. Both engines' shape guards also check that this rule is where they expect it.
- **The proof:** a new lane, `router-enforcing-cni`, runs the router behind its default policy on the `cilium-cni` profile and proves a warm claim, a tokenless 401 and a ServiceAccount-token 200. Before the fix it crash-looped on every run; with only the policy removed it served; with the fix it passes on both engines. The cross-engine parity test (`TestEnginesApplyTheSameDocuments`) passes.

### A reference reads the fields its target declares, not only its outputs

- **The defect:** `planton apply -f <dir>` (set deploy) and the e2e runner resolved a `valueFrom` only against the target's captured outputs. The control plane resolves every path against the target's stored resource (spec, metadata and outputs), so a reference to a namespace's `spec.name`, the default path of every namespaced kind's `namespace` field, resolved on the platform and failed everywhere else.
- **The fix:**
  - `refcheck.CanonicalValueFromPath` puts a path in proto field names with bare indexes. It shares its walk with `ResolveValueFromPath`, mirroring the control plane's canonicalizer.
  - `manifestgraph.ManifestFields` and `LookupKey` give a reference the referenced manifest's spec and metadata beside its outputs.
  - `setdeploy` and the e2e runner (`runner.ReferableValues`, `DependencyState.Referable`, the mapping-eval deployer) supply both.
  - A missing manifest field is reported as not set in its manifest, not as a missing output.

### The verifiers check what their scenarios promise

- The controller's probe claims only once its warm pool holds a ready member, then asserts the claim bound that member (`agents.x-k8s.io/launch-type: warm`) with the declared pod label. The claim verifier's binding check is now shared (`proveClaimBinding`).
- A Sandbox asking for `service: true` must have a headless Service whose endpoints are its pod.
- Every declared volume claim template must exist at its declared size and be mounted in the sandbox's pod. For a claim, the size is what proves its own template replaced the pool template's.
- An unmanaged SandboxTemplate is judged only after the controller has reconciled it: its `agents.x-k8s.io/sandbox-template-ref-hash` label is the evidence, not a fixed wait.
- After a refused deploy, a Deployment of that name may remain only when it is another owner's, without this kind's mark (`KubectlResourceNotOwnedBy`). That is what an owned-elsewhere lane plants. Both the controller and the Agent Substrate control-plane verifiers use it.

## Proof

All nine controller lanes pass on both engines on kind v0.33 (Kubernetes 1.37): `minimal`, `full-surface`, `sdk-ready`, `upgrade` and the four refusals on the default cluster, plus `router-enforcing-cni` on the Cilium profile. The profile is `green` with both engines validated, and the CI matrix is regenerated.
