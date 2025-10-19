package module

import (
	"fmt"
	"strconv"

	gcpgkeclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkecluster/v1"
	gcpprovider "github.com/project-planton/project-planton/apis/project/planton/provider/gcp"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	GcpProviderConfig                     *gcpprovider.GcpProviderConfig
	GcpGkeCluster                         *gcpgkeclusterv1.GcpGkeCluster
	KubernetesPodSecondaryIpRangeName     string
	KubernetesServiceSecondaryIpRangeName string
	KubernetesLabels                      map[string]string
	GcpLabels                             map[string]string
	ContainerClusterLoggingComponentList  []string
	NetworkTag                            string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpgkeclusterv1.GcpGkeClusterStackInput) *Locals {
	locals := &Locals{}

	locals.GcpGkeCluster = stackInput.Target

	target := stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster.String(),
	}

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster.String(),
	}

	if locals.GcpGkeCluster.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = target.Metadata.Org
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if locals.GcpGkeCluster.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = target.Metadata.Env
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	if locals.GcpGkeCluster.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = target.Metadata.Id
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	locals.KubernetesPodSecondaryIpRangeName = fmt.Sprintf("gke-%s-pods", target.Metadata.Name)
	locals.KubernetesServiceSecondaryIpRangeName = fmt.Sprintf("gke-%s-services", target.Metadata.Name)
	locals.NetworkTag = fmt.Sprintf("gke-%s", target.Metadata.Name)

	locals.ContainerClusterLoggingComponentList = []string{"SYSTEM_COMPONENTS"}

	if target.Spec.IsWorkloadLogsEnabled {
		locals.ContainerClusterLoggingComponentList = append(locals.ContainerClusterLoggingComponentList,
			"WORKLOADS")
	}

	locals.GcpProviderConfig = stackInput.ProviderConfig

	return locals
}
