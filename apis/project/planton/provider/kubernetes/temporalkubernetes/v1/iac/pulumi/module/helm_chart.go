package module

import (
	"github.com/pkg/errors"
	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/temporalkubernetes/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context, locals *Locals,
	createdNamespace pulumi.Resource) error {

	values := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.TemporalKubernetes.Metadata.Name),
	}

	// ------------------------------------------------------------------ database
	if locals.TemporalKubernetes.Spec.Database.ExternalDatabase != nil {
		externalDatabase := locals.TemporalKubernetes.Spec.Database.ExternalDatabase

		// turn off bundled datastores
		values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		// pick sub-driver string
		subDriver := "postgres12"
		if locals.TemporalKubernetes.Spec.Database.Backend == temporalkubernetesv1.TemporalKubernetesDatabaseBackend_mysql {
			subDriver = "mysql8"
		}

		// fall-back names if the manifest omits them
		if locals.TemporalKubernetes.Spec.Database.DatabaseName == "" {
			locals.TemporalKubernetes.Spec.Database.DatabaseName = "temporal"
		}
		if locals.TemporalKubernetes.Spec.Database.VisibilityName == "" {
			locals.TemporalKubernetes.Spec.Database.VisibilityName = "temporal_visibility"
		}

		// build the SQL blocks exactly like upstream values files
		defaultSql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":         pulumi.String(subDriver),
				"host":           pulumi.String(externalDatabase.Host),
				"port":           pulumi.Int(externalDatabase.Port),
				"database":       pulumi.String(locals.TemporalKubernetes.Spec.Database.DatabaseName),
				"user":           pulumi.String(externalDatabase.Username),
				"existingSecret": pulumi.String(vars.DatabasePasswordSecretName),
			},
		}

		visibilitySql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":         pulumi.String(subDriver),
				"host":           pulumi.String(externalDatabase.Host),
				"port":           pulumi.Int(externalDatabase.Port),
				"database":       pulumi.String(locals.TemporalKubernetes.Spec.Database.VisibilityName),
				"user":           pulumi.String(externalDatabase.Username),
				"existingSecret": pulumi.String(vars.DatabasePasswordSecretName),
			},
		}

		values["server"] = pulumi.Map{
			"config": pulumi.Map{
				"persistence": pulumi.Map{
					"default":    defaultSql,
					"visibility": visibilitySql,
					"driver":     pulumi.String("sql"),
				},
			},
		}
	} else {
		// embedded datastore paths (unchanged)
		switch locals.TemporalKubernetes.Spec.Database.Backend {
		case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_cassandra:
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

		case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_mysql:
			values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(true)}
			values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_postgresql:
			values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
			values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(true)}
		}
	}

	// ---------------------------------------------------------- schema jobs
	values["schema"] = pulumi.Map{
		"createDatabase": pulumi.Map{
			"enabled": pulumi.Bool(!locals.TemporalKubernetes.Spec.Database.DisableAutoSchemaSetup),
		},
		"setup":  pulumi.Map{"enabled": pulumi.Bool(false)},
		"update": pulumi.Map{"enabled": pulumi.Bool(false)},
	}

	// -------------------------------------------------------------- web-UI
	if locals.TemporalKubernetes.Spec.DisableWebUi {
		values["web"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ---------------------------------------------------------- monitoring
	monitoring := locals.TemporalKubernetes.Spec.EnableMonitoringStack
	if locals.TemporalKubernetes.Spec.ExternalElasticsearch != nil &&
		locals.TemporalKubernetes.Spec.ExternalElasticsearch.Host != "" {
		monitoring = true
	}
	values["prometheus"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}
	values["grafana"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}
	values["kubePrometheusStack"] = pulumi.Map{"enabled": pulumi.Bool(monitoring)}

	// -------------------------------------------------------- elasticsearch
	es := locals.TemporalKubernetes.Spec.ExternalElasticsearch
	if es != nil && es.Host != "" {
		values["elasticsearch"] = pulumi.Map{
			"enabled":  pulumi.Bool(false),
			"host":     pulumi.String(es.Host),
			"port":     pulumi.Int(es.Port),
			"scheme":   pulumi.String("http"),
			"username": pulumi.String(es.User),
			"password": pulumi.String(es.Password),
		}
	} else if !locals.TemporalKubernetes.Spec.EnableEmbeddedElasticsearch {
		values["elasticsearch"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ------------------------------------------------------- install chart
	_, err := helmv3.NewChart(ctx,
		locals.TemporalKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    values,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create temporal helm chart")
	}

	return nil
}
