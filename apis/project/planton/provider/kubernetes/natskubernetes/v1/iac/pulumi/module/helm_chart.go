package module

import (
	"github.com/pkg/errors"
	natskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/natskubernetes/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs the official NATS Helm chart (v1.3.6) with values
// translated 1-to-1 from NatsKubernetesSpec.
//
// • server_container ➜ replicas, resources, JetStream PVC size
// • disable_jet_stream flag ➜ jetstream.enabled=false
// • auth.enabled ➜ token|basic secret references
// • tls_enabled ➜ reference to tls-<ns> Secret
// • disable_nats_box ➜ natsbox.enabled=false
func helmChart(ctx *pulumi.Context, locals *Locals,
	createdNamespace pulumi.Resource) error {

	values := pulumi.Map{}

	// ---------------------------------------------------------------- cluster
	sc := locals.NatsKubernetes.Spec.ServerContainer
	cluster := pulumi.Map{
		"size": pulumi.Int(sc.Replicas),
		"resources": pulumi.Map{
			"limits": pulumi.Map{
				"cpu":    pulumi.String(sc.Resources.Limits.Cpu),
				"memory": pulumi.String(sc.Resources.Limits.Memory),
			},
			"requests": pulumi.Map{
				"cpu":    pulumi.String(sc.Resources.Requests.Cpu),
				"memory": pulumi.String(sc.Resources.Requests.Memory),
			},
		},
	}
	values["cluster"] = cluster

	// ------------------------------------------------------------ jetstream
	if locals.NatsKubernetes.Spec.DisableJetStream {
		values["jetstream"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	} else {
		values["jetstream"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
			"fileStorage": pulumi.Map{
				"size": pulumi.String(sc.DiskSize),
			},
		}
	}

	// ---------------------------------------------------------------- auth
	auth := locals.NatsKubernetes.Spec.Auth
	if auth.Enabled {
		switch auth.Scheme {
		case natskubernetesv1.NatsKubernetesAuthScheme_bearer_token:
			values["auth"] = pulumi.Map{
				"enabled": pulumi.Bool(true),
				"token": pulumi.Map{
					"secret": pulumi.Map{
						"name": pulumi.String(locals.AuthSecretName),
						"key":  pulumi.String(vars.AuthSecretKey),
					},
				},
			}

		case natskubernetesv1.NatsKubernetesAuthScheme_basic_auth:
			values["auth"] = pulumi.Map{
				"enabled": pulumi.Bool(true),
				"basic": pulumi.Map{
					"secret": pulumi.Map{
						"name":        pulumi.String(locals.AuthSecretName),
						"userKey":     pulumi.String(vars.AuthSecretKeyUser),
						"passwordKey": pulumi.String(vars.AuthSecretKeyPassword),
					},
				},
			}

		default:
			// leave auth disabled
		}
	}

	// ---------------------------------------------------------------- tls
	if locals.NatsKubernetes.Spec.TlsEnabled {
		values["tls"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
			// chart expects secret.name
			"secret": pulumi.Map{
				"name": pulumi.String(locals.TlsSecretName),
			},
		}
	}

	// ----------------------------------------------------------- nats-box pod
	if locals.NatsKubernetes.Spec.DisableNatsBox {
		values["natsbox"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// --------------------------------------------------------- install chart
	_, err := helmv3.NewChart(ctx,
		locals.NatsKubernetes.Metadata.Name, // release name
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
		return errors.Wrap(err, "failed to deploy NATS Helm chart")
	}

	// ------------------------------------------- metrics endpoint export
	metricsEndpoint := pulumi.Sprintf("http://nats-prom.%s.svc.cluster.local:7777/metrics",
		locals.Namespace)
	ctx.Export(OpMetricsEndpoint, metricsEndpoint)

	return nil
}
