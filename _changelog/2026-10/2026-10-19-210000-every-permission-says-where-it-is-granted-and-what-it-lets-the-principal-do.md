# Every Permission Says Where It Is Granted and What It Lets the Principal Do

## Why

A connection to a cloud is granted from these manifests: a GCP project role, AWS managed policies, an Azure custom role. Two facts decide how each entry may be granted, and the manifests did not say either of them.

- **Where an entry is granted.** A project role cannot hold a permission that authorizes on an organization or a billing account. Until now, the only way to find those permissions was to ask Google one project at a time, and wait for it to refuse.
- **What an entry lets the principal do.** Deploying a Cloud Run service needs to run it as a service account. Making someone an owner of the project is power over the whole account. These must be granted differently: the first by default, the second only when someone asks for it.

Both facts are now data on every entry, derived and gated against the providers' own answers.

## The schema (`iac/catalogkindpermissions/v1`)

- **`PermissionReach`** (`target`, `organization`, `billing_account`) on every AWS statement and GCP group. AWS `organization` is AWS Organizations, which answers only in the management account.
- **`PermissionPrivilege`** (`ordinary`, `uses_identity`, `grants_access`) on every AWS, GCP and Azure entry.
- **`predefined_role`** on a GCP group a custom role cannot carry: a billing-account reach, or a permission Google publishes as unsupported in custom roles (fleets, deny policies, organization policies).
- **`gcp.api_enabled_by`** and **`gcp.caller_project_services`**:
  - Every kind turns on the APIs it needs. It does that itself, or names in `api_enabled_by` the parent kinds that do (a Cloud SQL database relies on its instance's kind).
  - A resource that lives in no project names in `caller_project_services` the APIs its calls need in the connection's own project (organization policies, folder firewall policies, billing budgets).

## The inventory and the line

- **The GCP fetcher keeps what its three scoped queries say:**
  - which permissions only an organization or a billing account lists;
  - which ones Google will not place in a custom role;
  - the permissions of the predefined roles the manifests name.

  Of the catalog's 1,079 GCP permissions, 1,058 apply to the project, 12 to the organization and 5 to a billing account. Thirteen can't be placed in a custom role.
- **`pkg/iac/actioninventory/privileged.yaml`, written by hand, draws the line once:**
  - *Uses an identity:* `iam.serviceAccounts.actAs`, `iam:PassRole`, Azure's user-assigned identity assign.
  - *Grants access:* IAM policies on projects, folders, organizations, billing accounts and service accounts; roles, users and keys; acting as another identity; Azure role assignments, role definitions and federated credentials.
  - *Ordinary:* granting access on one resource a kind manages, such as a secret's accessor, a bucket's reader or a service's public invoker. This exposes that resource but never makes the principal more powerful.

## A bug the gate would have caught

`GcpSccBigQueryExport` named `securitycenter.BigQueryExports.*`, but Google spells it `bigQueryExports`. A custom role refuses the mis-cased name. The GCP existence gate compared case-insensitively, as IAM action matching does elsewhere, so it let the name through. GCP names are now matched exactly, and the manifest and its controls note are corrected.

## The backfill

- All 574 AWS, GCP and Azure manifests carry both facts. 122 entries use an identity, 20 grant access, 20 reach the organization and 2 a billing account.
- **Two groups were split,** because a group holds one reach:
  - `GcpProject` creates the project on its parent (organization), apart from managing it.
  - Its billing link's two sides are separate groups: the project side, and the billing account side under `roles/billing.user`.
- 24 GCP kinds name the kinds that turn their APIs on, and four name caller-project APIs.

## Gates

- **New gates:**
  - reach against the scope-typed inventory;
  - privilege against `privileged.yaml`;
  - predefined roles, held to exactly the groups that need them, never a basic role, and holding every permission of the group;
  - API ownership;
  - the privileged list's names existing in their inventories;
  - the snapshot's new invariants.

  Each was proven to fail on a deliberately broken manifest.
- **Passing:** `go test` for `pkg/iac/permissions`, `pkg/iac/actioninventory`, `pkg/catalogbundle`, `pkg/skills/...`, `pkg/anatomy`, `pkg/vocabulary`, `pkg/finops/...`, `pkg/compliance/...` and `pkg/specpath/...`.
- **Rules:** the forge, update and audit rules, `architecture/catalog-kind.md`, `CONTRIBUTING.md` and the multi-cloud-catalog skill teach the new fields.
