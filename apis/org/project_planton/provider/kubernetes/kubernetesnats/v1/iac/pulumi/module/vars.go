package module

// vars groups together simple constants used across multiple files.
// Keeping them in a single struct mirrors Terraform's `variables.tf`.
// Note: Secret NAMES are now computed in locals.go to avoid conflicts
// when multiple instances share a namespace. Only secret KEYS are kept here.
var vars = struct {
	// Helm chart info
	HelmChartName    string
	HelmChartRepoUrl string
	HelmChartVersion string

	// Fixed NATS client port
	NatsClientPort int

	// Secret keys we create / reference (note: these are keys WITHIN secrets, not secret names)
	AdminAuthSecretKey        string // bearer-token mode
	NatsUserSecretKeyUsername string // basic-auth mode
	NatsUserSecretKeyPassword string // basic-auth mode

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

	NatsClientPort: 4222,

	// Bearer token mode: key within the secret
	AdminAuthSecretKey: "token",

	// Basic auth mode: keys within the secret
	NatsUserSecretKeyUsername: "user",
	NatsUserSecretKeyPassword: "password",

	// TLS Secret always uses the standard keys trusted by most charts:
	// { key: tls.crt / tls.key }. We export only the key name to stay within
	// a single KubernetesSecretKey field.
	TlsCertKey: "tls.crt",

	AdminUsername: "nats",

	// Fixed username for unauthenticated / anonymous clients.
	NoAuthUsername: "noauth",
	NoAuthPassword: "nopassword", // not used, but kept for consistency

	AdminUserPasswordEnvVarName: "NATS_ADMIN_PASSWORD", // env var name for admin password
}
