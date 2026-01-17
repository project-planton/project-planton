# OpenFgaRelationshipTuple Main Resources
# This file creates the OpenFGA relationship tuple resource.
#
# A relationship tuple represents a relationship between a user and an object
# through a specific relation. Together with an authorization model, tuples
# determine access decisions.
#
# Note: Relationship tuples are immutable. Changing any field will result in
# Terraform deleting the old tuple and creating a new one.
#
# Reference: https://registry.terraform.io/providers/openfga/openfga/latest/docs/resources/relationship_tuple

resource "openfga_relationship_tuple" "this" {
  store_id               = local.store_id
  authorization_model_id = local.authorization_model_id
  user                   = local.user
  relation               = local.relation
  object                 = local.object

  # Condition block (optional)
  dynamic "condition" {
    for_each = local.condition != null ? [local.condition] : []
    content {
      name         = condition.value.name
      context_json = condition.value.context_json
    }
  }
}
