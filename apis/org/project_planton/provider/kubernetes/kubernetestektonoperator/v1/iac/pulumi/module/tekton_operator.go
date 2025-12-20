package module

import (
	"github.com/pkg/errors"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// tektonOperator installs the Tekton Operator using release manifests.
func tektonOperator(ctx *pulumi.Context, locals *Locals,
	k8sProvider *pulumikubernetes.Provider) error {

	// --------------------------------------------------------------------
	// 1. Install Tekton Operator from release manifests
	// --------------------------------------------------------------------
	operatorManifests, err := yaml.NewConfigFile(ctx, "tekton-operator", &yaml.ConfigFileArgs{
		File: locals.OperatorReleaseURL,
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return errors.Wrap(err, "install tekton operator manifests")
	}

	// --------------------------------------------------------------------
	// 2. Create TektonConfig to configure components
	// --------------------------------------------------------------------
	// Determine the profile based on enabled components
	profile := "lite" // minimal profile
	if locals.EnablePipelines && locals.EnableTriggers && locals.EnableDashboard {
		profile = "all"
	} else if locals.EnablePipelines && locals.EnableTriggers {
		profile = "basic"
	}

	// Build TektonConfig YAML dynamically
	tektonConfigYAML := buildTektonConfigYAML(locals, profile)

	_, err = yaml.NewConfigGroup(ctx, "tekton-config", &yaml.ConfigGroupArgs{
		YAML: []string{tektonConfigYAML},
	}, pulumi.Provider(k8sProvider), pulumi.DependsOn([]pulumi.Resource{operatorManifests}))
	if err != nil {
		return errors.Wrap(err, "create tekton config")
	}

	return nil
}

// buildTektonConfigYAML constructs the TektonConfig YAML based on component settings.
func buildTektonConfigYAML(locals *Locals, profile string) string {
	// TektonConfig CRD that tells the operator which components to install
	// Note: Do not set fields that the operator manages automatically (e.g., pipeline.enable-api-fields)
	// to avoid Server-Side Apply field conflicts
	return `apiVersion: operator.tekton.dev/v1alpha1
kind: TektonConfig
metadata:
  name: ` + locals.TektonConfigName + `
spec:
  profile: ` + profile + `
  targetNamespace: ` + locals.ComponentsNamespace + `
`
}
