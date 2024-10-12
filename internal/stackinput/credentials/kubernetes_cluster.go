package credentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	kubernetesClusterKey  = "kubernetesCluster"
	kubernetesClusterYaml = "kubernetes-cluster.yaml"
)

func AddKubernetesCluster(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.KubernetesCluster != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.KubernetesCluster)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.KubernetesCluster)
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
