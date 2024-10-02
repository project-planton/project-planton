package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	kubernetesClusterKey = "kubernetesCluster"
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
