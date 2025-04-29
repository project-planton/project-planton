package module

import (
	"fmt"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	kafkakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	Namespace       string
	KafkaKubernetes *kafkakubernetesv1.KafkaKubernetes
	Labels          map[string]string

	IngressCertClusterIssuerName string
	// bootstrap
	IngressExternalBootstrapHostname string
	IngressInternalBootstrapHostname string
	IngressExternalBrokerHostnames   []string
	IngressInternalBrokerHostnames   []string
	IngressHostnames                 []string
	BootstrapKubeServiceFqdn         string
	BootstrapKubeServiceName         string

	// schema registry
	IngressSchemaRegistryCertSecretName   string
	IngressExternalSchemaRegistryHostname string
	IngressInternalSchemaRegistryHostname string
	IngressSchemaRegistryHostnames        []string
	SchemaRegistryKubeServiceFqdn         string

	// kowl dashboard
	IngressKowlCertSecretName                              string
	IngressExternalKowlHostname                            string
	KowlKubeServiceFqdn                                    string
	KafkaIngressPrivateListenerLoadBalancerAnnotationKey   string
	KafkaIngressPrivateListenerLoadBalancerAnnotationValue string
	KafkaIngressPublicListenerLoadBalancerAnnotationKey    string
	KafkaIngressPublicListenerLoadBalancerAnnotationValue  string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kafkakubernetesv1.KafkaKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.KafkaKubernetes = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KafkaKubernetes.String(),
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
	ctx.Export(OpUsername, pulumi.String(vars.AdminUsername))
	ctx.Export(OpPasswordSecretName, pulumi.String(vars.SaslPasswordSecretName))
	ctx.Export(OpPasswordSecretKey, pulumi.String(vars.SaslPasswordKeyInSecret))

	locals.BootstrapKubeServiceName = fmt.Sprintf("%s-kafka-bootstrap", locals.Namespace)

	locals.BootstrapKubeServiceFqdn = fmt.Sprintf("%s.%s.svc", locals.BootstrapKubeServiceName, locals.Namespace)

	// schema registry related locals data
	if locals.KafkaKubernetes.Spec.SchemaRegistryContainer != nil &&
		locals.KafkaKubernetes.Spec.SchemaRegistryContainer.IsEnabled {

		locals.IngressSchemaRegistryCertSecretName = fmt.Sprintf("cert-%s-schema-registry", locals.Namespace)

		locals.IngressExternalSchemaRegistryHostname = fmt.Sprintf("%s-schema-registry.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

		locals.IngressInternalSchemaRegistryHostname = fmt.Sprintf("%s-schema-registry-internal.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

		ctx.Export(OpSchemaRegistryExternalUrl, pulumi.Sprintf("https://%s", locals.IngressExternalSchemaRegistryHostname))
		ctx.Export(OpSchemaRegistryInternalUrl, pulumi.Sprintf("https://%s", locals.IngressInternalSchemaRegistryHostname))

		locals.IngressSchemaRegistryHostnames = []string{
			locals.IngressExternalSchemaRegistryHostname,
			locals.IngressInternalSchemaRegistryHostname,
		}
		locals.SchemaRegistryKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", vars.SchemaRegistryKubeServiceName, locals.Namespace)
	}

	// kowl related locals data
	if locals.KafkaKubernetes.Spec.IsDeployKafkaUi {

		locals.IngressKowlCertSecretName = fmt.Sprintf("cert-%s-kowl", locals.Namespace)

		locals.IngressExternalKowlHostname = fmt.Sprintf("%s-kowl.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

		ctx.Export(OpKafkaUiExternalUrl, pulumi.Sprintf("https://%s", locals.IngressExternalKowlHostname))

		locals.KowlKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", vars.KowlKubeServiceName, locals.Namespace)
	}

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.IsEnabled ||
		target.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalBootstrapHostname = fmt.Sprintf("%s-bootstrap.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

	locals.IngressInternalBootstrapHostname = fmt.Sprintf("%s-bootstrap-internal.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

	ctx.Export(OpBootstrapServerExternalHostname, pulumi.String(locals.IngressExternalBootstrapHostname))
	ctx.Export(OpBootstrapServerInternalHostname, pulumi.String(locals.IngressInternalBootstrapHostname))

	// Creating internal broker hostnames
	ingressInternalBrokerHostnames := make([]string, int(target.Spec.BrokerContainer.Replicas))
	for i := 0; i < int(target.Spec.BrokerContainer.Replicas); i++ {
		ingressInternalBrokerHostnames[i] = fmt.Sprintf("%s-broker-b%d-internal.%s", locals.Namespace, i, target.Spec.Ingress.DnsDomain)
	}
	locals.IngressInternalBrokerHostnames = ingressInternalBrokerHostnames

	// Creating external broker hostnames
	ingressExternalBrokerHostnames := make([]string, int(target.Spec.BrokerContainer.Replicas))
	for i := 0; i < int(target.Spec.BrokerContainer.Replicas); i++ {
		ingressExternalBrokerHostnames[i] = fmt.Sprintf("%s-broker-b%d.%s", locals.Namespace, i, target.Spec.Ingress.DnsDomain)
	}
	locals.IngressExternalBrokerHostnames = ingressExternalBrokerHostnames

	var ingressHostnames = []string{
		locals.IngressInternalBootstrapHostname,
		locals.IngressExternalBootstrapHostname,
	}

	ingressHostnames = append(ingressHostnames, locals.IngressInternalBrokerHostnames...)
	ingressHostnames = append(ingressHostnames, locals.IngressExternalBrokerHostnames...)
	locals.IngressHostnames = ingressHostnames

	//export ingress hostnames
	//ctx.Export(IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	//ctx.Export(IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = target.Spec.Ingress.DnsDomain

	switch stackInput.ProviderCredential.Provider {
	case kubernetesclustercredentialv1.KubernetesProvider_gcp_gke:
		locals.KafkaIngressPrivateListenerLoadBalancerAnnotationKey = "cloud.google.com/load-balancer-type"
		locals.KafkaIngressPrivateListenerLoadBalancerAnnotationValue = "Internal"
		locals.KafkaIngressPublicListenerLoadBalancerAnnotationKey = "cloud.google.com/load-balancer-type"
		locals.KafkaIngressPublicListenerLoadBalancerAnnotationValue = "External"
	}

	return locals
}
