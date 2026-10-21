# KubernetesAgentSubstrate Pulumi Module

Installs the Agent Substrate v0.5.0 control plane from Planton's digest-pinned build of upstream's release (`https://downloads.planton.ai/releases/agent-substrate/v0.5.0/`), composed the way upstream's installer composes it, with the typed spec applied onto upstream's own objects. The exact semantic twin of the Terraform module (`../tf/`).

## Layout

| File | Role |
|---|---|
| `module/vars.go` | The pinned release, upstream's fixed names, the file addresses |
| `module/settings.go` | The spec resolved to settings with upstream's defaults filled in |
| `module/documents.go` | Download, split on separator lines, key by composed identity |
| `module/patch.go` | Upstream's kustomize components and overlays, applied as written |
| `module/compose.go` | The document set, in four phases, with the passes every document gets |
| `module/workloads.go` | The typed settings on each workload |
| `module/splice.go` | Upstream's text markers, the agentgateway and Cloud SQL components, the image mirror |
| `module/config.go` | The spec-derived ConfigMaps and Secret, and the rollout digests |
| `module/pools.go`, `module/secrets.go` | The generated certificate authorities and signing key, as typed Secrets |
| `module/upgrade.go`, `module/preflight.go` | Upstream's upgrade policy and every refusal |
| `module/main.go` | Refusals, the kept definitions, the Secrets, then the phases |

## Order and Refusals

Everything that can refuse runs before anything registers: a namespace other than upstream's; a workload or namespace another installation holds (`workloadowner`); a minor-release move upstream cannot survive, read from the stamps on the API server and the bundled database's volume; and, in `keptcrds`, a definition someone else owns or a newer one. Then the namespaces, the definitions, the Secrets, and the prerequisites, database, system and runtime phases, each awaited before the next.

## Tests

- `compose_test.go` — hermetic, against `testdata/v0.5.0` (the published files): an empty section installs upstream as published; typed settings land where upstream's installer puts them; upstream's components and markers apply as upstream applies them; the rollout digests cover what each workload reads; the upgrade rule follows upstream's policy; the pools have upstream's wire shape; the Terraform twin pins the same release. Gated by `AGENT_SUBSTRATE_UPSTREAM_PIN=1`: the fixtures still equal what is published.
- `pools_upstream_test.go` — gated by `AGENT_SUBSTRATE_UPSTREAM_PARSE=1`: both engines' generated pools parse with upstream's own `localca` and `localjwtauthority` code.
- `parity_test.go` — gated by `AGENT_SUBSTRATE_TF_PARITY=1` with a kubeconfig: plans the Terraform twin and compares its documents and outputs with this module's composition, across five specs.

## Usage

```bash
planton pulumi up --manifest agent-substrate.yaml
```

## Outputs

See `module/outputs.go`; every output is described in `../../v1alpha1/outputs.proto`.
