# A Chart Names No Owner, and What Planton Provides Is "Platform"

## Why

A chart definition said who owned it twice: once in `spec.selector` (`kind: platform` or `kind: organization`) and again in the record's `metadata.org` once the platform stored it. When the two disagreed, the platform believed the wrong one. Ownership now has one home on the platform, the record's organization, so a chart definition carries none. A chart belongs to the organization that publishes it. The platform charts are the ones with no organization, and they come only from this repository's release.

The same idea runs through the vocabulary. The product said both "official" and "platform" for what Planton provides; it now says "platform" everywhere.

## Charts

- **No `selector` in any `Chart.yaml`:** the 18 charts under `charts/`, the `pkg/infrachart` fixtures and tests, the chart rule's template, and the skill's chart format, worked example and build contract.
- **No hand publishing.** The `charts/Makefile` `publish` target and the publish rule are gone. Every release bundles `charts/` as `infra-charts.zip`, and each control plane, or desktop daemon, seeds the bundle its release pins.

## The engine an installation starts on

- **The operator's `spec.bootstrap.iacProvisioner` is gone,** with its CRD field, sample line, the control plane's `PLANTON_BOOTSTRAP_RUNNER_PROVISIONER` setting and its tests. A self-hosted installation starts its organization on Planton's default engine, like every other kind of installation.
- **`KubernetesPlantonPlatform` has no `bootstrap.iac_provisioner`.** `secret_backend` is field 4 now. The Pulumi and OpenTofu modules, the reference page, the e2e manifest and the spec tests follow.

## "Platform", not "official"

- **The platform modules:** the headers of every `controls.yaml` and `iac/permissions.yaml`, the control, cost and permissions profile protos, `pkg/catalogbundle`, `pkg/iac`, `tools/modulebuild`, the release README and workflows, `planton module eject` (help, notes and its success line, "Ejected the platform … module"), the module-verify suites, and the skills (`org + platform`, "Platform presets", the platform chart fleet).
- **Planton's own distributions:** the runner and operator kinds install Planton's `planton-runner` and `planton-operator` charts, and the operator's images default to Planton's published images.
- **Kept as other companies' word:** a vendor's official Helm chart, image or icon, and the logo provenance grammar (`Official <Provider> product icon`).
- **The vocabulary guard enforces it.** `pkg/vocabulary` retires the old word in front of Planton's modules, presets, chart fleet and catalog. The three dated changelog entries written before the change keep their words.

## Gates

- `go test` passes for `pkg/infrachart`, `pkg/infrachart/validatecmd`, the `KubernetesPlantonPlatform` Pulumi module, `pkg/catalogbundle`, `pkg/iac/eject`, `pkg/explain/refgen` (reference drift), `pkg/vocabulary`, and the operator's `internal/component` and `internal/resources` (boot contract included).
- `KubernetesPlantonPlatform`'s own spec suite passes 97 of 97. Its sizing test still reports the runner's memory default (512Mi) against the operator's 2Gi, a mismatch already on `main` and not touched here.
