package verify

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/plantonhq/planton/pkg/failure"
)

// The Agent Substrate lanes are only as honest as the verifier's reading of
// their scenarios: who owns the namespace, which database and dataplane run,
// which grants the credential provider must hold, who the bootstrap owner
// is. Each committed scenario is read here offline, so a scenario edit that
// changes what a lane asserts shows up as a failing test, not as a lane that
// quietly proves something else. The refusal phrases the lanes look for are
// held to the texts the modules actually write, and the release to both
// modules' pin.

const agentSubstrateScenarios = "../../kubernetesagentsubstrate/e2e/scenarios"

func TestAgentSubstrateVerifierReadsTheScenarios(t *testing.T) {
	research := []agentSubstrateGrant{{Atespace: "research", AllowedNamespaces: []string{"research-credentials"}}}
	changed := []agentSubstrateGrant{
		{Atespace: "research", AllowedNamespaces: []string{"research-credentials", "shared-credentials"}},
		{Atespace: "support", AllowedNamespaces: []string{"support-credentials"}},
	}
	type lane struct {
		create, bundled, agentgateway, authz, keep, refused, policy bool
		provider, installed                                         string
		grants                                                      []agentSubstrateGrant
		crds                                                        []string
	}
	kind := lane{bundled: true, policy: true, provider: "disabled", crds: agentSubstrateCrds}
	with := func(edit func(*lane)) lane { l := kind; edit(&l); return l }
	cases := map[string]lane{
		"minimal.yaml":                     with(func(l *lane) { l.keep = true }),
		"agentgateway.yaml":                with(func(l *lane) { l.agentgateway = true }),
		"full-surface.yaml":                with(func(l *lane) { l.authz = true }),
		"credential-policy-change.yaml":    with(func(l *lane) { l.provider, l.grants = "kubernetes", research }),
		"credential-policy-change.to.yaml": with(func(l *lane) { l.provider, l.grants = "kubernetes", changed }),
		"external-database.yaml":           with(func(l *lane) { l.bundled = false }),
		"gke-keyless.yaml":                 with(func(l *lane) { l.create, l.policy = true, false }),
		"namespace-missing.yaml":           with(func(l *lane) { l.keep, l.refused = true, true }),
		"namespace-owned-elsewhere.yaml":   with(func(l *lane) { l.create, l.keep, l.refused = true, true, true }),
		"workload-owned-elsewhere.yaml":    with(func(l *lane) { l.keep, l.refused = true, true }),
		"crds-owned-elsewhere.yaml": with(func(l *lane) {
			l.create, l.refused, l.crds = true, true, []string{"workerpools.ate.dev"}
		}),
		"upgrade-refused.yaml":        with(func(l *lane) { l.keep, l.refused, l.installed = true, true, "v0.4.0" }),
		"volume-upgrade-refused.yaml": with(func(l *lane) { l.keep, l.refused, l.installed = true, true, "v0.4" }),
	}
	files, err := filepath.Glob(filepath.Join(agentSubstrateScenarios, "*.yaml"))
	if err != nil || len(files) != len(cases) {
		t.Fatalf("every scenario needs a case here: found %d files, %d cases (%v)", len(files), len(cases), err)
	}
	for name, want := range cases {
		t.Run(name, func(t *testing.T) {
			v := newAgentSubstrateVerifier(filepath.Join(agentSubstrateScenarios, name))
			got := lane{
				create: v.CreateNamespace, bundled: v.BundledDatabase, agentgateway: v.Agentgateway, authz: v.Authorization,
				keep: v.KeepOnUninstall, refused: v.DeployRefused, policy: v.EnforcesNetworkPolicy,
				provider: v.CredentialProvider, installed: v.ExpectedVersion, grants: v.Grants, crds: v.CRDs,
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("read %+v, want %+v", got, want)
			}
			if v.ChartVersion != agentSubstrateRelease || v.SourceLabel != agentSubstrateBundle {
				t.Errorf("stamps version=%q source=%q, want %s %s", v.ChartVersion, v.SourceLabel, agentSubstrateRelease, agentSubstrateBundle)
			}
		})
	}
}

// The authorization lane proves the owner promise only if the owner it
// names is the principal the probe calls as.
func TestAgentSubstrateFullSurfaceOwnsTheProbe(t *testing.T) {
	v := newAgentSubstrateVerifier(filepath.Join(agentSubstrateScenarios, "full-surface.yaml"))
	if !reflect.DeepEqual(v.BootstrapOwners, []string{v.probeOwner()}) {
		t.Errorf("full-surface names bootstrap owners %v; the probe calls as %s", v.BootstrapOwners, v.probeOwner())
	}
}

// Each refusal phrase the lanes look for must appear in the text the
// modules write: the shared failure constructors render it, and the upgrade
// rule and remedies are read from both engines' sources.
func TestAgentSubstrateRefusalTextsMatchTheModules(t *testing.T) {
	read := func(path string) string {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return string(src)
	}
	goSource := read("../../kubernetesagentsubstrate/iac/pulumi/module/upgrade.go")
	tfSource := read("../../kubernetesagentsubstrate/iac/tf/preflight.tf")
	rule := regexp.MustCompile(`ruleBeforeOne = "([^"]+)"`).FindStringSubmatch(goSource)
	if rule == nil {
		t.Fatal("upgrade.go no longer names ruleBeforeOne")
	}
	for _, phrase := range []string{"destroy this resource", "delete the volume (`kubectl -n %s delete pvc %s`)"} {
		if !strings.Contains(goSource, phrase) {
			t.Errorf("upgrade.go no longer says %q", phrase)
		}
	}
	if !strings.Contains(tfSource, "delete the volume (`kubectl -n ${local.namespace} delete pvc ${local.postgres_volume_claim}`)") || !strings.Contains(tfSource, "destroy this resource") {
		t.Error("preflight.tf no longer names the remedies the lanes look for")
	}

	rendered := map[string]string{
		failureWorkloadOwnedElsewhere:  failure.KubernetesWorkloadOwnedElsewhere(agentSubstrateKind, "Deployment", agentSubstrateNamespace, agentSubstrateAPIServer, "no owner label").Error(),
		failureNamespaceMissing:        failure.KubernetesNamespaceMissing(agentSubstrateKind, agentSubstrateNamespace).Error(),
		failureNamespaceOwnedElsewhere: failure.KubernetesNamespaceOwnedElsewhere(agentSubstrateKind, agentSubstrateNamespace, "no owner label").Error(),
	}
	for expectation, text := range rendered {
		wants, err := agentSubstrateRefusalTexts(expectation, "")
		if err != nil {
			t.Fatal(err)
		}
		for _, want := range wants {
			if !strings.Contains(text, want) {
				t.Errorf("%s: the module writes %q, which lacks %q", expectation, text, want)
			}
		}
	}

	for installed, record := range map[string]string{
		"v0.4.0": "Deployment " + agentSubstrateNamespace + "/" + agentSubstrateAPIServer,
		"v0.4":   "PersistentVolumeClaim " + agentSubstrateNamespace + "/" + agentSubstrateVolumeClaim + ", the bundled database's volume left by an earlier installation,",
	} {
		remedy := "back up what you need from the database, destroy this resource"
		if installed == "v0.4" {
			remedy = "delete the volume (`kubectl -n " + agentSubstrateNamespace + " delete pvc " + agentSubstrateVolumeClaim + "`)"
		}
		text := failure.KubernetesUpgradeRefused(agentSubstrateKind, "Agent Substrate", record, installed, agentSubstrateRelease, rule[1], remedy).Error()
		wants, err := agentSubstrateRefusalTexts(failureUpgradeRefused, installed)
		if err != nil {
			t.Fatal(err)
		}
		for _, want := range wants {
			if !strings.Contains(text, want) {
				t.Errorf("upgrade from %s: the module writes %q, which lacks %q", installed, text, want)
			}
		}
	}
}

func TestAgentSubstrateVerifierPinsTheModulesRelease(t *testing.T) {
	for path, pattern := range map[string]*regexp.Regexp{
		"../../kubernetesagentsubstrate/iac/pulumi/module/vars.go": regexp.MustCompile(`Release:\s+"([^"]+)"`),
		"../../kubernetesagentsubstrate/iac/tf/locals.tf":          regexp.MustCompile(`release\s*=\s*"([^"]+)"`),
		"internal/ateapipb/upstream_test.go":                       regexp.MustCompile(`upstreamRelease = "([^"]+)"`),
	} {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		match := pattern.FindSubmatch(src)
		if match == nil || string(match[1]) != agentSubstrateRelease {
			t.Errorf("%s pins %q; the verifier expects %s", path, match, agentSubstrateRelease)
		}
	}
}
