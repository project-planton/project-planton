# Deployment Component Lifecycle Management

## Overview

This directory contains the complete rule system for managing deployment components in Project Planton. These rules handle the entire lifecycle from creation to deletion, ensuring components match the **ideal state** defined in `architecture/deployment-component.md`.

## The Four Lifecycle Operations

Project Planton provides four atomic operations for deployment components:

| Operation | Purpose | When to Use |
|-----------|---------|-------------|
| **🔨 Forge** | Create new components | Component doesn't exist |
| **🔍 Audit** | Assess completeness | Check status, find gaps |
| **🔄 Update** | Enhance existing components | Fill gaps, add features, fix issues |
| **🗑️ Delete** | Remove components | Obsolete, deprecated, consolidating |

**Key Principle:** Each operation is atomic, well-documented, and follows the ideal state standard.

---

## Quick Decision Tree

```
Need to work with a deployment component?
│
├─ Does the component exist?
│  │
│  ├─ NO → Use @forge-project-planton-component
│  │        Creates complete, production-ready component
│  │        Expected result: 95-100% complete
│  │
│  └─ YES → What do you need to do?
│     │
│     ├─ Just checking status?
│     │  └─ Use @audit-project-planton-component
│     │     Shows completion %, identifies gaps
│     │     Generates timestamped report
│     │
│     ├─ Need to improve/fix?
│     │  └─ Use @update-project-planton-component
│     │     Fills gaps, adds features, fixes issues
│     │     6 scenarios: fill-gaps, proto-changed, etc.
│     │
│     └─ Want to remove?
│        └─ Use @delete-project-planton-component
│           Safe deletion with backup and confirmation
│           Checks for references, creates backup
```

---

## 1. Forge: Create New Components

### Purpose
Bootstrap complete, production-ready deployment components from scratch.

### When to Use
- Creating support for a new cloud provider resource
- Adding a new SaaS platform integration
- Creating a new Kubernetes workload or addon

### What It Creates

✅ **Proto API Definitions** - All 4 proto files with validations and tests
✅ **IaC Modules** - Both Pulumi and Terraform with feature parity
✅ **Documentation** - User-facing, research, and technical docs
✅ **Supporting Files** - Test manifests, debug scripts, build configs
✅ **Registry Entry** - Registered in cloud_resource_kind.proto
✅ **Validation** - Build and test validation passed

**Result:** 95-100% completion score

### Usage

```bash
@forge-project-planton-component <ComponentName> --provider <provider>
```

**Examples:**
```bash
@forge-project-planton-component MongodbAtlas --provider atlas
@forge-project-planton-component GcpStorageBucket --provider gcp
@forge-project-planton-component PostgresKubernetes --provider kubernetes --category workload
```

### Learn More
- **README:** [`forge/README.md`](forge/README.md)
- **Rule:** [`forge/forge-project-planton-component.mdc`](forge/forge-project-planton-component.mdc)
- **Flow Rules:** [`forge/flow/`](forge/flow/)

---

## 2. Audit: Assess Component Completeness

### Purpose
Evaluate components against the ideal state and generate actionable completion reports.

### When to Use
- After forge (verify component was created completely)
- After update (confirm improvements)
- Before update (identify what needs fixing)
- Regular quality checks
- Pre-commit validation

### What It Checks

Audits 9 categories:
1. Cloud Resource Registry
2. Folder Structure
3. Protobuf API Definitions
4. IaC Modules - Pulumi
5. IaC Modules - Terraform
6. Documentation - Research
7. Documentation - User-Facing
8. Supporting Files
9. Nice to Have Items

**Scoring:** Weighted (Critical 40%, Important 40%, Nice-to-Have 20%)

### Usage

```bash
@audit-project-planton-component <ComponentName>
```

**Examples:**
```bash
@audit-project-planton-component MongodbAtlas
@audit-project-planton-component GcpCertManagerCert
```

### Report Output

- **Overall completion percentage** (0-100%)
- **Category-by-category breakdown**
- **Quick wins** (easy improvements)
- **Critical gaps** (blocking issues)
- **Prioritized recommendations**
- **Comparison to complete components**

Reports saved to: `<component>/v1/docs/audit/<timestamp>.md`

### Learn More
- **README:** [`audit/README.md`](audit/README.md)
- **Rule:** [`audit/audit-project-planton-component.mdc`](audit/audit-project-planton-component.mdc)

---

## 3. Update: Enhance Existing Components

### Purpose
Improve existing components by filling gaps, adding features, refreshing docs, or fixing issues.

### When to Use
- Filling gaps identified by audit
- Adding new fields to proto schema
- Refreshing outdated documentation
- Modifying IaC deployment logic
- Fixing specific issues

### Update Scenarios

| Scenario | Use When | Command |
|----------|----------|---------|
| **Fill Gaps** | Audit shows <100% | `--scenario fill-gaps` |
| **Proto Changed** | Modified spec.proto | `--scenario proto-changed` |
| **Refresh Docs** | Docs outdated | `--scenario refresh-docs` |
| **Update IaC** | Change deployment | `--scenario update-iac` |
| **Fix Issue** | Specific problem | `--explain "..."` |
| **Auto** | Not sure | Let AI decide |

### Usage

```bash
@update-project-planton-component <ComponentName> [--scenario <type>] [--explain "<description>"]
```

**Examples:**
```bash
# Fill gaps from audit
@update-project-planton-component MongodbAtlas --scenario fill-gaps

# Propagate proto changes
@update-project-planton-component GcpCertManagerCert --scenario proto-changed

# Refresh documentation
@update-project-planton-component AwsRdsInstance --scenario refresh-docs

# Update IaC implementation
@update-project-planton-component GcpGkeCluster --scenario update-iac --explain "add multi-region support"

# Fix specific issue
@update-project-planton-component PostgresKubernetes --explain "examples use deprecated field names"
```

### Safety Features
- ✅ Dry-run mode (`--dry-run`)
- ✅ Backup creation (`--backup`)
- ✅ Validation checkpoints
- ✅ Automatic retry (up to 3 times)
- ✅ Conflict detection

### Learn More
- **README:** [`update/README.md`](update/README.md)
- **Rule:** [`update/update-project-planton-component.mdc`](update/update-project-planton-component.mdc)

---

## 4. Delete: Remove Components Safely

### Purpose
Completely remove deployment components with safety features to prevent accidents.

### When to Use
- Component is obsolete or deprecated
- Provider discontinued service
- Consolidating similar components
- Cleaning up test components

### Safety Features

🔍 **Dry-Run Mode** - Preview what would be deleted
💾 **Automatic Backup** - Creates timestamped backup
🔎 **Reference Check** - Warns if component is referenced
✋ **Confirmation Required** - Must type component name
📋 **Detailed Report** - Shows exactly what was deleted

### Usage

```bash
@delete-project-planton-component <ComponentName> [flags]
```

**Recommended Workflow:**
```bash
# Step 1: Preview (dry-run)
@delete-project-planton-component ObsoleteComponent --dry-run

# Step 2: Delete with backup
@delete-project-planton-component ObsoleteComponent --backup

# Step 3: Confirm (type: DELETE ObsoleteComponent)

# Step 4: Verify
make build && make test
```

**Quick Delete (with caution):**
```bash
@delete-project-planton-component TestComponent --force --backup
```

### What Gets Deleted
- ✅ Component folder (all files)
- ✅ Registry entry (cloud_resource_kind.proto enum)
- ✅ Generated proto stubs (regenerated)

### Learn More
- **README:** [`delete/README.md`](delete/README.md)
- **Rule:** [`delete/delete-project-planton-component.mdc`](delete/delete-project-planton-component.mdc)

---

## Common Workflows

### Workflow 1: Create and Validate

```bash
# 1. Create new component
@forge-project-planton-component NewComponent --provider aws

# 2. Verify completeness
@audit-project-planton-component NewComponent
# Expected: 95-100% complete

# 3. If gaps found, fill them
@update-project-planton-component NewComponent --scenario fill-gaps

# 4. Re-audit to verify
@audit-project-planton-component NewComponent
# Expected: 100% complete
```

### Workflow 2: Improve Existing Component

```bash
# 1. Check current state
@audit-project-planton-component ExistingComponent
# Result: 65% complete (missing Terraform, docs)

# 2. Fill identified gaps
@update-project-planton-component ExistingComponent --scenario fill-gaps

# 3. Verify improvement
@audit-project-planton-component ExistingComponent
# Result: 98% complete
```

### Workflow 3: Add Feature to Component

```bash
# 1. Edit spec.proto (add new fields)
vim apis/org/project_planton/provider/gcp/gcpcloudsl/v1/spec.proto

# 2. Propagate changes
@update-project-planton-component GcpCloudSql --scenario proto-changed

# 3. Test changes
# Deploy with iac/hack/manifest.yaml

# 4. Verify no regressions
@audit-project-planton-component GcpCloudSql
# Score should not decrease
```

### Workflow 4: Replace Component

```bash
# 1. Check what needs migration
@audit-project-planton-component OldComponent
# Result: 40% complete, not worth updating

# 2. Create new implementation
@forge-project-planton-component NewComponent --provider aws

# 3. Migrate users (manual step)
# Update references, notify users

# 4. Delete old component
@delete-project-planton-component OldComponent --backup
```

### Workflow 5: Quality Gate (Pre-Commit)

```bash
# Before committing changes to component
@audit-project-planton-component ModifiedComponent

# If score decreased:
# - Investigate what was lost
# - Fix issues
# - Re-audit

# If score maintained or improved:
# - Safe to commit
git add -A
git commit -m "feat: enhance ModifiedComponent"
```

---

## Integration Points

### With Git Workflows

```bash
# Feature branch workflow
git checkout -b feature/new-component

# Create component
@forge-project-planton-component NewComponent --provider gcp

# Validate
@audit-project-planton-component NewComponent

# Commit
git add -A
git commit -m "feat: add NewComponent for GCP"
git push origin feature/new-component
```

### With CI/CD

```yaml
# .github/workflows/component-quality.yml
on: [pull_request]
jobs:
  audit-components:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Audit modified components
        run: |
          # Detect modified components
          # Run audit on each
          # Fail if score < 80%
```

### With Makefiles

```makefile
# Makefile
.PHONY: audit-all
audit-all:
	@for component in $(shell find apis/org/project_planton/provider -name "v1" -type d); do \
		@audit-project-planton-component $$(basename $$(dirname $$component)); \
	done
```

---

## Reference Documents

### Core Documentation
- **Ideal State Definition:** [`../../architecture/deployment-component.md`](../../architecture/deployment-component.md)
  - Defines what "complete" means
  - Provides checklist for all requirements
  - Explains 80/20 principle

### Rule Documentation
- **Forge:** [`forge/README.md`](forge/README.md)
- **Audit:** [`audit/README.md`](audit/README.md)
- **Update:** [`update/README.md`](update/README.md)
- **Delete:** [`delete/README.md`](delete/README.md)

### Component-Specific Documentation
- **Research Document:** `<component>/v1/docs/README.md`
  - **Critical Reference:** Consult this when executing any rule on a component
  - Contains comprehensive research about the component
  - Explains design decisions and 80/20 scoping rationale
  - Documents deployment landscape and best practices
  - Provides context for understanding what the component does and why
  - **Use when:**
    - Updating component (understand current design)
    - Auditing component (assess research quality)
    - Deleting component (understand impact)
    - Making decisions about component behavior

### Flow Rules
- **Forge Flow:** [`forge/flow/`](forge/flow/) - 21 atomic rules for component creation

---

## Best Practices

### For New Components

1. **Always use forge** - Don't create manually
2. **Audit immediately** - Verify forge created everything
3. **Fill any gaps** - Use update if audit shows <95%
4. **Test locally** - Deploy with hack manifest
5. **Document decisions** - Update research docs if needed

### For Existing Components

1. **Audit first** - Understand current state
2. **Prioritize critical gaps** - Fix blockers first
3. **Update systematically** - Use appropriate scenario
4. **Re-audit after changes** - Verify improvements
5. **Track progress** - Keep audit reports for history

### For Quality Assurance

1. **Regular audits** - Weekly or monthly health checks
2. **Trend tracking** - Compare audit scores over time
3. **Pre-commit gates** - Audit before committing
4. **Team standards** - Minimum 80% for production
5. **Documentation** - Keep audit reports in git

### For Deletions

1. **Always use dry-run first** - Preview deletion
2. **Always use --backup** - Create safety net
3. **Check references** - Don't break other components
4. **Notify team** - Communicate deletions
5. **Document why** - Record rationale in commit message

---

## Troubleshooting

### "Component not found"

**Check:**
- Component name spelling (case-sensitive)
- Component registered in cloud_resource_kind.proto
- Folder exists at expected path

**Solution:**
```bash
# List all components
grep "^\s*[A-Z]" apis/org/project_planton/shared/cloudresourcekind/cloud_resource_kind.proto
```

### "Audit shows 0% but component exists"

**Check:**
- Files exist but might be empty
- Folder structure matches conventions
- Minimum file sizes met (proto >500 bytes, etc.)

**Solution:**
```bash
# Check file sizes
find apis/org/project_planton/provider/atlas/mongodbatlas/v1 -type f -exec ls -lh {} \;
```

### "Update fails with build errors"

**Check:**
- Proto syntax is valid
- Import paths are correct
- Generated stubs are current

**Solution:**
```bash
# Regenerate stubs
make protos

# Check build
make build

# Check tests
make test
```

### "Delete warns about references"

**Options:**
1. Fix references first (recommended)
2. Use --force (may break builds)
3. Cancel deletion

**Solution:**
```bash
# Find all references
grep -r "ComponentName" apis/
grep -r "ComponentName" docs/

# Update references
# Then delete
```

---

## Success Metrics

### Component Quality

- **100%** = Perfect, production-ready
- **95-99%** = Excellent, minor polish possible
- **80-94%** = Good, some improvements recommended
- **60-79%** = Fair, significant work needed
- **<60%** = Poor, major work or reconsider

### Team Productivity

- **Time to create:** <30 minutes (forge)
- **Time to audit:** <1 minute
- **Time to update:** 10-60 minutes (depends on scenario)
- **Time to delete:** <2 minutes

### Quality Standards

- **New components:** ≥95% at creation
- **Production components:** ≥80% minimum
- **Active development:** ≥90% target
- **Deprecated components:** Audit before delete

---

## Examples by Provider

### AWS Components

```bash
@forge-project-planton-component AwsRdsInstance --provider aws
@forge-project-planton-component AwsEksCluster --provider aws
@forge-project-planton-component AwsS3Bucket --provider aws
```

### GCP Components

```bash
@forge-project-planton-component GcpCloudSql --provider gcp
@forge-project-planton-component GcpGkeCluster --provider gcp
@forge-project-planton-component GcpStorageBucket --provider gcp
```

### Kubernetes Components

```bash
@forge-project-planton-component PostgresKubernetes --provider kubernetes --category workload
@forge-project-planton-component RedisKubernetes --provider kubernetes --category workload
@forge-project-planton-component CertManagerKubernetes --provider kubernetes --category addon
```

### SaaS Platform Components

```bash
@forge-project-planton-component MongodbAtlas --provider atlas
@forge-project-planton-component ConfluentKafka --provider confluent
@forge-project-planton-component SnowflakeDatabase --provider snowflake
```

---

## Getting Help

### Documentation
- Read the README for specific operation
- Check ideal state document for requirements
- Review flow rules for implementation details

### Examples
- See complete components (e.g., GcpCertManagerCert)
- Run audit on gold-standard components
- Compare incomplete vs complete components

### Support
- GitHub Discussions for questions
- GitHub Issues for bugs
- Team chat for quick questions

---

## Contributing

When adding new rules or improving existing ones:

1. **Follow existing patterns** - Consistency matters
2. **Update documentation** - Keep READMEs current
3. **Test thoroughly** - Validate on multiple components
4. **Add examples** - Show real usage
5. **Update ideal state** - If requirements change

---

**Ready to start?** Choose the operation you need and follow its README for detailed instructions!

| Operation | Command Template |
|-----------|-----------------|
| **Forge** | `@forge-project-planton-component <Name> --provider <provider>` |
| **Audit** | `@audit-project-planton-component <Name>` |
| **Update** | `@update-project-planton-component <Name> [--scenario <type>]` |
| **Delete** | `@delete-project-planton-component <Name> --backup` |

