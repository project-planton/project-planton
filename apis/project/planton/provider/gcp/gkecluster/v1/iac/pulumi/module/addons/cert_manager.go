package addons

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	certmanagerv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/certmanager/kubernetes/cert_manager/v1"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CertManager installs Cert Manager in the Kubernetes cluster using Helm, sets up the necessary Google Service Account (GSA),
// Kubernetes Service Account (KSA), and creates a self-signed ClusterIssuer.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdCluster: The GKE cluster where Cert Manager will be installed.
// - gcpProvider: The GCP provider for Pulumi.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation.
//
// The function performs the following steps:
// 1. Creates a Google Service Account (GSA) for Cert Manager with a description and display name.
// 2. Exports the email of the created GSA.
// 3. Creates a Workload Identity binding for the GSA to allow it to act as the Kubernetes Service Account (KSA).
// 4. Creates a namespace for Cert Manager and labels it with metadata from locals.
// 5. Creates a Kubernetes Service Account (KSA) and adds the Google Workload Identity annotation with the GSA email.
// 6. Deploys the Cert Manager Helm chart into the created namespace with specific values for CRDs, service account.
func CertManager(ctx *pulumi.Context, locals *localz.Locals,
	createdCluster *container.Cluster,
	gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {

	//create google service account required to create workload identity binding
	createdGoogleServiceAccount, err := serviceaccount.NewAccount(ctx,
		vars.CertManager.KsaName,
		&serviceaccount.AccountArgs{
			Project:     createdCluster.Project,
			Description: pulumi.String("cert-manager service account for solving dns challenges to issue certificates"),
			AccountId:   pulumi.String(vars.CertManager.KsaName),
			DisplayName: pulumi.String(vars.CertManager.KsaName),
		}, pulumi.Parent(createdCluster), pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create cert-manager google service account")
	}

	//export cert-manager gsa email
	ctx.Export(outputs.CertManagerGsaEmail, createdGoogleServiceAccount.Email)

	//create workload-identity binding
	_, err = serviceaccount.NewIAMBinding(ctx,
		fmt.Sprintf("%s-workload-identity", vars.CertManager.KsaName),
		&serviceaccount.IAMBindingArgs{
			ServiceAccountId: createdGoogleServiceAccount.Name,
			Role:             pulumi.String("roles/iam.workloadIdentityUser"),
			Members: pulumi.StringArray{
				pulumi.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]",
					createdCluster.Project,
					vars.CertManager.Namespace,
					vars.CertManager.KsaName),
			},
		},
		pulumi.Parent(createdGoogleServiceAccount),
		pulumi.DependsOn([]pulumi.Resource{createdCluster}))
	if err != nil {
		return errors.Wrap(err, "failed to create workload-identity binding for cert-manager")
	}

	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.CertManager.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.CertManager.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create cert-manager namespace")
	}

	//it is not straight forward to add the gsa email as one of the helm values.
	// so, instead, disable service account creation in helm release and create it separately add
	// the Google workload identity annotation to the service account which requires the email id
	// of the Google service account added as part of IAM module.
	createdKubernetesServiceAccount, err := corev1.NewServiceAccount(ctx,
		vars.CertManager.KsaName,
		&corev1.ServiceAccountArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:      pulumi.String(vars.CertManager.KsaName),
					Namespace: createdNamespace.Metadata.Name(),
					Annotations: pulumi.StringMap{
						vars.WorkloadIdentityKubeAnnotationKey: createdGoogleServiceAccount.Email,
					},
				}),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes service account")
	}

	//created helm-release
	createdCertManagerHelmRelease, err := helm.NewRelease(ctx, "cert-manager",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.CertManager.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.CertManager.HelmChartName),
			Version:         pulumi.String(vars.CertManager.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"installCRDs": pulumi.Bool(true),
				//https://cert-manager.io/docs/configuration/acme/dns01/#setting-nameservers-for-dns01-self-check
				//https://github.com/cert-manager/cert-manager/issues/1163#issuecomment-484171354
				"extraArgs": pulumi.StringArray{
					pulumi.String("--dns01-recursive-nameservers-only=true"),
					pulumi.String("--dns01-recursive-nameservers=8.8.8.8:53,1.1.1.1:53"),
				},
				"serviceAccount": pulumi.Map{
					"create": pulumi.Bool(false),
					"name":   pulumi.String(vars.CertManager.KsaName),
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.CertManager.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.DependsOn([]pulumi.Resource{createdKubernetesServiceAccount}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create cert-manager helm release")
	}

	//for each ingress-domain, create a cluster-issuer
	for _, i := range locals.GkeCluster.Spec.IngressDnsDomains {
		//do not create a cluster-issuer resource if tls is not enabled.
		if !i.IsTlsEnabled {
			continue
		}
		_, err := certmanagerv1.NewClusterIssuer(ctx,
			i.Name,
			&certmanagerv1.ClusterIssuerArgs{
				Metadata: metav1.ObjectMetaArgs{
					Name:   pulumi.String(i.Name),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				},
				Spec: certmanagerv1.ClusterIssuerSpecArgs{
					Acme: certmanagerv1.ClusterIssuerSpecAcmeArgs{
						PreferredChain: pulumi.String(""),
						PrivateKeySecretRef: certmanagerv1.ClusterIssuerSpecAcmePrivateKeySecretRefArgs{
							Name: pulumi.String(vars.CertManager.LetsEncryptClusterIssuerSecretName),
						},
						Server: pulumi.String(vars.CertManager.LetsEncryptServer),
						Solvers: certmanagerv1.ClusterIssuerSpecAcmeSolversArray{
							certmanagerv1.ClusterIssuerSpecAcmeSolversArgs{
								Dns01: certmanagerv1.ClusterIssuerSpecAcmeSolversDns01Args{
									CloudDNS: certmanagerv1.ClusterIssuerSpecAcmeSolversDns01CloudDNSArgs{
										Project: pulumi.String(i.DnsZoneGcpProjectId),
									},
								},
							},
							certmanagerv1.ClusterIssuerSpecAcmeSolversArgs{
								Http01: certmanagerv1.ClusterIssuerSpecAcmeSolversHttp01Args{
									Ingress: certmanagerv1.ClusterIssuerSpecAcmeSolversHttp01IngressArgs{
										Class: pulumi.String(vars.CertManager.Http01ChallengeSolverIngressClass),
									},
								},
							},
						},
					},
				},
			}, pulumi.Parent(createdCertManagerHelmRelease))
		if err != nil {
			return errors.Wrapf(err, "failed to create cluster-issuer for %s ingress-domain", i.Name)
		}
	}

	return nil
}