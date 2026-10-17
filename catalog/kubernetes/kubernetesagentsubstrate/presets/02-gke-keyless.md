# GKE Keyless

Agent Substrate on GKE with no key anywhere. The API server logs into **Cloud SQL** with IAM as its own Google service account, through upstream's Cloud SQL Auth Proxy. Snapshots go to a **GCS bucket**. The API server and the node agent act as their Google service accounts through Workload Identity, and the API server trusts the cluster's own token issuer.

Pick this on Google Cloud. The cluster must have been created with the pod-certificate APIs (GKE 1.36 with the `certificates.k8s.io/v1beta1` beta APIs, or 1.37).

Before you deploy:
- The **Cloud SQL** instance: PostgreSQL 18, a database named `atepg`, the `cloudsql.iam_authentication` flag on, an IAM user for the API server's service account, and upstream's setup script run once.
- The **service accounts**: the API server's needs `roles/cloudsql.client`, `roles/cloudsql.instanceUser`, and `storage.objectAdmin` and `storage.bucketViewer` on the bucket; the node agent's needs the two storage roles and `artifactregistry.reader`. Bind each to its Kubernetes service account (`ate-system/ate-api-server`, `ate-system/atelet`).

Change first: the four references, to your resources' names.

See [02-gke-keyless.yaml](./02-gke-keyless.yaml) for the manifest.
