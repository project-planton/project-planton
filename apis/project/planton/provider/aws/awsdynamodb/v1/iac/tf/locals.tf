locals {
  safe_stream_enabled = try(var.spec.stream_enabled, false)
}


