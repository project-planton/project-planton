############################################################
# Variables that mirror the provided tfvars
############################################################

variable "apiVersion" {
  type        = string
  description = "The API version for this resource."
}

variable "kind" {
  type        = string
  description = "The kind of resource (e.g., AwsDynamodb)."
}

variable "metadata" {
  type = object({
    env = object({
      name = string
      id   = string
    })
    version = object({
      id      = string
      message = string
    })
    name = string
    id   = string
    org  = string
  })
  description = <<-EOT
Metadata related to this resource, including:
- env: an object containing environment name and id
- version: an object containing version id and message
- name: a string representing the resource name
- id: a string representing the resource id
- org: a string representing the organization name
EOT
}

variable "spec" {
  type = object({
    billingMode = string
    hashKey = object({
      name = string
      type = string
    })
  })
  description = <<-EOT
Specification for the resource, including:
- billingMode: The billing mode for the DynamoDB table (e.g., PROVISIONED or PAY_PER_REQUEST)
- hashKey: An object defining the hash key for the table, including:
  - name: The name of the hash key attribute
  - type: The type of the attribute (e.g., "S" for string)
EOT
}
