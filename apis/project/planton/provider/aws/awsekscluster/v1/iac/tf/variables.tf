variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # The AWS region in which to create the EKS cluster.
    # This must be a valid AWS region where EKS is available.
    # Note: The EKS cluster will be recreated if this value is updated.
    # For a list of AWS regions, see: https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
    region = string

    # Security Groups for the EKS cluster
    security_groups = list(string)

    # Subnets for the EKS cluster
    subnets = list(string)

    # role arn for the EKS cluster
    role_arn = string

    # Worker Node Role ARN
    node_role_arn = string

    # Instance type for the EKS worker nodes
    instance_type = string

    # Desired size of the EKS worker node group
    desired_size = number

    # Maximum size of the EKS worker node group
    max_size = number

    # Minimum size of the EKS worker node group
    min_size = number

    # Description for tags
    tags = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })
  })
}