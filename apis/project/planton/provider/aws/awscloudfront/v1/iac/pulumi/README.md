Pulumi Module to Deploy AwsCloudFront

This Pulumi module provisions an AWS CloudFront distribution based on the AwsCloudFront resource manifest. Use the ProjectPlanton CLI to run previews, updates, refreshes, and destroys.

Notes
- Credentials are provided via stack input flags or files; do not embed secrets in the manifest `spec`.
- You can override manifest values with `--set key=value`.
- The manifest path below points to the existing Terraform hack manifest. If you maintain a top-level `iac/hack/manifest.yaml`, adjust the path accordingly.

Preview
```shell
project-planton pulumi preview \
  --manifest ../tf/hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .
```

Update (apply)
```shell
project-planton pulumi update \
  --manifest ../tf/hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir . \
  --yes
```

Refresh
```shell
project-planton pulumi refresh \
  --manifest ../tf/hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .
```

Destroy
```shell
project-planton pulumi destroy \
  --manifest ../tf/hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .
```

Optional credential flags
- `--aws-credential /path/to/aws-credential.yaml`
- `--gcp-credential /path/to/gcp-credential.yaml`
- `--azure-credential /path/to/azure-credential.yaml`
- `--kubernetes-cluster /path/to/cluster.yaml`
- `--confluent-credential /path/to/confluent-credential.yaml`
- `--mongodb-atlas-credential /path/to/mongodb-atlas-credential.yaml`
- `--snowflake-credential /path/to/snowflake-credential.yaml`

Example with overrides
```shell
project-planton pulumi preview \
  --manifest ../tf/hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir . \
  --set spec.origins[0].domain_name=my-bucket.s3.amazonaws.com \
  --set spec.default_cache_behavior.viewer_protocol_policy=REDIRECT_TO_HTTPS
```


