# Project Planton Changelog

This directory contains detailed changelog entries for Project Planton. Each significant change is documented in its own file with the following naming convention:

```
YYYY-MM-DD-brief-description.md
```

## Changelog Format

Each changelog entry should include:

1. **Title** - Brief description of the change
2. **Date** - Date of the change (YYYY-MM-DD)
3. **Type** - Feature, Fix, Breaking Change, Performance, Security, etc.
4. **PR** - Link to the pull request (if applicable)
5. **Summary** - High-level overview of the change
6. **Motivation** - Why the change was made
7. **What's New** - Detailed description of new functionality
8. **Implementation Details** - Technical details for maintainers
9. **Migration Guide** - Steps for users to adopt the change (if applicable)
10. **Examples** - Code snippets demonstrating the change
11. **Benefits** - Value provided to users

## Recent Changes

- [2025-10-11: Percona PostgreSQL Operator Support](./2025-10-11-percona-postgresql-operator.md) - Added support for deploying and managing the Percona Distribution for PostgreSQL Operator on Kubernetes clusters with automated deployment, high availability, backup, and disaster recovery features
- [2025-10-11: Percona Server MySQL Operator & Improved Error Handling](./2025-10-11-percona-server-mysql-operator-and-improved-error-handling.md) - Added support for Percona Server for MySQL Operator and significantly improved CLI error messaging with beautiful, actionable guidance for unsupported cloud resource kinds
- [2025-10-10: Percona Server MongoDB Operator Support](./2025-10-10-percona-server-mongodb-operator.md) - Added support for deploying and managing the Percona Server for MongoDB Operator on Kubernetes clusters with automated deployment, backup, and high availability features
- [2025-10-10: PostgreSQL Backup and Disaster Recovery](./2025-10-10-postgres-backup-configuration.md) - Added comprehensive backup and disaster recovery capabilities for PostgreSQL databases with operator-level defaults and per-database overrides
- [2025-09-16: Manifest Backend Configuration](./2025-09-16-manifest-backend-configuration.md) - Added support for embedding Pulumi and Terraform/Tofu backend configuration in manifest labels

## Notes

- Each changelog file represents a single logical change or feature
- Changes are ordered chronologically by filename
- These entries will be aggregated for release notes and web documentation
- Keep entries detailed but user-focused
