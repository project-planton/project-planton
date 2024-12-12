package flag

import (
	log "github.com/sirupsen/logrus"
)

type Flag string

const (
	AwsCredential          Flag = "aws-credential"
	AzureCredential        Flag = "azure-credential"
	ConfluentCredential    Flag = "confluent-credential"
	DockerCredential       Flag = "docker-credential"
	GcpCredential          Flag = "gcp-credential"
	KubernetesCluster      Flag = "kubernetes-cluster"
	ModuleDir              Flag = "module-dir"
	InputDir               Flag = "input-dir"
	VarFile                Flag = "var-file"
	MongodbAtlasCredential Flag = "mongodb-atlas-credential"
	SnowflakeCredential    Flag = "snowflake-credential"
	Stack                  Flag = "stack"
	Set                    Flag = "set"
	Manifest               Flag = "manifest"
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
