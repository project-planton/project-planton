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
    env = object({

      # name of the environment
      name = string

      # id of the environment
      id = string
    })

    # labels for the resource
    labels = object({

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

    # The ID of the GCP project where the Managed Zone is created.
    project_id = string

    # An optional list of GCP service accounts that are granted permissions to manage DNS records in the Managed Zone.
    # These accounts are typically workload identities, such as those used by cert-manager,
    # and are added when new environments are created or updated.
    iam_service_accounts = list(string)

    # The DNS records to be added to the Managed Zone.
    records = list(object({

      # Required.** The DNS record type (e.g., A, AAAA, CNAME).
      record_type = string

      # Required.** The name of the DNS record (e.g., "example.com." or "dev.example.com.").
      # This value should always end with a dot to signify a fully qualified domain name.
      name = string

      # The list of values for the DNS record.
      # If the record type is CNAME, each value in the list should end with a dot.
      values = list(string)

      # The Time To Live (TTL) for the DNS record, in seconds.
      ttl_seconds = number
    }))
  })
}
