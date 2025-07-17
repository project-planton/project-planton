package module

import (
    "fmt"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the root entry-point executed by the Pulumi engine. It receives the
// fully-parsed AwsDynamodbStackInput message, wires-up the AWS provider using the
// supplied credentials, prepares local helper data and orchestrates the creation
// of all sub-resources that together implement the requested AwsDynamodb target
// resource.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // ---------------------------------------------------------------------
    // 1. Configure the AWS provider using the credentials embedded in the
    //    StackInput. All subsequently-created AWS resources must reference this
    //    provider instance so that they are executed with the correct account,
    //    region and temporary session information.
    // ---------------------------------------------------------------------
    awsProvider, err := newAWSProvider(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising AWS provider")
    }

    // ---------------------------------------------------------------------
    // 2. Build helper data that is re-used across multiple sub-resources
    //    (naming conventions, base tags, etc.).
    // ---------------------------------------------------------------------
    locals, err := initializeLocals(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising locals")
    }

    // ---------------------------------------------------------------------
    // 3. Create the main DynamoDB table together with all of its optional
    //    components (SSE, TTL, Streams, indexes, …).  The heavy lifting is
    //    delegated to a specialised helper so that the Resources() entry-point
    //    stays concise and readable.
    // ---------------------------------------------------------------------
    table, err := dynamodbTable(ctx, locals, awsProvider)
    if err != nil {
        return errors.Wrap(err, "creating DynamoDB table")
    }

    // ---------------------------------------------------------------------
    // 4. Export every observable identifier requested by the StackOutputs
    //    protocol buffer so that callers (CLI / Controllers / Terraform Remote
    //    State / etc.) can consume them during follow-up reconciliation loops.
    // ---------------------------------------------------------------------
    ctx.Export(TableArn, table.Arn)
    ctx.Export(TableName, table.Name)
    ctx.Export(TableId, table.ID())

    // Streams – only when enabled.
    if table.StreamArn != nil {
        ctx.Export(StreamStreamArn, table.StreamArn)
    }
    if table.StreamLabel != nil {
        ctx.Export(StreamStreamLabel, table.StreamLabel)
    }

    // If a customer-managed KMS CMK has been configured we expose its ARN so
    // that external tooling (e.g. key rotation) can reference it.
    if kmsArn := locals.KmsKeyArn; kmsArn != "" {
        ctx.Export(KmsKeyArn, pulumi.String(kmsArn))
    }

    // Index names are fully known at declaration time – we simply echo them
    // back so that consumers can use them when building queries against the
    // table.
    if len(locals.GsiNames) > 0 {
        ctx.Export(GlobalSecondaryIndexNames, pulumi.ToStringArray(locals.GsiNames))
    }
    if len(locals.LsiNames) > 0 {
        ctx.Export(LocalSecondaryIndexNames, pulumi.ToStringArray(locals.LsiNames))
    }

    return nil
}

// ---------------------------------------------------------------------------
// Helper functions
// ---------------------------------------------------------------------------

// newAWSProvider creates a Pulumi AWS provider that is configured with the
// credentials and region information embedded in the stack-input. All Pulumi
// resources must be associated with this provider so that they inherit the
// correct context.
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

// dynamodbTable provisions the DynamoDB table together with the majority of the
// user-controlled features (capacity/billing model, encryption, streams,
// TTL, tags, …). Complex optional components (e.g. AutoScaling policies) can
// be off-loaded into their own helpers to keep the file readable.
func dynamodbTable(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*dynamodb.Table, error) {
    spec := locals.Target.GetSpec()
    if spec == nil {
        return nil, errors.New("awsdynamodb target.Spec must be set")
    }

    // ---------------------------
    // Attributes & key-schema
    // ---------------------------
    attrs := dynamodb.TableAttributeArray{}
    for _, a := range spec.GetAttributeDefinitions() {
        attrType := "S"
        switch a.GetAttributeType() {
        case awsdynamodbv1.AttributeType_STRING:
            attrType = "S"
        case awsdynamodbv1.AttributeType_NUMBER:
            attrType = "N"
        case awsdynamodbv1.AttributeType_BINARY:
            attrType = "B"
        }
        attrs = append(attrs, dynamodb.TableAttributeArgs{
            Name: pulumi.String(a.GetAttributeName()),
            Type: pulumi.String(attrType),
        })
    }

    var hashKey, rangeKey string
    for _, k := range spec.GetKeySchema() {
        if k.GetKeyType() == awsdynamodbv1.KeyType_HASH {
            hashKey = k.GetAttributeName()
        } else if k.GetKeyType() == awsdynamodbv1.KeyType_RANGE {
            rangeKey = k.GetAttributeName()
        }
    }

    if hashKey == "" {
        return nil, errors.New("primary partition (HASH) key must be defined in key_schema")
    }

    // ---------------------------
    // Base Table arguments
    // ---------------------------
    tArgs := &dynamodb.TableArgs{
        Name:       pulumi.String(spec.GetTableName()),
        Attributes: attrs,
        HashKey:    pulumi.String(hashKey),
        Tags:       pulumi.ToStringMap(locals.Labels),
    }

    if rangeKey != "" {
        tArgs.RangeKey = pulumi.StringPtr(rangeKey)
    }

    // Billing / capacity model.
    switch spec.GetBillingMode() {
    case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
        tArgs.BillingMode = pulumi.StringPtr("PAY_PER_REQUEST")
    case awsdynamodbv1.BillingMode_PROVISIONED:
        tArgs.BillingMode = pulumi.StringPtr("PROVISIONED")
        if pt := spec.GetProvisionedThroughput(); pt != nil {
            tArgs.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
            tArgs.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
        }
    default:
        return nil, errors.Errorf("unsupported billing mode %v", spec.GetBillingMode())
    }

    // Streams.
    if s := spec.GetStreamSpecification(); s != nil && s.GetStreamEnabled() {
        tArgs.StreamEnabled = pulumi.BoolPtr(true)
        viewType := "NEW_AND_OLD_IMAGES" // sane default
        switch s.GetStreamViewType() {
        case awsdynamodbv1.StreamViewType_NEW_IMAGE:
            viewType = "NEW_IMAGE"
        case awsdynamodbv1.StreamViewType_OLD_IMAGE:
            viewType = "OLD_IMAGE"
        case awsdynamodbv1.StreamViewType_STREAM_KEYS_ONLY:
            viewType = "KEYS_ONLY"
        case awsdynamodbv1.StreamViewType_NEW_AND_OLD_IMAGES:
            viewType = "NEW_AND_OLD_IMAGES"
        }
        tArgs.StreamViewType = pulumi.StringPtr(viewType)
    }

    // TTL.
    if ttl := spec.GetTtlSpecification(); ttl != nil {
        tArgs.Ttl = &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(ttl.GetTtlEnabled()),
            AttributeName: pulumi.String(ttl.GetAttributeName()),
        }
    }

    // SSE.
    if sse := spec.GetSseSpecification(); sse != nil && sse.GetEnabled() {
        sseArgs := &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if sse.GetSseType() == awsdynamodbv1.SSEType_KMS {
            sseArgs.KmsKeyArn = pulumi.StringPtr(sse.GetKmsMasterKeyId())
            // Capture for later export.
            locals.KmsKeyArn = sse.GetKmsMasterKeyId()
        }
        tArgs.ServerSideEncryption = sseArgs
    }

    // Tagging – merge user-defined tags with the platform labels.
    for k, v := range spec.GetTags() {
        tArgs.Tags[k] = pulumi.String(v)
    }

    // ---------------------------
    // Table creation
    // ---------------------------
    table, err := dynamodb.NewTable(ctx, fmt.Sprintf("%s-table", spec.GetTableName()), tArgs, pulumi.Provider(provider))
    if err != nil {
        return nil, errors.Wrap(err, "creating aws.dynamodb.Table")
    }

    // Record index names for outputs.
    for _, g := range spec.GetGlobalSecondaryIndexes() {
        locals.GsiNames = append(locals.GsiNames, g.GetIndexName())
    }
    for _, l := range spec.GetLocalSecondaryIndexes() {
        locals.LsiNames = append(locals.LsiNames, l.GetIndexName())
    }

    return table, nil
}
