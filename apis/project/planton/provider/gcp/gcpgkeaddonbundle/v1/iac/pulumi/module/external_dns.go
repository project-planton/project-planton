package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func externalDns(ctx *pulumi.Context, locals *Locals,
	gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {

	//create google service account required to create workload identity binding
	createdGoogleServiceAccount, err := serviceaccount.NewAccount(ctx,
		vars.ExternalDns.KsaName,
		&serviceaccount.AccountArgs{
			Project:     pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
			Description: pulumi.String("external-dns service account for managing dns-records in cloud dns zones"),
			AccountId:   pulumi.String(vars.ExternalDns.KsaName),
			DisplayName: pulumi.String(vars.ExternalDns.KsaName),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create external-dns google service account")
	}

	//export cert-manager gsa email
	ctx.Export(outputs.ExternalDnsGsaEmail, createdGoogleServiceAccount.Email)

	//create workload-identity binding
	_, err = serviceaccount.NewIAMBinding(ctx,
		fmt.Sprintf("%s-workload-identity", vars.ExternalDns.KsaName),
		&serviceaccount.IAMBindingArgs{
			ServiceAccountId: createdGoogleServiceAccount.Name,
			Role:             pulumi.String("roles/iam.workloadIdentityUser"),
			Members: pulumi.StringArray{
				pulumi.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]",
					locals.GcpGkeAddonBundle.Spec.ClusterProjectId,
					vars.ExternalDns.Namespace,
					vars.ExternalDns.KsaName),
			},
		},
		pulumi.Parent(createdGoogleServiceAccount))
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
	_, err = corev1.NewServiceAccount(ctx,
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
	return nil
}
