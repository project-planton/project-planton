# AwsClientVpn

AWS Client VPN endpoint provisioning for secure remote access into a VPC using OpenVPN. This resource sets up a Client VPN endpoint, target subnet associations, basic authorization rules, and optional connection logging.

## Spec fields (summary)
- **vpc_id**: VPC to attach the Client VPN to. Required.
- **subnets**: One or more subnet IDs in the VPC to associate as target networks. At least 1 required.
- **client_cidr_block**: IPv4 CIDR for client IP allocation (e.g., 10.0.0.0/22). Required.
- **authentication_type**: Authentication method. Certificate-based only in v1.
- **server_certificate_arn**: ACM certificate ARN for the VPN server. Required.
- **cidr_authorization_rules**: CIDR ranges clients are allowed to access through the VPN.
- **disable_split_tunnel**: When true, route all client traffic through VPN. Default false.
- **vpn_port**: Listener port. Defaults to 443. Allowed: 443 (TCP), 1194 (UDP).
- **transport_protocol**: TCP or UDP. Defaults to TCP.
- **log_group_name**: CloudWatch Logs group for connection logs (optional).
- **security_groups**: Security groups to attach to target network associations (optional).
- **dns_servers**: Up to two custom DNS IPs for clients (optional).

## Stack outputs
- **client_vpn_endpoint_id**: The Client VPN endpoint ID.
- **security_group_id**: Security group applied to associations.
- **subnet_association_ids**: Map of subnet ID -> association ID.

## How it works
This repo supports Pulumi and Terraform backends. The Pulumi module uses `aws.ec2clientvpn` resources to create the endpoint, associations, and rules, and exports observable outputs for downstream references.

## References
- AWS Client VPN: `https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html`
- Pulumi AWS (Client VPN): `https://www.pulumi.com/registry/packages/aws/api-docs/ec2clientvpn/`


