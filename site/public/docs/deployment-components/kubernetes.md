---
title: "Kubernetes Deployments"
description: "Deploy applications and databases to Kubernetes clusters"
icon: "cloud"
order: 1
badge: "Popular"
---

# Kubernetes Deployments

ProjectPlanton provides several deployment components for Kubernetes that work on any Kubernetes cluster - whether it's a local Kind cluster, GKE, EKS, AKS, or any other Kubernetes distribution.

## Available Kubernetes Components

### Databases

- **PostgresKubernetes** - PostgreSQL database
- **RedisKubernetes** - Redis cache/database
- **MongodbKubernetes** - MongoDB document database
- **MySqlKubernetes** - MySQL relational database

### Applications

- **MicroserviceKubernetes** - Deploy containerized microservices
- **KafkaKubernetes** - Apache Kafka message streaming
- **TemporalKubernetes** - Temporal workflow engine
- **AirflowKubernetes** - Apache Airflow orchestration

## Example: Deploy Redis to Kubernetes

Here's a complete example of deploying Redis to a Kubernetes cluster.

### Prerequisites

- A Kubernetes cluster (local or cloud)
- `kubectl` configured to access your cluster
- ProjectPlanton CLI installed

### Create the Manifest

Create `redis.yaml`:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: RedisKubernetes
metadata:
  name: session-store
  org: acme
  env: production
spec:
  container:
    replicas: 3
    resources:
      limits:
        cpu: 500m
        memory: 1Gi
      requests:
        cpu: 250m
        memory: 512Mi
    isPersistenceEnabled: true
    diskSize: 20Gi
```

### Deploy

```bash
# Validate first (optional but recommended)
project-planton validate --manifest redis.yaml

# Deploy with Pulumi
pulumi login --local  # Use local backend for this example
project-planton pulumi up --manifest redis.yaml --stack acme/platform/prod
```

### Verify

```bash
# Check pods
kubectl get pods -l app=session-store

# Check services
kubectl get svc -l app=session-store

# Check persistent volumes
kubectl get pvc -l app=session-store
```

## Example: Deploy PostgreSQL to Kubernetes

Create `postgres.yaml`:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: app-database
  org: acme
  env: production
spec:
  container:
    replicas: 1
    resources:
      limits:
        cpu: 1000m
        memory: 2Gi
      requests:
        cpu: 500m
        memory: 1Gi
    isPersistenceEnabled: true
    diskSize: 50Gi
  database:
    name: myapp
    username: appuser
    # Password will be generated and stored in Kubernetes secret
```

Deploy:

```bash
project-planton pulumi up --manifest postgres.yaml --stack acme/platform/prod
```

### Access the Database

```bash
# Get the generated password
kubectl get secret app-database-postgresql -o jsonpath='{.data.postgresql-password}' | base64 -d

# Port forward to access locally
kubectl port-forward svc/app-database-postgresql 5432:5432

# Connect with psql
psql -h localhost -U appuser -d myapp
```

## Example: Deploy a Microservice

Create `api-service.yaml`:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: api-service
  org: acme
  env: production
spec:
  container:
    image: myorg/api-service:v1.2.3
    replicas: 3
    resources:
      limits:
        cpu: 1000m
        memory: 512Mi
      requests:
        cpu: 500m
        memory: 256Mi
    env:
      - name: DATABASE_URL
        value: "postgresql://app-database-postgresql:5432/myapp"
      - name: REDIS_URL
        value: "redis://session-store-redis:6379"
  service:
    type: LoadBalancer
    port: 8080
  ingress:
    enabled: true
    host: api.example.com
    tls:
      enabled: true
```

Deploy:

```bash
project-planton pulumi up --manifest api-service.yaml --stack acme/platform/prod
```

## Resource Configuration

### CPU and Memory

Specify resources using Kubernetes units:

```yaml
resources:
  limits:
    cpu: 1000m     # 1 CPU core
    memory: 2Gi    # 2 gigabytes
  requests:
    cpu: 500m      # 0.5 CPU cores
    memory: 1Gi    # 1 gigabyte
```

### Persistence

Enable persistence for stateful workloads:

```yaml
isPersistenceEnabled: true
diskSize: 50Gi
storageClassName: standard  # Optional: specify storage class
```

### Replicas

For high availability:

```yaml
replicas: 3  # Run 3 instances
```

## Multi-Environment Deployments

Use the same manifest across environments by overriding values:

```bash
# Development
project-planton pulumi up \
  --manifest redis.yaml \
  --set spec.container.replicas=1 \
  --set spec.container.diskSize=10Gi \
  --stack acme/platform/dev

# Production
project-planton pulumi up \
  --manifest redis.yaml \
  --set spec.container.replicas=3 \
  --set spec.container.diskSize=50Gi \
  --stack acme/platform/prod
```

## Best Practices

1. **Use Resource Limits**: Always specify resource limits to prevent resource exhaustion
2. **Enable Persistence**: For stateful workloads, enable persistence to prevent data loss
3. **Run Multiple Replicas**: For production, run at least 3 replicas for high availability
4. **Use Secrets**: Store sensitive data in Kubernetes secrets, not in manifests
5. **Test Locally**: Use Kind or Minikube to test deployments locally first

## Troubleshooting

### Pods Not Starting

```bash
# Check pod status
kubectl get pods

# View pod logs
kubectl logs <pod-name>

# Describe pod for events
kubectl describe pod <pod-name>
```

### Insufficient Resources

If pods remain in `Pending` state:

```bash
# Check node resources
kubectl top nodes

# Describe pod to see events
kubectl describe pod <pod-name>
```

### Storage Issues

If PVCs are pending:

```bash
# Check PVC status
kubectl get pvc

# Describe PVC
kubectl describe pvc <pvc-name>

# Check if storage class exists
kubectl get storageclass
```

## Next Steps

- [Getting Started](../getting-started) - Learn the basics
- [Architecture](../concepts/architecture) - Understand how it works
- Explore other [deployment components](index)

