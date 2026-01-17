# OpenFgaAuthorizationModel Main Resources
# This file creates the OpenFGA authorization model resource.
#
# An authorization model defines the types, relations, and access rules for
# fine-grained authorization within an OpenFGA store.
#
# Note: Authorization models are immutable. Changing the model_json will
# create a new model (new ID) rather than updating the existing one.
#
# Reference: https://registry.terraform.io/providers/openfga/openfga/latest/docs/resources/authorization_model

resource "openfga_authorization_model" "this" {
  store_id   = local.store_id
  model_json = local.model_json
}
