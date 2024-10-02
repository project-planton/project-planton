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
	GitCredential          Flag = "git-credential"
	KubernetesCluster      Flag = "kubernetes-cluster"
	MongodbAtlasCredential Flag = "mongodb-atlas-credential"
	SnowflakeCredential    Flag = "snowflake-credential"
	Stack                  Flag = "stack"
	Target                 Flag = "target"
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
