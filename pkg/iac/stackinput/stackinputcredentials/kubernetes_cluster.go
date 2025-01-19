package stackinputcredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	KubernetesClusterKey  = "kubernetesCluster"
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
		stackInputContentMap[KubernetesClusterKey] = credentialContentMap
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
