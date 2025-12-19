# Kubernetes Tekton Operator - Pulumi Module

This directory contains the Pulumi implementation for deploying the Tekton Operator on Kubernetes.

## Overview

The Pulumi module installs the Tekton Operator using official release manifests. The operator manages the lifecycle of Tekton components (Pipelines, Triggers, Dashboard) via the TektonConfig CRD.

## Module Structure

```
iac/pulumi/
├── main.go                     # Pulumi program entrypoint
├── Pulumi.yaml                 # Project configuration
├── Makefile                    # Build automation
├── debug.sh                    # Debugging helper script
└── module/
    ├── main.go                 # Module orchestration
    ├── locals.go               # Local variables and label management
    ├── outputs.go              # Stack output constants
    ├── vars.go                 # Tekton operator constants
    └── tekton_operator.go      # Operator installation implementation
```

## Prerequisites

- **Pulumi CLI**: Install from [pulumi.com](https://www.pulumi.com/docs/get-started/install/)
- **Go**: Version 1.21 or later
- **Kubernetes Cluster**: Target cluster with kubectl access

## Quick Start

### 1. Configure Stack

Create a new Pulumi stack and configure it:

```bash
cd iac/pulumi
pulumi stack init dev
pulumi config set-secret kubeconfig "$(cat ~/.kube/config)"
```

### 2. Set Input Variables

Create a `stack-input.json` file:

```json
{
  "metadata": {
    "name": "tekton-operator",
    "id": "tekton-op-dev"
  },
  "spec": {
    "target_cluster": {
      "kubernetes_credential_id": "my-cluster-cred"
    },
    "container": {
      "resources": {
        "requests": {
          "cpu": "100m",
          "memory": "128Mi"
        },
        "limits": {
          "cpu": "500m",
          "memory": "512Mi"
        }
      }
    },
    "components": {
      "pipelines": true,
      "triggers": true,
      "dashboard": true
    }
  }
}
```

### 3. Deploy

```bash
pulumi up
```

### 4. Verify Deployment

```bash
kubectl get pods -n tekton-operator
kubectl get tektonconfig
kubectl get pods -n tekton-pipelines
```

## Module Components

### Tekton Operator Installation

Deploys the Tekton Operator from official release manifests:

```go
yaml.NewConfigFile(ctx, "tekton-operator", &yaml.ConfigFileArgs{
    File: vars.OperatorReleaseURL,
}, pulumi.Provider(k8sProvider))
```

### TektonConfig Configuration

Creates a TektonConfig CRD to configure which components to install:

```go
tektonConfigYAML := `apiVersion: operator.tekton.dev/v1alpha1
kind: TektonConfig
metadata:
  name: config
spec:
  profile: all  # or: basic, lite
  targetNamespace: tekton-pipelines
  pipeline:
    enable-api-fields: stable`
```

### Component Profiles

The module selects a profile based on enabled components:

- **all**: Pipelines + Triggers + Dashboard
- **basic**: Pipelines + Triggers
- **lite**: Pipelines only

## Constants (vars.go)

Tekton Operator configuration constants:

```go
var vars = struct {
    OperatorNamespace    string
    ComponentsNamespace  string
    OperatorReleaseURL   string
    OperatorVersion      string
    TektonConfigName     string
}{
    OperatorNamespace:    "tekton-operator",
    ComponentsNamespace:  "tekton-pipelines",
    OperatorReleaseURL:   "https://storage.googleapis.com/tekton-releases/operator/latest/release.yaml",
    OperatorVersion:      "latest",
    TektonConfigName:     "config",
}
```

## Stack Outputs

The module exports the following outputs:

- **namespace**: Kubernetes namespace where Tekton components run (`tekton-pipelines`)
- **tekton_config_name**: Name of the TektonConfig resource
- **pipelines_controller_service**: Pipelines controller service name (if enabled)
- **triggers_controller_service**: Triggers controller service name (if enabled)
- **dashboard_service**: Dashboard service name (if enabled)
- **dashboard_port_forward_command**: Port-forward command for dashboard access

Access outputs:

```bash
pulumi stack output namespace
pulumi stack output dashboard_port_forward_command
```

## Debugging

Use the included debug script to inspect the stack:

```bash
./debug.sh
```

View Pulumi logs:

```bash
pulumi logs
```

## Cleanup

Remove the Tekton Operator deployment:

```bash
pulumi destroy
```

> **Warning**: This will remove the operator AND all Tekton components it manages.

## Common Issues

### Operator Pod Not Starting

**Symptom**: Tekton operator pod remains in Pending state.

**Solution**: Check resource availability:

```bash
kubectl describe pod -n tekton-operator -l app=tekton-operator
kubectl top nodes
```

### TektonConfig Not Ready

**Symptom**: TektonConfig status shows not ready.

**Solution**: Check operator logs:

```bash
kubectl logs -n tekton-operator -l app=tekton-operator
kubectl describe tektonconfig config
```

### Dashboard Not Accessible

**Symptom**: Dashboard service not reachable.

**Solution**: Use port-forwarding:

```bash
kubectl port-forward svc/tekton-dashboard -n tekton-pipelines 9097:9097
```

Then access: http://localhost:9097

## References

- [Pulumi Kubernetes Provider](https://www.pulumi.com/registry/packages/kubernetes/)
- [Tekton Documentation](https://tekton.dev/docs/)
- [Tekton Operator GitHub](https://github.com/tektoncd/operator)
