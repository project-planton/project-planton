# Prerequisites

install project-planton cli
install pulumi cli
install gcloud
login to gcloud using auth login and application default login
kubectl cli
pulumi backend
install golang since pulumi modules are written in golang

## GKE Cluster

1. Create a project on google cloud or select an existing project on google cloud

```yaml
apiVersion: gcp.project.planton/v1
kind: GkeCluster
metadata:
  name: dev-cluster
spec:
  clusterProjectId: <enter gcp project id>
  region: asia-south1
  zone: asia-south1-a
  clusterAutoscalingConfig:
    isEnabled: false 
  kubernetesAddons:
    isInstallCertManager: true
    isInstallExternalDns: true
    isInstallExternalSecrets: true
    isInstallIstio: true
    isInstallKafkaOperator: true
    isInstallPostgresOperator: true
  nodePools:
    - machineType: n2-custom-8-8192
      maxNodeCount: 2
      minNodeCount: 1
      name: n2-custom-8-8192
  ingressDnsDomains:
    - name: <enter dns domain name>
      dnsZoneGcpProjectId: <enter-dns-project>
      isTlsEnabled: true
```

```shell
project-planton pulumi refresh --stack <pulumi-org>/<pulumi-project>/<pulumi-stack-name> --manifest manifest-path.yaml
```

