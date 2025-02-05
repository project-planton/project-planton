# Example 1: Basic GKE Addon Bundle

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeAddonBundle
metadata:
  name: basic-http-endpoint
spec:
  clusterProjectId: id-of-the-gcp-project
  istio:
    enabled: false
  isInstallIngressNginx: true
  isInstallCertManager: true
  isInstallExternalDns: true
  isInstallExternalSecrets: true
  isInstallKafkaOperator: true
  isInstallPostgresOperator: true
  isInstallSolrOperator: true
  isInstallElasticOperator: true
```
