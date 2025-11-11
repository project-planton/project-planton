# Create using CLI

Create a yaml manifest using one of the examples below. After the YAML is created, use the command below to apply with
ProjectPlanton:

```shell
project-planton pulumi up --manifest <yaml-path> --stack <stack-name>
```

Or, if using Terraform:

```shell
project-planton terraform apply --manifest <yaml-path> --stack <stack-name>
```

(You can also use a shorter form like `planton apply -f <yaml-path>` if your environment is configured accordingly.)

---

# Basic Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsCluster
metadata:
  name: my-basic-aws-ecs-cluster
  version:
    message: "Initial cluster"
spec:
  # Recommended to be true for production, enabling CloudWatch Container Insights
  enable_container_insights: true

  # Using default Fargate capacity
  capacity_providers:
    - "FARGATE"

  # Defaulting to false; set to true if you want ECS Exec
  enable_execute_command: false
```

This example creates a simple AWS ECS cluster that runs on Fargate, with container insights enabled.

---

# Example with Multiple Capacity Providers

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsCluster
metadata:
  name: my-mixed-cluster
  version:
    message: "Multi capacity providers cluster"
spec:
  # Enable CloudWatch insights to capture performance metrics
  enable_container_insights: true

  # Include both FARGATE and FARGATE_SPOT for cost savings when spot capacity is available
  capacity_providers:
    - "FARGATE"
    - "FARGATE_SPOT"

  # Keep ECS Exec disabled if not needed
  enable_execute_command: false
```

This example showcases how to combine FARGATE with FARGATE_SPOT, allowing you to reduce costs by taking advantage of
Spot capacity where possible.

---

# Example with ECS Exec Enabled

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsCluster
metadata:
  name: my-exec-enabled-cluster
  version:
    message: "Exec debugging example"
spec:
  # Container Insights are highly recommended for monitoring
  enable_container_insights: true

  # Single capacity provider for a basic Fargate cluster
  capacity_providers:
    - "FARGATE"

  # Enable ECS Exec to allow debugging into running containers
  enable_execute_command: true
```

When `enable_execute_command` is `true`, you can connect to running tasks for debugging or operational troubleshooting.

---

# Example with Minimal Configuration

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsCluster
metadata:
  name: shorty
  version:
    message: "Minimal cluster example"
spec:
# Accept all defaults; Insights defaults to true
# capacity_providers defaults to an empty list (no cluster capacity providers),
# which is typically updated manually or managed externally if needed
```

If you need a bare-bones cluster definition, you can omit many fields. However, you should explicitly consider capacity
providers and container insights for a more production-ready setup.

---

After creating a YAML manifest for your ECS cluster, apply the configuration using either Pulumi or Terraform with the
ProjectPlanton CLI. The CLI will validate your manifest against the Protobuf schema, generate the required
infrastructure code, and provision the cluster on AWS.

For example:

```shell
project-planton pulumi up --manifest minimal-ecs.yaml --stack myorg/dev
```

Or:

```shell
project-planton terraform apply --manifest minimal-ecs.yaml --stack myorg/dev
```

Upon completion, you can check the newly created ECS cluster in the AWS Console or with the AWS CLI:

```shell
aws ecs list-clusters
```

This confirms that your ECS cluster has been created and is ready for workloads.
