# KubernetesAgentSubstrateSandboxConfig Pulumi Module

This Pulumi module creates one cluster-scoped Agent Substrate `SandboxConfig` on a target cluster.

It builds the object from the manifest's projection (`pkg/iac/pulumi/pulumimodule/provider/kubernetes/manifestcr`), the same projection the generated Terraform module applies. The class, the default version and the versions reach the custom resource under their upstream keys (each architecture's asset files as the bare map upstream defines), `name` becomes `metadata.name`, and `labels` and `annotations` become the object's own metadata under Planton's identity labels. The module maps no field by hand, so every field the spec gains reaches the object with no change here.

## Prerequisites

- Agent Substrate on the cluster (see the `KubernetesAgentSubstrate` kind): the SandboxConfig definition and the `sandboxconfig-assets` admission policy, which accepts or refuses the object at apply.
- Go toolchain and the Pulumi CLI, and access to the target cluster.

## Usage

### With the Planton CLI

```bash
planton pulumi up --manifest ../../e2e/manifest.yaml
```

### Direct Pulumi usage

The entrypoint loads the `KubernetesAgentSubstrateSandboxConfigIacInput` from the `IAC_INPUT_YAML_FILE` environment variable (path to the IaC input, with the manifest under `target`) or `IAC_INPUT_YAML` (inline YAML content). The CLI builds that input from a manifest and runs Pulumi:

```bash
planton pulumi up --manifest ../../e2e/manifest.yaml --stack <org>/<project>/<stack>
```

## Outputs

| Output | Description |
|--------|-------------|
| `sandbox_config_name` | Name of the created SandboxConfig (the spec's `name`) |
| `sandbox_class` | The config's class (`gvisor` or `microvm`) |

## Module Structure

```
pulumi/
├── main.go          # Pulumi entrypoint (loads the IaC input)
├── Pulumi.yaml      # Pulumi project configuration
├── README.md        # This file
└── module/
    ├── main.go      # Applies the projected custom resource
    └── outputs.go   # Output name constants
```
