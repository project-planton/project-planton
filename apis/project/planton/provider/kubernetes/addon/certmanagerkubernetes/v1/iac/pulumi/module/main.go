package module

import (
	"github.com/pkg/errors"
	certmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/certmanagerkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources create all Pulumi resources for the Cert‑Manager Kubernetes add‑on.
func Resources(ctx *pulumi.Context, stackInput *certmanagerv1.CertManagerKubernetesStackInput) error {
	// set up a kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	spec := stackInput.Target.Spec

	// pick helm chart version based on release_channel
	chartVersion := vars.DefaultStableVersion
	releaseChannel := spec.GetReleaseChannel()
	switch releaseChannel {
	case "", "stable":
		chartVersion = vars.DefaultStableVersion
	case "latest", "edge", "fast":
		chartVersion = vars.DefaultLatestVersion
	default:
		chartVersion = releaseChannel // explicit tag such as v1.16.1
	}

	// create cert‑manager namespace
	ns, err := corev1.NewNamespace(ctx, vars.Namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.Namespace),
			},
		},
		pulumi.Provider(kubeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// build identity annotation map
	annotations := pulumi.StringMap{}
	var solverIdentity pulumi.StringInput

	if gke := spec.GetGke(); gke != nil {
		annotations["iam.gke.io/gcp-service-account"] = pulumi.String(gke.GsaEmail)
		solverIdentity = pulumi.String(gke.GsaEmail)
	} else if eks := spec.GetEks(); eks != nil {
		roleArn := eks.IrsaRoleArnOverride
		if roleArn == "" {
			return errors.New("eks.irsa_role_arn_override must be set (auto‑creation not implemented)")
		}
		annotations["eks.amazonaws.com/role-arn"] = pulumi.String(roleArn)
		solverIdentity = pulumi.String(roleArn)
	} else if aks := spec.GetAks(); aks != nil && aks.ManagedIdentityClientId != "" {
		annotations["azure.workload.identity/client-id"] = pulumi.String(aks.ManagedIdentityClientId)
		solverIdentity = pulumi.String(aks.ManagedIdentityClientId)
	}

	// create a ServiceAccount with the chosen annotations
	_, err = corev1.NewServiceAccount(ctx, vars.KsaName,
		&corev1.ServiceAccountArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:        pulumi.String(vars.KsaName),
				Namespace:   ns.Metadata.Name(),
				Annotations: annotations,
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns))
	if err != nil {
		return errors.Wrap(err, "failed to create service account")
	}

	// deploy cert‑manager helm chart
	_, err = helm.NewRelease(ctx, "cert-manager",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       ns.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"installCRDs": pulumi.Bool(true),
				"serviceAccount": pulumi.Map{
					"create": pulumi.Bool(false),
					"name":   pulumi.String(vars.KsaName),
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns))
	if err != nil {
		return errors.Wrap(err, "failed to install cert-manager helm release")
	}

	// export stack outputs
	ctx.Export(OpNamespace, ns.Metadata.Name())
	ctx.Export(OpReleaseName, pulumi.String(vars.HelmChartName))
	if solverIdentity != nil {
		ctx.Export(OpSolverIdentity, solverIdentity)
	}

	return nil
}
