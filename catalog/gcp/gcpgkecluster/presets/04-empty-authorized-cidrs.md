# Empty Authorized CIDR List

This regional Standard cluster configures an authorized-networks policy without
any external CIDR ranges. It enables the IAM-authenticated DNS endpoint as a
separate administrative access path and uses private nodes.

## When to Use

Use this when administrators access the cluster through the DNS endpoint and do
not need an office or VPN range in the IP endpoint's allowlist. Confirm the
administrators' IAM permissions before applying the policy.

## Key Configuration Choices

- **Configured empty policy:** `masterAuthorizedNetworks: {}` remains a restriction
  even with no CIDR entries. Omitting it would configure no authorized-networks
  restriction; `unrestrictedNetworks: {}` states that unrestricted choice explicitly.
- **Independent DNS access:** DNS endpoint access requires IAM authentication and
  does not depend on the IP endpoint's CIDR list.
- **Private nodes:** node egress requires an appropriate network path, such as Cloud
  NAT for public registries. Add a `GcpGkeNodePool` to supply compute.
- **Separate policy switches:** Google public-IP access and private-endpoint
  enforcement retain their catalog defaults unless explicitly configured.

## Placeholders to Replace

| Placeholder | Meaning |
|---|---|
| `my-gcp-project-123` | Your GCP project ID |
| `my-app-vpc` | Your `GcpVpcNetwork` resource name |
| `my-gke-subnet` | Your `GcpSubnetwork` resource name |
