# Pulumi Examples for KubernetesNats

This document provides Pulumi usage examples for deploying NATS on Kubernetes using the official NATS Helm chart.

## Example 1: Basic NATS Cluster with Default Settings

Deploy a simple NATS cluster with default replicas, resources, and JetStream enabled.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsBasic = new kubernetesnats.NatsKubernetes("nats-basic", {
    metadata: {
        name: "nats-basic",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-basic",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 3,
            resources: {
                limits: {
                    cpu: "1000m",
                    memory: "2Gi",
                },
                requests: {
                    cpu: "100m",
                    memory: "256Mi",
                },
            },
            diskSize: "10Gi",
        },
        disableJetStream: false,
        tlsEnabled: false,
        disableNatsBox: false,
    },
});

// Export the internal client URL
export const internalClientUrl = natsBasic.internalClientUrl;
export const namespace = natsBasic.namespace;
```

## Example 2: NATS Cluster with Bearer Token Authentication

Set up a secure NATS cluster using bearer token authentication.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsSecure = new kubernetesnats.NatsKubernetes("nats-secure", {
    metadata: {
        name: "nats-secure",
        org: "my-organization",
        env: "production",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-secure",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 5,
            resources: {
                limits: {
                    cpu: "2000m",
                    memory: "4Gi",
                },
                requests: {
                    cpu: "500m",
                    memory: "1Gi",
                },
            },
            diskSize: "20Gi",
        },
        auth: {
            enabled: true,
            scheme: "bearer_token",
        },
        tlsEnabled: true,
        disableJetStream: false,
    },
});

// Export connection details
export const namespace = natsSecure.namespace;
export const internalClientUrl = natsSecure.internalClientUrl;
export const authSecretName = "auth-nats";
export const authSecretKey = "token";
```

## Example 3: NATS Cluster with Basic Authentication

Deploy NATS with basic username/password authentication.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsBasicAuth = new kubernetesnats.NatsKubernetes("nats-basic-auth", {
    metadata: {
        name: "nats-basic-auth",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-basic-auth",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 3,
            resources: {
                limits: {
                    cpu: "1000m",
                    memory: "2Gi",
                },
                requests: {
                    cpu: "100m",
                    memory: "256Mi",
                },
            },
            diskSize: "10Gi",
        },
        auth: {
            enabled: true,
            scheme: "basic_auth",
        },
        tlsEnabled: false,
    },
});

export const namespace = natsBasicAuth.namespace;
export const username = "nats"; // Default admin username
export const authSecretName = "auth-nats";
export const passwordKey = "password";
```

## Example 4: NATS Cluster with External Access via Ingress

Deploy NATS configured with ingress to allow external clients to access messaging services.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsExternal = new kubernetesnats.NatsKubernetes("nats-external", {
    metadata: {
        name: "nats-external",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-external",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 3,
            resources: {
                limits: {
                    cpu: "1000m",
                    memory: "2Gi",
                },
                requests: {
                    cpu: "100m",
                    memory: "256Mi",
                },
            },
            diskSize: "10Gi",
        },
        ingress: {
            enabled: true,
            hostname: "nats.example.com",
        },
        auth: {
            enabled: true,
            scheme: "basic_auth",
        },
        tlsEnabled: true,
    },
});

// Export connection details
export const namespace = natsExternal.namespace;
export const externalHostname = natsExternal.externalHostname;
export const internalClientUrl = natsExternal.internalClientUrl;
export const externalClientUrl = `nats://nats.example.com:4222`;
```

## Example 5: Lightweight NATS Cluster Without JetStream

Set up a lightweight NATS messaging cluster without JetStream persistence.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsMinimal = new kubernetesnats.NatsKubernetes("nats-minimal", {
    metadata: {
        name: "nats-minimal",
        env: "development",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-minimal",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 1,
            resources: {
                limits: {
                    cpu: "500m",
                    memory: "512Mi",
                },
                requests: {
                    cpu: "100m",
                    memory: "128Mi",
                },
            },
            diskSize: "1Gi",
        },
        disableJetStream: true,
        tlsEnabled: false,
        disableNatsBox: true,
    },
});

export const namespace = natsMinimal.namespace;
export const internalClientUrl = natsMinimal.internalClientUrl;
```

## Example 6: High Availability NATS Cluster with Basic Auth and No-Auth User

Deploy a highly available NATS cluster with both authenticated and unauthenticated access.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kubernetesnats from "@project-planton/kubernetesnats";

const natsHA = new kubernetesnats.NatsKubernetes("nats-ha-metrics", {
    metadata: {
        name: "nats-ha-metrics",
        org: "acme-corp",
        env: "production",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: "nats-ha-metrics",
        },
        createNamespace: true,
        serverContainer: {
            replicas: 7,
            resources: {
                limits: {
                    cpu: "4000m",
                    memory: "8Gi",
                },
                requests: {
                    cpu: "1000m",
                    memory: "2Gi",
                },
            },
            diskSize: "50Gi",
        },
        auth: {
            enabled: true,
            scheme: "basic_auth",
            noAuthUser: {
                enabled: true,
                publishSubjects: [
                    "telemetry.>",
                    "metrics.>",
                ],
            },
        },
        tlsEnabled: true,
        disableJetStream: false,
        ingress: {
            enabled: true,
            hostname: "nats-ha.example.com",
        },
    },
});

// Export comprehensive connection details
export const namespace = natsHA.namespace;
export const externalHostname = natsHA.externalHostname;
export const internalClientUrl = natsHA.internalClientUrl;
export const metricsEndpoint = `http://nats-prom.${natsHA.namespace}.svc.cluster.local:7777/metrics`;
```

## Connecting to NATS from Pulumi Programs

### From within the Kubernetes cluster:

```typescript
import * as pulumi from "@pulumi/pulumi";

const nats = // ... your NATS deployment

// Use the internal client URL
const clientUrl = nats.internalClientUrl; // nats://<service>.<namespace>.svc.cluster.local:4222
```

### Retrieving Authentication Credentials:

For bearer token authentication:

```typescript
import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";

const nats = // ... your NATS deployment with bearer_token auth

const authSecret = k8s.core.v1.Secret.get("auth-secret", 
    pulumi.interpolate`${nats.namespace}/auth-nats`);

export const bearerToken = authSecret.data.apply(d => 
    Buffer.from(d["token"], "base64").toString());
```

For basic authentication:

```typescript
import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";

const nats = // ... your NATS deployment with basic_auth

const authSecret = k8s.core.v1.Secret.get("auth-secret", 
    pulumi.interpolate`${nats.namespace}/auth-nats`);

export const username = authSecret.data.apply(d => 
    Buffer.from(d["user"], "base64").toString());
export const password = authSecret.data.apply(d => 
    Buffer.from(d["password"], "base64").toString());
```

## Example 7: Using Pre-existing Namespace

Deploy NATS into an existing namespace that's managed separately.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";
import * as kubernetesnats from "@project-planton/kubernetesnats";

// Assume namespace is created elsewhere (e.g., GitOps, separate stack)
const existingNamespace = "shared-messaging";

const natsWithExistingNs = new kubernetesnats.NatsKubernetes("nats-existing-ns", {
    metadata: {
        name: "nats-existing-ns",
    },
    spec: {
        targetCluster: {
            clusterName: "my-gke-cluster",
        },
        namespace: {
            value: existingNamespace,
        },
        createNamespace: false, // Don't create the namespace
        serverContainer: {
            replicas: 3,
            resources: {
                limits: {
                    cpu: "1000m",
                    memory: "2Gi",
                },
                requests: {
                    cpu: "100m",
                    memory: "256Mi",
                },
            },
            diskSize: "10Gi",
        },
        auth: {
            enabled: true,
            scheme: "basic_auth",
        },
        tlsEnabled: true,
    },
});

export const namespace = existingNamespace;
export const internalClientUrl = natsWithExistingNs.internalClientUrl;
```

**Use Case**: Ideal for GitOps workflows or when namespace has specific ResourceQuotas, NetworkPolicies, or other configurations managed separately.

## Notes

- The deployment uses the official NATS Helm chart for production-grade message streaming
- JetStream provides persistent streaming capabilities with file-based storage
- For production deployments, use an odd number of replicas (3, 5, 7) for optimal quorum
- TLS encryption is recommended for production workloads
- The `nats-box` utility pod is deployed by default for debugging and testing
- Metrics are exposed at port 7777 for Prometheus scraping

