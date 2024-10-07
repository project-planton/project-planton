package manifest

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/manifest/manifestprotobuf"
	"github.com/plantoncloud/project-planton/internal/ulidgen"
	"github.com/plantoncloud/project-planton/internal/workspace"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func LoadWithOverrides(manifestPath string, valueOverrides map[string]string) (proto.Message, error) {
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

	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load manifest")
	}
	for key, value := range valueOverrides {
		manifest, err = manifestprotobuf.SetProtoField(manifest, key, value)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to set %s=%s", key, value)
		}
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
