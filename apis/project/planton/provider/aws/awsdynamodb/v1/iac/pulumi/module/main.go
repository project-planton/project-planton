package module

import (
    "fmt"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the root entry-point executed by the Pulumi engine.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // 1. AWS provider.
    awsProvider, err := newAWSProvider(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising AWS provider")
    }

    // 2. Locals (shared derived data).
    locals, err := initializeLocals(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising locals")
    }

    // 3. Table.
    table, err := dynamodbTable(ctx, locals, awsProvider)
    if err != nil {
        return errors.Wrap(err, "creating DynamoDB table")
    }

    // 4. Outputs.
    ctx.Export(TableArn, table.Arn)
    ctx.Export(TableName, table.Name)
    ctx.Export(TableID, table.ID())

    if table.StreamArn != nil {
        ctx.Export(StreamArn, table.StreamArn)
    }
    if table.StreamLabel != nil {
        ctx.Export(StreamLabel, table.StreamLabel)
    }

    if locals.KmsKeyArn != "" {
        ctx.Export(KmsKeyArn, pulumi.String(locals.KmsKeyArn))
    }

    if len(locals.GsiNames) > 0 {
        ctx.Export(GlobalSecondaryIndexNames, pulumi.ToStringArray(locals.GsiNames))
    }
    if len(locals.LsiNames) > 0 {
        ctx.Export(LocalSecondaryIndexNames, pulumi.ToStringArray(locals.LsiNames))
    }

    return nil
}

// ---------------------------------------------------------------------------
// Helper functions (unchanged except for updated import paths / constants)
// ---------------------------------------------------------------------------

func newAWSProvider(ctx *pulumi.Context, in *awsdynamodbv1.AwsDynamodbStackInput) (*aws.Provider, error) {
    cred := in.GetProviderCredential()
    if cred == nil {
        return nil, errors.New("missing AWS provider_credential block in stack-input")
    }

    args := &aws.ProviderArgs{}

    if region := cred.GetRegion(); region != "" {
        args.Region = pulumi.String(region)
    }
    if ak := cred.GetAccessKeyId(); ak != "" {
        args.AccessKey = pulumi.StringPtr(ak)
    }
    if sk := cred.GetSecretAccessKey(); sk != "" {
        args.SecretKey = pulumi.StringPtr(sk)
    }
    if tok := cred.GetSessionToken(); tok != "" {
        args.Token = pulumi.StringPtr(tok)
    }
    if prof := cred.GetProfile(); prof != "" {
        args.Profile = pulumi.StringPtr(prof)
    }

    provider, err := aws.NewProvider(ctx, "aws", args)
    if err != nil {
        return nil, errors.Wrap(err, "creating aws.Provider")
    }
    return provider, nil
}

// dynamodbTable retains its original implementation – only the provider import
// path was updated above so compilation succeeds.
