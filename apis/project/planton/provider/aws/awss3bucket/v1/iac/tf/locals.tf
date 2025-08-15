locals {
  bucket_name = var.metadata.name
  is_public   = try(var.spec.is_public, false)
}


