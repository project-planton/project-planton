---
title: "Custom Domain"
description: "Serve your organization's Planton console at your own address, such as planton.acme.com, with one DNS record"
icon: settings
order: 30
tags:
  - Custom Domain
  - DNS
  - Teams
  - Sign-In
---

# Custom Domain

A custom domain serves your organization's Planton console at an address you own, such as `planton.acme.com`. You set the name, add one DNS record, and Planton proves it, certifies it, and serves your console there, usually within minutes.

## Why It Matters

Your team reaches Planton the way it reaches every other internal tool: at your company's own address, one your people recognize and your network policies can name. Nothing about how Planton runs changes. You keep the same organization, the same people and the same data, and planton.ai keeps working beside it.

## Before You Start

- **The Team plan.** Custom Domain is a Team plan capability on planton.ai. Self-hosted installations already run on a domain of their own.
- **An owner of the organization.** Only an organization's owners set or remove its custom domain. Every member can see it.
- **A name under your domain.** Use a subdomain such as `planton.acme.com`. A root domain (`acme.com`), a wildcard, and names under Planton's own domains are refused, each with a sentence saying what to set instead.
- **Access to your DNS.** You add one record wherever `acme.com`'s DNS is hosted.

## Set It Up

In the console, open **Settings → Custom Domain**, type the name, and save. Planton shows the one record to add, which DNS host holds your zone, and that host's steps.

From the terminal:

```bash
planton custom-domain set planton.acme.com
```

The CLI names the organization it is changing, prints the record and your DNS host's steps, then waits for the domain to go Live. Pass `--org <slug>` to choose an organization other than your current one.

### The DNS Record

Add a **CNAME** from your name to your organization's own label on Planton's edge:

| Type | Name | Value |
|------|------|-------|
| CNAME | `planton.acme.com` | `<your-org-label>.plantonedge.com` |

The exact value is shown in Settings and by the CLI. Keep the record **DNS only**. A proxied or flattened CNAME hides where the name points, so Planton cannot prove it.

If your own Cloudflare must keep proxying the name, keep the CNAME (proxied) and add a **TXT** record beside it: `_planton-verification.planton.acme.com` holding your organization's label. A proxied record hides where the name points; the TXT proves it is yours. Settings and the CLI show both.

Save the name and add the record in either order: Planton asks your zone's own name servers before it asks public DNS, so the domain goes Live about a minute and a half after the record is in. Don't open the address or look it up before then. A public resolver asked about a name before its record exists remembers "no such name" for your zone's negative TTL, often 30 minutes or more. If that happens, Settings says so, with the time it clears, and nothing needs changing.

### When Planton Already Manages Your DNS

If your `acme.com` zone is itself deployed through Planton, as a Cloudflare, Route 53, Cloud DNS, Azure DNS or DigitalOcean DNS zone, add the record through Planton too: a DNS record component in that zone, or a record on the zone component. Planton's record of your zone stays true, and the zone's next deploy keeps the record. A coding agent with the Planton skill does this for you: it finds the zone, shows you the record, and applies it on your go-ahead.

## Check Its Status

```bash
planton custom-domain get       # the domain, its record and its state
planton custom-domain verify    # check DNS, the certificate and the console now
```

**Check Now** in Settings does the same as `verify`. A custom domain is in one of three states:

- **Pending**: on its way. The record, the certificate or the edge isn't there yet.
- **Live**: the name points at your organization, has a certificate, and your console answers on it.
- **Needs Attention**: something definite stands in the way. Planton names the one thing to fix and the exact records to add or remove.

| What Planton Says | The Fix |
|---|---|
| Your zone doesn't serve a record for the name yet | Add the CNAME |
| Your zone serves the record, but public DNS still remembers there was none, until a time it names | Nothing: it goes Live once public DNS catches up |
| The CNAME points somewhere else | Change its value to your organization's label |
| An address record stands where the CNAME belongs | Remove the A/AAAA record and add the CNAME |
| The name answers with Cloudflare's addresses | Set the record to DNS only with CNAME flattening off, or keep it proxied and add the TXT record beside the CNAME |
| The name is a DNS zone of its own | Add the TXT record in that zone and use your host's apex flattening, or remove the delegation and add the CNAME in the parent zone |
| A CAA record forbids Planton's certificate authorities | Add CAA records allowing `letsencrypt.org` and `pki.goog`, or remove the CAA record |
| The certificate could not be issued | Planton keeps trying. If it stays, remove the domain and set it again |
| Planton's edge refuses to serve the name | Contact Planton support with the name |
| Certified, but something else answers | Remove the proxy or CDN in front of the name |

Planton keeps checking a Live domain. One failed check never raises an alarm. When a problem persists, the organization's owners are told what broke and how to fix it.

## Signing In at Your Custom Domain

People open `https://planton.acme.com` and sign in through planton.ai, which hands them into a session confined to your organization. Everyday work happens there. Inviting people and creating credentials (such as API keys) open at planton.ai, where your organization is unchanged. See [Authentication and Authorization](/docs/security/authentication-and-authorization#signing-in-at-a-custom-domain).

## With a Coding Agent

Hand the whole setup to your agent:

```bash
planton custom-domain prompt --copy
```

This copies a prompt with your record, your DNS host, and the steps to finish. **Copy Prompt for Your Agent** in Settings copies the same words. An agent using the [Planton skill](/docs/coding-agents) sets the domain, adds the record (through Planton when Planton manages the zone), and checks until it is Live:

```bash
planton custom-domain set planton.acme.com --no-wait -o json
planton custom-domain verify --wait -o json
```

`verify` exits `0` once the domain is Live, `2` while it isn't yet (with the one thing to fix when there is one), and `3` when the organization has no custom domain. Over MCP, the same steps are `apply_custom_domain`, `verify_custom_domain`, `get_custom_domain`, `get_custom_domain_setup_guide` and `delete_custom_domain`. An agent's API key needs the Organization & Members scope:

```bash
planton api-key new --name my-agent --scope organization_members=read-write
```

## Change or Remove It

A custom domain's name is fixed once set. To use a different name, remove the current one, then set the new one:

```bash
planton custom-domain delete
planton custom-domain set planton-new.acme.com
```

Removing it stops serving your console at that address. Your organization stays at planton.ai.

A custom domain is also a manifest, so `planton apply -f` sets it like any other resource:

```yaml
apiVersion: resource-manager.planton.ai/v1alpha1
kind: CustomDomain
metadata:
  org: acme
spec:
  hostname: planton.acme.com
```

## CLI Reference

| Command | Description |
|---------|-------------|
| `planton custom-domain set <name>` | Set the custom domain and wait for it to go Live (`--no-wait`, `--timeout 15m`) |
| `planton custom-domain get` | Show the custom domain, its record and its state |
| `planton custom-domain verify` | Check it now (`--wait` to wait for Live) |
| `planton custom-domain prompt` | Print the setup prompt for a coding agent (`--copy`) |
| `planton custom-domain delete` | Remove the custom domain |

## Related Documentation

- [Teams and Access](/docs/teams-and-access) — Members, teams and roles
- [Billing](/docs/teams-and-access/billing) — Plans and what each includes
- [Authentication and Authorization](/docs/security/authentication-and-authorization) — How sign-in and permissions work
- [Coding Agents](/docs/coding-agents) — The Planton skill and MCP tools
