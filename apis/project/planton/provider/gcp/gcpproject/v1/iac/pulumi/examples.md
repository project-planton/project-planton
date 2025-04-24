# GcpProject Examples

Below are several YAML examples demonstrating how to create and configure Google Cloud projects using ProjectPlanton's
`GcpProject` resource. After creating a manifest, you can apply it with Pulumi or Terraform via the ProjectPlanton CLI,
just like any other resource in the ProjectPlanton ecosystem.

```shell
# Pulumi
project-planton pulumi up --manifest <yaml-path> --stack <stack-name>

# Terraform
project-planton terraform apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic Example (Organization as Parent)

This example creates a Google Cloud project under an organization ID, links a billing account, and enables one API. Note
that you must specify exactly one of `orgId` or `folderId`.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpProject
metadata:
  name: my-basic-gcp-project
spec:
  projectId: my-basic-12345
  name: My Basic Project
  orgId: "987654321012"
  billingAccountId: "0123AB-4567CD-89EFGH"
  enabledApis:
    - "compute.googleapis.com"
```

---

## Example with Folder as Parent

Instead of an organization, this project will be placed under a folder.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpProject
metadata:
  name: my-folder-gcp-project
spec:
  projectId: my-folder-proj-9876
  name: Folder Parent Project
  folderId: "345678901234"
  billingAccountId: "0123AB-4567CD-89EFGH"
  enabledApis:
    - "storage.googleapis.com"
```

---

## Example with Multiple APIs, Custom Labels, and Default Network Disabled

This project enables multiple Google Cloud APIs, adds metadata labels, and removes the default VPC network immediately
after creation.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpProject
metadata:
  name: multi-api-and-labels
spec:
  projectId: multi-api-labels-1234
  name: Multi API Labels Project
  orgId: "123456789012"
  billingAccountId: "0123AB-4567CD-89EFGH"
  labels:
    env: "dev"
    costCenter: "finops"
  disableDefaultNetwork: true
  enabledApis:
    - "compute.googleapis.com"
    - "iam.googleapis.com"
    - "cloudkms.googleapis.com"
```

---

## Example Granting Owner Role to a Specific Member

This project grants the Owner role to a specified user during creation (for automation or administrative tasks).

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpProject
metadata:
  name: gcp-project-with-owner
spec:
  projectId: with-owner-123
  name: ProjectWithOwner
  orgId: "123456789012"
  ownerMember: "user:devops@example.com"
  enabledApis:
    - "compute.googleapis.com"
```

---

After deploying any of these manifests, you can confirm the newly created project in your GCP account:

```shell
gcloud projects list
gcloud projects describe <your-project-id>
```

You should see the project with the specified parent organization or folder, optional billing account, labels, and
enabled APIs. From there, you can continue configuring additional resources in your GCP environment or integrate your
new project with other components in the ProjectPlanton ecosystem.
