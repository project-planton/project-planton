package module

import (
	"github.com/pkg/errors"
	elasticsearchv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/elasticsearch/kubernetes/elasticsearch/v1"
	kibanav1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/elasticsearch/kubernetes/kibana/v1beta1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func elasticsearch(ctx *pulumi.Context,
	locals *Locals,
	namespaceInput pulumi.StringInput,
	createdNamespace *kubernetescorev1.Namespace,
	kubernetesProvider *kubernetes.Provider) error {

	var volumeClaimTemplates = &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesArray{}
	if locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.PersistenceEnabled {
		volumeClaimTemplates = &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesArray{
			elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesArgs{
				Metadata: &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesMetadataArgs{
					Name: pulumi.String("elasticsearch-data"),
				},
				Spec: &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesSpecArgs{
					AccessModes: pulumi.StringArray{
						pulumi.String("ReadWriteOnce"),
					},
					Resources: &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesSpecResourcesArgs{
						Requests: pulumi.Map{
							"storage": pulumi.String(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.DiskSize),
						},
					},
				},
			},
		}
	}

	// Build resource options based on whether namespace was created
	esResourceOpts := []pulumi.ResourceOption{pulumi.Provider(kubernetesProvider)}
	if createdNamespace != nil {
		esResourceOpts = append(esResourceOpts, pulumi.Parent(createdNamespace))
	}

	createdElasticSearch, err := elasticsearchv1.NewElasticsearch(ctx, locals.KubernetesElasticsearch.Metadata.Name, &elasticsearchv1.ElasticsearchArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String(locals.KubernetesElasticsearch.Metadata.Name),
			Namespace: namespaceInput,
			Labels:    pulumi.ToStringMap(locals.Labels),
			Annotations: pulumi.StringMap{
				"pulumi.com/patchForce": pulumi.String("true"),
			},
		},
		Spec: &elasticsearchv1.ElasticsearchSpecArgs{
			NodeSets: elasticsearchv1.ElasticsearchSpecNodeSetsArray{
				elasticsearchv1.ElasticsearchSpecNodeSetsArgs{
					Name:  pulumi.String("elasticsearch"),
					Count: pulumi.Int(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.Replicas),
					Config: pulumi.Map{
						"node.roles":            pulumi.ToStringArray([]string{"master", "data", "ingest"}),
						"node.store.allow_mmap": pulumi.Bool(false),
					},
					PodTemplate: pulumi.Map{
						"metadata": pulumi.Map{
							"labels": pulumi.StringMap{
								"role": pulumi.String("master"),
							},
						},
						"spec": pulumi.Map{
							"containers": pulumi.Array{
								pulumi.Map{
									"name": pulumi.String("elasticsearch"),
									"resources": pulumi.Map{
										"requests": pulumi.Map{
											"memory": pulumi.String(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.Resources.Requests.Memory),
											"cpu":    pulumi.String(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.Resources.Requests.Cpu),
										},
										"limits": pulumi.Map{
											"memory": pulumi.String(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.Resources.Limits.Memory),
											"cpu":    pulumi.String(locals.KubernetesElasticsearch.Spec.Elasticsearch.Container.Resources.Limits.Cpu),
										},
									},
								},
							},
						},
					},
					VolumeClaimTemplates: volumeClaimTemplates,
				},
			},
			Version: pulumi.String(vars.ElasticsearchVersion),
			Http: &elasticsearchv1.ElasticsearchSpecHttpArgs{
				Tls: &elasticsearchv1.ElasticsearchSpecHttpTlsArgs{
					SelfSignedCertificate: &elasticsearchv1.ElasticsearchSpecHttpTlsSelfSignedCertificateArgs{
						Disabled: pulumi.Bool(true),
					},
				},
			},
		},
	}, esResourceOpts...)
	if err != nil {
		return errors.Wrapf(err, "failed to create elastic search")
	}

	if locals.KubernetesElasticsearch.Spec.Kibana.Enabled {
		// Build resource options for Kibana
		kibanaResourceOpts := []pulumi.ResourceOption{
			pulumi.Provider(kubernetesProvider),
			pulumi.DependsOn([]pulumi.Resource{createdElasticSearch}),
		}
		if createdNamespace != nil {
			kibanaResourceOpts = append(kibanaResourceOpts, pulumi.Parent(createdNamespace))
		}

		_, err = kibanav1.NewKibana(ctx, locals.KubernetesElasticsearch.Metadata.Name, &kibanav1.KibanaArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.KubernetesElasticsearch.Metadata.Name),
				Namespace: namespaceInput,
				Labels:    pulumi.ToStringMap(locals.Labels),
				Annotations: pulumi.StringMap{
					"pulumi.com/patchForce": pulumi.String("true"),
				},
			},
			Spec: &kibanav1.KibanaSpecArgs{
				Version: pulumi.String(vars.ElasticsearchVersion),
				Count:   pulumi.Int(locals.KubernetesElasticsearch.Spec.Kibana.Container.Replicas),
				PodTemplate: pulumi.Map{
					"spec": pulumi.Map{
						"containers": pulumi.Array{
							pulumi.Map{
								"name": pulumi.String("kibana"),
								"resources": pulumi.Map{
									"requests": pulumi.Map{
										"memory": pulumi.String(locals.KubernetesElasticsearch.Spec.Kibana.Container.Resources.Requests.Memory),
										"cpu":    pulumi.String(locals.KubernetesElasticsearch.Spec.Kibana.Container.Resources.Requests.Cpu),
									},
									"limits": pulumi.Map{
										"memory": pulumi.String(locals.KubernetesElasticsearch.Spec.Kibana.Container.Resources.Limits.Memory),
										"cpu":    pulumi.String(locals.KubernetesElasticsearch.Spec.Kibana.Container.Resources.Limits.Cpu),
									},
								},
							},
						},
					},
				},
				ElasticsearchRef: kibanav1.KibanaSpecElasticsearchRefArgs{
					Name:      createdElasticSearch.Metadata.Name().Elem(),
					Namespace: namespaceInput,
				},
				Http: kibanav1.KibanaSpecHttpArgs{
					Tls: kibanav1.KibanaSpecHttpTlsArgs{
						SelfSignedCertificate: kibanav1.KibanaSpecHttpTlsSelfSignedCertificateArgs{
							Disabled: pulumi.Bool(true),
						},
					},
				},
			},
		}, kibanaResourceOpts...)
		if err != nil {
			return errors.Wrapf(err, "failed to create kibana instance for the elastic search instance")
		}
	}

	return nil
}
