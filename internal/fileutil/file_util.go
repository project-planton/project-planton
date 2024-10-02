package fileutil

import (
	"github.com/pkg/errors"
	"os"
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
