# Pulumi Module to Deploy AwsEc2Instance

This module provisions a single EC2 instance on AWS using ProjectPlanton's Pulumi integration.

## CLI usage

```shell
# Preview
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .

# Update (apply)
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir . \
  --yes

# Refresh
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .

# Destroy
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack <org>/<project>/<stack> \
  --module-dir .
```

## Debugging

You can debug the Pulumi program with Delve. A `debug.sh` helper is provided. To enable it, uncomment the `runtime.options.binary` line in `Pulumi.yaml`:

```yaml
runtime:
  name: go
  options:
    binary: ./debug.sh
```

Then run your Pulumi commands as usual. For detailed steps, see `docs/pages/docs/guide/debug-pulumi-modules.mdx`.


