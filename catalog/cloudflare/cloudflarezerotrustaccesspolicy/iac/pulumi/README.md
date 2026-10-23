# Pulumi Module: Cloudflare Zero Trust Access Policy

Provisions a single `cloudflare.ZeroTrustAccessPolicy` — a reusable, account-scoped
decision plus its access rules, attached to applications by reference.

## Layout

```
iac/pulumi/
├── main.go            # entrypoint (loads iac-input, calls module.Resources)
├── Pulumi.yaml
├── Makefile
└── module/
    ├── main.go            # Resources(): provider setup + policy()
    ├── locals.go          # iac-input references
    ├── policy.go          # the cloudflare.ZeroTrustAccessPolicy + rule mappers
    └── outputs.go         # output constant names
```

## Inputs

A `CloudflareZeroTrustAccessPolicyIacInput` (target + provider config). Required:
`account_id`, `name`, `decision`, and at least one `include` rule.

## Outputs

- `policy_id` — referenced by an application's policies list.

## Requirements

- API token with **Account → Access: Apps and Policies → Edit**.
