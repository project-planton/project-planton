resource "google_dns_managed_zone" "managed_zone" {
  name        = local.managed_zone_name
  project     = var.spec.project_id
  description = "managed-zone for ${var.metadata.name}"
  dns_name    = local.zone_dns_name
  visibility  = "public"
}

resource "google_project_iam_binding" "dns_admin" {
  count   = length(local.iam_binding_members) > 0 ? 1 : 0
  project = var.spec.project_id
  role    = "roles/dns.admin"
  members = local.iam_binding_members

  depends_on = [
    google_dns_managed_zone.managed_zone
  ]
}

resource "google_dns_record_set" "records" {
  for_each = {for idx, rec in var.spec.records : idx => rec}

  managed_zone = google_dns_managed_zone.managed_zone.name
  project      = var.spec.project_id
  name         = each.value.name
  type         = each.value.record_type
  ttl          = each.value.ttl_seconds
  rrdatas      = each.value.values

  depends_on = [
    google_dns_managed_zone.managed_zone
  ]
}
