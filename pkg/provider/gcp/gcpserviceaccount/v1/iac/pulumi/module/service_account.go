package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// serviceAccount provisions the account and (optionally) a key.
// Returns the created account and key pointers (key may be nil).
func serviceAccount(
	ctx *pulumi.Context,
	locals *Locals,
) (*serviceaccount.Account, *serviceaccount.Key, error) {

	// Build arguments for the account.
	accountArgs := &serviceaccount.AccountArgs{
		AccountId:   pulumi.String(locals.GcpServiceAccount.Spec.ServiceAccountId),
		DisplayName: pulumi.String(locals.GcpServiceAccount.Metadata.Name),
	}

	// project_id is optional in the spec.
	if locals.GcpServiceAccount.Spec.ProjectId != "" {
		accountArgs.Project = pulumi.StringPtr(locals.GcpServiceAccount.Spec.ProjectId)
	}

	// Create the service account.
	createdServiceAccount, err := serviceaccount.NewAccount(
		ctx,
		locals.GcpServiceAccount.Metadata.Name,
		accountArgs,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create service account")
	}

	// Optionally create a key.
	var createdKey *serviceaccount.Key
	if locals.GcpServiceAccount.Spec.GetCreateKey() {
		createdKey, err = serviceaccount.NewKey(
			ctx,
			fmt.Sprintf("%s-key", locals.GcpServiceAccount.Metadata.Name),
			&serviceaccount.KeyArgs{
				ServiceAccountId: createdServiceAccount.Name,
			},
			pulumi.DependsOn([]pulumi.Resource{createdServiceAccount}),
		)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create service account key")
		}
	}

	return createdServiceAccount, createdKey, nil
}
