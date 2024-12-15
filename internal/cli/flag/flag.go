package flag

import (
	log "github.com/sirupsen/logrus"
)

type Flag string

const (
	AutoApprove            Flag = "auto-approve"
	AwsCredential          Flag = "aws-credential"
	AzureCredential        Flag = "azure-credential"
	BackendConfig          Flag = "backend-config"
	BackendType            Flag = "backend-type"
	ConfluentCredential    Flag = "confluent-credential"
	DockerCredential       Flag = "docker-credential"
	GcpCredential          Flag = "gcp-credential"
	InputDir               Flag = "input-dir"
	KubernetesCluster      Flag = "kubernetes-cluster"
	Manifest               Flag = "manifest"
	ModuleDir              Flag = "module-dir"
	MongodbAtlasCredential Flag = "mongodb-atlas-credential"
	Set                    Flag = "set"
	SnowflakeCredential    Flag = "snowflake-credential"
	Stack                  Flag = "stack"
	Yes                    Flag = "yes"
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
