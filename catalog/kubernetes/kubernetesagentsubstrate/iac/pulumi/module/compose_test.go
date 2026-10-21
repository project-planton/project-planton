package module

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	kubernetesprovider "github.com/plantonhq/planton/catalog/kubernetes"
	v1 "github.com/plantonhq/planton/catalog/kubernetes/kubernetesagentsubstrate/v1alpha1"
	"github.com/plantonhq/planton/pkg/failure"
	"github.com/plantonhq/planton/shared"
	foreignkeyv1 "github.com/plantonhq/planton/shared/foreignkey/v1"
)

// The composition is pinned here against the real v0.5.0 files Planton
// publishes (testdata/v0.5.0, held equal to the published bytes by a gated
// test): an empty settings section installs upstream as published; every
// typed setting lands where upstream's installer puts it; upstream's own
// components (agentgateway, cordon, size10, the Cloud SQL sidecar) and text
// markers apply as upstream applies them; the rollout hashes cover exactly
// what each workload reads; the upgrade rule follows upstream's policy; and
// the generated pools have upstream's wire shape. The live lanes prove the
// behavior on a cluster.

func literal(v string) *foreignkeyv1.StringValueOrRef {
	return &foreignkeyv1.StringValueOrRef{LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: v}}
}

func i32(v int32) *int32 { return &v }

func fixtures(t *testing.T) releaseFileSet {
	t.Helper()
	root := filepath.Join("testdata", vars.Release)
	set := releaseFileSet{files: map[string]string{}}
	for _, path := range releaseFiles {
		raw, err := os.ReadFile(filepath.Join(root, path))
		if err != nil {
			t.Fatalf("fixture %s: %v", path, err)
		}
		set.files[path] = string(raw)
	}
	raw, err := os.ReadFile(filepath.Join(root, "setup.sql"))
	if err != nil {
		t.Fatal(err)
	}
	set.setupSQL = string(raw)
	return set
}

func minimalSpec() *v1.KubernetesAgentSubstrateSpec {
	return &v1.KubernetesAgentSubstrateSpec{
		Namespace: literal("ate-system"),
		Platform: &v1.KubernetesAgentSubstratePlatform{Kind: &v1.KubernetesAgentSubstratePlatform_Generic{Generic: &v1.KubernetesAgentSubstrateGeneric{
			ServiceAccountIssuer: "https://kubernetes.default.svc.cluster.local",
		}}},
		Database: &v1.KubernetesAgentSubstrateDatabase{Kind: &v1.KubernetesAgentSubstrateDatabase_Bundled{Bundled: &v1.KubernetesAgentSubstrateBundledDatabase{}}},
		Storage: &v1.KubernetesAgentSubstrateStorage{Kind: &v1.KubernetesAgentSubstrateStorage_S3{S3: &v1.KubernetesAgentSubstrateS3Storage{
			Endpoint:    literal("http://main-s3.seaweedfs.svc.cluster.local:8333"),
			Bucket:      "snapshots",
			Credentials: &v1.KubernetesAgentSubstrateS3Credentials{SecretName: literal("substrate-s3")},
		}}},
		Egress: &v1.KubernetesAgentSubstrateEgress{CredentialProvider: &v1.KubernetesAgentSubstrateCredentialProvider{
			Kind: &v1.KubernetesAgentSubstrateCredentialProvider_Disabled{Disabled: &v1.KubernetesAgentSubstrateCredentialProviderDisabled{}},
		}},
	}
}

func gkeSpec() *v1.KubernetesAgentSubstrateSpec {
	spec := minimalSpec()
	spec.Platform = &v1.KubernetesAgentSubstratePlatform{Kind: &v1.KubernetesAgentSubstratePlatform_Gke{Gke: &v1.KubernetesAgentSubstrateGke{
		ClusterId:               literal("projects/acme/locations/us-central1/clusters/agents"),
		ApiServerServiceAccount: literal("ate-api-server@acme.iam.gserviceaccount.com"),
		NodeAgentServiceAccount: literal("atelet@acme.iam.gserviceaccount.com"),
	}}}
	spec.Database = &v1.KubernetesAgentSubstrateDatabase{Kind: &v1.KubernetesAgentSubstrateDatabase_CloudSql{CloudSql: &v1.KubernetesAgentSubstrateCloudSqlDatabase{
		InstanceConnectionName: literal("acme:us-central1:substrate"),
	}}}
	spec.Storage = &v1.KubernetesAgentSubstrateStorage{Kind: &v1.KubernetesAgentSubstrateStorage_Gcs{Gcs: &v1.KubernetesAgentSubstrateGcsStorage{Bucket: literal("acme-snapshots")}}}
	return spec
}

func composeSpec(t *testing.T, spec *v1.KubernetesAgentSubstrateSpec) (*settings, *composition) {
	t.Helper()
	s := resolveSettings(spec)
	s.identityLabels = identityLabels(&shared.CatalogObjectMetadata{Name: "agent-substrate"})
	c, err := compose(s, fixtures(t))
	if err != nil {
		t.Fatalf("compose: %v", err)
	}
	return s, c
}

func find(c *composition, kind, name string) document {
	for _, d := range c.all() {
		if str(d, "kind") == kind && str(d, "metadata", "name") == name {
			return d
		}
	}
	return nil
}

func upstreamDocument(t *testing.T, path, kind, name string) document {
	t.Helper()
	docs, err := splitDocuments(fixtures(t).files[path])
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range docs {
		if str(d, "kind") == kind && str(d, "metadata", "name") == name {
			return d
		}
	}
	t.Fatalf("%s has no %s %s", path, kind, name)
	return nil
}

func args(t *testing.T, d document, container string, init bool) []string {
	t.Helper()
	c, err := containerNamed(podSpec(d), container, init)
	if err != nil {
		t.Fatal(err)
	}
	return stringList(c["args"])
}

func envOf(t *testing.T, d document, container string, init bool) map[string]interface{} {
	t.Helper()
	c, err := containerNamed(podSpec(d), container, init)
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]interface{}{}
	for _, e := range c["env"].([]interface{}) {
		m := e.(map[string]interface{})
		if v, ok := m["value"]; ok {
			out[m["name"].(string)] = v
		} else {
			out[m["name"].(string)] = m["valueFrom"]
		}
	}
	return out
}

func TestEmptySectionsInstallUpstreamAsPublished(t *testing.T) {
	_, c := composeSpec(t, minimalSpec())
	for _, w := range []struct{ path, kind, name, container string }{
		{fileAPIServer, "Deployment", vars.APIServerName, vars.APIServerName},
		{fileController, "Deployment", vars.ControllerName, vars.ControllerName},
		{fileRouter, "Deployment", vars.RouterName, vars.RouterName},
		{fileEgress, "Deployment", vars.EgressName, "ext-proc"},
		{filePodCertificates, "Deployment", vars.PodCertificatesName, "controller"},
	} {
		got := args(t, find(c, w.kind, w.name), w.container, false)
		want := args(t, upstreamDocument(t, w.path, w.kind, w.name), w.container, false)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s args moved off upstream's with nothing set:\n got %v\nwant %v", w.name, got, want)
		}
	}
	if got := find(c, "Deployment", vars.APIServerName)["spec"].(map[string]interface{})["replicas"]; got != int64(2) {
		t.Errorf("API server replicas = %v, upstream's is 2", got)
	}
}

func TestEveryWorkloadCarriesTheOwnershipLabel(t *testing.T) {
	_, c := composeSpec(t, minimalSpec())
	seen := 0
	for _, d := range c.all() {
		if !isWorkload(d) {
			continue
		}
		seen++
		if str(d, "metadata", "labels", failure.WorkloadOwnerLabel) != kindName {
			t.Errorf("%s lacks %s=%s, which the ownership check reads", documentID(d), failure.WorkloadOwnerLabel, kindName)
		}
	}
	if seen != 7 {
		t.Errorf("composed %d workloads; the bundled minimal install runs 7 (pod certificates, postgres, API server, controller, node agent, router, egress)", seen)
	}
	for _, w := range ownedWorkloads(resolveSettings(minimalSpec())) {
		if find(c, string(w.Kind), w.Name) == nil {
			t.Errorf("the ownership check reads %s %s, which the composition does not apply", w.Kind, w.Name)
		}
	}
}

func TestNamespacesAndPhases(t *testing.T) {
	_, c := composeSpec(t, minimalSpec())
	if len(c.namespaces) != 1 || str(c.namespaces[0], "metadata", "name") != vars.PodCertificatesNamespace {
		t.Fatalf("without create_namespace only the pod-certificate namespace is this resource's; got %d namespaces", len(c.namespaces))
	}
	if str(c.namespaces[0], "metadata", "labels", failure.WorkloadOwnerLabel) != kindName {
		t.Error("the pod-certificate namespace must carry the ownership mark the namespace check reads")
	}
	spec := minimalSpec()
	spec.CreateNamespace = true
	_, owned := composeSpec(t, spec)
	if len(owned.namespaces) != 2 {
		t.Fatalf("with create_namespace both namespaces are this resource's; got %d", len(owned.namespaces))
	}
	if len(c.runtime) != 1 || str(c.runtime[0], "kind") != "SandboxConfig" {
		t.Errorf("the runtime phase is the default gVisor SandboxConfig alone, applied last so it is deleted first")
	}
	for _, d := range c.all() {
		if str(d, "kind") == "Namespace" && str(d, "metadata", "name") == vars.Namespace {
			t.Error("ate-system is applied only with create_namespace (ate-controller.yaml repeats it)")
		}
	}
}

func TestGenericPlatformDropsTheGkeOnlyPieces(t *testing.T) {
	_, c := composeSpec(t, minimalSpec())
	if find(c, "PodMonitoring", "atenet-router-envoy") != nil {
		t.Error("PodMonitoring is GKE's managed Prometheus; a generic cluster has no such definition")
	}
	nodeAgent := find(c, "DaemonSet", vars.NodeAgentName)
	for _, a := range args(t, nodeAgent, vars.NodeAgentName, false) {
		if strings.HasPrefix(a, "--image-credential-provider-") {
			t.Errorf("generic clusters run no node image credential provider, yet %s remains", a)
		}
	}
	if namedItem(podSpec(nodeAgent)["volumes"], "name", "image-credential-provider-bin") != nil {
		t.Error("the GKE host mount of the image credential provider must go off GKE")
	}
	otel := find(c, "ConfigMap", configMapOtel)["data"].(map[string]interface{})
	if _, ok := otel["OTEL_EXPORTER_OTLP_ENDPOINT"]; ok || otel["OTEL_TRACES_EXPORTER"] != "none" || otel["OTEL_METRICS_EXPORTER"] != "prometheus" {
		t.Errorf("a generic cluster with no endpoint pushes nothing and keeps the Prometheus scrape; got %v", otel)
	}
	auth := find(c, "ConfigMap", configMapAuthentication)["data"].(map[string]interface{})["authentication.yaml"]
	want := "actorIdentityJWTProvider: kubernetes\njwtProviders:\n- name: kubernetes\n  issuer: https://kubernetes.default.svc.cluster.local\n  audiences: [api.ate-system.svc]\n  certificateAuthorityFile: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\n  discoveryTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token"
	if auth != want {
		t.Errorf("authentication.yaml differs from upstream's for an in-cluster issuer:\n%v", auth)
	}
}

func TestGkeKeepsItsPiecesAndWiresCloudSQLAndGcs(t *testing.T) {
	s, c := composeSpec(t, gkeSpec())
	if find(c, "PodMonitoring", "atenet-router-envoy") == nil {
		t.Error("GKE keeps the router's PodMonitoring")
	}
	if str(find(c, "ServiceAccount", vars.APIServerName), "metadata", "annotations", "iam.gke.io/gcp-service-account") != "ate-api-server@acme.iam.gserviceaccount.com" ||
		str(find(c, "ServiceAccount", vars.NodeAgentName), "metadata", "annotations", "iam.gke.io/gcp-service-account") != "atelet@acme.iam.gserviceaccount.com" {
		t.Error("each Kubernetes service account carries its Google service account")
	}
	api := find(c, "Deployment", vars.APIServerName)
	if _, err := containerNamed(podSpec(api), "cloud-sql-proxy", true); err != nil {
		t.Errorf("Cloud SQL adds upstream's proxy sidecar: %v", err)
	}
	if got := envOf(t, api, "objectstore-plugin", true)["ATE_STORAGE_BACKEND"]; got != "gcs" {
		t.Errorf("GCS keeps upstream's backend, got %v", got)
	}
	env := apiEnvConfigMap(s)
	if env["ATE_API_POSTGRES_CLOUDSQL_INSTANCE"] != "acme:us-central1:substrate" || env["CSQL_PROXY_PRIVATE_IP"] != "true" || env["CSQL_PROXY_AUTO_IAM_AUTHN"] != "true" {
		t.Errorf("the proxy's settings are upstream's: %v", env)
	}
	if got := apiEnvSecret(s)["ATE_API_POSTGRES_READ_WRITE_CONNECTION_STRING"]; got != "user=ate-api-server@acme.iam host=127.0.0.1 port=5432 dbname=atepg sslmode=disable" {
		t.Errorf("the IAM login DSN is upstream's cloudSQLDSN: %q", got)
	}
	if !strings.Contains(authenticationYAML(s), "issuer: https://container.googleapis.com/v1/projects/acme/locations/us-central1/clusters/agents\n  audiences: [api.ate-system.svc]") || strings.Contains(authenticationYAML(s), "discoveryTokenFile") {
		t.Errorf("GKE trusts its own issuer with no in-cluster discovery: %s", authenticationYAML(s))
	}
	if find(c, "StatefulSet", vars.PostgresName) != nil {
		t.Error("Cloud SQL installs no bundled database")
	}
}

func TestBundledDatabase(t *testing.T) {
	s, c := composeSpec(t, minimalSpec())
	postgres := find(c, "StatefulSet", vars.PostgresName)
	claim := namedItemInMeta(postgres["spec"].(map[string]interface{})["volumeClaimTemplates"].([]interface{}), "data")
	if str(claim, "metadata", "labels", vars.MinorLabel) != "0.5" {
		t.Error("the volume claim template carries the minor that writes the data, so the claim carries it past a destroy")
	}
	if got := lookup(claim, "spec", "resources", "requests", "storage"); got != "500Gi" {
		t.Errorf("the volume defaults to upstream's 500Gi, got %v", got)
	}
	container, _ := containerNamed(podSpec(postgres), vars.PostgresName, false)
	hook := stringList(lookup(container, "lifecycle", "postStart", "exec", "command").([]interface{}))
	if !strings.HasSuffix(hook[2], setupCommand+"\n") || !strings.Contains(hook[2], "createdb") {
		t.Errorf("the start hook keeps upstream's createdb and runs the setup script after it:\n%s", hook[2])
	}
	if find(c, "ConfigMap", vars.PostgresSetupConfigMap)["data"].(map[string]interface{})["setup.sql"] == "" {
		t.Error("the setup script is upstream's, mounted for the hook")
	}
	secret := apiEnvSecret(s)
	if !strings.HasPrefix(secret["ATE_API_POSTGRES_READ_WRITE_CONNECTION_STRING"], "postgresql://substrate_readwrite_user@postgres.ate-system.svc:5432/atepg?sslmode=verify-full") ||
		secret["ATE_API_POSTGRES_SCHEMA"] != "substrate" {
		t.Errorf("the bundled DSNs are upstream's: %v", secret)
	}
	storage := envOf(t, find(c, "DaemonSet", vars.NodeAgentName), "objectstore-plugin", true)
	if storage["ATE_STORAGE_BACKEND"] != "s3" || storage["AWS_ENDPOINT_URL"] != "http://main-s3.seaweedfs.svc.cluster.local:8333" || storage["AWS_S3_USE_PATH_STYLE"] != "true" {
		t.Errorf("S3 reaches the node agent's sidecar in the AWS SDK's own variables: %v", storage)
	}
	if ref, ok := storage["AWS_SECRET_ACCESS_KEY"].(map[string]interface{}); !ok || lookup(ref, "secretKeyRef", "name") != "substrate-s3" {
		t.Errorf("the S3 secret key is read from the Secret, never written: %v", storage["AWS_SECRET_ACCESS_KEY"])
	}
}

func TestExternalDatabaseReadsYourSecret(t *testing.T) {
	spec := minimalSpec()
	spec.Database = &v1.KubernetesAgentSubstrateDatabase{Kind: &v1.KubernetesAgentSubstrateDatabase_External{External: &v1.KubernetesAgentSubstrateExternalDatabase{
		ReadWriteConnectionString: &v1.KubernetesAgentSubstrateSecretKeyRef{Name: literal("substrate-db"), Key: "rw"},
	}}}
	s, c := composeSpec(t, spec)
	env := envOf(t, find(c, "Deployment", vars.APIServerName), vars.APIServerName, false)
	for _, name := range []string{"ATE_API_POSTGRES_READ_WRITE_CONNECTION_STRING", "ATE_API_POSTGRES_OWNER_CONNECTION_STRING"} {
		if lookup(env[name].(map[string]interface{}), "secretKeyRef", "key") != "rw" {
			t.Errorf("%s comes from your Secret (the owner falls back to read/write)", name)
		}
	}
	if _, ok := apiEnvSecret(s)["ATE_API_POSTGRES_READ_WRITE_CONNECTION_STRING"]; ok {
		t.Error("your connection string never lands in the module's Secret")
	}
}

func TestTypedSettingsLandWhereUpstreamPutsThem(t *testing.T) {
	spec := minimalSpec()
	spec.ApiServer = &v1.KubernetesAgentSubstrateApiServer{
		Replicas: i32(3), LogLevel: v1.LogLevel_debug,
		Authorization: &v1.KubernetesAgentSubstrateAuthorization{BootstrapOwners: []string{"spiffe://cluster.local/ns/ops/sa/admin"}},
	}
	spec.Router = &v1.KubernetesAgentSubstrateRouter{AllowedClientSpiffeIds: []string{"spiffe://cluster.local/ns/app/sa/web"}}
	spec.NodeAgent = &v1.KubernetesAgentSubstrateNodeAgent{AssetStagingCapacityGib: i32(16)}
	spec.PodCertificates = &v1.KubernetesAgentSubstratePodCertificates{WorkersPerSigner: i32(4)}
	_, c := composeSpec(t, spec)
	api := args(t, find(c, "Deployment", vars.APIServerName), vars.APIServerName, false)
	for _, want := range []string{"--log-level=debug", "--experimental-enable-authz=true", "--authz-bootstrap-owners=spiffe://cluster.local/ns/ops/sa/admin"} {
		if !contains(api, want) {
			t.Errorf("API server args lack %s: %v", want, api)
		}
	}
	router := args(t, find(c, "Deployment", vars.RouterName), vars.RouterName, false)
	if !contains(router, "--ingress-allowed-spiffe-ids=spiffe://cluster.local/ns/app/sa/web") || contains(router, "--ingress-allowed-spiffe-ids=spiffe://cluster.local/ns/ate-e2e-router-client/sa/router-client") {
		t.Errorf("the allowed SPIFFE IDs replace upstream's in place: %v", router)
	}
	nodeAgent := find(c, "DaemonSet", vars.NodeAgentName)
	if !contains(args(t, nodeAgent, "objectstore-plugin", true), "--asset-staging-capacity=16Gi") ||
		lookup(namedItem(podSpec(nodeAgent)["volumes"], "name", "asset-staging"), "emptyDir", "sizeLimit") != "17Gi" {
		t.Error("the staging capacity and its volume (capacity + 1Gi) move together")
	}
	if envOf(t, find(c, "Deployment", vars.PodCertificatesName), "controller", false)["WORKERS_PER_SIGNER"] != "4" {
		t.Error("workers per signer reaches the manifest's WORKERS_PER_SIGNER")
	}
}

func TestUpstreamComponentsApplyAsUpstreamAppliesThem(t *testing.T) {
	spec := minimalSpec()
	spec.ClusterSize = v1.KubernetesAgentSubstrateSpec_size10
	spec.CordonControlPlane = true
	s, c := composeSpec(t, spec)
	podcert := args(t, find(c, "Deployment", vars.PodCertificatesName), "controller", false)
	if !contains(podcert, "--kube-api-qps=100") || !contains(podcert, "--kube-api-burst=200") {
		t.Errorf("size10 raises the pod-certificate controller's rate: %v", podcert)
	}
	if !strings.Contains(find(c, "ConfigMap", configMapPostgres)["data"].(map[string]interface{})["postgresql.conf"].(string), "max_connections = 300") {
		t.Error("size10 replaces postgresql.conf with upstream's")
	}
	if !strings.HasSuffix(apiEnvSecret(s)["ATE_API_POSTGRES_READ_WRITE_CONNECTION_STRING"], size10PoolParams) {
		t.Error("size10 widens the read/write pool")
	}
	for _, name := range []string{vars.APIServerName, vars.ControllerName, vars.RouterName, vars.EgressName} {
		if lookup(podSpec(find(c, "Deployment", name)), "nodeSelector", "ate.dev/workloadType") != "ate-control-plane" {
			t.Errorf("cordon pins %s to the control-plane nodes", name)
		}
	}
	if lookup(podSpec(find(c, "StatefulSet", vars.PostgresName)), "nodeSelector", "ate.dev/workloadType") != "ate-postgres" {
		t.Error("cordon pins the database to its own node")
	}
}

func TestAgentgatewayDataplane(t *testing.T) {
	spec := minimalSpec()
	spec.Router = &v1.KubernetesAgentSubstrateRouter{Dataplane: v1.KubernetesAgentSubstrateRouter_agentgateway}
	spec.Egress.CredentialProvider = &v1.KubernetesAgentSubstrateCredentialProvider{Kind: &v1.KubernetesAgentSubstrateCredentialProvider_External{
		External: &v1.KubernetesAgentSubstrateExternalCredentialProvider{Name: "secretmanager.googleapis.com", Address: "gsm.ate-system.svc:50051"},
	}}
	_, c := composeSpec(t, spec)
	router := podSpec(find(c, "Deployment", vars.RouterName))["containers"].([]interface{})
	if len(router) != 1 || router[0].(map[string]interface{})["name"] != "agentgateway" {
		t.Errorf("agentgateway replaces the router binary and Envoy: %v", len(router))
	}
	if find(c, "ConfigMap", "atenet-router-envoy-config") != nil {
		t.Error("upstream's component deletes the router's Envoy configuration")
	}
	egress := podSpec(find(c, "Deployment", vars.EgressName))
	if _, ok := egress["initContainers"]; ok {
		t.Error("agentgateway mints its own leaves, so Envoy's minter goes")
	}
	config := find(c, "ConfigMap", agentgatewayEgressConfig)["data"].(map[string]interface{})["config.yaml"].(string)
	if !strings.Contains(config, `uriAuthority: "secretmanager.googleapis.com"`) || !strings.Contains(config, `hostname: "gsm.ate-system.svc"`) || strings.Contains(config, "#ATE_AGENTGATEWAY") {
		t.Errorf("the credential provider is spliced into the egress policy at upstream's marker:\n%s", config)
	}
}

// The API server rejects a pod template with two volumes, containers or
// mounts of one name, and kustomize's merge never leaves one: a JSON patch
// that appends a volume upstream's manifest already declares (agentgateway's
// servicedns-ca) collapses into it at the next strategic merge.
func TestEveryPodTemplateNamesEachVolumeContainerAndMountOnce(t *testing.T) {
	agentgateway := fullSurfaceSpec()
	agentgateway.Router.Dataplane = v1.KubernetesAgentSubstrateRouter_agentgateway
	agentgateway.Egress.AdditionalExtProcService = nil
	for name, spec := range map[string]*v1.KubernetesAgentSubstrateSpec{
		"minimal": minimalSpec(), "gke": gkeSpec(), "full surface": fullSurfaceSpec(), "agentgateway": agentgateway,
	} {
		_, c := composeSpec(t, spec)
		for _, d := range c.all() {
			if d["kind"] != "Deployment" && d["kind"] != "DaemonSet" && d["kind"] != "StatefulSet" {
				continue
			}
			unique := func(what string, items []interface{}, key string) {
				seen := map[interface{}]bool{}
				for _, item := range items {
					v := item.(map[string]interface{})[key]
					if seen[v] {
						t.Errorf("%s: %s %v names %s %v twice", name, d["kind"], str(d, "metadata", "name"), what, v)
					}
					seen[v] = true
				}
			}
			pod := podSpec(d)
			volumes, _ := pod["volumes"].([]interface{})
			unique("volume", volumes, "name")
			for _, field := range []string{"containers", "initContainers"} {
				containers, _ := pod[field].([]interface{})
				unique("container", containers, "name")
				for _, container := range containers {
					mounts, _ := container.(map[string]interface{})["volumeMounts"].([]interface{})
					unique("mount path", mounts, "mountPath")
				}
			}
		}
	}
}

func TestEnvoyEgressSplicesAndFlags(t *testing.T) {
	spec := minimalSpec()
	spec.Egress = &v1.KubernetesAgentSubstrateEgress{
		CredentialProvider: &v1.KubernetesAgentSubstrateCredentialProvider{Kind: &v1.KubernetesAgentSubstrateCredentialProvider_KubernetesSecrets{
			KubernetesSecrets: &v1.KubernetesAgentSubstrateKubernetesSecretsProvider{Grants: []*v1.KubernetesAgentSubstrateNamespaceGrant{{Atespace: "team-a", AllowedNamespaces: []string{"team-a"}}}},
		}},
		EnvoyConcurrency:         i32(4),
		AdditionalExtProcService: &v1.KubernetesAgentSubstrateExtProcService{Namespace: "policy", Service: "filter", Port: 9443},
	}
	_, c := composeSpec(t, spec)
	egress := find(c, "Deployment", vars.EgressName)
	extProc := args(t, egress, "ext-proc", false)
	for _, want := range []string{"--credential-provider-name=k8s.io", "--credential-provider-address=k8s-credential-provider.ate-system.svc:50051", "--credential-provider-server-name=k8s-credential-provider.ate-system.svc"} {
		if !contains(extProc, want) {
			t.Errorf("ext-proc lacks upstream's inject flag %s", want)
		}
	}
	if envoy := args(t, egress, "envoy", false); envoy[len(envoy)-2] != "--concurrency" || envoy[len(envoy)-1] != "4" {
		t.Errorf("Envoy's concurrency is appended as upstream's patch does: %v", envoy)
	}
	envoyYAML := find(c, "ConfigMap", vars.EgressName)["data"].(map[string]interface{})["envoy.yaml"].(string)
	if strings.Count(envoyYAML, "cluster_name: additional_egress_ext_proc") != 3 || markerLineLeft(envoyYAML, "#ATE_MITM_EXTPROC") || !strings.Contains(envoyYAML, "address: filter.policy.svc.cluster.local") {
		t.Error("the extra filter splices into both intercepting chains and its cluster at upstream's markers")
	}
	policy := find(c, "ConfigMap", configMapPolicy)["data"].(map[string]interface{})["namespace-policy.yaml"]
	if policy != "policies:\n- atespace: \"team-a\"\n  allowedNamespaces:\n  - \"team-a\"\n" {
		t.Errorf("the grants are upstream's policy format: %q", policy)
	}
	if find(c, "Deployment", vars.CredentialProviderName) == nil {
		t.Error("the Kubernetes Secrets provider installs upstream's provider")
	}
}

func TestImagesMoveToTheMirrorKeepingPathAndDigest(t *testing.T) {
	spec := gkeSpec()
	spec.ImageRegistry = "mirror.example.com"
	s, c := composeSpec(t, spec)
	for _, d := range c.all() {
		if !isWorkload(d) {
			continue
		}
		for _, field := range []string{"initContainers", "containers"} {
			items, _ := podSpec(d)[field].([]interface{})
			for _, item := range items {
				if image := item.(map[string]interface{})["image"].(string); !strings.HasPrefix(image, "mirror.example.com/") {
					t.Errorf("%s runs %s off the mirror", documentID(d), image)
				}
			}
		}
	}
	if got := mirrored("postgres:18-alpine@sha256:abc", "m.io"); got != "m.io/library/postgres:18-alpine@sha256:abc" {
		t.Errorf("a bare Docker Hub name reads as docker.io/library: %s", got)
	}
	if got := mirrored("envoyproxy/envoy:v1", "m.io"); got != "m.io/envoyproxy/envoy:v1" {
		t.Errorf("a Docker Hub path keeps its namespace: %s", got)
	}
	out := outputValues(s, c.images)
	if !strings.HasPrefix(out[OpGvisorWorkerImage], "mirror.example.com/plantonhq/planton/agent-substrate/ateom-gvisor:v0.5.0@sha256:") {
		t.Errorf("the worker image output follows the mirror: %s", out[OpGvisorWorkerImage])
	}
}

func TestRolloutHashesCoverWhatEachWorkloadReads(t *testing.T) {
	derived := map[string]bool{configMapAuthentication: true, configMapOtel: true, configMapPolicy: true, configMapPostgres: true, vars.PostgresSetupConfigMap: true}
	spec := minimalSpec()
	spec.Egress.CredentialProvider = &v1.KubernetesAgentSubstrateCredentialProvider{Kind: &v1.KubernetesAgentSubstrateCredentialProvider_KubernetesSecrets{
		KubernetesSecrets: &v1.KubernetesAgentSubstrateKubernetesSecretsProvider{},
	}}
	_, c := composeSpec(t, spec)
	for _, d := range c.all() {
		if !isWorkload(d) {
			continue
		}
		raw, _ := json.Marshal(podSpec(d))
		var read []string
		for name := range derived {
			if strings.Contains(string(raw), `"name":"`+name+`"`) {
				read = append(read, name)
			}
		}
		listed := map[string]bool{}
		for _, name := range configReaders[str(d, "metadata", "name")] {
			listed[name] = true
			if !strings.Contains(string(raw), `"name":"`+name+`"`) {
				t.Errorf("%s is hashed over %s, which it does not read", str(d, "metadata", "name"), name)
			}
		}
		for _, name := range read {
			if !listed[name] {
				t.Errorf("%s reads %s at start but no change to it would roll the workload", str(d, "metadata", "name"), name)
			}
		}
		if len(listed) > 0 && str(d, "spec", "template", "metadata", "annotations", vars.ConfigHashAnnotation) == "" {
			t.Errorf("%s lacks its configuration hash", str(d, "metadata", "name"))
		}
	}
	api := find(c, "Deployment", vars.APIServerName)
	if str(api, "spec", "template", "metadata", "annotations", vars.EnvHashAnnotation) == "" || str(api, "metadata", "annotations", vars.VersionAnnotation) != vars.Release {
		t.Error("the API server carries upstream's env hash and the installed-version stamp")
	}
	sum := sha256.Sum256([]byte("configmap\nA=1\nB=2\nsecret\nC=3\n"))
	if envHash(map[string]string{"B": "2", "A": "1"}, map[string]string{"C": "3"}) != hex.EncodeToString(sum[:]) {
		t.Error("the env hash is upstream's: sha256 of the sorted ConfigMap lines, then the sorted Secret lines")
	}
}

func TestUpgradeRuleFollowsUpstreamsPolicy(t *testing.T) {
	for _, c := range []struct{ from, to, want string }{
		{"", "v0.5.0", ""},
		{"v0.5.0", "v0.5.0", ""},
		{"v0.5.1", "v0.5.0", ""},
		{"v0.4.0", "v0.5.0", ruleBeforeOne},
		{"v0.6.0", "v0.5.0", ruleBeforeOne},
		{"v0.5", "v0.5.0", ""},
		{"v0.9.0", "v1.0.0", ruleBeforeOne},
		{"v1.2.0", "v1.3.0", ""},
		{"v1.2.0", "v1.4.0", ruleSkip},
		{"v1.3.0", "v1.2.0", ""},
		{"v1.4.0", "v1.2.0", ruleRollback},
		{"v1.4.0", "v2.0.0", ruleMajor},
		{"garbled", "v0.5.0", ""},
	} {
		if got := upgradeRule(c.from, c.to); got != c.want {
			t.Errorf("%s -> %s: rule %q, want %q", c.from, c.to, got, c.want)
		}
	}
	err := deploymentUpgradeRefused("v0.4.0", ruleBeforeOne).(*failure.Failure)
	if !strings.Contains(err.Observed, "Deployment ate-system/ate-api-server records Agent Substrate v0.4.0 on this cluster, and this KubernetesAgentSubstrate installs v0.5.0") ||
		!strings.Contains(err.NextStep, "kubectl -n ate-system delete pvc data-postgres-0") {
		t.Errorf("the refusal names the stamp, both releases and the way across:\n%s", err)
	}
}

func TestPoolsHaveUpstreamsWireShape(t *testing.T) {
	_, edKey, _ := ed25519.GenerateKey(rand.Reader)
	der, _ := x509.MarshalPKCS8PrivateKey(edKey)
	keyPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}))
	template := &x509.Certificate{SerialNumber: big.NewInt(1), IsCA: true, BasicConstraintsValid: true}
	certDER, _ := x509.CreateCertificate(rand.Reader, template, template, edKey.Public(), edKey)
	certPEM := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}))

	var pool struct {
		CAs []struct {
			ID                 string
			SigningKeyPKCS8    []byte
			RootCertificateDER []byte
		}
		ActiveForSigning string
	}
	if err := json.Unmarshal([]byte(caPoolJSON(keyPEM, certPEM)), &pool); err != nil {
		t.Fatal(err)
	}
	if pool.ActiveForSigning != "1" || len(pool.CAs) != 1 || pool.CAs[0].ID != "1" {
		t.Fatalf("one CA with ID 1, active: %+v", pool)
	}
	if _, err := x509.ParsePKCS8PrivateKey(pool.CAs[0].SigningKeyPKCS8); err != nil {
		t.Errorf("the key field is base64 PKCS #8 DER: %v", err)
	}
	if _, err := x509.ParseCertificate(pool.CAs[0].RootCertificateDER); err != nil {
		t.Errorf("the certificate field is base64 DER: %v", err)
	}

	ecKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	sec1, _ := x509.MarshalECPrivateKey(ecKey)
	pkcs8, err := toPKCS8PEM(string(pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: sec1})))
	if err != nil || !strings.HasPrefix(pkcs8, "-----BEGIN PRIVATE KEY-----") {
		t.Fatalf("the provider's SEC 1 key re-encodes as PKCS #8: %v", err)
	}
	spki, _ := x509.MarshalPKIXPublicKey(ecKey.Public())
	publicPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: spki}))
	var jwt struct {
		Authorities []struct {
			ID, Algorithm   string
			SigningKeyPKCS8 []byte
		}
		ActiveForSigning string
	}
	if err := json.Unmarshal([]byte(jwtPoolJSON("ES256", pkcs8, publicPEM)), &jwt); err != nil {
		t.Fatal(err)
	}
	if jwt.ActiveForSigning != jwtKeyID(publicPEM) || jwt.Authorities[0].Algorithm != "ES256" || len(jwt.ActiveForSigning) != 43 {
		t.Errorf("one ES256 key, active, named by a base64url SHA-256: %+v", jwt)
	}
	if _, err := base64.RawURLEncoding.DecodeString(jwt.ActiveForSigning); err != nil {
		t.Error("the key ID is base64url without padding")
	}
}

func TestFixturesArePublishedBytes(t *testing.T) {
	if os.Getenv("AGENT_SUBSTRATE_UPSTREAM_PIN") != "1" {
		t.Skip("set AGENT_SUBSTRATE_UPSTREAM_PIN=1 to compare the fixtures with what downloads.planton.ai serves")
	}
	local := fixtures(t)
	for _, path := range releaseFiles {
		served, err := fetchText(ReleaseFileURL(path))
		if err != nil {
			t.Fatal(err)
		}
		if served != local.files[path] {
			t.Errorf("testdata %s differs from the published file", path)
		}
	}
	setup, err := fetchText(SetupSQLURL())
	if err != nil {
		t.Fatal(err)
	}
	if setup != local.setupSQL {
		t.Error("testdata setup.sql differs from upstream's at the tag")
	}
	if response, err := http.Head(ReleaseFileURL("SHA256SUMS")); err != nil || response.StatusCode != 200 {
		t.Error("the pinned release is not sealed (SHA256SUMS is absent)")
	}
}

// markerLineLeft reports a line that still starts with marker, the lines
// upstream's splice replaces (its prose mentions the marker after "# ").
func markerLineLeft(text, marker string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), marker) {
			return true
		}
	}
	return false
}

func contains(list []string, want string) bool {
	for _, item := range list {
		if item == want {
			return true
		}
	}
	return false
}

// fullSurfaceSpec sets every field the generic platform admits, the shape
// the parity test and the full-surface lane exercise.
func fullSurfaceSpec() *v1.KubernetesAgentSubstrateSpec {
	resources := func(cpu string) *kubernetesprovider.ContainerResources {
		return &kubernetesprovider.ContainerResources{
			Requests: &kubernetesprovider.CpuMemory{Cpu: cpu, Memory: "128Mi"},
			Limits:   &kubernetesprovider.CpuMemory{Memory: "512Mi"},
		}
	}
	f64 := func(v float64) *float64 { return &v }
	str := func(v string) *string { return &v }
	return &v1.KubernetesAgentSubstrateSpec{
		Namespace: literal("ate-system"),
		Crds:      &v1.KubernetesAgentSubstrateCrds{Install: boolPtr(true), KeepOnUninstall: boolPtr(false)},
		Platform: &v1.KubernetesAgentSubstratePlatform{Kind: &v1.KubernetesAgentSubstratePlatform_Generic{Generic: &v1.KubernetesAgentSubstrateGeneric{
			ServiceAccountIssuer: "https://kubernetes.default.svc.cluster.local",
		}}},
		Database: &v1.KubernetesAgentSubstrateDatabase{
			Kind: &v1.KubernetesAgentSubstrateDatabase_Bundled{Bundled: &v1.KubernetesAgentSubstrateBundledDatabase{
				StorageClass: "standard", VolumeSize: str("10Gi"), Resources: resources("250m"),
			}},
			PoolMaxConnections: i32(16),
		},
		Storage: &v1.KubernetesAgentSubstrateStorage{Kind: &v1.KubernetesAgentSubstrateStorage_S3{S3: &v1.KubernetesAgentSubstrateS3Storage{
			Endpoint: literal("http://main-s3.seaweedfs.svc.cluster.local:8333"), Bucket: "snapshots", Region: str("eu-west-1"), ForcePathStyle: boolPtr(true),
			Credentials: &v1.KubernetesAgentSubstrateS3Credentials{SecretName: literal("substrate-s3"), AccessKeyIdKey: str("id"), SecretAccessKeyKey: str("secret")},
		}}},
		ApiServer: &v1.KubernetesAgentSubstrateApiServer{
			Replicas: i32(3), Resources: resources("200m"), LogLevel: v1.LogLevel_debug,
			Authorization:          &v1.KubernetesAgentSubstrateAuthorization{Enabled: boolPtr(true), BootstrapOwners: []string{"spiffe://cluster.local/ns/ops/sa/admin"}},
			ActorJwt:               &v1.KubernetesAgentSubstrateActorJwt{Algorithm: v1.KubernetesAgentSubstrateActorJwt_rs256, Issuer: "https://idp.example.com"},
			DrainDelay:             str("10s"),
			DrainTimeout:           str("20s"),
			TemplateResyncInterval: str("30s"),
		},
		Controller: &v1.KubernetesAgentSubstrateController{Resources: resources("100m"), LogLevel: v1.LogLevel_warn, WorkerSyncConcurrency: i32(4)},
		NodeAgent: &v1.KubernetesAgentSubstrateNodeAgent{
			Resources: resources("100m"), LogLevel: v1.LogLevel_error,
			NodeSelector:           map[string]string{"kubernetes.io/os": "linux"},
			Tolerations:            []*kubernetesprovider.WorkloadToleration{{Key: "sandbox", Operator: "Exists", Effect: "NoSchedule"}},
			ActorStatsPollInterval: str("30s"), DrainDelay: str("1s"), DrainTimeout: str("4m"),
			ImageCache: &v1.KubernetesAgentSubstrateImageCache{
				GcPeriod: str("10m"), HighPercent: i32(90), LowPercent: i32(70), MaxBytes: i64(10737418240), MinAge: str("5m"), DryRun: true,
			},
			AssetStagingCapacityGib: i32(12),
		},
		Router: &v1.KubernetesAgentSubstrateRouter{
			Replicas: i32(2), Resources: resources("100m"), LogLevel: v1.LogLevel_debug,
			AllowedClientSpiffeIds: []string{"spiffe://cluster.local/ns/app/sa/web"},
			RouteTimeout:           str("2m"),
			ParkedRequests: &v1.KubernetesAgentSubstrateParkedRequests{
				Budget: str("8s"), Max: i32(2048), RetryInterval: str("50ms"), RetryFactor: f64(1.5), RetryJitter: f64(0.2),
			},
			ExtProcMaxRequests: i32(500), DrainDelay: str("10s"), DrainTimeout: "25s",
		},
		Egress: &v1.KubernetesAgentSubstrateEgress{
			CredentialProvider: &v1.KubernetesAgentSubstrateCredentialProvider{Kind: &v1.KubernetesAgentSubstrateCredentialProvider_KubernetesSecrets{
				KubernetesSecrets: &v1.KubernetesAgentSubstrateKubernetesSecretsProvider{
					Grants:    []*v1.KubernetesAgentSubstrateNamespaceGrant{{Atespace: "team-a", AllowedNamespaces: []string{"team-a"}}},
					Resources: resources("50m"), LogLevel: v1.LogLevel_warn,
					KubeApiQps: f64(20), KubeApiBurst: i32(40), DrainGrace: str("3s"),
				},
			}},
			Replicas: i32(2), Resources: resources("100m"), LogLevel: v1.LogLevel_debug,
			PolicyCacheTtl: str("5s"), InterceptCertificateTtl: str("30m"), EnvoyConcurrency: i32(2),
			AdditionalExtProcService: &v1.KubernetesAgentSubstrateExtProcService{Namespace: "policy", Service: "filter", Port: 9443},
			DrainDelay:               str("5s"),
		},
		PodCertificates: &v1.KubernetesAgentSubstratePodCertificates{Resources: resources("50m"), WorkersPerSigner: i32(2), KubeApiQps: f64(50), KubeApiBurst: i32(100)},
		Telemetry: &v1.KubernetesAgentSubstrateTelemetry{
			OtlpEndpoint:     "http://otel-collector.observability.svc:4317",
			TracesExporters:  []string{"otlp"},
			MetricsExporters: []string{"otlp", "prometheus"},
			LogsExporters:    []string{"console"},
			TracesSampler:    "parentbased_traceidratio", TracesSamplerRatio: f64(0.25),
			MetricExportIntervalMs: i32(10000), MetricExportTimeoutMs: i32(5000),
		},
		ClusterSize:        v1.KubernetesAgentSubstrateSpec_size10,
		CordonControlPlane: true,
		ImageRegistry:      "mirror.example.com",
		ImagePullSecrets:   []string{"mirror-pull"},
	}
}

func boolPtr(v bool) *bool { return &v }
func i64(v int64) *int64   { return &v }

func TestTheFullSurfaceComposes(t *testing.T) {
	s, c := composeSpec(t, fullSurfaceSpec())
	if len(c.all()) == 0 || s.apiServer.jwtAlgorithm != "RS256" {
		t.Fatal("the every-setting spec composes")
	}
	router := args(t, find(c, "Deployment", vars.RouterName), vars.RouterName, false)
	for _, want := range []string{"--parked-request-retry-factor=1.5", "--parked-request-retry-jitter=0.2", "--drain-timeout=25s", "--log-level=debug"} {
		if !contains(router, want) {
			t.Errorf("router args lack %s: %v", want, router)
		}
	}
}

// TestTerraformTwinPinsTheSameRelease holds the upstream contract equal
// across the engines offline; the gated parity test proves the rest.
func TestTerraformTwinPinsTheSameRelease(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "tf", "locals.tf"))
	if err != nil {
		t.Fatal(err)
	}
	locals := string(raw)
	for name, value := range map[string]string{
		"release":                    vars.Release,
		"bundle_name":                vars.BundleName,
		"namespace":                  vars.Namespace,
		"pod_certificates_namespace": vars.PodCertificatesNamespace,
		"version_annotation":         vars.VersionAnnotation,
		"minor_label":                vars.MinorLabel,
		"env_hash_annotation":        vars.EnvHashAnnotation,
		"config_hash_annotation":     vars.ConfigHashAnnotation,
		"postgres_setup_config_map":  vars.PostgresSetupConfigMap,
		"postgres_setup_mount_path":  vars.PostgresSetupMountPath,
		"release_base_url":           releaseBaseURL,
		"setup_sql_base_url":         setupSQLBaseURL,
	} {
		found := false
		for _, line := range strings.Split(locals, "\n") {
			fields := strings.Fields(line)
			if len(fields) >= 3 && fields[0] == name && fields[1] == "=" {
				found = strings.Trim(fields[2], `"`) == value
				break
			}
		}
		if !found {
			t.Errorf("the Terraform module's local %s must equal the Pulumi module's %q", name, value)
		}
	}
	for _, path := range releaseFiles {
		if !strings.Contains(locals, `"`+path+`"`) {
			t.Errorf("the Terraform module does not read %s, which the Pulumi module does", path)
		}
	}
}
