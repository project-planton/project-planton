package module

import (
	"fmt"
	"github.com/pkg/errors"
	gcpdnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpdnszone/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpdnszone/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dns"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

func Resources(ctx *pulumi.Context, stackInput *gcpdnszonev1.GcpDnsZoneStackInput) error {
	//create gcp provider using credentials from the input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create a descriptive variable for the api-resource in the input
	gcpDnsZone := stackInput.Target

	//replace dots with hyphens to create valid managed-zone name
	managedZoneName := strings.ReplaceAll(gcpDnsZone.Metadata.Name, ".", "-")

	//create managed-zone
	createdManagedZone, err := dns.NewManagedZone(ctx,
		managedZoneName,
		&dns.ManagedZoneArgs{
			Name:        pulumi.String(managedZoneName),
			Project:     pulumi.String(gcpDnsZone.Spec.ProjectId),
			Description: pulumi.String(fmt.Sprintf("managed-zone for %s", gcpDnsZone.Metadata.Name)),
			//dns-name should have a dot at the end
			DnsName:    pulumi.Sprintf("%s.", gcpDnsZone.Metadata.Name),
			Visibility: pulumi.String("public"),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to add zone for %s domain", gcpDnsZone.Metadata.Name)
	}

	//export important managed-zone attributes as outputs
	ctx.Export(outputs.MANAGED_ZONE_NAME, createdManagedZone.Name)
	ctx.Export(outputs.GCP_PROJECT_ID, createdManagedZone.Project)
	// todo: export nameservers as per new outputs handler logic

	//create IAM binding for the gcp service-accounts to be granted permissions to manage the records in the zone.
	//with this binding each gcp service-account will be granted permissions to create/delete/update dns-records.
	if len(gcpDnsZone.Spec.IamServiceAccounts) > 0 {
		//each service account needs to be prefixed w/ 'serviceAccount:' for the input in the binding resource.
		iamBindingMembers := make([]pulumi.StringInput, 0)
		for _, serviceAccountEmail := range gcpDnsZone.Spec.IamServiceAccounts {
			iamBindingMembers = append(iamBindingMembers,
				pulumi.Sprintf("serviceAccount:%s", serviceAccountEmail))
		}

		// todo: the correct resource to use is https://cloud.google.com/dns/docs/zones/iam-per-resource-zones#gcloud
		// but the resource is not yet available in the gcp provider.
		// as a temporary workaround, granting dns admin role to all the service accounts to the entire project.
		// this method grants much broader permissions which allow the service account to control all the zones in the project.
		_, err := projects.NewIAMBinding(ctx,
			managedZoneName,
			&projects.IAMBindingArgs{
				Members: pulumi.StringArray(iamBindingMembers),
				Project: createdManagedZone.Project,
				Role:    pulumi.String("roles/dns.admin"),
			}, pulumi.Parent(createdManagedZone))
		if err != nil {
			return errors.Wrapf(err, "failed to create dns-admin iam-binding resource on gcp-project")
		}
	}

	//create dns-records in the created managed-zone
	for index, dnsRecord := range gcpDnsZone.Spec.Records {
		_, err := dns.NewRecordSet(ctx,
			fmt.Sprintf("dns-record-%d", index),
			&dns.RecordSetArgs{
				ManagedZone: createdManagedZone.Name,
				Name:        pulumi.String(dnsRecord.Name),
				Project:     createdManagedZone.Project,
				Rrdatas:     pulumi.ToStringArray(dnsRecord.Values),
				Ttl:         pulumi.IntPtr(int(dnsRecord.TtlSeconds)),
				Type:        pulumi.String(dnsRecord.RecordType.String()),
			}, pulumi.Parent(createdManagedZone))
		if err != nil {
			return errors.Wrapf(err, "failed to add %s rec", dnsRecord)
		}
	}
	return nil
}
