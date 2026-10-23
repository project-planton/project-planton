package kindscope

import (
	"strings"

	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

// Document is what a scope decision reads from one manifest document.
type Document struct {
	APIVersion string
	Kind       string
	Name       string
	// Namespace is the document's own metadata.namespace ("" when it names
	// none).
	Namespace string
	// Defines is set on a CustomResourceDefinition document: the kind it
	// serves and its scope, which answers documents of that kind in the same
	// manifest.
	Defines *DefinedKind
}

// DefinedKind is the kind a CustomResourceDefinition serves.
type DefinedKind struct {
	Group      string
	Kind       string
	Namespaced bool
}

type documentHeader struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Group string `json:"group"`
		Scope string `json:"scope"`
		Names struct {
			Kind string `json:"kind"`
		} `json:"names"`
	} `json:"spec"`
}

// ParseDocuments splits a multi-document manifest on its "---" separators,
// the way both engines' modules split it, and reads each document's header.
// Empty documents are skipped.
func ParseDocuments(manifestYaml string) ([]Document, error) {
	var docs []Document
	for _, raw := range strings.Split("\n"+manifestYaml, "\n---") {
		if strings.TrimSpace(raw) == "" {
			continue
		}
		var header documentHeader
		if err := yaml.Unmarshal([]byte(raw), &header); err != nil {
			return nil, errors.Wrap(err, "the manifest contains a document that is not valid YAML")
		}
		if header.APIVersion == "" && header.Kind == "" && header.Metadata.Name == "" {
			continue
		}
		doc := Document{
			APIVersion: header.APIVersion,
			Kind:       header.Kind,
			Name:       header.Metadata.Name,
			Namespace:  header.Metadata.Namespace,
		}
		if GroupOf(header.APIVersion) == "apiextensions.k8s.io" && header.Kind == "CustomResourceDefinition" {
			doc.Defines = &DefinedKind{
				Group:      header.Spec.Group,
				Kind:       header.Spec.Names.Kind,
				Namespaced: header.Spec.Scope != "Cluster",
			}
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// ClusterLookup answers a kind neither the built-in table nor the manifest's
// own CRDs know: whether the cluster serves it, and if so whether it is
// namespaced. It is called only when such a kind appears, so a module reads
// the cluster only when the answer is not already local.
type ClusterLookup func(group, kind string) (namespaced bool, served bool, err error)

// Unplaced returns the documents that a manifest without an anchor namespace
// cannot place: namespaced (or of unknown scope) and naming no namespace of
// their own. An empty result means every document applies as written.
func Unplaced(docs []Document, lookup ClusterLookup) ([]Document, error) {
	defined := map[string]bool{}
	for _, d := range docs {
		if d.Defines != nil {
			defined[Key(d.Defines.Group, d.Defines.Kind)] = d.Defines.Namespaced
		}
	}
	var unplaced []Document
	for _, d := range docs {
		if d.Namespace != "" {
			continue
		}
		group := GroupOf(d.APIVersion)
		key := Key(group, d.Kind)
		if clusterScoped[key] {
			continue
		}
		if builtinGroups[group] {
			unplaced = append(unplaced, d)
			continue
		}
		if namespaced, ok := defined[key]; ok {
			if namespaced {
				unplaced = append(unplaced, d)
			}
			continue
		}
		namespaced, served, err := lookup(group, d.Kind)
		if err != nil {
			return nil, err
		}
		if !served || namespaced {
			unplaced = append(unplaced, d)
		}
	}
	return unplaced, nil
}

// Describe names documents for a refusal: "apps/v1 Deployment web, v1 ConfigMap app".
func Describe(docs []Document) string {
	parts := make([]string, 0, len(docs))
	for _, d := range docs {
		parts = append(parts, strings.TrimSpace(d.APIVersion+" "+d.Kind+" "+d.Name))
	}
	return strings.Join(parts, ", ")
}
