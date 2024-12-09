package module

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

func enhancedMonitoring(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*iam.Role, error) {

	// Define the IAM policy document for enhanced monitoring
	enhancedMonitoringPolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Actions: []string{
					"sts:AssumeRole",
				},
				Effect: to.StringPtr("Allow"),
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type:        "Service",
						Identifiers: []string{"monitoring.rds.amazonaws.com"},
					},
				},
			},
		},
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Errorf("failed to get iam policy document")
	}

	regexReplaceChars := "[^a-zA-Z0-9-]"
	enhancedMonitoringRoleName := fmt.Sprintf("%s-emr", locals.AwsRdsCluster.Metadata.Id)
	if len(locals.AwsRdsCluster.Spec.EnhancedMonitoringAttributes) > 0 {
		normalizedAttributes := normalizeAttributes(locals.AwsRdsCluster.Spec.EnhancedMonitoringAttributes, regexReplaceChars)
		enhancedMonitoringRoleNameFull := strings.Join(normalizedAttributes, "_")
		enhancedMonitoringRoleName = truncateID(enhancedMonitoringRoleNameFull, 64)
	}

	// Create IAM Role for Enhanced Monitoring
	enhancedMonitoringRole, err := iam.NewRole(ctx, enhancedMonitoringRoleName, &iam.RoleArgs{
		Name:             pulumi.String(enhancedMonitoringRoleName),
		AssumeRolePolicy: pulumi.String(enhancedMonitoringPolicy.Json),
		Tags:             pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Errorf("failed to create enhanced monitoring role")
	}

	// Attach Amazon's managed policy for RDS enhanced monitoring
	_, err = iam.NewRolePolicyAttachment(ctx, "enhanced-monitoring-policy-attachment", &iam.RolePolicyAttachmentArgs{
		Role:      enhancedMonitoringRole.Name,
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole"),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Errorf("failed to create enhanced monitoring policy attachment")
	}

	return enhancedMonitoringRole, nil
}

// normalizeAttributes applies normalization to each attribute
func normalizeAttributes(attributes []string, regexReplaceChars string) []string {
	var normalized []string
	for _, attr := range attributes {
		normalized = append(normalized, strings.ToLower(strings.ReplaceAll(attr, regexReplaceChars, "")))
	}
	return normalized
}

// truncateID truncates the ID and appends a hash
func truncateID(id string, lengthLimit int) string {
	if lengthLimit >= len(id) {
		return id
	}
	hash := md5.Sum([]byte(id))
	return id[:lengthLimit-5] + hex.EncodeToString(hash[:])[:5]
}
