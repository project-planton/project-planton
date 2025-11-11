package module

import (
	"strconv"

	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	elasticoperatorkubernetesv1 "github.com/project-planton/project-planton/pkg/provider/kubernetes/addon/elasticoperatorkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/shared/cloudresourcekind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	ElasticOperatorKubernetes *elasticoperatorkubernetesv1.ElasticOperatorKubernetes
	KubeLabels                map[string]string
}

func initializeLocals(_ *pulumi.Context, in *elasticoperatorkubernetesv1.ElasticOperatorKubernetesStackInput) *Locals {
	var l Locals
	l.ElasticOperatorKubernetes = in.Target

	l.KubeLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: in.Target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_ElasticOperatorKubernetes.String(),
	}

	if id := in.Target.Metadata.Id; id != "" {
		l.KubeLabels[kuberneteslabelkeys.ResourceId] = id
	}
	if org := in.Target.Metadata.Org; org != "" {
		l.KubeLabels[kuberneteslabelkeys.Organization] = org
	}
	if env := in.Target.Metadata.Env; env != "" {
		l.KubeLabels[kuberneteslabelkeys.Environment] = env
	}
	return &l
}
