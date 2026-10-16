package failure

import (
	"fmt"
	"strings"
)

// signatureKubernetesForbidden is the fragment of KubernetesForbidden's
// observation Explain looks for before speaking.
const signatureKubernetesForbidden = "the identity this deploy runs as"

// signatureKubernetesFieldNotInDefinition is the fragment of
// KubernetesFieldNotInDefinition's observation Explain looks for before
// speaking.
const signatureKubernetesFieldNotInDefinition = "as a field its installed definition does not know"

// CustomResourceDefinitionsResource is the resource name the API server uses
// for CRDs in its authorization texts.
const CustomResourceDefinitionsResource = "customresourcedefinitions"

// definitionOwners names, for a custom resource kind the catalog declares,
// the catalog resource that installs its definition and the field on it that
// picks the definition's version. A kind absent here gets the generic next
// step (upgrade whatever installs the definition).
//
// Kept to kinds whose definition another catalog resource owns; a module that
// applies its own chart's definitions never hits this failure, because its
// definitions and its declaration come from the same release.
var definitionOwners = map[string]string{
	"PlantonPlatform": "KubernetesPlantonOperator (spec.chart_version)",
}

// KubernetesFieldNotInDefinition: the API server rejected a field of a custom
// resource because the kind's installed definition (its CRD) predates the
// field. The catalog release the declaration comes from knows the field; the
// definition on the cluster does not — the definition was installed from an
// older release of the chart that owns it, and nothing upgrades that chart
// implicitly. Under server-side apply the server refuses outright (an unknown
// field is never pruned silently), but its sentence names only the mechanism.
// kind is the API server's Kind; field is the dotted path without a leading
// dot (spec.ingress.gatewayRef); raw is the server's text, kept whole.
func KubernetesFieldNotInDefinition(kind, field, raw string) *Failure {
	owner, known := definitionOwners[kind]
	next := fmt.Sprintf("upgrade the resource that installs the %s definition to the current catalog default or newer, apply it, then re-apply this resource", kind)
	if known {
		next = fmt.Sprintf("set %s to the current catalog default or newer, apply it, then re-apply this resource", owner)
	}
	return &Failure{
		Observed: fmt.Sprintf("the API server rejected %s on %s %s (the API server answered: %s)", field, kind, signatureKubernetesFieldNotInDefinition, raw),
		Meaning:  fmt.Sprintf("the definition of %s installed on this cluster predates %s: the catalog release this declaration comes from knows the field, the chart that installed the definition is older than it, and nothing upgrades that chart implicitly", kind, field),
		NextStep: next,
	}
}

// WorkloadOwnerLabel is the label a kind stamps on the workloads it installs
// under fixed, upstream-chosen names, and the mark the workload-ownership
// check reads: a Deployment carrying it with the kind's own name was created
// by that kind (planton.ai/kind, the identity label the Pulumi modules
// stamp).
const WorkloadOwnerLabel = "planton.ai/kind"

// KubernetesWorkloadOwnedElsewhere: a kind that installs an upstream's own
// workloads under fixed names (one install per cluster is the upstream
// contract) found one of them already running, created by something else: a
// managed add-on that installs the same software, a `kubectl apply` of the
// upstream's release, another tool. Server-side apply would take it over
// without a word, two owners would then fight over one controller, and a
// later destroy would delete workloads this kind never created. So the
// deploy stops before writing anything. owner is what the cluster says about
// who wrote it (the field managers), or the honest absence of any mark.
func KubernetesWorkloadOwnedElsewhere(kind, namespace, name, owner string) *Failure {
	return &Failure{
		Observed: fmt.Sprintf("Deployment %s/%s already runs on the cluster and was not created by this %s (%s)", namespace, name, kind, owner),
		Meaning:  "another installation of the same software runs here, such as a managed add-on or a kubectl apply of the upstream release; installing over it would take over its workloads, and destroying this resource later would delete them",
		NextStep: fmt.Sprintf("remove the other installation first, or keep it and do not install %s on this cluster; to hand over an installation you made yourself, import it, then mark it as this kind's (`kubectl -n %s label deployment %s %s=%s`) and re-run", kind, namespace, name, WorkloadOwnerLabel, kind),
	}
}

// KubernetesNamespaceMissing: a kind that installs into an upstream-fixed
// namespace was told another component owns it (create_namespace false), and
// the namespace does not exist. Nothing could start; the usual cause is that
// the namespace's owner has not been deployed yet, or the two were deployed
// without a reference ordering them.
func KubernetesNamespaceMissing(kind, namespace string) *Failure {
	return &Failure{
		Observed: fmt.Sprintf("namespace %s does not exist on the cluster, and this %s does not create it (spec.create_namespace is false)", namespace, kind),
		Meaning:  fmt.Sprintf("everything installs into %s, so its owner must create it first; nothing could start without it", namespace),
		NextStep: fmt.Sprintf("deploy the KubernetesNamespace named %s first and reference it from spec.namespace so the two deploy in order, or set spec.create_namespace to true to let this resource create and own it", namespace),
	}
}

// KubernetesNamespaceCreatedHere: create_namespace was turned off on an
// installation whose namespace this kind created. The engines would delete
// the namespace they hold in state, and with it the kind's own workloads and
// everything other components put there, so the change is refused with the
// path that moves ownership without losing anything.
func KubernetesNamespaceCreatedHere(kind, namespace string) *Failure {
	return &Failure{
		Observed: fmt.Sprintf("namespace %s was created by this %s, and spec.create_namespace is now false", namespace, kind),
		Meaning:  "turning create_namespace off would delete the namespace, and with it this installation's workloads and everything other components put there",
		NextStep: fmt.Sprintf("keep spec.create_namespace true; to move the namespace to a KubernetesNamespace, destroy this resource first (its resource definitions and every sandbox outside %s are kept), deploy the KubernetesNamespace, then deploy this resource again with create_namespace false", namespace),
	}
}

// KubernetesNamespaceOwnedElsewhere: a kind asked to create an
// upstream-fixed namespace (create_namespace true) found it already there,
// created by something else. Taking it over would make this resource its
// owner, and destroying the resource would delete the namespace with
// everything other components put in it. owner is what the cluster says
// about who wrote it.
func KubernetesNamespaceOwnedElsewhere(kind, namespace, owner string) *Failure {
	return &Failure{
		Observed: fmt.Sprintf("namespace %s already exists on the cluster and was not created by this %s (%s)", namespace, kind, owner),
		Meaning:  "with spec.create_namespace true this resource would take the namespace over, and destroying it later would delete the namespace and everything other components put in it",
		NextStep: fmt.Sprintf("set spec.create_namespace to false and reference the namespace's owner from spec.namespace; or, if you created it yourself for this installation, mark it as this kind's (`kubectl label namespace %s %s=%s`) and re-run", namespace, WorkloadOwnerLabel, kind),
	}
}

// KubernetesForbidden: the API server refused a verb to the identity the
// deploy authenticates as. The server's own text names the identity, the verb,
// the resource, and the scope, which is the whole observation; what it never
// says is why the module needs that right and where the module lists every
// right it needs. scope is the server's phrase ("at the cluster scope" or
// `in the namespace "x"`); raw is the server's text, kept whole.
//
// A denied write on CustomResourceDefinitions gets the fuller answer: a module
// on the catalog's derive branch applies its CRDs itself, as kept resources
// apart from the release or manifest it installs, so it needs cluster-scoped
// CRD rights a namespace-admin identity never has, and the user has a way
// around it. The words name no source (chart or upstream manifest) because
// the Terraform twin's runner explains the raw text without knowing which.
func KubernetesForbidden(user, verb, resource, group, scope, raw string) *Failure {
	qualified := resource
	if group != "" {
		qualified = resource + "." + group
	}
	observed := fmt.Sprintf("%s (%s) may not %s %s %s (the API server answered: %s)", signatureKubernetesForbidden, user, verb, qualified, scope, raw)
	if resource == CustomResourceDefinitionsResource {
		return &Failure{
			Observed: observed,
			Meaning:  "the module applies its CustomResourceDefinitions itself, as kept resources apart from the release or manifest it installs, so it needs cluster-scoped rights on CRDs that a namespace-admin identity does not have",
			NextStep: fmt.Sprintf("grant the rules in the module's iac/permissions.yaml (customresourcedefinitions: get, list, create, update, patch, and delete when crds.keep_on_uninstall is false) to %s, or run the deploy with an identity that has them; or set spec.crds.install to false and have a cluster administrator apply the CRDs (for a Helm chart, `helm template --include-crds` renders them; the kind's README names where its CRDs come from)", user),
		}
	}
	return &Failure{
		Observed: observed,
		Meaning:  fmt.Sprintf("the module creates or manages %s and the identity running the deploy was never granted %s on them %s", qualified, verb, strings.TrimSpace(scope)),
		NextStep: fmt.Sprintf("grant %s on %s %s to %s (the module's iac/permissions.yaml lists every rule it needs), or run the deploy with an identity that has it", verb, qualified, strings.TrimSpace(scope), user),
	}
}
