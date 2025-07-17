package module

import (
    "github.com/pkg/errors"
    awsddb "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// stream extracts DynamoDB Streams related information from the provided
// aws.dynamodb.Table resource and returns Pulumi outputs for the latest
// stream ARN and stream label respectively. No additional Pulumi resources
// are required because Streams are configured directly on the table.
//
// When Streams are disabled, the returned outputs will resolve to undefined
// (nil) values which is exactly what the AWS provider exposes. The caller
// can decide whether or not to export these values.
func stream(
    ctx *pulumi.Context,
    locals *Locals,
    table *awsddb.Table,
) (pulumi.StringPtrOutput, pulumi.StringPtrOutput, error) {

    if ctx == nil {
        return pulumi.StringPtrOutput{}, pulumi.StringPtrOutput{}, errors.New("stream: nil Pulumi context")
    }
    if locals == nil {
        return pulumi.StringPtrOutput{}, pulumi.StringPtrOutput{}, errors.New("stream: nil locals")
    }
    if table == nil {
        return pulumi.StringPtrOutput{}, pulumi.StringPtrOutput{}, errors.New("stream: nil DynamoDB table reference")
    }

    // Inspect the requested StreamSpecification. Even though the presence or
    // absence of Streams does not change the way we read outputs, checking the
    // flag allows the function to stay future-proof (e.g. to attach extra
    // permissions when Streams are enabled).
    if spec := locals.Target.GetSpec().GetStreamSpecification(); spec != nil {
        if !spec.GetStreamEnabled() {
            // Streams explicitly disabled.
            return table.StreamArn, table.StreamLabel, nil
        }
        // Streams enabled – fall-through and return the outputs.
        return table.StreamArn, table.StreamLabel, nil
    }

    // No stream specification provided – treat as disabled; return whatever
    // the provider gives us (expected to be nil).
    return table.StreamArn, table.StreamLabel, nil
}
