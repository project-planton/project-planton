
locals {

  # replace '.' with '-' in hosted zone name

  normalized_r53_zone_name = replace(var.metadata.name, ".", "-")

  # Normalize record names to produce stable keys for for_each.
  # 1. trim starting and trailing dot
  # 2. replace '.' with '_'
  # 3. replace '*' with 'wildcard'

  normalized_records = {
    for rec in local.records :
    format(
      "%s_%s",
      replace(
        replace(
          trim(rec.name, "."),
          ".", "_"
        ),
        "*", "wildcard"
      ),
      rec.record_type
    ) => rec
  }

  records = var.spec.records

}
