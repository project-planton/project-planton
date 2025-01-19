package fileutil

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"text/template"
)

func IsExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrapf(err, "failed to lookup Pulumi.yaml file")
	}
	return true, nil
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
