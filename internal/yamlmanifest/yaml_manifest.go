package yamlmanifest

import (
	"bytes"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

// ScanApiResourcesYaml reads the input yaml file and loads all the resource definition into an array
func ScanApiResourcesYaml(manifestPath string) ([][]byte, error) {
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	// Create a new YAML decoder
	dec := yaml.NewDecoder(bytes.NewReader(manifestData))

	var docs [][]byte
	var doc map[interface{}]interface{}

	// Decode each document
	for {
		// Decode the next document
		err := dec.Decode(&doc)

		// If the error is EOF, we've reached the end of the stream
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrapf(err, "failed to decode")
		}

		// Marshal the document back into YAML and append it to docs
		docYAML, err := yaml.Marshal(doc)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to yaml marshal decoded resource")
		}

		docs = append(docs, docYAML)
	}

	return docs, nil
}
