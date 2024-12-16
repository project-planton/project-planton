package addons

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

func ExternalDns(ctx *pulumi.Context, locals *localz.Locals,
	createdCluster *container.Cluster, gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {

	//create google service account required to create workload identity binding
	createdGoogleServiceAccount, err := serviceaccount.NewAccount(ctx,
		vars.ExternalDns.KsaName,
		&serviceaccount.AccountArgs{
			Project:     createdCluster.Project,
			Description: pulumi.String("external-dns service account for managing dns-records in cloud dns zones"),
			AccountId:   pulumi.String(vars.ExternalDns.KsaName),
			DisplayName: pulumi.String(vars.ExternalDns.KsaName),
		}, pulumi.Parent(createdCluster), pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create external-dns google service account")
	}

	//export cert-manager gsa email
	ctx.Export(outputs.EXTERNAL_DNS_GSA_EMAIL, createdGoogleServiceAccount.Email)

	//create workload-identity binding
	_, err = serviceaccount.NewIAMBinding(ctx,
		fmt.Sprintf("%s-workload-identity", vars.ExternalDns.KsaName),
		&serviceaccount.IAMBindingArgs{
			ServiceAccountId: createdGoogleServiceAccount.Name,
			Role:             pulumi.String("roles/iam.workloadIdentityUser"),
			Members: pulumi.StringArray{
				pulumi.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]",
					createdCluster.Project,
					vars.ExternalDns.Namespace,
					vars.ExternalDns.KsaName),
			},
		},
		pulumi.Parent(createdGoogleServiceAccount),
		pulumi.DependsOn([]pulumi.Resource{createdCluster}))
	if err != nil {
		return errors.Wrap(err, "failed to create workload-identity binding for external-dns")
	}

	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.ExternalDns.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.ExternalDns.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create external-dns namespace")
	}

	//create kubernetes service account to be used by the external-dns.
	//it is not straight forward to add the gsa email as one of the helm values.
	// so, instead, disable service account creation in helm release and create it separately add
	// the Google workload identity annotation to the service account which requires the email id
	// of the Google service account added as part of IAM module.
	createdKubernetesServiceAccount, err := corev1.NewServiceAccount(ctx,
		vars.ExternalDns.KsaName,
		&corev1.ServiceAccountArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:      pulumi.String(vars.ExternalDns.KsaName),
					Namespace: createdNamespace.Metadata.Name(),
					Annotations: pulumi.StringMap{
						vars.WorkloadIdentityKubeAnnotationKey: createdGoogleServiceAccount.Email,
					},
				}),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes service account")
	}

	for _, i := range locals.GkeCluster.Spec.IngressDnsDomains {
		//created helm-release
		_, err := helm.NewRelease(ctx,
			fmt.Sprintf("external-dns-%s", i.Name),
			&helm.ReleaseArgs{
				Name:            pulumi.Sprintf("external-dns-%s", strings.ReplaceAll(i.Name, ".", "-")),
				Namespace:       createdNamespace.Metadata.Name(),
				Chart:           pulumi.String(vars.ExternalDns.HelmChartName),
				Version:         pulumi.String(vars.ExternalDns.HelmChartVersion),
				CreateNamespace: pulumi.Bool(false),
				Atomic:          pulumi.Bool(false),
				CleanupOnFail:   pulumi.Bool(true),
				WaitForJobs:     pulumi.Bool(true),
				Timeout:         pulumi.Int(180),
				Values: pulumi.Map{
					"txtOwnerId": pulumi.String(locals.GkeCluster.Metadata.Name),
					"serviceAccount": pulumi.Map{
						"create": pulumi.Bool(false),
						"name":   pulumi.String(vars.ExternalDns.KsaName),
					},
					"domainFilters": pulumi.ToStringArray([]string{
						i.Name,
					}),
					//https://kubernetes-sigs.github.io/external-dns/v0.13.1/tutorials/gateway-api/#manifest-with-rbac
					"sources": pulumi.StringArray{
						pulumi.String("service"),
						pulumi.String("gateway-httproute"),
					},
					"provider": pulumi.String(vars.ExternalDns.GcpCloudDnsProviderName),
					"extraArgs": pulumi.StringArray{
						pulumi.String("--google-zone-visibility=public"),
						pulumi.Sprintf("--google-project=%s", i.DnsZoneGcpProjectId),
					},
				},
				RepositoryOpts: helm.RepositoryOptsArgs{
					Repo: pulumi.String(vars.ExternalDns.HelmChartRepo),
				},
			}, pulumi.Parent(createdNamespace),
			pulumi.DependsOn([]pulumi.Resource{createdKubernetesServiceAccount}),
			pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
		if err != nil {
			return errors.Wrap(err, "failed to create external-dns helm release")
		}
	}
	return nil
}
