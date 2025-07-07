package module

import (
	"fmt"
	"strconv"

	gcpgkenodepoolv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkenodepool/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals aggregates frequently‑used values so that downstream functions
// can reference them without additional look‑ups or helpers.
// Keeping this struct simple (no getters) helps mimic Terraform's "locals".
type Locals struct {
	GcpGkeNodePool *gcpgkenodepoolv1.GcpGkeNodePool

	// Derived convenience values
	GcpLabels        map[string]string
	KubernetesLabels map[string]string
	NetworkTag       string
	ClusterName      string
}

// initializeLocals builds the Locals struct from the generated stack‑input message.
// It follows exactly the pattern used by the existing GKE‑cluster module.
func initializeLocals(ctx *pulumi.Context, stackInput *gcpgkenodepoolv1.GcpGkeNodePoolStackInput) *Locals {
	locals := &Locals{}

	locals.GcpGkeNodePool = stackInput.Target

	// Attempt to resolve the parent cluster name from the foreign‑key field.
	// We check both the literal value and any reference string; fallback to empty.
	if stackInput.Target.Spec.ClusterName != nil {
		locals.ClusterName = stackInput.Target.Spec.ClusterName.GetValue()
	}

	// Base label maps
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: locals.GcpGkeNodePool.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpGkeNodePool.String(),
	}
	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: locals.GcpGkeNodePool.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpGkeNodePool.String(),
	}

	// Optional metadata fields
	if locals.GcpGkeNodePool.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpGkeNodePool.Metadata.Org
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = locals.GcpGkeNodePool.Metadata.Org
	}
	if locals.GcpGkeNodePool.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpGkeNodePool.Metadata.Env
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = locals.GcpGkeNodePool.Metadata.Env
	}
	if locals.GcpGkeNodePool.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GcpGkeNodePool.Metadata.Id
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = locals.GcpGkeNodePool.Metadata.Id
	}

	// Network tag follows the same "gke-<clusterName>" convention as the cluster module.
	if locals.ClusterName != "" {
		locals.NetworkTag = fmt.Sprintf("gke-%s", locals.ClusterName)
	} else {
		locals.NetworkTag = fmt.Sprintf("gke-%s", locals.GcpGkeNodePool.Metadata.Name)
	}

	return locals
}
