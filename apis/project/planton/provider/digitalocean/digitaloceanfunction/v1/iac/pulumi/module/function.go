package module

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// function provisions the DigitalOcean Function and exports its outputs.
func function(
	ctx *pulumi.Context,
	locals *Locals,
	doProvider *digitalocean.Provider,
) (*digitalocean.AppSpecFunction, error) {
	return nil, nil
}
