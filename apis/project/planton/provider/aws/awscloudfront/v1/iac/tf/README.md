# Terraform Module to Deploy AwsCloudFront

This Terraform module provisions an `AwsCloudFront` distribution using the ProjectPlanton CLI (tofu) with the default local backend.

## CLI
```bash
project-planton tofu init --manifest hack/manifest.yaml
project-planton tofu plan --manifest hack/manifest.yaml
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

Credentials are provided via stack input through the CLI, not in the `spec` of the manifest.

See `iac/hack/manifest.yaml` for the minimal example manifest.


