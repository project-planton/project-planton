package module

import (
	"github.com/pkg/errors"
	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/temporalkubernetes/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs Temporal via the official Helm chart and wires-in
// only the minimal values derived from the API spec. Everything else is
// left to the chart defaults so Terraform-minded users can inspect / copy
// the “values.yaml” semantics.
func helmChart(ctx *pulumi.Context, locals *Locals,
	createdNamespace pulumi.Resource) error {

	values := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.TemporalKubernetes.Metadata.Name),
	}

	// ---------------------------------------------------------------------
	// persistence
	// ---------------------------------------------------------------------
	db := locals.TemporalKubernetes.Spec.Database
	if db.ExternalDatabase != nil {
		ext := db.ExternalDatabase

		// disable bundled datastores
		values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		sqlBlock := pulumi.Map{
			"driver":   pulumi.String(sqlDriver(db.Backend)),
			"host":     pulumi.String(ext.Host),
			"port":     pulumi.Int(ext.Port),
			"database": pulumi.String(db.DatabaseName),
			"user":     pulumi.String(ext.User),
			"password": pulumi.String(ext.Password),
		}
		sqlVis := pulumi.Map{
			"driver":   pulumi.String(sqlDriver(db.Backend)),
			"host":     pulumi.String(ext.Host),
			"port":     pulumi.Int(ext.Port),
			"database": pulumi.String(db.VisibilityName),
			"user":     pulumi.String(ext.User),
			"password": pulumi.String(ext.Password),
		}

		values["server"] = pulumi.Map{
			"config": pulumi.Map{
				"persistence": pulumi.Map{
					"default":    sqlBlock,
					"visibility": sqlVis,
					"driver":     pulumi.String("sql"),
				},
			},
		}
	} else {
		switch db.Backend {
		case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_cassandra:
			replicas := int(locals.TemporalKubernetes.Spec.CassandraReplicas)
			if replicas == 0 {
				replicas = vars.DefaultCassandraReplicas
			}

			if replicas == 1 {
				// dev=true → single-pod StatefulSet
				values["cassandra"] = pulumi.Map{
					"enabled": pulumi.Bool(true),
					"configuration": pulumi.Map{
						"dev": pulumi.Bool(true),
					},
				}
			} else {
				// production-style multi-node cluster
				values["cassandra"] = pulumi.Map{
					"enabled": pulumi.Bool(true),
					"configuration": pulumi.Map{
						"dev":          pulumi.Bool(false),
						"cluster_size": pulumi.Int(replicas),
					},
				}
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

	// ---------------------------------------------------------------------
	// schema setup (keep setup/update jobs so CQL files are mounted)
	// ---------------------------------------------------------------------
	values["schema"] = pulumi.Map{
		"createDatabase": pulumi.Map{
			"enabled": pulumi.Bool(!db.DisableAutoSchemaSetup),
		},
		"setup":  pulumi.Map{"enabled": pulumi.Bool(true)},
		"update": pulumi.Map{"enabled": pulumi.Bool(true)},
	}

	// ---------------------------------------------------------------------
	// web-ui
	// ---------------------------------------------------------------------
	if locals.TemporalKubernetes.Spec.DisableWebUi {
		values["web"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ---------------------------------------------------------------------
	// monitoring stack (Prometheus + Grafana)
	// ---------------------------------------------------------------------
	monitoringEnabled := locals.TemporalKubernetes.Spec.EnableMonitoringStack
	if locals.TemporalKubernetes.Spec.ExternalElasticsearch != nil &&
		locals.TemporalKubernetes.Spec.ExternalElasticsearch.Host != "" {
		// force on when external ES is provided
		monitoringEnabled = true
	}
	values["prometheus"] = pulumi.Map{"enabled": pulumi.Bool(monitoringEnabled)}
	values["grafana"] = pulumi.Map{"enabled": pulumi.Bool(monitoringEnabled)}
	values["kubePrometheusStack"] = pulumi.Map{"enabled": pulumi.Bool(monitoringEnabled)}

	// ---------------------------------------------------------------------
	// elasticsearch – embedded vs external
	// ---------------------------------------------------------------------
	es := locals.TemporalKubernetes.Spec.ExternalElasticsearch
	embedES := locals.TemporalKubernetes.Spec.EnableEmbeddedElasticsearch
	switch {
	case es != nil && es.Host != "":
		values["elasticsearch"] = pulumi.Map{
			"enabled":  pulumi.Bool(false),
			"host":     pulumi.String(es.Host),
			"port":     pulumi.Int(es.Port),
			"scheme":   pulumi.String("http"),
			"username": pulumi.String(es.User),
			"password": pulumi.String(es.Password),
		}
	case !embedES:
		values["elasticsearch"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ---------------------------------------------------------------------
	// install chart
	// ---------------------------------------------------------------------
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

// sqlDriver maps proto enum → Temporal Helm sql driver string.
func sqlDriver(backend temporalkubernetesv1.TemporalKubernetesDatabaseBackend) string {
	switch backend {
	case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_mysql:
		return "mysql8"
	case temporalkubernetesv1.TemporalKubernetesDatabaseBackend_postgresql:
		return "postgres12"
	default:
		return "unknown"
	}
}
