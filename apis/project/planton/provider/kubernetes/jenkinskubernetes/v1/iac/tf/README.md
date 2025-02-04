# Terraform Module to Deploy Jenkins on Kubernetes

```shell
project-planton tofu init --manifest hack/manifest.yaml --backend-type s3 \
  --backend-config="bucket=planton-cloud-tf-state-backend" \
  --backend-config="dynamodb_table=planton-cloud-tf-state-backend-lock" \
  --backend-config="region=ap-south-2" \
  --backend-config="key=kubernetes-stacks/test-jenkins-server.tfstate"
```

```shell
project-planton tofu plan --manifest hack/manifest.yaml
```

```shell
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
```

```shell
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```
