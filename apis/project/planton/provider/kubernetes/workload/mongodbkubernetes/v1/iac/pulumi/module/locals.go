package module

import (
	"fmt"
	"strconv"

	mongodbkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/mongodbkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	IngressExternalHostname  string
	KubePortForwardCommand   string
	KubeServiceFqdn          string
	KubeServiceName          string
	MongodbKubernetes        *mongodbkubernetesv1.MongodbKubernetes
	Namespace                string
	MongodbPodSelectorLabels map[string]string
	KubernetesLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *mongodbkubernetesv1.MongodbKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.MongodbKubernetes = stackInput.Target

	target := stackInput.Target

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_MongodbKubernetes.String(),
	}

	if target.Metadata.Id != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// Priority order:
	// 1. Default: metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))
	ctx.Export(OpUsername, pulumi.String(vars.RootUsername))
	ctx.Export(OpPasswordSecretName, pulumi.String(target.Metadata.Name))
	ctx.Export(OpPasswordSecretKey, pulumi.String(vars.MongodbRootPasswordKey))

	locals.KubeServiceName = target.Metadata.Name

	// Percona operator uses these labels for pod selection
	// These labels are automatically applied by the operator to MongoDB pods
	locals.MongodbPodSelectorLabels = map[string]string{
		"app.kubernetes.io/name":       "percona-server-mongodb",
		"app.kubernetes.io/instance":   target.Metadata.Name,
		"app.kubernetes.io/managed-by": "percona-server-mongodb-operator",
	}

	//export kubernetes service name
	ctx.Export(OpService, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(OpKubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, target.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled ||
		target.Spec.Ingress.Hostname == "" {
		return locals
	}

	locals.IngressExternalHostname = target.Spec.Ingress.Hostname

	ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))

	return locals
}
