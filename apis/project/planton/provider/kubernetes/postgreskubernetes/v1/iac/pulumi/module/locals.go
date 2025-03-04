package module

import (
	"fmt"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
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
	PostgresKubernetes      *postgreskubernetesv1.PostgresKubernetes
	PostgresPodSectorLabels map[string]string
	Labels                  map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *postgreskubernetesv1.PostgresKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.PostgresKubernetes = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: string(apiresourcekind.PostgresKubernetesKind),
	}

	if target.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.PostgresPodSectorLabels = map[string]string{
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
	}

	ctx.Export(outputs.UsernameSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			target.Metadata.Name))
	ctx.Export(outputs.UsernameSecretKey, pulumi.String("username"))

	ctx.Export(outputs.PasswordSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			target.Metadata.Name))
	ctx.Export(outputs.PasswordSecretKey, pulumi.String("password"))

	locals.KubeServiceName = fmt.Sprintf("%s-master", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.IsEnabled ||
		target.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(outputs.ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.InternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
