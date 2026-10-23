package module

import (
	"context"
	"os"

	"github.com/pkg/errors"
	kubernetesprovider "github.com/plantonhq/planton/catalog/kubernetes"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/iacinput"
	"github.com/plantonhq/planton/pkg/kubernetes/execcredential"
	"github.com/plantonhq/planton/pkg/kubernetes/kindscope"
	"github.com/plantonhq/planton/pkg/kubernetes/kubeconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

var crdGVR = schema.GroupVersionResource{
	Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions",
}

// refuseUnplacedDocuments is the no-anchor half of the namespace semantics:
// with spec.namespace omitted, a namespaced document that names no namespace
// of its own has nowhere it was asked to go, so the deploy stops before
// anything registers -- naming the documents -- instead of letting the
// provider put them in "default". Scope comes from pkg/kubernetes/kindscope,
// the table the Terraform module's kindscope.tf is generated from; the
// cluster's CRDs are read only when a kind is neither built in nor defined
// by the manifest itself. Destroy runs no check: what was applied is removed
// as it was applied.
func refuseUnplacedDocuments(locals *Locals, providerConfig *kubernetesprovider.KubernetesProviderConfig) error {
	if locals.Namespace != "" || iacinput.IsDestroy() {
		return nil
	}
	docs, err := kindscope.ParseDocuments(locals.Spec.GetManifestYaml())
	if err != nil {
		return err
	}
	var clusterCRDs map[string]bool
	lookup := func(group, kind string) (bool, bool, error) {
		if clusterCRDs == nil {
			read, err := readClusterCRDScopes(providerConfig)
			if err != nil {
				return false, false, err
			}
			clusterCRDs = read
		}
		namespaced, served := clusterCRDs[kindscope.Key(group, kind)]
		return namespaced, served, nil
	}
	unplaced, err := kindscope.Unplaced(docs, lookup)
	if err != nil {
		return err
	}
	if len(unplaced) > 0 {
		return errors.Errorf("spec.namespace is not set, and these namespaced documents name no metadata.namespace of their own: %s -- give each its namespace, or set spec.namespace as their anchor",
			kindscope.Describe(unplaced))
	}
	return nil
}

// readClusterCRDScopes reads the cluster's CustomResourceDefinitions through
// the connection the provider uses: each served kind and whether it is
// namespaced.
func readClusterCRDScopes(providerConfig *kubernetesprovider.KubernetesProviderConfig) (map[string]bool, error) {
	restConfig, err := kubeconfig.RESTConfig(providerConfig, os.Getenv(execcredential.CommandPathEnvVar))
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve the cluster connection to read its CRDs")
	}
	client, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the cluster client to read its CRDs")
	}
	list, err := client.Resource(crdGVR).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list the cluster's CRDs to decide document scope")
	}
	scopes := map[string]bool{}
	for _, crd := range list.Items {
		group, _, _ := unstructured.NestedString(crd.Object, "spec", "group")
		kind, _, _ := unstructured.NestedString(crd.Object, "spec", "names", "kind")
		scope, _, _ := unstructured.NestedString(crd.Object, "spec", "scope")
		scopes[kindscope.Key(group, kind)] = scope != "Cluster"
	}
	return scopes, nil
}
