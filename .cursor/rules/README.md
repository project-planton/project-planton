# Cursor Rules: How to use Forge rules in this repo

This directory contains automation rules used by Cursor to scaffold and validate ProjectPlanton resources. You can invoke any rule in chat by mentioning it with an at-annotation (e.g., `@001-spec-proto`).

## Ways to run rules

1) Run a single rule in a chat
- Start a new chat and write something like:
  - `@001-spec-proto Add AWS CloudFront` (the first rule needs the provider/kind context)
- Subsequent steps in the same chat can be invoked simply by rule name:
  - `@002-spec-validate`
  - `@003-spec-tests`

2) Use the Mega Rule (recommended for convenience)
- Use `@forge` to orchestrate all steps in sequence (001 → 015) for a new resource.
- This is the most convenient way when you want end-to-end scaffolding with minimal prompts.

3) Split runs across two chats (recommended balance)
- Chat A (Rules 001–008)
  - Kick off with the provider/kind context once: `@001-spec-proto Add AWS CloudFront`
  - Then: `@002-spec-validate`, `@003-spec-tests`, `@004-stack-outputs`, `@005-api`, `@006-stack-input`, `@007-build-and-docs`, `@008-hack-manifest`
- Chat B (Rules 009–015)
  - Then: `@009-pulumi-module`, `@010-pulumi-entrypoint`, `@011-pulumi-e2e`, `@012-pulumi-docs`, `@013-terraform-module`, `@014-terraform-e2e`, `@015-terraform-docs`
- Benefit: Clear separation between API scaffolding and IaC steps; fewer long steps per chat while still retaining convenience.

## Rule index (Forge)

- 001-spec-proto.mdc — Generate spec.proto (no validations)
- 002-spec-validate.mdc — Add buf.validate rules and CEL
- 003-spec-tests.mdc — Add and run spec-level tests
- 004-stack-outputs.mdc — Generate stack_outputs.proto
- 005-api.mdc — Generate api.proto wiring metadata/spec/status
- 006-stack-input.mdc — Generate stack_input.proto
- 007-build-and-docs.mdc — Generate BUILD.bazel, README.md, examples.md (root resource folder)
- 008-hack-manifest.mdc — Generate top-level iac/hack/manifest.yaml
- 009-pulumi-module.mdc — Scaffold pulumi module/ (controller, locals, outputs, resources)
- 010-pulumi-entrypoint.mdc — Add pulumi entrypoint, Pulumi.yaml, Makefile, BUILD
- 011-pulumi-e2e.mdc — Pulumi end-to-end (make local; preview/update/refresh/destroy)
- 012-pulumi-docs.mdc — Pulumi README, examples.md, debug.sh
- 013-terraform-module.mdc — Terraform tf/ (variables via CLI generator, main/locals/outputs/provider)
- 014-terraform-e2e.mdc — Terraform end-to-end (make local; init/plan/apply/destroy)
- 015-terraform-docs.mdc — Terraform README (local backend) and examples.md

## Using the Mega Rule

- Invoke `@forge` in a new chat to run 001 → 015 automatically.
- Only the first step (001) requires the provider/kind context, e.g.: `@001-spec-proto Add AWS CloudFront`.
- After that, the Mega Rule passes context across steps.

## Tips
- Start a new chat when you switch to a different resource to keep context clean.
- If a step fails (lint/build/test), the rule will auto-refine up to 3 times. Provide the minimal additional hint if asked.
- For e2e rules, local CLI must align with current protos. The rules already call `make local` before running CLI operations.
