# dns_settings.tf

locals {
  # Unset numeric/string leaves arrive as the proto zero values (0 / ""); the
  # provider's validators reject them (e.g. ns_ttl must be 30-86400), so every
  # optional leaf is sent as null unless the user actually set it -- Cloudflare
  # then keeps its default. Mirrors the Pulumi module's presence handling.
  dns_settings_soa = try(var.spec.dns_settings.soa, null) == null ? null : {
    expire  = var.spec.dns_settings.soa.expire > 0 ? var.spec.dns_settings.soa.expire : null
    min_ttl = var.spec.dns_settings.soa.min_ttl > 0 ? var.spec.dns_settings.soa.min_ttl : null
    mname   = var.spec.dns_settings.soa.mname != "" ? var.spec.dns_settings.soa.mname : null
    refresh = var.spec.dns_settings.soa.refresh > 0 ? var.spec.dns_settings.soa.refresh : null
    retry   = var.spec.dns_settings.soa.retry > 0 ? var.spec.dns_settings.soa.retry : null
    rname   = var.spec.dns_settings.soa.rname != "" ? var.spec.dns_settings.soa.rname : null
    ttl     = var.spec.dns_settings.soa.ttl > 0 ? var.spec.dns_settings.soa.ttl : null
  }

  # type is required whenever the block is sent (the spec's rule holds it).
  dns_settings_nameservers = try(var.spec.dns_settings.nameservers, null) == null ? null : {
    type              = var.spec.dns_settings.nameservers.type
    ns_set            = var.spec.dns_settings.nameservers.ns_set > 0 ? var.spec.dns_settings.nameservers.ns_set : null
    nameserver_set_id = var.spec.dns_settings.nameservers.nameserver_set_id != "" ? var.spec.dns_settings.nameservers.nameserver_set_id : null
  }

  dns_settings_internal_dns = (
    try(var.spec.dns_settings.internal_dns.reference_zone_id, "") != ""
    ? { reference_zone_id = var.spec.dns_settings.internal_dns.reference_zone_id }
    : null
  )
}

# Zone-wide DNS settings, provisioned only when the spec supplies a dns_settings block.
# The three booleans are always sent: the provider marks them Optional and not
# Computed, so leaving one unset while the zone holds a value would plan its
# removal on every refresh. ns_ttl has the same shape and the server echoes
# its default (86400) into state, which is why the spec asks a drift-free
# manifest to declare it. Advanced Nameservers are chosen through
# nameservers.type ("cloudflare.advanced"), the only switch Cloudflare accepts
# for them.
resource "cloudflare_zone_dns_settings" "main" {
  count   = local.has_dns_settings ? 1 : 0
  zone_id = cloudflare_zone.main.id

  flatten_all_cnames  = var.spec.dns_settings.flatten_all_cnames
  multi_provider      = var.spec.dns_settings.multi_provider
  secondary_overrides = var.spec.dns_settings.secondary_overrides
  ns_ttl              = var.spec.dns_settings.ns_ttl > 0 ? var.spec.dns_settings.ns_ttl : null
  zone_mode           = local.zone_mode

  soa          = local.dns_settings_soa
  nameservers  = local.dns_settings_nameservers
  internal_dns = local.dns_settings_internal_dns

  # foundation_dns is a retired switch the module never sends (Advanced
  # Nameservers are nameservers.type). The API still reports it, and the
  # provider marks it Optional and not Computed, so without this every refresh
  # would plan to null the reported value. When a provider release drops the
  # attribute, validation names this line and it goes.
  lifecycle {
    ignore_changes = [foundation_dns]
  }
}
