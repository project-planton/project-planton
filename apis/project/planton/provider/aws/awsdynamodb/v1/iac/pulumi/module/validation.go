package module

import (
    "fmt"

    "github.com/pkg/errors"
    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// ValidateStackInput executes additional, cross-field runtime checks that are
// inconvenient (or impossible) to describe in protobuf/CEL rules alone. It
// should be invoked by main.go before any provider resources are created so we
// can fail fast and return a descriptive error back to the caller.
func ValidateStackInput(in *awsdynamodbpb.AwsDynamodbStackInput) error {
    if in == nil {
        return errors.New("stack input is nil")
    }

    if in.Target == nil {
        return errors.New("target AwsDynamodb resource must be provided")
    }

    spec := in.Target.GetSpec()
    if spec == nil {
        return errors.New("AwsDynamodb.spec must be provided")
    }

    // Delegate to spec-specific validation logic.
    if err := validateDynamoDBSpec(spec); err != nil {
        return errors.Wrap(err, "invalid AwsDynamodb.spec")
    }

    return nil
}

// validateDynamoDBSpec enforces semantic rules that depend on multiple fields
// of the AwsDynamodbSpec message.
func validateDynamoDBSpec(spec *awsdynamodbpb.AwsDynamodbSpec) error {
    if spec == nil {
        return errors.New("spec is nil")
    }

    // ---------------------------------------------------------------------
    // Billing mode ↔ provisioned throughput consistency.
    // ---------------------------------------------------------------------
    switch spec.BillingMode {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        if spec.ProvisionedThroughput == nil {
            return errors.New("provisioned_throughput must be set when billing_mode is PROVISIONED")
        }
        for i, gsi := range spec.GlobalSecondaryIndexes {
            if gsi.ProvisionedThroughput == nil {
                return fmt.Errorf("global_secondary_indexes[%d].provisioned_throughput must be set when billing_mode is PROVISIONED", i)
            }
        }

    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        if spec.ProvisionedThroughput != nil {
            return errors.New("provisioned_throughput must be unset when billing_mode is PAY_PER_REQUEST")
        }
        for i, gsi := range spec.GlobalSecondaryIndexes {
            if gsi.ProvisionedThroughput != nil {
                return fmt.Errorf("global_secondary_indexes[%d].provisioned_throughput must be unset when billing_mode is PAY_PER_REQUEST", i)
            }
        }

    default:
        return errors.New("billing_mode must be either PROVISIONED or PAY_PER_REQUEST")
    }

    // ---------------------------------------------------------------------
    // Streams configuration checks.
    // ---------------------------------------------------------------------
    if spec.StreamSpecification != nil && spec.StreamSpecification.StreamEnabled {
        if spec.StreamSpecification.StreamViewType == awsdynamodbpb.StreamViewType_STREAM_VIEW_TYPE_UNSPECIFIED {
            return errors.New("stream_specification.stream_view_type must be specified when streams are enabled")
        }
    }

    // ---------------------------------------------------------------------
    // TTL configuration checks.
    // ---------------------------------------------------------------------
    if spec.TtlSpecification != nil && spec.TtlSpecification.TtlEnabled {
        if spec.TtlSpecification.AttributeName == "" {
            return errors.New("ttl_specification.attribute_name must be provided when TTL is enabled")
        }
    }

    // ---------------------------------------------------------------------
    // Server-side encryption (SSE) rules.
    // ---------------------------------------------------------------------
    if spec.SseSpecification != nil {
        sse := spec.SseSpecification
        if !sse.Enabled {
            // When disabled, type and CMK must be unset.
            if sse.SseType != awsdynamodbpb.SSEType_SSE_TYPE_UNSPECIFIED || sse.KmsMasterKeyId != "" {
                return errors.New("sse_type and kms_master_key_id must be unset when SSE is disabled")
            }
        } else {
            // Enabled → type must be set.
            if sse.SseType == awsdynamodbpb.SSEType_SSE_TYPE_UNSPECIFIED {
                return errors.New("sse_type must be specified when SSE is enabled")
            }
            switch sse.SseType {
            case awsdynamodbpb.SSEType_KMS:
                if sse.KmsMasterKeyId == "" {
                    return errors.New("kms_master_key_id must be provided when sse_type is KMS")
                }
            case awsdynamodbpb.SSEType_AES256:
                if sse.KmsMasterKeyId != "" {
                    return errors.New("kms_master_key_id must be empty when sse_type is AES256")
                }
            }
        }
    }

    return nil
}
