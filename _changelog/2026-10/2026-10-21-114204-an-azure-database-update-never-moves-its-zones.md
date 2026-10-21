# An Azure Database Update Never Moves Its Zones

## PostgreSQL: any update after the first deploy failed

An AzurePostgresqlFlexibleServer deployed without a `zone` could not be
updated at all. Adding one tag failed with Azure's "`zone` can only be changed
when exchanged with the zone specified in
`high_availability.0.standby_availability_zone`".

azurerm models the server's `zone` and `standby_availability_zone` as plain
optional fields (`commonschema.ZoneSingleOptional()`, not computed). Azure
assigns a zone at creation and Read stores it in state. The next plan then
compares "unset" with Azure's zone and tries to clear it, which Azure refuses.
The least-privilege lane found it on dev, on the OpenTofu engine.

## PostgreSQL and MySQL: an unrelated edit could fail the database over

When Azure fails over on its own, primary and standby swap zones. A manifest
that names both zones then sends the creation-time pair on every later update.
Both providers read that as a swap and run a planned failover, so a tag edit
caused a short outage. MySQL's zone fields are computed, so MySQL did not have
the bug above, but it did have this one.

## The contract: zones are placement chosen at creation

Both kinds now ignore later changes to `zone` and
`high_availability.standby_availability_zone` on both engines: Pulumi uses
`IgnoreChanges` and OpenTofu uses `lifecycle.ignore_changes`. This is the same
create-time-only contract `AzureDiskSnapshot` uses for its source. Both are
sent at create, and editing them later does nothing: no error, no
replacement, no failover. Zones move only through an Azure failover, which is
an Azure operation and never a side effect of a spec edit.

- The spec field comments say this.
- Each kind has a new `GUIDE.md` covering placement, what a failover does to
  the zones, how to fail over on purpose, and how to move to new zones (a new
  server).
- The kind and IaC READMEs carry the rule.
