package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func iamRole(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*iam.Role, error) {

	// Fetch AWS partition, region, and account information
	partition, err := aws.GetPartition(ctx, nil, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get partition")
	}
	region, err := aws.GetRegion(ctx, nil, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get region")
	}
	callerIdentity, err := aws.GetCallerIdentity(ctx, nil, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account id")
	}

	assumeRolePolicy := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
		Statements: iam.GetPolicyDocumentStatementArray{
			iam.GetPolicyDocumentStatementArgs{
				Actions: pulumi.StringArray{pulumi.String("sts:AssumeRole")},
				Principals: iam.GetPolicyDocumentStatementPrincipalArray{
					iam.GetPolicyDocumentStatementPrincipalArgs{
						Type:        pulumi.String("Service"),
						Identifiers: pulumi.StringArray{pulumi.String("lambda.amazonaws.com")},
					},
				},
			},
		},
	}, pulumi.Provider(awsProvider))

	createdIamRole, err := iam.NewRole(ctx,
		"iam-role",
		&iam.RoleArgs{
			Name:             pulumi.Sprintf("%s-lambda-iam-role", locals.AwsLambda.Metadata.Id),
			AssumeRolePolicy: assumeRolePolicy.Json(),
		}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create iam role")
	}

	// Attach policies to the role
	_, err = iam.NewRolePolicyAttachment(ctx,
		"cloudwatch-logs",
		&iam.RolePolicyAttachmentArgs{
			Role:      createdIamRole.Name,
			PolicyArn: pulumi.Sprintf("arn:%s:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole", partition.Partition),
		}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
	if err != nil {
		return nil, errors.Wrap(err, "failed to attach cloud watch logs role policy")
	}

	if locals.AwsLambda.Spec.Function.VpcConfig != nil {
		_, err = iam.NewRolePolicyAttachment(ctx,
			"vpc-access",
			&iam.RolePolicyAttachmentArgs{
				Role:      createdIamRole.Name,
				PolicyArn: pulumi.Sprintf("arn:%s:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole", partition.Partition),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to attach vpc access role policy")
		}
	}

	if locals.AwsLambda.Spec.Function.TracingConfigMode != "" {
		_, err = iam.NewRolePolicyAttachment(ctx,
			"xray",
			&iam.RolePolicyAttachmentArgs{
				Role:      createdIamRole.Name,
				PolicyArn: pulumi.Sprintf("arn:%s:iam::aws:policy/AWSXRayDaemonWriteAccess", partition.Partition),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to attach xray daemon role policy")
		}
	}

	if locals.AwsLambda.Spec.IamRole == nil {
		return createdIamRole, nil
	}

	if locals.AwsLambda.Spec.IamRole.CloudwatchLambdaInsightsEnabled {
		_, err = iam.NewRolePolicyAttachment(ctx,
			"cloudwatch-insights",
			&iam.RolePolicyAttachmentArgs{
				Role:      createdIamRole.Name,
				PolicyArn: pulumi.Sprintf("arn:%s:iam::aws:policy/CloudWatchLambdaInsightsExecutionRolePolicy", partition.Partition),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to attach cloud watch insights role policy")
		}
	}

	if locals.AwsLambda.Spec.IamRole.SsmParameterNames != nil &&
		len(locals.AwsLambda.Spec.IamRole.SsmParameterNames) > 0 {
		var ssmParameterArns []string
		for _, ssmParameterName := range locals.AwsLambda.Spec.IamRole.SsmParameterNames {
			ssmParameterArns = append(ssmParameterArns, fmt.Sprintf("arn:%s:ssm:%s:%s:parameter%s", partition.Partition, region.Name, callerIdentity.AccountId, ssmParameterName))
		}

		// Create the IAM Policy Document for accessing SSM parameters
		ssmPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("ssm:GetParameter"),
						pulumi.String("ssm:GetParameters"),
						pulumi.String("ssm:GetParametersByPath"),
					},
					Resources: pulumi.ToStringArray(ssmParameterArns),
				},
			},
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get ssm policy document")
		}

		// Create the IAM Policy for the Lambda Function
		ssmPolicy, err := iam.NewPolicy(ctx,
			"ssm",
			&iam.PolicyArgs{
				Name:        pulumi.Sprintf("%s-ssm-policy-%s", locals.AwsLambda.Metadata.Id, region.Name),
				Description: pulumi.String("Provides read access to specific SSM parameters for Lambda."),
				Policy:      ssmPolicyDocument.Json(),
				Tags:        pulumi.ToStringMap(locals.AwsLambda.Metadata.Labels),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new ssm policy")
		}

		// Attach the SSM policy to the Lambda IAM role
		_, err = iam.NewRolePolicyAttachment(ctx,
			"ssm",
			&iam.RolePolicyAttachmentArgs{
				Role:      createdIamRole.Name,
				PolicyArn: ssmPolicy.Arn,
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to attach ssm policy to role")
		}
	}

	if locals.AwsLambda.Spec.IamRole.CustomIamPolicyArns != nil {
		for i, CustomIamPolicyArn := range locals.AwsLambda.Spec.IamRole.CustomIamPolicyArns {
			_, err = iam.NewRolePolicyAttachment(ctx,
				fmt.Sprintf("custom-%d", i),
				&iam.RolePolicyAttachmentArgs{
					Role:      createdIamRole.Name,
					PolicyArn: pulumi.String(CustomIamPolicyArn),
				}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
			if err != nil {
				return nil, errors.Wrap(err, "failed to attach custom policy to role")
			}
		}
	}

	if locals.AwsLambda.Spec.IamRole.InlineIamPolicy != "" {
		_, err = iam.NewRolePolicy(ctx,
			"inline",
			&iam.RolePolicyArgs{
				Role:   createdIamRole.Name,
				Policy: pulumi.String(locals.AwsLambda.Spec.IamRole.InlineIamPolicy),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdIamRole))
		if err != nil {
			return nil, errors.Wrap(err, "failed to create inline policy to role")
		}
	}
	return createdIamRole, nil
}
