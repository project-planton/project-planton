package module

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func iamRole(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	roleName := locals.AwsIamRole.Metadata.Name
	spec := locals.AwsIamRole.Spec

	// Create the core IAM role
	iamRole, err := iam.NewRole(ctx, roleName, &iam.RoleArgs{
		Name:             pulumi.String(roleName),
		AssumeRolePolicy: pulumi.String(spec.TrustPolicyJson),
		Description:      pulumi.String(spec.Description),
		Path:             pulumi.String(spec.Path),
		Tags:             pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create IAM role")
	}

	// Attach any AWS-managed or customer-managed IAM policies by ARN
	for idx, policyArn := range spec.ManagedPolicyArns {
		attachName := fmt.Sprintf("%s-attach-%d", roleName, idx)
		_, err := iam.NewRolePolicyAttachment(ctx, attachName, &iam.RolePolicyAttachmentArgs{
			Role:      iamRole.Name,
			PolicyArn: pulumi.String(policyArn),
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrapf(err, "failed to attach policy ARN %s", policyArn)
		}
	}

	// Create inline policies if specified
	for policyName, policyJson := range spec.InlinePolicyJsons {
		inlineName := fmt.Sprintf("%s-inline-%s", roleName, policyName)
		_, err := iam.NewRolePolicy(ctx, inlineName, &iam.RolePolicyArgs{
			Role:   iamRole.Name,
			Policy: pulumi.String(policyJson),
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrapf(err, "failed to create inline policy %s", policyName)
		}
	}

	// Export final outputs
	ctx.Export(OpRoleArn, iamRole.Arn)
	ctx.Export(OpRoleName, iamRole.Name)

	return nil
}

// parseJson is an optional helper if you want to do JSON validation before apply
func parseJson(raw string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
