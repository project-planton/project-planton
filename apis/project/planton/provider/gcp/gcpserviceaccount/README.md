# GcpServiceAccount Component

The **GcpServiceAccount** API resource lets you declaratively create and manage Google Cloud service accounts—plus their
keys and IAM role bindings—using the same Project Planton workflow you already use for buckets, clusters, and other
components.  
With a few lines of YAML you can:

* provision the service account in a target project (or the provider-default project)
* optionally generate a JSON key and surface it as a stack output
* grant project-level or organization-level IAM roles in one step
* deploy through either **Pulumi** or **Terraform** with identical inputs

---

## Why We Built It

Every GKE cluster, Cloud Function, or workload identity eventually needs a service account. Doing it “by hand” means
juggling gcloud commands, IAM policies, key lifecycle, and diffs across environments.  
This component removes that toil by giving you:

* **Consistency** – one manifest style for every environment
* **Least-privilege by default** – supply only the roles you really need
* **Key hygiene** – create a key only when `createKey: true` and get the base-64 payload straight from stack outputs
* **Unified IaC** – the same spec deploys via Pulumi or Terraform without edits

---

## Key Features

| Capability                          | Details                                                                                               |
|-------------------------------------|-------------------------------------------------------------------------------------------------------|
| **Service account provisioning**    | Creates the account using `serviceAccountId` in the specified `projectId`.                            |
| **Optional key creation**           | Set `createKey: true` to receive a JSON key (`key_base64`) in stack outputs.                          |
| **IAM role binding (project)**      | Attach one or more roles via `projectIamRoles`; bindings are idempotent.                              |
| **IAM role binding (organization)** | Provide `orgId` and `orgIamRoles` for org-level privileges.                                           |
| **Cross-environment parity**        | Identical manifest applies in dev, staging, and prod—only metadata .name or project IDs change.       |
| **Pulumi & Terraform support**      | Pick your IaC engine at deploy time; the spec stays the same.                                         |
| **Built-in validation**             | Field-level checks (length, charset, required combos) prevent invalid GCP names before the plan runs. |

---

## Example Manifest

<details>
<summary>Minimal service account (no key, project roles only)</summary>

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpServiceAccount
metadata:
  name: log-writer-sa
spec:
  serviceAccountId: log-writer
  projectId: my-gcp-project
  projectIamRoles:
    - roles/logging.logWriter
```

</details>

<details open>
<summary>Service account with key + project & org bindings</summary>

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpServiceAccount
metadata:
  name: billing-access-sa
spec:
  serviceAccountId: billing-access
  projectId: billing-prod
  orgId: "123456789012"
  createKey: true
  projectIamRoles:
    - roles/billing.user
    - roles/logging.logWriter
  orgIamRoles:
    - roles/resourcemanager.organizationViewer
```

</details>

---

## Deployment Steps

1. **Validate**
   ```bash
   project-planton validate --manifest billing-access-sa.yaml
   ```

2. **Deploy (Pulumi)**
   ```bash
   project-planton pulumi up \
     --manifest billing-access-sa.yaml \
     --stack org/planton/prod
   ```  
   *or*

   **Deploy (Terraform)**
   ```bash
   project-planton terraform apply \
     --manifest billing-access-sa.yaml \
     --stack org/planton/prod
   ```

3. **Retrieve outputs**
   ```bash
   project-planton stack outputs --stack org/planton/prod
   # returns service account email and key_base64 (if createKey was true)
   ```

---

## Benefits Recap

* **Single-source manifests** replace ad-hoc scripts and web-console clicks.
* **Safer rollouts** thanks to schema validation and idempotent IAM bindings.
* **Faster onboarding** for new environments—just change `projectId` and apply.
* **Audit-friendly**: every change is version-controlled and surfaced as Stack Job history.

---

## Next Steps

* Explore more examples in [`examples.md`](./examples.md).
* Read the Protobuf spec for advanced field descriptions.
* Fork the default Pulumi/Terraform modules if you need custom tagging or VPC-scoped roles.

Happy deploying!
