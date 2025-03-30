# AWS ECS Service Example

Below is a sample ECS service manifest that runs a container on AWS Fargate. This configuration sets up a public-facing
service with an Application Load Balancer.

```yaml
apiVersion: aws.project-planton.org/v1
kind: EcsService
metadata:
  name: my-ecs-service
  version:
    # recommended to track your service versions
    message: "v1.0"
spec:
  launch_type: FARGATE
  cluster_name: my-ecs-cluster
  container_image: "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-webapp:latest"
  container_port: 8080
  cpu: 512
  memory: 1024
  desired_count: 2
  environment_variables:
    - key: LOG_LEVEL
      value: debug
  secret_variables:
    - key: DB_PASSWORD
      value: arn:aws:ssm:us-east-1:123456789012:parameter/my-secret-pass
  network:
    vpc_id: "vpc-0123456789abcdef0"
    subnet_ids:
      - "subnet-0123456789abcdef0"
      - "subnet-abcdef0123456789"
    security_group_ids:
      - "sg-0123456789abcdef0"
    assign_public_ip: true
  ingress:
    is_public: true
    domain_name: "myapp.example.com"
    path: "/"
    health_check_path: "/health"
  auto_scaling:
    is_enabled: true
    min_count: 2
    max_count: 5
```

## Usage

1. **Validate the Manifest** (optional but recommended):
   ```bash
   project-planton validate --manifest ./ecsservice.yaml
   ```

2. **Deploy via Pulumi**:
   ```bash
   project-planton pulumi up --manifest ./ecsservice.yaml --stack org/project/stack
   ```
   or **Deploy via Terraform**:
   ```bash
   project-planton terraform apply --manifest ./ecsservice.yaml --stack org/project/stack
   ```

3. **Verify**:
    - Check ECS in the AWS Console or use AWS CLI:
      ```bash
      aws ecs describe-services --cluster my-ecs-cluster --services my-ecs-service
      ```
    - If the service is internet-facing and your DNS is properly configured, you should be able to access the domain (
      `myapp.example.com`) once AWS finishes provisioning the load balancer and DNS.

This example demonstrates a typical public-facing ECS service on AWS Fargate. Adjust `cpu`, `memory`, subnets, security
groups, and domain names to suit your application’s requirements.
