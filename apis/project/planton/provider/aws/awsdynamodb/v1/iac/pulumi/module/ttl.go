package awsdynamodb

import (
    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// buildTtl converts the protobuf-based Time-to-Live (TTL) specification into
// the shape expected by the Pulumi AWS provider. The helper intentionally
// returns nil when the caller did not request a TTL configuration so the
// resulting TableArgs omit the ttl block entirely – this prevents needless
// provider-level diffs. When the message is present, the function faithfully
// reflects the user intent:
//   • enabled  – attribute name must be forwarded together with Enabled=true
//   • disabled – an explicit Enabled=false is sent to ensure a previously
//                configured TTL is removed.
//
// Example usage inside a higher-level builder:
//
//   tableArgs := &dynamodb.TableArgs{
//       // … other arguments …
//       Ttl: buildTtl(spec.TtlSpecification),
//   }
func buildTtl(ttlSpec *awsdynamodbpb.TimeToLiveSpecification) *dynamodb.TableTtlArgs {
    if ttlSpec == nil {
        // Caller did not specify TTL settings – omit the block.
        return nil
    }

    // When TTL is enabled we must provide both the flag and the attribute name.
    if ttlSpec.TtlEnabled {
        return &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(true),
            AttributeName: pulumi.StringPtr(ttlSpec.AttributeName),
        }
    }

    // TTL specification is present but disabled. Returning an explicit block
    // with Enabled=false instructs the provider to disable TTL if it was
    // previously enabled.
    return &dynamodb.TableTtlArgs{
        Enabled: pulumi.Bool(false),
    }
}
