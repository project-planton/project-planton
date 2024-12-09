Here are a few examples for the `GcpCloudCdn` API resource, showcasing different configurations that demonstrate basic setups, CDN caching, and security configurations. Since this resource does not have a highly complex spec like other resources, I'll include a few simple configurations and note when details might not make sense yet.

---

### Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudCdn
metadata:
  name: basic-cloud-cdn
spec:
  gcpCredentialId: my-gcp-cred
  gcpProjectId: my-gcp-project
```

---

### Example with Custom Project and Credentials

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudCdn
metadata:
  name: custom-cloud-cdn
spec:
  gcpCredentialId: custom-gcp-credentials
  gcpProjectId: custom-gcp-project-id
```

---

### Example with Empty Spec

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudCdn
metadata:
  name: empty-spec-cdn
spec: {}
```
