# Terraform Module to Deploy AwsClientVpn

This Terraform module provisions an `AwsClientVpn` endpoint using the ProjectPlanton CLI (tofu) with the default local backend.

## CLI
```bash
project-planton tofu init --manifest hack/manifest.yaml
project-planton tofu plan --manifest hack/manifest.yaml
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

Credentials are provided via stack input through the CLI, not in the `spec` of the manifest.

See `iac/tf/hack/manifest.yaml` for the minimal example manifest.


