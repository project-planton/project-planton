# Terraform Module to Deploy AwsRdsCluster

This module provisions an AWS RDS Cluster (Aurora MySQL/PostgreSQL or Multi-AZ DB Cluster) aligned with the Project Planton API.

## CLI (local backend)

```shell
project-planton tofu init --manifest ../hack/manifest.yaml
project-planton tofu plan --manifest ../hack/manifest.yaml
project-planton tofu apply --manifest ../hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest ../hack/manifest.yaml --auto-approve
```

Credentials are passed via the stack input through the CLI, not in `spec`.

## Files
- `variables.tf` (generated; do not edit)
- `provider.tf` — provider setup
- `locals.tf` — computed locals and flags
- `security_group.tf` — optional managed SG and rules
- `subnet_group.tf` — DB subnet group when subnet IDs provided
- `cluster_param_group.tf` — optional cluster parameter group
- `rds_cluster.tf` — main cluster resource
- `outputs.tf` — outputs matching `AwsRdsClusterStackOutputs`

## Examples
See `../../examples.md` for example manifests.


