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
	AdminAuthSecretName       string // bearer-token mode
	AdminAuthSecretKey        string // bearer-token mode
	NatsUserSecretKeyUsername string // basic-auth mode
	NatsUserSecretKeyPassword string // basic-auth mode

	NoAuthUserSecretName string // basic-auth mode

	TlsCertKey string // key holding cert+key bundle in TLS Secret

	AdminUsername string
	// Username assigned to unauthenticated ("no-auth") clients.
	NoAuthUsername              string
	NoAuthPassword              string
	AdminUserPasswordEnvVarName string
}{
	HelmChartName:    "nats",
	HelmChartRepoUrl: "https://nats-io.github.io/k8s/helm/charts",
	HelmChartVersion: "1.3.6",

	NatsClientPort:      4222,
	AdminAuthSecretName: "auth-nats", // Name of the secret holding admin credentials

	// Bearer token mode: { name: auth-<ns>, key: token }
	AdminAuthSecretKey: "token",

	// Basic auth mode: { name: auth-<ns>, key: user | password }
	NatsUserSecretKeyUsername: "user",
	NatsUserSecretKeyPassword: "password",

	NoAuthUserSecretName: "no-auth-user",

	// TLS Secret always uses the standard keys trusted by most charts:
	// { name: tls-<ns>, key: tls.crt / tls.key }. We export only the name +
	// the "tls.crt" key to stay within a single KubernetesSecretKey field.
	TlsCertKey: "tls.crt",

	AdminUsername: "nats",

	// Fixed username for unauthenticated / anonymous clients.
	NoAuthUsername: "noauth",
	NoAuthPassword: "nopassword", // not used, but kept for consistency

	AdminUserPasswordEnvVarName: "NATS_ADMIN_PASSWORD", // env var name for admin password
}
