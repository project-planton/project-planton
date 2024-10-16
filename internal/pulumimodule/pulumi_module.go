package pulumimodule

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"strings"
)

const (
	DownloadUrl = "https://raw.githubusercontent.com/project-planton/project-planton/6f050fbdff8580150cf396da581b1566a198531e/module-git-repos.yaml"
)

type DefaultPulumiModules struct {
	Atlas      map[string]string `yaml:"atlas"`
	Aws        map[string]string `yaml:"aws"`
	Confluent  map[string]string `yaml:"confluent"`
	Gcp        map[string]string `yaml:"gcp"`
	Kubernetes map[string]string `yaml:"kubernetes"`
	Snowflake  map[string]string `yaml:"snowflake"`
}

func GetCloneUrl(kindName string) (string, error) {
	defaultModules, err := downloadModuleInfo(DownloadUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to download module info")
	}
	cloneUrl, err := getCloneUrlFromModules(defaultModules, kindName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get clone url for the kind name")
	}
	return cloneUrl, nil
}

func downloadModuleInfo(url string) (DefaultPulumiModules, error) {
	resp, err := http.Get(url)
	if err != nil {
		return DefaultPulumiModules{}, errors.Wrap(err, "failed to fetch yaml")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DefaultPulumiModules{}, errors.Wrap(err, "failed to read response body")
	}

	var modules DefaultPulumiModules
	err = yaml.Unmarshal(body, &modules)
	if err != nil {
		return DefaultPulumiModules{}, errors.Wrapf(err, "failed to unmarshal yaml")
	}

	return modules, nil
}

func normalizeString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(strings.TrimSpace(s)), "_", ""), "-", "")
}

func getCloneUrlFromModules(modules DefaultPulumiModules, kindName string) (string, error) {
	normalizedKindName := normalizeString(kindName)

	for _, moduleMap := range []map[string]string{
		modules.Atlas,
		modules.Aws,
		modules.Confluent,
		modules.Gcp,
		modules.Kubernetes,
		modules.Snowflake,
	} {
		for key, url := range moduleMap {
			if normalizeString(key) == normalizedKindName {
				return url, nil
			}
		}
	}

	return "", errors.New("clone url not found for the provided kind name")
}
