###############################################################################
# Data Sources
###############################################################################
data "aws_partition" "current" {}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}


###############################################################################
# IAM Role - for Lambda
###############################################################################
# Only create the IAM Role if var.spec.iam_role is non-null (locals.create_iam_role)
resource "aws_iam_role" "lambda" {
  count = local.create_iam_role ? 1 : 0

  name               = "${local.resource_id}-lambda-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role[count.index].json
}

# Trust policy for Lambda to assume this role
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

###############################################################################
# Attach the AWSLambdaBasicExecutionRole policy
###############################################################################
resource "aws_iam_role_policy_attachment" "cloudwatch_logs" {
  count = local.create_iam_role ? 1 : 0

  role       = aws_iam_role.lambda[count.index].name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

###############################################################################
# Optional: VPC Access (AWSLambdaVPCAccessExecutionRole)
###############################################################################
resource "aws_iam_role_policy_attachment" "vpc_access" {
  count = (
  local.create_iam_role && local.lambda_vpc_config != null
  ) ? 1 : 0

  role       = aws_iam_role.lambda[count.index].name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

###############################################################################
# Optional: X-Ray Daemon Write Access
###############################################################################
resource "aws_iam_role_policy_attachment" "xray" {
  count = (
  local.create_iam_role && local.lambda_tracing_config_mode != null && local.lambda_tracing_config_mode != ""
  ) ? 1 : 0

  role       = aws_iam_role.lambda[count.index].name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/AWSXRayDaemonWriteAccess"
}

###############################################################################
# Optional: CloudWatch Lambda Insights
###############################################################################
resource "aws_iam_role_policy_attachment" "cloudwatch_lambda_insights" {
  count = (
  local.create_iam_role && local.iam_role_cloudwatch_lambda_insights_enabled
  ) ? 1 : 0

  role       = aws_iam_role.lambda[count.index].name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/CloudWatchLambdaInsightsExecutionRolePolicy"
}

###############################################################################
# Optional: SSM Parameter Read Access
###############################################################################
data "aws_iam_policy_document" "ssm" {
  count = (
  local.create_iam_role && length(local.iam_role_ssm_parameter_names) > 0
  ) ? 1 : 0

  statement {
    actions = [
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:GetParametersByPath",
    ]
    resources = [
      for param_name in local.iam_role_ssm_parameter_names :
      "arn:${data.aws_partition.current.partition}:ssm:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:parameter${param_name}"
    ]
  }
}

resource "aws_iam_policy" "ssm" {
  count = (
  local.create_iam_role && length(local.iam_role_ssm_parameter_names) > 0
  ) ? 1 : 0

  name        = "${local.resource_id}-ssm-policy-${data.aws_region.current.name}"
  description = "Provides read access to specific SSM parameters for Lambda."
  policy      = data.aws_iam_policy_document.ssm[0].json
  tags        = local.final_tags
}

resource "aws_iam_role_policy_attachment" "ssm_attach" {
  count = (
  local.create_iam_role && length(local.iam_role_ssm_parameter_names) > 0
  ) ? 1 : 0

  role       = aws_iam_role.lambda[count.index].name
  policy_arn = aws_iam_policy.ssm[count.index].arn
}

###############################################################################
# Optional: Custom IAM Policy ARNs
###############################################################################
resource "aws_iam_role_policy_attachment" "custom" {
  count = local.create_iam_role ? length(local.iam_role_custom_iam_policy_arns) : 0

  role = aws_iam_role.lambda[0].name
  policy_arn = local.iam_role_custom_iam_policy_arns[count.index]
}

###############################################################################
# Optional: Inline IAM Policy
###############################################################################
resource "aws_iam_role_policy" "inline" {
  count = (
  local.create_iam_role && local.iam_role_inline_iam_policy != null && local.iam_role_inline_iam_policy != ""
  ) ? 1 : 0

  name   = "${local.resource_id}-inline"
  role   = aws_iam_role.lambda[count.index].name
  policy = local.iam_role_inline_iam_policy
}
