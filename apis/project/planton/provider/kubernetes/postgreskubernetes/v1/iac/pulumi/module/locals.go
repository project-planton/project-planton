package module

import (
	"fmt"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
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
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_PostgresKubernetes.String(),
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

	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	locals.PostgresPodSectorLabels = map[string]string{
		"team": vars.TeamId,
	}

	ctx.Export(OpUsernameSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			target.Metadata.Name))
	ctx.Export(OpUsernameSecretKey, pulumi.String("username"))

	ctx.Export(OpPasswordSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			target.Metadata.Name))
	ctx.Export(OpPasswordSecretKey, pulumi.String("password"))

	locals.KubeServiceName = fmt.Sprintf("%s-master", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(OpService, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(OpKubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled ||
		target.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(OpInternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
