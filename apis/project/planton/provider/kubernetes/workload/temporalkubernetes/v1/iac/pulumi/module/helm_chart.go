package module

import (
	"github.com/pkg/errors"
	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/temporalkubernetes/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context, locals *Locals,
	createdNamespace pulumi.Resource) error {

	values := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.TemporalKubernetes.Metadata.Name),
	}

	// ---------------------------------------------------------------- database
	db := locals.TemporalKubernetes.Spec.Database

	if db.ExternalDatabase != nil {
		ext := db.ExternalDatabase

		// disable embedded datastores
		values["cassandra"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["mysql"] = pulumi.Map{"enabled": pulumi.Bool(false)}
		values["postgresql"] = pulumi.Map{"enabled": pulumi.Bool(false)}

		// choose sub-driver
		subDriver := "postgres12"
		if db.Backend == temporalkubernetesv1.TemporalKubernetesDatabaseBackend_mysql {
			subDriver = "mysql8"
		}

		// default DB names if not provided
		defaultDB := db.DatabaseName
		visibilityDB := db.VisibilityName
		if defaultDB == "" {
			defaultDB = "temporal"
		}
		if visibilityDB == "" {
			visibilityDB = "temporal_visibility"
		}

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

		defaultSql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":         pulumi.String(subDriver),
				"host":           pulumi.String(ext.Host),
				"port":           pulumi.Int(ext.Port),
				"database":       pulumi.String(defaultDB),
				"user":           pulumi.String(user),
				"existingSecret": pulumi.String(vars.DatabasePasswordSecretName),
				"tls":            tls,
			},
		}

		visibilitySql := pulumi.Map{
			"driver": pulumi.String("sql"),
			"sql": pulumi.Map{
				"driver":         pulumi.String(subDriver),
				"host":           pulumi.String(ext.Host),
				"port":           pulumi.Int(ext.Port),
				"database":       pulumi.String(visibilityDB),
				"user":           pulumi.String(user),
				"existingSecret": pulumi.String(vars.DatabasePasswordSecretName),
				"tls":            tls,
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
		// embedded datastore paths
		switch db.Backend {
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

	// ------------------------------------------------------ schema jobs
	values["schema"] = pulumi.Map{
		"createDatabase": pulumi.Map{
			"enabled": pulumi.Bool(!db.DisableAutoSchemaSetup), // true by default
		},
		"setup":  pulumi.Map{"enabled": pulumi.Bool(true)},
		"update": pulumi.Map{"enabled": pulumi.Bool(true)},
	}

	// -------------------------------------------------------------- web-UI
	if locals.TemporalKubernetes.Spec.DisableWebUi {
		values["web"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// ---------------------------------------------------- monitoring stack
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

	// ------------------------------------------------- search attributes
	if len(locals.TemporalKubernetes.Spec.SearchAttributes) > 0 {
		searchAttrsMap := pulumi.Map{}
		for _, attr := range locals.TemporalKubernetes.Spec.SearchAttributes {
			typeName := mapSearchAttributeType(attr.Type)
			searchAttrsMap[attr.Name] = pulumi.String(typeName)
		}

		// Configure via server dynamic config
		if serverCfg, ok := values["server"].(pulumi.Map); ok {
			if configMap, ok := serverCfg["config"].(pulumi.Map); ok {
				configMap["customSearchAttributes"] = searchAttrsMap
			} else {
				serverCfg["config"] = pulumi.Map{
					"customSearchAttributes": searchAttrsMap,
				}
			}
		} else {
			values["server"] = pulumi.Map{
				"config": pulumi.Map{
					"customSearchAttributes": searchAttrsMap,
				},
			}
		}
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

// mapSearchAttributeType converts proto enum to Temporal search attribute type string
func mapSearchAttributeType(attrType temporalkubernetesv1.TemporalKubernetesSearchAttributeType) string {
	switch attrType {
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_keyword_type:
		return "Keyword"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_text_type:
		return "Text"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_int_type:
		return "Int"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_double_type:
		return "Double"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_bool_type:
		return "Bool"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_datetime_type:
		return "Datetime"
	case temporalkubernetesv1.TemporalKubernetesSearchAttributeType_keyword_list_type:
		return "KeywordList"
	default:
		return "Keyword"
	}
}
