# GcpServiceAccount Examples

The `GcpServiceAccount` component lets you declaratively create Google Cloud service accounts, optionally generate a
key, and grant IAM roles at the project or organization level—all through the familiar ProjectPlanton YAML workflow. The
examples below illustrate typical scenarios.

---

## 1. Minimal Service Account (no key, default project)

Creates a service account in the provider’s default project without generating a key or granting any IAM roles.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpServiceAccount
metadata:
  name: analytics-sa
spec:
  serviceAccountId: analytics-sa
```

---

## 2. Service Account with Project-Level Roles and Key

Creates the service account in a specific project, generates a JSON key, and binds two project-level roles.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpServiceAccount
metadata:
  name: logging-writer-sa
spec:
  serviceAccountId: logging-writer
  projectId: my-app-prod-1234
  createKey: true
  projectIamRoles:
    - roles/logging.logWriter
    - roles/monitoring.metricWriter
```

---

## 3. Service Account with Organization-Level Roles

Creates the account, skips key creation, assigns one project role and one organization-level role.  
Useful when the service account must operate across multiple projects in the same organization.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpServiceAccount
metadata:
  name: org-auditor-sa
spec:
  serviceAccountId: org-auditor
  projectId: shared-infra-5678
  orgId: "123456789012"
  createKey: false
  projectIamRoles:
    - roles/viewer
  orgIamRoles:
    - roles/resourcemanager.organizationViewer
```

---

### Deploying the Examples

```bash
# Validate the manifest
project-planton validate --manifest example.yaml

# Deploy with Pulumi
project-planton pulumi up --manifest example.yaml --stack org/project/stack

# …or with Terraform
project-planton terraform apply --manifest example.yaml --stack org/project/stack
```

These examples demonstrate how the `GcpServiceAccount` component streamlines service-account management—eliminating
manual steps for key generation and IAM binding while fitting seamlessly into ProjectPlanton’s multi-cloud workflow.
