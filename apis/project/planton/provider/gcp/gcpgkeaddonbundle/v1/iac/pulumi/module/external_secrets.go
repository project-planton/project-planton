package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	externalsecretsv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/externalsecrets/kubernetes/external_secrets/v1beta1"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// externalSecrets installs the External Secrets operator in the Kubernetes cluster using Helm, sets up the necessary
// Google Service Account (GSA), Kubernetes Service Account (KSA), and creates a ClusterSecretStore for GCP Secrets Manager.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdCluster: The GKE cluster where External Secrets will be installed.
// - gcpProvider: The GCP provider for Pulumi.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation.
//
// The function performs the following steps:
// 1. Creates a Google Service Account (GSA) for External Secrets with a description and display name.
// 2. Exports the email of the created GSA.
// 3. Creates a Workload Identity binding for the GSA to allow it to act as the Kubernetes Service Account (KSA).
// 4. Creates a namespace for External Secrets and labels it with metadata from locals.
// 5. Creates a Kubernetes Service Account (KSA) and adds the Google Workload Identity annotation with the GSA email.
// 6. Deploys the External Secrets Helm chart into the created namespace with specific values for CRDs, environment variables, and RBAC.
// 7. Creates a ClusterSecretStore to configure the GCP project from which secrets need to be looked up.
// 8. Handles errors and returns any errors encountered during the creation of resources or Helm release deployment.
func externalSecrets(ctx *pulumi.Context, locals *Locals, gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {

	//create google service account required to create workload identity binding
	createdGoogleServiceAccount, err := serviceaccount.NewAccount(ctx,
		vars.ExternalSecrets.KsaName,
		&serviceaccount.AccountArgs{
			Project:     pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
			Description: pulumi.String("external-secrets service account for solving dns challenges to issue certificates"),
			AccountId:   pulumi.String(vars.ExternalSecrets.KsaName),
			DisplayName: pulumi.String(vars.ExternalSecrets.KsaName),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create external-secrets google service account")
	}

	//export external-secrets gsa email
	ctx.Export(OpExternalSecretsGsaEmail, createdGoogleServiceAccount.Email)

	//add iam binding for secrets accessor role
	_, err = projects.NewIAMBinding(ctx,
		"external-secrets-secrets-accessor-binding",
		&projects.IAMBindingArgs{
			Members: pulumi.StringArray{
				pulumi.Sprintf("serviceAccount:%s", createdGoogleServiceAccount.Email),
			},
			Project: pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
			Role:    pulumi.String("roles/secretmanager.secretAccessor"),
		}, pulumi.Parent(createdGoogleServiceAccount))
	if err != nil {
		return errors.Wrap(err, "failed to add secrets accessor IAM binding")
	}

	//create workload-identity binding
	_, err = serviceaccount.NewIAMBinding(ctx,
		fmt.Sprintf("%s-workload-identity", vars.ExternalSecrets.KsaName),
		&serviceaccount.IAMBindingArgs{
			ServiceAccountId: createdGoogleServiceAccount.Name,
			Role:             pulumi.String("roles/iam.workloadIdentityUser"),
			Members: pulumi.StringArray{
				pulumi.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]",
					locals.GcpGkeAddonBundle.Spec.ClusterProjectId,
					vars.ExternalSecrets.Namespace,
					vars.ExternalSecrets.KsaName),
			},
		},
		pulumi.Parent(createdGoogleServiceAccount))
	if err != nil {
		return errors.Wrap(err, "failed to create workload-identity binding for external-secrets")
	}

	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.ExternalSecrets.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.ExternalSecrets.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed  namespace")
	}

	//create kubernetes service account to be used by the external-secrets.
	//it is not straight forward to add the gsa email as one of the helm values.
	// so, instead, disable service account creation in helm release and create it separately add
	// the Google workload identity annotation to the service account which requires the email id
	// of the Google service account added as part of IAM module.
	createdKubernetesServiceAccount, err := corev1.NewServiceAccount(ctx,
		vars.ExternalSecrets.KsaName,
		&corev1.ServiceAccountArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:      pulumi.String(vars.ExternalSecrets.KsaName),
					Namespace: createdNamespace.Metadata.Name(),
					Annotations: pulumi.StringMap{
						vars.WorkloadIdentityKubeAnnotationKey: createdGoogleServiceAccount.Email,
					},
				}),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes service account")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, "external-secrets",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.ExternalSecrets.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.ExternalSecrets.HelmChartName),
			Version:         pulumi.String(vars.ExternalSecrets.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"customResourceManagerDisabled": pulumi.Bool(false),
				"crds": pulumi.Map{
					"create": pulumi.Bool(true),
				},
				"env": pulumi.Map{
					"POLLER_INTERVAL_MILLISECONDS": pulumi.Int(vars.ExternalSecrets.SecretsPollingIntervalSeconds * 1000),
					"LOG_LEVEL":                    pulumi.String("info"),
					"LOG_MESSAGE_KEY":              pulumi.String("msg"),
					"METRICS_PORT":                 pulumi.Int(3001),
				},
				"rbac": pulumi.Map{
					"create": pulumi.Bool(true),
				},
				"serviceAccount": pulumi.Map{
					"create": pulumi.Bool(false),
					"name":   pulumi.String(vars.ExternalSecrets.KsaName),
				},
				"replicaCount": pulumi.Int(1),
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.ExternalSecrets.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.DependsOn([]pulumi.Resource{createdKubernetesServiceAccount}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}

	//create cluster-secret-store to configure the gcp project from which the secrets need to be looked up
	_, err = externalsecretsv1.NewClusterSecretStore(ctx, "cluster-secret-store",
		&externalsecretsv1.ClusterSecretStoreArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:   pulumi.String(vars.ExternalSecrets.GcpSecretsManagerClusterSecretStoreName),
				Labels: pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: externalsecretsv1.ClusterSecretStoreSpecArgs{
				Provider: externalsecretsv1.ClusterSecretStoreSpecProviderArgs{
					Gcpsm: externalsecretsv1.ClusterSecretStoreSpecProviderGcpsmArgs{
						ProjectID: pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
					},
				},
				RefreshInterval: pulumi.Int(vars.ExternalSecrets.SecretsPollingIntervalSeconds),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.DependsOn([]pulumi.Resource{createdKubernetesServiceAccount}))
	if err != nil {
		return errors.Wrap(err, "failed to create cluster-secret-store")
	}

	return nil
}
