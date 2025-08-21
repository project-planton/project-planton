# Pulumi Module to Deploy AwsRdsCluster

This Pulumi program deploys an AWS RDS Cluster (Aurora MySQL/PostgreSQL or Multi-AZ DB Cluster) using the Project Planton API and module.

## Requirements
- Project Planton CLI built locally
- Valid AWS credential provided via the CLI stack input (not in `spec`)

## CLI commands

Preview:

```shell
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

Update (apply):

```shell
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes
```

Refresh:

```shell
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

Destroy:

```shell
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes
```

## Examples

See `../examples.md` for sample manifests.

## Debugging

Optionally enable debugging by setting a binary in `Pulumi.yaml` and using the `debug.sh` script.


