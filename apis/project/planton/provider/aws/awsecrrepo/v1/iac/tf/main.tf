resource "aws_ecr_repository" "this" {
  name                 = var.spec.repository_name
  image_tag_mutability = local.image_tag_mutability
  force_delete         = try(var.spec.force_delete, false)

  encryption_configuration {
    encryption_type = local.encryption_type
    kms_key         = local.kms_key_id
  }

  tags = local.tags
}


