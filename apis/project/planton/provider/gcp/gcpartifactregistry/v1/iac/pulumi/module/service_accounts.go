package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistry/v1/iac/pulumi/module/outputs"
	pulumigcp "github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func serviceAccounts(ctx *pulumi.Context, locals *Locals, gcpProvider *pulumigcp.Provider) (createdReaderServiceAccount,
	createdWriterServiceAccount *serviceaccount.Account, err error) {

	createdServiceAccountSuffixRandomString, err := random.NewRandomString(ctx, "service-account-suffix",
		&random.RandomStringArgs{
			Special: pulumi.Bool(false),
			Lower:   pulumi.Bool(true),
			Upper:   pulumi.Bool(false),
			Number:  pulumi.Bool(true),
			Length:  pulumi.Int(6), //increasing this can result in violation of service account id length <30
		})
	if err != nil {
		return nil, nil,
			errors.Wrap(err, "failed to create random suffix for service account")
	}

	//create a name for the google service account to be used for "read"
	//operations on the artifact-registry repositories.
	readerServiceAccountName := pulumi.Sprintf("%s-%s-ro", locals.GcpArtifactRegistry.Metadata.Name,
		createdServiceAccountSuffixRandomString.Result)

	//create google service account to be used for "read"
	//operations on the artifact-registry repositories.
	createdReaderServiceAccount, err = serviceaccount.NewAccount(ctx,
		"reader-service-account",
		&serviceaccount.AccountArgs{
			Project:     pulumi.String(locals.GcpArtifactRegistry.Spec.ProjectId),
			AccountId:   readerServiceAccountName,
			DisplayName: readerServiceAccountName,
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return nil, nil,
			errors.Wrap(err, "failed create new reader service account")
	}

	//create a json credentials key for the google service account to be used for "read"
	//operations on the artifact-registry repositories.
	createdReaderServiceAccountKey, err := serviceaccount.NewKey(ctx,
		"reader-service-account",
		&serviceaccount.KeyArgs{
			ServiceAccountId: createdReaderServiceAccount.Name,
			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
		}, pulumi.Parent(createdReaderServiceAccount))
	if err != nil {
		return nil, nil, errors.Wrap(err,
			"failed to create json key for reader service account")
	}

	//export outputs for email and private key as outputs for the "reader" service account
	ctx.Export(outputs.ReaderServiceAccountEmail, createdReaderServiceAccount.Email)
	ctx.Export(outputs.ReaderServiceAccountKey, createdReaderServiceAccountKey.PrivateKey)

	//create a name for the google service account to be used for "write"
	//operations on the artifact-registry repositories.
	writerServiceAccountName := pulumi.Sprintf("%s-%s-rw", locals.GcpArtifactRegistry.Metadata.Name,
		createdServiceAccountSuffixRandomString.Result)

	//create google service account to be used for "write"
	//operations on the artifact-registry repositories.
	createdWriterServiceAccount, err = serviceaccount.NewAccount(ctx,
		"writer-service-account",
		&serviceaccount.AccountArgs{
			Project:     pulumi.String(locals.GcpArtifactRegistry.Spec.ProjectId),
			AccountId:   writerServiceAccountName,
			DisplayName: writerServiceAccountName,
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return nil, nil,
			errors.Wrap(err, "failed create new writer service account")
	}

	//create a json credentials key for the google service account to be used for "write"
	//operations on the artifact-registry repositories.
	createdWriterServiceAccountKey, err := serviceaccount.NewKey(ctx,
		"writer-service-account",
		&serviceaccount.KeyArgs{
			ServiceAccountId: createdWriterServiceAccount.Name,
			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
		}, pulumi.Parent(createdWriterServiceAccount))
	if err != nil {
		return nil, nil, errors.Wrap(err,
			"failed to create json key for writer service account")
	}

	//export outputs for email and private key as outputs for the "writer" service account
	ctx.Export(outputs.WriterServiceAccountEmail, createdWriterServiceAccount.Email)
	ctx.Export(outputs.WriterServiceAccountKey, createdWriterServiceAccountKey.PrivateKey)

	return createdReaderServiceAccount, createdWriterServiceAccount, nil
}
