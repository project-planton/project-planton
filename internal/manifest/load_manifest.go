package manifest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/ulidgen"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

	kindName, err := crkreflect.ExtractKindFromTargetManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to extract cloudResourceKind from %s stack input yaml", manifestPath)
	}

	cloudResourceKind := crkreflect.KindFromString(kindName)

	manifest := crkreflect.ToMessageMap[cloudResourceKind]

	if manifest == nil {
		return nil, formatUnsupportedResourceError(kindName)
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

// formatUnsupportedResourceError creates a helpful error message when a cloud resource kind is not supported
func formatUnsupportedResourceError(kindName string) error {
	// Create colored output functions
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	var msg strings.Builder

	msg.WriteString("\n")
	msg.WriteString(red("╔═══════════════════════════════════════════════════════════════════════════════╗") + "\n")
	msg.WriteString(red("║") + bold("                ⚠️  UNSUPPORTED CLOUD RESOURCE KIND                           ") + red("║") + "\n")
	msg.WriteString(red("╚═══════════════════════════════════════════════════════════════════════════════╝") + "\n\n")

	msg.WriteString(yellow("Resource Kind:") + " " + bold(kindName) + "\n\n")

	msg.WriteString(red("❌ This cloud resource kind is not recognized.\n\n"))

	msg.WriteString(cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"))
	msg.WriteString(bold("                           🔧 HOW TO FIX\n"))
	msg.WriteString(cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"))

	msg.WriteString(yellow("1. Check your manifest for typos in the 'kind' field\n\n"))
	msg.WriteString("   Common mistakes:\n")
	msg.WriteString("   • Extra characters (e.g., 'AwsEksCluster" + bold("s") + "')\n")
	msg.WriteString("   • Wrong capitalization (e.g., 'Aws" + bold("EKS") + "Cluster')\n")
	msg.WriteString("   • Misspelled resource name (e.g., 'AwsEks" + bold("Clster") + "')\n\n")

	msg.WriteString(yellow("2. If the kind is correct, update your CLI to the latest version:\n\n"))
	msg.WriteString("   " + green("brew update && brew upgrade project-planton") + "\n\n")
	msg.WriteString("   Or if you haven't installed via Homebrew:\n\n")
	msg.WriteString("   " + green("brew install project-planton/tap/project-planton") + "\n\n")
	msg.WriteString("   Then verify:\n\n")
	msg.WriteString("   " + green("project-planton version") + "\n\n")

	msg.WriteString(yellow("3. Retry your command\n\n"))

	msg.WriteString(cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"))

	msg.WriteString(fmt.Sprintf(bold("💡 TIP: ") + "If you're developing a new cloud resource, ensure the proto files\n"))
	msg.WriteString(fmt.Sprintf("   are compiled and the CLI binary is rebuilt.\n\n"))

	return errors.New(msg.String())
}
