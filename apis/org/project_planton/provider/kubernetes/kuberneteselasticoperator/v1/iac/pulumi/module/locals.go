package module

import (
	"strconv"

	kuberneteselasticoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteselasticoperator/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	KubernetesElasticOperator *kuberneteselasticoperatorv1.KubernetesElasticOperator
	KubeLabels                map[string]string
}

func initializeLocals(_ *pulumi.Context, in *kuberneteselasticoperatorv1.KubernetesElasticOperatorStackInput) *Locals {
	var l Locals
	l.KubernetesElasticOperator = in.Target

	l.KubeLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: in.Target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesElasticOperator.String(),
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
