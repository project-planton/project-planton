# Overview

The **AwsEcsCluster** API resource provides a standardized, streamlined way to deploy and manage an Amazon ECS cluster on
AWS. It encapsulates key configurations—such as Fargate support, capacity providers, and ECS Exec options—into a single,
uniform resource that seamlessly integrates with the ProjectPlanton multi-cloud deployment framework.

## Purpose

Deploying ECS clusters on AWS can involve numerous configurations, from defining capacity providers to enabling
container insights for production monitoring. The **AwsEcsCluster** resource aims to:

- **Simplify Cluster Provisioning**: Provide an easy-to-use interface for creating ECS clusters tailored to Fargate or
  EC2-based workloads.
- **Enable Production-Ready Monitoring**: Optionally enable CloudWatch Container Insights to gain observability into
  containerized workloads.
- **Streamline Debugging**: Allow for ECS Exec configuration so teams can `exec` into running containers for
  troubleshooting and operational tasks.
- **Encourage Standards**: Enforce consistent naming and recommended defaults to maintain a uniform, predictable setup
  across multiple deployments and environments.

## Key Features

### Container Insights

- **CloudWatch Container Insights**: Easily enable container-level metrics, logs, and traces. Helps teams quickly
  diagnose issues and optimize resource usage.

### Capacity Providers

- **Fargate and Fargate Spot**: Effortlessly configure a Fargate-only cluster or mix in Spot capacity. This allows cost
  optimization while retaining a serverless container environment.
- **EC2 Capacity Providers**: For workloads requiring direct EC2 control, the AwsEcsCluster resource can also integrate
  with existing auto-scaling groups, offering maximum flexibility.

### ECS Exec Support

- **Enable Execute Command**: Optionally allow shell access into running ECS tasks for live debugging. This feature is
  extremely useful during development or in production for resolving critical issues.

### Seamless Integration

- **ProjectPlanton CLI**: Manifests validated and deployed using the same familiar CLI, ensuring a consistent developer
  experience.
- **Multi-Cloud Ready**: Combine AWS ECS with other providers or resources in the same workflow, using the uniform
  ProjectPlanton resource model.

## Benefits

- **Consistent Deployments**: A single resource definition for ECS clusters reduces the likelihood of misconfiguration
  or fragmentation across environments.
- **Improved Observability**: Built-in support for enabling CloudWatch Container Insights ensures teams can monitor and
  troubleshoot services without complex manual setup.
- **Cost Optimization**: Leverage capacity providers and advanced scheduling options (e.g., Spot, Fargate) to balance
  performance and budget constraints.
- **Future-Proof**: As AWS evolves ECS capabilities, updates to the AwsEcsCluster resource keep your deployments aligned
  with best practices and new features.

## Example Usage

Below is a minimal YAML snippet demonstrating how to configure and deploy an ECS cluster using ProjectPlanton:

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEcsCluster
metadata:
  name: my-aws-ecs-cluster
  version:
    message: "Initial ECS cluster deployment"
spec:
  enable_container_insights: true
  capacity_providers:
    - "FARGATE"
    - "FARGATE_SPOT"
  enable_execute_command: true
```

You would then deploy this using the ProjectPlanton CLI (either Pulumi or Terraform under the hood):

```bash
project-planton pulumi up --manifest awsecscluster.yaml --stack org/project/stack
```

Or:

```bash
project-planton terraform apply --manifest awsecscluster.yaml --stack org/project/stack
```

This deploys an ECS cluster with Container Insights enabled, Spot capacity, and support for ECS Exec in your chosen AWS
region.

---

Happy deploying! If you have any questions or issues, feel free to open an issue in our GitHub repository or reach out
via the ProjectPlanton community channels.
