package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

func optionGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*rds.OptionGroup, error) {
	majorEngineVersion := locals.AwsRdsInstance.Spec.MajorEngineVersion
	if majorEngineVersion == "" && locals.AwsRdsInstance.Spec.EngineVersion != "" {
		versionParts := strings.Split(locals.AwsRdsInstance.Spec.EngineVersion, ".")
		// If the engine is "postgres", take the first part, otherwise take the first two parts
		if locals.AwsRdsInstance.Spec.Engine == "postgres" {
			if len(versionParts) >= 1 {
				majorEngineVersion = versionParts[0]
			}
		} else {
			if len(versionParts) >= 2 {
				majorEngineVersion = strings.Join(versionParts[0:2], ".")
			}
		}
	}

	optionArray := rds.OptionGroupOptionArray{}
	for _, option := range locals.AwsRdsInstance.Spec.Options {
		optionSettingsArray := rds.OptionGroupOptionOptionSettingArray{}
		for _, optionSetting := range option.OptionSettings {
			optionSettingsArray = append(optionSettingsArray, &rds.OptionGroupOptionOptionSettingArgs{
				Name:  pulumi.String(optionSetting.Name),
				Value: pulumi.String(optionSetting.Value),
			})
		}
		optionArray = append(optionArray, &rds.OptionGroupOptionArgs{
			DbSecurityGroupMemberships:  pulumi.ToStringArray(option.DbSecurityGroupMemberships),
			OptionName:                  pulumi.String(option.OptionName),
			OptionSettings:              optionSettingsArray,
			Port:                        pulumi.Int(option.Port),
			Version:                     pulumi.String(option.Version),
			VpcSecurityGroupMemberships: pulumi.ToStringArray(option.VpcSecurityGroupMemberships),
		})
	}

	// Create RDS Option Group (optional based on the engine type)
	rdsOptionGroup, err := rds.NewOptionGroup(ctx, "rds-options-group", &rds.OptionGroupArgs{
		NamePrefix:         pulumi.Sprintf("%s-", locals.AwsRdsInstance.Metadata.Id),
		EngineName:         pulumi.String(locals.AwsRdsInstance.Spec.Engine),
		MajorEngineVersion: pulumi.String(majorEngineVersion),
		Tags:               pulumi.ToStringMap(locals.Labels),
		Options:            optionArray,
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rds option group")
	}

	return rdsOptionGroup, nil
}
