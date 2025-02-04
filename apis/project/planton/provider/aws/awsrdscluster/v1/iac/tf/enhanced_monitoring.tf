# Derive a raw name for the Enhanced Monitoring role:
#  1) If user provides spec.enhanced_monitoring_attributes, normalize and join them
#  2) Otherwise, use "<resource_id>-emr"
locals {
  enhanced_monitoring_role_name_pre = (length(var.spec.enhanced_monitoring_attributes) > 0
    ? lower(join(
      "_",
      [
        for attr in var.spec.enhanced_monitoring_attributes :
        regexreplace(attr, "[^a-zA-Z0-9-]", "")
      ]
    ))
    : format("%s-emr", local.resource_id))

  # If the raw name is over 64 characters, truncate and append a short MD5 suffix
  enhanced_monitoring_role_name_final = (length(local.enhanced_monitoring_role_name_pre) > 64
    ?
    substr(local.enhanced_monitoring_role_name_pre, 0, 64 - 5)  + substr(md5(local.enhanced_monitoring_role_name_pre), 0, 5)
    : local.enhanced_monitoring_role_name_pre)
}

# Create assume role policy for Enhanced Monitoring
data "aws_iam_policy_document" "enhanced_monitoring_assume_role" {
  count = var.spec.enhanced_monitoring_role_enabled ? 1 : 0

  statement {
    effect = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = ["monitoring.rds.amazonaws.com"]
    }
  }
}

# Create IAM Role for Enhanced Monitoring
resource "aws_iam_role" "enhanced_monitoring" {
  count              = var.spec.enhanced_monitoring_role_enabled ? 1 : 0
  name               = local.enhanced_monitoring_role_name_final
  assume_role_policy = data.aws_iam_policy_document.enhanced_monitoring_assume_role[0].json
  tags               = local.final_labels
}

# Attach Amazon's managed policy for RDS Enhanced Monitoring
resource "aws_iam_role_policy_attachment" "enhanced_monitoring_policy_attachment" {
  count      = var.spec.enhanced_monitoring_role_enabled ? 1 : 0
  role       = aws_iam_role.enhanced_monitoring[0].name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole"
}
