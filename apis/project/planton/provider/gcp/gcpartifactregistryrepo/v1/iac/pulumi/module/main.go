package module

import (
	"github.com/pkg/errors"
	gcpartifactregistryrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistryrepo/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpartifactregistryrepov1.GcpArtifactRegistryRepoStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create google provider using the credentials from the input
	googleProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to create google provider")
	}

	//create reader and writer service accounts and their keys
	addedReaderServiceAccount, addedWriterServiceAccount, err := serviceAccounts(ctx, locals, googleProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create service accounts")
	}

	//create docker repository and grant reader and writer roles for the service accounts on the repo
	if err := repo(ctx, locals, googleProvider, addedReaderServiceAccount, addedWriterServiceAccount); err != nil {
		return errors.Wrap(err, "failed to create docker repo")
	}

	return nil
}
