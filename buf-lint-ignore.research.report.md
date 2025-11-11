
# **Analysis and Configuration for Buf Lint in a Bazel-Managed Monorepo**

## **I. Executive Summary: The Three-Tiered Solution for Buf and Bazel**

This report details a comprehensive strategy for resolving a critical file discovery conflict between the Buf CLI (v1.55.1) and a Bazel 8 monorepo. The core problem manifests when running buf lint from the workspace root. Buf's default behavior is to traverse all directories, including Bazel's "convenience symlinks" (e.g., bazel-project-planton). This traversal incorrectly leads buf into Bazel's external/ dependency cache, which contains a mix of duplicate, test, and intentionally broken Protobuf files from build-time dependencies like gazelle and bazel\_tools. This results in a "symbol... already defined" compilation error before linting can even begin.

The conflict stems from two incompatible default behaviors: Buf's default setting to follow symlinks 1 and Bazel's architectural use of non-source, non-hermetic symlinks at the workspace root.2

This report presents three distinct,-phased solutions to resolve this conflict definitively:

1. **Immediate Fix (Tactical):** The use of the buf lint \--disable-symlinks flag. This command-line argument immediately instructs buf not to traverse the problematic symlinks, resolving the error.1  
2. **Preferred Fix (Declarative):** A re-configuration of the root buf.yaml to utilize the v2 modules key. This approach declaratively scopes Buf's file discovery to *only* the specified Protobuf source directories (provider/ and shared/), making buf entirely ignore the Bazel symlinks.4  
3. **Idiomatic Fix (Strategic):** A full integration of buf into the Bazel build graph using rules\_buf. This method replaces filesystem scanning with hermetic, cached, and incremental buf\_lint\_test targets, which is the "best practice" for a Bazel-centric monorepo.5

The final recommendation is to implement Solution 1 immediately to unblock development, followed by Solution 2 to establish a correct declarative configuration for standalone buf use. Finally, Solution 3 should be prioritized as the long-term, strategic integration that aligns with the monorepo's existing Bazel-centric workflow.

## **II. Deconstructing the Root Cause: Buf's Discovery vs. Bazel's Architecture**

The error observed—symbol "blaze\_deps.SourceLocation" already defined—is not a linting violation but a fatal compilation error. The buf CLI first builds a complete "image" of all discovered .proto files, functionally equivalent to a google.protobuf.FileDescriptorSet, before it can apply any linting rules.6 The error indicates that during this image-building phase, buf has found two different files (one in .../external/gazelle++go\_deps+... and one in .../external/bazel\_tools/...) that both define the same blaze\_deps.SourceLocation symbol.

This is a classic duplicate symbol problem, common in complex build environments where multiple dependencies provide the same definitions.7 The error paths prove conclusively that buf is traversing Bazel's external dependency cache, which is not part of the project's Protobuf source code.

### **Why Previous Attempts Failed: A v2 vs. v1 Mindset**

The attempts to resolve this issue were unsuccessful because they were based on configuration paradigms that do not apply to this specific problem:

1. **Failure 1: lint.ignore:** The attempt to add an ignore field to the lint section of a v2 buf.yaml failed because that key does not exist in the v2 lint configuration schema.4 While a lint.ignore key existed in v1beta1 and v1 configuration, it was for *ignoring specific lint rule violations*, not for excluding files from discovery.10 Since the failure is a compilation error, no lint rules are even being evaluated.  
2. **Failure 2: .bufignore:** The use of a .bufignore file failed for a more subtle reason. Bazel's "convenience symlinks" often point to non-deterministic, user-specific paths within /private/var/tmp/.11 It is futile to create an ignore file for a path (e.g., bazel-project-planton) that is itself a symlink to an unpredictable, machine-local directory.

### **The True Culprit: Symlink Traversal**

The root cause is a direct and known conflict between buf and bazel. The buf CLI *intentionally* added support for following symlinks to accommodate certain build system layouts.11 The official documentation for buf lint, buf build, and buf breaking confirms this default behavior: "By default, symlinks are followed in this CLI...".1

This default behavior is fundamentally incompatible with Bazel's creation of "convenience symlinks" in the workspace root. This exact conflict has been documented by the community: a user on GitHub reported, "Is there a way to make buf ignore the symlinks? It is picking up convenience symlinks created by bazel." This report led directly to the creation of a new flag to resolve the issue.11 The project's current problem is not a unique misconfiguration but the canonical conflict buf's maintainers are aware of.

## **III. Solution 1 (Immediate Fix): The \--disable-symlinks Flag**

The most direct and immediate solution is to use the CLI flag created specifically for this problem.

### **Authoritative Documentation and Implementation**

The buf lint command (version 1.55.1) accepts a boolean flag: \--disable-symlinks.1 The flag's official description is: "Do not follow symlinks when reading sources or configuration from the local filesystem".1

As established, this flag was introduced to address conflicts with tools like Bazel that create symlinks in the workspace.11

To resolve the error, the buf lint command should be invoked from the project root as follows:

Bash

buf lint \--disable-symlinks

This command instructs buf to treat bazel-project-planton and other symlinks as simple files (which it will ignore) rather than as directories to be traversed. This immediately stops buf from descending into the external/ cache, resolving the duplicate symbol error.

### **Analysis: Pros and Cons**

* **Pros:**  
  * **Trivial Implementation:** It requires no configuration changes and can be communicated to the team immediately.  
  * **Immediate Resolution:** It unblocks local development and CI/CD pipelines that rely on buf lint.  
* **Cons:**  
  * **Procedural, Not Declarative:** This fix is a command-line argument. It must be appended to *every* buf command (lint, build, breaking) that runs from the root. If a developer or script forgets the flag, the build will fail. Configuration should ideally be declarative (in a file) rather than procedural (in a command).  
  * **Inefficient:** This flag does not solve the underlying inefficiency. buf is still being invoked with . as its default input 1, meaning it will still scan the entire project root. It will read cmd/, internal/, pkg/, build/, site/, etc., only to find no .proto files. For a large monorepo, this is a needless performance penalty. This flag tells buf *what not* to traverse, but the better solution is to tell it *only* where to look.

## **IV. Solution 2 (Preferred Fix): Declarative Scoping with v2 modules Configuration**

The preferred and more robust solution is to move from buf's v1-compatibility mode to a modern v2 buf.yaml configuration. This changes the file discovery model from "exclusion" (trying to ignore bad directories) to "inclusion" (declaratively specifying *only* the correct directories).

### **The v2 Paradigm Shift: From Exclusion to Inclusion**

The v2 buf.yaml configuration was introduced specifically to provide first-class support for monorepos.13 Its core concept is the *workspace*, which is defined by a buf.yaml at the root and contains one or more *modules*.4

When a buf.yaml file *lacks* a modules key, buf (version 1.55.1) defaults to its v1 behavior: it assumes the input is the current directory (.).1 This is the source of the problem.

By adding the modules key, buf is instructed to *only* look for .proto files within the directories specified in each module's path entry. All other directories at the root—including bazel-project-planton, cmd/, and internal/—are completely and efficiently ignored.

The project's structure, with protos in provider/ (519+ files) and shared/ (16 files), is a perfect use case for a multi-module workspace. Protos in provider/ can import protos from shared/ (e.g., import "shared/common.proto";) without any special deps configuration, as they are part of the same workspace and buf resolves them locally.14

### **Implementation (Declarative)**

The buf.yaml file at the project root should be replaced with the following v2 configuration. This file declaratively defines the two source directories as modules, lists BSR dependencies, and configures linting for the entire workspace.

YAML

\# buf.yaml  
version: v2

\# Define the modules in this workspace.  
\# By adding this 'modules' key, Buf will \*only\* look  
\# in these directories and will ignore everything else  
\# at the root (including Bazel's symlinks).  
modules:  
  \- path: provider  
    \# This 'name' is the module's BSR identity.\[4, 16\]  
    \# It maps the 'provider/' directory to your primary BSR repository.  
    name: buf.build/project-planton/apis  
  \- path: shared  
    \# It is best practice for 'shared' to be its own module.  
    \# This allows other projects to depend \*only\* on shared types.  
    name: buf.build/project-planton/apis-shared

\# Define all BSR dependencies for the \*entire workspace\*.  
\# These are shared by all modules defined above.  
deps:  
  \- buf.build/bufbuild/protovalidate  
  \# Your custom plugin is also a BSR dependency.  
  \- buf.build/project-planton/optional-linter:v0.1.0

\# Workspace-level lint configuration  
lint:  
  use:  
    \- STANDARD  
  \# Add your custom linter rule category here.  
  \# This assumes your plugin 'optional-linter'  
  \# defines a rule or category named 'PROJECT\_OPTIONAL'.  
  \- buf.build/project-planton/optional-linter/PROJECT\_OPTIONAL

\# Workspace-level breaking change detection  
breaking:  
  use:  
    \- FILE

With this configuration, running buf lint (with no flags) will now work correctly, efficiently, and idempotently.

### **Analysis: Pros and Cons**

* **Pros:**  
  * **Declarative and Idempotent:** The build scope is defined in a version-controlled file, not on the command line.  
  * **Solves the Root Problem:** buf lint no longer scans the root; it only scans provider/ and shared/. The Bazel symlinks are never seen.  
  * **Efficient:** File discovery is now limited to *only* the 535+ .proto source files, which will be significantly faster than scanning the entire monorepo.  
* **Cons:**  
  * **Requires v2 Adoption:** This requires adopting the v2 modules structure. This may necessitate creating a new BSR repository for the shared module if one does not already exist, which is a best practice.16  
  * **Still Standalone:** This configuration perfects the use of buf as a *standalone* tool. It does not integrate linting into the Bazel build graph, so it remains a separate, non-hermetic, non-cached operation.

## **V. Solution 3 (Idiomatic Fix): Full Bazel Integration with rules\_buf**

The most strategic and idiomatic solution for a sophisticated Bazel monorepo is to make API linting a part of the build graph itself. This approach aligns buf with the hermetic, cached, and incremental principles of Bazel.

### **The Strategic Goal: A Hermetic, Cached, Incremental Build**

This solution fundamentally changes the execution model. Instead of buf lint *pulling* files by scanning the filesystem, Bazel will *push* hermetic proto\_library targets *to* the linting rule.17 This design completely bypasses the filesystem discovery problem.

The rules\_buf repository provides Bazel rules for buf.5 The key rule is buf\_lint\_test, which runs buf lint as a Bazel test target.5 The existing Gazelle setup can be configured to generate a buf\_lint\_test target for every proto\_library target.5

The result is that bazel test //... will run buf lint hermetically, in parallel, and cache the results. A change to a single .proto file will only trigger the *one* buf\_lint\_test target associated with it, not a full scan of all 500+ files.

### **Implementation (Strategic)**

This is a multi-step integration:

**Step 1: Configure MODULE.bazel for rules\_buf (bzlmod)**

The MODULE.bazel file must be updated to add rules\_buf and its dependencies, and to configure the buf extension to download the CLI and BSR dependencies.

Code snippet

\# MODULE.bazel

\#... other bazel\_dep calls (rules\_go, gazelle, etc.)...  
bazel\_dep(name \= "rules\_proto", version \= "...")  
bazel\_dep(name \= "protobuf", repo\_name \= "com\_google\_protobuf", version \= "...")

\# Add rules\_buf  
bazel\_dep(name \= "rules\_buf", version \= "0.4.0") \# 

\# \--- Buf Extension \---  
\# Load and configure the buf extension   
buf \= use\_extension("@rules\_buf//buf:extensions.bzl", "buf")

\# 1\. Configure the Buf CLI version to be used by the toolchain  
buf.toolchains(version \= "v1.50.0") \# 

\# 2\. Add BSR dependencies (replaces 'deps:' in buf.yaml for Bazel builds)  
\# These are downloaded and made available as Bazel targets (e.g., @buf\_deps//...)  
buf.dependency(  
    module \= "buf.build/bufbuild/protovalidate",  
    \# Use an immutable commit hash for hermeticity   
    revision \= "eac44469a7af47e7839a7f1f3d7ac004", \# Example, get latest  
)  
buf.dependency(  
    module \= "buf.build/project-planton/optional-linter",  
    revision \= "...", \# Use the immutable revision for v0.1.0  
)

\# 3\. Expose the BSR dependencies as a Bazel repository  
use\_repo(buf, "buf\_deps") \# 

\# 4\. Expose the Buf CLI toolchain itself  
use\_repo(buf, "rules\_buf\_toolchains") \# 

**Step 2: Configure Gazelle to Generate buf\_lint\_test Rules**

The Gazelle configuration in the root BUILD.bazel file must be updated to enable the rules\_buf extension.

Code snippet

\# //:BUILD.bazel  
load("@gazelle//:def.bzl", "gazelle")

gazelle(  
    name \= "gazelle",  
    \#... other args...  
    languages \= \[  
        "@io\_bazel\_rules\_go//go:go\_language",  
        "@rules\_proto//proto:proto\_language",  
        \# Add the 'buf' language to enable the extension \[20\]  
        "@rules\_buf//gazelle/buf:buf\_language",  
    \],  
)

The rules\_buf Gazelle extension is designed to work *with* proto\_library targets.5 The project's current "proto generation disabled" state for Gazelle may mean it is not generating proto\_library targets. This is a prerequisite; buf\_lint\_test *attaches* to an existing proto\_library target.

**Step 3: Run Gazelle and Test**

After updating the configuration, Gazelle must be run:

Bash

bazel run //:gazelle

This command will update all BUILD.bazel files in provider/ and shared/. For every proto\_library target, it will add a corresponding buf\_lint\_test target.5

This can be verified by running:  
bazel query 'kind(buf\_lint\_test, //...)'  
The linting workflow is now integrated into the standard Bazel test command:

Bash

\# Run all tests, including buf lint checks  
bazel test //...

\# Or, run only the buf lint checks  
bazel test $(bazel query 'kind(buf\_lint\_test, //...)')

This workflow will use the buf.yaml file from Solution 2 for its linting rule configuration. The buf\_lint\_test rule automatically discovers and applies the buf.yaml file from the workspace root.5 Therefore, Solution 2 is a prerequisite for Solution 3\.

### **Analysis: Pros and Cons**

* **Pros:**  
  * **Fully Idiomatic:** This is the "best practice" for using buf within a Bazel-centric monorepo.  
  * **Hermetic:** Linting inputs are defined *only* by the build graph, not by filesystem scanning.  
  * **Incremental and Cached:** bazel test will only re-lint the proto\_library targets that have changed, providing a massive performance increase for CI.  
  * **Consolidated:** MODULE.bazel manages dependencies, and buf.yaml manages lint configuration.  
* **Cons:**  
  * **Medium Effort:** This requires a non-trivial, one-time setup of MODULE.bazel and Gazelle.  
  * **Dependency on proto\_library:** This solution is contingent on Gazelle generating proto\_library targets.

## **VI. Addressing Secondary Objective: The \--path Flag**

A secondary research objective was to determine if buf lint supports a \--path flag to restrict the scanning scope.

### **Answer and Analysis**

Yes, the buf lint command supports a \--path flag. This flag limits buf to only processing the specified files or directories.1 The command would be:

Bash

buf lint \--path provider/ \--path shared/

This command *would* have successfully solved the immediate problem. It explicitly tells buf to only lint the provider/ and shared/ directories, thus ignoring the root-level symlinks.

However, this flag is considered an "advanced feature" and is "mostly intended to be used by other systems like editors".9 The documentation itself recommends letting buf discover files.21

This flag is the *procedural* equivalent of Solution 2\. Solution 2 (the v2 modules config) is superior because it is *declarative*. It bakes the provider/ and shared/ paths into the buf.yaml file, so the command remains a simple, flag-less buf lint. By implementing Solution 2, the workspace is *correctly configured* to be discoverable, aligning with buf's recommended practice.

## **VII. Comparative Analysis and Final Recommendation**

The four viable solutions (including the \--path flag) present clear trade-offs in implementation effort, performance, and adherence to build-system principles.

**Table: Comparison of Buf \+ Bazel Integration Strategies**

| Strategy | Implementation Effort | Solves Error? | Bazel-Idiomatic? | CI Performance | Configuration Style |
| :---- | :---- | :---- | :---- | :---- | :---- |
| **1\. \--disable-symlinks Flag** | Trivial (Flag) | **Yes** | No | **Slow** (Full Scan, No Traversal) | Procedural (CLI Flag) |
| **2\. \--path Flag** | Trivial (Flag) | **Yes** | No | **Fast** (Scoped Scan) | Procedural (CLI Flag) |
| **3\. v2 modules Config** | Low (Config Edit) | **Yes** | Partially | **Fast** (Scoped Scan) | **Declarative (File)** |
| **4\. rules\_buf Integration** | Medium (Bazel Setup) | **Yes** | **Yes (Best Practice)** | **Fastest** (Incremental & Cached) | **Declarative (Build Graph)** |

### **Final Phased Recommendation**

This analysis leads to a clear, three-phase recommendation to resolve the issue, establish a robust configuration, and achieve a best-in-class strategic integration.

1. **Phase 1: Immediate Mitigation (Today)**  
   * **Action:** Adopt the buf lint \--disable-symlinks flag for all local and CI invocations of buf.  
   * **Rationale:** This unblocks all developers and CI pipelines immediately with zero configuration change.1  
2. **Phase 2: Declarative Standalone (This Week)**  
   * **Action:** Refactor the root buf.yaml to implement the v2 modules configuration as detailed in Solution 2\.  
   * **Rationale:** This is the *correct* declarative fix for standalone buf usage. It makes buf lint (with no flags) work correctly, efficiently, and idempotently, eliminating the need for the \--disable-symlinks flag.4  
3. **Phase 3: Strategic Integration (This Quarter)**  
   * **Action:** Execute the rules\_buf integration detailed in Solution 3\.  
   * **Rationale:** This is the long-term, "best-in-class" solution. It fully aligns buf with the existing Bazel-centric workflow, unlocking hermeticity, incrementalism, and caching. This will provide a significant performance improvement for CI/CD and a more robust, integrated developer experience.5

#### **Works cited**

1. buf lint \- Buf Docs \- Buf.Build, accessed on November 11, 2025, [https://buf.build/docs/reference/cli/buf/lint/](https://buf.build/docs/reference/cli/buf/lint/)  
2. Commands and Options | Bazel, accessed on November 11, 2025, [https://bazel.build/docs/user-manual](https://bazel.build/docs/user-manual)  
3. buf-lint(1) \- Arch manual pages, accessed on November 11, 2025, [https://man.archlinux.org/man/buf-lint.1.en](https://man.archlinux.org/man/buf-lint.1.en)  
4. buf.yaml \- Buf Docs \- Buf.Build, accessed on November 11, 2025, [https://buf.build/docs/configuration/v2/buf-yaml/](https://buf.build/docs/configuration/v2/buf-yaml/)  
5. Bazel \- Buf Docs \- Buf.Build, accessed on November 11, 2025, [https://buf.build/docs/cli/build-systems/bazel/](https://buf.build/docs/cli/build-systems/bazel/)  
6. Buf Docs \- buf build, accessed on November 11, 2025, [https://buf.build/docs/reference/cli/buf/build/](https://buf.build/docs/reference/cli/buf/build/)  
7. Duplicate symbol issue with bazel and blacklisted\_protos \#7046 \- GitHub, accessed on November 11, 2025, [https://github.com/protocolbuffers/protobuf/issues/7046](https://github.com/protocolbuffers/protobuf/issues/7046)  
8. How to avoid duplicating work across bazel configurations? \- Stack Overflow, accessed on November 11, 2025, [https://stackoverflow.com/questions/66540087/how-to-avoid-duplicating-work-across-bazel-configurations](https://stackoverflow.com/questions/66540087/how-to-avoid-duplicating-work-across-bazel-configurations)  
9. Linting – Overview \- Buf Docs, accessed on November 11, 2025, [https://docs.bufbuild.ru/lint/](https://docs.bufbuild.ru/lint/)  
10. buf.yaml \- Buf Docs, accessed on November 11, 2025, [https://buf.build/docs/configuration/v1/buf-yaml/](https://buf.build/docs/configuration/v1/buf-yaml/)  
11. Evaluate adding symlink support · Issue \#104 · bufbuild/buf \- GitHub, accessed on November 11, 2025, [https://github.com/bufbuild/buf/issues/104](https://github.com/bufbuild/buf/issues/104)  
12. buf breaking \- Buf Docs, accessed on November 11, 2025, [https://buf.build/docs/reference/cli/buf/breaking/](https://buf.build/docs/reference/cli/buf/breaking/)  
13. Introducing the next generation of the Buf CLI: still v1 and backwards-compatible, accessed on November 11, 2025, [https://buf.build/blog/buf-cli-next-generation](https://buf.build/blog/buf-cli-next-generation)  
14. Modules and workspaces \- Buf Docs \- Buf.Build, accessed on November 11, 2025, [https://buf.build/docs/cli/modules-workspaces/](https://buf.build/docs/cli/modules-workspaces/)  
15. buf lint, accessed on November 11, 2025, [https://docs.bufbuild.ru/reference/cli/buf/lint/](https://docs.bufbuild.ru/reference/cli/buf/lint/)  
16. Buf Schema Registry repositories \- Buf Docs \- Buf.Build, accessed on November 11, 2025, [https://buf.build/docs/bsr/repositories/](https://buf.build/docs/bsr/repositories/)  
17. Buf \- Protobuf and gRPC rules for Bazel \- rules\_proto\_grpc, accessed on November 11, 2025, [https://rules-proto-grpc.com/en/5.2.1/lang/buf.html](https://rules-proto-grpc.com/en/5.2.1/lang/buf.html)  
18. Bazel rules · Issue \#681 · bufbuild/buf \- GitHub, accessed on November 11, 2025, [https://github.com/bufbuild/buf/issues/681](https://github.com/bufbuild/buf/issues/681)  
19. Bazel rules \- Buf, accessed on November 11, 2025, [https://buf.build/blog/bazel-rules](https://buf.build/blog/bazel-rules)  
20. Linting and code standards \- Buf Docs, accessed on November 11, 2025, [https://buf.build/docs/lint/](https://buf.build/docs/lint/)
