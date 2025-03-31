# Create using CLI

Create a YAML manifest using one of the examples below. After the YAML is created, apply it with ProjectPlanton:

```shell
project-planton pulumi up --manifest <yaml-path> --stack <stack-name>
```

Or, if using Terraform:

```shell
project-planton terraform apply --manifest <yaml-path> --stack <stack-name>
```

(You can also use the shorter form `planton apply -f <yaml-path>` if your environment is configured accordingly.)

---

# Basic Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: my-basic-ecs-service
  version:
    message: "Initial ECS service"
spec:
  cluster_name: "my-basic-ecs-cluster"
  service_name: "my-service"
  image: "amazonlinux:2"
  container_port: 80
  cpu: 512
  memory: 1024
  subnets:
    - "subnet-1234abcd"
    - "subnet-5678efgh"
```

This defines a simple ECS service that:
• Uses an existing ECS cluster named `my-basic-ecs-cluster`.
• Deploys the Amazon Linux 2 image, listening on port 80.
• Runs on Fargate with 512 CPU units and 1024 MiB memory.
• Spans two subnets for high availability.

---

# Example with Environment Variables

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: my-service-with-env
  version:
    message: "ECS service with environment variables"
spec:
  cluster_name: "my-mixed-cluster"
  service_name: "env-service"
  image: "123456789012.dkr.ecr.us-east-1.amazonaws.com/myapp:latest"
  container_port: 8080
  cpu: 512
  memory: 1024
  subnets:
    - "subnet-1a2b3c4d"
    - "subnet-5e6f7g8h"
  security_groups:
    - "sg-0123456789abcdef0"
  environment:
    - name: "APP_ENV"
      value: "production"
    - name: "LOG_LEVEL"
      value: "INFO"
```

Here, the ECS service is defined with custom environment variables, making it easy to configure runtime settings like
environment or log level. It also references both subnets and a security group for proper network isolation.

---

# Example with Multiple Task Replicas

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: high-availability-service
  version:
    message: "High availability example"
spec:
  cluster_name: "my-mixed-cluster"
  service_name: "ha-service"
  image: "nginx:stable"
  container_port: 80
  desired_count: 3
  cpu: 512
  memory: 1024
  subnets:
    - "subnet-aaaabbbb"
    - "subnet-ccccdddd"
  security_groups:
    - "sg-11112222333344445"
```

In this example, the service runs three task replicas for higher availability. Each task is placed in one of the
specified subnets, and the security group is applied to manage inbound or outbound traffic as required.

---

# Example with Public IP Assignment

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: public-ecs-service
  version:
    message: "Service with public IP"
spec:
  cluster_name: "my-basic-ecs-cluster"
  service_name: "public-service"
  image: "amazonlinux:2"
  container_port: 80
  cpu: 256
  memory: 512
  subnets:
    - "subnet-1111aaaa"
  security_groups:
    - "sg-2222bbbb"
  assign_public_ip: true
```

Setting `assign_public_ip` to `true` grants the Fargate tasks a public IP. This is useful for quick testing or for
services that must be internet-facing without a load balancer. However, for production, it’s often safer to place tasks
in private subnets behind a load balancer.

---

# Example with Custom IAM Roles

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: custom-roles-service
  version:
    message: "Service with custom IAM roles"
spec:
  cluster_name: "my-exec-enabled-cluster"
  service_name: "custom-role-service"
  image: "123456789012.dkr.ecr.us-east-1.amazonaws.com/myapp:latest"
  container_port: 8080
  cpu: 1024
  memory: 2048
  subnets:
    - "subnet-xyz12345"
    - "subnet-abc67890"
  security_groups:
    - "sg-a1b2c3d4"
  task_execution_role_arn: "arn:aws:iam::123456789012:role/ecsTaskExecutionRole"
  task_role_arn: "arn:aws:iam::123456789012:role/myAppTaskRole"
```

Here, the ECS service is granted two distinct IAM roles:
• `task_execution_role_arn` for pulling private images and writing logs.
• `task_role_arn` for application-level AWS API calls within the container.

---

# Minimal Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: minimal-ecs-service
  version:
    message: "Minimal example"
spec:
  cluster_name: "arn:aws:ecs:us-east-1:123456789012:cluster/simple-cluster"
  service_name: "min-service"
  image: "amazonlinux:2"
  cpu: 256
  memory: 512
  subnets:
    - "subnet-11112222"
```

A bare-bones deployment that sets only the required fields:
• ECS cluster reference, service name, image, CPU, memory, and subnets.

---

After choosing one of these examples, apply it with ProjectPlanton using Pulumi or Terraform:

```shell
project-planton pulumi up --manifest ecs-service.yaml --stack myorg/dev
```

or

```shell
project-planton terraform apply --manifest ecs-service.yaml --stack myorg/dev
```

When the command completes successfully, your ECS service will be created. You can confirm by checking the AWS console
or by using the AWS CLI:

```shell
aws ecs list-services --cluster <your-cluster-name-or-arn>
```

This confirms that your ECS service is up and running on the specified ECS cluster, ready to serve traffic or perform
background processing.

---

Happy deploying with ProjectPlanton!
