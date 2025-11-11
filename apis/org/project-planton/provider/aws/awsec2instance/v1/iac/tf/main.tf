# Core EC2 instance and related resources for AwsEc2Instance

# IAM Instance Profile attachment is provided via var.spec.iam_instance_profile_arn

resource "aws_instance" "this" {
  ami                    = var.spec.ami_id
  instance_type          = var.spec.instance_type
  subnet_id              = try(var.spec.subnet_id.value, null)
  vpc_security_group_ids = local.security_group_ids_values

  key_name = local.needs_key_name ? try(var.spec.key_name, null) : null

  iam_instance_profile = local.use_ssm ? local.iam_instance_profile_name : null

  ebs_optimized          = try(var.spec.ebs_optimized, null)
  disable_api_termination = try(var.spec.disable_api_termination, null)

  root_block_device {
    volume_size = var.spec.root_volume_size_gb
  }

  user_data = try(var.spec.user_data, null)

  tags = merge({
    Name = var.spec.instance_name
  }, local.safe_tags)
}


