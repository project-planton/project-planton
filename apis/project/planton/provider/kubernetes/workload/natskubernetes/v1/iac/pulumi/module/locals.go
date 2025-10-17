package module

import (
	"fmt"
	"strconv"

	natskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/natskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps all frequently-used, derived values in one place –
// similar to a Terraform “locals {}” block.
type Locals struct {
	Namespace         string
	Labels            map[string]string
	NatsKubernetes    *natskubernetesv1.NatsKubernetes
	ClientURLInternal string
	ClientURLExternal string
	TlsSecretName     string
	TlsSecretKey      string
}

// initializeLocals builds the Locals struct and immediately exports the
// values required by NatsKubernetesStackOutputs.
func initializeLocals(ctx *pulumi.Context,
	stackInput *natskubernetesv1.NatsKubernetesStackInput) *Locals {

	locals := &Locals{}
	locals.NatsKubernetes = stackInput.Target
	target := stackInput.Target

	// ------------------------------- labels ----------------------------------
	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_NatsKubernetes.String(),
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

	// ------------------------------ namespace --------------------------------
	locals.Namespace = target.Metadata.Name
	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {

		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// ------------------------- internal client URL ---------------------------
	// Helm chart installs a Service named "<name>-nats".  Port 4222 is fixed.
	serviceName := fmt.Sprintf("%s-nats", target.Metadata.Name)
	locals.ClientURLInternal = fmt.Sprintf("nats://%s.%s.svc.cluster.local:%d",
		serviceName, locals.Namespace, vars.NatsClientPort)
	ctx.Export(OpClientUrlInternal, pulumi.String(locals.ClientURLInternal))

	// ------------------------------ ingress ----------------------------------
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Enabled &&
		target.Spec.Ingress.Hostname != "" {

		locals.ClientURLExternal = fmt.Sprintf("nats://%s:%d",
			target.Spec.Ingress.Hostname, vars.NatsClientPort)
		ctx.Export(OpClientUrlExternal, pulumi.String(locals.ClientURLExternal))
	}

	// -------------------- auth / token secret outputs ------------------------
	// Secret names are deterministic so callers / automation can pre-bake RBAC.
	ctx.Export(OpAuthSecretName, pulumi.String(vars.AdminAuthSecretName))
	ctx.Export(OpAuthSecretKey, pulumi.String(vars.AdminAuthSecretKey))

	// ----------------------- TLS certificate secret --------------------------
	if target.Spec.TlsEnabled {
		locals.TlsSecretName = fmt.Sprintf("tls-%s", locals.Namespace)
		locals.TlsSecretKey = vars.TlsCertKey
		ctx.Export(OpTlsSecretName, pulumi.String(locals.TlsSecretName))
		ctx.Export(OpTlsSecretKey, pulumi.String(locals.TlsSecretKey))
	}

	// ------------------------ jet-stream domain ------------------------------
	if !target.Spec.DisableJetStream {
		localsJetDomain := fmt.Sprintf("%s", locals.Namespace) // simple default
		ctx.Export(OpJetStreamDomain, pulumi.String(localsJetDomain))
	}

	return locals
}
