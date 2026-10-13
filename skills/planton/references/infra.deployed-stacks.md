# Deployed Infra Stacks — Working Copies, Fixing Failures, Saving Changes

A chart is a blueprint; an **Infra Stack** is one chart installed into one
environment. A stack is named within its environment: the same chart
installed into qa and prod is two stacks that share a name, so every
command that names a stack also names its environment (`--env`). The moment
a stack is created it carries its OWN copy of the
templates and values — detached from the chart it came from. Editing the
chart never changes a deployed stack, and editing a stack never changes the
chart. This file governs the second mode of your work: modifying something
that is already running.

## Am I working on a chart or on a deployed stack?

Check for the binding file before assuming:

```
cat .planton/stack.yaml
```

The workspace file tree does not show hidden paths, so LOOK for it with the
shell — never conclude from the tree alone. A folder carrying this file is a
**working copy of a deployed Infra Stack**; the file names the stack
(`stackId`, `stackName`), its organization, and its environment:

```yaml
stackId: infstk_01jb7k4m2x9q3r5t6v8w0yzabc
stackName: eks-mumbai
org: acme
env: prod
```

No binding file → a plain chart folder; this reference does not apply.

What the binding changes:

- **Your edits target the running stack.** Saving updates the stack and the
  platform starts a deployment pipeline for it immediately — there is no
  separate "deploy" step to offer. A save never destroys anything: a
  resource you delete from the templates stays deployed, and stays the
  stack's, until the stack is undeployed. The organization's catalog
  policy never blocks these saves: it refuses only NEW creations of
  disabled kinds, and a deployed stack's Infra Components already exist —
  never warn about availability on a working-copy save
  (`catalog.availability.md`).
- **The folder is laid out fresh from the server** — by the app when the
  conversation starts, or by your own checkout — so what you read here is
  the deployed truth as of that moment, placed for THIS conversation alone.
- **The chart is out of the picture.** Never suggest publishing this folder
  as a chart or `chart publish` unless the user explicitly wants to turn
  their stack's state into a reusable chart.

## When you are NOT in a working copy: check one out

A conversation about a deployed stack does not always start inside its
working copy — in your own workspace, the request may simply name the
stack. Never reconstruct a stack's files from its record by hand; pull
the real thing:

```
planton infra stack checkout <id-or-name> --env <env> --output-dir <stack-dir>
```

A name alone is ambiguous when the chart is installed in several
environments; pass the environment the user means (an id needs none).

The checkout lays out the working copy — templates, values, and the hidden
binding — as its own top-level subfolder of your workspace, and everything
in this reference applies to it from that moment. It writes many files at
once, so the composing declaration (SKILL.md's live-screen rule) comes
FIRST. Re-running it against the same folder refreshes the managed files
from server truth and leaves any other files alone.

## Opening posture: diagnose FIRST when there is a failure

When the conversation opens from a failed pipeline (the opening message
names a pipeline id), **the four-step diagnosis is your FIRST ACT — run it
before any reply that merely acknowledges the request.** The user already
asked for the investigation by arriving here; a first turn that restates
the failure and waits is a failed turn. Consent gates CHANGES, never
diagnosis — reading status, logs, and records requires nobody's permission.

1. Run the four-step diagnosis from `craft.planton-cli.md` (status → failed
   node → logs → Infra Job when needed), dumping large records to
   `.scratch/` files and reading those (the large-records rule in
   `craft.planton-cli.md`).
2. Explain what went wrong in the user's language — name the component and
   the cause, not the plumbing.
3. Recommend the fix with a reason, and say what it will do when applied
   (updated stack, new pipeline, roughly how long).
4. **Ask what they want** — even when the fix is obvious. Modifying a
   deployed stack is always consent-gated.

The same compile loop applies while editing: `planton chart build . -o json`
works on a working copy exactly as on a chart folder (the binding file is
invisible to it). Keep the folder green before proposing to save.

## Saving: `chart install` IS the update

The Helm-consistent verb updates the existing stack when the name matches
in that environment — it never creates a duplicate. The same name with
another `--env` is a different stack: installing there creates (or updates)
that environment's own stack and never touches this one, and a stack's
environment can never be changed:

```
planton chart install <stackName> . --org <org> --env <env> \
  -m "one line saying WHY, like a commit message" --plain
```

Read every flag from the binding file — never guess the name, org, or env,
and never omit `--org`/`--env` (your CLI context may point elsewhere; the
stack's environment is not yours to move). The `-m` message is mandatory
practice: it becomes the run's name, where the whole team reads why it
happened.

What one invocation does, in order: builds the folder server-side (refusing
on errors, before anything is saved), prints where each param's value came
from, updates the stack, starts the deployment pipeline,
prints its id, and **follows the run to completion** — the exit code is
non-zero when the followed pipeline fails. That single command is your save
AND your monitor. To show the user exactly what a save would apply first,
run the same command with `--dry-run`: it prints every param and the
rendered documents and creates nothing.

A stack has one run at a time. A save while a run of the stack is queued,
running or waiting for approval is refused, naming that run ("… already
has a run in progress (<id>) …" or "… has a run waiting for approval
(<id>) …"). Never retry blindly: tell the user, and let them approve or
reject the waiting run, or cancel it (`planton infra pipeline cancel <id>`,
a platform mutation — confirm first), then save again.

When holding one long-running command open is unsuitable, split it:

```
planton chart install <stackName> . --org <org> --env <env> -m "…" --no-follow
planton infra pipeline status <infpipe_id>     # poll between other work
planton infra pipeline logs <infpipe_id>       # when a node fails
```

## What a run applies, and what it leaves alone

A stack run applies only the resources whose inputs moved since their last
successful deploy. Before a resource runs, the platform compares what it
would deploy (its settings, the values it reads from other resources, its
secrets' current versions and its variables, the account it deploys into,
the module version, whether anything it sits on was rebuilt) with what its
last successful deploy applied. An exact match runs no job: the node reads
**Unchanged since <job>**, naming the job that last applied it, and hands
that job's outputs to its dependents. So a save that changes one value runs
that resource, and a dependent only when a value it reads changed; a
re-apply with nothing changed runs nothing. Say so when you narrate: unchanged nodes are successes, not
skipped work.

Three more node outcomes to read correctly:

- **Unmanaged** — a resource whose lifecycle events are disabled
  (`planton.dev/infra-job.disable-on-lifecycle-events: "true"`). The stack
  records it and never applies or destroys it: a deploy leaves it alone,
  and an undeploy or purge releases it from the stack and leaves it in the
  cloud.
- **Owned elsewhere** — a component another live stack or service owns.
  An undeploy or purge leaves it in place for its owner. A deploy that
  would change it is refused with a sentence naming the owner ("… is owned
  by infra stack 'x', so this run may not change it -- change it through
  that stack."); relay it, and change the component through that owner.
- **A resource changed outside Planton** (edited in the cloud console,
  deleted by hand) is invisible to the comparison, because nothing Planton
  deploys moved — a plain re-apply reads it unchanged. Refresh that
  component first (`planton tofu refresh <ic_id>`, or `planton pulumi
  refresh`): a refresh that finds the drift clears what the component
  recorded as applied. Then re-apply the stack
  (`planton infra stack deploy <stackName> --env <env>`), and that
  resource runs its job. Both are platform mutations: confirm first.

## Narrate every transition — the user is watching the app

The app watches the stack's pipelines: the moment your install creates
one, the screen announces it and navigates to the run. Say what you did the
moment you did it, so the voice and the screen agree:

- After the install starts: state that the stack was updated and a new
  pipeline is running — and that they are being taken to it.
- On success: say it plainly, summarize what changed.
- On failure: diagnose again (same four steps), explain, recommend, ask.

## When the fix does not take

A correct fix can still fail with the SAME error as before. Verified failure
mode: an Infra Component whose deploy failed can be left in a failed state
that keeps feeding the old configuration to new pipeline runs — the stack's
record and even the component's record show your corrected values, yet the
run fails with the stale ones. Recognize it by the signature (identical
error after a save that provably changed the relevant field), then clear it:

```
planton search infra-components --org <org> -e <env>      # find the component id
planton get infra-component <ic_id> -o yaml               # confirm the record already carries your fix
planton infra component cleanup <ic_id>                   # clear the failed state (confirm with the user first)
```

Then save again (a fresh `chart install`, its own confirmation). Cleanup only
clears the failed deployment state — it does not touch cloud infrastructure —
but it is still a platform mutation: say what it does and get a yes.

A different non-taking failure has its own playbook: a rerun that fails
because a provider resource **already exists** (an earlier run created it in
the cloud, then died before the state recorded it). Do not delete-and-retry
— import the orphan into state with the platform's import commands. Read
`infra.state-import.md`.

## Iterating until green

Failures can be layered. After the first fix attempt, ask ONE question that
sets the collaboration mode for the rest of the session:

> "Want me to keep iterating until this deploys green — asking you before
> each save — or check in with you after every diagnosis?"

Honor the answer. Even in keep-iterating mode, EVERY save still gets its own
explicit confirmation (what will change and why, one line each); what the
mode changes is whether you pause after each diagnosis or proceed to the
proposed fix directly.

## Boundaries

- Never `chart install` without the user's go-ahead for that specific change.
- Never change the target env/org to "make it work" — a failure that smells
  like wrong-environment is a finding to report, not to route around.
- Undeploying, purging, or deleting the stack is out of scope here —
  surface the need, let the user drive it in the product. What to tell them
  when they ask: undeploy destroys everything the stack owns, including
  components its templates no longer declare, dependents first; unmanaged
  resources are released and stay in the cloud. It is refused while any
  owned resource carries `planton.dev/prevent-destroy: "true"` (lift it by
  removing the annotation from the templates and saving) or while another
  stack still reads one of its resources (the refusal names it). Purge is
  undeploy, then delete, only if the undeploy succeeded. Delete alone is
  refused while the stack still owns components or has a run in flight.
