package module

import (
	civocredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/civocredential/v1"
	civocomputeinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civocomputeinstance/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles pointers frequently reused across the module.
type Locals struct {
	CivoCredentialSpec  *civocredentialv1.CivoCredentialSpec
	CivoComputeInstance *civocomputeinstancev1.CivoComputeInstance
}

// initializeLocals mirrors the simple pattern used in other Planton modules.
func initializeLocals(
	_ *pulumi.Context,
	stackInput *civocomputeinstancev1.CivoComputeInstanceStackInput,
) *Locals {
	return &Locals{
		CivoComputeInstance: stackInput.Target,
		CivoCredentialSpec:  stackInput.ProviderCredential,
	}
}
