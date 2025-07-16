package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// nodePool provisions the Kubernetes node‑pool and exports its IDs.
func nodePool(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.KubernetesNodePool, error) {

	// 1. Build argument structs directly from proto fields.
	labels := pulumi.StringMap{}
	for k, v := range locals.DigitalOceanLabels {
		labels[k] = pulumi.String(v)
	}

	tags := pulumi.StringArray{}
	for _, t := range locals.DigitalOceanKubernetesNodePool.Spec.Tags {
		tags = append(tags, pulumi.String(t))
	}

	nodePoolArgs := &digitalocean.KubernetesNodePoolArgs{
		ClusterId: pulumi.String(locals.DigitalOceanKubernetesNodePool.Spec.Cluster.GetValue()),
		Name:      pulumi.String(locals.DigitalOceanKubernetesNodePool.Spec.NodePoolName),
		Size:      pulumi.String(locals.DigitalOceanKubernetesNodePool.Spec.SizeSlug),
		Labels:    labels,
		Tags:      tags,
		NodeCount: pulumi.IntPtr(int(locals.DigitalOceanKubernetesNodePool.Spec.NodeCount)),
	}

	if locals.DigitalOceanKubernetesNodePool.Spec.AutoScale {
		nodePoolArgs.AutoScale = pulumi.BoolPtr(true)
		nodePoolArgs.MinNodes = pulumi.IntPtr(int(locals.DigitalOceanKubernetesNodePool.Spec.MinNodes))
		nodePoolArgs.MaxNodes = pulumi.IntPtr(int(locals.DigitalOceanKubernetesNodePool.Spec.MaxNodes))
	}

	// 2. Create the node‑pool.
	createdNodePool, err := digitalocean.NewKubernetesNodePool(
		ctx,
		"node_pool",
		nodePoolArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean kubernetes node pool")
	}

	// 3. Export stack outputs.
	ctx.Export(OpNodePoolId, createdNodePool.ID())

	return createdNodePool, nil
}
