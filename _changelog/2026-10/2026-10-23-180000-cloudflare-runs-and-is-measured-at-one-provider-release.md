# Cloudflare Runs, and Is Measured, at One Provider Release

## Why

The catalog's lock moved every OpenTofu module to cloudflare/cloudflare 5.28.0, while the Pulumi modules still bridged 5.23 and the parity check still measured 5.23. Release 5.28.0 renamed the Zero Trust organization's `mfa_ssh_piv_key_requirements` to `mfa_piv_key_requirements`, so every OpenTofu apply of CloudflareZeroTrustOrganization failed with "Unsupported argument" (#610). Separately, Cloudflare stops accepting the DNS settings' `foundation_dns` on 2026-11-23, and CloudflareDnsZone sent it on every zone with `dns_settings` (#584). Nothing caught the rename: the validate job that runs on changes found no changed modules in the lock sweep's pull request, does not run on pushes to main, and the nightly run that did fail reports to nobody.

## What changed in the kinds

- **CloudflareZeroTrustOrganization** names the PIV block `mfa_piv_key_requirements`, accepts the authenticators `piv_key` and `ssh_fido2_key` (Cloudflare retired `ssh_piv_key`), and gains `strict_service_token_auth`, `warp_auth_non_browser_401` and `service_token_inactivity`.
- **CloudflareDnsZone** no longer has `foundation_dns`. Advanced Nameservers are `dns_settings.nameservers.type: cloudflare.advanced`, beside the new `custom` type and `nameserver_set_id`. A `nameservers` block must name its type, as Cloudflare requires. Both modules ignore the retired switch the API still reports, so a zone with settings plans nothing on its second apply (proven live on both engines).
- **CloudflareDnsZone and CloudflareDnsRecord** records gain `include_shadow_metadata`.
- **The Pulumi SDK moves to pulumi-cloudflare v6.23.1**, which bridges 5.28.0, so both engines run one provider. The gaps the older SDK forced are closed: CloudflareWorker's `cache_options`, `exports`, `package_dependencies`, rate-limit `mitigation_timeout` and traces `propagation_policy`; CloudflareEmailRoutingAddress's `status`; CloudflareRuleset's `vary` (whose `default` the spec now requires, as the provider does); and the `cloudflare_account_member` rule on CloudflareZeroTrustAccessGroup and CloudflareZeroTrustAccessPolicy.
- Every Cloudflare module requires `~> 5.28`.

## What changed in the framework

- **Parity is measured at the release the modules run.** `make generate-provider-schemas` takes every version from the canonical lock, and `TestCommittedArtifactsMatchCanonicalLock` fails when a committed schema disagrees with it. Every provider's schema now matches its lock; the arguments those releases added are the parity baseline's migration work list.
- **The OpenTofu validate job covers what it claims.** It runs on pushes to main as well as pull requests, counts a module whose lock moved as changed, fails when it cannot compute the change list, and pushes no longer cancel each other's runs. A red nightly run opens or comments on one tracking issue.
- **Doctrine.** A retired field is deleted and the message renumbered (`architecture/specification-guidelies.md`, "Retiring a Field"); a provider move regenerates the parity schema, validates every module of that provider, and moves the Pulumi SDK to the matching bridge (`catalog/_providers/README.md`, the update rule, the pin guard).
