package module

import (
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkeaddonbundle/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
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

	locals.GkeAddonBundle = stackInput.Target

	target := stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: string(apiresourcekind.GkeAddonBundleKind),
	}

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: string(apiresourcekind.GkeAddonBundleKind),
	}

	if target.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = target.Metadata.Id
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = target.Metadata.Org
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = target.Metadata.Env
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	locals.GcpCredentialSpec = stackInput.GcpCredential

	return locals
}
