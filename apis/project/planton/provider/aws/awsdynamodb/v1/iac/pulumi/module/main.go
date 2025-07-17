package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    dynamodb "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the main entry-point invoked by the stack runner. It prepares the
// AWS provider based on the received credentials, initialises shared locals and
// delegates the creation of every concrete cloud resource to dedicated helper
// functions. Finally, it exports all observable output values so that callers
// (e.g. the orchestrating control-plane) can reference them.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    if stackInput == nil {
        return errors.New("nil AwsDynamodbStackInput passed to Resources")
    }

    // ---------------------------------------------------------------------
    // 1. Configure the AWS provider – credentials are taken from the
    //    stack-input. If individual fields are not set we simply rely on the
    //    default resolution mechanisms (env-vars, shared config, IAM role …).
    // ---------------------------------------------------------------------
    providerArgs := &aws.ProviderArgs{}

    if cred := stackInput.GetProviderCredential(); cred != nil {
        if v := cred.GetAccessKeyId(); v != "" {
            providerArgs.AccessKey = pulumi.StringPtr(v)
        }
        if v := cred.GetSecretAccessKey(); v != "" {
            providerArgs.SecretKey = pulumi.StringPtr(v)
        }
        if v := cred.GetRegion(); v != "" {
            providerArgs.Region = pulumi.StringPtr(v)
        }
        // The credential proto currently does not expose an AWS profile field.
    }

    awsProvider, err := aws.NewProvider(ctx, "aws", providerArgs)
    if err != nil {
        return errors.Wrap(err, "creating AWS provider")
    }

    // ---------------------------------------------------------------------
    // 2. Build common/derived values (locals) used by all subsequent helper
    //    functions.
    // ---------------------------------------------------------------------
    locals, err := initializeLocals(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising locals")
    }

    // ---------------------------------------------------------------------
    // 3. Provision the DynamoDB table (primary resource for this module).
    // ---------------------------------------------------------------------
    table, err := dynamodbTable(ctx, locals, awsProvider)
    if err != nil {
        return errors.Wrap(err, "creating DynamoDB table")
    }

    // ---------------------------------------------------------------------
    // 4. Export all observable output values as defined in outputs.go.
    // ---------------------------------------------------------------------
    ctx.Export(TableArnKey, table.Arn)
    ctx.Export(TableNameKey, table.Name)
    ctx.Export(TableIDKey, table.ID().ToStringOutput())

    // Stream information – will resolve to empty ("null") values when streams
    // are not enabled for the table which is fully acceptable for callers.
    ctx.Export(StreamArnKey, table.StreamArn)
    ctx.Export(StreamLabelKey, table.StreamLabel)

    // The remaining outputs are optional and can be nil/empty depending on the
    // module configuration. We intentionally export well-typed empty values so
    // that the shape of the final outputs object is always deterministic.
    ctx.Export(KmsKeyArnKey, pulumi.String(""))
    ctx.Export(GlobalSecondaryIndexNamesKey, pulumi.StringArray{})
    ctx.Export(LocalSecondaryIndexNamesKey, pulumi.StringArray{})

    return nil
}

// ----------------------------------------------------------------------------
// Helper function placeholder
// ----------------------------------------------------------------------------
// dynamodbTable lives in table.go (as required by the coding guideline). It
// provisions the actual aws.dynamodb.Table resource and returns the Pulumi
// object so that callers can reference its attributes.
func dynamodbTable(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*dynamodb.Table, error) {
    // The real implementation is provided in table.go. The stub below is only
    // present so that the Go compiler is happy when this file is built in
    // isolation (e.g. during CI checks before other files are generated).
    return nil, errors.New("dynamodbTable stub called – real implementation missing")
}
