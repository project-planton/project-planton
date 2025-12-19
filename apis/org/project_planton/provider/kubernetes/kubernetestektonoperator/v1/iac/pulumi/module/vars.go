package module

// vars groups all operator-level constants so version bumps or repo changes
// remain one-liner edits.
var vars = struct {
	OperatorNamespace   string
	ComponentsNamespace string
	OperatorReleaseURL  string
	OperatorVersion     string
	TektonConfigName    string
}{
	OperatorNamespace:   "tekton-operator",
	ComponentsNamespace: "tekton-pipelines",
	OperatorReleaseURL:  "https://storage.googleapis.com/tekton-releases/operator/latest/release.yaml",
	OperatorVersion:     "latest",
	TektonConfigName:    "config",
}
