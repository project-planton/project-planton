package module

// vars groups all operator-level constants so version bumps or repo changes
// remain one-liner edits.
var vars = struct {
	OperatorNamespace        string
	ComponentsNamespace      string
	OperatorReleaseURLFormat string
	TektonConfigName         string
}{
	OperatorNamespace:   "tekton-operator",
	ComponentsNamespace: "tekton-pipelines",
	// Release URL format: %s is replaced with the version (e.g., v0.78.0)
	// https://github.com/tektoncd/operator/releases
	OperatorReleaseURLFormat: "https://storage.googleapis.com/tekton-releases/operator/previous/%s/release.yaml",
	TektonConfigName:         "config",
}
