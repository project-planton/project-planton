---
title: "Infra Jobs"
description: "The atomic execution unit that provisions, updates, and destroys cloud infrastructure — running IaC operations with full visibility and control."
icon: pipeline
order: 50
tags:
  - Infrastructure
  - Infra Jobs
  - Infrastructure
---

# Infra Jobs

An Infra Job is the execution unit behind every infrastructure change in Planton. When you create an Infra Component, update its configuration, or destroy it, an Infra Job runs the Infrastructure-as-Code operations that make the change real — initializing the module, refreshing state, previewing changes, and applying them. Every step is tracked, every operation is logged, and the entire execution is preserved for audit.

## Why Infra Jobs Exist

Running infrastructure automation manually is fragile. You start with `terraform apply` (or `tofu apply`), then realize you should preview first, then add a drift check to catch changes made in the console, then add approval gates for production — and before long you are maintaining custom CI pipelines just to deploy a database safely.

Infra Jobs codify this workflow into the platform. The right operations run in the right order, every time. Credentials are resolved automatically. A resource's flow control and its environment's protection decide where a job waits for a person. And every execution — successful or failed — is recorded with full logs and resource diffs for troubleshooting and compliance.

## How an Infra Job Executes

Every Infra Job follows a defined sequence of operations. Which operations run depends on what the job is doing.

### For Creating or Updating Infrastructure

1. **Initialize** — Set up the IaC module, download providers, configure the state backend.
2. **Preview** — Generate a plan against the live cloud showing what will change. Additions, modifications, and deletions are displayed before anything happens, and the plan writes nothing.
3. **Apply** — Execute the changes. Create new resources, modify existing ones, update configurations.

### For Destroying Infrastructure

1. **Initialize** — Same module and state backend setup.
2. **Preview** — Show what will be destroyed.
3. **Destroy** — Remove the resources from the cloud provider.

### For Checking for Drift

1. **Initialize** — Same module and state backend setup.
2. **Drift Check** — Compare what the state recorded when Planton last applied the resource with what is really in the cloud, setting by setting. This catches drift — changes made outside of Planton, such as manual edits in the cloud console. It writes nothing: not the state file, not the resource.

The resource then reads **In Sync**, **Drifted**, or **Couldn't Check**, and a drifted resource is put back by its next plan or deploy. See [Infra Stacks](/docs/infrastructure/infra-stacks#changed-outside-planton-drift) for checking a whole stack and reconciling it.

### For Importing Existing Resources

1. **Initialize** — Same module and state backend setup.
2. **Import** — Read the resource from the cloud provider and write it to the IaC state file. The actual infrastructure is not modified.
3. **Refresh** — Confirm the imported resource's state matches what exists on the provider.

Import Infra Jobs do not include preview or apply steps. The import operation writes to state only. See [Importing Resources](/docs/infrastructure/importing-resources) for the full guide, including provisioner-specific identifiers and CLI commands.

A resource's [flow control](/docs/infrastructure/flow-control) can make a job wait for approval before it begins, or pause between preview and apply until someone approves the plan; a protected environment adds that pause to every direct change made there.

<!-- SCREENSHOT: Infra Job detail page
  Page: /resource/infra-hub/infra-job/[infraJobId]
  Action: Show an Infra Job in progress with operation steps visible
  Focus: The operation steps with status indicators and expandable logs
  Alt: Infra Job detail page showing initialize completed, refresh in progress, preview and apply pending
-->

## What Gets Resolved Before Execution

Before an Infra Job starts running, the platform resolves four things automatically.

### IaC Module

The Pulumi, Terraform, or OpenTofu module that contains the provisioning logic for the catalog kind. If the Infra Component specifies a custom module, that module is used. Otherwise, the platform looks up the registered module for the catalog kind — first in the organization's module registry, then in the platform-level registry.

### Provider Credentials

The cloud provider credentials needed to provision the infrastructure — an AWS connection, a GCP connection, or whatever the resource type requires. These are resolved from the Infra Component's configuration, which specifies the connection to use. See [Connections](/docs/connections) for how connections are managed.

### State Backend

Where the IaC state file is stored. Each Infra Component gets a dedicated state file, stored in a backend (such as an S3 bucket or Pulumi Cloud) configured at the organization level. This ensures that one resource's state operations never interfere with another's.

### Flow Control

Whether the job waits for approval before it begins, and whether it pauses between preview and apply. It comes from the Infra Component itself — its annotations or its typed `flow_control` — and is off by default; a protected environment adds the pause to a direct change. It always resolves, so it never stops a job from starting. See [Flow Control](/docs/infrastructure/flow-control).

If the module, the credentials or the state backend cannot be resolved — missing credentials, no registered module, no state backend configured — the Infra Job fails its preflight check and does not execute. The preflight report identifies exactly what is missing.

A connection that exists but isn't [authorized](/docs/connections/environment-mappings) for the job's environment is refused the same way, before any resource runs, and the reason names the fix: *"Nothing ran: environment staging may not use the aws connection aws-prod. Authorize it (planton connection auth create --provider aws --connection aws-prod --scope environment --environments staging), or make it the organization's (--scope organization), then run the job again."*

## Two Deployment Paths

### Direct

When you create, update, or destroy a single Infra Component, one Infra Job runs for that resource. This is the simplest path — one resource, one job, one execution.

### Orchestrated

When an [Infra Pipeline](/docs/infrastructure/infra-pipelines) deploys multiple Infra Components, each resource gets its own Infra Job. The pipeline coordinates execution order based on the dependency graph — ensuring that a VPC's Infra Job completes before the database that depends on it starts.

## Monitoring Progress

### Web Console

The Infra Job detail page shows real-time progress as the job executes. Each operation (initialize, refresh, preview, apply/destroy) has its own status indicator — queued, running, succeeded, or failed. Expanding an operation reveals the full logs, including resource-level details showing what is being created, modified, or deleted.

The Infra Component detail page includes an Infra Jobs tab listing all jobs that have run for that resource, with timestamps, operation types, and results.

### CLI

Stream logs from a running Infra Job:

```bash
planton infra job stream-progress-events <infra-job-id>
```

This command (aliased as `logs`) streams progress events in real time, including resource-level changes and operation transitions.

## Controlling Execution

### Pausing and Resuming

When the resource's flow control or a protected environment requires approval — either before execution starts or between preview and apply — the Infra Job pauses and waits. Resume it from the web console or CLI:

```bash
planton infra job resume <infra-job-id>
```

### Cancelling

Cancel a running Infra Job to stop execution. The currently in-flight IaC operation completes to avoid leaving resources in an inconsistent state, then remaining operations are skipped:

```bash
planton infra job cancel <infra-job-id>
```

### Re-running

Re-run a completed or failed Infra Job to repeat the same operation. Useful after fixing external issues like quota limits or permission errors. A resource an Infra Stack owns is rerun through its stack's run instead (`planton infra pipeline rerun <pipeline-id>`), which repeats the reviewed revision:

```bash
planton infra job rerun <infra-job-id>
```

`planton rerun <run-id>` reruns any run from its id alone — an Infra Job (`ij_`), a stack run (`infpipe_`), a service run (`svcpipe_`) or an external CI run (`extrun_`) — and follows it to the end; `--no-follow` returns once it starts.

## Preflight Checks

Before committing to execution, verify that all four essentials are in place for a given catalog kind and environment:

```bash
planton infra job preflight-checks --catalog-kind <kind>
```

The report shows whether the IaC module, provider credentials and state backend can all be resolved. This is useful when setting up a new environment or debugging why an Infra Job failed to start.

## Using the CLI

```bash
# Create an Infra Job for an Infra Component
planton infra job create <infra-component-id> --operation update --tail

# List Infra Jobs for a resource
planton infra job list <infra-component-id>

# Stream logs from a running job (alias: logs)
planton infra job stream-progress-events <infra-job-id>

# Cancel a running job
planton infra job cancel <infra-job-id>

# Resume a paused job
planton infra job resume <infra-job-id>

# Re-run a completed or failed job
planton infra job rerun <infra-job-id>

# Run preflight checks
planton infra job preflight-checks --catalog-kind <kind>

# View the IaC input for a completed job
planton infra job iac-input <infra-job-id>
```

The `--operation` flag on `infra job create` accepts: `refresh`, `preview`, `update`, `destroy`, `destroy_preview`. The default is `preview`. Any other value is refused as **Unknown Operation**, naming the valid ones, and no job is created.

Every `infra-job` subcommand names its argument on its usage line (`<infra-component-id>` or `<infra-job-id>`), and `--help` shows an example.

The summary printed when a job finishes names each operation by the engine that ran it: `tofu apply` and `tofu destroy` for OpenTofu (`terraform apply` for a resource set to Terraform), `pulumi up` and `pulumi destroy` for Pulumi.

Additional flags for `create`, `resume`, and `rerun`:

- `--tail` / `-t` — Follow the job's progress events after creation
- `--show-stack-summary` — Display the resource summary at completion (default: on)
- `--show-stack-diff` — Display detailed diffs for changed resources
- `--show-outputs` — Display the outputs at completion (default: on). A secret the resource generates shows as its `$secret/` reference; the value itself is kept in your secret store (see [Where Secrets Live](/docs/secrets/where-secrets-live#secrets-a-resource-generates))
- `--version-message` / `-m` — A description for the job, similar to a commit message

`planton infra ij` is the short form of `planton infra job`.

## Related Documentation

- [Infra Components](/docs/infrastructure/infra-components) — The infrastructure that Infra Jobs provision
- [Importing Resources](/docs/infrastructure/importing-resources) — Importing existing infrastructure into IaC state
- [Infra Pipelines](/docs/infrastructure/infra-pipelines) — Orchestrating multiple Infra Jobs across resources
- [Flow Control](/docs/infrastructure/flow-control) — Governance policies that control Infra Job execution
- [Connections](/docs/connections) — How provider credentials are managed and resolved
