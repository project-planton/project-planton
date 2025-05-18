package module

// vars groups together simple constants used across multiple files.
// Keeping them in a single struct mirrors Terraform’s `variables.tf`.
var vars = struct {
	// Helm chart info
	HelmChartName    string
	HelmChartRepoUrl string
	HelmChartVersion string

	// Fixed NATS client port
	NatsClientPort int

	// Secret keys we create / reference
	AuthSecretKey         string // bearer-token mode
	AuthSecretKeyUser     string // basic-auth mode
	AuthSecretKeyPassword string // basic-auth mode
	TlsCertKey            string // key holding cert+key bundle in TLS Secret
}{
	HelmChartName:    "nats",
	HelmChartRepoUrl: "https://nats-io.github.io/k8s/helm/charts",
	HelmChartVersion: "1.3.6",

	NatsClientPort: 4222,

	// Bearer token mode: { name: auth-<ns>, key: token }
	AuthSecretKey: "token",

	// Basic auth mode: { name: auth-<ns>, key: user | password }
	AuthSecretKeyUser:     "user",
	AuthSecretKeyPassword: "password",

	// TLS Secret always uses the standard keys trusted by most charts:
	// { name: tls-<ns>, key: tls.crt / tls.key }. We export only the name +
	// the "tls.crt" key to stay within a single KubernetesSecretKey field.
	TlsCertKey: "tls.crt",
}
