// pkg/crkreflect/cmd/genkindmap/main.go
// Generates pkg/crkreflect/kind_map_gen.go.
//
// Run:  go run ./pkg/crkreflect/cmd/genkindmap
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/project-planton/project-planton/pkg/crkreflect"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
)

// -----------------------------------------------------------------------------
// helpers
// -----------------------------------------------------------------------------

// pascalFromSnake converts "digital_ocean" → "DigitalOcean".
func pascalFromSnake(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) == 0 {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

// fixDigitCase: Neo4j → Neo4J; CloudflareKVNamespace → CloudflareKvNamespace
// (matches message names in generated code)
func fixDigitCase(s string) string {
	r := []rune(s)

	// 1. upper‑case the letter that follows any digit (existing rule)
	for i := 0; i < len(r)-1; i++ {
		if r[i] >= '0' && r[i] <= '9' && r[i+1] >= 'a' && r[i+1] <= 'z' {
			r[i+1] -= 'a' - 'A'
		}
	}

	// 2. inside runs of ≥2 consecutive upper‑case letters,
	//    lower‑case every rune except the first so that "KVNamespace" → "KvNamespace".
	for i := 1; i < len(r)-1; i++ {
		if r[i-1] >= 'A' && r[i-1] <= 'Z' &&
			r[i] >= 'A' && r[i] <= 'Z' &&
			r[i+1] >= 'A' && r[i+1] <= 'Z' {
			r[i] += 'a' - 'A'
		}
	}

	return string(r)
}

// lowerNoSep: AwsAlb → awsalb
func lowerNoSep(s string) string { return strings.ToLower(strings.ReplaceAll(s, "_", "")) }

// -----------------------------------------------------------------------------
// types for template
// -----------------------------------------------------------------------------

type importInfo struct{ Alias, Path string }
type entry struct {
	KindConst, Alias, MessageType string
}

// -----------------------------------------------------------------------------
// main
// -----------------------------------------------------------------------------

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	provEntries := map[string][]entry{} // provider raw name ("digital_ocean") → []entry
	k8sAddon, k8sWorkload := []entry{}, []entry{}

	imports := []importInfo{}
	aliasByPath := map[string]string{}

	for _, cloudResourceKind := range crkreflect.KindsList() {
		provider := crkreflect.GetProvider(cloudResourceKind)
		if provider == cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified ||
			provider == cloudresourcekind.CloudResourceProvider_test {
			// skip test values
			continue
		}

		kindName := cloudResourceKind.String() // e.g. AwsAlb

		provRaw := provider.String()                     // "digital_ocean"
		provSlug := strings.ReplaceAll(provRaw, "_", "") // "digitalocean"

		lowerKind := lowerNoSep(kindName) // awsalb
		importAlias := lowerKind + "v1"   // awsalbv1

		// kubernetes special‑case
		var importPath string
		if provRaw == cloudresourcekind.CloudResourceProvider_kubernetes.String() {
			kubernetesResourceType := crkreflect.GetKubernetesResourceCategory(cloudResourceKind)

			importPath = fmt.Sprintf(
				"github.com/project-planton/project-planton/apis/project/planton/provider/%s/%s/%s/v1",
				provSlug, kubernetesResourceType, lowerKind)

			putUniqueEntry(kubernetesResourceType == cloudresourcekind.KubernetesCloudResourceCategory_addon,
				&k8sAddon, &k8sWorkload,
				entry{
					KindConst:   kindName,
					Alias:       uniqueAlias(importPath, importAlias, &imports, aliasByPath),
					MessageType: fixDigitCase(kindName),
				})
			continue
		}

		// non‑kubernetes
		importPath = fmt.Sprintf(
			"github.com/project-planton/project-planton/apis/project/planton/provider/%s/%s/v1",
			provSlug, lowerKind)

		alias := uniqueAlias(importPath, importAlias, &imports, aliasByPath)
		provEntries[provRaw] = append(provEntries[provRaw], entry{
			KindConst:   kindName,
			Alias:       alias,
			MessageType: fixDigitCase(kindName),
		})
	}

	// deterministic output
	for _, list := range provEntries {
		sort.Slice(list, func(i, j int) bool { return list[i].KindConst < list[j].KindConst })
	}
	sort.Slice(k8sAddon, func(i, j int) bool { return k8sAddon[i].KindConst < k8sAddon[j].KindConst })
	sort.Slice(k8sWorkload, func(i, j int) bool { return k8sWorkload[i].KindConst < k8sWorkload[j].KindConst })
	sort.Slice(imports, func(i, j int) bool { return imports[i].Alias < imports[j].Alias })

	// render
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, struct {
		Imports     []importInfo
		ProvEntries map[string][]entry
		K8sAddon    []entry
		K8sWorkload []entry
		Providers   []string
	}{
		Imports:     imports,
		ProvEntries: provEntries,
		K8sAddon:    k8sAddon,
		K8sWorkload: k8sWorkload,
		Providers:   sortedKeys(provEntries),
	}); err != nil {
		return errors.Wrap(err, "execute template")
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		src = buf.Bytes() // keep raw on formatting failure
	}

	outPath := filepath.Join("pkg", "crkreflect", "kind_map_gen.go")
	if err := os.WriteFile(outPath, src, 0o644); err != nil {
		return errors.Wrapf(err, "write %s", outPath)
	}
	fmt.Printf("created %s\n", outPath)
	return nil
}

func uniqueAlias(path, base string, imports *[]importInfo, seen map[string]string) string {
	if a, ok := seen[path]; ok {
		return a
	}
	alias := base
	for i := 1; aliasExists(*imports, alias); i++ {
		alias = fmt.Sprintf("%s_%d", base, i)
	}
	*imports = append(*imports, importInfo{alias, path})
	seen[path] = alias
	return alias
}

func aliasExists(imps []importInfo, a string) bool {
	for _, v := range imps {
		if v.Alias == a {
			return true
		}
	}
	return false
}

func putUniqueEntry(isAddon bool, addon, workload *[]entry, e entry) {
	if isAddon {
		*addon = append(*addon, e)
	} else {
		*workload = append(*workload, e)
	}
}

func sortedKeys(m map[string][]entry) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// -----------------------------------------------------------------------------
// template
// -----------------------------------------------------------------------------

var tpl = template.Must(template.New("").Funcs(template.FuncMap{
	"pascal": pascalFromSnake,
}).Parse(`// Code generated by pkg/crkreflect/cmd/genkindmap; DO NOT EDIT.
package crkreflect

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"google.golang.org/protobuf/proto"
{{- range .Imports }}
	{{ .Alias }} "{{ .Path }}"
{{- end }}
)

func merge(maps ...map[cloudresourcekind.CloudResourceKind]proto.Message) map[cloudresourcekind.CloudResourceKind]proto.Message {
	out := make(map[cloudresourcekind.CloudResourceKind]proto.Message)
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

{{/* provider maps */}}
{{- range $prov, $list := .ProvEntries }}
var Provider{{ pascal $prov }}Map = map[cloudresourcekind.CloudResourceKind]proto.Message{
{{- range $list }}
	cloudresourcekind.CloudResourceKind_{{ .KindConst }}: &{{ .Alias }}.{{ .MessageType }}{},
{{- end }}
}
{{ end }}

{{/* kubernetes */}}
var ProviderKubernetesAddonMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
{{- range .K8sAddon }}
	cloudresourcekind.CloudResourceKind_{{ .KindConst }}: &{{ .Alias }}.{{ .MessageType }}{},
{{- end }}
}

var ProviderKubernetesWorkloadMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
{{- range .K8sWorkload }}
	cloudresourcekind.CloudResourceKind_{{ .KindConst }}: &{{ .Alias }}.{{ .MessageType }}{},
{{- end }}
}

var ProviderKubernetesMap = merge(ProviderKubernetesAddonMap, ProviderKubernetesWorkloadMap)

var ToMessageMap = merge(
{{- range .Providers }}
	Provider{{ pascal . }}Map,
{{- end }}
	ProviderKubernetesMap,
)
`))
