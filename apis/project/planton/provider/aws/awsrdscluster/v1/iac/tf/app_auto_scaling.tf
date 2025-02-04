resource "aws_appautoscaling_target" "replicas" {
  count              = var.spec.auto_scaling.is_enabled ? 1 : 0
  service_namespace  = "rds"
  resource_id        = "cluster:${aws_rds_cluster.this[0].id}"
  scalable_dimension = "rds:cluster:ReadReplicaCount"
  min_capacity       = (try(var.spec.auto_scaling.min_capacity, 0) > 0 ? var.spec.auto_scaling.min_capacity : 1)
  max_capacity       = (try(var.spec.auto_scaling.max_capacity, 0) > 0 ? var.spec.auto_scaling.max_capacity : 5)

  depends_on = [
    aws_rds_cluster.this
  ]
}

resource "aws_appautoscaling_policy" "replicas_policy" {
  count = var.spec.auto_scaling.is_enabled ? 1 : 0

  name               = local.resource_id
  policy_type        = (try(var.spec.auto_scaling.policy_type, "") != "" ? var.spec.auto_scaling.policy_type :
    "TargetTrackingScaling")
  resource_id        = aws_appautoscaling_target.replicas[0].resource_id
  scalable_dimension = aws_appautoscaling_target.replicas[0].scalable_dimension
  service_namespace  = aws_appautoscaling_target.replicas[0].service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = (try(var.spec.auto_scaling.target_metrics, "") != "" ?
        var.spec.auto_scaling.target_metrics : "RDSReaderAverageCPUUtilization")
    }
    target_value       = (try(var.spec.auto_scaling.target_value, 0) > 0 ? var.spec.auto_scaling.target_value : 75)
    scale_in_cooldown  = (try(var.spec.auto_scaling.scale_in_cooldown, 0) > 0 ? var.spec.auto_scaling.scale_in_cooldown
      : 300)
    scale_out_cooldown = (try(var.spec.auto_scaling.scale_out_cooldown, 0) > 0 ?
      var.spec.auto_scaling.scale_out_cooldown : 300)
    disable_scale_in   = false
  }

  depends_on = [
    aws_appautoscaling_target.replicas
  ]
}
