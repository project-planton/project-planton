package awsdynamodb

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// BuildBillingMode converts the BillingMode enum coming from the user-facing
// protobuf spec into the Pulumi input expected by aws.dynamodb.Table.  The AWS
// provider treats the absence of "billingMode" as "PROVISIONED", therefore we
// only return a concrete value for PAY_PER_REQUEST.
func BuildBillingMode(bm awsdynamodbpb.BillingMode) (pulumi.StringPtrInput, error) {
    switch bm {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        // Returning nil avoids an extraneous diff because the provider already
        // assumes PROVISIONED when the attribute is missing.
        return nil, nil
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        return pulumi.StringPtr("PAY_PER_REQUEST"), nil
    default:
        return nil, fmt.Errorf("unsupported billing mode %v", bm)
    }
}

// BuildTableCapacity populates the ReadCapacity and WriteCapacity attributes on
// TableArgs when the table uses the PROVISIONED billing model. If the table is
// on-demand (PAY_PER_REQUEST) the fields are intentionally left nil so that the
// provider ignores them.
func BuildTableCapacity(args *dynamodb.TableArgs, bm awsdynamodbpb.BillingMode, pt *awsdynamodbpb.ProvisionedThroughput) error {
    if bm == awsdynamodbpb.BillingMode_PROVISIONED {
        if pt == nil {
            return fmt.Errorf("provisioned_throughput must be provided when billing_mode is PROVISIONED")
        }
        args.ReadCapacity = pulumi.IntPtr(int(pt.ReadCapacityUnits))
        args.WriteCapacity = pulumi.IntPtr(int(pt.WriteCapacityUnits))
    }
    return nil
}

// BuildGSIProvisionedThroughput returns the read & write capacity settings that
// have to be applied to a dynamodb.TableGlobalSecondaryIndexArgs instance.  For
// PAY_PER_REQUEST tables the returned inputs are nil because GSIs inherit the
// on-demand billing model.
func BuildGSIProvisionedThroughput(bm awsdynamodbpb.BillingMode, pt *awsdynamodbpb.ProvisionedThroughput) (pulumi.IntPtrInput, pulumi.IntPtrInput, error) {
    if bm == awsdynamodbpb.BillingMode_PROVISIONED {
        if pt == nil {
            return nil, nil, fmt.Errorf("provisioned_throughput must be provided for the global secondary index when billing_mode is PROVISIONED")
        }
        return pulumi.IntPtr(int(pt.ReadCapacityUnits)), pulumi.IntPtr(int(pt.WriteCapacityUnits)), nil
    }

    // PAY_PER_REQUEST – capacities are not applicable.
    return nil, nil, nil
}
