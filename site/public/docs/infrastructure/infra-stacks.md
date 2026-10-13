---
title: "Infra Stacks"
sidebar_title: "Infra Stacks"
description: "An Infra Chart rendered into one environment: the configured instance that deploys and manages a collection of Infra Components through deploy and undeploy pipelines."
icon: infrastructure
order: 40
tags:
  - Infrastructure
  - Infra Stacks
  - Deployment
---

# Infra Stacks

An Infra Chart defines what you *can* deploy — a reusable template with parameterized placeholders. But a template alone does not answer the questions that matter for a specific deployment: which region, which domain name, how many availability zones, which environment. An Infra Stack captures those answers. It is a configured instance of a template with your specific parameter values filled in, ready to deploy as real infrastructure.

The relationship mirrors how Helm works: an Infra Chart is to an Infra Stack as a Helm chart is to a Helm release. The chart is reusable; the stack is the deployed instance with specific values. This separation enables teams to capture infrastructure patterns once and instantiate them many times — across different environments, regions, or customers — each instance tracked as a distinct Infra Stack.

## Why Infra Stacks Exist

Before Infra Stacks, the platform had Infra Charts for reusability but no persistent record of each deployment instance. Consider an organization that deploys five ECS environments from the same chart — development, staging, production-us, production-eu, and production-asia. Each uses different parameter values. Without a stack record, fundamental questions have no answers:

- Which parameters deployed production-us last month?
- Did staging update to the latest chart version?
- Who changed the load balancer domain in production-eu?

Infra Stacks solve this by introducing a persistent record of every template instantiation. Each stack captures the chart's templates, the exact parameter values used, the environment, the rendered catalog object manifests, the dependency graph, and a link to every pipeline that deployed it. Pipelines are transient — they run and complete. Infra Stacks persist as the source of truth for what was configured and when.

## One Chart, One Environment

You select an Infra Chart from the catalog, pick the environment it deploys into, provide parameter values, and the stack renders the chart's templates into concrete catalog object manifests.

When the stack is created, the complete template is copied from the chart into the stack. This ensures that chart updates do not break or silently change existing stacks — each stack is a self-contained snapshot of the template it was created from, and it can be checked out to disk and edited with no chart in reach.

The web console provides a creation wizard: browse the chart catalog, pick the environment, fill in the parameter form (auto-generated from the chart's parameter definitions), preview the rendered output, and deploy.

<!-- SCREENSHOT: Infra Stack creation from chart
  Page: /[org]/infra-stacks/create/[infraChartId]
  Action: Show the creation wizard with chart parameters form
  Focus: The parameter form and environment selector
  Alt: Infra Stack creation wizard showing chart parameters and environment selection for an AWS ECS environment chart
-->

## Automatic Pipeline Triggering

Every save of an Infra Stack — create, update, apply, or `planton chart install` — starts a deploy run, an [Infra Pipeline](/docs/infrastructure/infra-pipelines). The pipeline takes the stack's Infra Component dependency graph and deploys each resource after everything it consumes, up to four at once. There is no separate "apply" step — an Infra Stack is the declaration of an environment's infrastructure, and the platform acts on it. A save never destroys anything: a resource you remove from the templates stays deployed, and stays the stack's, until the stack is undeployed.

Deploying again without changing anything, and tearing the environment down, are explicit operations on the stack: `deploy` and `undeploy` are a pair, and both hand back the stack with the new run's id.

### One Run at a Time

A stack has one run at a time. A save, deploy, undeploy, purge or rerun is refused while a run of the stack is queued, running or waiting for approval, and the refusal names that run. Approve or reject a run that waits for approval, or cancel it (`planton infra pipeline cancel <pipeline-id>`), then try again. A run whose workflow is gone is ended by the platform within about ten minutes, so a stack is never left busy.

## Dependency Graph Visualization

Each Infra Stack maintains a dependency graph that visualizes the relationships between its Infra Components. The web console displays this as an interactive DAG where you can see:

- Resource nodes with their current deployment status (queued, running, succeeded, failed)
- Dependency edges showing which resources must be deployed before others
- Real-time updates as a pipeline progresses through the graph

This visualization makes it easy to understand the deployment topology and diagnose failures — if a resource fails, you can immediately see which downstream resources are blocked.

<!-- SCREENSHOT: Infra Stack DAG visualization
  Page: /[org]/infra-stack/[infraStackSlug]
  Action: Show the Infra Stack detail page with the DAG tab active
  Focus: The dependency graph visualization showing resources and their connections
  Alt: Infra Stack DAG visualization showing an AWS VPC connected to subnets, security groups, and an RDS instance
-->

## One Stack per Environment

An Infra Stack lives in exactly one environment, and its name is unique within that environment. Installing the same chart into staging and production gives you two stacks that share a name: each keeps its own values, its own history and its own pipelines, and changing one never touches the other. Every command that names a stack by name also names its environment with `--env`; a stack's id identifies it on its own. A stack's environment is fixed when it is created.

The environment also decides who may change the stack. Stacks follow the same access rules as the Infra Components they deploy: a permission granted on an environment covers that environment's stacks, and a protected environment stops organization-wide deploy and destroy permissions at its boundary.

## Infra Stack Lifecycle

### Create

Create an Infra Stack from a chart with the CLI:

```bash
planton chart install my-stack ./my-chart -f values.yaml --org <org> --env <env> -m "why"
```

A values file lists only the parameters it changes; every other parameter keeps the chart's value (see [how parameter values resolve](/docs/infrastructure/infra-charts)). Add `--dry-run` to see every parameter and the rendered documents first, without creating anything.

Or from the web console's chart catalog by selecting a chart, picking the environment, filling in parameters, and deploying. The message becomes the run's own name, so a person reading the run later knows why it happened.

### Redeploy

Trigger a new pipeline without changing the stack's configuration. Useful for retrying after a failed deployment, or re-running after fixing external issues like quota limits or permissions:

```bash
planton infra stack deploy <stack-name-or-id> --env <env>
```

A redeploy runs only what needs to run (see [What a Run Applies](#what-a-run-applies)). To put back a resource someone changed outside Planton, refresh that Infra Component first, then redeploy the stack.

### Undeploy

Destroy everything the stack owns without deleting the stack record. The stack configuration is preserved and can be redeployed later — useful for temporarily tearing down infrastructure to save costs:

```bash
planton infra stack undeploy <stack-name-or-id> --env <env>
```

An undeploy covers every Infra Component the stack created, including components its templates no longer declare. It reads each resource from its stored record, never from the templates, and destroys dependents first: a workload goes before the cluster it runs on. Each record is settled the moment its own destroy ends:

- A resource the stack holds as **Unmanaged** is released from the stack and stays in the cloud. An undeploy never destroys it.
- A resource whose record is already gone reads **Already Gone**.
- A resource another stack or service now owns is left in place for its owner.

An undeploy is refused before anything starts when:

- any resource the stack owns carries `planton.dev/prevent-destroy: "true"` (see [Protecting a Resource from Teardown](#protecting-a-resource-from-teardown));
- another stack still reads one of its resources — through a reference to it, or with workloads placed in its namespace. The refusal names the reading stack;
- the stack owns no Infra Components: it "owns no infra components, so there is nothing to undeploy".

A protected environment asks for approval before an undeploy starts, as it does before a deploy.

### Purge

Undeploy the stack, then delete its record — the record is deleted only if the undeploy succeeded:

```bash
planton infra stack purge <stack-name-or-id> --env <env>
```

A purge is refused whenever an undeploy would be, with one difference: a purge of a stack that owns nothing simply deletes it.

### Delete

Deleting an Infra Stack never destroys infrastructure, so it is refused while the stack still owns Infra Components — undeploy it first, or purge it to do both. A delete is also refused while a run of the stack is in flight.

## Protecting a Resource from Teardown

Annotate a resource with `planton.dev/prevent-destroy: "true"` and it is never torn down through the platform:

```yaml
metadata:
  name: prod-db
  annotations:
    planton.dev/prevent-destroy: "true"
```

While any resource the stack owns carries the annotation, the stack's undeploy and purge are refused before anything starts, naming the resource. A person's destroy or purge of that Infra Component is refused too, as is creating, rerunning or resuming a destroy job for it. Adding or removing the annotation changes no deployment: no job runs for that change alone.

To lift the protection, remove the annotation from the stack's templates and apply the stack. For a resource the stack no longer declares, remove it from the Infra Component itself. Deleting a record alone is not a destroy, so the annotation does not stand in its way.

## What a Run Applies

A stack's deploy applies only the resources whose inputs moved since their last successful deploy. Before a resource runs, Planton compares what it would deploy — its settings, the values it reads from other resources, its secrets' current versions and its variables, the account it deploys into, the module version, and whether anything it sits on was rebuilt — with what its last successful deploy applied. An exact match runs no job: the resource reads **Unchanged since** the job that last applied it, and hands that job's outputs to the resources that depend on it. Change one value and that resource runs, along with any dependent whose inputs it changes; redeploy with nothing changed and nothing runs. When Planton cannot know exactly what a resource would deploy — a module that follows a branch, a secret whose version cannot be read, a connection that signs in through the runner's own machine — the resource always runs.

A change made outside Planton, such as an edit in the cloud console, moves nothing Planton deploys, so a redeploy reads that resource as unchanged. Refresh the Infra Component first: a refresh that finds the change clears what the resource recorded as applied, and the next redeploy runs its job.

Two more outcomes appear on a run's resources:

- **Unmanaged** — a resource whose lifecycle events are disabled (`planton.dev/infra-job.disable-on-lifecycle-events: "true"`, see [Flow Control](/docs/infrastructure/flow-control)). The stack records it and never applies or destroys it: a deploy leaves it alone, and an undeploy or purge releases it from the stack and leaves it in the cloud.
- **Owned elsewhere** — a resource another live Infra Stack or Service owns. An undeploy or purge leaves it in place for its owner, and a deploy that would change it is refused, naming the owner to change it through.

## Using the CLI

```bash
# Create an Infra Stack from a chart (triggers deployment automatically)
planton chart install my-stack ./chart-dir -f values.yaml

# Deploy an existing stack (starts a deploy run and follows it)
planton infra stack deploy <stack-name-or-id> --env <env>

# List pipelines for a stack
planton infra stack infra-pipelines <stack-id>

# Get the last pipeline for a stack
planton infra stack last-pipeline <stack-id>

# Undeploy (destroy resources, keep stack)
planton infra stack undeploy <stack-name-or-id> --env <env>

# Purge (destroy resources and delete stack)
planton infra stack purge <stack-name-or-id> --env <env>

# Get stack details
planton infra stack get <stack-name-or-id> --env <env>

# Check a stack out as a chart-shaped folder you can edit and install again
planton infra stack checkout <stack-name-or-id>
```

## When to Use Infra Stacks vs. Direct Infra Components

| Scenario | Recommended Approach |
|----------|---------------------|
| Complete environment (dev, staging, production) | Infra Stack — coordinated deployment, dependency handling, one recorded configuration |
| Single resource needed quickly | Direct Infra Component — no orchestration overhead |
| Reproducible deployments | Infra Stack — configuration captured as a persistent record |
| Quick experiments | Direct Infra Component — fast iteration, easy to delete |
| Production infrastructure | Infra Stack — coordinated changes, approval gates, protected teardown |
| Adding to existing infrastructure | Direct Infra Component — targeted change without touching other resources |

## Related Documentation

- [Infra Charts](/docs/infrastructure/infra-charts) — The templates that Infra Stacks instantiate
- [Infra Pipelines](/docs/infrastructure/infra-pipelines) — How Infra Stack deployments are orchestrated
- [Infra Components](/docs/infrastructure/infra-components) — The deployed instances that Infra Stacks own and manage
- [Infra Jobs](/docs/infrastructure/infra-jobs) — The atomic execution units within pipelines
- [Flow Control](/docs/infrastructure/flow-control) — Governance policies that affect pipeline execution
