package main

import (
	"github.com/pkg/errors"
	cronjobkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/cronjobkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/cronjobkubernetes/v1/iac/pulumi/module"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &cronjobkubernetesv1.CronJobKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return module.Resources(ctx, stackInput)
	})
}
