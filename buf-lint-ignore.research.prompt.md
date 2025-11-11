# Research Prompt: Buf Lint Configuration for Bazel-Managed Go Monorepo

## Project Context

**Environment:**
- Bazel 8.4.2 with bzlmod (MODULE.bazel)
- Buf CLI 1.55.1 (v2 configuration)
- Go 1.24.0 monorepo with protobuf definitions
- Gazelle for BUILD.bazel generation (proto generation disabled)
- Rules_go for Go compilation
- Nogo for static analysis

**Directory Structure:**
Project root contains:
- `provider/` - 519+ proto files organized by cloud provider (aws/, gcp/, kubernetes/, etc.)
- `shared/` - 16 proto files for common types
- Bazel symlinks: `bazel-project-planton`, `bazel-bin`, `bazel-out`, `bazel-testlogs` (pointing to `/private/var/tmp/_bazel_*/execroot/_main/`)
- Generated code in `pkg/`, Go source in `cmd/` and `internal/`
- Build artifacts in `build/`, docs in `site/` and `docs/`

**Current Configuration:**
- `buf.yaml` (v2) at project root
- Module name: `buf.build/project-planton/apis`
- Dependency: `buf.build/bufbuild/protovalidate`
- Custom plugin: `buf.build/project-planton/optional-linter:v0.1.0`
- `.bufignore` file created with 18 directory patterns

## The Problem

When running `buf lint`, it scans Bazel's symlinked external dependencies at `bazel-project-planton/external/` which contain:
1. Duplicate proto definitions from test data (gazelle, rules_go, buildtools)
2. Intentionally broken/test proto files
3. Proto files with duplicate symbols across different external repos

**Error Pattern:**
```
bazel-project-planton/external/gazelle++go_deps+com_github_bazelbuild_buildtools/deps_proto/deps.proto:24:9:
symbol "blaze_deps.SourceLocation" already defined at bazel-project-planton/external/bazel_tools/src/main/protobuf/deps.proto:24:9
```

## What Has Been Tried

1. Adding `ignore` field in `lint` section of `buf.yaml` - Resulted in "field not found" error for v2
2. Creating `.bufignore` file with directory patterns - Still scans external dependencies (patterns ignored)
3. Bazel's `.bazelproject` already excludes these directories but doesn't affect Buf

## Research Objectives

**Primary:** Find authoritative documentation/solutions for:
1. How Buf v2 handles directory exclusion (`.bufignore` syntax, glob patterns, symlink handling)
2. Best practices for Buf + Bazel integration where proto files live in workspace root
3. Whether Buf follows symlinks by default and how to prevent it
4. Alternative approaches: buf.work.yaml, module-specific configuration, or running buf in subdirectories

**Secondary:**
1. Known issues with Buf scanning Bazel external directories
2. Community solutions from projects using both Buf and Bazel with colocated proto files
3. Whether buf lint supports `--path` or similar flags to restrict scanning scope

**Desired Outcome:**
A working configuration that makes `buf lint` process only `provider/` and `shared/` directories while completely ignoring Bazel's generated/symlinked directories and external dependencies.

