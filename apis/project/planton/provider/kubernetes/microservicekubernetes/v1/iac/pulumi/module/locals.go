package module

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	dockercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/dockercredential/v1"
	microservicekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/microservicekubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/microservicekubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/pulmod/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	IngressExternalHostname      string
	IngressHostnames             []string
	IngressInternalHostname      string
	KubePortForwardCommand       string
	KubeServiceFqdn              string
	KubeServiceName              string
	Namespace                    string
	MicroserviceKubernetes       *microservicekubernetesv1.MicroserviceKubernetes
	ImagePullSecretData          map[string]string
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *microservicekubernetesv1.MicroserviceKubernetesStackInput) (*Locals, error) {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	microserviceKubernetes := stackInput.Target

	//assign value for the locals variable to make it available across the project
	locals.MicroserviceKubernetes = stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   stackInput.Target.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: "microservice_kubernetes",
	}

	if microserviceKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = microserviceKubernetes.Metadata.Org
	}

	if microserviceKubernetes.Metadata.Env != nil {
		locals.Labels[kuberneteslabelkeys.Environment] = microserviceKubernetes.Metadata.Env.Id
	}

	if stackInput.DockerCredential != nil &&
		dockercredentialv1.DockerRepoProvider_gcp_artifact_registry == stackInput.DockerCredential.Provider {
		decodedStringBytes, err := b64.StdEncoding.DecodeString(stackInput.DockerCredential.GcpArtifactRegistry.ServiceAccountKeyBase64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode gcp service account key base64")
		}
		dockerConfigAuth := fmt.Sprintf("_json_key:%s", string(decodedStringBytes))

		dockerConfigAuth = b64.StdEncoding.EncodeToString([]byte(dockerConfigAuth))

		dockerRepoHostname := fmt.Sprintf("%s-docker.pkg.dev", stackInput.DockerCredential.GcpArtifactRegistry.GcpRegion)

		locals.ImagePullSecretData = map[string]string{".dockerconfigjson": fmt.Sprintf(`
			{
  				"auths": {
    				"%s": {
      					"username": "_json_key",
						"auth": "%s"
					}
  				}
			}`, dockerRepoHostname, dockerConfigAuth)}
	}

	//decide on the namespace
	locals.Namespace = microserviceKubernetes.Metadata.Id

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.KubeServiceName = microserviceKubernetes.Spec.Version

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.KubePortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if microserviceKubernetes.Spec.Ingress == nil ||
		!microserviceKubernetes.Spec.Ingress.IsEnabled ||
		microserviceKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals, nil
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", microserviceKubernetes.Metadata.Id,
		microserviceKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", microserviceKubernetes.Metadata.Id,
		microserviceKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	//export ingress hostnames
	ctx.Export(outputs.IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = microserviceKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = microserviceKubernetes.Metadata.Id

	return locals, nil
}
