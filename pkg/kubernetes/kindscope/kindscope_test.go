package kindscope

import (
	"errors"
	"testing"
)

const mixed = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: reader
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: widgets.example.com
spec:
  group: example.com
  scope: Cluster
  names:
    kind: Widget
---
apiVersion: example.com/v1
kind: Widget
metadata:
  name: one
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: placed
  namespace: apps
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: unplaced
`

func noCluster(group, kind string) (bool, bool, error) {
	return false, false, errors.New("the cluster must not be read when every kind is known locally")
}

func TestUnplacedNamesOnlyNamespacedDocumentsWithoutANamespace(t *testing.T) {
	docs, err := ParseDocuments(mixed)
	if err != nil {
		t.Fatal(err)
	}
	unplaced, err := Unplaced(docs, noCluster)
	if err != nil {
		t.Fatal(err)
	}
	if got := Describe(unplaced); got != "v1 ConfigMap unplaced" {
		t.Fatalf("unplaced = %q, want only the ConfigMap that names no namespace", got)
	}
}

func TestClusterScopedOnlyManifestNeedsNoAnchor(t *testing.T) {
	docs, err := ParseDocuments("apiVersion: storage.k8s.io/v1\nkind: StorageClass\nmetadata:\n  name: fast\n---\napiVersion: v1\nkind: Namespace\nmetadata:\n  name: team\n")
	if err != nil {
		t.Fatal(err)
	}
	unplaced, err := Unplaced(docs, noCluster)
	if err != nil || len(unplaced) != 0 {
		t.Fatalf("unplaced = %v (err %v), want none", unplaced, err)
	}
}

func TestClusterCRDsAnswerKindsTheManifestDoesNotDefine(t *testing.T) {
	docs, err := ParseDocuments("apiVersion: cloud.google.com/v1\nkind: ComputeClass\nmetadata:\n  name: default\n---\napiVersion: example.com/v1\nkind: Gadget\nmetadata:\n  name: g\n")
	if err != nil {
		t.Fatal(err)
	}
	lookup := func(group, kind string) (bool, bool, error) {
		if Key(group, kind) == "cloud.google.com/ComputeClass" {
			return false, true, nil
		}
		return false, false, nil // the cluster serves no Gadget: unknown scope
	}
	unplaced, err := Unplaced(docs, lookup)
	if err != nil {
		t.Fatal(err)
	}
	if got := Describe(unplaced); got != "example.com/v1 Gadget g" {
		t.Fatalf("unplaced = %q, want only the kind of unknown scope", got)
	}
}

func TestTableHoldsTheCoreClusterScopedKinds(t *testing.T) {
	for _, k := range []struct{ apiVersion, kind string }{
		{"v1", "Namespace"},
		{"rbac.authorization.k8s.io/v1", "ClusterRole"},
		{"apiextensions.k8s.io/v1", "CustomResourceDefinition"},
		{"storage.k8s.io/v1", "StorageClass"},
	} {
		if !IsBuiltinClusterScoped(k.apiVersion, k.kind) {
			t.Errorf("%s %s is missing from the table", k.apiVersion, k.kind)
		}
	}
	if IsBuiltinClusterScoped("v1", "ConfigMap") {
		t.Error("ConfigMap is namespaced")
	}
}
