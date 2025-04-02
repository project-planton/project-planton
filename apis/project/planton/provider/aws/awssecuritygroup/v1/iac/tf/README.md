# AWS Security Group Terraform Module

This Terraform module creates an AWS Security Group. It is designed to map closely to the Protobuf definitions for
`AwsSecurityGroupSpec` and its nested fields (`SecurityGroupRule`).

## Inputs

| Name          | Description                                                                      | Type                                                                                                                                                                                                                                                                                                                                       | Default | Required |
|---------------|----------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|----------|
| `name`        | Name of the Security Group.                                                      | `string`                                                                                                                                                                                                                                                                                                                                   | n/a     | yes      |
| `vpc_id`      | ID of the VPC where this Security Group will be created.                         | `string`                                                                                                                                                                                                                                                                                                                                   | n/a     | yes      |
| `description` | Short explanation of this Security Group’s purpose (<= 255 chars).               | `string`                                                                                                                                                                                                                                                                                                                                   | n/a     | yes      |
| `ingress`     | List of ingress rules, each defined as an object mapping to `SecurityGroupRule`. | <code>list(object({<br/> protocol = string,<br/> from_port = number,<br/> to_port = number,<br/> ipv4_cidrs = list(string),<br/> ipv6_cidrs = list(string),<br/> source_security_group_ids = list(string),<br/> destination_security_group_ids = list(string),<br/> self_reference = bool,<br/> rule_description = string,<br/> }))</code> | `[]`    | no       |
| `egress`      | List of egress rules, each defined as an object mapping to `SecurityGroupRule`.  | Same type as `ingress`.                                                                                                                                                                                                                                                                                                                    | `[]`    | no       |
| `tags`        | Additional tags to apply to the Security Group.                                  | `map(string)`                                                                                                                                                                                                                                                                                                                              | `{}`    | no       |

## Outputs

| Name                  | Description                                                 |
|-----------------------|-------------------------------------------------------------|
| `security_group_id`   | The ID of the newly created Security Group.                 |
| `vpc_id`              | The VPC ID in which this Security Group is created.         |
| `internet_gateway_id` | Placeholder output, returns an empty string in this module. |
| `private_subnets`     | Placeholder output, returns an empty list in this module.   |
| `public_subnets`      | Placeholder output, returns an empty list in this module.   |

## Notes

• This module does **not** manage AWS credentials or provider configurations. It expects them to be handled
externally.  
• The input structure is intentionally similar to the Protobuf specification (`AwsSecurityGroupSpec`), ensuring
consistency with the rest of ProjectPlanton.
