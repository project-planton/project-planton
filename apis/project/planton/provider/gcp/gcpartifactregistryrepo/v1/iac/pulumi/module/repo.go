package module

import (
	"fmt"
	pulumigcp "github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// repo creates a repository and also grants reader role to the reader service account and
// writer, admin roles to writer service account.
func repo(ctx *pulumi.Context, locals *Locals, gcpProvider *pulumigcp.Provider,
	readerServiceAccount *serviceaccount.Account, writerServiceAccount *serviceaccount.Account) error {
	//create a variable with descriptive name for the api-resource in the input
	gcpArtifactRegistryRepo := locals.GcpArtifactRegistryRepo

	//todo: might be better to add a random suffix to the repo name to avoid conflicts
	repoName := gcpArtifactRegistryRepo.Metadata.Name

	//create repository
	createdRepo, err := artifactregistry.NewRepository(ctx,
		repoName,
		&artifactregistry.RepositoryArgs{
			Project:      pulumi.String(gcpArtifactRegistryRepo.Spec.ProjectId),
			Location:     pulumi.String(gcpArtifactRegistryRepo.Spec.Region),
			RepositoryId: pulumi.String(gcpArtifactRegistryRepo.Metadata.Name),
			Format:       pulumi.String(gcpArtifactRegistryRepo.Spec.RepoFormat.String()),
			Labels:       pulumi.ToStringMap(locals.GcpLabels),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create repo")
	}

	if gcpArtifactRegistryRepo.Spec.EnablePublicAccess {
		//grant "reader" role for all users on the repo to make it public
		_, err = artifactregistry.NewRepositoryIamMember(ctx,
			fmt.Sprintf("%s-reader-for-all-users", repoName),
			&artifactregistry.RepositoryIamMemberArgs{
				Project:    pulumi.String(gcpArtifactRegistryRepo.Spec.ProjectId),
				Location:   pulumi.String(gcpArtifactRegistryRepo.Spec.Region),
				Repository: createdRepo.RepositoryId,
				Role:       pulumi.String("roles/artifactregistry.reader"),
				//"allUsers" is a special identifier on google identity system which is used
				//for granting permissions to for everyone.
				Member: pulumi.Sprintf("allUsers"),
			}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to grant reader role on repo for reader service account")
		}
	} else {
		//grant "reader" role for the writer service account on the repo
		_, err = artifactregistry.NewRepositoryIamMember(ctx,
			fmt.Sprintf("%s-reader", repoName),
			&artifactregistry.RepositoryIamMemberArgs{
				Project:    pulumi.String(gcpArtifactRegistryRepo.Spec.ProjectId),
				Location:   pulumi.String(gcpArtifactRegistryRepo.Spec.Region),
				Repository: createdRepo.RepositoryId,
				Role:       pulumi.String("roles/artifactregistry.reader"),
				Member:     pulumi.Sprintf("serviceAccount:%s", readerServiceAccount.Email),
			}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to grant reader role on repo for reader service account")
		}
	}

	//grant "writer" role for the writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-writer",
		repoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistryRepo.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistryRepo.Spec.Region),
		Repository: createdRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.writer"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant writer role on repo for writer service account")
	}

	//grant "admin" role for writer service account on the repo
	_, err = artifactregistry.NewRepositoryIamMember(ctx, fmt.Sprintf("%s-admin",
		repoName), &artifactregistry.RepositoryIamMemberArgs{
		Project:    pulumi.String(gcpArtifactRegistryRepo.Spec.ProjectId),
		Location:   pulumi.String(gcpArtifactRegistryRepo.Spec.Region),
		Repository: createdRepo.RepositoryId,
		Role:       pulumi.String("roles/artifactregistry.repoAdmin"),
		Member:     pulumi.Sprintf("serviceAccount:%s", writerServiceAccount.Email),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to grant admin role on repo for writer service account")
	}

	//export important attributes of the repository as outputs
	ctx.Export(OpRepoName, createdRepo.RepositoryId)
	//export docker-repo hostname
	ctx.Export(OpRepoHostname, pulumi.Sprintf(
		"%s-docker.pkg.dev", createdRepo.Location))
	//export complete repo url based on the attributes of the created docker-repo
	// ex: artifactregistry://us-central1-docker.pkg.dev/my-gcp-project-id/my-company-docker-repo
	ctx.Export(OpRepoUrl, pulumi.Sprintf(
		"%s-docker.pkg.dev/%s/%s",
		createdRepo.Location, createdRepo.Project, createdRepo.Name))

	return nil
}
