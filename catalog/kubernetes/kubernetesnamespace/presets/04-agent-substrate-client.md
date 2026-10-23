# Agent Substrate Client Namespace

This preset creates the namespace a service that calls a KubernetesAgentSubstrate runs in, labeled so the substrate's NetworkPolicy admits its API calls.

## When to Use

- The namespace of a product's server that creates and calls actors (pair it with the KubernetesDeployment preset **05-agent-substrate-client**)

## Key Configuration Choices

- **`ate.dev/api-access: "true"`** — the label the substrate exports as `api_access_namespace_label`; without it the API's NetworkPolicy drops every connection from the namespace
- **Baseline pod security** — the client is an ordinary service; nothing here needs privileges

## Placeholders to Replace

| Placeholder | Description | Where to Find |
|---|---|---|
| `my-substrate-client-namespace` | The namespace name | Your naming convention |

## Related Presets

- **01-standard** — a namespace with no substrate access
