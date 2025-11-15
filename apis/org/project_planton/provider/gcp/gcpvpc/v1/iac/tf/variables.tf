variable "metadata" {
  description = "Metadata for the GCP VPC resource"
  type = object({
    name = string
    id   = string
    org  = optional(string)
    env  = optional(string)
  })
}

variable "spec" {
  description = "Specification for the GCP VPC"
  type = object({
    project_id = object({
      value = string
    })
    auto_create_subnetworks = optional(bool, false)
    routing_mode            = optional(number, 0) # 0=REGIONAL (default), 1=GLOBAL
  })
}
