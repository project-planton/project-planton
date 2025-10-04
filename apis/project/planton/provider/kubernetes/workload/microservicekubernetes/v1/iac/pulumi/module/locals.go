package module

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	microservicekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/microservicekubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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

	locals.MicroserviceKubernetes = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes.String(),
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

	// Priority order:
	// 1. Default: metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// Priority 1: StackInput (used by Planton Cloud - takes precedence)
	// If present, use it and don't check the label at all
	if stackInput.DockerConfigJson != "" {
		locals.ImagePullSecretData = map[string]string{".dockerconfigjson": stackInput.DockerConfigJson}
	} else {
		// Priority 2: Label with file path (for open-source users)
		// Only checked if stackInput.DockerConfigJson is empty
		if dockerConfigFilePath := target.Metadata.Labels[kuberneteslabels.DockerConfigJsonFileLabelKey]; dockerConfigFilePath != "" {
			dockerConfigJson, err := loadDockerConfigFromFile(dockerConfigFilePath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to load docker config from file specified in label: %s", dockerConfigFilePath)
			}
			locals.ImagePullSecretData = map[string]string{".dockerconfigjson": dockerConfigJson}
		}
		// Priority 3: If neither set, ImagePullSecretData remains nil (no image pull secret)
	}

	locals.KubeServiceName = target.Spec.Version

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
		return locals, nil
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	//export ingress hostnames
	ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(OpInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = target.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = locals.Namespace

	if locals.MicroserviceKubernetes.Spec.Container.App.Image == nil {
		return nil, errors.New("spec.container.app.image is required")
	}

	if locals.MicroserviceKubernetes.Spec.Availability == nil {
		locals.MicroserviceKubernetes.Spec.Availability = &microservicekubernetesv1.MicroserviceKubernetesAvailability{
			MinReplicas: 1,
		}
	}

	return locals, nil
}

// loadDockerConfigFromFile reads docker config JSON from the specified file path.
// Returns error if file doesn't exist or can't be read.
func loadDockerConfigFromFile(filePath string) (string, error) {
	// Expand ~ to home directory if present
	if strings.HasPrefix(filePath, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "failed to get user home directory")
		}
		filePath = filepath.Join(homeDir, filePath[2:])
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.Errorf("docker config file does not exist: %s", filePath)
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read docker config file: %s", filePath)
	}

	// Validate it's not empty
	if len(content) == 0 {
		return "", errors.Errorf("docker config file is empty: %s", filePath)
	}

	// Optional: Basic JSON validation
	var js json.RawMessage
	if err := json.Unmarshal(content, &js); err != nil {
		return "", errors.Wrapf(err, "docker config file contains invalid JSON: %s", filePath)
	}

	return string(content), nil
}
