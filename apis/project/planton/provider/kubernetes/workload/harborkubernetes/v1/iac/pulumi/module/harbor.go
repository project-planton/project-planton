package module

import (
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/harborkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func harbor(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
	kubernetesProvider pulumi.ProviderResource) error {

	// https://github.com/goharbor/harbor-helm/blob/main/values.yaml
	helmValues := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.HarborKubernetes.Metadata.Name),
		"commonLabels":     pulumi.ToStringMap(locals.KubernetesLabels),
	}

	// Configure Harbor Core container
	if locals.HarborKubernetes.Spec.CoreContainer != nil {
		coreValues := pulumi.Map{
			"replicas": pulumi.Int(int(locals.HarborKubernetes.Spec.CoreContainer.Replicas)),
		}

		if locals.HarborKubernetes.Spec.CoreContainer.Resources != nil {
			coreValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.HarborKubernetes.Spec.CoreContainer.Resources)
		}

		if locals.HarborKubernetes.Spec.CoreContainer.Image != nil {
			coreValues["image"] = pulumi.Map{
				"repository": pulumi.String(locals.HarborKubernetes.Spec.CoreContainer.Image.Repo),
				"tag":        pulumi.String(locals.HarborKubernetes.Spec.CoreContainer.Image.Tag),
			}
		}

		helmValues["core"] = coreValues
	}

	// Configure Harbor Portal container
	if locals.HarborKubernetes.Spec.PortalContainer != nil {
		portalValues := pulumi.Map{
			"replicas": pulumi.Int(int(locals.HarborKubernetes.Spec.PortalContainer.Replicas)),
		}

		if locals.HarborKubernetes.Spec.PortalContainer.Resources != nil {
			portalValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.HarborKubernetes.Spec.PortalContainer.Resources)
		}

		helmValues["portal"] = portalValues
	}

	// Configure Harbor Registry container
	if locals.HarborKubernetes.Spec.RegistryContainer != nil {
		registryValues := pulumi.Map{
			"replicas": pulumi.Int(int(locals.HarborKubernetes.Spec.RegistryContainer.Replicas)),
		}

		if locals.HarborKubernetes.Spec.RegistryContainer.Resources != nil {
			registryValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.HarborKubernetes.Spec.RegistryContainer.Resources)
		}

		helmValues["registry"] = registryValues
	}

	// Configure Harbor Jobservice container
	if locals.HarborKubernetes.Spec.JobserviceContainer != nil {
		jobserviceValues := pulumi.Map{
			"replicas": pulumi.Int(int(locals.HarborKubernetes.Spec.JobserviceContainer.Replicas)),
		}

		if locals.HarborKubernetes.Spec.JobserviceContainer.Resources != nil {
			jobserviceValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.HarborKubernetes.Spec.JobserviceContainer.Resources)
		}

		helmValues["jobservice"] = jobserviceValues
	}

	// Configure database (PostgreSQL)
	if locals.HarborKubernetes.Spec.Database != nil {
		if locals.HarborKubernetes.Spec.Database.IsExternal {
			// External PostgreSQL configuration
			if locals.HarborKubernetes.Spec.Database.ExternalDatabase != nil {
				ext := locals.HarborKubernetes.Spec.Database.ExternalDatabase
				helmValues["database"] = pulumi.Map{
					"type": pulumi.String("external"),
					"external": pulumi.Map{
						"host":                 pulumi.String(ext.Host),
						"port":                 pulumi.String(ext.GetPort()),
						"username":             pulumi.String(ext.Username),
						"password":             pulumi.String(ext.Password),
						"coreDatabase":         pulumi.String(ext.GetCoreDatabase()),
						"clairDatabase":        pulumi.String(ext.GetClairDatabase()),
						"notaryServerDatabase": pulumi.String(ext.GetNotaryServerDatabase()),
						"notarySignerDatabase": pulumi.String(ext.GetNotarySignerDatabase()),
						"sslmode":              pulumi.String(map[bool]string{true: "require", false: "disable"}[ext.UseSsl]),
					},
				}
				// Disable internal PostgreSQL
				helmValues["postgresql"] = pulumi.Map{
					"enabled": pulumi.Bool(false),
				}
			}
		} else {
			// Self-managed PostgreSQL configuration
			if locals.HarborKubernetes.Spec.Database.ManagedDatabase != nil {
				managed := locals.HarborKubernetes.Spec.Database.ManagedDatabase
				postgresValues := pulumi.Map{
					"enabled": pulumi.Bool(true),
				}

				if managed.Container != nil {
					if managed.Container.Resources != nil {
						postgresValues["resources"] = containerresources.ConvertToPulumiMap(
							managed.Container.Resources)
					}

					if managed.Container.PersistenceEnabled {
						postgresValues["persistence"] = pulumi.Map{
							"enabled": pulumi.Bool(true),
							"size":    pulumi.String(managed.Container.DiskSize),
						}
					} else {
						postgresValues["persistence"] = pulumi.Map{
							"enabled": pulumi.Bool(false),
						}
					}
				}

				helmValues["postgresql"] = postgresValues
			}
		}
	}

	// Configure cache (Redis)
	if locals.HarborKubernetes.Spec.Cache != nil {
		if locals.HarborKubernetes.Spec.Cache.IsExternal {
			// External Redis configuration
			if locals.HarborKubernetes.Spec.Cache.ExternalCache != nil {
				ext := locals.HarborKubernetes.Spec.Cache.ExternalCache
				redisConfig := pulumi.Map{
					"type": pulumi.String("external"),
					"external": pulumi.Map{
						"addr":     pulumi.String(ext.Host),
						"password": pulumi.String(ext.Password),
					},
				}

				if ext.UseSentinel {
					redisConfig["external"].(pulumi.Map)["sentinelMasterSet"] = pulumi.String(ext.SentinelMasterSet)
				}

				helmValues["redis"] = redisConfig

				// Disable internal Redis
				helmValues["redis"].(pulumi.Map)["internal"] = pulumi.Map{
					"enabled": pulumi.Bool(false),
				}
			}
		} else {
			// Self-managed Redis configuration
			if locals.HarborKubernetes.Spec.Cache.ManagedCache != nil {
				managed := locals.HarborKubernetes.Spec.Cache.ManagedCache
				redisValues := pulumi.Map{
					"type": pulumi.String("internal"),
					"internal": pulumi.Map{
						"enabled": pulumi.Bool(true),
					},
				}

				if managed.Container != nil {
					if managed.Container.Resources != nil {
						redisValues["internal"].(pulumi.Map)["resources"] = containerresources.ConvertToPulumiMap(
							managed.Container.Resources)
					}
				}

				helmValues["redis"] = redisValues
			}
		}
	}

	// Configure storage
	if locals.HarborKubernetes.Spec.Storage != nil {
		storageConfig := pulumi.Map{}

		switch locals.HarborKubernetes.Spec.Storage.Type {
		case harborkubernetesv1.HarborKubernetesStorageType_s3:
			if locals.HarborKubernetes.Spec.Storage.S3 != nil {
				s3 := locals.HarborKubernetes.Spec.Storage.S3
				storageConfig = pulumi.Map{
					"type": pulumi.String("s3"),
					"s3": pulumi.Map{
						"bucket":         pulumi.String(s3.Bucket),
						"region":         pulumi.String(s3.Region),
						"accesskey":      pulumi.String(s3.AccessKey),
						"secretkey":      pulumi.String(s3.SecretKey),
						"regionendpoint": pulumi.String(s3.EndpointUrl),
						"encrypt":        pulumi.Bool(s3.Encrypt),
						"secure":         pulumi.Bool(s3.Secure),
					},
				}
			}
		case harborkubernetesv1.HarborKubernetesStorageType_gcs:
			if locals.HarborKubernetes.Spec.Storage.Gcs != nil {
				gcs := locals.HarborKubernetes.Spec.Storage.Gcs
				storageConfig = pulumi.Map{
					"type": pulumi.String("gcs"),
					"gcs": pulumi.Map{
						"bucket":    pulumi.String(gcs.Bucket),
						"keydata":   pulumi.String(gcs.KeyData),
						"chunksize": pulumi.Int(int(gcs.GetChunkSize())),
					},
				}
			}
		case harborkubernetesv1.HarborKubernetesStorageType_azure:
			if locals.HarborKubernetes.Spec.Storage.Azure != nil {
				azure := locals.HarborKubernetes.Spec.Storage.Azure
				storageConfig = pulumi.Map{
					"type": pulumi.String("azure"),
					"azure": pulumi.Map{
						"accountname": pulumi.String(azure.AccountName),
						"accountkey":  pulumi.String(azure.AccountKey),
						"container":   pulumi.String(azure.Container),
					},
				}
			}
		case harborkubernetesv1.HarborKubernetesStorageType_oss:
			if locals.HarborKubernetes.Spec.Storage.Oss != nil {
				oss := locals.HarborKubernetes.Spec.Storage.Oss
				storageConfig = pulumi.Map{
					"type": pulumi.String("oss"),
					"oss": pulumi.Map{
						"bucket":          pulumi.String(oss.Bucket),
						"endpoint":        pulumi.String(oss.Endpoint),
						"accesskeyid":     pulumi.String(oss.AccessKeyId),
						"accesskeysecret": pulumi.String(oss.AccessKeySecret),
						"secure":          pulumi.Bool(oss.Secure),
					},
				}
			}
		case harborkubernetesv1.HarborKubernetesStorageType_filesystem:
			if locals.HarborKubernetes.Spec.Storage.Filesystem != nil {
				fs := locals.HarborKubernetes.Spec.Storage.Filesystem
				storageConfig = pulumi.Map{
					"type": pulumi.String("filesystem"),
					"filesystem": pulumi.Map{
						"size": pulumi.String(fs.DiskSize),
					},
				}
			}
		}

		helmValues["persistence"] = pulumi.Map{
			"imageChartStorage": storageConfig,
		}
	}

	// Disable ingress in Helm chart (we manage it separately using Gateway API)
	helmValues["expose"] = pulumi.Map{
		"type": pulumi.String("clusterIP"),
	}

	// Merge with custom helm values if provided
	if locals.HarborKubernetes.Spec.HelmValues != nil {
		for k, v := range locals.HarborKubernetes.Spec.HelmValues {
			helmValues[k] = pulumi.String(v)
		}
	}

	// Deploy Harbor using Helm chart
	_, err := helmv3.NewRelease(ctx,
		locals.HarborKubernetes.Metadata.Name,
		&helmv3.ReleaseArgs{
			Name:      pulumi.String(locals.HarborKubernetes.Metadata.Name),
			Namespace: createdNamespace.Metadata.Name(),
			Chart:     pulumi.String("harbor"),
			RepositoryOpts: helmv3.RepositoryOptsArgs{
				Repo: pulumi.String("https://helm.goharbor.io"),
			},
			Values: helmValues,
		},
		pulumi.Provider(kubernetesProvider),
		pulumi.Parent(createdNamespace),
	)

	return err
}
