# The four agent-sandbox resources are proven live on both engines

## Why

`KubernetesAgentSandbox`, `KubernetesAgentSandboxTemplate`, `KubernetesAgentSandboxWarmPool` and `KubernetesAgentSandboxClaim` were built and checked offline. Their lanes now ran on kind v0.33 (Kubernetes 1.37), on both engines, behind the controller each installs as its prerequisite.

## What the lanes proved

- **Sandbox:**
  - Every field a full-surface manifest sets reads back from the live object.
  - The workspace claim exists at its declared size and is mounted.
  - `service: true` gives a headless Service resolving to the sandbox's pod.
  - Suspending deletes the pod and keeps the same claim (marked with its UID); resuming brings the workspace file back.
- **Template:**
  - Managed with no rules, the controller applies the secure default (ingress only from the router).
  - Managed with rules, the policy carries exactly the declared rules.
  - Unmanaged, the controller reconciles the template and creates no policy.
- **Warm pool:** it fills to its size, and the scale endpoint an autoscaler drives reports it.
- **Claim:**
  - A claim asking for nothing a running pod cannot take binds a warm member.
  - One setting variables and its own volume starts its own sandbox (cold), with the pod label, the variable and its 2Gi workspace.
  - One its template's policy forbids is refused with `EnvVarsInjectionRejected`.

## Fixed on the way

- **Empty pools:** a warm pool sized 0 is proven by holding no members. The API omits a zero `readyReplicas`, so the fill check could never see "0".

All four profiles are `green` with both engines validated, and the CI matrix is regenerated.
