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
)

// AgentSandboxVerifier proves a Sandbox: its fields read back, it becomes
// Ready when Running, and destroy removes it with its pod and claims. The
// suspend lane (two acts) proves the promise behind operatingMode: suspending
// deletes the pod and keeps the same claim, and resuming brings the files
// back.
type AgentSandboxVerifier struct {
	Name, Namespace, ManifestPath string
	Suspended                     bool
	// SuspendLane is the behavioral-suspend scenario: act one writes the
	// marker, act two (Suspended) proves it survives.
	SuspendLane bool
	// ClaimTemplates are the volume claim template names; each claim is
	// `<template>-<sandbox>`.
	ClaimTemplates []string
}

func newAgentSandboxVerifier(manifestPath, name, namespace string) *AgentSandboxVerifier {
	spec := manifestSpecMap(manifestPath)
	v := &AgentSandboxVerifier{
		Name: name, Namespace: namespace, ManifestPath: manifestPath,
		Suspended:   stringOr(specField(spec, "operatingMode"), "Running") == "Suspended",
		SuspendLane: strings.Contains(manifestPath, "behavioral-suspend"),
	}
	templates, _ := specField(spec, "volumeClaimTemplates").([]interface{})
	for _, t := range templates {
		tm, _ := t.(map[string]interface{})
		meta, _ := specField(tm, "metadata").(map[string]interface{})
		if n, ok := specField(meta, "name").(string); ok {
			v.ClaimTemplates = append(v.ClaimTemplates, n)
		}
	}
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
	for _, template := range v.ClaimTemplates {
		if err := KubectlResourceExists(ctx, kubeconfig, "pvc", v.claimName(template), v.Namespace); err != nil {
			return errors.Wrapf(err, "the controller did not create the claim for volume claim template %q", template)
		}
	}
	if v.Suspended {
		return v.proveSuspended(ctx, kubeconfig)
	}
	if err := kubectlWait(ctx, kubeconfig, agentSandboxResource, v.Name, v.Namespace, "condition=Ready", 5*time.Minute); err != nil {
		return errors.Wrap(err, "the Sandbox never became Ready -- the controller did not bring its pod up")
	}
	fmt.Printf("  [verify] the Sandbox is Ready: its pod %q runs\n", v.Name)
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
	for _, template := range v.ClaimTemplates {
		if err := KubectlResourceAbsent(ctx, kubeconfig, "pvc", v.claimName(template), v.Namespace); err != nil {
			return errors.Wrapf(err, "the Sandbox is gone but its claim for %q remains", template)
		}
	}
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
		// Give the controller a reconcile to create one, then require none.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(20 * time.Second):
		}
		if err := KubectlResourceAbsent(ctx, kubeconfig, "networkpolicy", v.policyName(), v.Namespace); err != nil {
			return errors.Wrap(err, "the template is Unmanaged, but the controller created its NetworkPolicy")
		}
		fmt.Printf("  [verify] UNMANAGED: the controller created no NetworkPolicy for the template\n")
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
