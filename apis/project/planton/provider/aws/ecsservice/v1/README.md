# AWS ECS Service (EcsService)

## Introduction

The **EcsService** component manages containerized workloads on AWS ECS using a single manifest. Whether you choose a
serverless Fargate launch type or EC2-based tasks, this resource standardizes configuration with a consistent,
Protobuf-based schema. You can validate and deploy your ECS workloads using the Project Planton CLI, which wraps Pulumi
or Terraform to simplify the entire process.

---

## Resource Definition

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: ...
spec:
  launch_type: ...
  cluster_name: ...
  container_image: ...
  container_port: ...
  cpu: ...
  memory: ...
  desired_count: ...
  environment_variables: ...
  secret_variables: ...
  network:
    vpc_id: ...
    subnet_ids: ...
    security_group_ids: ...
    assign_public_ip: ...
  ingress:
    is_public: ...
    domain_name: ...
    path: ...
    health_check_path: ...
  auto_scaling:
    is_enabled: ...
    min_count: ...
    max_count: ...
```

### Fields Overview

| Field                               | Description                                                                                     | Required | Default   |
|-------------------------------------|-------------------------------------------------------------------------------------------------|----------|-----------|
| **apiVersion**                      | Must be `aws.project-planton.org/v1`.                                                           | Yes      | -         |
| **kind**                            | Must be `EcsService`.                                                                           | Yes      | -         |
| **metadata.name**                   | Logical name for this resource.                                                                 | Yes      | -         |
| **spec.launch_type**                | ECS launch type: `FARGATE` or `EC2`.                                                            | No       | `FARGATE` |
| **spec.cluster_name**               | ECS cluster name or ARN. If empty, a default cluster may be inferred.                           | No       | -         |
| **spec.container_image**            | Docker image to run, such as `nginx:latest` or an ECR image.                                    | Yes      | -         |
| **spec.container_port**             | Port the container listens on (e.g., `8080`).                                                   | Yes      | -         |
| **spec.cpu**                        | CPU units for each task.                                                                        | No       | `256`     |
| **spec.memory**                     | Memory in MiB for each task.                                                                    | No       | `512`     |
| **spec.desired_count**              | Number of tasks to run.                                                                         | No       | `1`       |
| **spec.environment_variables**      | List of key-value pairs for non-sensitive environment variables.                                | No       | []        |
| **spec.secret_variables**           | List of key-value pairs referencing AWS Secrets Manager or SSM.                                 | No       | []        |
| **spec.network.vpc_id**             | VPC ID for tasks. If omitted, a default might be used.                                          | No       | -         |
| **spec.network.subnet_ids**         | List of subnet IDs.                                                                             | No       | -         |
| **spec.network.security_group_ids** | Security groups for tasks.                                                                      | No       | -         |
| **spec.network.assign_public_ip**   | Whether to assign a public IP (for Fargate in a public subnet).                                 | No       | `false`   |
| **spec.ingress.is_public**          | Controls if the service is exposed to the internet.                                             | No       | `false`   |
| **spec.ingress.domain_name**        | Domain name if you want a custom FQDN (e.g. `api.example.com`).                                 | No       | -         |
| **spec.ingress.path**               | Route path for the application if using path-based routing.                                     | No       | `/`       |
| **spec.ingress.health_check_path**  | Path for ALB health checks.                                                                     | No       | `/`       |
| **spec.auto_scaling.is_enabled**    | Enables or disables horizontal scaling. If `false`, the service stays at the **desired_count**. | No       | `false`   |
| **spec.auto_scaling.min_count**     | Minimum number of tasks.                                                                        | No       | `1`       |
| **spec.auto_scaling.max_count**     | Maximum number of tasks.                                                                        | No       | `1`       |

---

## How It Works

1. **Define the Manifest**  
   Write a YAML manifest that follows the schema above.
2. **Validate and Deploy**  
   Use `project-planton validate --manifest <file>` to ensure correctness, then either:
    - **Pulumi**: `project-planton pulumi up --manifest <file> --stack <stack_name>`
    - **Terraform**: `project-planton terraform apply --manifest <file> --stack <stack_name>`
3. **Inspect Outputs**  
   On success, the CLI returns relevant ECS service information (e.g., cluster name, load balancer DNS).

---

## Additional Notes

- **Launch Types**: Fargate is the simplest route (no EC2 management); for advanced setups or specialized hardware, you
  can opt for EC2.
- **Secrets**: Securely pass credentials and other sensitive info via `secret_variables`.
- **Auto-Scaling**: If `auto_scaling.is_enabled` is `true`, ensure `min_count` and `max_count` reflect realistic bounds.
- **Networking**: For Fargate with a public IP, you’ll usually pick subnets that have internet access and a suitable
  security group.
