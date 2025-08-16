# Pulumi Module to Deploy AwsIamRole

## CLI usage (ProjectPlanton pulumi)

```bash
# Preview
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Update (apply)
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes

# Refresh
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Destroy
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

## Debugging

This module includes a `debug.sh` helper. To enable debugging, edit `Pulumi.yaml` and uncomment the `runtime.options.binary` line so Pulumi runs the program via the script:

```yaml
name: aws-module-test-pulumi-project
runtime:
  name: go
#  options:
#    binary: ./debug.sh
```

Then make the script executable and run your command (e.g., `preview` or `update`). See `docs/pages/docs/guide/debug-pulumi-modules.mdx` for full instructions.

```bash
chmod +x debug.sh
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```


