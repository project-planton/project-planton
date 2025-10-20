package flag

import (
	log "github.com/sirupsen/logrus"
)

type Flag string

const (
	AutoApprove              Flag = "auto-approve"
	AwsProviderConfig        Flag = "aws-provider-config"
	AzureProviderConfig      Flag = "azure-provider-config"
	BackendConfig            Flag = "backend-config"
	BackendType              Flag = "backend-type"
	ConfluentProviderConfig  Flag = "confluent-provider-config"
	Destroy                  Flag = "destroy"
	Diff                     Flag = "diff"
	Force                    Flag = "force"
	GcpProviderConfig        Flag = "gcp-provider-config"
	InputDir                 Flag = "input-dir"
	KubernetesProviderConfig Flag = "kubernetes-provider-config"
	KustomizeDir             Flag = "kustomize-dir"
	Manifest                 Flag = "manifest"
	ModuleDir                Flag = "module-dir"
	AtlasProviderConfig      Flag = "atlas-provider-config"
	OutputFile               Flag = "output-file"
	Overlay                  Flag = "overlay"
	Set                      Flag = "set"
	SnowflakeProviderConfig  Flag = "snowflake-provider-config"
	Stack                    Flag = "stack"
	Yes                      Flag = "yes"
)

func HandleFlagErrAndValue(err error, flag Flag, flagVal string) {
	if err != nil {
		log.Fatalf("error parsing %s flag. err %v", flag, err)
	}
	if flagVal == "" {
		log.Fatalf("please provide %s", flag)
	}
}

func HandleFlagErr(err error, flag Flag) {
	if err != nil {
		log.Fatalf("error parsing %s flag. err %v", flag, err)
	}
}
