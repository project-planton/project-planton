# Overview

The **AwsEcsService** API resource provides a standardized and straightforward way to deploy containerized applications
onto an existing Amazon ECS cluster on AWS. By focusing on essential configurations like image definition, compute
capacity (Fargate or EC2), networking, and environment variables, it makes running services on ECS far more accessible
within the ProjectPlanton multi-cloud deployment framework.

## Purpose

Deploying ECS services typically involves handling multiple moving parts—task definitions, networking, autoscaling,
IAM roles, and more. The **AwsEcsService** resource aims to streamline that process by:

- **Simplifying ECS Deployments**: Offer an easy-to-use interface for spinning up microservices on ECS (Fargate or EC2).
- **Aligning with Best Practices**: Provide recommended defaults (e.g., CPU, memory) to ensure users have a
  production-ready
  baseline without repetitive configuration.
- **Promoting Consistency**: Enforce standardized naming and validations, reducing misconfigurations across
  multiple services and environments.

## Key Features

### Single-Container Focus

- **Minimal, Opinionated Spec**: Focuses on the 80-20 use case—a single-container service—while still exposing fields
  for
  resource requirements, environment variables, and networking.

### Flexible Compute Options

- **Fargate or EC2**: Operate serverless via AWS Fargate, or integrate with your existing EC2-backed ECS environment.
- **Resource Control**: Define CPU and memory precisely, aligned with ECS constraints (e.g., 256, 512, 1024 CPU units).

### Automatic Networking Setup

- **Subnets & Security Groups**: Attach your service to specific VPC subnets, choosing whether to assign a public IP.
- **Public or Private**: Easily configure production deployments in private subnets, or set up a publicly accessible
  service when needed.

### Environment Management

- **Environment Variables**: Pass configuration to your container, including references to secrets from AWS Secrets
  Manager or SSM.
- **Role Separation**: Separate `task_execution_role_arn` (for pulling container images and writing logs) from
  `task_role_arn` (for runtime AWS API access).

### Seamless Integration

- **ProjectPlanton CLI**: Deploy the same resource across multiple stacks using either Pulumi or Terraform under the
  hood.
- **Multi-Cloud Ready**: Combine AwsEcsService on AWS with other providers in the same manifest, adopting ProjectPlanton’s
  uniform resource model.

## Benefits

- **Reduced Complexity**: A single definition for your ECS service—container image, CPU/memory, subnets, and more—means
  fewer files and less overhead.
- **Scalable & Available**: Scale out by adjusting `desired_count` to meet traffic demands without repeatedly editing
  multiple YAML or JSON templates.
- **Infrastructure Consistency**: Enforce naming conventions, validations, and recommended defaults for CPU/memory
  allocations so your deployments remain predictable and repeatable.
- **Enhanced Observability**: Integrate seamlessly with ECS cluster features like CloudWatch metrics and logs—no extra
  manual setup needed.

## Example Usage

Below is a minimal YAML snippet demonstrating how to configure and deploy an ECS service using ProjectPlanton:

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsService
metadata:
  name: my-aws-ecs-service
  version:
    message: "Initial ECS service deployment"
spec:
  cluster_name: "arn:aws:ecs:us-east-1:123456789012:cluster/my-mixed-cluster"
  service_name: "my-service"
  image: "123456789012.dkr.ecr.us-east-1.amazonaws.com/myapp:latest"
  container_port: 80
  desired_count: 2
  cpu: 512
  memory: 1024
  subnets:
    - subnet-1abc234d
    - subnet-2abc345e
  security_groups:
    - sg-111aaabbb
  assign_public_ip: false
  environment:
    - name: REDIS_URL
      value: "redis://my-redis-cache:6379"
```

### Deploying with ProjectPlanton

Once your YAML manifest is ready, you can deploy using ProjectPlanton’s CLI. ProjectPlanton will validate the manifest
against the Protobuf schema and orchestrate everything in Pulumi or Terraform.

- **Using Pulumi**:
  ```bash
  project-planton pulumi up --manifest awsecsservice.yaml --stack org/project/my-stack
  ```
- **Using Terraform**:
  ```bash
  project-planton terraform apply --manifest awsecsservice.yaml --stack org/project/my-stack
  ```

ProjectPlanton will provision the ECS service, create or update the necessary IAM roles (if specified), assign the
service to the given subnets and security groups, and ensure you have the correct number of running tasks.

---

Happy deploying! If you have questions or run into issues, feel free to open an issue on our GitHub repository or
reach out through our community channels for support.
