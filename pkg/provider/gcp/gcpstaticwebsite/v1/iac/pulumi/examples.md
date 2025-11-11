

### Example 1: Basic Static Website on GCP Storage

This example sets up a simple static website hosted on Google Cloud Storage with public access. It provisions a storage bucket configured to serve static content as a public website.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: basic-static-website
spec:
  projectId: my-gcp-project
  bucket_config:
    bucketName: my-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
```

---

### Example 2: Static Website with Custom Domain and SSL

This example provisions a static website on GCP Storage with a custom domain and enables SSL for secure connections. It assumes that the SSL certificate is managed via GCP.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: custom-domain-static-website
spec:
  projectId: my-custom-domain-project
  bucket_config:
    bucketName: custom-domain-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    custom_domain:
      domain_name: www.mysite.com
      ssl_enabled: true
```

---

### Example 3: Static Website with Private Access

This example sets up a static website on Google Cloud Storage but configures the bucket for private access. This is useful for situations where you want to restrict access to the static content.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: private-static-website
spec:
  projectId: private-gcp-project
  bucket_config:
    bucketName: my-private-static-site
    public_access: false
    index_page: index.html
    error_page: 403.html
    security_settings:
      viewer_permissions: ["user1@example.com", "user2@example.com"]
```

---

### Example 4: Static Website with Versioning Enabled

This example configures the GCP Storage bucket for a static website with versioning enabled, which allows for multiple versions of objects to be stored and accessed. This setup is useful for rollback or archiving purposes.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: versioned-static-website
spec:
  projectId: versioned-site-project
  bucket_config:
    bucketName: versioned-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    versioningEnabled: true
```

---

### Example 5: Static Website with CORS Configuration

This example configures a static website hosted on GCP Storage with Cross-Origin Resource Sharing (CORS) settings. This setup allows the website to serve resources to specified domains.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: cors-enabled-static-website
spec:
  projectId: cors-site-project
  bucket_config:
    bucketName: cors-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    cors_settings:
      allowed_origins:
        - https://www.allowed-domain.com
      allowed_methods:
        - GET
        - POST
      max_age_seconds: 3600
```

---

### Example 6: Static Website with Lifecycle Rules for Object Deletion

This example provisions a GCP Storage bucket for a static website and configures lifecycle rules to automatically delete objects after a specified period.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: lifecycle-static-website
spec:
  projectId: lifecycle-project
  bucket_config:
    bucketName: lifecycle-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    lifecycle_rules:
      - action: Delete
        age: 30
```

---

### Example 7: Multi-Region Static Website for High Availability

This example sets up a static website hosted on a multi-region Google Cloud Storage bucket, providing high availability and redundancy for the website content.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: multi-region-static-website
spec:
  projectId: ha-project
  bucket_config:
    bucketName: multi-region-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    region: multi-region
```

---

### Example 8: Static Website with Logging Enabled

This example configures a static website on GCP Storage with logging enabled to track and monitor access to the website content.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: logging-enabled-static-website
spec:
  projectId: logging-project
  bucket_config:
    bucketName: logging-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    logging:
      log_bucket_name: website-access-logs
      log_object_prefix: access-log
```

---

### Example 9: Static Website with Custom Error Pages

This example provisions a static website on GCP Storage and configures custom error pages to be displayed for specific HTTP errors, such as 404 (Not Found) and 403 (Forbidden).

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpStaticWebsite
metadata:
  name: custom-error-pages-static-website
spec:
  projectId: custom-error-project
  bucket_config:
    bucketName: custom-error-static-site
    public_access: true
    index_page: index.html
    error_page: 404.html
    custom_error_pages:
      404: custom-404.html
      403: custom-403.html
```

---

### Applying the Configurations

To deploy any of these GCP Static Website configurations, create a YAML file with the desired example content and use the following command to apply the configuration:

```shell
planton apply -f <yaml-path>
```
