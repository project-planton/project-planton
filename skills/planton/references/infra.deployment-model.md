# The Deployment Model — What Happens After Compose

You compose the chart; the platform runs it. Knowing the machinery lets you
set expectations and diagnose failures — but it is background knowledge, not
conversation material. **Educate only when it serves the user's next step**
(they clicked deploy and wonder what is happening; a pipeline failed and you
are explaining where; they ask). Never lecture the model at someone who just
wants their infrastructure up.

## The path from chart to real infrastructure

```
InfraChart (the template you compose)
  → deploy: InfraStack (the chart rendered with the user's values — the
    record of exactly what was requested)
  → InfraPipeline (executes the stack's dependency graph in order)
  → each pipeline node = one InfraComponent (one manifest from the chart)
  → each node's deploy = an Infra Job (one IaC engine run for that component)
  → real cloud infrastructure
```

Practical implications worth sharing at the right moment:

- Deploying a chart does not "run the chart" — it creates an Infra Stack, whose
  pipeline deploys component by component in dependency order. Independent
  components run in parallel, at most four at once per run; a failed node
  ends everything downstream of it ("Not run: <Kind>/<slug> failed") while
  independent branches carry on.
- Every change to a stack is a plan read first: what applying would do to
  each resource (Create, Update, Destroy, Keep in Cloud, Adopt, Unchanged…),
  before anything runs. Applying records a numbered revision and starts one
  run that applies what the stack declares, then removes what it no longer
  declares — destroyed, or kept in the cloud as chosen.
- A stack has one run at a time: an apply, undeploy, purge or rerun is
  refused while a run of that stack is queued, running or waiting for
  approval, and the refusal names the run. Relay it: approve or reject that
  run, or cancel it (`planton infra pipeline cancel <id>`), then retry.
- Undeploy tears down everything the stack owns — including components its
  templates no longer declare — dependents first, so a cluster outlives the
  workloads on it. It is refused while any owned resource carries
  `planton.dev/prevent-destroy: "true"`, or while another stack still reads
  one of its resources (the refusal names that stack).
- A failed pipeline is diagnosed node by node: find the failed node, read its
  Infra Job's error and logs (`craft.planton-cli.md` has the exact commands).
- Applying a chart fix creates a new run; resources whose inputs did not
  move run no job and read "Unchanged since <job>" (no duplicate
  infrastructure). A resource the templates stop declaring is destroyed
  only when its plan says Destroy and the person agreed — or kept running
  in the cloud when they chose Keep in Cloud.

## Every resource deploys through an open-source module

Each Infra Component is deployed by its kind's IaC module in the open-source
repository `github.com/plantonhq/planton`, at:

```
catalog/<provider>/<kind-lowercase>/iac/
  ├── tf/       # the Terraform/OpenTofu module
  └── pulumi/   # the Pulumi (Go) module
```

Which engine runs is the deploying organization's setting — **the default is
OpenTofu (`tf/`) unless the org or manifest explicitly says otherwise**. Never
set a provisioner in a chart.

The module source is the ground truth for what a spec field actually DOES in
the cloud — one level deeper than the schema. Use it when:

- A spec field's cloud-side effect is ambiguous and the choice matters.
- A deploy failed inside the engine and the error names cloud-provider
  concepts the manifest never mentions — map the failing provider resource back
  to the module code that creates it, then back to the spec field feeding it.
- You need to know a default the schema does not state (what the module does
  when a field is empty).

To inspect: ask the user's permission ONCE per session to clone the
open-source repo read-only (`git clone --depth 1
https://github.com/plantonhq/planton /tmp/planton-oss` — say why), then read
the kind's `iac/tf/` (or `iac/pulumi/` when that engine is in play). Cloning
writes only to the local temp directory and touches no infrastructure; it
still gets one honest ask because it downloads a large repository onto their
machine. Combine what the module does with what the user wants to produce the
most accurate chart — this is what separates a partner from a form-filler.
