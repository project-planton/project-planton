package module

import (
	"fmt"
	mongodbkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/mongodbkubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	IngressExternalHostname  string
	IngressInternalHostname  string
	KubePortForwardCommand   string
	KubeServiceFqdn          string
	KubeServiceName          string
	KubernetesLabels         map[string]string
	MongodbKubernetes        *mongodbkubernetesv1.MongodbKubernetes
	Namespace                string
	MongodbPodSelectorLabels map[string]string
	Labels                   map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *mongodbkubernetesv1.MongodbKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.MongodbKubernetes = stackInput.Target

	mongodbKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: mongodbKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "mongodb_kubernetes",
	}

	if mongodbKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = mongodbKubernetes.Metadata.Id
	}

	if mongodbKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = mongodbKubernetes.Metadata.Org
	}

	if mongodbKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = mongodbKubernetes.Metadata.Env
	}

	locals.Namespace = mongodbKubernetes.Metadata.Name

	if mongodbKubernetes.Metadata.Labels != nil &&
		mongodbKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = mongodbKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))
	ctx.Export(outputs.Username, pulumi.String(vars.RootUsername))
	ctx.Export(outputs.PasswordSecretName, pulumi.String(mongodbKubernetes.Metadata.Name))
	ctx.Export(outputs.PasswordSecretKey, pulumi.String(vars.MongodbRootPasswordKey))

	locals.KubeServiceName = mongodbKubernetes.Metadata.Name

	locals.MongodbPodSelectorLabels = map[string]string{
		"app.kubernetes.io/component": "mongodb",
		"app.kubernetes.io/instance":  mongodbKubernetes.Metadata.Id,
		"app.kubernetes.io/name":      "mongodb",
	}

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, mongodbKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if mongodbKubernetes.Spec.Ingress == nil ||
		!mongodbKubernetes.Spec.Ingress.IsEnabled ||
		mongodbKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", mongodbKubernetes.Metadata.Id,
		mongodbKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", mongodbKubernetes.Metadata.Id,
		mongodbKubernetes.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(outputs.ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.InternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
