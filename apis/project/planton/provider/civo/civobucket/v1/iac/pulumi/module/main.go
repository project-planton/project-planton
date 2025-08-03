package module

import (
	civobucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civobucket/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civobucketv1.CivoBucketStackInput,
) error {
	return nil
}
