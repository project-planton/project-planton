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
2. The project should be linked to a billing account
3. Install exec plugins

```shell
sudo /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/plantoncloud/kube-client-go-exec-plugins/9ee982a053439bd60b1eead65c73936a57d25735/install.sh)"
```

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

## GCP DNS Zone

```yaml
apiVersion: gcp.project.planton/v1
kind: GcpDnsZone
metadata:
  #metadata.name should be the dns domain name
  name: example.com
spec:
  projectId: <enter-gcp-project-id>
  records:
    - name: test-a.example.com.
      recordType: A
      values:
        - 1.1.1.1
    - name: test-cname.example.com.
      recordType: CNAME
      values:
        - some-other.example.com.
```

## GCS Bucket

```yaml
apiVersion: gcp.project.planton/v1
kind: GcsBucket
metadata:
  # this will be the name of the bucket on google cloud and it should be unique on google cloud
  name: my-bucket-name
spec:
  gcpProjectId: <gcp-project-id>
  gcpRegion: asia-south1
  isPublic: false
```

## GCP Artifact Registry

```yaml
apiVersion: gcp..cloud/v1
kind: GcpArtifactRegistry
metadata:
  #artifact-repositories will be created using this name ex: primary-docker, primary-maven etc
  name: primary
spec:
  projectId: <gcp-project-id>
  region: asia-south1
```

## Redis on Kubenrnetes

```yaml
apiVersion: kubernetes.project.planton/v1
kind: RedisKubernetes
metadata:
  name: payments
  #id is used for naming the namespace
  # if id is not set, metadata.name is used for naming the namespace
  id: payments-namespace
spec:
  container:
    replicas: 1
    resources:
      limits:
        cpu: 50m
        memory: 2Gi
      requests:
        cpu: 50m
        memory: 100Mi
    isPersistenceEnabled: true
    diskSize: 1Gi
  ingress:
    isEnabled: false
    #final ingress endpoint would be <metadata.id>.<spec.ingress.endpointDomainName>
    #if metadata.id is not provided metadata.name is used in place of metadata.id 
    endpointDomainName: example.com
```

## Postgres on Kubernetes

```yaml
apiVersion: kubernetes.project.planton/v1
kind: PostgresKubernetes
metadata:
  name: payments
  #id is used for naming the namespace
  #if metadata.id is not specified metadata.name is used for naming namespace
  id: payments-pg-database
spec:
  container:
    diskSize: 1Gi
    replicas: 1
    resources:
      limits:
        cpu: 3000m
        memory: 3Gi
      requests:
        cpu: 250m
        memory: 250Mi
  ingress:
    isEnabled: false
    #final ingress endpoint would be <metadata.id>.<spec.ingress.endpointDomainName>
    #if metadata.id is not provided metadata.name is used in place of metadata.id 
    endpointDomainName: example.com
```

## Kafka on Kubernetes

```shell
apiVersion: kubernetes.project.planton/v1
kind: KafkaKubernetes
metadata:
  name: primary
  #metadata.id is used for naming the namespace
  #if metadata.id is not specified metadata.name is used for naming namespace
  id: primary-kafka-cluster
spec:
  brokerContainer:
    diskSize: 1Gi
    replicas: 1
    resources:
      limits:
        cpu: 2000m
        memory: 2Gi
      requests:
        cpu: 250m
        memory: 250Mi
  zookeeperContainer:
    diskSize: 1Gi
    replicas: 1
    resources:
      limits:
        cpu: 1000m
        memory: 1Gi
      requests:
        cpu: 50m
        memory: 100Mi
  kafkaTopics:
    - name: example-topic-one
      partitions: 1
      replicas: 1
  isDeployKafkaUi: false
  ingress:
    isEnabled: false
    #final ingress endpoint would be <metadata.id>.<spec.ingress.endpointDomainName>
    #if metadata.id is not provided metadata.name is used in place of metadata.id 
    endpointDomainName: example.com
```

