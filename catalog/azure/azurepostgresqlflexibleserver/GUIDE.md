# Azure PostgreSQL Flexible Server -- Operational Guide

Judgment calls that matter when you run PostgreSQL flexible servers in production.

## Zones are placement, decided at creation

`zone` and `highAvailability.standbyAvailabilityZone` place the primary and the standby when the server is created, and that is the only time they are applied. Leave them unset unless you have a reason: Azure picks a zone with capacity for the SKU, and for ZONE_REDUNDANT high availability it puts the standby in a different zone on its own. Set them when placement must line up with something else that is zonal -- an AKS node pool or a VM in the same zone keeps the app-to-database hop inside one zone.

## Later edits are ignored on purpose

Both engines ignore later changes to the two zone fields. Editing them after creation does nothing -- no error, no replacement, no failover. Two reasons, both live-relevant:

- **An unset zone is still a zone.** Azure assigns one at creation, and the provider records it. Without the guard, the next plan -- for a tag edit, say -- would try to clear it, and Azure refuses every zone change that is not a swap with the standby.
- **Azure swaps the zones when it fails over.** After an automatic failover, the primary runs in what was the standby's zone. Re-sending the creation-time pair would read as a swap request, and the provider would run a planned failover -- a short outage triggered by an unrelated edit.

So the zones in your manifest describe where the server was created, not where it runs today. Ask Azure for the live placement:

```
az postgres flexible-server show --resource-group <rg> --name <server> \
  --query "{primary: availabilityZone, standby: highAvailability.standbyAvailabilityZone}"
```

## Failing over on purpose is an Azure operation

To move the primary back to its original zone after an automatic failover, or to rehearse one, run a planned failover yourself -- it waits for the standby to catch up, so no committed data is lost:

```
az postgres flexible-server restart --resource-group <rg> --name <server> --failover Planned
```

Expect a short connection drop, as with any failover; run it in a quiet window and make sure clients reconnect. Nothing in Planton needs to change afterwards, because the manifest's zones are never re-applied.

## Moving to different zones means a new server

A server cannot move to zones outside its current primary/standby pair in place. Create a new server with the placement you want -- a point-in-time restore (`createMode: POINT_IN_TIME_RESTORE`) into the new zone is the low-effort path -- cut clients over, then delete the old one.
