# PostgreSQL on Kubernetes: The Modern Landscape

## Introduction: The Great Database Debate (Resolved)

For years, asking "Should I run a database in Kubernetes?" would spark heated debates in engineering circles. The conventional wisdom was clear: **don't do it**. Kubernetes was designed for stateless applications, and databases are the antithesis of stateless.

That advice is now outdated.

Today, running PostgreSQL on Kubernetes isn't just viable—it's become the strategic choice for organizations building internal Database-as-a-Service (DBaaS) platforms. Companies like Zalando, GitLab, and countless others run production PostgreSQL databases on Kubernetes at scale.

**Why the shift?** Two compelling reasons:

1. **Economics**: Managed services like Amazon RDS can be prohibitively expensive, especially at scale. They also impose their own maintenance schedules and limitations that may conflict with your operational needs.

2. **Automation**: Kubernetes' declarative API enables automation and extensibility that would be difficult to achieve otherwise. Your database becomes part of your GitOps workflow, managed alongside your applications with the same tools and processes.

The goal isn't to lift-and-shift a database into a container just to colocate it with your app. It's to build an internal, on-demand, cost-effective data platform that integrates seamlessly with your infrastructure-as-code workflow.

This guide explains the landscape of options and why we chose Zalando Postgres Operator as the default for our IaC modules.

---

## The Maturity Spectrum: From Broken to Production-Ready

Running PostgreSQL on Kubernetes requires solving three fundamental challenges:

1. **Data Persistence**: Pods are ephemeral; their filesystems die with them. Your database must survive pod restarts.
2. **Stable Identity**: Database nodes aren't interchangeable. A primary must be identifiable as the primary, and replicas must maintain their roles.
3. **Day 2 Operations**: Production databases require extensive operational knowledge—automated failover, backups, point-in-time recovery, upgrades, and more.

Different deployment approaches solve these problems with varying degrees of success. Think of it as a maturity ladder:

### Level 0: The Anti-Pattern (Don't Do This)

**Simple Pods or Deployments** with PostgreSQL will cause data loss. Period.

- **Problem 1**: If you run `kubectl run postgres`, the database stores data in the container's filesystem. When the pod dies, your data is gone forever.
- **Problem 2**: If you use a Deployment with a PersistentVolumeClaim and set `replicas: 3`, all three pods will mount the **same** volume. PostgreSQL's data directory expects exclusive access. Multiple processes fighting for the same PGDATA directory will corrupt your database instantly and irreversibly.

**Verdict**: This is not a deployment strategy. It's a disaster waiting to happen.

### Level 1: The Foundation (StatefulSet)

**StatefulSet** is the correct Kubernetes primitive for stateful applications. It solves the first two problems:

- **Stable Identity**: Pods get predictable names (`postgres-0`, `postgres-1`) that persist across restarts.
- **Stable Storage**: Each pod gets its own PersistentVolumeClaim via `volumeClaimTemplates`. When `postgres-1` restarts, it reattaches to its original volume.

**What's Missing**: StatefulSet has zero PostgreSQL-specific knowledge. It can restart `postgres-0`, but it doesn't know that `postgres-0` was the primary. It can't promote a replica, reconfigure replication, or handle failover.

**Verdict**: Necessary but not sufficient. You've solved the Kubernetes problem, but you're stuck with massive operational toil.

### Level 2: The Packaging Layer (Helm Charts)

**Helm** is a package manager that bundles the YAML complexity into a single deployable chart. Popular charts like Bitnami's PostgreSQL-HA add tools like:

- **repmgr**: Automates replication management and failover
- **PgBouncer**: Connection pooling and load balancing
- **pgpool-II**: Connection routing to the current primary

**What's Missing**: Helm is a "fire-and-forget" installer. It has no running component that actively manages your cluster after deployment. If `repmgr` fails or `pgpool-II` loses track of the primary, Helm doesn't know and can't react.

**Verdict**: Better than manual StatefulSets, but Day 2 operations (failovers, upgrades, backups) remain problematic.

### Level 3: The Production Solution (Kubernetes Operators)

**Kubernetes Operators** are the state-of-the-art approach. An operator is software that runs in your cluster, continuously watching your database's state and autonomously managing its entire lifecycle.

Think of it as encoding an expert PostgreSQL SRE's knowledge into a control loop that never sleeps.

**How Operators Work**:

1. **Custom Resource Definition (CRD)**: You define `kind: PostgresCluster` instead of manually creating StatefulSets, Services, and ConfigMaps.
2. **Controller**: A pod running the operator software watches for changes to your PostgresCluster objects.
3. **Reconciliation Loop**: The operator continuously compares your desired state (what you declared) with actual state (what's running) and takes action to reconcile any differences.

**What This Enables**:

- **Automated Failover**: Primary dies? The operator detects it, promotes the most up-to-date replica, reconfigures other replicas, and updates the service endpoint—all in seconds.
- **Automated Backups**: Declare `backups: enabled` in your manifest, and the operator schedules continuous backups with point-in-time recovery capability.
- **Automated Upgrades**: Change `version: 14` to `version: 15`, and the operator performs a safe, rolling upgrade following PostgreSQL best practices.

**Verdict**: This is the only approach that provides production-grade, autonomous lifecycle management.

---

## The Operator Landscape: Which One Should You Use?

The PostgreSQL-on-Kubernetes ecosystem has matured around several production-ready operators. Here's the landscape:

### CloudNativePG (by EDB)

**Philosophy**: Pure Kubernetes-native. No external dependencies like Patroni.

**Architecture**: Uses the Kubernetes API directly for coordination and leader election. Doesn't use StatefulSets—manages Pods and PVCs directly for granular control.

**Strengths**:
- True open-source (Apache 2.0) with strong community momentum
- Built-in Barman Cloud for backups with PITR
- Zero-data-loss configurations with synchronous replication
- First-class volume snapshot support

**Best For**: Teams that prioritize "pure" Kubernetes-native solutions and want to avoid external orchestration tools.

### Crunchy Data PGO (Postgres Operator)

**Philosophy**: Battle-tested enterprise solution with best-in-class backup capabilities.

**Architecture**: Uses Patroni for high availability and has deep integration with pgBackRest for backups.

**Strengths**:
- One of the oldest and most mature operators
- Exceptional pgBackRest integration (backups enabled by default)
- Asymmetric topologies (different resources for primary vs. replicas)
- Synchronous replication support

**Licensing Gotcha**: While the source code is Apache 2.0, the pre-built container images require a commercial subscription for production use under Crunchy's "Developer Program" terms.

**Best For**: Organizations that want a commercial, vendor-supported solution with enterprise SLAs.

### Zalando Postgres Operator

**Philosophy**: Proven at scale in production at Zalando. Batteries-included with integrated connection pooling.

**Architecture**: Uses Patroni for HA, packaged in the "Spilo" container (PostgreSQL + Patroni + WAL-G). Uses Kubernetes API as the Distributed Configuration Store.

**Strengths**:
- 100% open-source (MIT license—no production restrictions)
- **Integrated PgBouncer**: Built-in connection pooling managed by the operator
- WAL-G for backups with full PITR to object storage (S3, GCS, Azure)
- Fast in-place major version upgrades
- Battle-tested at scale in Zalando's production for years

**Best For**: Teams that need integrated connection pooling and want a fully open-source, production-proven solution.

### Percona Operator for PostgreSQL

**Philosophy**: 100% open-source fork of Crunchy PGO with integrated monitoring.

**Architecture**: Forked from Crunchy PGO, inheriting its mature Patroni + pgBackRest design.

**Strengths**:
- All the benefits of PGO without licensing restrictions
- 100% open-source (Apache 2.0) including container images
- Out-of-the-box integration with Percona Monitoring & Management (PMM)
- Enterprise-grade monitoring with zero setup

**Best For**: Teams that want PGO's architecture without commercial restrictions, especially if integrated monitoring is important.

### KubeDB (by AppsCode)

**Philosophy**: Multi-database platform for unified management.

**Architecture**: Polyglot operator that manages PostgreSQL, MySQL, MongoDB, Elasticsearch, Redis, and more.

**Strengths**:
- Single control plane for all database types
- Unified CRDs and operational patterns
- KubeStash for backups with PITR
- Integrated PgBouncer support

**Licensing Gotcha**: Uses an "open-core" model. Essential Day 2 features (backups, scaling, TLS, pooling) are only in the paid Enterprise Edition.

**Best For**: Organizations willing to pay for a commercial platform to manage multiple database types with a single vendor.

---

## Why We Chose Zalando for Default IaC Modules

After evaluating the landscape, we selected **Zalando Postgres Operator** as the default for our Pulumi and Terraform modules for several reasons:

### 1. Truly Open Source

Zalando is **MIT licensed** with no production restrictions. Unlike Crunchy PGO's container images or KubeDB's feature gating, every feature is available without a commercial subscription. For an open-source IaC project, this alignment is critical.

### 2. Production-Proven at Scale

Zalando has run this operator in production for years, managing hundreds of PostgreSQL clusters. It's not a lab project—it's battle-tested infrastructure.

### 3. Integrated Connection Pooling

The built-in PgBouncer integration is a differentiator. Most applications need connection pooling, and having the operator automatically deploy and configure PgBouncer alongside your database is a significant operational win.

### 4. Comprehensive Backup & DR

WAL-G integration provides continuous backups with point-in-time recovery to object storage (S3, GCS, Azure, Cloudflare R2). The operator also supports the **standby-then-promote pattern** for cross-cluster disaster recovery—critical for real-world scenarios where the source cluster may be destroyed.

### 5. Operational Simplicity

The operator manages the entire lifecycle declaratively. Deploy a cluster, scale replicas, perform major version upgrades, configure backups—all through simple manifest changes. This aligns perfectly with GitOps workflows.

---

## The API: Kubernetes-Native, Manager-Agnostic

Our `PostgresKubernetes` API (see `api.proto`) is designed to be **manager-agnostic**. While our default IaC modules use Zalando operator under the hood, the API abstracts those details. You declare your desired PostgreSQL cluster—version, resources, storage, backups—and the IaC layer translates that to the appropriate operator manifests.

This design allows future flexibility. If Zalando isn't the right fit for your use case, you could implement alternative IaC modules using CloudNativePG, Percona, or even a different approach entirely, all while keeping the same user-facing API.

---

## Deep Dive: Zalando Operator

For a comprehensive guide to how Zalando Postgres Operator works—its architecture, day-to-day operations, disaster recovery patterns, connection pooling, monitoring, and how it maps to our API—see the **[Zalando Operator Guide](./zalando-operator.md)**.

That guide is essential reading if you're:
- Operating PostgreSQL clusters deployed via our IaC modules
- Troubleshooting issues
- Understanding disaster recovery procedures
- Configuring backups and restores
- Planning upgrades or scaling operations

---

## Conclusion: The Paradigm Has Shifted

Running PostgreSQL on Kubernetes is no longer a controversial choice—it's a mature, well-supported pattern with robust tooling. The key is understanding the maturity spectrum and choosing the right level of abstraction for your needs.

**For production workloads**, only Kubernetes Operators (Level 3) provide the autonomous, continuous lifecycle management required. Among those operators, **Zalando Postgres Operator** offers the best balance of open-source licensing, production readiness, and operational features for most teams.

Our IaC modules abstract the complexity, giving you production-grade PostgreSQL clusters through simple, declarative manifests. Under the hood, Zalando operator handles the hard parts—high availability, automated failover, backups, disaster recovery, and upgrades—so you can focus on your applications.

Welcome to the modern era of database operations. Kubernetes and PostgreSQL love each other now. 🐘❤️☸️

---

## Further Reading

- **[Zalando Operator Deep Dive](./zalando-operator.md)** - Comprehensive guide to Zalando's architecture and operations
- [Zalando Postgres Operator Documentation](https://opensource.zalando.com/postgres-operator/) - Official upstream docs
- [PostgreSQL on Kubernetes: The Good Parts](https://www.crunchydata.com/blog/stateful-postgres-storage-using-kubernetes) - Crunchy Data's perspective
- [CloudNativePG Documentation](https://cloudnative-pg.io/) - Alternative operator documentation

