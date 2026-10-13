# Deployed Infra Stacks — Working Copies, Plans, Fixing Failures

A chart is a blueprint; an **Infra Stack** is one chart declared into one
environment. A stack is named within its environment: the same chart
installed into qa and prod is two stacks that share a name, so every
command that names a stack also names its environment (`--env`). A stack
carries its OWN declaration — the templates, the chart's `values.yaml`
(every value declared, with its default) and the values set for its
environment — detached from the chart it came from. Editing the chart
never changes a deployed stack, and editing a stack never changes the
chart. This file governs the second mode of your work: changing something
that is already running.

Every change to a stack is a **plan** a person reads first. A plan says
what applying would do to each resource, before anything runs; applying
it runs exactly that plan, and only while nothing it was made from has
moved. Each applied plan is a numbered **revision** ("r12 · Raise
replicas").

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

A working copy holds:

| File | What it is |
|------|------------|
| `templates/**` | The stack's templates. |
| `values.yaml` | The chart's declarations and defaults, as the stack carries them. |
| `values.<env>.yaml` | Only the values set for the stack's environment (a `params:` list of name and value). Anything not listed takes the default from `values.yaml`. |
| `.planton/stack.yaml` | The binding. |

To change a value for this environment, edit `values.<env>.yaml`; to
change a default every environment shares, edit `values.yaml` (a plan
shows such a change as a value change, marked "the chart's default
changed").

What the binding changes:

- **Your edits target the running stack** — through a plan you show the
  user before anything applies (below). The organization's catalog policy
  never blocks these changes: it refuses only NEW creations of disabled
  kinds, and a deployed stack's Infra Components already exist — never
  warn about availability on a working copy (`catalog.availability.md`).
- **The folder is laid out fresh from the server** — by the app when the
  conversation starts, or by your own checkout — so what you read here is
  the deployed truth as of that moment, placed for THIS conversation alone.
- **The chart is out of the picture.** Never suggest publishing this folder
  as a chart or `chart publish` unless the user explicitly wants to turn
  their stack's declaration into a reusable chart.

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

The checkout lays out the working copy as its own top-level subfolder of
your workspace, and everything in this reference applies to it from that
moment. It writes many files at once, so the composing declaration
(SKILL.md's live-screen rule) comes FIRST. Re-running it against the same
folder refreshes the managed files (`Chart.yaml`, `values.yaml`,
`values.<env>.yaml`, `templates/`, `.planton/`) from server truth and
leaves any other files alone.

## Opening posture: diagnose FIRST when there is a failure

When the conversation opens from a failed run (the opening message names a
pipeline id), **the four-step diagnosis is your FIRST ACT — run it before
any reply that merely acknowledges the request.** The user already asked
for the investigation by arriving here; a first turn that restates the
failure and waits is a failed turn. Consent gates CHANGES, never
diagnosis — reading status, logs, records and plans requires nobody's
permission.

1. Run the four-step diagnosis from `craft.planton-cli.md` (status → failed
   node → logs → Infra Job when needed), dumping large records to
   `.scratch/` files and reading those (the large-records rule in
   `craft.planton-cli.md`).
2. Explain what went wrong in the user's language — name the component and
   the cause, not the plumbing.
3. Recommend the fix with a reason, and show its plan (below).
4. **Ask what they want** — even when the fix is obvious. Changing a
   deployed stack is always consent-gated.

The same compile loop applies while editing: `planton chart build . -o json`
works on a working copy exactly as on a chart folder (the binding file is
invisible to it). Keep the folder green before you plan.

## Plan, show, consent, apply

**1. Plan.** Planning writes nothing, so do it freely:

```
planton infra stack plan <stackName> . --env <env> --org <org>
```

Read every flag from the binding file — never guess the name, org, or env,
and never omit `--org`/`--env` (your CLI context may point elsewhere; the
stack's environment is not yours to move). With no folder argument it
plans the declaration the stack already holds — what moved underneath it
(a module release, a secret's new version, an upstream output). `--set
name=value`, `--set-file name=path` and `-f <file>` change values on top of
the ones the environment already sets.

The plan lists each resource by what applying does to it, in reading
order, each with its reason:

| Row | Meaning |
|-----|---------|
| **Blocked** | It cannot apply as declared; the row says why and the way out. |
| **Destroy** | No longer in the templates: destroyed in the cloud once everything the plan applies has applied. |
| **Update** | Its settings, a secret's version, a value it reads or its module moved. |
| **Update Record** | Only Planton's record changes; nothing in the cloud (an annotation like `planton.dev/prevent-destroy` added). |
| **Create** | New. |
| **Adopt** | It already exists and no stack manages it; this stack takes it over only when chosen, and the apply may change it to match the template. |
| **Keep in Cloud** | No longer in the templates, and kept: the stack lets go of it, and it keeps running and billing in the account. |
| **May Change** | It reads a value known only after an earlier resource applies. |
| **Always Runs** | Its job runs on every apply: whether it changes is known only when it runs. |
| **Unchanged** / **Unmanaged** | Nothing moved / the platform records it and never applies it. Collapsed into one line. |

A removed resource whose destroy is blocked — protected by
`planton.dev/prevent-destroy`, its own deletion protection on, or still
read by another stack or by a resource the templates declare — blocks only
its destroy: choosing Keep in Cloud applies the plan around it. A Blocked
row on a resource the templates declare (another stack owns it, a secret
it reads is missing) stops the whole plan until it is resolved.

Exit codes: `0` planned; `1` the declaration does not build, a resource it
declares is blocked, or a run is in flight; `2` the plan could not be
made. With `--detailed-exitcode`, `4` means applying would change
something. `-o markdown` prints the plan for a pull request or an answer;
`-o json` for a program.

**2. Show.** Put the plan in front of the user in their words: what is
created, updated and destroyed — **every Destroy row by name** — and what
it costs them if a resource holds data. A plan of the declaration the
stack already holds also has a page in the console (the plan prints its
link); a plan of an edited folder lives only in your output.

**3. Consent.** Ask, and name each destroy. Consent to destroy is per
resource: the user agrees to destroy THIS bucket, not "whatever the plan
says". A resource that holds data a destroy cannot give back is marked in
the plan and is destroyed only when named (`--destroy Kind/name`).
Anything the user wants kept goes `--keep-in-cloud Kind/name`; an existing
resource the stack should take over goes `--adopt Kind/name` — a stack
never takes over a resource nobody chose to hand it.

**4. Apply** exactly what was shown:

```
planton infra stack apply <stackName> . --env <env> --org <org> \
  -m "one line saying WHY, like a commit message" \
  --plan-digest <digest from the plan you showed> \
  --destroy AwsS3Bucket/old-logs --keep-in-cloud AwsRoute53Zone/legacy --plain
```

The apply plans again and goes ahead only while the fresh plan is the one
named by `--plan-digest` — otherwise it refuses, saying the plan is out of
date; plan again and show the new plan. Without a terminal (your session,
CI), a destroy happens only when named with `--destroy` or covered by the
digest of a plan a person reviewed; a data-bearing destroy is always
named. `--plan-digest` or any `--destroy` also answers the final question;
`--yes` alone answers it but consents to no destroy. The `-m` message is
mandatory practice: it names the revision and its run ("r13 · Raise
replicas"), where the whole team reads why it happened.

What one apply does: records the declaration and a numbered revision, then
starts ONE run that applies what the stack declares and — only when all
of it applied — removes what the plan removed (destroyed, or kept in the
cloud), prints the run's id, and **follows it to completion**; the exit
code is non-zero when the followed run fails. A plan that changes only the
declaration (a comment in a template, a value set to its default) records
the revision and starts no run. A plan that changes nothing applies
nothing: "No changes".

`planton chart install <stackName> . --env <env>` takes the same path (plan,
show, ask, apply), and creates the stack when the name is new; `--dry-run`
stops at the plan. `planton apply -f <stack manifest>` applies a plan only
when nothing in it is destroyed, adopted or blocked; otherwise it refuses
and names the `planton infra stack apply` command to review it with.

A stack has one run at a time. A plan of a stack whose run is in flight
answers with that run instead of rows, and an apply is refused naming it
("… already has a run in progress (<id>) …" or "… has a run waiting for
approval (<id>) …"). Never retry blindly: tell the user, and let them
approve or reject the waiting run, or cancel it (`planton infra pipeline
cancel <id>`, a platform mutation — confirm first), then plan again.

When holding one long-running command open is unsuitable, split it:

```
planton infra stack apply <stackName> . --env <env> -m "…" --plan-digest <d> --no-follow
planton infra pipeline status <infpipe_id>     # poll between other work
planton infra pipeline logs <infpipe_id>       # when a node fails
```

`planton infra stack history <stackName> --env <env>` lists the revisions,
newest first: number, message, outcome (Applying, Applied, Applied With
Removal Failures, Failed, Cancelled), who applied it, and its run; the live
revision — the newest that applied — is marked.

## What a run applies, and what it leaves alone

A stack run applies only the resources whose inputs moved since their last
successful deploy — the same comparison the plan shows as Unchanged. An
exact match runs no job: the node reads **Unchanged since <job>**, naming
the job that last applied it, and hands that job's outputs to its
dependents. So a change to one value runs that resource, and a dependent
only when a value it reads changed; a re-apply with nothing changed runs
nothing. Say so when you narrate: unchanged nodes are successes, not
skipped work. `--reapply Kind/name` (or `--reapply-all`) runs a job anyway.

More node outcomes to read correctly:

- **Unmanaged** — a resource whose lifecycle events are disabled
  (`planton.dev/infra-job.disable-on-lifecycle-events: "true"`). The stack
  records it and never applies or destroys it: removing it from the
  templates plans Keep in Cloud, and an undeploy releases it.
- **Kept in Cloud** — a removed resource the user chose to keep, or one
  marked `planton.dev/keep-in-cloud: "true"`: the stack let go of it, no job
  ran, and it keeps running; another stack or a service can adopt it.
- **Owned elsewhere** — a component another live stack or service owns.
  A teardown leaves it in place for its owner; a plan that would change it
  is Blocked, naming the owner. Change it through that owner.
- **A resource changed outside Planton** (edited in the cloud console,
  deleted by hand) is invisible to the comparison, because nothing Planton
  deploys moved — a plain re-plan reads it Unchanged. Refresh that
  component first (`planton tofu refresh <ic_id>`, or `planton pulumi
  refresh`): a refresh that finds the drift clears what the component
  recorded as applied. Then plan the stack again (`planton infra stack plan
  <stackName> --env <env>`) — that resource plans Update — and apply it.
  Both are platform mutations: confirm first.

A run whose declared resources applied but a removal failed ends
**"Applied; 1 removal failed"**: what the stack declares is live, and the
revision reads Applied With Removal Failures. Diagnose the failed destroy
like any failed node.

## Narrate every transition — the user is watching the app

The app watches the stack's runs: the moment your apply starts one, the
screen announces it and navigates to the run. Say what you did the moment
you did it, so the voice and the screen agree:

- After the apply: say which revision was recorded and that its run is
  running — and that they are being taken to it.
- On success: say it plainly, summarize what changed.
- On failure: diagnose again (same four steps), explain, recommend, ask.

## When the fix does not take

A correct fix can still fail with the SAME error as before. Verified failure
mode: an Infra Component whose deploy failed can be left in a failed state
that keeps feeding the old configuration to new runs — the stack's record
and even the component's record show your corrected values, yet the run
fails with the stale ones. Recognize it by the signature (identical error
after an apply that provably changed the relevant field), then clear it:

```
planton search infra-components --org <org> -e <env>      # find the component id
planton get infra-component <ic_id> -o yaml               # confirm the record already carries your fix
planton infra component cleanup <ic_id>                   # clear the failed state (confirm with the user first)
```

Then have the user rerun the run (Rerun on its page in the console): a
rerun repeats the same revision, so what applied stays Unchanged and what
failed runs again. A rerun is refused once a newer revision exists or the
stack was undeployed since — plan again and apply instead. Cleanup only clears the failed
deployment state — it does not touch cloud infrastructure — but it is
still a platform mutation: say what it does and get a yes.

A different non-taking failure has its own playbook: a rerun that fails
because a provider resource **already exists** (an earlier run created it in
the cloud, then died before the state recorded it). Do not destroy-and-retry
— import the orphan into state with the platform's import commands. Read
`infra.state-import.md`.

## Iterating until green

Failures can be layered. After the first fix attempt, ask ONE question that
sets the collaboration mode for the rest of the session:

> "Want me to keep iterating until this deploys green — showing you each
> plan before I apply it — or check in with you after every diagnosis?"

Honor the answer. Even in keep-iterating mode, EVERY apply still gets its
own explicit consent on its plan (what will change and why, one line each,
every destroy by name); what the mode changes is whether you pause after
each diagnosis or proceed to the proposed fix's plan directly.

## Boundaries

- Never apply without the user's go-ahead for that specific plan.
- Never change the target env/org to "make it work" — a failure that smells
  like wrong-environment is a finding to report, not to route around.
- Undeploying, purging, or deleting the stack is the user's to drive in the
  product; surface the need. What to tell them when they ask: undeploy
  destroys everything the stack owns, including components its templates
  no longer declare, dependents first, and keeps the stack and its
  revisions; unmanaged resources and ones marked `planton.dev/keep-in-cloud`
  are released and stay in the cloud. It is refused while any owned
  resource carries `planton.dev/prevent-destroy: "true"` (lift it by
  removing the annotation and applying) or while another stack still reads
  one of its resources (the refusal names it). Applying the stack again
  later plans each resource anew — any the undeploy released plans Adopt.
  Purge is undeploy, then delete, only if the undeploy succeeded. Delete
  alone is refused while the stack still owns components or has a run in
  flight. From a session, each of these runs only with `--yes`, and only
  after the user said yes.
