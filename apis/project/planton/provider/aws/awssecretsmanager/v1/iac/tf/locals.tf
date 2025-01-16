locals {
  aws_tags = {
    Resource      = "true"
    Organization  = var.metadata.org
    Environment   = var.metadata.env.id
    ResourceKind  = "aws-secrets-manager"
    ResourceId    = var.metadata.id
  }
}
