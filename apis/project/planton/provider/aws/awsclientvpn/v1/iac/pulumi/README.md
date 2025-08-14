# Pulumi Module to Deploy AwsClientVpn

This Pulumi program deploys an `AwsClientVpn` endpoint using the ProjectPlanton CLI.

## CLI commands

```bash
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

## Examples
See `examples.md` in this directory for example manifests and flows. These mirror the root-level examples for AwsClientVpn.

## Debugging
For local debugging, a `debug.sh` helper is provided. To enable it, uncomment the following in `Pulumi.yaml`:

```yaml
# options:
#   binary: ./debug.sh
```

Then run the preview/update commands as usual; Pulumi will execute the compiled binary under Delve.

For more details, refer to the docs page: docs/pages/docs/guide/debug-pulumi-modules.mdx
