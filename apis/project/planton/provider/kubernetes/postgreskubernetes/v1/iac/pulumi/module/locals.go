package module

import (
	"fmt"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1/iac/pulumi/module/outputs"
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
	PostgresKubernetes      *postgreskubernetesv1.PostgresKubernetes
	PostgresPodSectorLabels map[string]string
	Labels                  map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *postgreskubernetesv1.PostgresKubernetesStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	postgresKubernetes := stackInput.Target

	//assign value for the local variable to make it available across the module.
	locals.PostgresKubernetes = postgresKubernetes

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   postgresKubernetes.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: "postgres_kubernetes",
	}

	if postgresKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = postgresKubernetes.Metadata.Org
	}

	if postgresKubernetes.Metadata.Env != nil {
		locals.Labels[kuberneteslabelkeys.Environment] = postgresKubernetes.Metadata.Env.Id

	}

	//decide on the namespace
	locals.Namespace = postgresKubernetes.Metadata.Id

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.PostgresPodSectorLabels = map[string]string{
		"planton.cloud/resource-kind": "postgres_kubernetes",
		"planton.cloud/resource-id":   postgresKubernetes.Metadata.Id,
	}

	ctx.Export(outputs.UsernameSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			postgresKubernetes.Metadata.Id))
	ctx.Export(outputs.UsernameSecretKey, pulumi.String("username"))

	ctx.Export(outputs.PasswordSecretName,
		pulumi.Sprintf("postgres.db-%s.credentials.postgresql.acid.zalan.do",
			postgresKubernetes.Metadata.Id))
	ctx.Export(outputs.PasswordSecretKey, pulumi.String("password"))

	locals.KubeServiceName = fmt.Sprintf("%s-master", postgresKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if postgresKubernetes.Spec.Ingress == nil ||
		!postgresKubernetes.Spec.Ingress.IsEnabled ||
		postgresKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", postgresKubernetes.Metadata.Id,
		postgresKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", postgresKubernetes.Metadata.Id,
		postgresKubernetes.Spec.Ingress.DnsDomain)

	//export ingress hostnames
	ctx.Export(outputs.ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.InternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
