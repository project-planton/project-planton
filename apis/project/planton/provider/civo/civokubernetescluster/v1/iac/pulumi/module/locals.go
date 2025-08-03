package module

import (
	civocredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/civocredential/v1"
	civokubernetesclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civokubernetescluster/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	CivoCredentialSpec    *civocredentialv1.CivoCredentialSpec
	CivoKubernetesCluster *civokubernetesclusterv1.CivoKubernetesCluster
}

// initializeLocals copies stack‑input fields into the Locals struct.
func initializeLocals(_ *pulumi.Context, stackInput *civokubernetesclusterv1.CivoKubernetesClusterStackInput) *Locals {
	return &Locals{
		CivoCredentialSpec:    stackInput.ProviderCredential,
		CivoKubernetesCluster: stackInput.Target,
	}
}
