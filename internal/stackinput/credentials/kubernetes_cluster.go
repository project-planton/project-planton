package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"os"
)

const (
	kubernetesClusterKey  = "kubernetesCluster"
	kubernetesClusterYaml = "kubernetes-cluster.yaml"
)

func AddKubernetesCluster(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.KubernetesCluster != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.KubernetesCluster)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.KubernetesCluster)
		}
		stackInputContentMap[kubernetesClusterKey] = string(credentialContent)
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
