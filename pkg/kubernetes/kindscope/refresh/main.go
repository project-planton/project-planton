// Command refresh adds every built-in cluster-scoped kind a cluster's API
// discovery reports to pkg/kubernetes/kindscope/cluster_scoped.txt. Kinds a
// CustomResourceDefinition serves are left out (their scope comes from the CRD
// itself), and nothing already in the table is removed: a kind never changes
// scope, so the table is the union of every cluster it was read from. Run it
// against each Kubernetes version the e2e harness's cluster profiles run, then
// regenerate the OpenTofu copy:
//
//	go run ./pkg/kubernetes/kindscope/refresh --kubeconfig <path> [--kubeconfig <path>...]
//	PLANTON_REGEN_KINDSCOPE_TF=1 go test ./pkg/iac/tofu/generators/ -run TestKindScopeTFDrift
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/plantonhq/planton/pkg/kubernetes/kindscope"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	tablePath  = "pkg/kubernetes/kindscope/cluster_scoped.txt"
	groupsPath = "pkg/kubernetes/kindscope/builtin_groups.txt"
)

const groupsHeader = `# The API groups Kubernetes itself serves ("core" is the core group), read
# from real clusters' API discovery by pkg/kubernetes/kindscope/refresh and
# kept as a union. A kind of one of these groups that cluster_scoped.txt does
# not list is namespaced. Never edit by hand.
`

const header = `# Built-in Kubernetes kinds that are cluster-scoped, as "<group>/<Kind>" (the
# bare kind for the core group). Read from real clusters' API discovery by
# pkg/kubernetes/kindscope/refresh and kept as a union -- a kind never changes
# scope -- so a cluster newer or older than any one source is answered too.
# Never edit by hand; regenerate the OpenTofu copy after a refresh.
`

type kubeconfigs []string

func (k *kubeconfigs) String() string     { return strings.Join(*k, ",") }
func (k *kubeconfigs) Set(v string) error { *k = append(*k, v); return nil }

func main() {
	var paths kubeconfigs
	flag.Var(&paths, "kubeconfig", "a cluster to read discovery from (repeatable)")
	flag.Parse()
	if len(paths) == 0 {
		fatal("at least one --kubeconfig is required")
	}
	keys := map[string]bool{}
	for _, k := range kindscope.BuiltinClusterScoped() {
		keys[k] = true
	}
	groups := map[string]bool{}
	for _, g := range kindscope.BuiltinGroups() {
		groups[g] = true
	}
	for _, path := range paths {
		added, addedGroups, err := read(path)
		if err != nil {
			fatal("%s: %v", path, err)
		}
		for _, k := range added {
			keys[k] = true
		}
		for _, g := range addedGroups {
			groups[g] = true
		}
	}
	groupLines := make([]string, 0, len(groups))
	for g := range groups {
		if g == "" {
			g = "core"
		}
		groupLines = append(groupLines, g)
	}
	sort.Strings(groupLines)
	if err := os.WriteFile(groupsPath, []byte(groupsHeader+strings.Join(groupLines, "\n")+"\n"), 0o644); err != nil {
		fatal("writing %s: %v", groupsPath, err)
	}
	sorted := make([]string, 0, len(keys))
	for k := range keys {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	if err := os.WriteFile(tablePath, []byte(header+strings.Join(sorted, "\n")+"\n"), 0o644); err != nil {
		fatal("writing %s: %v", tablePath, err)
	}
	fmt.Printf("%s: %d cluster-scoped kinds\n", tablePath, len(sorted))
}

// read returns the cluster-scoped kinds one cluster serves and the API groups
// it serves, minus everything its CustomResourceDefinitions serve.
func read(kubeconfigPath string) ([]string, []string, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, nil, err
	}
	crds, err := apiextensionsclientset.NewForConfigOrDie(config).ApiextensionsV1().CustomResourceDefinitions().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("listing CRDs: %w", err)
	}
	served := map[string]bool{}
	crdGroups := map[string]bool{}
	for _, crd := range crds.Items {
		served[kindscope.Key(crd.Spec.Group, crd.Spec.Names.Kind)] = true
		crdGroups[crd.Spec.Group] = true
	}
	client, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	lists, err := client.ServerPreferredResources()
	if err != nil {
		return nil, nil, fmt.Errorf("discovery: %w", err)
	}
	var out, groups []string
	for _, list := range lists {
		group := kindscope.GroupOf(list.GroupVersion)
		// A group a CustomResourceDefinition serves answers from the CRD; every
		// other group's kinds are reported here with their real scope.
		if !crdGroups[group] {
			groups = append(groups, group)
		}
		for _, r := range list.APIResources {
			if r.Namespaced || strings.Contains(r.Name, "/") {
				continue
			}
			key := kindscope.Key(group, r.Kind)
			if !served[key] {
				out = append(out, key)
			}
		}
	}
	return out, groups, nil
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "kindscope refresh: "+format+"\n", args...)
	os.Exit(1)
}
