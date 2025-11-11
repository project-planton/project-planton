package module

import (
	"fmt"

	"github.com/pkg/errors"
	digitaloceandatabaseclusterv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/digitalocean/digitaloceandatabasecluster/v1"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// cluster provisions the managed database cluster and exports its outputs.
func cluster(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.DatabaseCluster, error) {

	// 1. Translate proto enum to engine slug.
	var engineSlug string
	switch locals.DigitalOceanDatabaseCluster.Spec.Engine {
	case digitaloceandatabaseclusterv1.DigitalOceanDatabaseEngine_postgres:
		engineSlug = "pg"
	case digitaloceandatabaseclusterv1.DigitalOceanDatabaseEngine_mysql:
		engineSlug = "mysql"
	case digitaloceandatabaseclusterv1.DigitalOceanDatabaseEngine_redis:
		engineSlug = "redis"
	case digitaloceandatabaseclusterv1.DigitalOceanDatabaseEngine_mongodb:
		engineSlug = "mongodb"
	default:
		return nil, errors.Errorf("unsupported database engine: %v", locals.DigitalOceanDatabaseCluster.Spec.Engine)
	}

	// 2. Convert label map → slice of "key:value" tags.
	var tagInputs pulumi.StringArray
	if len(locals.DigitalOceanLabels) > 0 {
		for k, v := range locals.DigitalOceanLabels {
			tagInputs = append(tagInputs, pulumi.String(k+":"+v))
		}
	}

	// 3. Build resource arguments straight from proto fields.
	clusterArgs := &digitalocean.DatabaseClusterArgs{
		Engine:    pulumi.String(engineSlug),
		Name:      pulumi.String(locals.DigitalOceanDatabaseCluster.Spec.ClusterName),
		Region:    pulumi.String(locals.DigitalOceanDatabaseCluster.Spec.Region.String()),
		Version:   pulumi.String(locals.DigitalOceanDatabaseCluster.Spec.EngineVersion),
		Size:      pulumi.String(locals.DigitalOceanDatabaseCluster.Spec.SizeSlug),
		NodeCount: pulumi.Int(int(locals.DigitalOceanDatabaseCluster.Spec.NodeCount)),
		Tags:      tagInputs,
	}

	// Optional storage override.
	if locals.DigitalOceanDatabaseCluster.Spec.StorageGib != 0 {
		clusterArgs.StorageSizeMib = pulumi.String(fmt.Sprintf("%dMib",
			locals.DigitalOceanDatabaseCluster.Spec.StorageGib*1024))
	}

	// Optional VPC attachment.
	if locals.DigitalOceanDatabaseCluster.Spec.Vpc != nil &&
		locals.DigitalOceanDatabaseCluster.Spec.Vpc.GetValue() != "" {
		clusterArgs.PrivateNetworkUuid = pulumi.StringPtr(locals.DigitalOceanDatabaseCluster.Spec.Vpc.GetValue())
	}

	// NOTE: DigitalOcean API does not yet expose a direct "disable public network" switch.
	// Leaving the flag as‑is; see explanation in the differences section.
	_ = locals.DigitalOceanDatabaseCluster.Spec.EnablePublicConnectivity

	// 4. Provision the cluster.
	createdCluster, err := digitalocean.NewDatabaseCluster(
		ctx,
		"cluster",
		clusterArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean database cluster")
	}

	// 5. Export stack outputs.
	ctx.Export(OpClusterId, createdCluster.ID())
	ctx.Export(OpConnectionUri, createdCluster.Uri)
	ctx.Export(OpHost, createdCluster.Host)
	ctx.Export(OpPort, createdCluster.Port)
	ctx.Export(OpDatabaseUser, createdCluster.User)
	ctx.Export(OpDatabasePassword, createdCluster.Password)

	return createdCluster, nil
}
