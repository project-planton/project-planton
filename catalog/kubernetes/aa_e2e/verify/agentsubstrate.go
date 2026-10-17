package verify

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/plantonhq/planton/catalog/kubernetes/aa_e2e/verify/internal/ateapipb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"sigs.k8s.io/yaml"
)

// Agent Substrate's fixed names and pinned release. The module installs
// upstream's own release, so these are upstream's names; the release mirrors
// the module's pin (the definition stamps and the version output carry it).
const (
	agentSubstrateKind               = "KubernetesAgentSubstrate"
	agentSubstrateRelease            = "v0.5.0"
	agentSubstrateBundle             = "agent-substrate"
	agentSubstrateNamespace          = "ate-system"
	agentSubstratePodCertNamespace   = "podcertificate-controller-system"
	agentSubstrateAPIServer          = "ate-api-server"
	agentSubstrateCredentialProvider = "k8s-credential-provider"
	agentSubstrateNodeAgent          = "atelet"
	agentSubstratePostgres           = "postgres"
	agentSubstrateVolumeClaim        = "data-postgres-0"
	agentSubstrateDefaultSandbox     = "gvisor-default"
	agentSubstratePolicyConfigMap    = "k8s-credential-provider-namespace-policy"
	agentSubstrateVersionAnnotation  = "planton.ai/agent-substrate-version"
	agentSubstrateMinorLabel         = "planton.ai/agent-substrate-minor"
	agentSubstrateConfigHash         = "planton.ai/config-hash"

	// The probe the verifier creates to call the API as a product would: a
	// namespace carrying the API-access label with two service accounts (a
	// client, named a bootstrap owner where a lane turns authorization on,
	// and a stranger who is not), a namespace without the label, and a
	// small client pod in each for the in-cluster reachability checks.
	agentSubstrateProbeNamespace    = "e2e-ate-probe"
	agentSubstrateOutsiderNamespace = "e2e-ate-outsider"
	agentSubstrateProbeClient       = "e2e-ate-client"
	agentSubstrateProbeStranger     = "e2e-ate-stranger"
	agentSubstrateProbePod          = "e2e-ate-shell"
	agentSubstrateProbeAtespace     = "e2e-ate-probe"
	agentSubstrateProbeImage        = "busybox:1.36"

	// The upgrade refusal, this kind's own class.
	failureUpgradeRefused = "upgrade-refused"
)

// The workloads every install runs, by namespace.
var agentSubstrateDeployments = [][2]string{
	{agentSubstratePodCertNamespace, "podcertificate-controller"},
	{agentSubstrateNamespace, agentSubstrateAPIServer},
	{agentSubstrateNamespace, "ate-controller"},
	{agentSubstrateNamespace, "atenet-router"},
	{agentSubstrateNamespace, "atenet-egress"},
}

// The pod-certificate controller's three signers, each of which publishes
// its root as a ClusterTrustBundle named "<signer, / as :>:primary-bundle".
var agentSubstrateSigners = []string{
	"servicedns.podcert.ate.dev/identity",
	"podidentity.podcert.ate.dev/identity",
	"postgres.podcert.ate.dev/identity",
}

var agentSubstrateCrds = []string{
	"csidriverconfigs.ate.dev",
	"sandboxconfigs.ate.dev",
	"workerpools.ate.dev",
}

// AgentSubstrateVerifier checks an Agent Substrate install by what it
// promises a product, using only the kind's outputs: a token minted for the
// output audience, trusted through the output bundle, dialed at the output
// address under the output server name, creates, reads and deletes an
// atespace. Around that it asserts every workload rolled out, the
// definitions are stamped, the version stamps are on, the API admits only
// namespaces carrying the output label, the router answers, and a lane's
// declared settings are what runs. Destroy asserts the keep promises: kept
// definitions, and the bundled database's volume while another resource owns
// the namespace.
type AgentSubstrateVerifier struct {
	CreateNamespace    bool
	BundledDatabase    bool
	Agentgateway       bool
	CredentialProvider string
	Grants             []agentSubstrateGrant
	Authorization      bool
	BootstrapOwners    []string
	// EnforcesNetworkPolicy marks a lane whose cluster enforces
	// NetworkPolicy (the pod-certificates kind profile), so an unlabeled
	// namespace must be refused.
	EnforcesNetworkPolicy bool
	// ExpectedVersion is the stamp a refused-upgrade lane planted.
	ExpectedVersion string
	helmCRDLifecycle
}

type agentSubstrateGrant struct {
	Atespace          string   `json:"atespace"`
	AllowedNamespaces []string `json:"allowedNamespaces"`
}

func newAgentSubstrateVerifier(manifestPath string) *AgentSubstrateVerifier {
	spec := manifestSpecMap(manifestPath)
	v := &AgentSubstrateVerifier{CredentialProvider: "disabled"}
	v.CreateNamespace, _ = specField(spec, "createNamespace").(bool)
	if db, ok := specField(spec, "database").(map[string]interface{}); ok {
		_, v.BundledDatabase = specField(db, "bundled").(map[string]interface{})
	}
	if router, ok := specField(spec, "router").(map[string]interface{}); ok {
		v.Agentgateway = stringOr(specField(router, "dataplane"), "envoy") == "agentgateway"
	}
	if egress, ok := specField(spec, "egress").(map[string]interface{}); ok {
		if provider, ok := specField(egress, "credentialProvider").(map[string]interface{}); ok {
			if ks, ok := specField(provider, "kubernetesSecrets").(map[string]interface{}); ok {
				v.CredentialProvider = "kubernetes"
				v.Grants = readAgentSubstrateGrants(ks)
			} else if _, ok := specField(provider, "external").(map[string]interface{}); ok {
				v.CredentialProvider = "external"
			}
		}
	}
	if api, ok := specField(spec, "apiServer").(map[string]interface{}); ok {
		if authz, ok := specField(api, "authorization").(map[string]interface{}); ok {
			v.Authorization = boolOr(specField(authz, "enabled"), true)
			if owners, ok := specField(authz, "bootstrapOwners").([]interface{}); ok {
				for _, owner := range owners {
					if s, ok := owner.(string); ok {
						v.BootstrapOwners = append(v.BootstrapOwners, s)
					}
				}
			}
		}
	}
	v.EnforcesNetworkPolicy = manifestAnnotation(manifestPath, "planton.dev/e2e-cluster-profile") == "pod-certificates"
	v.ExpectedVersion = manifestAnnotation(manifestPath, "planton.dev/e2e-expect-installed-version")

	crds := agentSubstrateCrds
	if expected := manifestAnnotation(manifestPath, "planton.dev/e2e-expect-crds"); expected != "" {
		crds = strings.Split(expected, ",")
	}
	v.helmCRDLifecycle = readHelmCRDLifecycle(manifestPath, agentSubstrateRelease, agentSubstrateBundle, crds)
	return v
}

func readAgentSubstrateGrants(provider map[string]interface{}) []agentSubstrateGrant {
	var grants []agentSubstrateGrant
	list, _ := specField(provider, "grants").([]interface{})
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		g := agentSubstrateGrant{Atespace: stringOr(specField(m, "atespace"), "")}
		namespaces, _ := specField(m, "allowedNamespaces").([]interface{})
		for _, ns := range namespaces {
			if s, ok := ns.(string); ok {
				g.AllowedNamespaces = append(g.AllowedNamespaces, s)
			}
		}
		grants = append(grants, g)
	}
	return grants
}

// probeOwner is the principal the probe client calls the API as.
func (v *AgentSubstrateVerifier) probeOwner() string {
	return "system:serviceaccount:" + agentSubstrateProbeNamespace + ":" + agentSubstrateProbeClient
}

// VerifyExists cannot prove the promise without the outputs; the harness
// calls VerifyExistsWithOutputs.
func (v *AgentSubstrateVerifier) VerifyExists(ctx context.Context, kubeconfig string) error {
	return errors.New("the Agent Substrate verifier calls the API through the kind's outputs, and none were passed")
}

// VerifyExistsWithOutputs runs every check against the installed control
// plane, reading the API's coordinates from the outputs alone.
func (v *AgentSubstrateVerifier) VerifyExistsWithOutputs(ctx context.Context, kubeconfig string, outputs map[string]interface{}) error {
	out := func(key string) string { return fmt.Sprint(outputs[key]) }
	fmt.Printf("  [verify] Agent Substrate %s in %s (bundled database=%t, agentgateway=%t, credential provider=%s, authorization=%t)\n",
		out("version"), out("namespace"), v.BundledDatabase, v.Agentgateway, v.CredentialProvider, v.Authorization)
	if out("version") != agentSubstrateRelease || out("namespace") != agentSubstrateNamespace {
		return errors.Errorf("the outputs say %s in %s; the kind installs %s in %s", out("version"), out("namespace"), agentSubstrateRelease, agentSubstrateNamespace)
	}

	if err := v.verifyRollouts(ctx, kubeconfig); err != nil {
		return err
	}
	if err := v.verifyEstablishedAndStamped(ctx, kubeconfig); err != nil {
		return err
	}
	if err := v.verifyTrustBundles(ctx, kubeconfig); err != nil {
		return err
	}
	if err := KubectlResourceExists(ctx, kubeconfig, "sandboxconfigs.ate.dev", agentSubstrateDefaultSandbox, ""); err != nil {
		return errors.Wrap(err, "the default SandboxConfig every worker pool names was not applied")
	}
	if err := v.verifyStamps(ctx, kubeconfig); err != nil {
		return err
	}
	if err := v.verifyNamespaceOwner(ctx, kubeconfig); err != nil {
		return err
	}
	if err := v.verifyDataplane(ctx, kubeconfig); err != nil {
		return err
	}
	if v.CredentialProvider == "kubernetes" {
		if err := v.verifyGrants(ctx, kubeconfig); err != nil {
			return err
		}
	}

	if err := v.applyProbe(ctx, kubeconfig, out("api_access_namespace_label")); err != nil {
		return err
	}
	if err := v.verifyAPIRoundTrip(ctx, kubeconfig, outputs); err != nil {
		return err
	}
	if err := v.verifyReachability(ctx, kubeconfig, out("api_address")); err != nil {
		return err
	}
	return v.verifyRouter(ctx, kubeconfig, out("router_endpoint"))
}

func (v *AgentSubstrateVerifier) verifyRollouts(ctx context.Context, kubeconfig string) error {
	workloads := make([][3]string, 0, 8)
	for _, d := range agentSubstrateDeployments {
		workloads = append(workloads, [3]string{"deployment", d[0], d[1]})
	}
	workloads = append(workloads, [3]string{"daemonset", agentSubstrateNamespace, agentSubstrateNodeAgent})
	if v.CredentialProvider == "kubernetes" {
		workloads = append(workloads, [3]string{"deployment", agentSubstrateNamespace, agentSubstrateCredentialProvider})
	}
	if v.BundledDatabase {
		workloads = append(workloads, [3]string{"statefulset", agentSubstrateNamespace, agentSubstratePostgres})
	}
	for _, w := range workloads {
		out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "rollout", "status",
			w[0]+"/"+w[2], "-n", w[1], "--timeout=8m").CombinedOutput()
		if err != nil {
			return errors.Wrapf(err, "%s %s/%s never finished rolling out: %s", w[0], w[1], w[2], firstLines(string(out), 4))
		}
	}
	fmt.Printf("  [verify] every workload rolled out (%d)\n", len(workloads))
	return nil
}

// verifyTrustBundles checks every signer published its root: the bundles
// every Substrate client and the database trust.
func (v *AgentSubstrateVerifier) verifyTrustBundles(ctx context.Context, kubeconfig string) error {
	for _, signer := range agentSubstrateSigners {
		name := strings.ReplaceAll(signer, "/", ":") + ":primary-bundle"
		deadline := time.Now().Add(3 * time.Minute)
		for {
			got, err := kubectlGetJSONPath(ctx, kubeconfig, "clustertrustbundle", name, "", "{.spec.signerName}")
			if err == nil && strings.TrimSpace(got) == signer {
				break
			}
			if time.Now().After(deadline) {
				return errors.Errorf("ClusterTrustBundle %s was never published for signer %s (last: %q, %v)", name, signer, strings.TrimSpace(got), err)
			}
			time.Sleep(5 * time.Second)
		}
	}
	fmt.Printf("  [verify] the three signers published their trust bundles\n")
	return nil
}

// verifyStamps checks the marks the upgrade refusal reads: the installed
// release on the API server, and the minor on the bundled database's volume.
func (v *AgentSubstrateVerifier) verifyStamps(ctx context.Context, kubeconfig string) error {
	stamp, err := kubectlGetJSONPath(ctx, kubeconfig, "deployment", agentSubstrateAPIServer, agentSubstrateNamespace,
		"{.metadata.annotations.planton\\.ai/agent-substrate-version}")
	if err != nil {
		return errors.Wrap(err, "reading the API server's release stamp")
	}
	if strings.TrimSpace(stamp) != agentSubstrateRelease {
		return errors.Errorf("the API server carries %s=%q, want %s", agentSubstrateVersionAnnotation, strings.TrimSpace(stamp), agentSubstrateRelease)
	}
	if v.BundledDatabase {
		minor, err := kubectlGetJSONPath(ctx, kubeconfig, "pvc", agentSubstrateVolumeClaim, agentSubstrateNamespace,
			"{.metadata.labels.planton\\.ai/agent-substrate-minor}")
		if err != nil {
			return errors.Wrap(err, "reading the database volume's minor stamp")
		}
		if want := strings.TrimPrefix(agentSubstrateRelease[:strings.LastIndex(agentSubstrateRelease, ".")], "v"); strings.TrimSpace(minor) != want {
			return errors.Errorf("the database volume carries %s=%q, want %s", agentSubstrateMinorLabel, strings.TrimSpace(minor), want)
		}
	}
	fmt.Printf("  [verify] the release stamps are on (API server, and the database volume where bundled)\n")
	return nil
}

func (v *AgentSubstrateVerifier) verifyNamespaceOwner(ctx context.Context, kubeconfig string) error {
	owner, err := kubectlGetJSONPath(ctx, kubeconfig, "namespace", agentSubstrateNamespace, "",
		"{.metadata.labels.planton\\.ai/kind}")
	if err != nil {
		return errors.Wrapf(err, "reading namespace %s", agentSubstrateNamespace)
	}
	owner = strings.TrimSpace(owner)
	if v.CreateNamespace && owner != agentSubstrateKind {
		return errors.Errorf("namespace %s carries planton.ai/kind=%q, but this resource created it", agentSubstrateNamespace, owner)
	}
	if !v.CreateNamespace && owner == agentSubstrateKind {
		return errors.Errorf("namespace %s carries this kind's mark, but another resource owns it", agentSubstrateNamespace)
	}
	return nil
}

// verifyDataplane checks the router and the egress gateway run the declared
// proxy: agentgateway replaces Envoy in both.
func (v *AgentSubstrateVerifier) verifyDataplane(ctx context.Context, kubeconfig string) error {
	for _, name := range []string{"atenet-router", "atenet-egress"} {
		containers, err := kubectlGetJSONPath(ctx, kubeconfig, "deployment", name, agentSubstrateNamespace,
			"{.spec.template.spec.containers[*].name}")
		if err != nil {
			return errors.Wrapf(err, "reading %s's containers", name)
		}
		has := strings.Contains(" "+containers+" ", " agentgateway ")
		if has != v.Agentgateway {
			return errors.Errorf("%s runs containers [%s], but the manifest sets dataplane agentgateway=%t", name, strings.TrimSpace(containers), v.Agentgateway)
		}
	}
	return nil
}

// verifyGrants checks the credential provider runs the declared grants: its
// policy holds exactly them, and the running pod started from that policy
// (its configuration digest equals the Deployment's).
func (v *AgentSubstrateVerifier) verifyGrants(ctx context.Context, kubeconfig string) error {
	policy, err := kubectlGetJSONPath(ctx, kubeconfig, "configmap", agentSubstratePolicyConfigMap, agentSubstrateNamespace,
		"{.data.namespace-policy\\.yaml}")
	if err != nil {
		return errors.Wrap(err, "reading the credential provider's namespace policy")
	}
	var live struct {
		Policies []agentSubstrateGrant `json:"policies"`
	}
	if err := yaml.Unmarshal([]byte(policy), &live); err != nil {
		return errors.Wrapf(err, "the credential provider's namespace policy is not YAML: %s", firstLines(policy, 6))
	}
	want := v.Grants
	if want == nil {
		want = []agentSubstrateGrant{}
	}
	if live.Policies == nil {
		live.Policies = []agentSubstrateGrant{}
	}
	if !reflect.DeepEqual(live.Policies, want) {
		return errors.Errorf("the credential provider's policy grants %+v; the manifest declares %+v", live.Policies, want)
	}
	declared, err := kubectlGetJSONPath(ctx, kubeconfig, "deployment", agentSubstrateCredentialProvider, agentSubstrateNamespace,
		"{.spec.template.metadata.annotations.planton\\.ai/config-hash}")
	if err != nil {
		return errors.Wrap(err, "reading the credential provider's configuration digest")
	}
	declared = strings.TrimSpace(declared)
	if declared == "" {
		return errors.Errorf("the credential provider's pod template carries no %s, so a policy change would not restart it", agentSubstrateConfigHash)
	}
	// Pods of the previous configuration may still be terminating just after
	// the rollout, so the running set is polled until it settles.
	deadline := time.Now().Add(2 * time.Minute)
	for {
		running, err := kubectlGetJSONPathList(ctx, kubeconfig, "pod", agentSubstrateNamespace, "app="+agentSubstrateCredentialProvider,
			`{range .items[*]}{.metadata.annotations.planton\.ai/config-hash}{"\n"}{end}`)
		if err != nil {
			return errors.Wrap(err, "reading the running credential provider's configuration digest")
		}
		stale := ""
		for _, digest := range running {
			if digest != declared {
				stale = digest
			}
		}
		if stale == "" && len(running) > 0 {
			break
		}
		if time.Now().After(deadline) {
			return errors.Errorf("a running credential provider started from configuration %q, but the declared one is %s -- the policy change did not roll it", stale, declared)
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("  [verify] the credential provider runs the declared grants (%d)\n", len(want))
	return nil
}

func (v *AgentSubstrateVerifier) applyProbe(ctx context.Context, kubeconfig, accessLabel string) error {
	key, value, ok := strings.Cut(accessLabel, "=")
	if !ok {
		return errors.Errorf("the api_access_namespace_label output %q is not key=value", accessLabel)
	}
	shell := func(namespace string) string {
		return fmt.Sprintf(`---
apiVersion: v1
kind: Pod
metadata:
  name: %s
  namespace: %s
spec:
  containers:
    - name: shell
      image: %s
      command: ["sleep", "3600"]
`, agentSubstrateProbePod, namespace, agentSubstrateProbeImage)
	}
	manifest := fmt.Sprintf(`apiVersion: v1
kind: Namespace
metadata:
  name: %[1]s
  labels:
    %[2]q: %[3]q
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: %[4]s
  namespace: %[1]s
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: %[5]s
  namespace: %[1]s
---
apiVersion: v1
kind: Namespace
metadata:
  name: %[6]s
`, agentSubstrateProbeNamespace, key, value, agentSubstrateProbeClient, agentSubstrateProbeStranger, agentSubstrateOutsiderNamespace)
	manifest += shell(agentSubstrateProbeNamespace) + shell(agentSubstrateOutsiderNamespace)
	if err := kubectlApplyStdin(ctx, kubeconfig, manifest); err != nil {
		return errors.Wrap(err, "creating the API probe")
	}
	for _, namespace := range []string{agentSubstrateProbeNamespace, agentSubstrateOutsiderNamespace} {
		if err := kubectlWait(ctx, kubeconfig, "pod", agentSubstrateProbePod, namespace, "condition=Ready", 3*time.Minute); err != nil {
			return errors.Wrapf(err, "the probe shell in %s never became Ready", namespace)
		}
	}
	return nil
}

// verifyAPIRoundTrip calls the API exactly as a product does, from the
// outputs alone.
func (v *AgentSubstrateVerifier) verifyAPIRoundTrip(ctx context.Context, kubeconfig string, outputs map[string]interface{}) error {
	address := fmt.Sprint(outputs["api_address"])
	serverName := fmt.Sprint(outputs["api_server_name"])
	audience := fmt.Sprint(outputs["api_token_audience"])
	bundle := fmt.Sprint(outputs["api_trust_bundle_name"])

	pem, err := kubectlGetJSONPath(ctx, kubeconfig, "clustertrustbundle", bundle, "", "{.spec.trustBundle}")
	if err != nil {
		return errors.Wrapf(err, "reading the API's trust bundle %s (the api_trust_bundle_name output)", bundle)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM([]byte(pem)) {
		return errors.Errorf("ClusterTrustBundle %s holds no certificate", bundle)
	}

	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return errors.Wrapf(err, "the api_address output %q is not host:port", address)
	}
	labels := strings.Split(host, ".")
	if len(labels) < 2 {
		return errors.Errorf("the api_address output %q does not name a Service", address)
	}
	localPort, err := freeLocalPort()
	if err != nil {
		return err
	}
	stop, err := startPortForward(ctx, kubeconfig, "svc/"+labels[0], labels[1], fmt.Sprintf("%d:%s", localPort, port))
	if err != nil {
		return errors.Wrap(err, "forwarding a local port to the API")
	}
	defer stop()

	conn, err := grpc.NewClient(fmt.Sprintf("127.0.0.1:%d", localPort),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: roots, ServerName: serverName, MinVersion: tls.VersionTLS12})))
	if err != nil {
		return errors.Wrap(err, "dialing the API")
	}
	defer func() { _ = conn.Close() }()
	api := ateapipb.NewControlClient(conn)

	tokenless, cancel := context.WithTimeout(ctx, 30*time.Second)
	_, err = api.GetAtespace(tokenless, &ateapipb.GetAtespaceRequest{Atespace: &ateapipb.ObjectRef{Name: agentSubstrateProbeAtespace}})
	cancel()
	if status.Code(err) != codes.Unauthenticated {
		return errors.Errorf("a call without a token returned %v; the API must refuse it as Unauthenticated", err)
	}

	token, err := mintToken(ctx, kubeconfig, agentSubstrateProbeClient, audience)
	if err != nil {
		return err
	}
	call := func(token string) (context.Context, context.CancelFunc) {
		c, cancel := context.WithTimeout(ctx, 30*time.Second)
		return withBearer(c, token), cancel
	}

	if v.Authorization {
		stranger, err := mintToken(ctx, kubeconfig, agentSubstrateProbeStranger, audience)
		if err != nil {
			return err
		}
		c, cancel := call(stranger)
		_, err = api.CreateAtespace(c, &ateapipb.CreateAtespaceRequest{Atespace: &ateapipb.Atespace{Metadata: &ateapipb.ResourceMetadata{Name: agentSubstrateProbeAtespace}}})
		cancel()
		strangerID := "system:serviceaccount:" + agentSubstrateProbeNamespace + ":" + agentSubstrateProbeStranger
		if status.Code(err) != codes.PermissionDenied || !strings.Contains(err.Error(), strangerID) {
			return errors.Errorf("with authorization on, a caller who is not an owner (%s) got %v; the API must deny it by that principal", strangerID, err)
		}
		fmt.Printf("  [verify] with authorization on, a non-owner was denied by its principal %s\n", strangerID)
	}

	deadline := time.Now().Add(3 * time.Minute)
	for {
		c, cancel := call(token)
		_, err = api.CreateAtespace(c, &ateapipb.CreateAtespaceRequest{Atespace: &ateapipb.Atespace{Metadata: &ateapipb.ResourceMetadata{Name: agentSubstrateProbeAtespace}}})
		cancel()
		if err == nil || status.Code(err) == codes.AlreadyExists {
			break
		}
		if time.Now().After(deadline) {
			return errors.Wrapf(err, "creating atespace %s as %s through %s (server name %s, token audience %s)", agentSubstrateProbeAtespace, v.probeOwner(), address, serverName, audience)
		}
		time.Sleep(5 * time.Second)
	}
	c, cancel := call(token)
	got, err := api.GetAtespace(c, &ateapipb.GetAtespaceRequest{Atespace: &ateapipb.ObjectRef{Name: agentSubstrateProbeAtespace}})
	cancel()
	if err != nil {
		return errors.Wrapf(err, "reading back atespace %s", agentSubstrateProbeAtespace)
	}
	if got.GetMetadata().GetName() != agentSubstrateProbeAtespace || got.GetMetadata().GetUid() == "" {
		return errors.Errorf("atespace %s read back as %v", agentSubstrateProbeAtespace, got.GetMetadata())
	}
	c, cancel = call(token)
	_, err = api.DeleteAtespace(c, &ateapipb.DeleteAtespaceRequest{Atespace: &ateapipb.ObjectRef{Name: agentSubstrateProbeAtespace}})
	cancel()
	if err != nil {
		return errors.Wrapf(err, "deleting atespace %s", agentSubstrateProbeAtespace)
	}
	fmt.Printf("  [verify] the API created, read and deleted an atespace for %s, trusted through %s as %s\n", v.probeOwner(), bundle, serverName)
	return nil
}

// verifyReachability checks the API's network policy from inside the
// cluster: a namespace carrying the access label connects; on a cluster that
// enforces NetworkPolicy, one without it does not.
func (v *AgentSubstrateVerifier) verifyReachability(ctx context.Context, kubeconfig, address string) error {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return errors.Wrapf(err, "the api_address output %q is not host:port", address)
	}
	connect := func(namespace string) error {
		out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "exec", agentSubstrateProbePod, "-n", namespace, "--",
			"nc", "-z", "-w", "5", host, port).CombinedOutput()
		if err != nil {
			return errors.Errorf("%v: %s", err, firstLines(string(out), 2))
		}
		return nil
	}
	var last error
	for attempt := 0; attempt < 12; attempt++ {
		if last = connect(agentSubstrateProbeNamespace); last == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if last != nil {
		return errors.Wrapf(last, "a pod in a namespace carrying the access label could not reach %s", address)
	}
	if v.EnforcesNetworkPolicy {
		if err := connect(agentSubstrateOutsiderNamespace); err == nil {
			return errors.Errorf("a pod in a namespace without the access label reached %s; the API's network policy must refuse it", address)
		}
		fmt.Printf("  [verify] the API admits the labeled namespace and refuses an unlabeled one\n")
		return nil
	}
	fmt.Printf("  [verify] the API admits the labeled namespace\n")
	return nil
}

// verifyRouter checks the router answers HTTP at its output endpoint. With no
// actor behind the request, any HTTP status proves it serves.
func (v *AgentSubstrateVerifier) verifyRouter(ctx context.Context, kubeconfig, endpoint string) error {
	deadline := time.Now().Add(2 * time.Minute)
	var last string
	for time.Now().Before(deadline) {
		out, _ := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "exec", agentSubstrateProbePod, "-n", agentSubstrateProbeNamespace, "--",
			"wget", "-S", "-q", "-O", "/dev/null", "-T", "5", endpoint+"/").CombinedOutput()
		last = string(out)
		if strings.Contains(last, "HTTP/") {
			fmt.Printf("  [verify] the router answered HTTP at %s\n", endpoint)
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return errors.Errorf("the router never answered HTTP at %s (the router_endpoint output); last: %s", endpoint, firstLines(last, 3))
}

func (v *AgentSubstrateVerifier) VerifyAbsent(ctx context.Context, kubeconfig string) error {
	for _, d := range agentSubstrateDeployments {
		if err := KubectlResourceAbsent(ctx, kubeconfig, "deployment", d[1], d[0]); err != nil {
			return err
		}
	}
	if err := KubectlResourceAbsent(ctx, kubeconfig, "namespace", agentSubstratePodCertNamespace, ""); err != nil {
		return errors.Wrap(err, "this resource always creates the pod-certificate namespace, so destroy must delete it")
	}
	if !v.DeployRefused {
		if v.CreateNamespace {
			if err := KubectlResourceAbsent(ctx, kubeconfig, "namespace", agentSubstrateNamespace, ""); err != nil {
				return errors.Wrap(err, "this resource created the namespace, so destroy must delete it")
			}
		} else {
			if err := KubectlResourceExists(ctx, kubeconfig, "namespace", agentSubstrateNamespace, ""); err != nil {
				return errors.Wrap(err, "another resource owns the namespace, so destroying the control plane must leave it")
			}
			if v.BundledDatabase {
				if err := KubectlResourceExists(ctx, kubeconfig, "pvc", agentSubstrateVolumeClaim, agentSubstrateNamespace); err != nil {
					return errors.Wrap(err, "the bundled database's volume must outlive the destroy while another resource owns the namespace")
				}
				fmt.Printf("  [verify] DESTROY: the bundled database's volume outlived the control plane\n")
			}
		}
		if v.InstallCrds && v.KeepOnUninstall {
			if err := KubectlResourceAbsent(ctx, kubeconfig, "sandboxconfigs.ate.dev", agentSubstrateDefaultSandbox, ""); err != nil {
				return errors.Wrap(err, "destroy removes the default SandboxConfig even when the definitions are kept")
			}
		}
	}
	if err := v.verifyDestroyPosture(ctx, kubeconfig); err != nil {
		return err
	}
	for _, namespace := range []string{agentSubstrateProbeNamespace, agentSubstrateOutsiderNamespace} {
		if out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "delete", "namespace",
			namespace, "--ignore-not-found", "--wait=true", "--timeout=3m").CombinedOutput(); err != nil {
			return errors.Wrapf(err, "removing the probe namespace %s: %s", namespace, firstLines(string(out), 3))
		}
	}
	return nil
}

// VerifyExpectedDeployFailure pins a refused deploy to its three-part text:
// this kind's refusals here, the definition primitive's in helmCRDLifecycle.
func (v *AgentSubstrateVerifier) VerifyExpectedDeployFailure(ctx context.Context, kubeconfig, expectation string, deployErr error) error {
	text := strings.Join(strings.Fields(strings.ReplaceAll(deployErr.Error(), "|", " ")), " ")
	wants, err := agentSubstrateRefusalTexts(expectation, v.ExpectedVersion)
	if err != nil {
		return v.verifyRefusal(ctx, kubeconfig, expectation, deployErr)
	}
	for _, part := range []string{"observed:", "meaning:", "next step:"} {
		if !strings.Contains(text, part) {
			return errors.Errorf("the refusal lacks its %q part; got: %s", part, firstLines(text, 12))
		}
	}
	for _, want := range wants {
		if !strings.Contains(text, want) {
			return errors.Errorf("expected the %s refusal saying %q; got: %s", expectation, want, firstLines(text, 12))
		}
	}
	// Nothing was taken over or rewritten.
	switch expectation {
	case failureWorkloadOwnedElsewhere:
		owner, err := kubectlGetJSONPath(ctx, kubeconfig, "deployment", agentSubstrateAPIServer, agentSubstrateNamespace,
			"{.metadata.labels.planton\\.ai/kind}")
		if err != nil {
			return errors.Wrap(err, "reading the other installation's API server after the refusal")
		}
		if strings.TrimSpace(owner) == agentSubstrateKind {
			return errors.New("the other installation's API server gained this kind's mark despite the refusal")
		}
	case failureUpgradeRefused:
		if manifestVolumeLane(v) {
			break
		}
		stamp, err := kubectlGetJSONPath(ctx, kubeconfig, "deployment", agentSubstrateAPIServer, agentSubstrateNamespace,
			"{.metadata.annotations.planton\\.ai/agent-substrate-version}")
		if err != nil {
			return errors.Wrap(err, "reading the earlier installation's API server after the refusal")
		}
		if strings.TrimSpace(stamp) != v.ExpectedVersion {
			return errors.Errorf("the API server's release stamp became %q despite the refusal; it must still read %s", strings.TrimSpace(stamp), v.ExpectedVersion)
		}
	}
	fmt.Printf("  [verify] REFUSED as expected (%s), with what was observed, what it means and the next step\n", expectation)
	return nil
}

// manifestVolumeLane reports a refused upgrade read from the database volume
// rather than the API server (the planted stamp is a minor, "v0.4").
func manifestVolumeLane(v *AgentSubstrateVerifier) bool {
	return strings.Count(v.ExpectedVersion, ".") == 1
}

// agentSubstrateRefusalTexts are the phrases each of this kind's refusals
// must carry, in the words the modules write.
func agentSubstrateRefusalTexts(expectation, installed string) ([]string, error) {
	switch expectation {
	case failureWorkloadOwnedElsewhere:
		return []string{
			"Deployment " + agentSubstrateNamespace + "/" + agentSubstrateAPIServer + " already runs on the cluster and was not created by this " + agentSubstrateKind,
			"another installation of the same software runs here",
			"label deployment " + agentSubstrateAPIServer,
		}, nil
	case failureNamespaceMissing:
		return []string{
			"namespace " + agentSubstrateNamespace + " does not exist on the cluster, and this " + agentSubstrateKind + " does not create it",
			"its owner must create it first",
			"set spec.create_namespace to true",
		}, nil
	case failureNamespaceOwnedElsewhere:
		return []string{
			"namespace " + agentSubstrateNamespace + " already exists on the cluster and was not created by this " + agentSubstrateKind,
			"would take the namespace over",
			"set spec.create_namespace to false",
		}, nil
	case failureUpgradeRefused:
		if installed == "" {
			return nil, errors.New("an upgrade-refused lane names the planted version in planton.dev/e2e-expect-installed-version")
		}
		record := "Deployment " + agentSubstrateNamespace + "/" + agentSubstrateAPIServer
		remedy := "destroy this resource"
		if strings.Count(installed, ".") == 1 {
			record = "PersistentVolumeClaim " + agentSubstrateNamespace + "/" + agentSubstrateVolumeClaim
			remedy = "delete the volume (`kubectl -n " + agentSubstrateNamespace + " delete pvc " + agentSubstrateVolumeClaim + "`)"
		}
		return []string{
			record,
			"records Agent Substrate " + installed + " on this cluster, and this " + agentSubstrateKind + " installs " + agentSubstrateRelease,
			"Agent Substrate does not move from " + installed + " to " + agentSubstrateRelease + " in place: before 1.0 every minor release is a reinstall",
			remedy,
		}, nil
	}
	return nil, errors.Errorf("%s is not one of this kind's refusals", expectation)
}

func mintToken(ctx context.Context, kubeconfig, serviceAccount, audience string) (string, error) {
	out, err := exec.CommandContext(ctx, "kubectl", "--kubeconfig", kubeconfig, "create", "token", serviceAccount,
		"-n", agentSubstrateProbeNamespace, "--audience", audience, "--duration=10m").CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "minting a token for %s with audience %s: %s", serviceAccount, audience, firstLines(string(out), 2))
	}
	return strings.TrimSpace(string(out)), nil
}

func freeLocalPort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, errors.Wrap(err, "finding a free local port")
	}
	defer func() { _ = l.Close() }()
	_, port, _ := net.SplitHostPort(l.Addr().String())
	return strconv.Atoi(port)
}

// withBearer attaches a token the way upstream's clients do.
func withBearer(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
}
