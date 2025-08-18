## Terraform Module to Deploy AwsDynamodb

Run the module via the ProjectPlanton CLI (tofu) using the default local backend.

```shell
project-planton tofu init --manifest hack/manifest.yaml
project-planton tofu plan --manifest hack/manifest.yaml
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

- Credentials are provided via stack input (by the CLI), not in the manifest `spec`.
- Manifest file: `../hack/manifest.yaml`


