# GCP Cloud CDN Examples

This document provides comprehensive examples for deploying Google Cloud CDN using Project Planton. These examples demonstrate the 80/20 configuration principle, covering common use cases with simple configurations and advanced scenarios with detailed tuning.

## Table of Contents

1. [Basic Static Website CDN (GCS Backend)](#1-basic-static-website-cdn-gcs-backend)
2. [Static Website with Custom Domain and SSL](#2-static-website-with-custom-domain-and-ssl)
3. [CDN with Custom Cache Configuration](#3-cdn-with-custom-cache-configuration)
4. [CDN with Cache Key Optimization](#4-cdn-with-cache-key-optimization)
5. [CDN with Signed URLs for Private Content](#5-cdn-with-signed-urls-for-private-content)
6. [Cloud Run Backend with CDN](#6-cloud-run-backend-with-cdn)
7. [Compute Engine Backend with Health Checks](#7-compute-engine-backend-with-health-checks)
8. [External Origin (Hybrid/Multi-Cloud)](#8-external-origin-hybridmulti-cloud)
9. [Production CDN with Cloud Armor](#9-production-cdn-with-cloud-armor)
10. [Advanced Configuration with Negative Caching](#10-advanced-configuration-with-negative-caching)

---

## 1. Basic Static Website CDN (GCS Backend)

The simplest use case: serve a static website from Google Cloud Storage with Cloud CDN enabled.

**Use case:** Personal blog, marketing site, documentation portal

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: my-static-site-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    gcs_bucket:
      bucket_name: my-static-website-bucket
      enable_uniform_access: true
  # Optional: Using defaults for cache_mode (CACHE_ALL_STATIC), default_ttl_seconds (3600)
```

**What this creates:**
- BackendBucket pointing to GCS bucket
- Cloud CDN enabled with default cache settings
- Global HTTPS load balancer with auto-generated IP

---

## 2. Static Website with Custom Domain and SSL

Same as above, but with custom domain and Google-managed SSL certificate.

**Use case:** Production website with branded domain

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: production-website-cdn
spec:
  gcp_project_id: my-company-prod
  backend:
    gcs_bucket:
      bucket_name: www-mycompany-com-bucket
      enable_uniform_access: true
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 3600  # 1 hour
  max_ttl_seconds: 86400     # 1 day
  enable_negative_caching: true
  frontend_config:
    custom_domains:
      - www.mycompany.com
      - mycompany.com
    ssl_certificate:
      google_managed:
        domains:
          - www.mycompany.com
          - mycompany.com
    enable_https_redirect: true
```

**After deployment:**
1. Point DNS A records for `www.mycompany.com` and `mycompany.com` to the `global_ip_address` output
2. Google will automatically provision SSL certificates (takes 10-15 minutes)
3. Users accessing `http://mycompany.com` will be redirected to `https://mycompany.com`

---

## 3. CDN with Custom Cache Configuration

Fine-tune caching behavior with custom TTLs and cache mode.

**Use case:** API responses or dynamic content with specific caching requirements

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: api-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    gcs_bucket:
      bucket_name: api-static-responses-bucket
  cache_mode: USE_ORIGIN_HEADERS  # Only cache if origin sends Cache-Control headers
  default_ttl_seconds: 1800        # 30 minutes default
  max_ttl_seconds: 3600            # 1 hour maximum (overrides origin headers)
  client_ttl_seconds: 900          # 15 minutes for browser cache
  enable_negative_caching: true    # Cache 404 errors
```

**When to use `USE_ORIGIN_HEADERS`:**
- Your origin (GCS bucket or backend service) sets proper `Cache-Control` headers on all responses
- You want fine-grained control per-file (different TTLs for different content)
- You're confident all cacheable content has appropriate headers

---

## 4. CDN with Cache Key Optimization

Prevent "cache shattering" by controlling what goes into the cache key.

**Use case:** Website with analytics parameters, user tracking, or session IDs in query strings

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: optimized-cache-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    gcs_bucket:
      bucket_name: webapp-assets-bucket
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 3600
  max_ttl_seconds: 86400
  enable_negative_caching: true
  advanced_config:
    cache_key_policy:
      include_query_string: true
      # Only these query params affect caching; all others are ignored
      query_string_whitelist:
        - version
        - lang
        - page
      include_protocol: true
      include_host: true
    enable_request_coalescing: true
```

**Problem this solves:**

Without cache key optimization:
```
Request 1: /app.js?version=1.0&user_id=123&session=abc → Cache entry 1
Request 2: /app.js?version=1.0&user_id=456&session=xyz → Cache entry 2 (cache MISS!)
```

With query_string_whitelist = ["version"]:
```
Request 1: /app.js?version=1.0&user_id=123&session=abc → Cache entry (version=1.0)
Request 2: /app.js?version=1.0&user_id=456&session=xyz → Cache HIT! (same cache entry)
```

---

## 5. CDN with Signed URLs for Private Content

Serve paid or user-specific content with time-limited, cryptographically signed URLs.

**Use case:** Paid video courses, premium downloads, user-uploaded files

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: private-content-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    gcs_bucket:
      bucket_name: premium-videos-bucket
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 7200  # 2 hours (videos change infrequently)
  max_ttl_seconds: 86400     # 1 day
  enable_negative_caching: true
  advanced_config:
    signed_url_config:
      enabled: true
      keys:
        - key_name: primary-key-2024
          key_value: aGVsbG93b3JsZHRoaXNpc2Fwcm9kdWN0aW9ua2V5  # Base64-encoded 128-bit key
        - key_name: backup-key-2024
          key_value: YmFja3Vwa2V5Zm9yY2RucmVkaXN0cmlidXRpb24  # For key rotation
```

**Generate signing keys:**
```bash
# Generate a secure 128-bit key
openssl rand -base64 16
```

**Application integration:**
Your application generates signed URLs with expiration time:
```python
# Python example (using google-auth library)
from google.auth.transport import requests
from google.auth import compute_engine
import base64
import time

def sign_url(url, key_name, key_value, expires_in=3600):
    expiration = int(time.time()) + expires_in
    url_to_sign = f"{url}?Expires={expiration}&KeyName={key_name}"
    # Create signature...
    return url_to_sign
```

---

## 6. Cloud Run Backend with CDN

Accelerate a serverless Cloud Run application with Cloud CDN.

**Use case:** API with cacheable responses, dynamic web app with static assets

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: cloud-run-api-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    cloud_run_service:
      service_name: my-api-service
      region: us-central1
  cache_mode: USE_ORIGIN_HEADERS  # Cloud Run app sets Cache-Control headers
  default_ttl_seconds: 300         # 5 minutes for uncached responses
  max_ttl_seconds: 3600            # 1 hour max
  client_ttl_seconds: 600          # 10 minutes for browsers
  enable_negative_caching: true
  frontend_config:
    custom_domains:
      - api.mycompany.com
    ssl_certificate:
      google_managed:
        domains:
          - api.mycompany.com
    enable_https_redirect: true
```

**Cloud Run application must:**
1. Send `Cache-Control` headers for cacheable responses:
   ```python
   # Python Flask example
   @app.route('/api/data')
   def get_data():
       resp = make_response(jsonify(data))
       resp.headers['Cache-Control'] = 'public, max-age=300'  # Cache for 5 minutes
       return resp
   ```

2. Send `Cache-Control: no-store` for dynamic/user-specific responses:
   ```python
   @app.route('/api/user/profile')
   def get_profile():
       resp = make_response(jsonify(user_data))
       resp.headers['Cache-Control'] = 'no-store, private'  # Never cache
       return resp
   ```

---

## 7. Compute Engine Backend with Health Checks

CDN in front of Compute Engine VMs (Managed Instance Group).

**Use case:** Traditional web application on VMs, monolithic application

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: webapp-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    compute_service:
      instance_group_name: webapp-mig-us-central1
      protocol: HTTP
      port: 8080
      health_check:
        path: /healthz
        port: 8080
        check_interval_seconds: 10
        timeout_seconds: 5
        healthy_threshold: 2
        unhealthy_threshold: 3
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 3600
  max_ttl_seconds: 86400
  enable_negative_caching: true
  advanced_config:
    cache_key_policy:
      include_query_string: true
      query_string_whitelist:
        - page
        - category
        - sort
      include_protocol: true
      include_host: true
  frontend_config:
    custom_domains:
      - www.mywebapp.com
    ssl_certificate:
      google_managed:
        domains:
          - www.mywebapp.com
    enable_https_redirect: true
```

**Prerequisites:**
- Managed Instance Group `webapp-mig-us-central1` must exist in the project
- VMs must respond to health checks at `/healthz` on port 8080
- VMs must serve HTTP traffic on port 8080

---

## 8. External Origin (Hybrid/Multi-Cloud)

Use Cloud CDN to accelerate content from an origin outside GCP (AWS S3, on-prem server, another cloud).

**Use case:** Gradual migration to GCP, multi-cloud architecture, hybrid deployments

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: hybrid-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    external_origin:
      hostname: assets.mycompany.com  # Could be AWS S3, on-prem, etc.
      port: 443
      protocol: HTTPS
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 3600
  max_ttl_seconds: 86400
  enable_negative_caching: true
  frontend_config:
    custom_domains:
      - cdn.mycompany.com
    ssl_certificate:
      google_managed:
        domains:
          - cdn.mycompany.com
    enable_https_redirect: true
```

**What this enables:**
- External origin (e.g., AWS S3 bucket at `assets.mycompany.com`) remains the source of truth
- Google Cloud CDN caches content at its global edge network
- End users benefit from Google's edge locations without migrating data

**Use case:** Gradual AWS → GCP migration:
1. Phase 1: Keep data in AWS S3, use GCP Cloud CDN (this example)
2. Phase 2: Migrate data to GCS, switch backend to `gcs_bucket`
3. Users see no downtime during migration

---

## 9. Production CDN with Cloud Armor

Full production deployment with WAF/DDoS protection via Cloud Armor.

**Use case:** E-commerce site, high-value application requiring security

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: production-secure-cdn
spec:
  gcp_project_id: my-company-prod
  backend:
    gcs_bucket:
      bucket_name: ecommerce-static-assets
      enable_uniform_access: true
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 3600
  max_ttl_seconds: 86400
  client_ttl_seconds: 1800
  enable_negative_caching: true
  advanced_config:
    cache_key_policy:
      include_query_string: true
      query_string_whitelist:
        - product_id
        - category
        - page
      include_protocol: true
      include_host: true
    enable_request_coalescing: true
    serve_while_stale_seconds: 60  # Serve stale content for 1 min while revalidating
  frontend_config:
    custom_domains:
      - www.mystore.com
      - mystore.com
    ssl_certificate:
      google_managed:
        domains:
          - www.mystore.com
          - mystore.com
    cloud_armor:
      enabled: true
      security_policy_name: ecommerce-security-policy  # Pre-existing Cloud Armor policy
    enable_https_redirect: true
```

**Cloud Armor security policy must exist:**
```bash
# Create a Cloud Armor policy (if not already exists)
gcloud compute security-policies create ecommerce-security-policy \
  --description "WAF and DDoS protection for e-commerce site"

# Add rules (example: block SQL injection patterns)
gcloud compute security-policies rules create 1000 \
  --security-policy ecommerce-security-policy \
  --expression "evaluatePreconfiguredExpr('sqli-stable')" \
  --action deny-403
```

---

## 10. Advanced Configuration with Negative Caching

Granular control over error response caching to optimize origin load and user experience.

**Use case:** API with different error handling strategies for different status codes

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudCdn
metadata:
  name: advanced-api-cdn
spec:
  gcp_project_id: my-project-123456
  backend:
    cloud_run_service:
      service_name: advanced-api
      region: us-central1
  cache_mode: USE_ORIGIN_HEADERS
  default_ttl_seconds: 600
  max_ttl_seconds: 3600
  client_ttl_seconds: 300
  enable_negative_caching: true
  advanced_config:
    cache_key_policy:
      include_query_string: true
      query_string_whitelist:
        - api_version
        - format
      include_protocol: true
      include_host: true
      included_headers:
        - Accept
        - Accept-Encoding
    negative_caching_policies:
      - code: 404
        ttl_seconds: 600  # Cache "Not Found" for 10 minutes (reduces origin load)
      - code: 500
        ttl_seconds: 10   # Cache "Internal Server Error" for 10 seconds only
      - code: 502
        ttl_seconds: 10   # Cache "Bad Gateway" for 10 seconds only
      - code: 503
        ttl_seconds: 30   # Cache "Service Unavailable" for 30 seconds
    serve_while_stale_seconds: 120  # Serve stale content for 2 minutes during origin outage
    enable_request_coalescing: true
  frontend_config:
    custom_domains:
      - api.example.com
    ssl_certificate:
      google_managed:
        domains:
          - api.example.com
    enable_https_redirect: true
```

**Why this configuration?**

- **404 cached for 10 minutes:** Invalid API endpoints won't hammer origin (e.g., bot scanning)
- **500/502 cached for 10 seconds:** Transient errors are cached briefly to protect origin, but not so long that legitimate users are affected
- **503 cached for 30 seconds:** Service temporarily unavailable (e.g., during deployment) - cache longer than 500 errors
- **Serve-while-stale 120 seconds:** If origin goes down, CDN serves cached content for 2 minutes, giving origin time to recover

---

## Deployment Instructions

### Using Pulumi

```bash
# Navigate to the component directory
cd apis/org/project_planton/provider/gcp/gcpcloudcdn/v1/iac/pulumi

# Set your GCP project
export GCP_PROJECT_ID=my-project-123456

# Create a manifest file (e.g., manifest.yaml) with your desired configuration
# Then deploy
planton pulumi up --manifest manifest.yaml
```

### Using Terraform

```bash
# Navigate to the Terraform module directory
cd apis/org/project_planton/provider/gcp/gcpcloudcdn/v1/iac/tf

# Initialize Terraform
terraform init

# Create a tfvars file or pass variables
terraform plan -var-file=dev.tfvars

# Apply the configuration
terraform apply -var-file=dev.tfvars
```

---

## Post-Deployment Verification

### Check CDN Status

```bash
# Get backend bucket details
gcloud compute backend-buckets describe <backend-name> --global

# Check if CDN is enabled
gcloud compute backend-buckets describe <backend-name> --global \
  --format="get(cdnPolicy.cacheMode)"
```

### Test Cache Behavior

```bash
# First request (cache MISS)
curl -I https://your-cdn-url.com/test.jpg
# Look for: X-Cache: MISS

# Second request (cache HIT)
curl -I https://your-cdn-url.com/test.jpg
# Look for: X-Cache: HIT
```

### Monitor Cache Hit Ratio

Navigate to Cloud Console → Network Services → Cloud CDN

Or query via gcloud:
```bash
gcloud logging read "resource.type=http_load_balancer
  AND jsonPayload.statusDetails=response_from_cache" \
  --limit 100 --format json
```

---

## Common Patterns

### Pattern: Versioned URLs (Best Practice)

Instead of cache invalidation, use versioned filenames:

```yaml
# Configure long TTLs
spec:
  cache_mode: CACHE_ALL_STATIC
  default_ttl_seconds: 31536000  # 1 year
  max_ttl_seconds: 31536000      # 1 year
```

Your build process generates hashed filenames:
```
# Build output
app.a83f5b2c.js
app.f19d84e1.css
logo.9c3f4a7b.png
```

HTML references the versioned files:
```html
<script src="/app.a83f5b2c.js"></script>
<link rel="stylesheet" href="/app.f19d84e1.css">
```

**Benefit:** No cache invalidation needed. New deployment = new filenames = instant updates.

---

## Troubleshooting

### Low Cache Hit Ratio

**Symptoms:** High origin load, slow response times, high costs

**Common causes:**
1. **Cache shattering:** Too many query parameters in cache key
   - **Fix:** Use `query_string_whitelist` in `cache_key_policy`
2. **Unsupported Vary headers:** Origin sends `Vary: Cookie`
   - **Fix:** Remove unsupported Vary headers from origin (Cloud CDN only supports Vary: Accept, Accept-Encoding, Origin)
3. **Incorrect cache mode:** Using `USE_ORIGIN_HEADERS` without sending Cache-Control
   - **Fix:** Switch to `CACHE_ALL_STATIC` or ensure origin sends proper headers

### SSL Certificate Not Provisioning

**Symptoms:** `https://` doesn't work, certificate pending

**Common causes:**
1. DNS not pointed to load balancer IP
   - **Fix:** Point A record to `global_ip_address` output
2. Certificate validation in progress (takes 10-15 minutes)
   - **Fix:** Wait for Google to complete validation

### Origin Errors Not Cached

**Symptoms:** 404/500 errors hitting origin repeatedly

**Fix:** Enable `enable_negative_caching` and configure `negative_caching_policies`

---

*For more information, see the [Research Document](docs/README.md) for architectural deep-dive and best practices.*
