package main

import (
	"github.com/pkg/errors"
	awsvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1/iac/pulumi/module"
	"github.com/project-planton/project-planton/pkg/pulmod/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &awsvpcv1.AwsVpcStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return module.Resources(ctx, stackInput)
	})
}
