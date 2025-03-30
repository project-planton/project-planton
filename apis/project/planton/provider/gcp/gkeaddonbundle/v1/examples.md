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
  installIngressNginx: true
  installCertManager: true
  installExternalDns: true
  installExternalSecrets: true
  installKafkaOperator: true
  installPostgresOperator: true
  installSolrOperator: true
  installElasticOperator: true
```
