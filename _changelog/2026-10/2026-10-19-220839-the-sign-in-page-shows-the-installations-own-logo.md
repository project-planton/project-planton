# The Sign-In Page Shows the Installation's Own Logo

**Date**: October 19, 2026

## Summary

A self-hosted installation's sign-in page now draws the installation's own logo in place of Planton's mark, read from the installation itself. The console gains the control plane's in-cluster address for its own server, so the screens it draws before anyone signs in can carry the logo too. The Planton skill teaches agents to put a company's logo on Planton with `planton branding`.

## Problem Statement

The identity server's theme carried Planton's mark as a file of its own, so a company running Planton could not see its logo on the page its people open every morning. Changing it would have meant a theme change and an identity-server restart.

### Pain Points

- The sign-in page was the one self-hosted surface no product setting could reach.
- The console's server had only the public address of the control plane, which need not resolve or hairpin from inside the cluster.
- The skill taught that every self-hosted change is an edit to the declaration, which is not true of what the product itself holds.

## Solution

- **The theme reads the logo from the installation.** `--keycloak-logo-url` points at `/storage/v1/branding/sign-in.png`, a path-absolute address on the same front door as the page. The control plane answers it with the dark-theme version of the installation's logo, else the logo, else Planton's mark (the page is dark by design). The theme ships no image; nothing is fetched from a third party, and a logo change needs no theme change and no restart.
- **`API_INTERNAL_ENDPOINT`.** The console receives the control plane's gRPC-Web door at its in-cluster address, always, beside the browser's `API_ENDPOINT`: the same split horizon as `IAM_ISSUER_URL` / `IAM_ISSUER_INTERNAL_URL`. One helper, `ControlPlaneBrowserAPIInternalOrigin`, spells that in-cluster door for the storage relay and the console alike.
- **The skill.** A new `self-hosted.branding.md` reference, the logo verbs in the CLI craft, a corrected doctrine (the declaration is the handle for what the operator runs; directory mappings and the logo are held by the product), and an eval question.

## Impact

- Installations whose platform serves `sign-in.png` show their own logo on the sign-in page; one that does not yet serves it shows an empty logo area until its platform is upgraded.
- The console pod gains one environment variable; the boot contract fixture is refreshed.
- Agents answer "put our logo on Planton" with one command instead of an edit to the declaration.
