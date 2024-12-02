# kubernetes-crd-pulumi-types

Using [crd2pulumi](https://github.com/pulumi/crd2pulumi) makes is possible to generate strong types for 
Kubernetes Custom Resource Definitions, which makes it easy to create those custom resources in pulumi modules.

1. Install `crd2pulumi` cli.

```shell
brew install pulumi/tap/crd2pulumi
```

2. Generate Pulumi Types for CRDs

```bash
make build
```

## Golang Bug

There is an open issue [#89](https://github.com/pulumi/crd2pulumi/issues/89) about `goPath` not working as expected. As a result, the kubernetes resources are created
inside `pkg/kubernetes` when the `goPath` is specified as `pkg/istio` or `pkg/certmanager`.
