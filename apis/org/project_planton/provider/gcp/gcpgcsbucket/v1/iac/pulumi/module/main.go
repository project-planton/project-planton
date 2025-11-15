package module

import (
	"fmt"

	"github.com/pkg/errors"
	gcpgcsbucketv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgcsbucket/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpgcsbucketv1.GcpGcsBucketStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	// Build bucket arguments
	bucketArgs := &storage.BucketArgs{
		ForceDestroy: pulumi.Bool(true),
		Labels:       pulumi.ToStringMap(locals.GcpLabels),
		Location:     pulumi.String(locals.GcpGcsBucket.Spec.Location),
		Name:         pulumi.String(locals.GcpGcsBucket.Metadata.Name),
		Project:      pulumi.String(locals.GcpGcsBucket.Spec.GcpProjectId),
		UniformBucketLevelAccess: &storage.BucketUniformBucketLevelAccessArgs{
			Enabled: pulumi.Bool(locals.GcpGcsBucket.Spec.UniformBucketLevelAccessEnabled),
		},
	}

	// Set storage class if specified
	if locals.GcpGcsBucket.Spec.StorageClass != gcpgcsbucketv1.GcpGcsStorageClass_GCP_GCS_STORAGE_CLASS_UNSPECIFIED {
		bucketArgs.StorageClass = pulumi.String(storageClassToString(locals.GcpGcsBucket.Spec.StorageClass))
	}

	// Configure versioning if enabled
	if locals.GcpGcsBucket.Spec.VersioningEnabled {
		bucketArgs.Versioning = &storage.BucketVersioningArgs{
			Enabled: pulumi.Bool(true),
		}
	}

	// Configure lifecycle rules
	if len(locals.GcpGcsBucket.Spec.LifecycleRules) > 0 {
		var lifecycleRules storage.BucketLifecycleRuleArray
		for _, rule := range locals.GcpGcsBucket.Spec.LifecycleRules {
			lifecycleRules = append(lifecycleRules, buildLifecycleRule(rule))
		}
		bucketArgs.LifecycleRules = lifecycleRules
	}

	// Configure encryption if specified
	if locals.GcpGcsBucket.Spec.Encryption != nil && locals.GcpGcsBucket.Spec.Encryption.KmsKeyName != "" {
		bucketArgs.Encryption = &storage.BucketEncryptionArgs{
			DefaultKmsKeyName: pulumi.String(locals.GcpGcsBucket.Spec.Encryption.KmsKeyName),
		}
	}

	// Configure CORS if specified
	if len(locals.GcpGcsBucket.Spec.CorsRules) > 0 {
		var corsRules storage.BucketCorsArray
		for _, corsRule := range locals.GcpGcsBucket.Spec.CorsRules {
			corsRules = append(corsRules, buildCorsRule(corsRule))
		}
		bucketArgs.Cors = corsRules
	}

	// Configure website if specified
	if locals.GcpGcsBucket.Spec.Website != nil {
		bucketArgs.Website = &storage.BucketWebsiteArgs{
			MainPageSuffix: pulumi.String(locals.GcpGcsBucket.Spec.Website.MainPageSuffix),
			NotFoundPage:   pulumi.String(locals.GcpGcsBucket.Spec.Website.NotFoundPage),
		}
	}

	// Configure retention policy if specified
	if locals.GcpGcsBucket.Spec.RetentionPolicy != nil {
		bucketArgs.RetentionPolicy = &storage.BucketRetentionPolicyArgs{
			RetentionPeriod: pulumi.Int(int(locals.GcpGcsBucket.Spec.RetentionPolicy.RetentionPeriodSeconds)),
			IsLocked:        pulumi.Bool(locals.GcpGcsBucket.Spec.RetentionPolicy.IsLocked),
		}
	}

	// Configure requester pays if enabled
	if locals.GcpGcsBucket.Spec.RequesterPays {
		bucketArgs.RequesterPays = pulumi.Bool(true)
	}

	// Configure logging if specified
	if locals.GcpGcsBucket.Spec.Logging != nil {
		bucketArgs.Logging = &storage.BucketLoggingArgs{
			LogBucket:       pulumi.String(locals.GcpGcsBucket.Spec.Logging.LogBucket),
			LogObjectPrefix: pulumi.String(locals.GcpGcsBucket.Spec.Logging.LogObjectPrefix),
		}
	}

	// Configure public access prevention if specified
	if locals.GcpGcsBucket.Spec.PublicAccessPrevention != "" {
		bucketArgs.PublicAccessPrevention = pulumi.String(locals.GcpGcsBucket.Spec.PublicAccessPrevention)
	}

	// Create the bucket
	createdBucket, err := storage.NewBucket(ctx,
		locals.GcpGcsBucket.Metadata.Name,
		bucketArgs,
		pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create bucket resource")
	}

	ctx.Export(OpBucketId, createdBucket.ID())

	// Configure IAM bindings if specified
	if len(locals.GcpGcsBucket.Spec.IamBindings) > 0 {
		for i, binding := range locals.GcpGcsBucket.Spec.IamBindings {
			iamBindingArgs := &storage.BucketIAMBindingArgs{
				Bucket:  createdBucket.Name,
				Role:    pulumi.String(binding.Role),
				Members: pulumi.ToStringArray(binding.Members),
			}

			// Add condition if specified
			if binding.Condition != "" {
				iamBindingArgs.Condition = &storage.BucketIAMBindingConditionArgs{
					Expression: pulumi.String(binding.Condition),
					Title:      pulumi.String(fmt.Sprintf("condition-%d", i)),
				}
			}

			_, err = storage.NewBucketIAMBinding(ctx,
				fmt.Sprintf("%s-iam-%d", locals.GcpGcsBucket.Metadata.Name, i),
				iamBindingArgs,
				pulumi.Parent(createdBucket),
				pulumi.Provider(gcpProvider))
			if err != nil {
				return errors.Wrapf(err, "failed to create IAM binding %d", i)
			}
		}
	}

	return nil
}

// storageClassToString converts the proto enum to GCS storage class string
func storageClassToString(class gcpgcsbucketv1.GcpGcsStorageClass) string {
	switch class {
	case gcpgcsbucketv1.GcpGcsStorageClass_STANDARD:
		return "STANDARD"
	case gcpgcsbucketv1.GcpGcsStorageClass_NEARLINE:
		return "NEARLINE"
	case gcpgcsbucketv1.GcpGcsStorageClass_COLDLINE:
		return "COLDLINE"
	case gcpgcsbucketv1.GcpGcsStorageClass_ARCHIVE:
		return "ARCHIVE"
	default:
		return "STANDARD"
	}
}

// buildLifecycleRule converts proto lifecycle rule to Pulumi lifecycle rule
func buildLifecycleRule(rule *gcpgcsbucketv1.GcpGcsLifecycleRule) storage.BucketLifecycleRuleArgs {
	lifecycleRule := storage.BucketLifecycleRuleArgs{
		Action: &storage.BucketLifecycleRuleActionArgs{
			Type: pulumi.String(rule.Action.Type),
		},
		Condition: &storage.BucketLifecycleRuleConditionArgs{},
	}

	// Set storage class for SetStorageClass action
	if rule.Action.StorageClass != gcpgcsbucketv1.GcpGcsStorageClass_GCP_GCS_STORAGE_CLASS_UNSPECIFIED {
		lifecycleRule.Action.StorageClass = pulumi.String(storageClassToString(rule.Action.StorageClass))
	}

	// Set condition fields
	if rule.Condition.AgeDays > 0 {
		lifecycleRule.Condition.Age = pulumi.Int(int(rule.Condition.AgeDays))
	}
	if rule.Condition.CreatedBefore != "" {
		lifecycleRule.Condition.CreatedBefore = pulumi.String(rule.Condition.CreatedBefore)
	}
	if rule.Condition.NumNewerVersions > 0 {
		lifecycleRule.Condition.NumNewerVersions = pulumi.Int(int(rule.Condition.NumNewerVersions))
	}
	if len(rule.Condition.MatchesStorageClass) > 0 {
		var storageClasses pulumi.StringArray
		for _, class := range rule.Condition.MatchesStorageClass {
			storageClasses = append(storageClasses, pulumi.String(storageClassToString(class)))
		}
		lifecycleRule.Condition.MatchesStorageClasses = storageClasses
	}

	return lifecycleRule
}

// buildCorsRule converts proto CORS rule to Pulumi CORS rule
func buildCorsRule(corsRule *gcpgcsbucketv1.GcpGcsCorsRule) storage.BucketCorsArgs {
	return storage.BucketCorsArgs{
		Methods:         pulumi.ToStringArray(corsRule.Methods),
		Origins:         pulumi.ToStringArray(corsRule.Origins),
		ResponseHeaders: pulumi.ToStringArray(corsRule.ResponseHeaders),
		MaxAgeSeconds:   pulumi.Int(int(corsRule.MaxAgeSeconds)),
	}
}
