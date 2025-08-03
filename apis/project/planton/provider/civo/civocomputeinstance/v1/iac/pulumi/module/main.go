package module

import (
	civocomputeinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civocomputeinstance/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civocomputeinstancev1.CivoComputeInstanceStackInput,
) error {
	return nil
}
