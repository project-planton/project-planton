// Package kindscope answers one question about a Kubernetes manifest document:
// does its kind live in a namespace? Modules that apply arbitrary documents
// need the answer before anything applies -- a namespaced document with no
// namespace of its own and no anchor would otherwise land silently in
// "default" -- and both engines must answer identically.
//
// The answer comes from three places, in order:
//   - the built-in kinds: the cluster-scoped ones (cluster_scoped.txt) and the
//     API groups Kubernetes itself serves (builtin_groups.txt), whose other
//     kinds are namespaced -- both read from the API discovery of real
//     clusters by the refresh tool and kept as a union, because a kind never
//     changes scope between Kubernetes versions;
//   - the CustomResourceDefinitions in the same manifest (their spec.scope);
//   - the cluster's own CustomResourceDefinitions, which the module reads
//     only when it needs them.
//
// Everything else is treated as namespaced: a kind the table and the CRDs do
// not know (an aggregated API, say) must carry its own namespace or an
// anchor, which refuses a guess rather than making one.
//
// The OpenTofu copy of the table is generated into the modules that apply
// documents (pkg/iac/tofu/generators, KindScopeTF), with a drift test.
package kindscope

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed cluster_scoped.txt
var clusterScopedSource string

//go:embed builtin_groups.txt
var builtinGroupsSource string

// clusterScoped is the built-in table, keyed as Key(group, kind).
var clusterScoped = parse(clusterScopedSource)

// builtinGroups are the API groups Kubernetes itself serves ("" is the core
// group, written "core" in the file). A kind of one of these groups that the
// table does not list is namespaced.
var builtinGroups = parseGroups(builtinGroupsSource)

// Key is the table's spelling of a kind: "<group>/<Kind>", or the bare kind
// for the core group ("Namespace").
func Key(group, kind string) string {
	if group == "" {
		return kind
	}
	return group + "/" + kind
}

// GroupOf returns the API group of an apiVersion ("rbac.authorization.k8s.io/v1"
// -> "rbac.authorization.k8s.io", "v1" -> "").
func GroupOf(apiVersion string) string {
	if group, _, ok := strings.Cut(apiVersion, "/"); ok {
		return group
	}
	return ""
}

// IsBuiltinGroup reports whether Kubernetes itself serves an API group.
func IsBuiltinGroup(group string) bool {
	return builtinGroups[group]
}

// BuiltinGroups lists the built-in API groups, sorted ("" first, the core
// group).
func BuiltinGroups() []string {
	groups := make([]string, 0, len(builtinGroups))
	for g := range builtinGroups {
		groups = append(groups, g)
	}
	sort.Strings(groups)
	return groups
}

// IsBuiltinClusterScoped reports whether a built-in kind is cluster-scoped.
func IsBuiltinClusterScoped(apiVersion, kind string) bool {
	return clusterScoped[Key(GroupOf(apiVersion), kind)]
}

// BuiltinClusterScoped lists the table's keys, sorted.
func BuiltinClusterScoped() []string {
	keys := make([]string, 0, len(clusterScoped))
	for k := range clusterScoped {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func parse(source string) map[string]bool {
	out := map[string]bool{}
	for _, line := range strings.Split(source, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		out[line] = true
	}
	return out
}

func parseGroups(source string) map[string]bool {
	out := map[string]bool{}
	for line := range parse(source) {
		if line == "core" {
			line = ""
		}
		out[line] = true
	}
	return out
}
