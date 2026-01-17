# OpenFgaAuthorizationModel Local Values
# This file computes local values from the input variables.

locals {
  # Store ID from spec (used in the openfga_authorization_model resource)
  store_id = var.spec.store_id

  # Model JSON from spec (the authorization model definition)
  model_json = var.spec.model_json
}
