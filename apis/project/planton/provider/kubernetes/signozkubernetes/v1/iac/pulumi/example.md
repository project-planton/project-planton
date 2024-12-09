# SignozKubernetes API-Resource Examples

Below are two examples demonstrating how to configure and deploy the `SignozKubernetes` API resource using various specifications. Follow the instructions to create and apply each YAML configuration using the Planton CLI.

---

## Example 1: Basic Signoz Deployment

### Description

This example demonstrates a basic deployment of Signoz within a Kubernetes cluster. It sets up the foundational infrastructure required for Signoz without additional configurations such as environment variables or secrets. This is suitable for initial deployments or testing purposes.

### Create and Apply

1. **Create a YAML file** using the example below.
2. **Apply the configuration** using the following command:

    ```shell
    planton apply -f <yaml-path>
    ```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: SignozKubernetes
metadata:
  name: basic-signoz
spec:
  kubernetes_cluster_credential_id: my-cluster-credential
```

---

## Example 2: Signoz Deployment with Enhanced Configuration

### Description

This example illustrates how to deploy Signoz with additional stack job settings. Although the `SignozKubernetesSpec` is minimal, this configuration includes stack job settings to customize the deployment process, such as specifying deployment parameters or environment configurations managed by Planton Cloud.

### Create and Apply

1. **Create a YAML file** using the example below.
2. **Apply the configuration** using the following command:

    ```shell
    planton apply -f <yaml-path>
    ```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: SignozKubernetes
metadata:
  name: enhanced-signoz
spec:
  kubernetes_cluster_credential_id: my-cluster-credential
  stack_job_settings:
    retries: 3
    timeout: 600
```

*Note: Replace `<yaml-path>` with the actual path to your YAML configuration file when applying the configurations.*
