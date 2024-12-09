# StackJobRunnerKubernetes API-Resource Examples

## Example 1: Basic Stack Job Runner

This example demonstrates the most basic configuration for deploying a Stack Job Runner in a Kubernetes cluster. It uses default Kubernetes credentials and provisions the necessary resources to run the stack job.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: StackJobRunnerKubernetes
metadata:
  name: basic-stack-job-runner
spec:
  kubernetesClusterCredentialId: cluster-credential-12345
```

## Example 2: Custom Kubernetes Cluster Credential

In this example, a custom Kubernetes cluster credential is used to set up the Kubernetes provider for running the stack job runner. This configuration allows for more control over the deployment environment.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: StackJobRunnerKubernetes
metadata:
  name: custom-cluster-job-runner
spec:
  kubernetesClusterCredentialId: cluster-credential-98765
```

## Example 3: Stack Job Runner with Multiple Environments

This example demonstrates deploying a Stack Job Runner with multiple environment credentials. It provides flexibility for running the stack job in different environments, such as development or production.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: StackJobRunnerKubernetes
metadata:
  name: multi-env-job-runner
spec:
  kubernetesClusterCredentialId: cluster-credential-112233
```

## Example 4: Empty Spec (Not Fully Implemented)

In this example, the `spec` field is left empty, indicating that this particular module or API-resource is not fully implemented. It serves as a placeholder until more specific configurations are supported.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: StackJobRunnerKubernetes
metadata:
  name: empty-spec-job-runner
spec: {}
```
