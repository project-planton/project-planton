package module

import (
	"fmt"
	kuberneteshttpendpointv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kuberneteshttpendpoint/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	EndpointDomainName     string
	IngressCertSecretName  string
	KubernetesHttpEndpoint *kuberneteshttpendpointv1.KubernetesHttpEndpoint
	Labels                 map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kuberneteshttpendpointv1.KubernetesHttpEndpointStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.KubernetesHttpEndpoint = stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Environment:  stackInput.Target.Metadata.Env,
		kuberneteslabelkeys.Organization: stackInput.Target.Metadata.Org,
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   stackInput.Target.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesHttpEndpoint.String(),
	}

	locals.EndpointDomainName = locals.KubernetesHttpEndpoint.Metadata.Name

	locals.IngressCertSecretName = fmt.Sprintf("cert-%s", locals.KubernetesHttpEndpoint.Metadata.Name)

	ctx.Export(OpNamespace, pulumi.String(vars.IstioIngressNamespace))

	return locals
}
