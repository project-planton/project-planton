package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// cluster provisions the Kubernetes cluster itself and exports its outputs.
func cluster(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.KubernetesCluster, error) {

	// 1. Collect tags from the spec.
	var tags pulumi.StringArray
	for _, t := range locals.DigitalOceanKubernetesCluster.Spec.Tags {
		tags = append(tags, pulumi.String(t))
	}

	// 4. Build the cluster arguments straight from proto fields.
	clusterArgs := &digitalocean.KubernetesClusterArgs{
		Name:         pulumi.String(locals.DigitalOceanKubernetesCluster.Spec.ClusterName),
		Region:       pulumi.String(locals.DigitalOceanKubernetesCluster.Spec.Region.String()),
		Version:      pulumi.String(locals.DigitalOceanKubernetesCluster.Spec.KubernetesVersion),
		Ha:           pulumi.BoolPtr(locals.DigitalOceanKubernetesCluster.Spec.HighlyAvailable),
		AutoUpgrade:  pulumi.BoolPtr(locals.DigitalOceanKubernetesCluster.Spec.AutoUpgrade),
		SurgeUpgrade: pulumi.BoolPtr(!locals.DigitalOceanKubernetesCluster.Spec.DisableSurgeUpgrade),
		VpcUuid:      pulumi.String(locals.DigitalOceanKubernetesCluster.Spec.Vpc.GetValue()),
		Tags:         tags,
		NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
			AutoScale: pulumi.BoolPtr(false),
			MaxNodes:  pulumi.IntPtr(2),
			MinNodes:  pulumi.IntPtr(1),
			Name:      pulumi.String("default"),
			NodeCount: pulumi.IntPtr(1),
			Size:      pulumi.String("s-2vcpu-4gb"),
		},
	}

	// 5. Create the cluster.
	createdCluster, err := digitalocean.NewKubernetesCluster(
		ctx,
		"cluster",
		clusterArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean kubernetes cluster")
	}

	// 6. Export required stack outputs.
	ctx.Export(OpClusterId, createdCluster.ID())
	ctx.Export(OpApiServerEndpoint, createdCluster.Endpoint)

	// Use the first kubeconfig in the list.
	kubeconfig := createdCluster.KubeConfigs.Index(pulumi.Int(0)).RawConfig()
	ctx.Export(OpKubeconfig, kubeconfig)

	return createdCluster, nil
}
