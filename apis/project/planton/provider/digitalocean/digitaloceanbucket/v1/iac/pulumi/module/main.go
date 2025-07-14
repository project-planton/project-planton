package module

import (
	"github.com/pkg/errors"
	digitaloceanbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanbucket/v1"
	"github.com/pulumi/pulumi-digitalocean/sdk/v5/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *digitaloceanbucketv1.DigitalOceanBucketStackInput) error {
	digitaloceanCredential := stackInput.ProviderCredential
	//create digitalocean provider using the credentials from the input
	_, err := digitalocean.NewProvider(ctx,
		"digitalocean",
		&digitalocean.ProviderArgs{
			ClientId:       pulumi.String(digitaloceanCredential.ClientId),
			ClientSecret:   pulumi.String(digitaloceanCredential.ClientSecret),
			SubscriptionId: pulumi.String(digitaloceanCredential.SubscriptionId),
			TenantId:       pulumi.String(digitaloceanCredential.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create digitalocean provider")
	}
	return nil
}
