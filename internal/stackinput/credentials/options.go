package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/cli/flag"
	"github.com/spf13/pflag"
)

type StackInputCredentialOptions struct {
	AwsCredential          string
	AzureCredential        string
	ConfluentCredential    string
	DockerCredential       string
	GcpCredential          string
	KubernetesCluster      string
	MongodbAtlasCredential string
	SnowflakeCredential    string
}

type StackInputCredentialOption func(*StackInputCredentialOptions)

func WithAwsCredential(awsCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.AwsCredential = awsCredential
	}
}

func WithAzureCredential(azureCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.AzureCredential = azureCredential
	}
}

func WithConfluentCredential(confluentCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.ConfluentCredential = confluentCredential
	}
}

func WithDockerCredential(dockerCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.DockerCredential = dockerCredential
	}
}

func WithGcpCredential(gcpCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.GcpCredential = gcpCredential
	}
}

func WithKubernetesCluster(kubernetesCluster string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.KubernetesCluster = kubernetesCluster
	}
}

func WithMongodbAtlasCredential(mongodbAtlasCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.MongodbAtlasCredential = mongodbAtlasCredential
	}
}

func WithSnowflakeCredential(snowflakeCredential string) StackInputCredentialOption {
	return func(opts *StackInputCredentialOptions) {
		opts.SnowflakeCredential = snowflakeCredential
	}
}

func BuildOptions(commandFlagSet *pflag.FlagSet) ([]StackInputCredentialOption, error) {
	resp := make([]StackInputCredentialOption, 0)

	awsCredential, err := commandFlagSet.GetString(string(flag.AwsCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.AwsCredential)
	}
	if awsCredential != "" {
		resp = append(resp, WithAwsCredential(awsCredential))
	}

	azureCredential, err := commandFlagSet.GetString(string(flag.AzureCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.AzureCredential)
	}
	if azureCredential != "" {
		resp = append(resp, WithAzureCredential(azureCredential))
	}

	confluentCredential, err := commandFlagSet.GetString(string(flag.ConfluentCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.ConfluentCredential)
	}
	if confluentCredential != "" {
		resp = append(resp, WithConfluentCredential(confluentCredential))
	}

	dockerCredential, err := commandFlagSet.GetString(string(flag.DockerCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.DockerCredential)
	}
	if dockerCredential != "" {
		resp = append(resp, WithDockerCredential(dockerCredential))
	}

	gcpCredential, err := commandFlagSet.GetString(string(flag.GcpCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.GcpCredential)
	}
	if gcpCredential != "" {
		resp = append(resp, WithGcpCredential(gcpCredential))
	}

	kubernetesCluster, err := commandFlagSet.GetString(string(flag.KubernetesCluster))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.KubernetesCluster)
	}
	if kubernetesCluster != "" {
		resp = append(resp, WithKubernetesCluster(kubernetesCluster))
	}

	mongodbAtlasCredential, err := commandFlagSet.GetString(string(flag.MongodbAtlasCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.MongodbAtlasCredential)
	}
	if mongodbAtlasCredential != "" {
		resp = append(resp, WithMongodbAtlasCredential(mongodbAtlasCredential))
	}

	snowflakeCredential, err := commandFlagSet.GetString(string(flag.SnowflakeCredential))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s flag", flag.SnowflakeCredential)
	}
	if snowflakeCredential != "" {
		resp = append(resp, WithSnowflakeCredential(snowflakeCredential))
	}
	return resp, nil
}
