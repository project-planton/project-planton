# OpenFGA Authorization Model Examples

## Basic Document Authorization

A simple model for document access control with viewer, editor, and owner roles:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaAuthorizationModel
metadata:
  name: document-authz-v1
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  modelJson: |
    {
      "schema_version": "1.1",
      "type_definitions": [
        {
          "type": "user",
          "relations": {}
        },
        {
          "type": "document",
          "relations": {
            "viewer": {"this": {}},
            "editor": {"this": {}},
            "owner": {"this": {}}
          },
          "metadata": {
            "relations": {
              "viewer": {"directly_related_user_types": [{"type": "user"}]},
              "editor": {"directly_related_user_types": [{"type": "user"}]},
              "owner": {"directly_related_user_types": [{"type": "user"}]}
            }
          }
        }
      ]
    }
```

## Google Drive-like Model

A more complex model with folders, documents, and inheritance:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaAuthorizationModel
metadata:
  name: drive-authz-v1
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  modelJson: |
    {
      "schema_version": "1.1",
      "type_definitions": [
        {
          "type": "user",
          "relations": {}
        },
        {
          "type": "folder",
          "relations": {
            "viewer": {
              "union": {
                "child": [
                  {"this": {}},
                  {"computedUserset": {"relation": "viewer", "object": "", "relation": "parent"}}
                ]
              }
            },
            "editor": {"this": {}},
            "owner": {"this": {}},
            "parent": {"this": {}}
          },
          "metadata": {
            "relations": {
              "viewer": {"directly_related_user_types": [{"type": "user"}]},
              "editor": {"directly_related_user_types": [{"type": "user"}]},
              "owner": {"directly_related_user_types": [{"type": "user"}]},
              "parent": {"directly_related_user_types": [{"type": "folder"}]}
            }
          }
        },
        {
          "type": "document",
          "relations": {
            "viewer": {
              "union": {
                "child": [
                  {"this": {}},
                  {"computedUserset": {"relation": "viewer", "object": "", "relation": "parent"}}
                ]
              }
            },
            "editor": {"this": {}},
            "owner": {"this": {}},
            "parent": {"this": {}}
          },
          "metadata": {
            "relations": {
              "viewer": {"directly_related_user_types": [{"type": "user"}]},
              "editor": {"directly_related_user_types": [{"type": "user"}]},
              "owner": {"directly_related_user_types": [{"type": "user"}]},
              "parent": {"directly_related_user_types": [{"type": "folder"}]}
            }
          }
        }
      ]
    }
```

## Multi-Tenant SaaS Model

A model for multi-tenant applications with organizations and teams:

```yaml
apiVersion: open-fga.project-planton.org/v1
kind: OpenFgaAuthorizationModel
metadata:
  name: saas-authz-v1
  org: my-org
  env: production
spec:
  storeId: "01HXYZ..."
  modelJson: |
    {
      "schema_version": "1.1",
      "type_definitions": [
        {
          "type": "user",
          "relations": {}
        },
        {
          "type": "organization",
          "relations": {
            "admin": {"this": {}},
            "member": {"this": {}}
          },
          "metadata": {
            "relations": {
              "admin": {"directly_related_user_types": [{"type": "user"}]},
              "member": {"directly_related_user_types": [{"type": "user"}]}
            }
          }
        },
        {
          "type": "team",
          "relations": {
            "member": {"this": {}},
            "organization": {"this": {}}
          },
          "metadata": {
            "relations": {
              "member": {"directly_related_user_types": [{"type": "user"}]},
              "organization": {"directly_related_user_types": [{"type": "organization"}]}
            }
          }
        },
        {
          "type": "project",
          "relations": {
            "viewer": {"this": {}},
            "editor": {"this": {}},
            "admin": {"this": {}},
            "team": {"this": {}},
            "organization": {"this": {}}
          },
          "metadata": {
            "relations": {
              "viewer": {"directly_related_user_types": [{"type": "user"}, {"type": "team#member"}]},
              "editor": {"directly_related_user_types": [{"type": "user"}, {"type": "team#member"}]},
              "admin": {"directly_related_user_types": [{"type": "user"}]},
              "team": {"directly_related_user_types": [{"type": "team"}]},
              "organization": {"directly_related_user_types": [{"type": "organization"}]}
            }
          }
        }
      ]
    }
```

## Deployment

All examples require Terraform/Tofu as the provisioner:

```bash
# Create OpenFGA credentials
cat > openfga-creds.yaml << EOF
apiUrl: https://api.fga.example.com
apiToken: your-api-token
EOF

# Deploy the authorization model
project-planton apply --manifest model.yaml \
  --openfga-provider-config openfga-creds.yaml \
  --provisioner tofu
```

## Versioning Strategy

Since authorization models are immutable, use metadata names that include versions:

```yaml
metadata:
  name: document-authz-v2  # Increment version for new model
```

This makes it easy to:
- Track model evolution over time
- Roll back to previous versions if needed
- A/B test different authorization models
