package module

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"strconv"

	cronjobkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workloads/cronjobkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals holds local references and configurations used by the module.
type Locals struct {
	CronJobKubernetes   *cronjobkubernetesv1.CronJobKubernetes
	Namespace           string
	Labels              map[string]string
	ImagePullSecretData map[string]string
}

// initializeLocals configures the Locals struct from the given stack input.
func initializeLocals(ctx *pulumi.Context, stackInput *cronjobkubernetesv1.CronJobKubernetesStackInput) (*Locals, error) {
	locals := &Locals{}
	target := stackInput.Target

	locals.CronJobKubernetes = target
	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_CronJobKubernetes.String(),
	}

	// add resource id if present
	if target.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	// add organization if present
	if target.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	// add environment if present
	if target.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// default namespace to resource's name
	locals.Namespace = target.Metadata.Name
	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	// export namespace as an output
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// handle docker config json if specified
	if stackInput.DockerConfigJson != "" {
		locals.ImagePullSecretData = map[string]string{
			".dockerconfigjson": stackInput.DockerConfigJson,
		}
	}

	return locals, nil
}
