package module

import (
	"github.com/pkg/errors"
	natskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/natskubernetes/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs the upstream NATS Helm chart and tailors it to the
// NatsKubernetesSpec supplied by the API-resource.
func helmChart(ctx *pulumi.Context, locals *Locals, parent pulumi.Resource) error {
	sc := locals.NatsKubernetes.Spec.ServerContainer
	if sc == nil {
		return errors.New("spec.serverContainer must be set")
	}
	if sc.Resources == nil || sc.Resources.Requests == nil || sc.Resources.Limits == nil {
		return errors.New("serverContainer.resources.{requests,limits} must be set")
	}

	// ---------------------------------------------------------------------
	// Helm values
	// ---------------------------------------------------------------------
	values := pulumi.Map{}
	config := pulumi.Map{}

	// ----------------------- container resources -------------------------
	values["container"] = pulumi.Map{
		"merge": pulumi.Map{
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
		},
	}

	// ---------------------------- clustering -----------------------------
	clusterEnabled := sc.Replicas > 1
	config["cluster"] = pulumi.Map{
		"enabled":  pulumi.Bool(clusterEnabled),
		"replicas": pulumi.Int(sc.Replicas),
	}

	// ---------------------------- JetStream ------------------------------
	if locals.NatsKubernetes.Spec.DisableJetStream {
		config["jetstream"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	} else {
		if sc.DiskSize == "" {
			return errors.New("serverContainer.diskSize must be set when JetStream is enabled")
		}
		config["jetstream"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
			"fileStore": pulumi.Map{
				"enabled": pulumi.Bool(true),
				"pvc": pulumi.Map{
					"size": pulumi.String(sc.DiskSize),
				},
			},
		}
	}

	// --------------------------- authentication --------------------------
	if auth := locals.NatsKubernetes.Spec.Auth; auth != nil && auth.Enabled {
		switch auth.Scheme {
		case natskubernetesv1.NatsKubernetesAuthScheme_bearer_token:
			values["auth"] = pulumi.Map{
				"enabled": pulumi.Bool(true),
				"token": pulumi.Map{
					"users": pulumi.Array{
						pulumi.Map{
							"existingSecret": pulumi.Map{
								"name": pulumi.String(vars.AdminAuthSecretName),
								"key":  pulumi.String(vars.AdminAuthSecretKey),
							},
						},
					},
				},
			}
		case natskubernetesv1.NatsKubernetesAuthScheme_basic_auth:
			// ----- 1. admin user via main secret -----
			userList := pulumi.Array{
				pulumi.Map{
					"existingSecret": pulumi.Map{
						"name":        pulumi.String(vars.AdminAuthSecretName),
						"userKey":     pulumi.String(vars.NatsUserSecretKeyUsername),
						"passwordKey": pulumi.String(vars.NatsUserSecretKeyPassword),
					},
				},
			}

			// ----- 2. optional noauth user via **second** secret -----
			if na := auth.NoAuthUser; na != nil && na.Enabled {
				userList = append(userList, pulumi.Map{
					"existingSecret": pulumi.Map{
						"name":        pulumi.String(vars.NoAuthUserSecretName), // auth-noauth-<ns>
						"userKey":     pulumi.String(vars.NatsUserSecretKeyUsername),
						"passwordKey": pulumi.String(vars.NatsUserSecretKeyPassword),
					},
					"permissions": pulumi.Map{
						"publish":   pulumi.ToStringArray(na.PublishSubjects),
						"subscribe": pulumi.Array{},
					},
				})

				// top-level pointer because users come from the include file
				config["patch"] = pulumi.Array{
					pulumi.Map{
						"op":    pulumi.String("add"),
						"path":  pulumi.String("/no_auth_user"),
						"value": pulumi.String(vars.NoAuthUsername),
					},
				}
			}

			// hand off users array to Helm chart
			values["auth"] = pulumi.Map{
				"enabled": pulumi.Bool(true),
				"basic":   pulumi.Map{"users": userList},
			}
		}
	}

	// ------------------------------- TLS ---------------------------------
	if locals.NatsKubernetes.Spec.TlsEnabled {
		values["tls"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
			"secret":  pulumi.Map{"name": pulumi.String(locals.TlsSecretName)},
		}
	}

	// --------------------------- nats-box --------------------------------
	if locals.NatsKubernetes.Spec.DisableNatsBox {
		values["natsbox"] = pulumi.Map{"enabled": pulumi.Bool(false)}
	}

	// Attach chart-level config and deploy.
	values["config"] = config

	_, err := helmv3.NewChart(ctx,
		locals.NatsKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    values,
			FetchArgs: helmv3.FetchArgs{Repo: pulumi.String(vars.HelmChartRepoUrl)},
		}, pulumi.Parent(parent))
	if err != nil {
		return errors.Wrap(err, "deploying NATS chart failed")
	}

	ctx.Export(OpMetricsEndpoint, pulumi.Sprintf(
		"http://nats-prom.%s.svc.cluster.local:7777/metrics", locals.Namespace))
	return nil
}
