package module

import (
	"github.com/pkg/errors"
	kubernetestemporalv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetestemporal/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider pulumi.ProviderResource) error {

	values := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.KubernetesTemporal.Metadata.Name),
	}

	// ---------------------------------------------------------------- database
	db := locals.KubernetesTemporal.Spec.Database

	if db.ExternalDatabase != nil {
		ext := db.ExternalDatabase

		// disable embedded datastores
		values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		// choose sub-driver
		subDriver := "postgres12"
		if db.Backend == kubernetestemporalv1.KubernetesTemporalDatabaseBackend_mysql {
			subDriver = "mysql8"
		}

		// default DB names if not provided
		defaultDB := db.GetDatabaseName()
		visibilityDB := db.GetVisibilityName()

		// username field
		user := ext.Username
		if user == "" {
			user = ext.Username
		}

		// common TLS section (SSL on, host-verification off)
		tls := pulumi.Map{
			"enabled":                pulumi.Bool(true),
			"enableHostVerification": pulumi.Bool(false),
		}

		// Determine which secret to use for database password
		// If secretRef is provided, use the existing secret; otherwise, use the secret we created
		dbSecretName := locals.DatabasePasswordSecretName
		dbSecretKey := vars.DatabasePasswordSecretKey
		if ext.Password != nil && ext.Password.GetSecretRef() != nil {
			secretRef := ext.Password.GetSecretRef()
			dbSecretName = secretRef.Name
			dbSecretKey = secretRef.Key
		}

		defaultSql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":            pulumi.String(subDriver),
				"host":              pulumi.String(ext.Host),
				"port":              pulumi.Int(ext.Port),
				"database":          pulumi.String(defaultDB),
				"user":              pulumi.String(user),
				"existingSecret":    pulumi.String(dbSecretName),
				"existingSecretKey": pulumi.String(dbSecretKey),
				"tls":               tls,
			},
		}

		visibilitySql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":            pulumi.String(subDriver),
				"host":              pulumi.String(ext.Host),
				"port":              pulumi.Int(ext.Port),
				"database":          pulumi.String(visibilityDB),
				"user":              pulumi.String(user),
				"existingSecret":    pulumi.String(dbSecretName),
				"existingSecretKey": pulumi.String(dbSecretKey),
				"tls":               tls,
			},
		}

		values["server"] = pulumi.Map{
			"config": pulumi.Map{
				"services": pulumi.Map{
					"frontend": pulumi.Map{
						"rpc": pulumi.Map{
							"grpcPort": pulumi.Int(vars.FrontendGrpcPort),
							"httpPort": pulumi.Int(vars.FrontendHttpPort),
						},
					},
				},
				"persistence": pulumi.Map{
					"default":    defaultSql,
					"visibility": visibilitySql,
					"driver":     pulumi.String("sql"),
				},
			},
		}
	} else {
		// embedded datastore paths
		switch db.Backend {
		case kubernetestemporalv1.KubernetesTemporalDatabaseBackend_cassandra:
			values["cassandra"] = pulumi.Map{
				"enabled":      pulumi.Bool(true),
				"replicaCount": pulumi.Int(1),
				"config": pulumi.Map{
					"dev":          pulumi.Bool(true),
					"cluster_size": pulumi.Int(1),
				},
			}
			values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		case kubernetestemporalv1.KubernetesTemporalDatabaseBackend_mysql:
			values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(true)}
			values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		case kubernetestemporalv1.KubernetesTemporalDatabaseBackend_postgresql:
			values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(true)}
		}
	}

	// ------------------------------------------------------ schema jobs
	values["schema"] = pulumi.Map{
		"createDatabase": pulumi.Map{
			"enabled": pulumi.Bool(!db.DisableAutoSchemaSetup), // true by default
		},
		"setup":  pulumi.Map{"enabled": pulumi.Bool(true)},
		"update": pulumi.Map{"enabled": pulumi.Bool(true)},
	}

	// -------------------------------------------------------------- web-UI
	if locals.KubernetesTemporal.Spec.DisableWebUi {
		values["web"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ---------------------------------------------------- monitoring stack
	monitoring := locals.KubernetesTemporal.Spec.EnableMonitoringStack
	if locals.KubernetesTemporal.Spec.ExternalElasticsearch != nil &&
		locals.KubernetesTemporal.Spec.ExternalElasticsearch.Host != "" {
		monitoring = true
	}
	values["prometheus"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}
	values["grafana"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}
	values["kubePrometheusStack"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}

	// -------------------------------------------------------- elasticsearch
	es := locals.KubernetesTemporal.Spec.ExternalElasticsearch
	if es != nil && es.Host != "" {
		esValues := pulumi.Map{
			"enabled":  pulumi.Bool(false),
			"host":     pulumi.String(es.Host),
			"port":     pulumi.Int(es.Port),
			"scheme":   pulumi.String("http"),
			"username": pulumi.String(es.User),
		}

		// Handle password - either as plain string or from existing secret
		if es.Password != nil {
			if es.Password.GetSecretRef() != nil {
				// Use existing Kubernetes secret for password
				secretRef := es.Password.GetSecretRef()
				esValues["existingSecret"] = pulumi.String(secretRef.Name)
				esValues["existingSecretKey"] = pulumi.String(secretRef.Key)
			} else if es.Password.GetStringValue() != "" {
				// Use plain string password
				esValues["password"] = pulumi.String(es.Password.GetStringValue())
			}
		}

		values["elasticsearch"] = esValues
	} else if !locals.KubernetesTemporal.Spec.EnableEmbeddedElasticsearch {
		values["elasticsearch"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ----------------------------------------------------------- version
	// determine which version to use: spec.version if provided, otherwise default
	chartVersion := vars.HelmChartVersion
	if locals.KubernetesTemporal.Spec.Version != "" {
		chartVersion = locals.KubernetesTemporal.Spec.Version
	}

	// ------------------------------------------------------- install chart
	_, err := helmv3.NewChart(ctx,
		locals.KubernetesTemporal.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(chartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    values,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create temporal helm chart")
	}

	return nil
}
