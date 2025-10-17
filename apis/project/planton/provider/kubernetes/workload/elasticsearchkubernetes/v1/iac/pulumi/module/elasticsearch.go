package module

import (
	"github.com/pkg/errors"
	elasticsearchv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/elasticsearch/kubernetes/elasticsearch/v1"
	kibanav1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/elasticsearch/kubernetes/kibana/v1beta1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func elasticsearch(ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	var volumeClaimTemplates = &elasticsearchv1.ElasticsearchSpecNodeSetsVolumeClaimTemplatesArray{}
	if locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.PersistenceEnabled {
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
							"storage": pulumi.String(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.DiskSize),
						},
					},
				},
			},
		}
	}

	createdElasticSearch, err := elasticsearchv1.NewElasticsearch(ctx, locals.ElasticsearchKubernetes.Metadata.Name, &elasticsearchv1.ElasticsearchArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String(locals.ElasticsearchKubernetes.Metadata.Name),
			Namespace: createdNamespace.Metadata.Name(),
			Labels:    pulumi.ToStringMap(locals.Labels),
			Annotations: pulumi.StringMap{
				"pulumi.com/patchForce": pulumi.String("true"),
			},
		},
		Spec: &elasticsearchv1.ElasticsearchSpecArgs{
			NodeSets: elasticsearchv1.ElasticsearchSpecNodeSetsArray{
				elasticsearchv1.ElasticsearchSpecNodeSetsArgs{
					Name:  pulumi.String("elasticsearch"),
					Count: pulumi.Int(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.Replicas),
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
											"memory": pulumi.String(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.Resources.Requests.Memory),
											"cpu":    pulumi.String(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.Resources.Requests.Cpu),
										},
										"limits": pulumi.Map{
											"memory": pulumi.String(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.Resources.Limits.Memory),
											"cpu":    pulumi.String(locals.ElasticsearchKubernetes.Spec.Elasticsearch.Container.Resources.Limits.Cpu),
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
	}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrapf(err, "failed to create elastic search")
	}

	if locals.ElasticsearchKubernetes.Spec.Kibana.Enabled {
		_, err = kibanav1.NewKibana(ctx, locals.ElasticsearchKubernetes.Metadata.Name, &kibanav1.KibanaArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.ElasticsearchKubernetes.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
				Annotations: pulumi.StringMap{
					"pulumi.com/patchForce": pulumi.String("true"),
				},
			},
			Spec: &kibanav1.KibanaSpecArgs{
				Version: pulumi.String(vars.ElasticsearchVersion),
				Count:   pulumi.Int(locals.ElasticsearchKubernetes.Spec.Kibana.Container.Replicas),
				PodTemplate: pulumi.Map{
					"spec": pulumi.Map{
						"containers": pulumi.Array{
							pulumi.Map{
								"name": pulumi.String("kibana"),
								"resources": pulumi.Map{
									"requests": pulumi.Map{
										"memory": pulumi.String(locals.ElasticsearchKubernetes.Spec.Kibana.Container.Resources.Requests.Memory),
										"cpu":    pulumi.String(locals.ElasticsearchKubernetes.Spec.Kibana.Container.Resources.Requests.Cpu),
									},
									"limits": pulumi.Map{
										"memory": pulumi.String(locals.ElasticsearchKubernetes.Spec.Kibana.Container.Resources.Limits.Memory),
										"cpu":    pulumi.String(locals.ElasticsearchKubernetes.Spec.Kibana.Container.Resources.Limits.Cpu),
									},
								},
							},
						},
					},
				},
				ElasticsearchRef: kibanav1.KibanaSpecElasticsearchRefArgs{
					Name:      createdElasticSearch.Metadata.Name().Elem(),
					Namespace: createdNamespace.Metadata.Name(),
				},
				Http: kibanav1.KibanaSpecHttpArgs{
					Tls: kibanav1.KibanaSpecHttpTlsArgs{
						SelfSignedCertificate: kibanav1.KibanaSpecHttpTlsSelfSignedCertificateArgs{
							Disabled: pulumi.Bool(true),
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace), pulumi.DependsOn([]pulumi.Resource{createdElasticSearch}))
		if err != nil {
			return errors.Wrapf(err, "failed to create kibana instance for the elastic search instance")
		}
	}

	return nil
}
