package module

import (
	"github.com/pkg/errors"
	gcpcloudsqlv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudsql/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi program entry-point for the GcpCloudSql component.
func Resources(ctx *pulumi.Context, stackInput *gcpcloudsqlv1.GcpCloudSqlStackInput) error {
	locals := initializeLocals(stackInput)

	// Set up the GCP provider from the supplied credential.
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup google provider")
	}

	// Create the Cloud SQL database instance.
	createdInstance, err := databaseInstance(ctx, locals, gcpProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create cloud-sql instance")
	}

	// Export stack outputs
	ctx.Export(OpInstanceName, createdInstance.Name)
	ctx.Export(OpConnectionName, createdInstance.ConnectionName)
	ctx.Export(OpSelfLink, createdInstance.SelfLink)

	// Export IPs if they exist
	createdInstance.IpAddresses.ApplyT(func(ipAddresses interface{}) error {
		if ipAddresses == nil {
			return nil
		}

		// Export first public IP and first private IP found
		for _, ipAddress := range ipAddresses.([]interface{}) {
			ipMap := ipAddress.(map[string]interface{})
			ipType := ipMap["type"].(string)
			ipAddr := ipMap["ipAddress"].(string)

			if ipType == "PRIMARY" {
				ctx.Export(OpPublicIp, pulumi.String(ipAddr))
			} else if ipType == "PRIVATE" {
				ctx.Export(OpPrivateIp, pulumi.String(ipAddr))
			}
		}
		return nil
	})

	return nil
}
