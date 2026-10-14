---
title: "Flow Control"
description: "Governance policies that control how Infra Jobs execute — manual approvals, preview requirements, and deployment pauses."
icon: guide
order: 60
tags:
  - Infrastructure
  - Flow Control
  - Governance
---

# Flow Control

A Flow Control policy defines the rules for how [Infra Jobs](/docs/infrastructure/infra-jobs) execute. It controls whether jobs require manual approval, whether a preview step runs before every apply, and whether the system pauses between preview and apply to give you a chance to review the plan. Different environments need different levels of rigor — development should be fast, production should be careful — and Flow Control makes that distinction automatic.

## Why Flow Control Exists

Without governance, every infrastructure change follows the same path: submit configuration, run the IaC operations, apply the result. That works for development, but production deserves more care. You want to see the plan before it runs. You want someone to approve destructive changes.

Flow Control policies encode these decisions into the platform. Define the rules once, attach them to the right scope, and every Infra Job in that scope automatically follows them. No per-deployment configuration, no manual steps to remember, no governance gaps.

## The Four Controls

Every Flow Control policy consists of four settings. Each is a boolean toggle that modifies how Infra Jobs execute.

### Manual Approval Required

When enabled, every Infra Job created in this scope pauses immediately after creation and waits for manual approval before any operations begin. Nothing executes until a human explicitly approves it.

This differs from "Pause After Preview" (below), which allows the preview to run automatically and only pauses before the apply step. Manual Approval Required pauses before any operation starts.

### Lifecycle Events Disabled

When enabled, Infra Jobs are not automatically created when an Infra Component is created or updated: the resource is unmanaged, and Planton records it without deploying anything. You can still create Infra Jobs manually through the CLI or web console, but the automatic trigger from lifecycle events is suppressed. On a single resource, set it with the annotation `planton.dev/infra-job.disable-on-lifecycle-events: "true"`.

Inside an Infra Stack or a Service deployment, such a resource reads **Unmanaged**: the run records it and never applies or destroys it. A stack's plan reads it Unmanaged; removing it from the templates plans **Keep** (named by its provider: Keep in AWS), with the choice locked; a stack's undeploy or purge releases it and leaves it in the cloud; a service's deletion leaves it in the cloud too. Tearing it down is a deliberate act on that Infra Component alone.

This is useful for shared infrastructure that multiple teams depend on, where accidental changes could have wide impact. Only explicit, intentional Infra Jobs run.

### Preview Before Apply

When enabled, a preview step always runs before update or destroy operations. The preview shows exactly what will change — resources to create, modify, or delete — before anything happens. This is enabled by default.

Disabling it removes the preview step, making Infra Jobs faster but removing the safety net of seeing the plan first. Only disable this for well-understood, low-risk resources in non-production environments.

### Pause After Preview

When enabled (and Preview Before Apply is also enabled), the Infra Job pauses after the preview completes but before the apply or destroy runs. A team member must explicitly approve the previewed plan before the job continues.

This is the most common production governance pattern: let the preview run automatically so you can see the plan, then require a human to approve it before any changes are made to real infrastructure.

<!-- SCREENSHOT: Flow Control settings in Infra Job detail
  Page: /resource/infra-hub/infra-job/[infraJobId] (Config tab)
  Action: Show the Flow Control settings display with boolean indicators
  Focus: The flow control settings section with check/cross icons
  Alt: Infra Job configuration tab showing resolved flow control settings with enabled/disabled indicators for each control
-->

## Annotations a Stack's Plan Reads

Two annotations on a resource decide what an [Infra Stack](/docs/infrastructure/infra-stacks) may do when it lets go of the resource. A plan reads both before anything runs, and adding or removing either changes no deployment — the plan reads **Update Record**: only Planton's record changes.

| Annotation | Whenever the stack lets go of the resource | In a plan that removes it |
|------------|---------------------------------------------|---------------------------|
| `planton.dev/keep-on-provider: "true"` | Released, never destroyed: removed from the templates, undeployed or purged, it keeps running on its provider, and another stack can adopt it or a service can take it over. | **Keep** (named by its provider: Keep in AWS), the choice locked: "Marked planton.dev/keep-on-provider: its stack never destroys it". |
| `planton.dev/prevent-destroy: "true"` | Refused: an undeploy or purge is refused before anything starts, naming the resource. | **Blocked** from destroying — choose Keep to apply the plan, or remove the annotation to destroy it. |

Keep answers "this resource outlives its stack"; prevent-destroy answers "stop anyone from tearing this down by mistake". A removed resource without either is destroyed when its plan says Destroy and you agreed — a resource that holds data only when you named it.

## Policy Resolution Hierarchy

Flow Control policies are attached to a scope using a selector. The platform resolves which policy applies to a given Infra Job by checking four levels, from most specific to least:

1. **Infra Component** — A policy targeting a specific Infra Component by ID
2. **Environment** — A policy targeting all resources in a specific environment
3. **Organization** — A policy targeting all resources in the organization
4. **Platform** — A default policy for the entire platform

The first match wins. Policies are not merged across levels — if an environment-level policy exists, the organization and platform policies are ignored entirely for resources in that environment.

If no policy matches at any level, the platform default applies.

## Common Patterns

### Development: Fast Iteration

Optimize for speed. Let changes apply immediately without manual steps:

- Manual Approval Required: **off**
- Lifecycle Events Disabled: **off**
- Preview Before Apply: **on** (the plan still appears in logs)
- Pause After Preview: **off** (no manual approval step)

### Production: Full Control

Maximize safety. Require preview and manual approval before any change takes effect:

- Manual Approval Required: **off** (let the job start and preview automatically)
- Lifecycle Events Disabled: **off**
- Preview Before Apply: **on**
- Pause After Preview: **on** (require human approval of the plan)

### Shared Infrastructure: Explicit Only

For resources that multiple teams depend on, disable automatic triggers entirely:

- Manual Approval Required: **on**
- Lifecycle Events Disabled: **on** (no automatic Infra Jobs)
- Preview Before Apply: **on**
- Pause After Preview: **on**

## Using the Web Console

Flow Control settings are visible on the Infra Job detail page in the configuration tab. The display shows each of the four controls with an enabled/disabled indicator, reflecting the resolved policy for that specific job.

To create or update Flow Control policies, use the CLI (below). Policy management through the web console is planned.

## Using the CLI

Flow Control policies are managed through the standard resource commands:

```bash
# Create or update a Flow Control policy from a YAML manifest
planton apply -f flow-control-policy.yaml

# Get a Flow Control policy by ID
planton get flow-control-policy <policy-id>

# Delete a Flow Control policy
planton delete flow-control-policy <policy-id>
```

To check which Flow Control policy applies to a given resource type and environment, use the Infra Job preflight check:

```bash
planton infra job preflight-checks --catalog-kind <kind>
```

The preflight report includes the resolved policy and its settings.

## Related Documentation

- [Infra Jobs](/docs/infrastructure/infra-jobs) — The execution units that Flow Control governs
- [Infra Components](/docs/infrastructure/infra-components) — The resources whose deployments are controlled
- [Infra Pipelines](/docs/infrastructure/infra-pipelines) — Pipeline-level approval gates for multi-resource deployments
- [Infra Stacks](/docs/infrastructure/infra-stacks) — Plans, revisions, and what a stack does with what it no longer declares
