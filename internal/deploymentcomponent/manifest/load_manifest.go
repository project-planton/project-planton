package manifest

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/pkg/ulidgen"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

func LoadManifest(manifestPath string) (proto.Message, error) {
	isUrl, err := isManifestPathUrl(manifestPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to determine if manifest path is url")
	}

	if isUrl {
		manifestPath, err = downloadManifest(manifestPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to download manifest using %s", manifestPath)
		}
	}

	manifestYamlBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read manifest file %s", manifestPath)
	}

	jsonBytes, err := yaml.YAMLToJSON(manifestYamlBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load yaml to json")
	}

	kindName, err := ExtractKindFromTargetManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to extract kind from %s stack input yaml", manifestPath)
	}

	manifest := DeploymentComponentMap[FindMatchingComponent(ConvertKindName(kindName))]

	if manifest == nil {
		return nil, errors.Errorf("deployment-component does not contain %s", ConvertKindName(kindName))
	}

	if err := protojson.Unmarshal(jsonBytes, manifest); err != nil {
		return nil, errors.Wrapf(err, "failed to load json into proto message from %s", manifestPath)
	}
	return manifest, nil
}

func downloadManifest(manifestUrl string) (string, error) {
	// Get the directory to save the downloaded file
	dir, err := workspace.GetManifestDownloadDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get manifest download directory")
	}

	// Generate a new ulid for the file name
	fileName := ulidgen.NewGenerator().Generate().String() + ".yaml"

	filePath := filepath.Join(dir, fileName)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to create file")
	}
	defer out.Close()

	// Download the file
	resp, err := http.Get(manifestUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to download manifest from %s", manifestUrl)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to write manifest to file")
	}

	// Return the absolute path of the downloaded file
	return filePath, nil
}

func isManifestPathUrl(manifestPath string) (bool, error) {
	// Attempt to parse the manifestPath as a URL
	parsedUrl, err := url.Parse(manifestPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to parse manifest path as URL")
	}

	// Check if the URL has a scheme and host
	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return false, nil
	}

	return true, nil
}
