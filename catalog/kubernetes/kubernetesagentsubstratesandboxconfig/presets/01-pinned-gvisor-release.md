# Pinned gVisor Release

One gVisor release, pinned to the byte, that your worker pools and actor templates name explicitly. It is the release Agent Substrate v0.5.0 ships as `gvisor-default` (gVisor nightly 2026-10-10, the build with the filesystem checkpoint Substrate's durable directories need), so the config starts out equal to the shipped one and stays where you put it when a later control plane ships a newer release.

## When to Use

- You validate a gVisor release before your agents run on it, and want it to stay put across control plane upgrades.
- A worker pool per tenant, or per environment, should run a runtime you chose rather than the class default.

## How It Works

Every node downloads the tarball for its own architecture ahead of time, verifies it against its SHA-256 and caches it by hash. Worker pools that name this config (`sandboxClasses[].configRef.name`) keep it from deletion, and the actors of templates that name it (`sandbox_config.config_name`) boot and restore with its default version. The pause image holds each sandbox's namespaces and is pinned by digest, as the cluster requires.

## Key Configuration Choices

- **A version named after what it pins.** `gvisor-2026-10-10` reads as a fact in a template or a status line; a counter does not.
- **Both architectures.** Keep a block for every architecture your workers run on; a version with no files for a node's architecture cannot boot there.
- **No class-default annotation.** This config is named explicitly; `gvisor-default` stays what pools that name no config resolve. Templates always name their config.

## Prerequisites

- **Agent Substrate** on the cluster.
- Worker nodes with internet egress to `gs://gvisor` (read anonymously), or the tarballs mirrored to your own bucket and the URLs changed to it.

## Placeholders to Replace

| Placeholder | Description |
|-------------|-------------|
| `gvisor-pinned` (name) | The config's name in the cluster, which worker pools and actor templates name; rename it freely (a DNS subdomain name). |
