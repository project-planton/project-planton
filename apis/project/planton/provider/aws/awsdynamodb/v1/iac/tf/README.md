# Terraform Module to Deploy AwsDynamodb

Use the ProjectPlanton CLI (tofu) to run this Terraform module. The minimal manifest is under `../hack/manifest.yaml`.

```bash
project-planton tofu init --manifest ../hack/manifest.yaml
project-planton tofu plan --manifest ../hack/manifest.yaml
project-planton tofu apply --manifest ../hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest ../hack/manifest.yaml --auto-approve
```

- Credentials are provided via stack input by the CLI, not in the manifest `spec`.
