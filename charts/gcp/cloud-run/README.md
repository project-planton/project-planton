# GCP Cloud Run Quick Start InfraChart

The GCP Cloud Run Quick Start InfraChart provides a streamlined way to rapidly provision all necessary cloud resources
to deploy
your GCP Cloud Run-based service on GCP. This chart specifically focuses on creating infrastructure components essential
for
deploying and managing GCP Cloud Run services, ensuring ease of use and quick onboarding for development teams.

The resources defined by this chart are available in the [templates](templates). Configuration parameters are
managed through [values.yaml](values.yaml).

---

## Included Cloud Resources

This GCP Cloud Run Quick Start Chart creates the following GCP resources:

1. **GCP VPC**:
    - Public subnets across two availability zones
    - DNS hostname and DNS support enabled
    - NAT Gateway enabled for internet access from private subnets

---

## Chart Input Values

The following values must be provided or will default as specified in [values.yaml](values.yaml):

| Input Parameter | Description                                | Example      | Required/Default |
|-----------------|--------------------------------------------|--------------|------------------|
| `org`           | Organization ID on PlantonCloud            | acmecorp     | Required         |
| `env`           | Name of your target deployment environment | dev, staging | Required         |

---

## Chart Customization

Resources created by this GCP Cloud Run Quick Start Chart can be customized post-deployment to fit specific
requirements.
Individual configurations may be modified directly in your GCP account or through Planton Cloud.

---

## Important Notes

- The GCP Cloud Run Quick Start Chart is intended for initial provisioning only. Subsequent changes or management of
  resources
  must be handled individually.
- Verify DNS configurations within Route 53 to ensure seamless ALB and SSL certificate operations.

---

© 2025 Planton Cloud. All rights reserved.

