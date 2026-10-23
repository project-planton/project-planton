package generators

import (
	"fmt"
	"strings"

	"github.com/plantonhq/planton/pkg/kubernetes/kindscope"
)

// KindScopeTFFileName is the generated file a module that applies arbitrary
// Kubernetes documents carries: the built-in cluster-scoped kinds, as a local.
const KindScopeTFFileName = "kindscope.tf"

// KindScopeTF renders the OpenTofu copy of pkg/kubernetes/kindscope's table, so
// the module decides a document's scope from the same list the Pulumi twin
// reads. Deterministic: the table is sorted.
func KindScopeTF() string {
	var b strings.Builder
	b.WriteString(`# Generated from pkg/kubernetes/kindscope/cluster_scoped.txt by
# pkg/iac/tofu/generators (KindScopeTF) -- never edit by hand. After a table
# refresh, regenerate with:
#   PLANTON_REGEN_KINDSCOPE_TF=1 go test ./pkg/iac/tofu/generators/ -run TestKindScopeTFDrift
#
# The built-in Kubernetes kinds that are cluster-scoped, keyed "<group>/<Kind>"
# (the bare kind for the core group), and the API groups Kubernetes itself
# serves ("" is the core group), whose other kinds are namespaced: the same
# tables the Pulumi module reads.
locals {
  kindscope_builtin_cluster_scoped = toset([
`)
	for _, key := range kindscope.BuiltinClusterScoped() {
		fmt.Fprintf(&b, "    %q,\n", key)
	}
	b.WriteString("  ])\n  kindscope_builtin_groups = toset([\n")
	for _, group := range kindscope.BuiltinGroups() {
		fmt.Fprintf(&b, "    %q,\n", group)
	}
	b.WriteString("  ])\n}\n")
	return b.String()
}
