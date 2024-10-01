# Example 1: Basic Kafka Kubernetes Setup

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: KafkaKubernetes
metadata:
  name: kafka-cluster-basic
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  kafka_topics:
    - name: my-topic
      partitions: 3
      replicas: 2
  broker_container:
    replicas: 1
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
    disk_size: 20Gi
  zookeeper_container:
    replicas: 3
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
    disk_size: 10Gi
  ingress:
    enabled: false
  is_deploy_kafka_ui: false
```

# Example 2: Kafka Kubernetes with Schema Registry and Kafka UI

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: KafkaKubernetes
metadata:
  name: kafka-cluster-with-schema-registry
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  kafka_topics:
    - name: my-topic
      partitions: 3
      replicas: 3
  broker_container:
    replicas: 2
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 2
        memory: 2Gi
    disk_size: 50Gi
  zookeeper_container:
    replicas: 3
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
    disk_size: 20Gi
  schema_registry_container:
    is_enabled: true
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
  ingress:
    enabled: true
    ingressClassName: "nginx"
    hosts:
      - host: kafka.mydomain.com
        paths:
          - /
  is_deploy_kafka_ui: true
```

# Example 3: Kafka Kubernetes with Minimal Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: KafkaKubernetes
metadata:
  name: kafka-minimal
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  broker_container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
    disk_size: 10Gi
  zookeeper_container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
    disk_size: 10Gi
  ingress:
    enabled: false
  is_deploy_kafka_ui: false
```

# Example 4: Kafka Kubernetes with Custom Topic Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: KafkaKubernetes
metadata:
  name: kafka-custom-topics
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  kafka_topics:
    - name: logs
      partitions: 5
      replicas: 3
      config:
        retention.ms: "86400000"
        cleanup.policy: "delete"
    - name: metrics
      partitions: 10
      replicas: 2
      config:
        retention.ms: "3600000"
        cleanup.policy: "compact"
  broker_container:
    replicas: 3
    resources:
      requests:
        cpu: 500m
        memory: 2Gi
      limits:
        cpu: 4
        memory: 4Gi
    disk_size: 100Gi
  zookeeper_container:
    replicas: 3
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 2
        memory: 2Gi
    disk_size: 20Gi
  ingress:
    enabled: true
  is_deploy_kafka_ui: true
```

# Example 5: Kafka Kubernetes with Schema Registry but No Kafka UI

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: KafkaKubernetes
metadata:
  name: kafka-with-schema-registry
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  kafka_topics:
    - name: transactions
      partitions: 3
      replicas: 2
  broker_container:
    replicas: 2
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 2
        memory: 2Gi
    disk_size: 50Gi
  zookeeper_container:
    replicas: 3
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
    disk_size: 20Gi
  schema_registry_container:
    is_enabled: true
    replicas: 2
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
  ingress:
    enabled: true
  is_deploy_kafka_ui: false
``` 
