package verify

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// The agent-sandbox resource kinds (Sandbox, SandboxTemplate, SandboxWarmPool,
// SandboxClaim) are projection kinds: every lane first proves each field its
// scenario sets reads back from the live object, then proves the behavior the
// resource promises, the way a product using agent-sandbox relies on it.
const (
	agentSandboxResource         = "sandboxes.agents.x-k8s.io"
	agentSandboxTemplateResource = "sandboxtemplates.extensions.agents.x-k8s.io"
	agentSandboxWarmPoolResource = "sandboxwarmpools.extensions.agents.x-k8s.io"
	agentSandboxClaimResource    = "sandboxclaims.extensions.agents.x-k8s.io"

	// The suspend lane's proof: a marker file in the workspace claim, and a
	// mark on the claim carrying its own UID (a recreated claim would lack
	// it).
	agentSandboxWorkspaceVolume = "workspace"
	agentSandboxWorkspacePath   = "/workspace"
	agentSandboxMarkerFile      = "/workspace/e2e-marker"
	agentSandboxMarkerContent   = "kept-across-suspend"
	agentSandboxClaimMark       = "e2e.planton.dev/first-act-uid"

	// agentSandboxTemplateRefHashLabel is the label upstream's template
	// controller stamps on every SandboxTemplate it reconciles, first thing.
	agentSandboxTemplateRefHashLabel = "agents.x-k8s.io/sandbox-template-ref-hash"
)

// AgentSandboxVerifier proves a Sandbox: its fields read back, it becomes
// Ready when Running with each declared claim at its declared size mounted in
// its pod (and, with service on, a headless Service whose endpoints are that
// pod), and destroy removes it with its pod and claims. The suspend lane (two
// acts) proves the promise behind operatingMode: suspending deletes the pod
// and keeps the same claim, and resuming brings the files back.
type AgentSandboxVerifier struct {
	Name, Namespace, ManifestPath string
	Suspended                     bool
	// SuspendLane is the behavioral-suspend scenario: act one writes the
	// marker, act two (Suspended) proves it survives.
	SuspendLane bool
	// Claims are the declared volume claim templates; each claim is
	// `<template>-<sandbox>`.
	Claims []agentSandboxVolumeClaim
	// Service is spec.service: the Sandbox asks for its headless Service.
	Service bool
}

// agentSandboxVolumeClaim is one declared volume claim template: its name and
// the storage it requests.
type agentSandboxVolumeClaim struct {
	Name, Storage string
}

// readVolumeClaimTemplates reads the volume claim templates a Sandbox or a
// claim declares.
func readVolumeClaimTemplates(spec map[string]interface{}) []agentSandboxVolumeClaim {
	var claims []agentSandboxVolumeClaim
	templates, _ := specField(spec, "volumeClaimTemplates").([]interface{})
	for _, t := range templates {
		tm, _ := t.(map[string]interface{})
		meta, _ := specField(tm, "metadata").(map[string]interface{})
		name, ok := specField(meta, "name").(string)
		if !ok {
			continue
		}
		claim := agentSandboxVolumeClaim{Name: name}
		if claimSpec, ok := specField(tm, "spec").(map[string]interface{}); ok {
			if resources, ok := specField(claimSpec, "resources").(map[string]interface{}); ok {
				if requests, ok := specField(resources, "requests").(map[string]interface{}); ok {
					claim.Storage = stringOr(specField(requests, "storage"), "")
				}
			}
		}
		claims = append(claims, claim)
	}
	return claims
}

func newAgentSandboxVerifier(manifestPath, name, namespace string) *AgentSandboxVerifier {
	spec := manifestSpecMap(manifestPath)
	v := &AgentSandboxVerifier{
		Name: name, Namespace: namespace, ManifestPath: manifestPath,
		Suspended:   stringOr(specField(spec, "operatingMode"), "Running") == "Suspended",
		SuspendLane: strings.Contains(manifestPath, "behavioral-suspend"),
		Claims:      readVolumeClaimTemplates(spec),
	}
	v.Service, _ = specField(spec, "service").(bool)
	return v
}

func (v *AgentSandboxVerifier) claimName(template string) string {
	return template + "-" + v.Name
}

func (v *AgentSandboxVerifier) VerifyExists(ctx context.Context, kubeconfig string) error {
	fmt.Printf("  [verify] Sandbox %q in %q (suspended=%t)\n", v.Name, v.Namespace, v.Suspended)
	if err := KubectlResourceExists(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace); err != nil {
		return err
	}
	if err := proveProjectionReadback(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace, v.ManifestPath); err != nil {
		return err
	}
	if v.Suspended {
		if err := proveVolumeClaims(ctx, kubeconfig, v.Namespace, v.Name, v.Claims, false); err != nil {
			return err
		}
		return v.proveSuspended(ctx, kubeconfig)
	}
	if err := kubectlWait(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace, "condition=Ready", 5*time.Minute); err != nil {
		return errors.Wrap(err, "the Sandbox never became Ready -- the controller did not bring its pod up")
	}
	fmt.Printf("  [verify] the Sandbox is Ready: its pod %q runs\n", v.Name)
	if err := proveVolumeClaims(ctx, kubeconfig, v.Namespace, v.Name, v.Claims, true); err != nil {
		return err
	}
	if v.Service {
		if err := proveHeadlessService(ctx, kubeconfig, v.Namespace, v.Name); err != nil {
			return err
		}
	}
	if v.SuspendLane {
		return v.markWorkspace(ctx, kubeconfig)
	}
	return nil
}

// markWorkspace is act one of the suspend lane: a file in the workspace, and
// the claim marked with its own UID.
func (v *AgentSandboxVerifier) markWorkspace(ctx context.Context, kubeconfig string) error {
	if err := kubectlExec(ctx, kubeconfig, v.Namespace, v.Name, "sh", "-c",
		fmt.Sprintf("echo %s > %s", agentSandboxMarkerContent, agentSandboxMarkerFile)); err != nil {
		return errors.Wrap(err, "writing the marker file into the workspace")
	}
	claim := v.claimName(agentSandboxWorkspaceVolume)
	uid, err := kubectlGetJSONPath(ctx, kubeconfig, "pvc", claim, v.Namespace, "{.metadata.uid}")
	if err != nil {
		return errors.Wrap(err, "reading the workspace claim's UID")
	}
	if out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "annotate", "pvc", claim, "-n", v.Namespace,
		"--overwrite", agentSandboxClaimMark+"="+strings.TrimSpace(uid)).CombinedOutput(); err != nil {
		return errors.Wrapf(err, "marking the workspace claim: %s", firstLines(string(out), 3))
	}
	fmt.Printf("  [verify] SUSPEND (act one): wrote %s and marked claim %q\n", agentSandboxMarkerFile, claim)
	return nil
}

// proveSuspended is act two: the pod is gone, the Sandbox and the same claim
// remain, and resuming the sandbox brings the marker file back.
func (v *AgentSandboxVerifier) proveSuspended(ctx context.Context, kubeconfig string) error {
	if err := kubectlWait(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace, "condition=Suspended", 3*time.Minute); err != nil {
		return errors.Wrap(err, "the Sandbox never reported Suspended")
	}
	if err := KubectlResourceAbsent(ctx, kubeconfig, "pod", v.Name, v.Namespace); err != nil {
		return errors.Wrap(err, "the Sandbox is Suspended but its pod still exists")
	}
	if !v.SuspendLane {
		fmt.Printf("  [verify] the Sandbox is Suspended with no pod\n")
		return nil
	}
	claim := v.claimName(agentSandboxWorkspaceVolume)
	uid, err := kubectlGetJSONPath(ctx, kubeconfig, "pvc", claim, v.Namespace, "{.metadata.uid}")
	if err != nil {
		return errors.Wrap(err, "the workspace claim is gone after suspending -- the volume was not kept")
	}
	mark, err := kubectlGetJSONPath(ctx, kubeconfig, "pvc", claim, v.Namespace,
		"{.metadata.annotations."+strings.ReplaceAll(agentSandboxClaimMark, ".", "\\.")+"}")
	if err != nil || strings.TrimSpace(mark) != strings.TrimSpace(uid) {
		return errors.Errorf("the workspace claim %q is not the one act one marked (mark %q, uid %q) -- it was recreated, so its data was lost", claim, mark, uid)
	}
	fmt.Printf("  [verify] SUSPEND: the pod is gone and the same workspace claim remains\n")

	// Resume out of band (destroy follows, so the drift is never planned
	// against) and read the file back from the new pod.
	if out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "patch", agentSandboxResource, v.Name, "-n", v.Namespace,
		"--type=merge", "-p", `{"spec":{"operatingMode":"Running"}}`).CombinedOutput(); err != nil {
		return errors.Wrapf(err, "resuming the Sandbox: %s", firstLines(string(out), 3))
	}
	if err := kubectlWait(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace, "condition=Ready", 5*time.Minute); err != nil {
		return errors.Wrap(err, "the resumed Sandbox never became Ready")
	}
	out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "exec", v.Name, "-n", v.Namespace, "--",
		"cat", agentSandboxMarkerFile).CombinedOutput()
	if err != nil || strings.TrimSpace(string(out)) != agentSandboxMarkerContent {
		return errors.Errorf("the resumed sandbox does not have its file back (%s: %q, err %v)", agentSandboxMarkerFile, firstLines(string(out), 2), err)
	}
	fmt.Printf("  [verify] SUSPEND: resumed, and %s came back with the sandbox\n", agentSandboxMarkerFile)
	return nil
}

func (v *AgentSandboxVerifier) VerifyAbsent(ctx context.Context, kubeconfig string) error {
	if err := KubectlResourceAbsent(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace); err != nil {
		return err
	}
	if err := KubectlResourceAbsent(ctx, kubeconfig, "pod", v.Name, v.Namespace); err != nil {
		return errors.Wrap(err, "the Sandbox is gone but its pod remains")
	}
	for _, claim := range v.Claims {
		if err := KubectlResourceAbsent(ctx, kubeconfig, "pvc", v.claimName(claim.Name), v.Namespace); err != nil {
			return errors.Wrapf(err, "the Sandbox is gone but its claim for %q remains", claim.Name)
		}
	}
	return nil
}

// proveVolumeClaims proves each declared claim of a sandbox exists at its
// declared size and, when the sandbox runs, is mounted in its pod: the
// workspace the agent writes to is the one asked for, not a default.
func proveVolumeClaims(ctx context.Context, kubeconfig, namespace, sandbox string, claims []agentSandboxVolumeClaim, running bool) error {
	if len(claims) == 0 {
		return nil
	}
	mounted := ""
	if running {
		out, err := kubectlGetJSONPath(ctx, kubeconfig, "pod", sandbox, namespace, "{.spec.volumes[*].persistentVolumeClaim.claimName}")
		if err != nil {
			return errors.Wrapf(err, "reading the claims sandbox %q's pod mounts", sandbox)
		}
		mounted = " " + strings.TrimSpace(out) + " "
	}
	for _, claim := range claims {
		pvc := claim.Name + "-" + sandbox
		size, err := kubectlGetJSONPath(ctx, kubeconfig, "pvc", pvc, namespace, "{.spec.resources.requests.storage}")
		if err != nil {
			return errors.Wrapf(err, "the controller did not create the claim %q for volume claim template %q", pvc, claim.Name)
		}
		if claim.Storage != "" && strings.TrimSpace(size) != claim.Storage {
			return errors.Errorf("claim %q requests %s, but its volume claim template declares %s", pvc, strings.TrimSpace(size), claim.Storage)
		}
		if running && !strings.Contains(mounted, " "+pvc+" ") {
			return errors.Errorf("claim %q exists, but sandbox %q's pod does not mount it (it mounts:%s)", pvc, sandbox, mounted)
		}
		fmt.Printf("  [verify] CLAIM: %q holds %s and is mounted in the sandbox's pod\n", pvc, strings.TrimSpace(size))
	}
	return nil
}

// proveHeadlessService proves the Service a Sandbox asks for: the Sandbox's
// status names it, it is headless (each sandbox is addressed by its own pod),
// and its endpoints are the sandbox's pod.
func proveHeadlessService(ctx context.Context, kubeconfig, namespace, sandbox string) error {
	service, err := awaitJSONPath(ctx, kubeconfig, agentSandboxResource, sandbox, namespace, "{.status.service}", 2*time.Minute)
	if err != nil {
		return errors.Wrap(err, "the Sandbox asks for a Service, but its status names none")
	}
	service = strings.TrimSpace(service)
	clusterIP, err := kubectlGetJSONPath(ctx, kubeconfig, "service", service, namespace, "{.spec.clusterIP}")
	if err != nil {
		return errors.Wrapf(err, "the Sandbox's status names Service %q, which does not exist", service)
	}
	if strings.TrimSpace(clusterIP) != "None" {
		return errors.Errorf("Service %q has cluster IP %q; a sandbox's Service must be headless", service, strings.TrimSpace(clusterIP))
	}
	podIP, err := kubectlGetJSONPath(ctx, kubeconfig, "pod", sandbox, namespace, "{.status.podIP}")
	if err != nil || strings.TrimSpace(podIP) == "" {
		return errors.Wrapf(err, "reading sandbox %q's pod IP", sandbox)
	}
	podIP = strings.TrimSpace(podIP)
	deadline := time.Now().Add(2 * time.Minute)
	for {
		addresses, err := kubectlGetJSONPathList(ctx, kubeconfig, "endpointslice", namespace,
			"kubernetes.io/service-name="+service, `{range .items[*].endpoints[*].addresses[*]}{@}{"\n"}{end}`)
		if err == nil && containsLine(strings.Join(addresses, "\n"), podIP) {
			break
		}
		if time.Now().After(deadline) {
			return errors.Errorf("headless Service %q never listed the sandbox's pod %s among its endpoints (has %v, %v)", service, podIP, addresses, err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
		}
	}
	fmt.Printf("  [verify] SERVICE: headless Service %q resolves to the sandbox's pod %s\n", service, podIP)
	return nil
}

// AgentSandboxTemplateVerifier proves a SandboxTemplate by what the
// controller does with it, since a template runs nothing: managed with no
// rules, the shared NetworkPolicy carries the secure default (ingress only
// from the router); managed with rules, it carries exactly the declared
// rules; unmanaged, there is none.
type AgentSandboxTemplateVerifier struct {
	Name, Namespace, ManifestPath string
	Managed, CustomRules          bool
}

func newAgentSandboxTemplateVerifier(manifestPath, name, namespace string) *AgentSandboxTemplateVerifier {
	spec := manifestSpecMap(manifestPath)
	_, custom := specField(spec, "networkPolicy").(map[string]interface{})
	return &AgentSandboxTemplateVerifier{
		Name: name, Namespace: namespace, ManifestPath: manifestPath,
		Managed:     stringOr(specField(spec, "networkPolicyManagement"), "Managed") == "Managed",
		CustomRules: custom,
	}
}

func (v *AgentSandboxTemplateVerifier) policyName() string {
	return v.Name + "-network-policy"
}

func (v *AgentSandboxTemplateVerifier) VerifyExists(ctx context.Context, kubeconfig string) error {
	fmt.Printf("  [verify] SandboxTemplate %q in %q (managed=%t, custom rules=%t)\n", v.Name, v.Namespace, v.Managed, v.CustomRules)
	if err := KubectlResourceExists(ctx, kubeconfig, agentSandboxTemplateResource, v.Name, v.Namespace); err != nil {
		return err
	}
	if err := proveProjectionReadback(ctx, kubeconfig, agentSandboxTemplateResource, v.Name, v.Namespace, v.ManifestPath); err != nil {
		return err
	}
	if !v.Managed {
		// Every reconcile stamps the template with its ref-hash label before
		// it decides whether to manage a policy, so the label is the evidence
		// the controller has seen this template and chosen not to create one.
		if _, err := awaitJSONPath(ctx, kubeconfig, agentSandboxTemplateResource, v.Name, v.Namespace,
			"{.metadata.labels."+strings.ReplaceAll(agentSandboxTemplateRefHashLabel, ".", "\\.")+"}", 3*time.Minute); err != nil {
			return errors.Wrap(err, "the controller never reconciled the template (no "+agentSandboxTemplateRefHashLabel+" label)")
		}
		if err := KubectlResourceAbsent(ctx, kubeconfig, "networkpolicy", v.policyName(), v.Namespace); err != nil {
			return errors.Wrap(err, "the template is Unmanaged, but the controller created its NetworkPolicy")
		}
		fmt.Printf("  [verify] UNMANAGED: the controller reconciled the template and created no NetworkPolicy for it\n")
		return nil
	}
	policy, err := awaitJSONPath(ctx, kubeconfig, "networkpolicy", v.policyName(), v.Namespace, "{.spec}", 3*time.Minute)
	if err != nil {
		return errors.Wrap(err, "the template is Managed, but the controller never created its NetworkPolicy")
	}
	var live map[string]interface{}
	if err := json.Unmarshal([]byte(policy), &live); err != nil {
		return errors.Wrap(err, "decoding the template's NetworkPolicy")
	}
	if !v.CustomRules {
		ingress, _ := live["ingress"].([]interface{})
		if !strings.Contains(policy, `"app":"sandbox-router"`) || len(ingress) == 0 {
			return errors.Errorf("the template's NetworkPolicy is not the secure default (ingress only from the router): %s", policy)
		}
		fmt.Printf("  [verify] MANAGED: the controller applied the secure-default NetworkPolicy (ingress only from the router)\n")
		return nil
	}
	want, err := expectedProjection(v.ManifestPath)
	if err != nil {
		return err
	}
	encoded, err := json.Marshal(want.Spec["networkPolicy"])
	if err != nil {
		return err
	}
	var declared map[string]interface{}
	if err := json.Unmarshal(encoded, &declared); err != nil {
		return err
	}
	var missing []string
	for _, direction := range []string{"ingress", "egress"} {
		if rules, ok := declared[direction]; ok {
			missing = append(missing, missingFromLive(rules, live[direction], "networkPolicy."+direction)...)
		}
	}
	if len(missing) > 0 {
		return errors.Errorf("the template's NetworkPolicy does not carry its declared rules:\n  %s", joinLines(missing))
	}
	fmt.Printf("  [verify] MANAGED: the controller's NetworkPolicy %q carries exactly the declared rules\n", v.policyName())
	return nil
}

func (v *AgentSandboxTemplateVerifier) VerifyAbsent(ctx context.Context, kubeconfig string) error {
	if err := KubectlResourceAbsent(ctx, kubeconfig, agentSandboxTemplateResource, v.Name, v.Namespace); err != nil {
		return err
	}
	return errors.Wrap(KubectlResourceAbsent(ctx, kubeconfig, "networkpolicy", v.policyName(), v.Namespace),
		"the template is gone but its NetworkPolicy remains")
}

// AgentSandboxWarmPoolVerifier proves a SandboxWarmPool fills to its size
// with ready members and serves the scale endpoint an autoscaler drives, and
// that destroy takes its unclaimed members with it.
type AgentSandboxWarmPoolVerifier struct {
	Name, Namespace, ManifestPath string
	// Replicas is the declared size, or the cluster's default of one.
	Replicas int64
}

func newAgentSandboxWarmPoolVerifier(manifestPath, name, namespace string) *AgentSandboxWarmPoolVerifier {
	spec := manifestSpecMap(manifestPath)
	return &AgentSandboxWarmPoolVerifier{
		Name: name, Namespace: namespace, ManifestPath: manifestPath,
		Replicas: intOr(specField(spec, "replicas"), 1),
	}
}

func (v *AgentSandboxWarmPoolVerifier) VerifyExists(ctx context.Context, kubeconfig string) error {
	fmt.Printf("  [verify] SandboxWarmPool %q in %q (replicas=%d)\n", v.Name, v.Namespace, v.Replicas)
	if err := KubectlResourceExists(ctx, kubeconfig, agentSandboxWarmPoolResource, v.Name, v.Namespace); err != nil {
		return err
	}
	if err := proveProjectionReadback(ctx, kubeconfig, agentSandboxWarmPoolResource, v.Name, v.Namespace, v.ManifestPath); err != nil {
		return err
	}
	want := fmt.Sprintf("%d", v.Replicas)
	if err := pollJSONPath(ctx, kubeconfig, agentSandboxWarmPoolResource, v.Name, v.Namespace, "{.status.readyReplicas}", want, 6*time.Minute); err != nil {
		return errors.Wrapf(err, "the pool never filled to %s ready members", want)
	}
	fmt.Printf("  [verify] FILL: the pool holds %s ready members\n", want)

	raw, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "get", "--raw",
		fmt.Sprintf("/apis/extensions.agents.x-k8s.io/v1beta1/namespaces/%s/sandboxwarmpools/%s/scale", v.Namespace, v.Name)).CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "reading the pool's scale subresource: %s", firstLines(string(raw), 3))
	}
	var scale struct {
		Spec   struct{ Replicas int64 } `json:"spec"`
		Status struct{ Replicas int64 } `json:"status"`
	}
	if err := json.Unmarshal(raw, &scale); err != nil {
		return errors.Wrap(err, "decoding the pool's scale subresource")
	}
	if scale.Spec.Replicas != v.Replicas || scale.Status.Replicas != v.Replicas {
		return errors.Errorf("the pool's scale endpoint reports spec %d and status %d, want %d -- an autoscaler would read the wrong size", scale.Spec.Replicas, scale.Status.Replicas, v.Replicas)
	}
	fmt.Printf("  [verify] SCALE: the pool's scale endpoint reports %d, the endpoint an autoscaler drives\n", v.Replicas)
	return nil
}

func (v *AgentSandboxWarmPoolVerifier) VerifyAbsent(ctx context.Context, kubeconfig string) error {
	if err := KubectlResourceAbsent(ctx, kubeconfig, agentSandboxWarmPoolResource, v.Name, v.Namespace); err != nil {
		return err
	}
	return errors.Wrap(awaitNoneNamed(ctx, kubeconfig, agentSandboxResource, v.Namespace, v.Name+"-", 3*time.Minute),
		"the pool is gone but its unclaimed members remain")
}

// AgentSandboxClaimVerifier proves a SandboxClaim binds the way it should: to
// a ready pool member (warm, named after the pool) when it asks for nothing a
// running pod cannot take, to a sandbox started for it (cold, named after the
// claim) when it sets variables or volumes, and not at all, with the
// template's refusal reason, when the template disallows what it asks.
type AgentSandboxClaimVerifier struct {
	Name, Namespace, ManifestPath string
	Pool                          string
	Cold                          bool
	// RefusedReason is the Ready=False reason the refusal lane expects.
	RefusedReason string
	PodLabels     map[string]string
	EnvNames      []string
	// Claims are the claim's own volume claim templates, which replace the
	// template's.
	Claims []agentSandboxVolumeClaim
}

func newAgentSandboxClaimVerifier(manifestPath, name, namespace string) *AgentSandboxClaimVerifier {
	spec := manifestSpecMap(manifestPath)
	v := &AgentSandboxClaimVerifier{Name: name, Namespace: namespace, ManifestPath: manifestPath, PodLabels: map[string]string{}}
	if ref, ok := specField(spec, "warmPoolRef").(map[string]interface{}); ok {
		switch n := specField(ref, "name").(type) {
		case map[string]interface{}:
			v.Pool = stringOr(specField(n, "value"), "")
			if v.Pool == "" {
				if from, ok := specField(n, "valueFrom").(map[string]interface{}); ok {
					v.Pool = stringOr(specField(from, "name"), "")
				}
			}
		case string:
			v.Pool = n
		}
	}
	env, _ := specField(spec, "env").([]interface{})
	for _, e := range env {
		em, _ := e.(map[string]interface{})
		if n, ok := specField(em, "name").(string); ok {
			v.EnvNames = append(v.EnvNames, n)
		}
	}
	v.Claims = readVolumeClaimTemplates(spec)
	v.Cold = len(env) > 0 || len(v.Claims) > 0
	if meta, ok := specField(spec, "additionalPodMetadata").(map[string]interface{}); ok {
		if labels, ok := specField(meta, "labels").(map[string]interface{}); ok {
			for k, val := range labels {
				v.PodLabels[k], _ = val.(string)
			}
		}
	}
	if strings.Contains(manifestPath, "behavioral-policy-refusal") {
		v.RefusedReason = "EnvVarsInjectionRejected"
	}
	return v
}

func (v *AgentSandboxClaimVerifier) VerifyExists(ctx context.Context, kubeconfig string) error {
	fmt.Printf("  [verify] SandboxClaim %q in %q from pool %q (cold=%t, refused=%t)\n", v.Name, v.Namespace, v.Pool, v.Cold, v.RefusedReason != "")
	if err := KubectlResourceExists(ctx, kubeconfig, agentSandboxClaimResource, v.Name, v.Namespace); err != nil {
		return err
	}
	if err := proveProjectionReadback(ctx, kubeconfig, agentSandboxClaimResource, v.Name, v.Namespace, v.ManifestPath); err != nil {
		return err
	}
	if v.RefusedReason != "" {
		if err := pollJSONPath(ctx, kubeconfig, agentSandboxClaimResource, v.Name, v.Namespace,
			`{.status.conditions[?(@.type=="Ready")].reason}`, v.RefusedReason, 3*time.Minute); err != nil {
			return errors.Wrapf(err, "the claim was not refused with %s", v.RefusedReason)
		}
		fmt.Printf("  [verify] REFUSED: the claim reports Ready=False with reason %s, as its template's policy requires\n", v.RefusedReason)
		return nil
	}
	if err := kubectlWait(ctx, kubeconfig, agentSandboxClaimResource, v.Name, v.Namespace, "condition=Ready", 6*time.Minute); err != nil {
		return errors.Wrap(err, "the claim never became Ready -- the controller did not bind it to a sandbox")
	}
	sandbox, err := proveClaimBinding(ctx, kubeconfig, v.Name, v.Namespace, v.Pool, v.Cold, v.PodLabels)
	if err != nil {
		return err
	}
	if len(v.EnvNames) > 0 {
		env, err := kubectlGetJSONPath(ctx, kubeconfig, "pod", sandbox, v.Namespace, "{.spec.containers[*].env[*].name}")
		if err != nil {
			return errors.Wrap(err, "reading the claimed sandbox's environment")
		}
		for _, name := range v.EnvNames {
			if !strings.Contains(" "+env+" ", " "+name+" ") {
				return errors.Errorf("the claimed sandbox's pod lacks the claim's variable %s (has: %s)", name, env)
			}
		}
	}
	if len(v.PodLabels) > 0 || len(v.EnvNames) > 0 {
		fmt.Printf("  [verify] BIND: the sandbox's pod carries the claim's labels and variables\n")
	}
	return proveVolumeClaims(ctx, kubeconfig, v.Namespace, sandbox, v.Claims, true)
}

// proveClaimBinding proves a Ready claim holds the sandbox it should: a ready
// member of its pool (warm, named after the pool) when it asks for nothing a
// running pod cannot take, or one started for it (cold, named after the
// claim); and that the sandbox's pod carries the claim's pod labels. It
// returns the claimed sandbox's name.
func proveClaimBinding(ctx context.Context, kubeconfig, claim, namespace, pool string, cold bool, podLabels map[string]string) (string, error) {
	sandbox, err := kubectlGetJSONPath(ctx, kubeconfig, agentSandboxClaimResource, claim, namespace, "{.status.sandbox.name}")
	if err != nil || strings.TrimSpace(sandbox) == "" {
		return "", errors.Wrap(err, "the claim is Ready but names no sandbox")
	}
	sandbox = strings.TrimSpace(sandbox)
	launch, err := kubectlGetJSONPath(ctx, kubeconfig, agentSandboxResource, sandbox, namespace,
		"{.metadata.labels.agents\\.x-k8s\\.io/launch-type}")
	if err != nil {
		return "", errors.Wrapf(err, "reading the claimed sandbox %q", sandbox)
	}
	launch = strings.TrimSpace(launch)
	switch {
	case cold && (launch != "cold" || sandbox != claim):
		return "", errors.Errorf("the claim sets variables or volumes, so it must start a sandbox of its own (cold, named %q); it got %q, launched %q", claim, sandbox, launch)
	case !cold && (launch != "warm" || !strings.HasPrefix(sandbox, pool+"-")):
		return "", errors.Errorf("the claim asks for nothing a running pod cannot take, so it must bind a ready member of pool %q (warm, named %s-...); it got %q, launched %q -- it started cold, the pool did not hand over a member in time", pool, pool, sandbox, launch)
	}
	fmt.Printf("  [verify] BIND: claim %q holds sandbox %q, launched %s\n", claim, sandbox, launch)

	for key, val := range podLabels {
		got, err := kubectlGetJSONPath(ctx, kubeconfig, "pod", sandbox, namespace,
			"{.metadata.labels."+strings.ReplaceAll(key, ".", "\\.")+"}")
		if err != nil || strings.TrimSpace(got) != val {
			return "", errors.Errorf("the claimed sandbox's pod lacks the claim's label %s=%s (got %q)", key, val, got)
		}
	}
	return sandbox, nil
}

func (v *AgentSandboxClaimVerifier) VerifyAbsent(ctx context.Context, kubeconfig string) error {
	if err := KubectlResourceAbsent(ctx, kubeconfig, agentSandboxClaimResource, v.Name, v.Namespace); err != nil {
		return err
	}
	// A claimed sandbox belongs to its claim; with the claim gone, no pod in
	// the namespace may still carry a claim's mark.
	return errors.Wrap(awaitNoneSelected(ctx, kubeconfig, "pod", v.Namespace, "agents.x-k8s.io/claim-uid", 3*time.Minute),
		"the claim is gone but its sandbox's pod remains")
}

// awaitNoneNamed waits until no object of kind in namespace has a name
// starting with prefix.
func awaitNoneNamed(ctx context.Context, kubeconfig, kind, namespace, prefix string, budget time.Duration) error {
	deadline := time.Now().Add(budget)
	for {
		out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "get", kind, "-n", namespace,
			"-o", "jsonpath={.items[*].metadata.name}").CombinedOutput()
		left := []string{}
		if err == nil {
			for _, name := range strings.Fields(string(out)) {
				if strings.HasPrefix(name, prefix) {
					left = append(left, name)
				}
			}
			if len(left) == 0 {
				return nil
			}
		}
		if time.Now().After(deadline) {
			return errors.Errorf("%s still present in %s after %s: %v (%v)", kind, namespace, budget, left, err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}

// awaitNoneSelected waits until no object of kind in namespace matches the
// label selector.
func awaitNoneSelected(ctx context.Context, kubeconfig, kind, namespace, selector string, budget time.Duration) error {
	deadline := time.Now().Add(budget)
	for {
		out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "get", kind, "-n", namespace,
			"-l", selector, "-o", "jsonpath={.items[*].metadata.name}").CombinedOutput()
		if err == nil && strings.TrimSpace(string(out)) == "" {
			return nil
		}
		if time.Now().After(deadline) {
			return errors.Errorf("%s matching %s still present in %s after %s: %s", kind, selector, namespace, budget, firstLines(string(out), 2))
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}

// awaitJSONPath reads a jsonpath of an object until it exists and is not
// empty, within budget.
func awaitJSONPath(ctx context.Context, kubeconfig, kind, name, namespace, jsonPath string, budget time.Duration) (string, error) {
	deadline := time.Now().Add(budget)
	var lastErr error
	for {
		out, err := kubectlGetJSONPath(ctx, kubeconfig, kind, name, namespace, jsonPath)
		if err == nil && strings.TrimSpace(out) != "" {
			return out, nil
		}
		lastErr = err
		if time.Now().After(deadline) {
			return "", errors.Wrapf(lastErr, "%s %s/%s %s stayed empty for %s", kind, namespace, name, jsonPath, budget)
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(3 * time.Second):
		}
	}
}

// kubectlExec runs a command in a pod's first container.
func kubectlExec(ctx context.Context, kubeconfig, namespace, pod string, command ...string) error {
	args := append([]string{"--kubeconfig", kubeconfig, "exec", pod, "-n", namespace, "--"}, command...)
	if out, err := exec.CommandContext(ctx, "kubectl", args...).CombinedOutput(); err != nil {
		return errors.Wrapf(err, "%s", firstLines(string(out), 3))
	}
	return nil
}
