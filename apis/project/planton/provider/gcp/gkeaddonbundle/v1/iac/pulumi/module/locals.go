package module

import (
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkeaddonbundle/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpCredentialSpec *gcpcredentialv1.GcpCredentialSpec
	GkeAddonBundle    *gkeaddonbundlev1.GkeAddonBundle
	KubernetesLabels  map[string]string
	GcpLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gkeaddonbundlev1.GkeAddonBundleStackInput) *Locals {
	locals := &Locals{}

	locals.GcpCredentialSpec = stackInput.GcpCredential
	locals.GkeAddonBundle = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceKind: "gke-cluster",
	}

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceKind: "gke-cluster",
	}

	if locals.GkeAddonBundle.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GkeAddonBundle.Metadata.Org
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = locals.GkeAddonBundle.Metadata.Org
	}

	if locals.GkeAddonBundle.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GkeAddonBundle.Metadata.Id
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = locals.GkeAddonBundle.Metadata.Id
	}

	return locals
}
