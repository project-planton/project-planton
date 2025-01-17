package file

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/util/shell"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// IsFileExists check if a file exists
func IsFileExists(f string) bool {
	if f == "" {
		return false
	}
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsDirExists check if a directory exists
func IsDirExists(d string) bool {
	if d == "" {
		return false
	}
	info, err := os.Stat(d)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		//todo: should return an error instead
		return false
	}
	return info.IsDir()
}

func GetAbsPath(filePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home dir")
	}
	if strings.HasPrefix(filePath, "~/") {
		filePath = filepath.Join(homeDir, filePath[2:])
	}
	return filePath, nil
}

// Unzip unzips file using shell command
// expects unzip package on the os
func Unzip(zipFile, dest string) error {
	unzipCmd := exec.Command("unzip", zipFile)
	unzipCmd.Dir = dest

	if err := shell.RunCmd(unzipCmd); err != nil {
		return errors.Wrap(err, "failed to unzip")
	}
	return nil
}

func RenderTemplate(input interface{}, templateString string) ([]byte, error) {
	log.Debugf("rendering template")
	t := template.New("template").Funcs(template.FuncMap(sprig.FuncMap()))
	t, err := t.Parse(templateString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}
	var renderedBytes bytes.Buffer
	if err := t.Execute(&renderedBytes, input); err != nil {
		return nil, errors.Wrapf(err, "failed to render template")
	}
	return renderedBytes.Bytes(), nil
}

func WriteFile(content []byte, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to create %s directory", filepath.Dir(outputPath))
	}
	err := os.WriteFile(outputPath, content, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to write %s file", outputPath)
	}
	return nil
}

func Download(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
