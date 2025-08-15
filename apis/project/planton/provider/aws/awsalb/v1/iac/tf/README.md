# AWS ALB Terraform module

This module provisions an AWS Application Load Balancer and, if SSL is enabled,
adds an HTTP->HTTPS redirect listener and an HTTPS listener. Optional Route53
records are created when DNS is enabled.

Generated `variables.tf` reflects the proto schema for `AwsAlb`.

Usage example (local backend):

```
terraform init
terraform plan -var-file=.terraform/terraform.tfvars
```


