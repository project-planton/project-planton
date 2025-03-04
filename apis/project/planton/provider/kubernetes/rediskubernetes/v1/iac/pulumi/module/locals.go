package module

import (
	"fmt"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	IngressExternalHostname string
	IngressInternalHostname string
	KubePortForwardCommand  string
	KubeServiceFqdn         string
	KubeServiceName         string
	Namespace               string
	RedisKubernetes         *rediskubernetesv1.RedisKubernetes
	RedisPodSelectorLabels  map[string]string
	Labels                  map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *rediskubernetesv1.RedisKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.RedisKubernetes = stackInput.Target

	redisKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: redisKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "redis_kubernetes",
	}

	if redisKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = redisKubernetes.Metadata.Id
	}

	if redisKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = redisKubernetes.Metadata.Org
	}

	if redisKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = redisKubernetes.Metadata.Env
	}

	locals.Namespace = redisKubernetes.Metadata.Name

	if redisKubernetes.Metadata.Labels != nil &&
		redisKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = redisKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.RedisPodSelectorLabels = map[string]string{
		"app.kubernetes.io/component": "master",
		"app.kubernetes.io/instance":  redisKubernetes.Metadata.Name,
		"app.kubernetes.io/name":      "redis",
	}

	locals.KubeServiceName = fmt.Sprintf("%s-master", redisKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if redisKubernetes.Spec.Ingress == nil ||
		!redisKubernetes.Spec.Ingress.IsEnabled ||
		redisKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		redisKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		redisKubernetes.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(outputs.ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.InternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
