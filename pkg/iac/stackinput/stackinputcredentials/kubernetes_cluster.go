package stackinputcredentials

import (
	"github.com/pkg/errors"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	kubernetesClusterKey  = "kubernetesCluster"
	kubernetesClusterYaml = "kubernetes-cluster.yaml"
)

func AddKubernetesCluster(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.KubernetesCluster != "" {
		credentialContent, err := os.ReadFile(credentialOptions.KubernetesCluster)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.KubernetesCluster)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[kubernetesClusterKey] = credentialContentMap
	}
	return stackInputContentMap, nil
}

func LoadKubernetesCluster(dir string) (string, error) {
	path := dir + "/" + kubernetesClusterYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}

func GetKubernetesClusterCredential(stackInputContentMap map[string]interface{}) (*kubernetesclustercredentialv1.KubernetesClusterCredentialSpec, error) {
	kubernetesCluster, ok := stackInputContentMap[kubernetesClusterKey]
	if !ok {
		return nil, nil
	}

	kubernetesClusterSpecBytes, err := yaml.Marshal(kubernetesCluster)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal kubernetes cluster content")
	}
	kubernetesClusterCredentialSpec := new(kubernetesclustercredentialv1.KubernetesClusterCredentialSpec)
	err = yaml.Unmarshal(kubernetesClusterSpecBytes, kubernetesClusterCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal kubernetes cluster content")
	}

	return kubernetesClusterCredentialSpec, nil
}
