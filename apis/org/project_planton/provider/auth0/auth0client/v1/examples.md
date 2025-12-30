# Auth0Client Examples

This document provides complete, copy-paste ready examples for configuring Auth0 applications using Project Planton.

## Single Page Applications (SPA)

### Basic SPA

Minimal SPA configuration:

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: my-spa
  org: my-organization
  env: development
spec:
  application_type: spa
  description: My React Application
  callbacks:
    - http://localhost:3000/callback
  web_origins:
    - http://localhost:3000
```

### Production SPA with Refresh Tokens

Full-featured SPA with secure refresh token configuration:

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: production-spa
  org: my-organization
  env: production
  labels:
    team: frontend
    security-level: high
spec:
  application_type: spa
  description: Production React Application
  callbacks:
    - https://app.example.com/callback
  allowed_logout_urls:
    - https://app.example.com
  web_origins:
    - https://app.example.com
  grant_types:
    - authorization_code
    - refresh_token
  oidc_conformant: true
  is_first_party: true
  refresh_token:
    rotation_type: rotating
    expiration_type: expiring
    token_lifetime: 2592000  # 30 days
    idle_token_lifetime: 1296000  # 15 days
    leeway: 60
```

### SPA with Cross-Origin Authentication

For embedded login forms:

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: embedded-login-spa
  org: my-organization
  env: production
spec:
  application_type: spa
  cross_origin_authentication: true
  cross_origin_loc: https://app.example.com/cross-origin-callback
  callbacks:
    - https://app.example.com/callback
  web_origins:
    - https://app.example.com
  allowed_origins:
    - https://app.example.com
```

## Native Applications (Mobile/Desktop)

### iOS Application

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: ios-app
  org: my-organization
  env: production
spec:
  application_type: native
  description: iOS Mobile Application
  callbacks:
    - com.example.myapp://callback
    - myapp://callback
  allowed_logout_urls:
    - com.example.myapp://logout
  grant_types:
    - authorization_code
    - refresh_token
  mobile:
    ios:
      team_id: ABCDE12345
      app_bundle_identifier: com.example.myapp
  native_social_login:
    apple:
      enabled: true
```

### Android Application

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: android-app
  org: my-organization
  env: production
spec:
  application_type: native
  description: Android Mobile Application
  callbacks:
    - com.example.myapp://callback
  allowed_logout_urls:
    - com.example.myapp://logout
  grant_types:
    - authorization_code
    - refresh_token
  mobile:
    android:
      app_package_name: com.example.myapp
      sha256_cert_fingerprints:
        - "D8:A0:1B:2C:3D:4E:5F:6G:7H:8I:9J:0K:1L:2M:3N:4O:5P:6Q:7R:8S"
```

### Cross-Platform Mobile App

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: mobile-app
  org: my-organization
  env: production
spec:
  application_type: native
  description: Cross-platform mobile app (React Native/Flutter)
  callbacks:
    - com.example.myapp://callback
    - myapp://callback
  allowed_logout_urls:
    - com.example.myapp://logout
    - myapp://logout
  grant_types:
    - authorization_code
    - refresh_token
  mobile:
    ios:
      team_id: ABCDE12345
      app_bundle_identifier: com.example.myapp
    android:
      app_package_name: com.example.myapp
      sha256_cert_fingerprints:
        - "D8:A0:1B:..."
  native_social_login:
    apple:
      enabled: true
    facebook:
      enabled: true
  refresh_token:
    rotation_type: rotating
    expiration_type: expiring
    token_lifetime: 2592000
```

## Regular Web Applications

### Basic Web Application

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: web-app
  org: my-organization
  env: production
spec:
  application_type: regular_web
  description: Node.js web application
  callbacks:
    - https://webapp.example.com/auth/callback
  allowed_logout_urls:
    - https://webapp.example.com
  grant_types:
    - authorization_code
    - refresh_token
```

### Web App with Custom JWT Settings

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: secure-web-app
  org: my-organization
  env: production
spec:
  application_type: regular_web
  description: Secure web application with custom JWT
  callbacks:
    - https://secure.example.com/callback
  allowed_logout_urls:
    - https://secure.example.com
  grant_types:
    - authorization_code
    - refresh_token
  jwt_configuration:
    lifetime_in_seconds: 3600  # 1 hour
    alg: RS256
    scopes:
      admin: Full administrative access
      read: Read-only access
```

### Web App with OIDC Backchannel Logout

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: oidc-web-app
  org: my-organization
  env: production
spec:
  application_type: regular_web
  description: Web app with backchannel logout
  callbacks:
    - https://app.example.com/callback
  allowed_logout_urls:
    - https://app.example.com
  oidc_backchannel_logout:
    backchannel_logout_urls:
      - https://app.example.com/backchannel-logout
```

## Machine-to-Machine (M2M) Applications

### Basic API Client

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: api-client
  org: my-organization
  env: production
spec:
  application_type: non_interactive
  description: Backend API client
  grant_types:
    - client_credentials
```

### M2M with Custom Token Lifetime

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: scheduled-job-client
  org: my-organization
  env: production
spec:
  application_type: non_interactive
  description: Scheduled job service account
  grant_types:
    - client_credentials
  jwt_configuration:
    lifetime_in_seconds: 86400  # 24 hours
    alg: RS256
  client_metadata:
    service_type: scheduled_job
    owner_team: platform
```

### M2M for CI/CD Pipeline

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: cicd-client
  org: my-organization
  env: production
  labels:
    purpose: cicd
spec:
  application_type: non_interactive
  description: CI/CD pipeline authentication
  grant_types:
    - client_credentials
  jwt_configuration:
    lifetime_in_seconds: 3600  # 1 hour
  is_token_endpoint_ip_header_trusted: true
```

## Organization-Enabled Applications

### App with Optional Organization

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: saas-app
  org: my-organization
  env: production
spec:
  application_type: regular_web
  organization_usage: allow
  callbacks:
    - https://saas.example.com/callback
```

### App Requiring Organization

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: b2b-saas-app
  org: my-organization
  env: production
spec:
  application_type: regular_web
  organization_usage: require
  organization_require_behavior: pre_login_prompt
  callbacks:
    - https://b2b-saas.example.com/callback
```

## Multi-Environment Configuration

### Development

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: my-app
  org: my-organization
  env: development
spec:
  application_type: spa
  callbacks:
    - http://localhost:3000/callback
    - http://localhost:4200/callback
  web_origins:
    - http://localhost:3000
    - http://localhost:4200
  sso_disabled: true  # Disable SSO for dev
```

### Staging

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: my-app
  org: my-organization
  env: staging
spec:
  application_type: spa
  callbacks:
    - https://staging.example.com/callback
  web_origins:
    - https://staging.example.com
  jwt_configuration:
    lifetime_in_seconds: 7200  # 2 hours for easier testing
```

### Production

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Client
metadata:
  name: my-app
  org: my-organization
  env: production
spec:
  application_type: spa
  callbacks:
    - https://app.example.com/callback
  allowed_logout_urls:
    - https://app.example.com
  web_origins:
    - https://app.example.com
  oidc_conformant: true
  jwt_configuration:
    lifetime_in_seconds: 3600  # 1 hour
  refresh_token:
    rotation_type: rotating
    expiration_type: expiring
```

## Deployment Commands

### Deploy a client

```bash
project-planton apply --manifest auth0-client.yaml \
  --auth0-provider-config auth0-creds.yaml
```

### Preview changes

```bash
project-planton plan --manifest auth0-client.yaml \
  --auth0-provider-config auth0-creds.yaml
```

### Destroy a client

```bash
project-planton destroy --manifest auth0-client.yaml \
  --auth0-provider-config auth0-creds.yaml
```

## Auth0 Credential File

Create `auth0-creds.yaml`:

```yaml
domain: your-tenant.auth0.com
clientId: your-m2m-client-id
clientSecret: your-m2m-client-secret
```

Ensure the Machine-to-Machine application has the following permissions:
- `create:clients`
- `read:clients`
- `update:clients`
- `delete:clients`
- `read:client_keys`


