# OpenFgaAuthorizationModel Variables
# This file defines all input variables for the OpenFgaAuthorizationModel Terraform module.
# These variables map to the OpenFgaAuthorizationModelSpec protobuf message.

variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name    = string
    id      = optional(string)
    org     = optional(string)
    env     = optional(string)
    labels  = optional(map(string))
    tags    = optional(list(string))
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "OpenFgaAuthorizationModel specification"
  type = object({
    # store_id is the unique identifier of the OpenFGA store where this model will be created.
    # The store must exist before creating an authorization model.
    # Note: Changing the store_id requires replacing the model.
    store_id = string

    # model_json is the authorization model definition in JSON format.
    # The JSON must conform to the OpenFGA authorization model schema.
    # Note: Changing the model requires replacing it (new model ID).
    model_json = string
  })
}
