# KubernetesAgentSubstrateWorkerPool Pulumi Module

This Pulumi module creates one namespaced Agent Substrate `WorkerPool` on a target cluster.

It builds the object from the manifest's projection (`pkg/iac/pulumi/pulumimodule/provider/kubernetes/manifestcr`), the same projection the generated Terraform module applies. The size, the worker image, the worker pod template and the sandbox class reach the custom resource under their upstream keys, `namespace` becomes `metadata.namespace`, and `labels` and `annotations` become the object's own metadata under Planton's identity labels. The module maps no field by hand, so every field the spec gains reaches the object with no change here.

## Prerequisites

- Agent Substrate on the cluster (see the `KubernetesAgentSubstrate` kind). The object applies with only the definition present; Substrate's controller is what runs the workers, once the pool's class resolves to a SandboxConfig.
- The target namespace (see `KubernetesNamespace`), enforcing no Pod Security level.
- Go toolchain and the Pulumi CLI, and access to the target cluster.

## Usage

### With the Planton CLI

```bash
planton pulumi up --manifest ../../e2e/manifest.yaml
```

### Direct Pulumi usage

The entrypoint loads the `KubernetesAgentSubstrateWorkerPoolIacInput` from the `IAC_INPUT_YAML_FILE` environment variable (path to the IaC input, with the manifest under `target`) or `IAC_INPUT_YAML` (inline YAML content). The CLI builds that input from a manifest and runs Pulumi:

```bash
planton pulumi up --manifest ../../e2e/manifest.yaml --stack <org>/<project>/<stack>
```

## Outputs

| Output | Description |
|--------|-------------|
| `worker_pool_name` | Name of the created WorkerPool (equals `metadata.name`) |
| `namespace` | Namespace the WorkerPool and its workers run in |

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
