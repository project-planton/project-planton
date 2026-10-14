---
title: "Infra Stacks"
sidebar_title: "Infra Stacks"
description: "An Infra Chart declared into one environment: every change is a plan you read first, and every applied plan is a numbered revision."
icon: infrastructure
order: 40
tags:
  - Infrastructure
  - Infra Stacks
  - Deployment
---

# Infra Stacks

An Infra Chart defines what you *can* deploy — a reusable template with parameterized placeholders. But a template alone does not answer the questions that matter for a specific deployment: which region, which domain name, how many availability zones, which environment. An Infra Stack captures those answers. It is a template declared into one environment with the values set there, applied as real infrastructure -- and every change to it is a plan you read before anything runs.

The relationship mirrors how Helm works: an Infra Chart is to an Infra Stack as a Helm chart is to a Helm release. The chart is reusable; the stack is the deployed instance with specific values. This separation enables teams to capture infrastructure patterns once and instantiate them many times — across different environments, regions, or customers — each instance tracked as a distinct Infra Stack.

## Why Infra Stacks Exist

Before Infra Stacks, the platform had Infra Charts for reusability but no persistent record of each deployment instance. Consider an organization that deploys five ECS environments from the same chart — development, staging, production-us, production-eu, and production-asia. Each uses different parameter values. Without a stack record, fundamental questions have no answers:

- Which parameters deployed production-us last month?
- Did staging update to the latest chart version?
- Who changed the load balancer domain in production-eu?

Infra Stacks answer them with a persistent record of every template instantiation. Each stack holds its declaration — the chart's templates, the chart's `values.yaml` (every value declared, with its default) and the values set for its environment — with its render, its dependency graph, and a numbered **revision** for every plan applied to it: who applied it, why, and how its run ended. Runs are transient; the stack and its revisions are the source of truth for what was declared and when.

## One Chart, One Environment

You select an Infra Chart from the catalog, pick the environment it deploys into, set the values that environment needs, and review the plan of what installing creates.

The stack carries its own copy of the chart's templates and `values.yaml`. Chart updates never change an existing stack silently — each stack is self-contained, and it can be checked out to disk and edited with no chart in reach. Only the values someone set are stored as the environment's values; every other value takes the default from the stack's `values.yaml`, so a default the stack's `values.yaml` changes reaches every environment that never chose otherwise — and the plan shows it as a value change first.

The web console provides a creation flow: browse the chart catalog, pick the environment, fill in the values form (generated from the chart's declarations, each default shown as a placeholder), then **Review Plan**.

<!-- SCREENSHOT: Infra Stack creation from chart
  Page: /[org]/infra-stacks/create/[infraChartId]
  Action: Show the creation wizard with chart parameters form
  Focus: The parameter form and environment selector
  Alt: Infra Stack creation wizard showing chart parameters and environment selection for an AWS ECS environment chart
-->

## Every Change Is a Plan

Every change to an Infra Stack — an edit in the console, `planton infra stack apply`, `planton chart install`, `planton apply -f`, an agent's tool — is first a **plan**: what applying the declaration would do to each resource, before anything runs.

| Row | What applying does |
|-----|--------------------|
| **Blocked** | Nothing, until it is resolved; the row says why and the way out. |
| **Destroy** | The resource is no longer in the templates: it is destroyed in the cloud once everything the plan applies has applied. |
| **Replace** | The resource stays in the templates, but its provider can't change it in place: it is destroyed and created again, and what it holds does not carry over. A request's preview finds it, and an approver names each one that holds data. |
| **Update** | Something that decides its deployment moved: its settings, a secret's version, a value it reads, its module. |
| **Update Record** | Only Planton's record changes; nothing in the cloud (for example, an annotation added). |
| **Create** | The resource is new. |
| **Adopt** | The resource already exists and no stack manages it; the stack takes it over only when you choose Adopt, and the apply may change it to match the template. |
| **Keep** (named by its provider: Keep in AWS) | The resource is no longer in the templates and is kept: Planton stops managing it, and it keeps running where it is, and billing there. |
| **May Change** | It reads a value known only after an earlier resource applies. |
| **Always Runs** | Its job runs on every apply: whether it changes is known only when it runs. |
| **Unchanged** | Nothing it depends on moved since its last successful deploy. |
| **Unmanaged** | Planton records it and never applies it. |

Applying a plan runs exactly that plan, and only while nothing it was made from has moved: the apply plans again and is refused when the fresh plan differs ("This plan is out of date…") — someone applied in the meantime, or a resource changed underneath. Plan again and review what changed.

An apply records a numbered **revision** named by its message ("r13 · Raise replicas") and starts one [Infra Pipeline](/docs/infrastructure/infra-pipelines) run: it applies what the stack declares, each resource after everything it consumes, up to four at once; then — only when all of it applied — it removes what the plan removed, dependents first: destroyed, or let go of and kept on its provider. A failed apply removes nothing. A plan that changes only the declaration (a comment in a template, a value set to its default) records its revision and starts no run; a plan that changes nothing applies nothing.

### Choosing What Happens to Each Resource

- **Destroy or Keep.** A removed resource is destroyed unless you keep it. A resource that holds data a destroy cannot give back is destroyed only when you name it: type its name in the console, `--destroy Kind/name` on the CLI. A removed resource whose destroy is blocked — `planton.dev/prevent-destroy`, its own deletion protection, or something that still reads it — blocks only its destroy: choose Keep and the plan applies.
- **Adopt.** A stack never takes over a resource it found; you choose Adopt for it, or remove it from the templates.
- **Re-apply.** Run an unchanged resource's job anyway, or every resource's (`--reapply Kind/name`, `--reapply-all`).

`planton apply -f` applies a plan only when nothing in it is destroyed, adopted or blocked; otherwise it refuses and names the `planton infra stack apply` command to review the plan with. It never files a request: in a protected environment it refuses and names that command.

### In a Protected Environment, a Change Is a Request

In a protected environment applying a plan files a **request** instead of writing anything. The plan says **This Plan Needs an Approver**, and the request carries a message its approvers read first:

```bash
planton infra stack apply app-infra --env prod -m "Raise replicas"
```

The request's run previews what the change reaches — each preview shows its engine's plan and changes nothing — and an approver reads that exact plan, what is destroyed and replaced first, in the console, on the phone, or in a terminal:

```bash
planton infra stack review app-infra --env prod
planton infra stack approve app-infra --env prod --replace AwsRedisElasticache/orders-cache
planton infra stack reject app-infra --env prod --reason "Not during the freeze"
planton infra stack withdraw app-infra --env prod   # the requester takes it back
```

- **Nothing applies until an approval**, and the stack's declaration does not move: a rejected request leaves only its revision in history.
- **The approver approves exactly the plan they read.** Each resource plans again and applies only when its plan matches its preview; one that moved since — a setting edited in the cloud, a secret's new version — stops before changing anything, shows its new plan and waits for a person. A resource that could only be previewed once an earlier change applied (May Change) applies under the approval unless its plan destroys or replaces something.
- **What holds data is named.** Approving a plan that destroys or replaces a resource holding data asks the approver to type each name (`--destroy` / `--replace Kind/name` in CI, with the `--review-digest` of the review read).
- **A second pair of eyes.** The person who asked cannot approve their own change while another approver exists. An undecided request expires after 7 days.
- **A terminal never waits on a person.** The apply follows while the request previews, then ends with where to review it and the commands that decide it; `-o json` prints `{run_id, review_url, state, expires_at, revision}`.

A request that is rejected, withdrawn, expires, cannot be previewed or goes out of date applies nothing. Plan it again as it was, with the choices its author made: `planton infra stack plan app-infra --env prod --revision 14`. Undeploy, purge and rerun into a protected environment are requests too.

### Revisions

Each applied plan is a revision: its number, message, author, the declaration it applied, the plan's rows as reviewed and chosen, and how its run ended — **Applying**, **Applied**, **Applied With Removal Failures** (what the stack declares applied; a removal failed), **Failed** or **Cancelled**. A request reads **Awaiting Approval** until someone decides it, and closes **Rejected**, **Withdrawn**, **Expired**, **Preview Failed** or **Out of Date** when it applies nothing. The stack's **live revision** is the newest that applied; an undeploy clears it. `planton infra stack history <name> --env <env>` lists them, newest first, with who requested and who approved each.

A run's Rerun (`planton infra pipeline rerun <pipeline-id>` in a terminal) repeats its revision — the same reviewed plan — so what applied stays Unchanged and what failed runs again. It asks first when the rerun still destroys something. A rerun is refused once a newer revision exists, or after an undeploy: plan again instead.

### One Run at a Time

A stack has one run at a time. While a run is queued, running or waiting for approval, a plan answers with that run instead of rows, and an apply, undeploy, purge or rerun is refused, naming the run — or the request waiting for approval, with the commands to review or withdraw it. Cancel a run with `planton infra pipeline cancel <pipeline-id>`, then try again. A run whose workflow is gone is ended by the platform within about ten minutes, so a stack is never left busy.

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

### Plan and Apply

Plan a stack from a chart, a checked-out folder, or the declaration it already holds:

```bash
planton infra stack plan my-stack ./my-chart --env <env> -f values-prod.yaml
planton infra stack apply my-stack ./my-chart --env <env> -f values-prod.yaml -m "why"
```

A values file lists only the values it sets; every other value keeps the chart's default (see [how values resolve](/docs/infrastructure/infra-charts)). With no chart, the plan is of the stack's own declaration — what moved underneath it — and `--set name=value` alone changes one value. `planton chart install my-stack ./my-chart --env <env>` takes the same path and creates the stack when the name is new; `--dry-run` stops at the plan.

At a terminal, the apply shows the plan and asks once, and asks for the name of anything that holds data before destroying it. Without a terminal (CI), name each destroy with `--destroy Kind/name`, or pass `--plan-digest` from a plan a person reviewed — it consents to that plan's ordinary destroys; a resource that holds data is still named. `--yes` answers the question but consents to no destroy.

`planton infra stack plan` exits `0` when planned, `1` when the declaration does not build, a declared resource is blocked or a run is in flight, and `2` when the plan could not be made; `--detailed-exitcode` adds `4` for a plan that would change something. `-o markdown` prints the plan for a pull request.

In the console, **Review Plan** opens the same plan — from the stack page, the Configuration tab, the create form, or Studio — and **Apply** runs it.

### Undeploy

Destroy everything the stack owns without deleting the stack record. The stack's declaration and revisions are kept, and applying its plan again brings it back — useful for temporarily tearing down infrastructure to save costs:

```bash
planton infra stack undeploy <stack-name-or-id> --env <env>
```

An undeploy covers every Infra Component the stack created, including components its templates no longer declare. It reads each resource from its stored record, never from the templates, and destroys dependents first: a workload goes before the cluster it runs on. Each record is settled the moment its own destroy ends:

- A resource the stack holds as **Unmanaged**, or one marked `planton.dev/keep-on-provider: "true"`, is released from the stack and stays on its provider. An undeploy never destroys it.
- A resource whose record is already gone reads **Already Gone**.
- A resource another stack or service now owns is left in place for its owner.

An undeploy is refused before anything starts when:

- any resource the stack owns carries `planton.dev/prevent-destroy: "true"` (see [Protecting a Resource from Teardown](#protecting-a-resource-from-teardown));
- another stack still reads one of its resources — through a reference to it, or with workloads placed in its namespace. The refusal names the reading stack;
- the stack owns no Infra Components: it "owns no infra components, so there is nothing to undeploy".

In a protected environment an undeploy is a request an approver decides, as an apply is: nothing is destroyed until then. The CLI asks before undeploying or purging; `--yes` answers ahead. Applying the stack again later plans each resource anew: any resource the undeploy released plans Adopt.

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

To lift the protection, remove the annotation from the stack's templates and apply the stack (the plan reads Update Record: nothing in the cloud changes). A plan that removes a protected resource reads it Blocked from destroying — choose Keep to apply around it. For a resource the stack no longer declares, remove the annotation from the Infra Component itself. Deleting a record alone is not a destroy, so the annotation does not stand in its way.

## Keeping a Resource in the Cloud

Annotate a resource with `planton.dev/keep-on-provider: "true"` and its stack never destroys it: whenever the stack lets go of it — removed from the templates, undeployed, purged — it is released, still running, and the plan reads Keep with the choice locked. A released resource can be adopted by another stack or taken over by a service. This is how a resource moves from an Infra Stack to a Service: remove it from the templates, keep it running on its provider, and the service's next deploy takes it over.

A kept resource keeps what it depends on. While a deployment stays, the namespace it lives in — and anything else it references — is never destroyed: the namespace's plan row says what depends on it, with **Keep This Too** in the console and the exact `--keep` flags on the CLI, and an undeploy is refused before anything is torn down, naming each kept resource and what it depends on.

## What a Run Applies

A stack's run applies only the resources whose inputs moved since their last successful deploy — the same comparison its plan shows. Before a resource runs, Planton compares what it would deploy — its settings, the values it reads from other resources, its secrets' current versions and its variables, the account it deploys into, the module version, and whether anything it sits on was rebuilt — with what its last successful deploy applied. An exact match runs no job: the resource reads **Unchanged since** the job that last applied it, and hands that job's outputs to the resources that depend on it. Change one value and that resource runs, along with any dependent whose inputs it changes. When Planton cannot know exactly what a resource would deploy — a module that follows a branch, a secret whose version cannot be read, a connection that signs in through the runner's own machine — the resource Always Runs.

A change made outside Planton, such as an edit in the cloud console, moves nothing Planton deploys, so a plan reads that resource Unchanged. Refresh the Infra Component first: a refresh that finds the change clears what the resource recorded as applied, and the next plan reads it Update.

More outcomes appear on a run's resources:

- **Unmanaged** — a resource whose lifecycle events are disabled (`planton.dev/infra-job.disable-on-lifecycle-events: "true"`, see [Flow Control](/docs/infrastructure/flow-control)). The stack records it and never applies or destroys it; removing it from the templates plans Keep, and an undeploy releases it.
- **Kept** — a removed resource the stack let go of; no job ran, and it keeps running.
- **Owned elsewhere** — a resource another live Infra Stack or Service owns. A teardown leaves it in place for its owner, and a plan that would change it reads Blocked, naming the owner to change it through.

## Using the CLI

```bash
# Plan: what applying would do, resource by resource (writes nothing)
planton infra stack plan my-stack ./chart-dir --env <env> -f values.yaml

# Apply exactly the plan shown, as a numbered revision, and follow its run
planton infra stack apply my-stack ./chart-dir --env <env> -f values.yaml -m "why"

# Install a chart as a stack: the same plan, review and apply
planton chart install my-stack ./chart-dir --env <env> -f values.yaml -m "why"

# A stack's revisions, newest first
planton infra stack history <stack-name> --env <env>

# List runs for a stack, and its last run
planton infra stack infra-pipelines <stack-id>
planton infra stack last-pipeline <stack-id>

# Undeploy (destroy resources, keep the stack and its revisions)
planton infra stack undeploy <stack-name-or-id> --env <env>

# Purge (destroy resources and delete the stack)
planton infra stack purge <stack-name-or-id> --env <env>

# Get stack details: its declaration and live revision
planton infra stack get <stack-name-or-id> --env <env>

# Check a stack out as a working copy: templates, values.yaml, values.<env>.yaml
planton infra stack checkout <stack-name-or-id> --env <env>
```

## When to Use Infra Stacks vs. Direct Infra Components

| Scenario | Recommended Approach |
|----------|---------------------|
| Complete environment (dev, staging, production) | Infra Stack — coordinated deployment, dependency handling, one recorded configuration |
| Single resource needed quickly | Direct Infra Component — no orchestration overhead |
| Reproducible deployments | Infra Stack — configuration captured as a persistent record |
| Quick experiments | Direct Infra Component — fast iteration, easy to delete |
| Production infrastructure | Infra Stack — reviewed plans, revisions, requests an approver decides on the exact plan, protected teardown |
| Adding to existing infrastructure | Direct Infra Component — targeted change without touching other resources |

## Related Documentation

- [Infra Charts](/docs/infrastructure/infra-charts) — The templates that Infra Stacks instantiate
- [Infra Pipelines](/docs/infrastructure/infra-pipelines) — How Infra Stack deployments are orchestrated
- [Infra Components](/docs/infrastructure/infra-components) — The deployed instances that Infra Stacks own and manage
- [Infra Jobs](/docs/infrastructure/infra-jobs) — The atomic execution units within pipelines
- [Flow Control](/docs/infrastructure/flow-control) — Governance policies that affect pipeline execution
