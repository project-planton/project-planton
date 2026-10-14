# An Unchanged Chart or Module Plans No Change

## Charts: a param without a default renders a typed stand-in

`planton chart validate` used to render a param with no default as the literal
`<name>`. No name rule accepts a leading `<`, so a chart whose resource names
come from a value every stack sets failed schema validation, even when nothing
was wrong with it.

Such a param now renders a typed stand-in, the same one the platform's chart
build binds:

| Type | Stand-in |
|---|---|
| string | the param's name as a slug (`bucket_name` renders `bucket-name`) |
| string_enum | its first allowed value |
| number | 1, because count-like fields refuse zero |
| bool | off |
| list | empty |

Each stand-in is the plainest value of its type that a kind's schema accepts
most often. The chart renders whole, and its verdict stays about the chart.
The gate also names each such param once, as a warning against `values.yaml`:
"Every stack sets `bucket_name` -- the chart declares no default". This
replaces the per-template "has no default value" warning.

## Three modules stop planning a change on every run

- **CloudflareWorker** defaulted an unset `compatibility_date` to the day of the
  run: `timestamp()` in OpenTofu and `time.Now()` in Pulumi. An unchanged
  Worker planned an update every day. Both engines now default it to the fixed
  date `2025-01-15`, which the spec declares as its recommended default.
- **AzureCosmosdbSqlRoleDefinition** and **AzureCosmosdbSqlRoleAssignment**
  (Pulumi) drew a fresh `uuid.New()` for an unpinned GUID on every run, so each
  run planned a replacement. The GUID is now derived from the resource's
  identity: a SHA-1 UUID over a per-module namespace and the metadata id, or
  the Pulumi stack when no id is present. The OpenTofu modules already kept the
  GUID that azurerm draws at create, so both engines now keep one GUID for the
  resource's life.

An existing Pulumi role definition or assignment without a pinned GUID is
replaced once, on its first run with this release, and keeps its GUID after
that.
