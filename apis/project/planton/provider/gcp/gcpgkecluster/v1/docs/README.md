# GKE Cluster Deployment: From Click-Ops to Production-Ready Infrastructure

## Introduction

The Kubernetes ecosystem presents a unique paradox. On one hand, Kubernetes is the open-source standard for container orchestration, designed to run anywhere. On the other hand, managing a production-grade Kubernetes cluster is notoriously complex—from certificate rotation and etcd backups to CNI plugins, CSI drivers, and version upgrades.

Google Kubernetes Engine (GKE) was created to resolve this paradox. As a managed Kubernetes service on Google Cloud Platform, GKE handles the operational complexity of the control plane while offering different levels of infrastructure management. The strategic question for platform teams isn't just *whether* to use GKE, but *how* to deploy it, *which mode* to choose, and *what configuration* to expose to application teams.

This document explores the deployment methods for GKE clusters across the industry, examines the evolution from manual provisioning to declarative infrastructure, compares Autopilot vs. Standard modes, and explains why Project Planton provides an opinionated Standard mode cluster with intelligent defaults.

## The GKE Architecture: Control Plane, Nodes, and Modes

Before diving into deployment methods, it's essential to understand what you're deploying.

A GKE cluster consists of two fundamental components:

1. **Control Plane**: Runs the Kubernetes API server, scheduler, controllers, and etcd. In GKE, Google fully manages this layer—provisioning, upgrading, securing, and maintaining high availability. The control plane's boot disk and etcd database are encrypted by default.

2. **Worker Nodes**: Typically Google Compute Engine VMs that run your containerized workloads. The control plane orchestrates these nodes via the Kubernetes API.

The key differentiator in GKE is the **operational mode**, which determines who manages the nodes:

- **GKE Standard Mode**: Google manages the control plane; you manage the nodes by defining node pools (groups of VMs with shared configuration like machine type, disk size, and autoscaling parameters). You're responsible for capacity planning, node upgrades, and machine type selection. Billing is based on the provisioned VMs, including all unused capacity.

- **GKE Autopilot Mode**: Google manages both the control plane *and* the nodes. You define workload resource requests (CPU, memory), and GKE automatically provisions, scales, and manages the underlying infrastructure. Billing is per-pod based on actual resource requests, eliminating the "bin-packing problem" and unused capacity waste. Autopilot enforces Google's security and operational best practices by default.

This architectural choice is immutable after cluster creation and fundamentally shapes the API surface and operational model.

## The Maturity Spectrum: How Teams Deploy GKE Clusters

Deployment methods for GKE exist on a spectrum, from manual and fragile to fully automated and production-ready.

### Level 0: The Anti-Pattern — Manual Console Provisioning

**What it is**: Using the Google Cloud Console (web UI) to click through forms, select options, and create a GKE cluster.

**Why teams start here**: It's the fastest path to "hello world." The console provides immediate visual feedback and requires no coding knowledge.

**The hidden cost**: Every click creates **undocumented infrastructure**. When the cluster fails or needs to be replicated (for staging, DR, or a second region), there's no record of the configuration. Teams end up with "snowflake clusters"—each one subtly different, impossible to version control, and a nightmare for compliance and audit.

**Verdict**: Acceptable for learning and experimentation. A production anti-pattern.

---

### Level 1: Imperative Scripting — gcloud CLI

**What it is**: Using the `gcloud container clusters create` command in shell scripts to provision clusters.

**What it solves**: Repeatability. The command can be saved in a script, version-controlled, and executed multiple times. It's the official GCP tool and supports all GKE features.

**What it doesn't solve**: State management. Running the same script twice doesn't update the cluster; it attempts to create a duplicate, which fails. There's no built-in mechanism to detect drift (manual changes made via the console) or perform updates. Deleting a cluster requires a separate `delete` command, making cleanup error-prone.

**Example**:
```bash
gcloud container clusters create prod-cluster \
  --region us-central1 \
  --num-nodes 3 \
  --machine-type n2-standard-4 \
  --enable-autoscaling \
  --enable-workload-identity \
  --release-channel stable
```

**Verdict**: Better than manual, but still imperative. Lacks the declarative power needed for production infrastructure.

---

### Level 2: Configuration Management — Ansible

**What it is**: Using Ansible's `gcp_container_cluster` module to provision GKE clusters as part of a configuration management playbook.

**What it solves**: For organizations already using Ansible to configure servers and applications, it provides a familiar interface. Ansible can be made *idempotent*—running the playbook multiple times converges to the desired state without errors.

**What it doesn't solve**: Infrastructure state tracking. Ansible doesn't maintain a durable state file of what it created. If a team member manually modifies a cluster outside Ansible, there's no native drift detection. It also lacks the rich ecosystem of reusable modules that modern IaC tools provide.

**Verdict**: A viable option for Ansible-centric organizations, but not the industry standard for cloud infrastructure.

---

### Level 3: Infrastructure as Code — Terraform and Pulumi

**What it is**: Declarative infrastructure management using tools that maintain a *state file* tracking all provisioned resources.

**Why it's production-ready**: 

- **Declarative**: You define the desired state (a cluster with specific properties), and the tool calculates and executes the minimal set of changes to converge from the current state to the desired state.
- **State Management**: A state file (stored in Google Cloud Storage or the Pulumi Service) tracks what was created, enabling updates, drift detection, and safe deletions.
- **Plan and Apply Workflow**: Changes are previewed (`terraform plan` / `pulumi preview`) before execution, preventing surprises.
- **Modularity**: Reusable modules encapsulate best practices (e.g., the official `terraform-google-kubernetes-engine` module).

**A Critical Best Practice — Cluster and Node Pool Separation**:

Both Terraform and Pulumi codify a non-obvious pattern: the cluster control plane and the node pools are separate resources. This separation is crucial because it allows node pools to be added, removed, or resized without destroying and recreating the entire cluster.

Terraform example:
```hcl
resource "google_container_cluster" "primary" {
  name     = "prod-cluster"
  location = "us-central1"

  # Remove the default node pool after creation
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_nodes" {
  cluster    = google_container_cluster.primary.name
  location   = google_container_cluster.primary.location
  node_count = 3

  node_config {
    machine_type = "n2-standard-4"
  }

  autoscaling {
    min_node_count = 1
    max_node_count = 10
  }
}
```

**Terraform vs. Pulumi**:

| Feature | Terraform | Pulumi |
|---------|-----------|--------|
| **Language** | HCL (Domain-Specific Language) | General-purpose (Go, Python, TypeScript) |
| **State Management** | Remote backend (GCS, S3) | Pulumi Service (default) or self-hosted |
| **Secret Handling** | **Secrets stored in plaintext** in state file; requires external tool (Vault, Secret Manager) | **Secrets encrypted by default** in Pulumi Service |
| **Modularity** | Modules (reusable HCL) | Functions, classes, and libraries in your language |
| **Multi-Environment** | Workspaces or directory separation | Stacks (e.g., `dev`, `staging`, `prod`) |
| **GKE Pattern** | `google_container_cluster` + `google_container_node_pool` | `gcp.container.Cluster` + `gcp.container.NodePool` |

**Verdict**: This is the industry standard for production GKE deployments. Terraform has the largest ecosystem; Pulumi offers the power of general-purpose languages and superior secret management.

---

### Level 4: Kubernetes-Native Control Plane — Config Connector and Crossplane

**What it is**: An operator-driven model where GCP resources (including GKE clusters) are defined as Kubernetes Custom Resources and continuously reconciled by an in-cluster controller.

**Google Cloud Config Connector**: A Google-provided Kubernetes operator that runs inside a GKE cluster. You define a `ContainerCluster` Custom Resource, apply it to the cluster, and the operator creates and manages the actual GKE cluster in GCP. If someone manually changes the cluster, the operator automatically reverts the change (continuous reconciliation).

**Crossplane**: An open-source, multi-cloud evolution of the Config Connector concept. It turns a Kubernetes cluster into a "universal control plane" for managing infrastructure across GCP, AWS, and Azure.

**What it solves**: 

- **Continuous Reconciliation**: Unlike Terraform (which only reconciles on `apply`), these operators continuously monitor and repair drift.
- **Platform Abstraction**: Platform teams can define high-level abstractions (e.g., `XCluster`) that hide the complexity of VPCs, node pools, and IAM from application developers.

**What it doesn't solve**: 

- **Bootstrapping Problem**: You need a Kubernetes cluster to run the operator that creates Kubernetes clusters. This creates a chicken-and-egg problem, often requiring an initial Terraform-provisioned "management cluster."
- **Operational Complexity**: Running a Kubernetes operator for infrastructure is powerful but adds operational overhead. You're now managing the health of the operator itself.

**Verdict**: Ideal for platform engineering teams building internal developer platforms (IDPs) with high levels of abstraction. Overkill for most teams; Terraform/Pulumi remains the pragmatic choice.

---

## Autopilot vs. Standard: The Operational Decision

Before choosing *how* to deploy GKE, you must choose *which mode* to deploy.

### Autopilot: The Fully Managed, Opinionated Path

**What it is**: A "serverless Kubernetes" experience. You define workloads with resource requests, and GKE handles the rest—provisioning nodes, scaling infrastructure, applying security policies, and performing upgrades.

**What's pre-configured (non-negotiable)**:
- Regional topology (high availability)
- VPC-native networking
- Workload Identity (secure GCP API access)
- Shielded GKE nodes (Secure Boot, vTPM)
- GKE Dataplane V2 (eBPF-based networking)
- Auto-upgrades via release channels

**Pricing**: Billed per-pod based on actual CPU, memory, and storage requests. You don't pay for system overhead or unused capacity.

**When to choose Autopilot**:
- Your priority is application development, not infrastructure management.
- Your workloads are stateless, containerized, and follow standard Kubernetes patterns.
- You want Google's security and operational best practices enforced by default.
- Your team is new to Kubernetes and wants to avoid the complexity of node pool management.

**When Autopilot isn't suitable**:
- You need privileged containers or specific host-level access (e.g., custom kernel modules).
- You require specific machine types (e.g., high-memory instances, GPUs) not available in Autopilot's abstraction.
- You have a large, predictable baseline workload and a platform team that can optimize Standard mode with Committed Use Discounts.

---

### Standard: The Full-Control, Node-Managed Path

**What it is**: Google manages the control plane; you manage the nodes by defining node pools, selecting machine types, configuring autoscaling, and managing upgrades.

**What you control**:
- Node pool configuration: machine type, disk size, count, autoscaling parameters
- Spot VMs for cost optimization
- Custom taints, labels, and node affinity
- Regional vs. zonal topology
- Advanced network policies and security settings

**Pricing**: Billed per-second for the underlying Compute Engine VMs. You pay for the entire VM, including unused capacity and system overhead.

**When to choose Standard**:
- You need granular control over infrastructure (specific machine types, custom OS images, kernel parameters).
- You have a dedicated platform team capable of managing node pools, capacity planning, and cost optimization.
- Your workloads require configurations Autopilot restricts (privileged access, specific Linux capabilities).
- You have large, stable workloads and can leverage Committed Use Discounts to reduce costs below Autopilot's per-pod pricing.

**Cost optimization in Standard mode**:

Elite cost performers combine multiple strategies:
- **Committed Use Discounts (CUDs)**: Purchase 1- or 3-year commitments for baseline workloads (e.g., `n2-standard-16` instances).
- **Spot VMs**: Use interruptible instances (up to 91% savings) for fault-tolerant batch jobs.
- **Cluster Autoscaler**: Dynamically scale node pools based on pod scheduling needs, avoiding over-provisioning.

---

## Project Planton's Choice: Standard Mode with Intelligent Defaults

Project Planton provides a **GKE Standard mode cluster** with an opinionated API designed around the 80/20 principle—exposing the 20% of configuration that 80% of production teams need.

### Why Standard Mode?

1. **Flexibility**: Standard mode supports a broader range of workloads, from stateless web applications to GPU-intensive ML pipelines, privileged infrastructure controllers, and legacy applications requiring specific node configurations.

2. **Cost Optimization Control**: Platform teams can implement sophisticated cost optimization strategies—mixing CUDs for baseline capacity, Spot VMs for batch workloads, and autoscaling for spikes.

3. **Multi-Tenant Platform Flexibility**: In a multi-tenant platform, different teams often have radically different requirements (e.g., a data science team needs GPU nodes; a web team needs high-CPU instances). Standard mode's node pool model elegantly supports this diversity.

4. **Transparent Learning Path**: By exposing node pools, machine types, and autoscaling as first-class configuration, Project Planton helps teams understand and control the infrastructure. Autopilot's abstraction, while convenient, can obscure operational realities.

That said, Autopilot is an excellent choice for many teams. Organizations already using Autopilot can continue to do so; Project Planton's Standard mode API is designed for teams that need or want infrastructure control.

---

### The Project Planton API: Essential Configuration

The `GcpGkeCluster` API exposes the critical fields for production clusters:

**Cluster Basics**:
- `cluster_project_id`: The GCP project for the cluster
- `region` / `zone`: Cluster location (regional for HA, zonal for dev/test)
- `is_workload_logs_enabled`: Toggle for pod log forwarding to Cloud Logging

**Shared VPC** (for multi-project organizations):
- `shared_vpc_config.is_enabled`: Whether to use a Shared VPC
- `shared_vpc_config.vpc_project_id`: The host project containing the VPC

**Cluster Autoscaling** (Node Auto-Provisioning):
- `cluster_autoscaling_config.is_enabled`: Enable automatic node pool creation
- `cpu_min_cores` / `cpu_max_cores`: Total CPU resource limits
- `memory_min_gb` / `memory_max_gb`: Total memory resource limits

**Node Pools** (the core of Standard mode):
- `name`: Identifier for the node pool (also added as a label for pod scheduling)
- `machine_type`: GCE machine type (e.g., `e2-medium`, `n2-standard-4`, `n2-custom-8-16384`)
- `min_node_count` / `max_node_count`: Autoscaling boundaries
- `is_spot_enabled`: Use Spot VMs for cost savings

---

### What's Missing (and Why)

The Project Planton API intentionally omits advanced fields that most teams don't need:

- **Release Channels**: While critical for production (enabling auto-upgrades), Project Planton's IaC modules set a sensible default (`STABLE` for prod, `REGULAR` for staging).
- **Networking Details** (VPC, subnets, IP ranges): These are assumed to be provisioned separately and referenced via Shared VPC config.
- **Workload Identity**: Enabled by default in the underlying IaC, as it's a non-negotiable security best practice.
- **Private Cluster Settings**: Platform-level decisions (private nodes, private endpoint, authorized networks) are configured in the IaC module, not exposed in the API, to ensure consistent security posture.
- **Monitoring and Logging**: GKE integrates with Cloud Operations (Stackdriver) by default; the API provides a simple toggle for workload logs.

This design reflects Project Planton's philosophy: **expose the configuration that varies across environments and workloads; hide the configuration that should be consistent and secure by default**.

---

## Production-Ready Patterns

The following patterns are essential for production GKE clusters, whether managed via Project Planton or directly via Terraform/Pulumi.

### High Availability: Regional Clusters

**Zonal Cluster**: Control plane and nodes in a single zone. If that zone fails, the cluster is unavailable.

**Regional Cluster** (recommended for production): Control plane replicated across three zones. If a zone fails, the control plane remains available, and workloads on nodes in other zones continue running.

**Cost**: Regional clusters have the same management fee as zonal clusters ($0.10/hour), making them the clear choice for production.

---

### Networking: VPC-Native and Private Clusters

**VPC-Native (IP Aliasing)**: Pod IPs are allocated from a secondary IP range in the VPC subnet, making them first-class VPC citizens. This enables:
- Native routing (no NAT overhead)
- Direct firewall rule targeting
- Integration with Cloud SQL, Memorystore, and other VPC-native services

**Private Clusters**: Production security baseline.
- **Private Nodes**: Nodes have only internal IPs (no public internet exposure).
- **Private Endpoint**: Kubernetes API server is internal-only (accessible from within the VPC or via VPN/Cloud Interconnect).
- **Master Authorized Networks**: If the API server has a public endpoint, this whitelist restricts access to trusted CIDR blocks.

---

### Node Pool Strategies

**Anti-pattern**: Using the default node pool created with the cluster. It has a generic configuration and can't be removed without recreating the cluster.

**Production pattern**: Remove the default node pool (`remove_default_node_pool = true` in Terraform) and create custom node pools for different workload classes:

- **Baseline Pool**: `n2-standard-4` or `n2-standard-8`, backed by Committed Use Discounts, for stable, long-running workloads.
- **Batch Pool**: `e2-medium` Spot VMs (`is_spot_enabled = true`), with `min_node_count = 0` and `max_node_count = 50`, for fault-tolerant, interruptible jobs.
- **GPU Pool**: `n1-standard-4` with NVIDIA T4 GPUs, for ML training and inference.

Each node pool includes autoscaling (`min_node_count` / `max_node_count`), auto-repair, and auto-upgrade.

---

### Security: Workload Identity and Pod Security

**Workload Identity**: The recommended method for GKE workloads to access GCP services (Cloud Storage, BigQuery, etc.). It links a Kubernetes Service Account (KSA) to a Google IAM Service Account (GSA), allowing pods to impersonate the GSA without storing static keys.

**Pod Security Standards (PSS)**: The modern replacement for PodSecurityPolicy (deprecated and removed in Kubernetes v1.25+). PSS enforces security policies (`Privileged`, `Baseline`, `Restricted`) via labels on namespaces, not complex cluster-wide resources.

**Binary Authorization** (advanced): Deploy-time policy enforcement. Only container images that have been cryptographically signed (e.g., in your CI/CD pipeline) are allowed to run.

---

## Deployment Workflow: From API to Running Cluster

Project Planton's deployment workflow follows a two-step pattern:

1. **Provision Infrastructure**: The IaC tool (Terraform or Pulumi) reads the `GcpGkeCluster` protobuf spec and provisions:
   - The GKE cluster control plane
   - Node pools with autoscaling
   - IAM bindings (for Shared VPC, Workload Identity)
   - Outputs (cluster endpoint, CA certificate, credentials)

2. **Deploy Workloads** (optional): The IaC tool can immediately connect to the new cluster's API endpoint and apply Kubernetes manifests:
   - Gateway API resources (for L7 load balancing)
   - ConfigConnector operator (if building a Kubernetes-native platform)
   - Application namespaces, RBAC, and resource quotas

This two-step approach is a production pattern for bootstrapping clusters with "Day 1.5" configuration, ensuring clusters are immediately usable by application teams.

---

## Cost Optimization Strategies

**For Standard Mode**:

1. **Committed Use Discounts (CUDs)**: Purchase 1- or 3-year commitments for baseline vCPU and memory. CUDs apply automatically to matching VMs in node pools. Typical savings: 30-50%.

2. **Spot VMs**: For batch jobs, CI/CD workers, and fault-tolerant workloads. Savings: up to 91%. Enable with `is_spot_enabled = true` in node pool configuration.

3. **Right-Sizing Node Pools**: Use E2 machine types for general-purpose workloads (lowest cost), N2 for balanced performance, and C2/M2 only when workload characteristics demand it.

4. **Cluster Autoscaler**: Set `min_node_count = 0` for node pools that support intermittent workloads. Nodes scale to zero when idle, eliminating waste.

5. **Monitor and Optimize**: Use GKE's built-in cost allocation (via labels) to understand which teams/applications drive costs, and optimize accordingly.

**For Autopilot Mode** (if chosen outside Project Planton):
- Purchase GKE-specific Committed Use Discounts (spend-based, not resource-based).
- Ensure workloads define accurate resource requests (over-requesting inflates costs).

---

## Backup and Disaster Recovery

**Backup for GKE**: An add-on service that must be enabled on the cluster. It provides:
- **Config Backup**: Kubernetes manifests (Deployments, Services, etc.), capturing cluster state.
- **Volume Backup**: Snapshots of PersistentVolumeClaims (PVCs).

**Disaster Recovery Strategy**:
1. **Infrastructure**: Store IaC code (Terraform/Pulumi) in Git. Cluster recreation is a `terraform apply`.
2. **Application State**: Enable Backup for GKE for stateful workloads. Schedule regular backups.
3. **Multi-Region**: For mission-critical applications, deploy to multiple regions with a global load balancer (Cloud Load Balancing with multi-region backends).

---

## Next Steps

This document provides a strategic overview of GKE deployment methods and production architecture. For deeper implementation details, refer to:

- **Terraform Module**: [terraform-google-kubernetes-engine](https://github.com/terraform-google-modules/terraform-google-kubernetes-engine) — official Terraform module with best practices
- **Pulumi Examples**: [GKE Tutorial](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/gke/)
- **GCP Documentation**: [GKE Best Practices](https://cloud.google.com/kubernetes-engine/docs/best-practices)
- **Project Planton IaC Modules**: [iac/pulumi/](../iac/pulumi/) — the implementation of Project Planton's GKE cluster provisioning

---

## Conclusion

The decision to deploy GKE is straightforward—it eliminates the operational burden of managing Kubernetes control planes. The deeper questions are:

1. **Which mode?** Autopilot for hands-off simplicity and enforced best practices. Standard for flexibility, control, and advanced cost optimization.

2. **How to deploy?** Move beyond manual and imperative approaches as quickly as possible. Adopt Terraform or Pulumi for declarative, state-managed infrastructure that can be versioned, reviewed, and safely updated.

3. **What to configure?** Focus on the 80/20—cluster location, node pools, machine types, autoscaling, and cost optimization strategies. Trust Google to handle the rest (control plane HA, security patches, etcd backups).

Project Planton's `GcpGkeCluster` API is designed around these principles, providing an opinionated, production-ready Standard mode cluster that balances simplicity with control, making Kubernetes infrastructure manageable for teams of all sizes.

