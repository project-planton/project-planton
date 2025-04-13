// File: variables.go
package module

// vars contains all the configurable or default values used in this Pulumi module.
var vars = struct {
	// The name of the Helm chart to install.
	Neo4jHelmChartName string

	// The Helm repo URL hosting the chart.
	Neo4jHelmChartRepoUrl string

	// The version of the Helm chart to install.
	Neo4jHelmChartVersion string

	// The secret key for the admin password inside the Kubernetes secret.
	Neo4jPasswordSecretKey string

	// The name of the Kubernetes secret that stores the admin password.
	Neo4jPasswordSecretName string
}{
	Neo4jHelmChartName:      "neo4j",
	Neo4jHelmChartRepoUrl:   "https://helm.neo4j.com/neo4j",
	Neo4jHelmChartVersion:   "2025.03.0",
	Neo4jPasswordSecretKey:  "password",
	Neo4jPasswordSecretName: "neo4j-admin-password",
}
