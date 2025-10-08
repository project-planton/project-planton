# System Prompt: Create ClickhouseOperatorKubernetes Deployment Component

## Objective

Create a new Kubernetes operator deployment component called `ClickhouseOperatorKubernetes` that deploys the **Altinity ClickHouse Operator** to a Kubernetes cluster. This follows the same pattern as other operator deployments in the project (KafkaOperatorKubernetes, SolrOperatorKubernetes, PostgresOperatorKubernetes).

## Context

Project Planton is a multi-cloud deployment framework following the Pareto principle (80/20 rule). Operators are deployed as separate, reusable components that can be referenced by workload deployments.

## Reference Implementations

Study these existing operator deployments to understand the pattern:
- `project-planton/apis/project/planton/provider/kubernetes/addon/kafkaoperatorkubernetes/v1/`
- `project-planton/apis/project/planton/provider/kubernetes/addon/solroperatorkubernetes/v1/`
- `project-planton/apis/project/planton/provider/kubernetes/addon/postgresoperatorkubernetes/v1/`

## Target Location

Create the component at:
```
project-planton/apis/project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1/
```

## Required Files Structure

```
clickhouseoperatorkubernetes/v1/
├── BUILD.bazel
├── api.proto
├── spec.proto
├── stack_input.proto
├── stack_outputs.proto
├── api.pb.go (generated)
├── spec.pb.go (generated)
├── stack_input.pb.go (generated)
├── stack_outputs.pb.go (generated)
└── iac/
    ├── pulumi/
    │   ├── BUILD.bazel
    │   ├── Makefile
    │   ├── Pulumi.yaml
    │   ├── debug.sh
    │   ├── main.go
    │   └── module/
    │       ├── BUILD.bazel
    │       ├── main.go
    │       ├── outputs.go
    │       ├── vars.go
    │       └── clickhouse_operator.go
    └── tf/
        ├── main.tf
        ├── provider.tf
        └── variables.tf
```

## Altinity ClickHouse Operator Details

### Official Resources
- **GitHub**: https://github.com/Altinity/clickhouse-operator
- **Documentation**: https://github.com/Altinity/clickhouse-operator/tree/master/docs
- **Helm Chart**: https://github.com/Altinity/clickhouse-operator/tree/master/deploy/helm/clickhouse-operator
- **CRDs**: ClickHouseInstallation (CHI), ClickHouseOperatorConfiguration

### Installation Method
Use Helm chart deployment (similar to SolrOperatorKubernetes pattern):
- **Chart Repository**: `https://docs.altinity.com/clickhouse-operator/`
- **Chart Name**: `clickhouse-operator`
- **Recommended Version**: Check latest stable version (currently ~0.23.x)
- **Namespace**: `clickhouse-operator`

### Key Configuration
The operator watches all namespaces by default. Key Helm values to consider:
- `operator.watchNamespaces`: Can be left empty for cluster-wide
- `operator.createCRD`: Set to true to automatically install CRDs
- Resource limits for the operator pod

## Proto Files Specification

### api.proto
```protobuf
syntax = "proto3";

package project.planton.provider.kubernetes.addon.clickhouseoperatorkubernetes.v1;

import "buf/validate/validate.proto";
import "project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1/spec.proto";
import "project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1/stack_outputs.proto";
import "project/planton/shared/status.proto";
import "project/planton/shared/metadata.proto";

//clickhouse-operator-kubernetes
message ClickhouseOperatorKubernetes {
  //api-version
  string api_version = 1 [
    (buf.validate.field).string.const = 'kubernetes.project-planton.org/v1'
  ];

  //resource-kind
  string kind = 2 [
    (buf.validate.field).string.const = 'ClickhouseOperatorKubernetes'
  ];

  //metadata
  project.planton.shared.ApiResourceMetadata metadata = 3 [
    (buf.validate.field).required = true
  ];

  //spec
  ClickhouseOperatorKubernetesSpec spec = 4 [
    (buf.validate.field).required = true
  ];

  //status
  ClickhouseOperatorKubernetesStatus status = 5;
}

//clickhouse-operator-kubernetes status
message ClickhouseOperatorKubernetesStatus {
  //stack-outputs
  ClickhouseOperatorKubernetesStackOutputs outputs = 1;
}
```

### spec.proto
Follow the minimal operator spec pattern (target cluster + container resources only):

```protobuf
syntax = "proto3";

package project.planton.provider.kubernetes.addon.clickhouseoperatorkubernetes.v1;

import "buf/validate/validate.proto";
import "project/planton/shared/kubernetes/kubernetes.proto";
import "project/planton/shared/kubernetes/options.proto";
import "project/planton/shared/kubernetes/target_cluster.proto";

message ClickhouseOperatorKubernetesSpec {
  // The Kubernetes cluster to install this operator on
  project.planton.shared.kubernetes.KubernetesAddonTargetCluster target_cluster = 1;
  
  // Container resource specifications for the operator
  ClickhouseOperatorKubernetesSpecContainer container = 2 [
    (buf.validate.field).required = true
  ];
}

message ClickhouseOperatorKubernetesSpecContainer {
  // CPU and memory resources for the operator pod
  project.planton.shared.kubernetes.ContainerResources resources = 1 [
    (project.planton.shared.kubernetes.default_container_resources) = {
      limits {
        cpu: "1000m"
        memory: "1Gi"
      },
      requests {
        cpu: "100m"
        memory: "256Mi"
      }
    }
  ];
}
```

### stack_input.proto
```protobuf
syntax = "proto3";

package project.planton.provider.kubernetes.addon.clickhouseoperatorkubernetes.v1;

import "project/planton/credential/kubernetesclustercredential/v1/spec.proto";
import "project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1/api.proto";

message ClickhouseOperatorKubernetesStackInput {
  ClickhouseOperatorKubernetes target = 1;
  project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec provider_credential = 2;
}
```

### stack_outputs.proto
```protobuf
syntax = "proto3";

package project.planton.provider.kubernetes.addon.clickhouseoperatorkubernetes.v1;

message ClickhouseOperatorKubernetesStackOutputs {
  // Namespace where the operator is installed
  string namespace = 1;
}
```

## Pulumi Module Implementation Pattern

### module/vars.go
```go
package module

var vars = struct {
    Namespace        string
    HelmChartName    string
    HelmChartRepo    string
    HelmChartVersion string
}{
    Namespace:        "clickhouse-operator",
    HelmChartName:    "clickhouse-operator",
    HelmChartRepo:    "https://docs.altinity.com/clickhouse-operator/",
    HelmChartVersion: "0.23.6", // Check for latest stable version
}
```

### module/clickhouse_operator.go
Key implementation steps:
1. Create namespace with proper labels
2. Deploy Helm chart with:
   - CRD installation enabled
   - Cluster-wide watch (all namespaces)
   - Resource limits from spec
   - Timeout and cleanup settings
3. Export namespace to stack outputs

Use the SolrOperatorKubernetes implementation as the closest reference, as it also handles CRD installation.

### module/main.go
Standard pattern:
1. Initialize locals (labels, metadata)
2. Create Kubernetes provider from credentials
3. Call operator installation function
4. Return any errors

## Build Integration

### BUILD.bazel Files
Follow the exact pattern from other operator deployments. Ensure proper dependencies on:
- `//apis/project/planton/credential/kubernetesclustercredential/v1`
- `//apis/project/planton/shared`
- `//apis/project/planton/shared/kubernetes`
- Protobuf validation
- Pulumi Kubernetes SDK v4

## Testing Considerations

After implementation, the operator should:
1. Deploy successfully to any Kubernetes cluster
2. Install all Altinity ClickHouse CRDs
3. Watch all namespaces for ClickHouseInstallation resources
4. Be reusable by workload deployments (future ClickHouseKubernetes component)

## Key Differences from Existing Patterns

The Altinity operator is more complex than Kafka/Postgres operators because:
- It manages ZooKeeper coordination for clusters
- It handles complex ClickHouse cluster topologies
- It has extensive CRD configurations

Keep the operator deployment simple - just get it installed. The complexity of configuring actual ClickHouse clusters will be in the separate ClickHouseKubernetes workload component.

## Success Criteria

1. All proto files compile without errors
2. Pulumi module successfully deploys the operator
3. Operator pod is running in `clickhouse-operator` namespace
4. CRDs are installed and ready
5. Pattern matches existing operator deployments (consistency)
6. Documentation is clear and follows project standards

## Commands to Run After Implementation

```bash
# Generate protobuf code
cd project-planton/apis/project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1
buf generate

# Build Pulumi module
cd iac/pulumi
make build

# Test deployment (with proper credentials)
project-planton pulumi up --manifest clickhouse-operator.yaml \
  --module-dir apis/project/planton/provider/kubernetes/addon/clickhouseoperatorkubernetes/v1/iac/pulumi
```

## Notes

- The operator deployment is stateless - it just installs the operator
- Actual ClickHouse clusters will be deployed separately using ClickHouseKubernetes workload component
- Keep spec.proto minimal - only target cluster and operator resources
- Follow existing conventions exactly for consistency
- Don't add features that aren't in other operator deployments

