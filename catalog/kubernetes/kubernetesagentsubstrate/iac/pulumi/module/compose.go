package module

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// releaseFileSet is every pinned file the composition reads, by path, plus
// upstream's database setup script.
type releaseFileSet struct {
	files    map[string]string
	setupSQL string
}

// composition is the document set the module applies besides the kept
// definitions and the generated Secrets, in four phases, each applied once
// the one before has rolled out and deleted in reverse:
//   - namespaces: ate-system when this resource owns it, and
//     podcertificate-controller-system, which it always owns;
//   - prerequisites: the controller's RBAC, the spec-derived ConfigMaps, the
//     SandboxConfig admission policy, and the pod-certificate controller,
//     which every later component's certificates come from;
//   - database: the shared OpenTelemetry ConfigMap every workload reads, and
//     the bundled PostgreSQL with its setup script;
//   - system: the API server, controller, node agent, router, egress
//     gateway, credential provider and network policies;
//   - runtime: the default gVisor SandboxConfig, applied last so a destroy
//     deletes it while the controller that clears its finalizer still runs.
//
// Every object is upstream's own with the typed settings applied, plus the
// few the module adds (the database setup ConfigMap). The Terraform twin
// builds the same groups in locals.tf; a gated parity test holds them equal.
type composition struct {
	namespaces    []document
	prerequisites []document
	database      []document
	system        []document
	runtime       []document

	// images is upstream's image list (images.json), for the worker-image
	// outputs.
	images map[string]string
}

func (c *composition) all() []document {
	var out []document
	for _, group := range [][]document{c.namespaces, c.prerequisites, c.database, c.system, c.runtime} {
		out = append(out, group...)
	}
	return out
}

// compose builds the document set from the pinned files.
func compose(s *settings, f releaseFileSet) (*composition, error) {
	parse := func(path string) ([]document, error) {
		text, ok := f.files[path]
		if !ok {
			return nil, errors.Errorf("the release file %s was not fetched", path)
		}
		return splitDocuments(text)
	}
	c := &composition{}

	var imageList struct {
		Images map[string]string `json:"images"`
	}
	if err := json.Unmarshal([]byte(f.files[fileImages]), &imageList); err != nil {
		return nil, errors.Wrap(err, "parsing the release's images.json")
	}
	c.images = imageList.Images

	// Namespaces. ate-controller.yaml repeats ate-system; one identity is
	// applied, and only when this resource owns it.
	namespaceDocs, err := parse(fileNamespace)
	if err != nil {
		return nil, err
	}
	if s.createNamespace {
		for _, d := range namespaceDocs {
			labelWith(d, s.identityLabels)
			c.namespaces = append(c.namespaces, d)
		}
	}

	// Prerequisites.
	roleDocs, err := parse(fileControllerRole)
	if err != nil {
		return nil, err
	}
	c.prerequisites = append(c.prerequisites, roleDocs...)
	c.prerequisites = append(c.prerequisites,
		configMap(configMapAPIEnv, apiEnvConfigMap(s)),
		configMap(configMapAuthentication, map[string]string{"authentication.yaml": authenticationYAML(s)}),
	)
	if s.egress.provider == "kubernetes" {
		c.prerequisites = append(c.prerequisites, configMap(configMapPolicy, map[string]string{"namespace-policy.yaml": policyYAML(s.egress.grants)}))
	}
	podCertDocs, err := parse(filePodCertificates)
	if err != nil {
		return nil, err
	}
	if s.size10 {
		size10, err := parseKustomization(f.files[filePodCertSize10])
		if err != nil {
			return nil, err
		}
		if podCertDocs, err = applyComponent(size10, podCertDocs, "Deployment/"+vars.PodCertificatesName); err != nil {
			return nil, err
		}
	}
	for _, d := range podCertDocs {
		if str(d, "kind") == "Namespace" {
			labelWith(d, s.identityLabels)
			c.namespaces = append(c.namespaces, d)
			continue
		}
		if str(d, "kind") == "Deployment" && str(d, "metadata", "name") == vars.PodCertificatesName {
			if err := patchPodCertificates(d, s); err != nil {
				return nil, err
			}
		}
		c.prerequisites = append(c.prerequisites, d)
	}
	validationDocs, err := parse(fileSandboxValidation)
	if err != nil {
		return nil, err
	}
	c.prerequisites = append(c.prerequisites, validationDocs...)

	// Database phase.
	otelDocs, err := parse(fileOtelConfig)
	if err != nil {
		return nil, err
	}
	configMaps := map[string]map[string]string{
		configMapAuthentication: {"authentication.yaml": authenticationYAML(s)},
	}
	for _, d := range otelDocs {
		data := otelData(s, stringMap(d["data"]))
		d["data"] = toAnyMap(data)
		configMaps[configMapOtel] = data
		c.database = append(c.database, d)
	}
	if s.egress.provider == "kubernetes" {
		configMaps[configMapPolicy] = map[string]string{"namespace-policy.yaml": policyYAML(s.egress.grants)}
	}
	if s.database.kind == "bundled" {
		postgresDocs, err := parse(filePostgres)
		if err != nil {
			return nil, err
		}
		setup := map[string]string{"setup.sql": f.setupSQL}
		c.database = append(c.database, configMap(vars.PostgresSetupConfigMap, setup))
		configMaps[vars.PostgresSetupConfigMap] = setup
		for _, d := range postgresDocs {
			switch {
			case str(d, "kind") == "ConfigMap" && str(d, "metadata", "name") == configMapPostgres:
				if s.size10 {
					if err := applySize10PostgresConfig(d, f.files[filePostgresSize10]); err != nil {
						return nil, err
					}
				}
				configMaps[configMapPostgres] = stringMap(d["data"])
			case str(d, "kind") == "StatefulSet" && str(d, "metadata", "name") == vars.PostgresName:
				if err := patchPostgres(d, s); err != nil {
					return nil, err
				}
			}
			c.database = append(c.database, d)
		}
	}

	// System phase.
	var system []document
	for _, path := range []string{fileAPIServer, fileController, fileNodeAgent, fileRouter, fileNetworkPolicies, fileEgress} {
		docs, err := parse(path)
		if err != nil {
			return nil, err
		}
		for _, d := range docs {
			if str(d, "kind") == "Namespace" {
				continue
			}
			system = append(system, d)
		}
	}
	if s.gke {
		monitoring, err := parse(fileRouterMonitoring)
		if err != nil {
			return nil, err
		}
		system = append(system, monitoring...)
	}
	if s.egress.provider == "kubernetes" {
		providerDocs, err := parse(fileCredentialProvider)
		if err != nil {
			return nil, err
		}
		system = append(system, providerDocs...)
	}
	if s.router.agentgateway {
		if system, err = applyAgentgateway(system, f); err != nil {
			return nil, err
		}
	}
	for _, d := range system {
		if err := patchSystemDocument(d, s); err != nil {
			return nil, err
		}
	}
	if s.database.kind == "cloudsql" {
		if system, err = applyCloudSQLSidecar(system, f.files[fileCloudSQLPatch]); err != nil {
			return nil, err
		}
	}
	c.system = system

	// Runtime phase.
	if c.runtime, err = parse(fileSandboxGvisor); err != nil {
		return nil, err
	}

	// Passes over every document: ownership, placement, images, rollout
	// hashes, version stamps.
	if s.cordon {
		cordon, err := parseKustomization(f.files[fileCordon])
		if err != nil {
			return nil, err
		}
		for _, group := range []*[]document{&c.prerequisites, &c.database, &c.system} {
			if *group, err = applyComponent(cordon, *group); err != nil {
				return nil, err
			}
		}
	}
	apiEnv, apiSecret := apiEnvConfigMap(s), apiEnvSecret(s)
	for _, d := range c.all() {
		mirrorImages(d, s.registry)
		if !isWorkload(d) {
			continue
		}
		labelWith(child(d, "metadata"), map[string]string{ownerLabel: kindName})
		if len(s.pullSecrets) > 0 {
			podSpec(d)["imagePullSecrets"] = pullSecretList(s.pullSecrets)
		}
		annotations := child(child(child(d, "spec"), "template"), "metadata")
		if readers, ok := configReaders[str(d, "metadata", "name")]; ok {
			child(annotations, "annotations")[vars.ConfigHashAnnotation] = configHash(readers, configMaps)
		}
		if str(d, "metadata", "name") == vars.APIServerName {
			child(annotations, "annotations")[vars.EnvHashAnnotation] = envHash(apiEnv, apiSecret)
			child(child(d, "metadata"), "annotations")[vars.VersionAnnotation] = vars.Release
		}
	}
	return c, nil
}

// labelWith adds labels to an object (or to a metadata map directly).
func labelWith(object document, labels map[string]string) {
	metadata := object
	if _, ok := object["metadata"]; ok {
		metadata = child(object, "metadata")
	}
	target := child(metadata, "labels")
	for k, v := range labels {
		target[k] = v
	}
}

// isWorkload reports whether a document runs pods under a fixed name.
func isWorkload(d document) bool {
	switch str(d, "kind") {
	case "Deployment", "DaemonSet", "StatefulSet":
		return true
	}
	return false
}

// configMap is a ConfigMap in ate-system.
func configMap(name string, data map[string]string) document {
	return document{
		"apiVersion": "v1",
		"kind":       "ConfigMap",
		"metadata":   document{"name": name, "namespace": vars.Namespace},
		"data":       toAnyMap(data),
	}
}

func toAnyMap(m map[string]string) document {
	out := document{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func stringMap(value interface{}) map[string]string {
	out := map[string]string{}
	m, _ := value.(map[string]interface{})
	for k, v := range m {
		if s, ok := v.(string); ok {
			out[k] = s
		}
	}
	return out
}

func pullSecretList(names []string) []interface{} {
	out := make([]interface{}, 0, len(names))
	for _, name := range names {
		out = append(out, document{"name": name})
	}
	return out
}
