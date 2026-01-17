# OpenFgaRelationshipTuple Local Values
# This file computes local values from the input variables.

locals {
  # Store ID from spec (used in the openfga_relationship_tuple resource)
  store_id = var.spec.store_id

  # Authorization model ID from spec (optional, uses latest if not specified)
  authorization_model_id = var.spec.authorization_model_id

  # User, relation, and object - the core tuple fields
  user     = var.spec.user
  relation = var.spec.relation
  object   = var.spec.object

  # Condition (optional)
  condition = var.spec.condition
}
