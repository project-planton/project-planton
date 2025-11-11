variable "metadata" {
  description = "metadata"
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = string

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # annotations for the resource
    annotations = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # tags for the resource
    tags = list(string)
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # Flag to indicate if the S3 bucket should have external (public) access.
    # When set to `true`, the bucket will be accessible publicly over the internet,
    # allowing anyone to access the objects stored within it.
    # When set to `false` (default), the bucket is private, and access is restricted
    # based on AWS Identity and Access Management (IAM) policies and bucket policies.
    # Public access should be used cautiously to avoid unintended data exposure.
    is_public = bool

    # The AWS region where the S3 bucket will be created.
    # This must be a valid AWS region where S3 is available.
    # Specifying the region is important because it affects data latency and costs.
    # For a list of AWS regions, see: https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
    aws_region = string
  })
}