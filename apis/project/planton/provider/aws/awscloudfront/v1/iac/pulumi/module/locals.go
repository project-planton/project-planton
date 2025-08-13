package module

import (
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals groups frequently used references and derived values.
type Locals struct {
	Ctx           *pulumi.Context
	Input         *awscloudfrontv1.AwsCloudFrontStackInput
	AwsCloudFront *awscloudfrontv1.AwsCloudFront
	Spec          *awscloudfrontv1.AwsCloudFrontSpec
}

func initializeLocals(ctx *pulumi.Context, in *awscloudfrontv1.AwsCloudFrontStackInput) *Locals {
	return &Locals{
		Ctx:           ctx,
		Input:         in,
		AwsCloudFront: in.Target,
		Spec:          in.Target.Spec,
	}
}
