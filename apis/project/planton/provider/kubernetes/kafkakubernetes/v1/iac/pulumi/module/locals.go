package module

import (
	"fmt"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	kafkakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
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

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	kafkaKubernetes := stackInput.Target

	//assign value for the locals variable to make it available across the project
	locals.KafkaKubernetes = stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceKind: "kafka_kubernetes",
		kuberneteslabelkeys.ResourceId:   kafkaKubernetes.Metadata.Id,
	}

	if kafkaKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = kafkaKubernetes.Metadata.Org
	}

	if kafkaKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = kafkaKubernetes.Metadata.Env

	}

	//decide on the namespace
	locals.Namespace = kafkaKubernetes.Metadata.Id

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))
	ctx.Export(outputs.Username, pulumi.String(vars.AdminUsername))
	ctx.Export(outputs.PasswordSecretName, pulumi.String(vars.SaslPasswordSecretName))
	ctx.Export(outputs.PasswordSecretKey, pulumi.String(vars.SaslPasswordKeyInSecret))

	locals.BootstrapKubeServiceName = fmt.Sprintf("%s-kafka-bootstrap", kafkaKubernetes.Metadata.Id)

	locals.BootstrapKubeServiceFqdn = fmt.Sprintf("%s.%s.svc", locals.BootstrapKubeServiceName, locals.Namespace)

	// schema registry related locals data
	if locals.KafkaKubernetes.Spec.SchemaRegistryContainer != nil &&
		locals.KafkaKubernetes.Spec.SchemaRegistryContainer.IsEnabled {

		locals.IngressSchemaRegistryCertSecretName = fmt.Sprintf("cert-%s-schema-registry", kafkaKubernetes.Metadata.Id)

		locals.IngressExternalSchemaRegistryHostname = fmt.Sprintf("%s-schema-registry.%s", kafkaKubernetes.Metadata.Id, kafkaKubernetes.Spec.Ingress.DnsDomain)

		locals.IngressInternalSchemaRegistryHostname = fmt.Sprintf("%s-schema-registry-internal.%s", kafkaKubernetes.Metadata.Id, kafkaKubernetes.Spec.Ingress.DnsDomain)

		ctx.Export(outputs.SchemaRegistryExternalUrl, pulumi.Sprintf("https://%s", locals.IngressExternalSchemaRegistryHostname))
		ctx.Export(outputs.SchemaRegistryInternalUrl, pulumi.Sprintf("https://%s", locals.IngressInternalSchemaRegistryHostname))

		locals.IngressSchemaRegistryHostnames = []string{
			locals.IngressExternalSchemaRegistryHostname,
			locals.IngressInternalSchemaRegistryHostname,
		}
		locals.SchemaRegistryKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", vars.SchemaRegistryKubeServiceName, locals.Namespace)
	}

	// kowl related locals data
	if locals.KafkaKubernetes.Spec.IsDeployKafkaUi {

		locals.IngressKowlCertSecretName = fmt.Sprintf("cert-%s-kowl", kafkaKubernetes.Metadata.Id)

		locals.IngressExternalKowlHostname = fmt.Sprintf("%s-kowl.%s", kafkaKubernetes.Metadata.Id, kafkaKubernetes.Spec.Ingress.DnsDomain)

		ctx.Export(outputs.KafkaUiExternalUrl, pulumi.Sprintf("https://%s", locals.IngressExternalKowlHostname))

		locals.KowlKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", vars.KowlKubeServiceName, locals.Namespace)
	}

	if kafkaKubernetes.Spec.Ingress == nil ||
		!kafkaKubernetes.Spec.Ingress.IsEnabled ||
		kafkaKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalBootstrapHostname = fmt.Sprintf("%s-bootstrap.%s", kafkaKubernetes.Metadata.Id, kafkaKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalBootstrapHostname = fmt.Sprintf("%s-bootstrap-internal.%s", kafkaKubernetes.Metadata.Id, kafkaKubernetes.Spec.Ingress.DnsDomain)

	ctx.Export(outputs.BootstrapServerExternalHostname, pulumi.String(locals.IngressExternalBootstrapHostname))
	ctx.Export(outputs.BootstrapServerInternalHostname, pulumi.String(locals.IngressInternalBootstrapHostname))

	// Creating internal broker hostnames
	ingressInternalBrokerHostnames := make([]string, int(kafkaKubernetes.Spec.BrokerContainer.Replicas))
	for i := 0; i < int(kafkaKubernetes.Spec.BrokerContainer.Replicas); i++ {
		ingressInternalBrokerHostnames[i] = fmt.Sprintf("%s-broker-b%d-internal.%s", kafkaKubernetes.Metadata.Id, i, kafkaKubernetes.Spec.Ingress.DnsDomain)
	}
	locals.IngressInternalBrokerHostnames = ingressInternalBrokerHostnames

	// Creating external broker hostnames
	ingressExternalBrokerHostnames := make([]string, int(kafkaKubernetes.Spec.BrokerContainer.Replicas))
	for i := 0; i < int(kafkaKubernetes.Spec.BrokerContainer.Replicas); i++ {
		ingressExternalBrokerHostnames[i] = fmt.Sprintf("%s-broker-b%d.%s", kafkaKubernetes.Metadata.Id, i, kafkaKubernetes.Spec.Ingress.DnsDomain)
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
	//ctx.Export(outputs.IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	//ctx.Export(outputs.IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = kafkaKubernetes.Spec.Ingress.DnsDomain

	switch stackInput.ProviderCredential.Provider {
	case kubernetesclustercredentialv1.KubernetesProvider_gcp_gke:
		locals.KafkaIngressPrivateListenerLoadBalancerAnnotationKey = "cloud.google.com/load-balancer-type"
		locals.KafkaIngressPrivateListenerLoadBalancerAnnotationValue = "Internal"
		locals.KafkaIngressPublicListenerLoadBalancerAnnotationKey = "cloud.google.com/load-balancer-type"
		locals.KafkaIngressPublicListenerLoadBalancerAnnotationValue = "External"
	}

	return locals
}
