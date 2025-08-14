# EKS Environment InfraChart

This chart provisions a **complete, production‑ready Kubernetes environment on AWS**:

* Custom VPC across two AZs with public / private subnets and NAT
* IAM roles for control‑plane and nodes
* Optional customer‑managed KMS key for secrets encryption
* Private or restricted API endpoint with CloudWatch control‑plane logs
* Managed node group with autoscaling, Spot or On‑Demand instances
* Optional Route 53 public zone
* Toggleable Kubernetes add‑ons (Cert‑Manager, External‑DNS, Istio, etc.)

Edit **values.yaml** to tailor the deployment; each `*Enabled` boolean cleanly removes its add‑on.

© 2025 Planton Cloud. All rights reserved.
