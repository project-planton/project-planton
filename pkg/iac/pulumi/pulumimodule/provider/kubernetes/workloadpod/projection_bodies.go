package workloadpod

import (
	kubernetesv1 "github.com/plantonhq/planton/catalog/kubernetes"
)

// ClusterTrustBundleProjectionBody renders a clusterTrustBundle projection as
// a manifest map, for kinds that write their pod spec into a custom resource
// or Helm values instead of typed args: the bundle by name or by signer (and
// label selector), written to "ca.crt" unless a path is set -- the same
// result clusterTrustBundleProjection builds for typed args.
func ClusterTrustBundleProjectionBody(source *kubernetesv1.ClusterTrustBundleVolumeSource) map[string]interface{} {
	path := source.GetPath()
	if path == "" {
		path = "ca.crt"
	}
	body := map[string]interface{}{"path": path}
	if name := source.GetName().GetValue(); name != "" {
		body["name"] = name
	}
	if source.GetSignerName() != "" {
		body["signerName"] = source.GetSignerName()
	}
	if selector := source.GetLabelSelector(); selector != nil {
		labelSelector := map[string]interface{}{}
		if len(selector.GetMatchLabels()) > 0 {
			matchLabels := map[string]interface{}{}
			for k, val := range selector.GetMatchLabels() {
				matchLabels[k] = val
			}
			labelSelector["matchLabels"] = matchLabels
		}
		if len(selector.GetMatchExpressions()) > 0 {
			expressions := make([]interface{}, 0, len(selector.GetMatchExpressions()))
			for _, e := range selector.GetMatchExpressions() {
				expression := map[string]interface{}{"key": e.GetKey(), "operator": e.GetOperator()}
				if len(e.GetValues()) > 0 {
					values := make([]interface{}, 0, len(e.GetValues()))
					for _, val := range e.GetValues() {
						values = append(values, val)
					}
					expression["values"] = values
				}
				expressions = append(expressions, expression)
			}
			labelSelector["matchExpressions"] = expressions
		}
		body["labelSelector"] = labelSelector
	}
	if source.GetOptional() {
		body["optional"] = true
	}
	return body
}

// PodCertificateProjectionBody renders a podCertificate projection as a
// manifest map, writing the credential bundle to "credential-bundle.pem"
// unless the source names its own paths -- the same result
// podCertificateProjection builds for typed args.
func PodCertificateProjectionBody(source *kubernetesv1.PodCertificateVolumeSource) map[string]interface{} {
	body := map[string]interface{}{
		"signerName": source.GetSignerName(),
		"keyType":    source.GetKeyType(),
	}
	bundlePath := source.GetCredentialBundlePath()
	if bundlePath == "" && source.GetKeyPath() == "" {
		bundlePath = "credential-bundle.pem"
	}
	if bundlePath != "" {
		body["credentialBundlePath"] = bundlePath
	}
	if source.GetKeyPath() != "" {
		body["keyPath"] = source.GetKeyPath()
		body["certificateChainPath"] = source.GetCertificateChainPath()
	}
	if source.MaxExpirationSeconds != nil {
		body["maxExpirationSeconds"] = int(source.GetMaxExpirationSeconds())
	}
	if len(source.GetUserAnnotations()) > 0 {
		annotations := map[string]interface{}{}
		for k, val := range source.GetUserAnnotations() {
			annotations[k] = val
		}
		body["userAnnotations"] = annotations
	}
	return body
}
