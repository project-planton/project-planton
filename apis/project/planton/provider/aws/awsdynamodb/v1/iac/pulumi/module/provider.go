package module

import (
    "strings"

    "github.com/pkg/errors"
    awscredentialpb "github.com/project-planton/project-planton/apis/project/planton/credential/awscredential/v1"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AwsProvider instantiates a Pulumi AWS provider using the given credential
// specification. The returned provider must be supplied to every AWS resource
// via pulumi.Provider(<returned provider>) so that the resources share the same
// configuration and authentication context.
//
// Any missing (empty) optional values are simply omitted from the provider
// arguments so that the Pulumi provider falls back to the usual environment /
// shared-config loading logic (e.g., AWS_PROFILE, default chain, etc.).
func AwsProvider(ctx *pulumi.Context, cred *awscredentialpb.AwsCredentialSpec) (*aws.Provider, error) {
    if cred == nil {
        return nil, errors.New("aws credential spec is nil")
    }

    // Prepare provider arguments, filling only the fields that have been
    // supplied – this avoids overriding any external AWS SDK resolution
    // mechanisms when the caller explicitly leaves a field blank.
    args := &aws.ProviderArgs{}

    if v := strings.TrimSpace(cred.GetRegion()); v != "" {
        args.Region = pulumi.StringPtr(v)
    }
    if v := strings.TrimSpace(cred.GetAccessKeyId()); v != "" {
        args.AccessKey = pulumi.StringPtr(v)
    }
    if v := strings.TrimSpace(cred.GetSecretAccessKey()); v != "" {
        args.SecretKey = pulumi.StringPtr(v)
    }
    // NOTE: The credential proto currently has no field for a session token.
    // The credential proto also does not define a dedicated profile field.

    // Create the provider. A single, deterministic name ("aws") is used so
    // that Pulumi creates only one provider instance per stack, avoiding
    // redundant providers when AwsProvider is called multiple times in a
    // program.
    p, err := aws.NewProvider(ctx, "aws", args)
    if err != nil {
        return nil, errors.Wrap(err, "creating AWS provider")
    }

    return p, nil
}
