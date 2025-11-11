```md
# AWS ALB Examples

Below are several examples demonstrating how to define an AWS ALB (Application Load Balancer) component in
ProjectPlanton. After creating one of these YAML manifests, apply it with your preferred IaC engine (Pulumi or
Terraform) using the ProjectPlanton CLI:

```shell
project-planton pulumi up --manifest <yaml-path> --stack <stack-name>
```

or

```shell
project-planton terraform apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic Internet-Facing ALB

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: basic-alb
spec:
  subnets:
    - subnet-12345abc
    - subnet-67890def
  securityGroups:
    - sg-abcdef1234567890
  listeners:
    - protocol: "HTTP"
      port: 80
```

This example:
• Creates an ALB in two subnets (`subnet-12345abc` and `subnet-67890def`).
• Attaches one security group (`sg-abcdef1234567890`).
• Sets up an HTTP listener on port 80.

---

## ALB with HTTPS

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: alb-with-https
spec:
  subnets:
    - subnet-1111aaaa
    - subnet-2222bbbb
  securityGroups:
    - sg-zzzzxxxxccccdddd
  listeners:
    - protocol: "HTTPS"
      port: 443
      certificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-efgh-ijkl-123456abcdef"
      sslPolicy: "ELBSecurityPolicy-2016-08"
```

Here:
• The listener uses HTTPS on port 443.
• A certificate ARN and SSL policy are specified for TLS termination.

---

## Internal ALB

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: internal-alb
spec:
  subnets:
    - subnet-internal1
    - subnet-internal2
  securityGroups:
    - sg-internal1234
  scheme: "internal"
  listeners:
    - protocol: "HTTP"
      port: 80
```

This ALB is internal-facing:
• `scheme` is set to `internal` for private/internal traffic.
• It can serve traffic only accessible within the specified subnets or VPC.

---

## ALB with Multiple Listeners

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: alb-multi-listener
spec:
  subnets:
    - subnet-abc12345
    - subnet-def67890
  securityGroups:
    - sg-1234abcd5678efgh
  listeners:
    - protocol: "HTTP"
      port: 80
    - protocol: "HTTPS"
      port: 443
      certificateArn: "arn:aws:acm:us-east-1:987654321098:certificate/fedcba98-7654-3210-fedc-ba9876543210"
      sslPolicy: "ELBSecurityPolicy-2015-05"
```

In this example:
• Both HTTP (port 80) and HTTPS (port 443) listeners are defined on the same ALB.
• HTTPS listener uses a specific SSL policy and certificate ARN.

---

## ALB with Custom Idle Timeout and Deletion Protection

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: alb-custom-config
spec:
  subnets:
    - subnet-custom1
    - subnet-custom2
  securityGroups:
    - sg-custom
  ipAddressType: "dualstack"
  enableDeletionProtection: true
  idleTimeoutSeconds: 120
  listeners:
    - protocol: "HTTP"
      port: 80
```

This configuration:
• Enables dualstack to support both IPv4 and IPv6.
• Sets a custom idle timeout of 120 seconds.
• Protects the ALB from accidental deletion by enabling deletion protection.

---

## Minimal ALB (No Listeners)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsAlb
metadata:
  name: alb-minimal
spec:
  subnets:
    - subnet-98765432
  securityGroups:
    - sg-5555aaaa9999bbbb
```

A minimal manifest with:
• Required subnets and security groups.
• No listeners defined in `spec`. You can add them later or manage them with a separate resource.

---

## After Deploying

Once you’ve applied your manifest with ProjectPlanton, you can confirm that the ALB is active via the AWS console or by
using the AWS CLI:

```shell
aws elbv2 describe-load-balancers --names <your-load-balancer-name>
```

You should see details about your Application Load Balancer, such as its DNS name and ARN. Feel free to adjust subnets,
security groups, or listeners to fit your needs.

```
