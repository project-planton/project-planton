// Package localz instead of locals to avoid naming collision w/ "locals" for the instance name created for the struct.
package localz

import (
	"fmt"
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gkeclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpCredentialSpec                     *gcpcredentialv1.GcpCredentialSpec
	GkeCluster                            *gkeclusterv1.GkeCluster
	KubernetesPodSecondaryIpRangeName     string
	KubernetesServiceSecondaryIpRangeName string
	KubernetesLabels                      map[string]string
	GcpLabels                             map[string]string
	ContainerClusterLoggingComponentList  []string
	NetworkTag                            string
}

func Initialize(ctx *pulumi.Context, stackInput *gkeclusterv1.GkeClusterStackInput) *Locals {
	gkeCluster := stackInput.Target

	locals := &Locals{}

	locals.GcpCredentialSpec = stackInput.GcpCredential
	locals.GkeCluster = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceKind: "gke-cluster",
	}

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceKind: "gke-cluster",
	}

	if locals.GkeCluster.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GkeCluster.Metadata.Org
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = locals.GkeCluster.Metadata.Org
	}

	if locals.GkeCluster.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GkeCluster.Metadata.Id
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = locals.GkeCluster.Metadata.Id
	}

	locals.KubernetesPodSecondaryIpRangeName = fmt.Sprintf("gke-%s-pods", gkeCluster.Metadata.Name)
	locals.KubernetesServiceSecondaryIpRangeName = fmt.Sprintf("gke-%s-services", gkeCluster.Metadata.Name)
	locals.NetworkTag = fmt.Sprintf("gke-%s", gkeCluster.Metadata.Name)

	locals.ContainerClusterLoggingComponentList = []string{"SYSTEM_COMPONENTS"}

	if gkeCluster.Spec.IsWorkloadLogsEnabled {
		locals.ContainerClusterLoggingComponentList = append(locals.ContainerClusterLoggingComponentList,
			"WORKLOADS")
	}

	return locals
}
