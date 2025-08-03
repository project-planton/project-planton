package module

import (
	civokubernetesclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civokubernetescluster/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civokubernetesclusterv1.CivoKubernetesClusterStackInput,
) error {
	return nil
}
