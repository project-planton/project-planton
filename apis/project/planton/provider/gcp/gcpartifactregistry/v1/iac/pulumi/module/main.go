package module

import (
	"github.com/pkg/errors"
	gcpartifactregistryv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistry/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpartifactregistryv1.GcpArtifactRegistryStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create google provider using the credentials from the input
	googleProvider, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to create google provider")
	}

	//create reader and writer service accounts and their keys
	addedReaderServiceAccount, addedWriterServiceAccount, err := serviceAccounts(ctx, locals, googleProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create service accounts")
	}

	//create docker repository and grant reader and writer roles for the service accounts on the repo
	if err := dockerRepo(ctx, locals, googleProvider, addedReaderServiceAccount, addedWriterServiceAccount); err != nil {
		return errors.Wrap(err, "failed to create docker repo")
	}

	//create maven repository and grant reader and writer roles for the service accounts on the repo
	if err := mavenRepo(ctx, locals, googleProvider, addedReaderServiceAccount, addedWriterServiceAccount); err != nil {
		return errors.Wrap(err, "failed to create maven repo")
	}

	//create npm repository and grant reader and writer roles for the service accounts on the repo
	if err := npmRepo(ctx, locals, googleProvider, addedReaderServiceAccount, addedWriterServiceAccount); err != nil {
		return errors.Wrap(err, "failed to create npm repo")
	}

	//create python repository and grant reader and writer roles for the service accounts on the repo
	if err := pythonRepo(ctx, locals, googleProvider, addedReaderServiceAccount, addedWriterServiceAccount); err != nil {
		return errors.Wrap(err, "failed to create python repo")
	}

	return nil
}
