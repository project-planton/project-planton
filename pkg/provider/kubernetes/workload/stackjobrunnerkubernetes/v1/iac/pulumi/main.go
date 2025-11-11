package main

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/stackinput"
	stackjobrunnerkubernetesv1 "github.com/project-planton/project-planton/pkg/provider/kubernetes/workload/stackjobrunnerkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/provider/kubernetes/workload/stackjobrunnerkubernetes/v1/iac/pulumi/module"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &stackjobrunnerkubernetesv1.StackJobRunnerKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return module.Resources(ctx, stackInput)
	})
}
