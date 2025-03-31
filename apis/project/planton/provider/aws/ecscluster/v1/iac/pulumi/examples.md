# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS VPC with default settings.

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsCluster
metadata:
  name: dev-cluster
spec:
  enableContainerInsights: true
  capacityProviders:
    - FARGATE
    - FARGATE_SPOT
  enableExecuteCommand: false
```

```shell
project-planton pulumi up --stack planton-cloud-state-backend/project-planton-prod/ecs-cluster.dev-cluster --manifest hack/dev-cluster.yaml
```
