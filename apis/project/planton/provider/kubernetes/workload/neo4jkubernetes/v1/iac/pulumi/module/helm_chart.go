package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs the Neo4j Helm chart with values derived from the spec.
func helmChart(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
) error {
	container := locals.Neo4jKubernetes.Spec.Container

	// honour ingress settings
	ingressEnabled := locals.Neo4jKubernetes.Spec.Ingress != nil &&
		locals.Neo4jKubernetes.Spec.Ingress.Enabled &&
		locals.Neo4jKubernetes.Spec.Ingress.Hostname != ""

	// optional external LB
	externalSvc := pulumi.Map{
		"enabled": pulumi.Bool(ingressEnabled),
	}
	if ingressEnabled {
		externalSvc["type"] = pulumi.String("LoadBalancer")
		externalSvc["annotations"] = pulumi.StringMap{
			"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressExternalHostname),
		}
	}

	_, err := helmv3.NewChart(ctx,
		locals.Neo4jKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.Neo4jHelmChartName),
			Version:   pulumi.String(vars.Neo4jHelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values: pulumi.Map{
				"neo4j": pulumi.Map{
					"name": pulumi.String(locals.Neo4jKubernetes.Metadata.Name),

					// let the chart create its own secret + password
					// (no passwordFromSecret / passwordKey provided)

					"resources": pulumi.Map{
						"cpu":    pulumi.String(container.Resources.Limits.Cpu),
						"memory": pulumi.String(container.Resources.Limits.Memory),
					},
					"acceptLicenseAgreement": pulumi.String("yes"),
				},

				"externalService": externalSvc,

				// persistence
				"volumes": pulumi.Map{
					"data": pulumi.Map{
						"mode": pulumi.String("defaultStorageClass"),
						"size": pulumi.String(container.DiskSize),
					},
				},

				// neo4j.conf overrides
				"config": pulumi.Map{
					"server.memory.heap.initial_size": pulumi.String(locals.Neo4jKubernetes.Spec.MemoryConfig.HeapMax),
					"server.memory.pagecache.size":    pulumi.String(locals.Neo4jKubernetes.Spec.MemoryConfig.PageCache),
				},

				"podLabels": convertstringmaps.ConvertGoStringMapToPulumiMap(locals.Labels),
			},
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.Neo4jHelmChartRepoUrl),
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to deploy neo4j helm chart")
	}

	// ---------------------------------------------------------------------
	// Export outputs
	// ---------------------------------------------------------------------
	ctx.Export(OpUsername, pulumi.String("neo4j"))

	// the chart creates: <release>-auth  with key "neo4j-password"
	secretName := fmt.Sprintf("%s-auth", locals.Neo4jKubernetes.Metadata.Name)
	ctx.Export(OpPasswordSecretName, pulumi.String(secretName))
	ctx.Export(OpPasswordSecretKey, pulumi.String("neo4j-password"))

	return nil
}
