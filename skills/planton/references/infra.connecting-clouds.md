---
title: Connecting a Cloud — Keyless With the Person's Own Sign-In, Explained the Whole Way
description: How an agent connects AWS, Google Cloud, Azure, DigitalOcean or Cloudflare to Planton for the person -- keyless with the sign-in already on their machine (`planton connect gcp`, `aws`, `azure`), a token for DigitalOcean and Cloudflare -- and keeps them informed and in control throughout -- what each instance offers, the plan to show before anything changes, the facts that answer every worry about access, the grant and its two opt-in tiers, reading the verdict, resuming, verify and cleanup, and every refusal with its fix. Read when the person asks to connect a cloud account, asks what access Planton needs or has, or a connection cannot deploy something.
---

# Connecting a Cloud

Read this when a person asks you to connect their cloud to Planton, asks what Planton can do in their account, or a connection cannot deploy something they asked for.

## The doctrine in one sentence

The person decides, fully informed: you explain what Planton will be able to do and that the setup runs as them in their own cloud, show the exact plan, ask, run it only on their yes, narrate each step, read the verdict back -- and every reassurance you give is a fact the product makes true, never a comfort you made up.

## Why the person is right to ask questions

A connection gives Planton a way into their cloud. A careful engineer, and every security lead, will want to know exactly what that way allows and how to take it back. Treat each question as the most important part of the job, not an obstacle to it. Connecting is half the work; the person understanding what they connected is the other half.

## What each Planton offers

| Where Planton runs | What the bare `planton connect gcp` (or `aws`, `azure`) does |
|---|---|
| planton.ai | Keyless: the trust is written in the person's cloud with their sign-in, Planton stores no key |
| A self-hosted Planton with a public address | Keyless, the same way |
| Planton Desktop's own instance, or a self-hosted Planton with a private address | Refuses keyless with the instance's reason, and names `planton connect gcp detect` (the Runner Identity method: deploys run with this machine's sign-in directly) |

DigitalOcean and Cloudflare have no keyless trust: the connection holds an API token the person creates (below). The four methods are named the same everywhere: Keyless, Stored Key, Runner Identity, Vault Broker.

## The journey, in order

1. **Explain first, in plain words** -- before any command that could change something. Use the facts in "What to tell the person" below.
2. **Show the plan.** `planton connect gcp --project <project> -o json` prints the plan and changes nothing (`"written": false`). Its keys: `runsAs` (the identity the setup acts as), `target` and `targetNoun` (the project, account or subscription it acts in), `connection` (the name and slug it will create), `grant` and `kindsGranted` of `kindsInCatalog` (what Planton may deploy), and per script `keeps` (what it leaves untouched) and `reachesOthers` (what it also changes for other connections). Tell the person each of these, in their words. `readScript` is the command that prints the exact script, if they want to read it.
3. **Ask.** Name the identity and the target in the question: "This runs as dev@acme.com in project acme-prod and creates a role only Planton's connection can use. Go ahead?"
4. **Run on their yes:** the `run` command from the plan (it carries `--yes`). The script runs on this machine, as them; narrate as it goes: the trust, the role, waiting for the cloud to accept the trust (a new trust can take a few minutes to spread), the connection, Verify.
5. **Read the verdict back** (below), and say what they can deploy now.

A person at the terminal can simply run `planton connect gcp`: it shows the same plan as a card and asks the same question.

- **`--yes` is the person's yes, given ahead.** Pass it only after they agreed to this plan.
- **`--accept-shared-changes`** is needed with `--yes` when `reachesOthers` is not empty: the role is bound to an identity other connections also sign in as, so the change reaches them too. Pass it only after the person accepted that consequence by name.
- **Target and sign-in flags:** GCP `--project` and `--configuration` (a gcloud named configuration); AWS `--account-id` and `--profile`; Azure `--subscription` (by id or name). Without them, the setup uses the cloud CLI's active sign-in and its default project, account or subscription -- never another sign-in it happens to find. `--name` sets the connection's name (default: `gcp.project.<project>`, `aws.account.<id>`, `azure.subscription.<id>`).

## What to tell the person -- every sentence is a product fact

- **It runs as you, on your machine.** The setup script runs here with the sign-in you already have. Your sign-in never leaves this machine; Planton receives only what the script reports back (on GCP the trust's provider name and the service account, on AWS the role's ARN, on Azure the app's client id, tenant and subscription).
- **Planton keeps no key.** A keyless connection stores no secret. Each deploy gets a short-lived token that only this connection's trust accepts.
- **It can do exactly what the grant says, nothing more.** The setup creates the connection's own role with exactly the permissions the catalog lists for the kinds granted. "Everything Planton can deploy" never includes Owner or Editor, never power over who has access in the account (IAM policies, roles, users, keys), and never permissions above the target (an organization, a folder, a billing account) -- those two tiers are separate opt-ins, `--access-granting` and `--organization-access`, each off unless the person asks.
- **You stay in control, in your own cloud.** The role is theirs: they can read it, narrow it or delete it at any time in their cloud's console. Verify always says what the connection can actually do after a change.
- **Taking it back is one script and one delete.** `planton connect gcp cleanup <slug>` prints what deleting the connection leaves in their cloud and the script that removes the trust and the role Planton set up; `planton delete GcpProviderConnection <slug>` deletes the connection.
- **The script is readable before it runs.** `planton connect gcp --project <project> --print-documents` prints it exactly as it will run, with every permission written out.

## The common worries, answered truthfully

| The person asks | The answer, and where they see it |
|---|---|
| "What exactly can Planton do in my account?" | Deploy, change and destroy the kinds the grant covers, with exactly the permissions the catalog lists for them. The plan's `grant` and `kindsGranted` say which; `--print-documents` shows every permission. |
| "Does Planton keep my credentials?" | No. Keyless stores nothing; each deploy gets a short-lived token the trust accepts. Their own sign-in never leaves the machine. (DigitalOcean and Cloudflare are the exception: their token is stored as an organization secret, because those clouds have no keyless trust.) |
| "Can Planton read my data?" | Answer from the documents, never with a blanket no: the grant includes what the kinds need to manage their resources, which for some kinds includes reading that resource's contents (a secret store's versions, for example). Show the permissions with `--print-documents` and name any that read data. |
| "How do I limit it?" | Grant only some kinds (`--kinds GcpGkeCluster,GcpGcsBucket`), or narrow the role in their cloud later; Verify then says which kinds can no longer deploy. A recorded grant is changed with Update the Grant in the connection's Permissions section. |
| "How do I revoke it?" | `planton connect gcp cleanup <slug>` for the removal script, then `planton delete GcpProviderConnection <slug>`. Deleting the trust in their cloud cuts Planton off at once. |
| "Does this touch my other roles or connections?" | The plan's `keeps` lists what stays untouched; `reachesOthers` lists what the change does to other connections that sign in as the same identity -- read it to them before asking. |
| "Why does it need this permission?" | Because a kind in the grant needs it to deploy: the catalog's permission files list each kind's permissions by purpose, and the connection's Permissions section in the console shows which kind needs what. |
| "Who in my organization can use this connection?" | By default it is authorized for every environment of the organization and becomes the organization's default for that cloud when none exists (never replacing one). `--no-authorize` and `--no-default` leave both alone. |

## The grant

- **Default: everything Planton can deploy** for that cloud -- every catalog kind, and every way Planton itself uses a connection.
- **`--kinds`** narrows it to the kinds named, as the catalog spells them (`GcpGkeCluster`, `AwsVpc`).
- **`--organization-access`** adds the permissions above the target that organization-level kinds need. **`--access-granting`** adds power over who has access in the account. Each is the person's explicit choice; say what it gives before passing it.
- **Running the command again keeps the connection's recorded grant.** To change it, the person uses Update the Grant in the connection's Permissions section.

## Reading the verdict

`-o json --yes` ends with `verdict`; a person sees the same as a headline and lists.

- `headline`: "171 of 186 GCP Kinds Can Deploy", counted against the grant.
- `blocked`: kinds the grant covers that cannot deploy, each with the permissions it is missing. The way forward is to apply the grant again from the connection's Permissions section, never to build a second role by hand.
- `outsideGrant`: kinds the person did not grant. Not a fault -- what they chose.
- `canDeploy`, `usesNotReady` (Planton's own uses the identity cannot serve yet), `catalogRelease`.
- `healthy: false` with `reason` and `providerReason`: Planton's sentence and the cloud's own words. `providerReasonCode: "Throttling"` means the cloud rate-limited Verify: run `planton connect aws verify <slug>` again in a minute.

`planton connect gcp verify <slug>` re-runs Verify any time: after the person changes the role, or to see the issuer and subject the trust must pin.

## When something stops

Every failure after the script ran names the command that resumes; running the same command again picks up where it stopped: the script adopts what it already made, and an existing keyless connection with the same trust goes straight to Verify.

| The refusal | What to do |
|---|---|
| Keyless Isn't Offered Here | This Planton cannot offer keyless (its reason is in the sentence). Offer `planton connect gcp detect` (Runner Identity) and say what differs: deploys run with this machine's sign-in, with that sign-in's full access. |
| A lapsed or missing sign-in (detection fails) | The sentence names the sign-in it tried. The person runs `gcloud auth login`, `aws sso login --profile <profile>` or `az login`, then the same command. |
| Run the Setup in Your Cloud's Shell (no bash) | The machine has no bash (Windows without Git Bash). Print the script with `--print-documents` and have the person run it in Google Cloud Shell, AWS CloudShell or Azure Cloud Shell, which have bash and are already signed in. |
| Install gcloud / aws / az | The cloud's CLI is missing. Install it and sign in, or use the cloud's own shell as above. |
| The Setup Stopped | The script's last line is the cloud's own reason. A permission refusal means the person's sign-in cannot create the trust: on GCP it must be able to create workload identity pools, service accounts and custom roles, bind roles on the project and turn on APIs (a project Owner can); on AWS create IAM identity providers, roles and policies; on Azure register an app in Entra ID and create and assign roles in the subscription (Owner or User Access Administrator). Nothing was created in Planton. |
| The script says the AWS CLI is signed in to another account | The profile belongs to a different account than the target. Pick the right one with `--profile`. |
| The script says the Azure CLI works in another tenant | Switch with `az account set --subscription <id>` (the person's choice to make), then run again. |
| The Trust Isn't Working | The cloud refused the connection's token; the sentence and the cloud's reason say why. The trust is set up; after fixing, the same command finishes. |
| Choose Another Name | The name belongs to a different connection. Pass `--name`. |
| Change the Grant From the Connection | A re-run keeps the recorded grant. Use Update the Grant in the connection's Permissions section. |
| This Setup Changes Other Connections | `--yes` without `--accept-shared-changes` while the change reaches other connections. Tell the person which, and ask. |

## DigitalOcean and Cloudflare: a token

`planton connect digitalocean --print-documents` (or `cloudflare`) prints the scopes the token needs and where to create it. The person creates the token and passes it in `DIGITALOCEAN_TOKEN` or `CLOUDFLARE_API_TOKEN`, or on stdin -- never as a flag, and never pasted into the conversation if you can avoid it. Then `planton connect digitalocean --region nyc3` (DigitalOcean needs the region resources deploy to unless they name one) stores it as the organization secret Planton Desktop uses too, connects, makes the connection ready and verifies it. Connecting again with a new token rotates it. Revoking the token on the provider's token page cuts Planton off.

## What never to do

- Never pass `--yes` or `--accept-shared-changes` without the person's yes to that exact plan.
- Never edit the server's script, and never compose a provider name, an ARN or a subject yourself: the script reports them.
- Never put a token in a command-line flag or a file in the repository.
- Never tell the person Planton "can't read anything" or "only has read access": answer from the permissions the documents list.
- Never switch the person's default subscription, project or profile for them; name the command and let them choose.
