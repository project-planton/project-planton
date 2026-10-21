---
title: Custom Domain — the Organization's Console at Its Own Address
description: How a Team organization on planton.ai serves its console at a name it owns (planton.acme.com) -- one DNS record, proven from public DNS, certified by Planton, Live in minutes -- the `planton custom-domain` verbs and their MCP twins, every state and cause with its fix, the key scope, and the agent's loop, which adds the record through Planton when the organization already manages that DNS zone as an infra component. Read when a person asks to put their Planton console at planton.<company>, to set up, check, fix or remove the organization's custom domain, why it is pending or needs attention (no record, a proxied or conflicting record, a CAA block), or what scope an agent's key needs. Not this: a service's own hostname (`service.serving-domains-custom.md`).
---

# Custom Domain

Read this when someone says "put our Planton at planton.acme.com", "set up our custom domain", or "why is our custom domain still pending". The people-facing steps are the console's **Settings → Custom Domain**; this file is the agent's version, with the reasons. A service's own address (acme.com serving an app) is a different thing: `service.serving-domains-custom.md`.

## The doctrine in one sentence

An organization on planton.ai serves its console at one name it owns, proven by one DNS record it adds and certified and served by Planton -- and when Planton already manages that name's DNS zone, the record goes in through Planton, never behind its back.

## What it is

- **Hosted Team plan.** Setting one needs the `custom_domain` capability on the organization's plan. Without it the refusal carries reason `capability_required`: the CLI prints "Custom Domain Is on the Team Plan" (exit 1) with the organization's billing page as the next step, and an MCP tool answers an error of category `not_available` whose suggestion starts "Tell the person; don't retry." Tell the person and do not retry. A self-hosted Planton refuses the same way with reason `custom_domain_not_served`: it already runs on its own domain. Re-applying, reading and removing never need the plan, so a lapsed plan never strands a Live address.
- **One per organization, set by its owners.** Only an organization's owners set or remove it; members can read it.
- **A name under the company's domain.** `planton.acme.com`, not `acme.com`: a root domain is refused (set a name under it), so are wildcards and names under Planton's own domains. The name is fixed once set: a different name means `delete`, then `set`. Re-applying the same name means "try again now".
- **One record proves it.** A CNAME from the name to the organization's label under Planton's edge zone (`<org-label>.plantonedge.com` on planton.ai; each Planton environment has its own edge zone, so write exactly the `record.value` the setup guide gives, never one composed from this page). When the organization's own Cloudflare must keep proxying the name, the CNAME stays (proxied) and a TXT record `_planton-verification.<name>` holding the label goes beside it: a proxied record hides where the name points, and the TXT proves it is the organization's. Never the TXT in place of the CNAME -- the CNAME carries the traffic. Planton reads public DNS, so a name nobody pointed here is never certified. A name another organization set moves to the one its DNS points at.
- **Signing in there.** People open the address and sign in through a hand-off from planton.ai into a session confined to that organization. Inviting people and minting credentials open at planton.ai from there.

## Commands and their tool twins

Every command takes the organization from `--org` or the saved context (`planton context get`), and `-o json|yaml`. With no context saved, pass `--org` on every command: setting the context is a change of its own, not needed here. JSON uses the API's field names (snake_case) and enum names (`live`, `record_missing`), byte for byte what the MCP tools answer.

| CLI | MCP twin | What it does |
|---|---|---|
| `planton custom-domain set <host> [--no-wait] [--timeout 15m]` | `apply_custom_domain` | Sets the name (a mutation). Prints the organization's name and slug, the record, and the DNS host's steps. In a terminal it then waits for Live; under `-o json` or without a terminal it never waits. |
| `planton custom-domain get` | `get_custom_domain` | The domain, its record and its recorded state. The tool wraps its answer: `{"custom_domain": {...}}`, or with none `{"custom_domain": null, "next": "apply_custom_domain"}` (the CLI exits 3). The other tools answer the bare record: apply and delete the CustomDomain, verify the verification, the guide the guide. |
| `planton custom-domain verify [--wait] [--timeout 15m]` | `verify_custom_domain` | Checks now: `custom_domain` (the record) beside `reading` (what DNS, the certificate and the console show this moment, with `reading.diagnosis`). A read: it never moves the recorded state, only hurries it when the reading is ahead. The recorded state and the exit code decide: right after Live, `reading` can still show the certificate being issued for a few seconds. |
| `planton custom-domain prompt [--copy]` (`-o json` for the whole guide) | `get_custom_domain_setup_guide`; prompt `setup_custom_domain(org)` (the guide's `agent_prompt`, verbatim) | The setup guide: `record`, `ownership_record`, `dns_host` (`kind`, `display_name`, `zone`, `dashboard_url`, `steps`, `command`) and `agent_prompt`. |
| `planton custom-domain delete [--yes]` | `delete_custom_domain` with `confirm: true` | Removes it (a mutation). The address stops serving the console; everyone keeps it at planton.ai. |

It is also a manifest, so `planton apply -f custom-domain.yaml` sets it like any other record:

```yaml
apiVersion: resource-manager.planton.ai/v1alpha1
kind: CustomDomain
metadata:
  org: acme
spec:
  hostname: planton.acme.com
```

**Exit codes** (`set` that waits, `verify`): **0** the recorded state is Live; **1** refused or failed; **2** not Live yet -- with a `reading.diagnosis` other than `record_missing` or `record_not_yet_public`, act on it; with one of those two, or none, it is still on its way or the wait timed out; **3** the organization has no custom domain. `get` exits 0 whenever a domain exists, whatever its state.

**The key an agent uses** needs the Organization & Members area: `planton api-key new --name <agent> --scope organization_members=read-write` (`=read` for get, verify and the guide alone). A refusal has two causes, worded as such: the key's scope, or the person is not an owner of the organization.

## States and causes

The recorded state is `custom_domain.status.observation.state`: **`pending`** (on its way: the record, the certificate or the edge is not there yet), **`live`** (points at this organization, certified, and the console answered on it), **`needs_attention`** (something definite stands in the way, named by the diagnosis). A diagnosis is the cause, one sentence, and the exact `records_to_add` and `records_to_remove`.

| Cause | What it means | The fix |
|---|---|---|
| `record_missing` | The name's own DNS zone does not serve a record for it yet | Add the CNAME |
| `record_not_yet_public` | The zone serves the record, but public DNS still remembers the name had none: someone looked it up before the record existed. `diagnosis.clears_at` says when that memory ends | Nothing: Planton keeps checking, and it goes Live once public DNS catches up |
| `record_points_elsewhere` | The name's CNAME points somewhere else | Change the CNAME's value to the organization's label |
| `record_conflicts` | An A or AAAA record stands where the CNAME belongs | Remove the address records (`records_to_remove`), add the CNAME |
| `record_proxied` | The name answers with Cloudflare's addresses: the organization's own Cloudflare proxies it or flattens its CNAME | Set the record to DNS only with CNAME flattening off, or keep it proxied and add the TXT beside the CNAME (`records_to_add` holds the TXT; the CNAME stays) |
| `record_at_zone_apex` | The name is its own DNS zone (delegated with NS records), and a CNAME can't sit at a zone's apex | Add the TXT in that zone and point the name at the target with the host's apex flattening, or remove the delegation and add the CNAME in the parent zone |
| `certificate_blocked_by_caa` | A CAA record forbids the certificate authorities Planton uses | Add CAA records allowing `letsencrypt.org` and `pki.goog`, or remove the CAA record |
| `certificate_failed` | The certificate could not be issued | Planton keeps trying; if it stays, `delete` and `set` again |
| `name_blocked` | Planton's edge refuses to serve the name | Contact Planton support with the name; nothing in DNS fixes it |
| `console_not_answering` | Certified, but something in front of it (another proxy or CDN) answers instead of Planton | Remove whatever sits in front of the name |

## The agent's loop

1. **Confirm the organization and the change.** Read first (`planton context get`, `planton custom-domain get -o json`), then say it in one sentence -- "This sets planton.acme.com as the console address of Acme (`acme`)" -- and wait for a yes: the read-then-confirm protocol of `cloud.exploration.md`, one confirmation per mutation. A go-ahead the person already gave for this domain and its DNS record counts: say what you are about to change, then do it, without asking again.
2. **Set it:** `planton custom-domain set planton.acme.com --no-wait -o json` (or `apply_custom_domain`).
3. **Read the guide:** `planton custom-domain prompt -o json` (or `get_custom_domain_setup_guide`): the record, the TXT for a name that stays proxied, and `dns_host.zone`, the zone the name lives in.
4. **Planton-managed DNS first** (below), else the DNS host's own tooling -- either way, show the exact change and add it on the person's yes. The guide's `dns_host.command` names the host's identifiers as `$`-variables; at Cloudflare, `$ZONE_ID` is `curl -s -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" "https://api.cloudflare.com/client/v4/zones?name=<dns_host.zone>" | jq -r '.result[0].id'`.
5. **Verify:** `planton custom-domain verify --wait --timeout 8m -o json` (8 minutes stays inside the time limit a coding agent's shell commands run under; on exit 2 with no diagnosis to act on, run it again), or `verify_custom_domain` every ~30 seconds for up to 15 minutes. Setting the domain and adding the record work in either order: Planton asks the name's own DNS zone before it asks public DNS, so a name set first is Live about a minute and a half after its record is in. What delays it is a lookup made elsewhere: a public resolver asked about the name before its record exists remembers "no such name" for the zone's negative TTL, often 30 minutes or more. So never look the name up yourself (`dig` through your own resolver, `curl`, a browser) before it is Live: `verify` reads it for you.
6. **Act on `reading.diagnosis`:** add `records_to_add`, remove `records_to_remove` (each its own confirmed change), and verify again. Done at exit 0 (the recorded state is `live`): tell the person their console answers at `https://planton.acme.com`.

## Planton-managed DNS first

Before touching DNS, look up whether the organization already manages the name's zone in Planton as an infra component. A record added at the DNS host behind Planton's back makes Planton's record of that zone lie, and the zone's next deploy can remove it (`cloud.exploration.md`, the platform-tracked path).

**Look.** Take the zone from `dns_host.zone` (planton.acme.com usually lives in `acme.com`). List the organization's infra components across environments -- `planton search infra-components --all-envs -o json`, or `list_infra_components` with `kinds` -- and find a zone kind whose zone name is that zone (an `-o json` answer that lists no results -- only `indexFreshness` -- means the organization manages none); read it with `planton get <Kind> <name> -o yaml` (or `get_infra_component`):

| Zone kind | Its zone name | Its records | Record kind (its zone reference) |
|---|---|---|---|
| `CloudflareDnsZone` | `spec.zoneName` | inline `spec.records` | `CloudflareDnsRecord` (`spec.zoneId`) |
| `AwsRoute53Zone` | `metadata.name` | separate kind only | `AwsRoute53DnsRecord` (`spec.zoneId`) |
| `GcpDnsZone` | `spec.dnsName` (trailing dot), else `metadata.name` | separate kind only | `GcpDnsRecord` (`spec.managedZone`) |
| `AzureDnsZone` | `spec.zoneName` | separate kind only | `AzureDnsRecord` (`spec.zoneName`) |
| `DigitalOceanDnsZone` | `spec.domainName` | inline `spec.records` | `DigitalOceanDnsRecord` (`spec.domain`) |

**When one exists**, propose one of two changes, show the manifest, and apply it through Planton on the person's yes:

- **A new DNS-record component in that zone**, its zone reference wired with `valueFrom` to the zone component (`infra.dependencies.md`; add `env` when the zone lives in another environment). The default, because the record gets its own lifecycle and leaves the zone's other records untouched. Ground each kind's fields with `planton explain <Kind>` before writing it.
- **An update to the zone component's records**, when the zone keeps its records inline (`CloudflareDnsZone`, `DigitalOceanDnsZone`) and the person prefers it there: a change to a deployed component, planned and shown like any other.

A CNAME in a Planton-managed Cloudflare zone (`planton validate` passes it):

```yaml
apiVersion: cloudflare.planton.dev/v1alpha1
kind: CloudflareDnsRecord
metadata:
  name: planton-console-cname
spec:
  zoneId:
    valueFrom:
      kind: CloudflareDnsZone
      name: acme-com              # the zone component's metadata.name
      env: shared                 # only when the zone lives in another environment
  name: planton.acme.com
  type: CNAME
  content:
    value: <org-label>.plantonedge.com   # record.value from the setup guide, exactly
  proxied: false                  # DNS only: a proxied CNAME is invisible to public DNS
  ttl: 1
```

Apply it with `planton apply -f <file>` (one confirmation), wait for its deploy to succeed, then verify. The same shape covers the TXT that sits beside a proxied CNAME (`type: TXT`, the ownership record's name and value) and a diagnosis's records.

**Only when no Planton-managed zone exists** use the DNS host's own way: `dns_host.command` (its CLI or API line) when the guide has one and that tool is signed in here, else hand the person `dns_host.steps` and `dns_host.dashboard_url`. Confirm the host first with `dig NS <zone> +short`.

## What not to do

- Never proxy or flatten the record, never touch any other record, and never add a CAA record that leaves out `letsencrypt.org` and `pki.goog`.
- Never look the name up through a public resolver, or open it, before it is Live (`dig NS <zone>` asks about the zone, which exists, and is fine). Waiting on `record_missing` or `record_not_yet_public` means waiting; never add the record twice.
- Never retry a `capability_required` refusal, and never offer the TXT as a way to take a name another organization still stands behind: it claims only a free name.
