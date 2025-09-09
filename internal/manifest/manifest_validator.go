package manifest

import (
	"buf.build/go/protovalidate"
	"fmt"
	"github.com/pkg/errors"
)

func Validate(manifestPath string) error {
	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		return errors.Wrap(err, "failed to load manifest")
	}

	spec, err := ExtractSpec(manifest)
	if err != nil {
		return errors.Wrap(err, "failed to extract spec from manifest")
	}

	v, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages(spec),
	)
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	return v.Validate(spec)
}
