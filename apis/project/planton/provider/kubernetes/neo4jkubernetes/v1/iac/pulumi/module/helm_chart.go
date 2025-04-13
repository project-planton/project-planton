package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"

	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs the official or community Neo4j Helm chart
// using the values derived from the user's inputs.
func helmChart(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
) error {
	// The container fields from the user's spec for convenience:
	container := locals.Neo4jKubernetes.Spec.Container

	// Install the Helm chart from a specified repo/URL.
	_, err := helmv3.NewChart(ctx,
		locals.Neo4jKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.Neo4jHelmChartName),
			Version:   pulumi.String(vars.Neo4jHelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values: pulumi.Map{
				"neo4j": pulumi.Map{
					"name": pulumi.String(locals.Neo4jKubernetes.Metadata.Name),
				},
				// Make sure to override the chart’s name to keep it consistent.
				"fullnameOverride": pulumi.String(locals.Neo4jKubernetes.Metadata.Name),

				// Some Helm charts handle resources and container config in slightly different ways.
				// We'll map container resources, disk size, etc.
				"image": pulumi.Map{
					// The official chart often expects "tag" for selecting the version of Neo4j.
					// We might allow this in the future if your proto includes an image version field.
					"tag": pulumi.String("5.5.0"),
				},
				"resources": containerresources.ConvertToPulumiMap(container.Resources),
				"persistence": pulumi.Map{
					"enabled": pulumi.Bool(container.IsPersistenceEnabled),
					"size":    pulumi.String(container.DiskSize),
				},

				// We’re referencing the password secret from admin_password.go
				"auth": pulumi.Map{
					"existingSecret":            pulumi.String(vars.Neo4jPasswordSecretName),
					"existingSecretPasswordKey": pulumi.String(vars.Neo4jPasswordSecretKey),
				},

				// Optional memory config from the proto for heap and page cache.
				// If not provided, the chart will apply default.
				"conf": pulumi.Map{
					"neo4j": pulumi.Map{
						"dbms.memory.heap.maxSize":   pulumi.String(locals.Neo4jKubernetes.Spec.MemoryConfig.HeapMax),
						"dbms.memory.pagecache.size": pulumi.String(locals.Neo4jKubernetes.Spec.MemoryConfig.PageCache),
					},
				},

				// Use provided labels for the pods (similar to the Redis reference).
				"podLabels": convertstringmaps.ConvertGoStringMapToPulumiMap(locals.Labels),
				"volumes": pulumi.Map{
					"data": pulumi.Map{
						"mode": pulumi.String("defaultStorageClass"),
					},
				},
			},
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.Neo4jHelmChartRepoUrl),
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create neo4j helm chart")
	}
	return nil
}
