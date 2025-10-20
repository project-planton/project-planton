package stackinputproviderconfig

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/spf13/pflag"
)

type StackInputProviderConfigOptions struct {
	AwsProviderConfig        string
	AzureProviderConfig      string
	ConfluentProviderConfig  string
	GcpProviderConfig        string
	KubernetesProviderConfig string
	AtlasProviderConfig      string
	SnowflakeProviderConfig  string
}

type StackInputProviderConfigOption func(*StackInputProviderConfigOptions)

func WithAwsProviderConfig(awsProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.AwsProviderConfig = awsProviderConfig
	}
}

func WithAzureProviderConfig(azureProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.AzureProviderConfig = azureProviderConfig
	}
}

func WithConfluentProviderConfig(confluentProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.ConfluentProviderConfig = confluentProviderConfig
	}
}

func WithGcpProviderConfig(gcpProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.GcpProviderConfig = gcpProviderConfig
	}
}

func WithKubernetesProviderConfig(kubernetesProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.KubernetesProviderConfig = kubernetesProviderConfig
	}
}

func WithAtlasProviderConfig(atlasProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.AtlasProviderConfig = atlasProviderConfig
	}
}

func WithSnowflakeProviderConfig(snowflakeProviderConfig string) StackInputProviderConfigOption {
	return func(opts *StackInputProviderConfigOptions) {
		opts.SnowflakeProviderConfig = snowflakeProviderConfig
	}
}

func BuildWithFlags(commandFlagSet *pflag.FlagSet) ([]StackInputProviderConfigOption, error) {
	resp := make([]StackInputProviderConfigOption, 0)

	awsProviderConfig, err := commandFlagSet.GetString(string(flag.AwsProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.AwsProviderConfig)
	}
	if awsProviderConfig != "" {
		resp = append(resp, WithAwsProviderConfig(awsProviderConfig))
	}

	azureProviderConfig, err := commandFlagSet.GetString(string(flag.AzureProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.AzureProviderConfig)
	}
	if azureProviderConfig != "" {
		resp = append(resp, WithAzureProviderConfig(azureProviderConfig))
	}

	confluentProviderConfig, err := commandFlagSet.GetString(string(flag.ConfluentProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.ConfluentProviderConfig)
	}
	if confluentProviderConfig != "" {
		resp = append(resp, WithConfluentProviderConfig(confluentProviderConfig))
	}

	gcpProviderConfig, err := commandFlagSet.GetString(string(flag.GcpProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.GcpProviderConfig)
	}
	if gcpProviderConfig != "" {
		resp = append(resp, WithGcpProviderConfig(gcpProviderConfig))
	}

	kubernetesProviderConfig, err := commandFlagSet.GetString(string(flag.KubernetesProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.KubernetesProviderConfig)
	}
	if kubernetesProviderConfig != "" {
		resp = append(resp, WithKubernetesProviderConfig(kubernetesProviderConfig))
	}

	atlasProviderConfig, err := commandFlagSet.GetString(string(flag.AtlasProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.AtlasProviderConfig)
	}
	if atlasProviderConfig != "" {
		resp = append(resp, WithAtlasProviderConfig(atlasProviderConfig))
	}

	snowflakeProviderConfig, err := commandFlagSet.GetString(string(flag.SnowflakeProviderConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.SnowflakeProviderConfig)
	}
	if snowflakeProviderConfig != "" {
		resp = append(resp, WithSnowflakeProviderConfig(snowflakeProviderConfig))
	}
	return resp, nil
}

func BuildWithInputDir(inputDir string) ([]StackInputProviderConfigOption, error) {
	resp := make([]StackInputProviderConfigOption, 0)

	awsProviderConfig, err := LoadAwsProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up aws-provider-config in %s", inputDir)
	}
	if awsProviderConfig != "" {
		resp = append(resp, WithAwsProviderConfig(awsProviderConfig))
	}

	azureProviderConfig, err := LoadAzureProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up azure-provider-config in %s", inputDir)
	}
	if azureProviderConfig != "" {
		resp = append(resp, WithAzureProviderConfig(azureProviderConfig))
	}

	confluentProviderConfig, err := LoadConfluentProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up confluent-provider-config in %s", inputDir)
	}
	if confluentProviderConfig != "" {
		resp = append(resp, WithConfluentProviderConfig(confluentProviderConfig))
	}

	gcpProviderConfig, err := LoadGcpProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up gcp-provider-config in %s", inputDir)
	}
	if gcpProviderConfig != "" {
		resp = append(resp, WithGcpProviderConfig(gcpProviderConfig))
	}

	kubernetesProviderConfig, err := LoadKubernetesProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up kubernetes-provider-config in %s", inputDir)
	}
	if kubernetesProviderConfig != "" {
		resp = append(resp, WithKubernetesProviderConfig(kubernetesProviderConfig))
	}

	atlasProviderConfig, err := LoadAtlasProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up atlas-provider-config in %s", inputDir)
	}
	if atlasProviderConfig != "" {
		resp = append(resp, WithAtlasProviderConfig(atlasProviderConfig))
	}

	snowflakeProviderConfig, err := LoadSnowflakeProviderConfig(inputDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to look up snowflake-provider-config in %s", inputDir)
	}
	if snowflakeProviderConfig != "" {
		resp = append(resp, WithSnowflakeProviderConfig(snowflakeProviderConfig))
	}
	return resp, nil
}
