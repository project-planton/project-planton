package module

import (
	"github.com/pkg/errors"
	externalsecretsv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/addon/externalsecretskubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources installs the External‑Secrets operator (ESO) into the target cluster.
func Resources(ctx *pulumi.Context, in *externalsecretsv1.ExternalSecretsKubernetesStackInput) error {
	// set up provider from credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, in.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	spec := in.Target.Spec
	// choose chart version – today we expose only a stable channel
	chartVersion := vars.DefaultStableVersion

	// 1. namespace
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

	// 2. identity annotations
	annotations := pulumi.StringMap{}
	var identity pulumi.StringInput

	if gke := spec.GetGke(); gke != nil && gke.GsaEmail != "" {
		annotations["iam.gke.io/gcp-service-account"] = pulumi.String(gke.GsaEmail)
		identity = pulumi.String(gke.GsaEmail)
	} else if eks := spec.GetEks(); eks != nil && eks.IrsaRoleArnOverride != "" {
		annotations["eks.amazonaws.com/role-arn"] = pulumi.String(eks.IrsaRoleArnOverride)
		identity = pulumi.String(eks.IrsaRoleArnOverride)
	} else if aks := spec.GetAks(); aks != nil && aks.ManagedIdentityClientId != "" {
		annotations["azure.workload.identity/client-id"] = pulumi.String(aks.ManagedIdentityClientId)
		identity = pulumi.String(aks.ManagedIdentityClientId)
	}

	// 3. service account
	sa, err := corev1.NewServiceAccount(ctx, vars.KsaName,
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

	// 4. translate optional container resources into helm values
	resVals := pulumi.Map{}
	if c := spec.GetContainer(); c != nil && c.Resources != nil {
		req := c.Resources.GetRequests()
		lim := c.Resources.GetLimits()

		if req != nil || lim != nil {
			reqVals := pulumi.Map{}
			if req != nil {
				if req.Cpu != "" {
					reqVals["cpu"] = pulumi.String(req.Cpu)
				}
				if req.Memory != "" {
					reqVals["memory"] = pulumi.String(req.Memory)
				}
			}
			limVals := pulumi.Map{}
			if lim != nil {
				if lim.Cpu != "" {
					limVals["cpu"] = pulumi.String(lim.Cpu)
				}
				if lim.Memory != "" {
					limVals["memory"] = pulumi.String(lim.Memory)
				}
			}
			resVals = pulumi.Map{
				"requests": reqVals,
				"limits":   limVals,
			}
		}
	}

	// 5. helm release
	_, err = helm.NewRelease(ctx, "external-secrets",
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
					"name":   pulumi.String(vars.KsaName),
					"create": pulumi.Bool(false),
				},
				"env": pulumi.Map{
					"POLLER_INTERVAL_MILLISECONDS": pulumi.Int(vars.DefaultSecretsPollIntervalSec * 1000),
				},
				"rbac":      pulumi.Map{"create": pulumi.Bool(true)},
				"resources": resVals,
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns),
		pulumi.DependsOn([]pulumi.Resource{sa}))
	if err != nil {
		return errors.Wrap(err, "failed to install external‑secrets helm release")
	}

	// 6. stack outputs
	ctx.Export(OpNamespace, ns.Metadata.Name())
	ctx.Export(OpReleaseName, pulumi.String(vars.HelmChartName))
	ctx.Export(OpOperatorSA, sa.Metadata.Name())
	if identity != nil {
		ctx.Export(OpIdentity, identity)
	}

	return nil
}
