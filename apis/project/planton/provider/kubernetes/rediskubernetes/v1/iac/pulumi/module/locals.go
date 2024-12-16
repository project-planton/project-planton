package module

import (
	"fmt"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
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

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	redisKubernetes := stackInput.Target

	//assign value for the local variable to make it available across the module.
	locals.RedisKubernetes = redisKubernetes

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   redisKubernetes.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: "redis_kubernetes",
	}

	if redisKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = redisKubernetes.Metadata.Org
	}

	if redisKubernetes.Metadata.Env != nil {
		locals.Labels[kuberneteslabelkeys.Environment] = redisKubernetes.Metadata.Env.Id
	}

	//decide on the namespace
	locals.Namespace = redisKubernetes.Metadata.Id

	ctx.Export(outputs.NAMESPACE, pulumi.String(locals.Namespace))

	locals.RedisPodSelectorLabels = map[string]string{
		"app.kubernetes.io/component": "master",
		"app.kubernetes.io/instance":  redisKubernetes.Metadata.Id,
		"app.kubernetes.io/name":      "redis",
	}

	locals.KubeServiceName = fmt.Sprintf("%s-master", redisKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.SERVICE, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KUBE_ENDPOINT, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PORT_FORWARD_COMMAND, pulumi.String(locals.KubePortForwardCommand))

	if redisKubernetes.Spec.Ingress == nil ||
		!redisKubernetes.Spec.Ingress.IsEnabled ||
		redisKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", redisKubernetes.Metadata.Id,
		redisKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", redisKubernetes.Metadata.Id,
		redisKubernetes.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(outputs.EXTERNAL_HOSTNAME, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.INTERNAL_HOSTNAME, pulumi.String(locals.IngressInternalHostname))

	return locals
}
