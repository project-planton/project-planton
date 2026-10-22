# The Agent Substrate control plane is proven live on both engines

## Why

`KubernetesAgentSubstrate` had never been installed on a live cluster. Its eleven lanes on the pod-certificates profile (kind v0.33, Kubernetes 1.37 with the pod-certificate APIs) now pass on both engines. On the way they found two settings combinations the spec accepted but Substrate cannot run, a download that failed on a host's throttling, and three verifier checks that tested the wrong thing.

## What the lanes proved

- The v0.5.0 install from Planton's published tree rolls out every workload, stamps its definitions and its release, and publishes the three signers' trust bundles.
- From the kind's outputs alone, the API over TLS (trusted through the output trust bundle and dialed under the output server name) creates, reads and deletes an atespace for a projected token's subject. It admits the labeled namespace and refuses an unlabeled one, and the router answers.
- **Bundled database:** keep and re-adopt over the kept volume.
- **External database:** an external PostgreSQL through connection strings.
- **Authorization on:** the owner creates while a stranger is denied by its principal.
- **Credential provider:** the Kubernetes Secrets provider runs exactly the declared grants after a grant change.
- **The agentgateway dataplane** serves.
- **Refusals:** each of the six refuses with what it saw, what it means and what to do, and changes nothing.

## What changed

### The spec refuses what Substrate cannot run

- **`size10_sizes_the_bundled_database`:** size10 sizes the bundled database for a dedicated node, its resources and its `postgresql.conf` (`shared_buffers = 35GB`) together. Explicit `database.bundled.resources` with size10 kept the configuration in a smaller container, which would run out of memory under load.
- **`router_ext_proc_covers_parked_requests` and `router_drain_covers_parked_budget`** mirror the router's own startup checks:
  - `ext_proc_max_requests` must be 0 or cover `parked_requests.max` (1024 when unset);
  - `drain_timeout` must cover `parked_requests.budget` (5s when unset).

  The router refused to start, in a crash loop, with values the spec accepted. The comment that called `ext_proc_max_requests: 0` "no limit" now says upstream derives it.

### A throttled download slows the deploy down instead of failing it

- **The defect:** every plan fetches the release's 23 files from `downloads.planton.ai`. A burst of plans met HTTP 429, which failed the deploy with a message blaming the pin.
- **The fix:**
  - Pulumi's `fetchText` retries dropped connections, 429 and 5xx with backoff, honoring `Retry-After`.
  - The Terraform `http` data sources carry the provider's `retry` block, whose policy retries the same answers.
  - Both engines' refusals now say whether the host is limiting requests, failing, or no longer serves the pin.

### The module's offline tests run on a fresh checkout

- **The defect:** the root `.gitignore` ignored every `generated/` directory, so upstream's `manifests/ate-install/generated/role.yaml` was missing from the module's v0.5.0 test snapshot, and five compose tests failed on any clean checkout.
- **The fix:** `generated/` directories under `testdata/` are kept, and the file came from the pinned release, checked against its `SHA256SUMS`.

### The verifiers check the right thing

- **The denied principal:** a denial is matched against the principal exactly as the API's authorizer names it (`user:` and the percent-encoded identity).
- **Refused deploys:** a planted Deployment may remain only while it is still the planted stand-in, every container running the stand-in's pause image. An upgrade refusal's stand-in is an earlier install of the same kind and carries its mark by design. `KubectlDeploymentLeftAsPlanted` replaces the owner-mark check for the controller and Substrate alike.

The profile is `green` with both engines validated, the reference is regenerated, and the CI matrix is regenerated. The GKE lane (`gke-keyless`) runs on a real cluster.
