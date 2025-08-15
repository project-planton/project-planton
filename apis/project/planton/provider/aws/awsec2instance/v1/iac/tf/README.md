# Terraform Module to Deploy AwsEc2Instance

This module provisions a single EC2 instance on AWS using ProjectPlanton's tofu (Terraform) integration.

## CLI usage (local backend)

```shell
project-planton tofu init --manifest hack/manifest.yaml
project-planton tofu plan --manifest hack/manifest.yaml
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

Notes:
- Provider credentials are supplied via stack input (through the CLI), not inside the manifest `spec`.
- The `hack/manifest.yaml` is created by the forge rule 008.


