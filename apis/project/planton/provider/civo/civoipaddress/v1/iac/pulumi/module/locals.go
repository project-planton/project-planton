package module

import (
	civocredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/civocredential/v1"
	civoipaddressv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civoipaddress/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals holds frequently accessed values so that downstream functions
// read like Terraform locals rather than deeply nested getters.
type Locals struct {
	CivoCredentialSpec *civocredentialv1.CivoCredentialSpec
	CivoIpAddress      *civoipaddressv1.CivoIpAddress
}

// initializeLocals prepares the struct in the simplest possible way.
func initializeLocals(_ *pulumi.Context, stackInput *civoipaddressv1.CivoIpAddressStackInput) *Locals {
	locals := &Locals{}
	locals.CivoIpAddress = stackInput.Target
	locals.CivoCredentialSpec = stackInput.ProviderCredential
	return locals
}
