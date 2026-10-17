package module

import (
	"github.com/pkg/errors"
	kubernetesagentsubstratev1alpha1 "github.com/plantonhq/planton/catalog/kubernetes/kubernetesagentsubstrate/v1alpha1"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/keptcrds"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/plantonhq/planton/pkg/kubernetes/helmcrds"
	pulumiyaml "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources installs the Agent Substrate control plane from Planton's build
// of the pinned release: upstream's own manifests with every image
// digest-pinned, composed the way upstream's installer composes them, with
// the typed settings applied onto upstream's own objects (compose.go).
//
// ORDER: everything that can refuse runs before anything registers
// (preflight.go): a namespace other than upstream's, a workload or namespace
// another installation holds, and a release move upstream cannot make. Then
// the namespaces, the kept definitions (ownership-checked, never downgraded),
// the generated and derived Secrets, and the four document phases, each
// awaited before the next: prerequisites (the pod-certificate controller
// every certificate comes from), database, system, runtime.
//
// DESTROY: in reverse. The default SandboxConfig goes first, while the
// controller that clears its finalizer still runs; the namespaces last. The
// definitions are kept unless crds.keep_on_uninstall is false, so worker
// pools and their actors are not deleted with the control plane; the bundled
// database's volume is kept by its StatefulSet, and the upgrade refusal reads
// its stamp.
func Resources(ctx *pulumi.Context, iacInput *kubernetesagentsubstratev1alpha1.KubernetesAgentSubstrateIacInput) error {
	locals := initializeLocals(iacInput)
	s := resolveSettings(locals.Spec)
	s.identityLabels = locals.IdentityLabels
	if err := preflight(iacInput.ProviderConfig, s); err != nil {
		return err
	}

	files, err := fetchRelease()
	if err != nil {
		return err
	}
	composed, err := compose(s, files)
	if err != nil {
		return err
	}

	createdCrds, err := keptcrds.Apply(ctx, keptcrds.Args{
		Source: helmcrds.Source{
			BundleURL:  CRDBundleURLTemplate(),
			BundleName: vars.BundleName,
			Version:    vars.Release,
		},
		Policy:          helmcrds.Policy{ExpectCRDs: true},
		ReleaseName:     locals.ResourceName,
		Namespace:       vars.Namespace,
		Install:         s.crdsInstall,
		KeepOnUninstall: s.crdsKeep,
		ProviderConfig:  iacInput.ProviderConfig,
		ProviderName:    "kubernetes-crd-upsert",
	})
	if err != nil {
		return errors.Wrap(err, "failed to apply the Agent Substrate resource definitions")
	}

	provider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx, iacInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	dependsOn := createdCrds
	group := func(suffix string, objects []document) error {
		if len(objects) == 0 {
			return nil
		}
		applied, err := pulumiyaml.NewConfigGroup(ctx, locals.ResourceName+"-"+suffix,
			&pulumiyaml.ConfigGroupArgs{Objs: objects},
			pulumi.Provider(provider), pulumi.DependsOn(dependsOn))
		if err != nil {
			return errors.Wrapf(err, "failed to apply the Agent Substrate %s", suffix)
		}
		dependsOn = []pulumi.Resource{applied}
		return nil
	}

	if err := group("namespaces", composed.namespaces); err != nil {
		return err
	}
	secrets, err := applySecrets(ctx, s, locals.IdentityLabels, pulumi.Provider(provider), pulumi.DependsOn(dependsOn))
	if err != nil {
		return err
	}
	dependsOn = append(dependsOn, secrets...)
	for _, phase := range []struct {
		suffix  string
		objects []document
	}{
		{"prerequisites", composed.prerequisites},
		{"database", composed.database},
		{"system", composed.system},
		{"runtime", composed.runtime},
	} {
		if err := group(phase.suffix, phase.objects); err != nil {
			return err
		}
	}

	exportOutputs(ctx, s, composed.images)
	return nil
}

// fetchRelease downloads every pinned file the composition reads.
func fetchRelease() (releaseFileSet, error) {
	set := releaseFileSet{files: map[string]string{}}
	for _, path := range releaseFiles {
		text, err := fetchText(ReleaseFileURL(path))
		if err != nil {
			return releaseFileSet{}, err
		}
		set.files[path] = text
	}
	setupSQL, err := fetchText(SetupSQLURL())
	if err != nil {
		return releaseFileSet{}, err
	}
	set.setupSQL = setupSQL
	return set, nil
}
