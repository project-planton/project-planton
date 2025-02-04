resource "aws_appautoscaling_target" "table_read" {
  count = (
    local.auto_scaling_is_enabled
    ? 1
    : 0
  )

  max_capacity       = coalesce(try(var.spec.auto_scale.read_capacity.max_capacity, null), 20)
  min_capacity       = coalesce(try(var.spec.auto_scale.read_capacity.min_capacity, null), 5)
  resource_id        = "table/${var.metadata.name}"
  scalable_dimension = "dynamodb:table:ReadCapacityUnits"
  service_namespace  = "dynamodb"

  depends_on = [aws_dynamodb_table.this]

  tags = local.final_labels
}

resource "aws_appautoscaling_policy" "table_read" {
  count = (
    local.auto_scaling_is_enabled
    ? 1
    : 0
  )

  name               = "DynamoDBReadCapacityUtilization:${aws_appautoscaling_target.table_read[0].id}"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.table_read[0].resource_id
  scalable_dimension = aws_appautoscaling_target.table_read[0].scalable_dimension
  service_namespace  = aws_appautoscaling_target.table_read[0].service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "DynamoDBReadCapacityUtilization"
    }
    target_value = coalesce(
      try(var.spec.auto_scale.read_capacity.target_utilization, null),
      50
    )
  }

  depends_on = [
    aws_dynamodb_table.this,
    aws_appautoscaling_target.table_read,
  ]
}

resource "aws_appautoscaling_target" "table_write" {
  count = (
    local.auto_scaling_is_enabled
    ? 1
    : 0
  )

  max_capacity       = coalesce(try(var.spec.auto_scale.write_capacity.max_capacity, null), 20)
  min_capacity       = coalesce(try(var.spec.auto_scale.write_capacity.min_capacity, null), 5)
  resource_id        = "table/${var.metadata.name}"
  scalable_dimension = "dynamodb:table:WriteCapacityUnits"
  service_namespace  = "dynamodb"

  depends_on = [aws_dynamodb_table.this]

  tags = local.final_labels
}

resource "aws_appautoscaling_policy" "table_write" {
  count = (
    local.auto_scaling_is_enabled
    ? 1
    : 0
  )

  name               = "DynamoDBWriteCapacityUtilization:${aws_appautoscaling_target.table_write[0].id}"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.table_write[0].resource_id
  scalable_dimension = aws_appautoscaling_target.table_write[0].scalable_dimension
  service_namespace  = aws_appautoscaling_target.table_write[0].service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "DynamoDBWriteCapacityUtilization"
    }
    target_value = coalesce(
      try(var.spec.auto_scale.write_capacity.target_utilization, null),
      50
    )
  }

  depends_on = [
    aws_dynamodb_table.this,
    aws_appautoscaling_target.table_write,
  ]
}


###############################################################################
# GSI Auto Scaling for Read and Write
###############################################################################

resource "aws_appautoscaling_target" "gsi_read" {
  for_each = (
    local.auto_scaling_is_enabled
    ? { for g in try(var.spec.global_secondary_indexes, []) : g.name => g }
    : {}
  )

  max_capacity       = coalesce(try(var.spec.auto_scale.read_capacity.max_capacity, null), 20)
  min_capacity       = coalesce(try(var.spec.auto_scale.read_capacity.min_capacity, null), 5)
  resource_id        = "table/${var.metadata.name}/index/${each.key}"
  scalable_dimension = "dynamodb:index:ReadCapacityUnits"
  service_namespace  = "dynamodb"

  depends_on = [aws_dynamodb_table.this]

  tags = local.final_labels
}

resource "aws_appautoscaling_policy" "gsi_read" {
  for_each = (
    local.auto_scaling_is_enabled
    ? aws_appautoscaling_target.gsi_read
    : {}
  )

  name               = "DynamoDBReadCapacityUtilization:${each.key}"
  policy_type        = "TargetTrackingScaling"
  resource_id        = each.value.resource_id
  scalable_dimension = each.value.scalable_dimension
  service_namespace  = each.value.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "DynamoDBReadCapacityUtilization"
    }
    target_value = coalesce(
      try(var.spec.auto_scale.read_capacity.target_utilization, null),
      50
    )
  }

  depends_on = [
    aws_dynamodb_table.this,
    aws_appautoscaling_target.gsi_read,
  ]
}

resource "aws_appautoscaling_target" "gsi_write" {
  for_each = (
    local.auto_scaling_is_enabled
    ? { for g in try(var.spec.global_secondary_indexes, []) : g.name => g }
    : {}
  )

  max_capacity       = coalesce(try(var.spec.auto_scale.write_capacity.max_capacity, null), 20)
  min_capacity       = coalesce(try(var.spec.auto_scale.write_capacity.min_capacity, null), 5)
  resource_id        = "table/${var.metadata.name}/index/${each.key}"
  scalable_dimension = "dynamodb:index:WriteCapacityUnits"
  service_namespace  = "dynamodb"

  depends_on = [aws_dynamodb_table.this]

  tags = local.final_labels
}

resource "aws_appautoscaling_policy" "gsi_write" {
  for_each = (
    local.auto_scaling_is_enabled
    ? aws_appautoscaling_target.gsi_write
    : {}
  )

  name               = "DynamoDBWriteCapacityUtilization:${each.key}"
  policy_type        = "TargetTrackingScaling"
  resource_id        = each.value.resource_id
  scalable_dimension = each.value.scalable_dimension
  service_namespace  = each.value.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "DynamoDBWriteCapacityUtilization"
    }
    target_value = coalesce(
      try(var.spec.auto_scale.write_capacity.target_utilization, null),
      50
    )
  }

  depends_on = [
    aws_dynamodb_table.this,
    aws_appautoscaling_target.gsi_write,
  ]
}
