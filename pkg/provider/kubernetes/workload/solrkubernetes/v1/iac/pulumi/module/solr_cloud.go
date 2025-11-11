package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/solroperator/kubernetes/solr/v1beta1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func solrCloud(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {
	//create solr-operator's solrcloud resource
	_, err := v1beta1.NewSolrCloud(ctx, "solr-cloud",
		&v1beta1.SolrCloudArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.SolrKubernetes.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: v1beta1.SolrCloudSpecArgs{
				Replicas: pulumi.Int(locals.SolrKubernetes.Spec.SolrContainer.Replicas),
				SolrImage: v1beta1.SolrCloudSpecSolrImageArgs{
					Repository: pulumi.String(locals.SolrKubernetes.Spec.SolrContainer.Image.Repo),
					Tag:        pulumi.String(locals.SolrKubernetes.Spec.SolrContainer.Image.Tag),
				},
				SolrJavaMem: pulumi.String(locals.SolrKubernetes.Spec.Config.JavaMem),
				SolrOpts:    pulumi.String(locals.SolrKubernetes.Spec.Config.Opts),
				SolrGCTune:  pulumi.String(locals.SolrKubernetes.Spec.Config.GarbageCollectionTuning),
				SolrModules: pulumi.ToStringArray(vars.SolrCloudSolrModules),
				CustomSolrKubeOptions: v1beta1.SolrCloudSpecCustomSolrKubeOptionsArgs{
					PodOptions: v1beta1.SolrCloudSpecCustomSolrKubeOptionsPodOptionsArgs{
						Resources: v1beta1.SolrCloudSpecCustomSolrKubeOptionsPodOptionsResourcesArgs{
							Limits: pulumi.ToMap(map[string]interface{}{
								//"cpu":    locals.SolrKubernetes.Spec.SolrContainer.Resources.Limits.Cpu,
								"memory": locals.SolrKubernetes.Spec.SolrContainer.Resources.Limits.Memory,
							}),
							Requests: pulumi.ToMap(map[string]interface{}{
								"cpu":    locals.SolrKubernetes.Spec.SolrContainer.Resources.Requests.Cpu,
								"memory": locals.SolrKubernetes.Spec.SolrContainer.Resources.Requests.Memory,
							}),
						},
					},
				},
				DataStorage: v1beta1.SolrCloudSpecDataStorageArgs{
					Ephemeral: nil,
					Persistent: v1beta1.SolrCloudSpecDataStoragePersistentArgs{
						ReclaimPolicy: pulumi.String("Delete"),
						PvcTemplate: v1beta1.SolrCloudSpecDataStoragePersistentPvcTemplateArgs{
							Spec: v1beta1.SolrCloudSpecDataStoragePersistentPvcTemplateSpecArgs{
								Resources: v1beta1.SolrCloudSpecDataStoragePersistentPvcTemplateSpecResourcesArgs{
									Requests: pulumi.ToMap(map[string]interface{}{
										"storage": locals.SolrKubernetes.Spec.SolrContainer.DiskSize,
									}),
								},
							},
						},
					},
				},
				ZookeeperRef: v1beta1.SolrCloudSpecZookeeperRefArgs{
					Provided: v1beta1.SolrCloudSpecZookeeperRefProvidedArgs{
						Replicas: pulumi.Int(locals.SolrKubernetes.Spec.ZookeeperContainer.Replicas),
						Persistence: v1beta1.SolrCloudSpecZookeeperRefProvidedPersistenceArgs{
							Spec: v1beta1.SolrCloudSpecZookeeperRefProvidedPersistenceSpecArgs{
								Resources: v1beta1.SolrCloudSpecZookeeperRefProvidedPersistenceSpecResourcesArgs{
									Requests: pulumi.Map{
										"storage": pulumi.String(locals.SolrKubernetes.Spec.ZookeeperContainer.DiskSize),
									},
								},
							},
						},
						ZookeeperPodPolicy: v1beta1.SolrCloudSpecZookeeperRefProvidedZookeeperPodPolicyArgs{
							Resources: v1beta1.SolrCloudSpecZookeeperRefProvidedZookeeperPodPolicyResourcesArgs{
								Limits: pulumi.ToMap(map[string]interface{}{
									//"cpu":    locals.SolrKubernetes.Spec.ZookeeperContainer.Resources.Limits.Cpu,
									"memory": locals.SolrKubernetes.Spec.ZookeeperContainer.Resources.Limits.Memory,
								}),
								Requests: pulumi.Map{
									"cpu":    pulumi.String(locals.SolrKubernetes.Spec.ZookeeperContainer.Resources.Requests.Cpu),
									"memory": pulumi.String(locals.SolrKubernetes.Spec.ZookeeperContainer.Resources.Requests.Memory),
								},
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create solr-cloud resource")
	}
	return nil
}
