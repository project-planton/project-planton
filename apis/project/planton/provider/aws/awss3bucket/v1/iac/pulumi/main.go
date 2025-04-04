package main

import (
	"github.com/pkg/errors"
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1/iac/pulumi/module"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &awss3bucketv1.AwsS3BucketStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return module.Resources(ctx, stackInput)
	})
}
