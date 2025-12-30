# Auth0Connection Examples

This document provides complete, copy-paste ready examples for configuring Auth0 connections using Project Planton.

## Database Connections

### Basic Auth0 Database

Minimal database connection with default settings:

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: user-database
  org: my-organization
  env: development
spec:
  strategy: auth0
  display_name: Sign Up with Email
  enabled_clients:
    - "your-app-client-id"
```

### Production Database with Security

Database connection with enhanced security settings:

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: production-database
  org: my-organization
  env: production
  labels:
    team: identity
    security-level: high
spec:
  strategy: auth0
  display_name: Create Account
  enabled_clients:
    - "web-app-client-id"
    - "mobile-app-client-id"
  database_options:
    password_policy: excellent
    requires_username: false
    disable_signup: false
    brute_force_protection: true
    password_history_size: 10
    password_no_personal_info: true
    password_dictionary: true
    mfa_enabled: true
```

## Social Connections

### Google OAuth

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: google-social
  org: my-organization
  env: production
spec:
  strategy: google-oauth2
  display_name: Continue with Google
  show_as_button: true
  enabled_clients:
    - "web-app-client-id"
  social_options:
    client_id: "123456789-abcdef.apps.googleusercontent.com"
    client_secret: "GOCSPX-your-secret-here"
    scopes:
      - openid
      - profile
      - email
```

### GitHub OAuth

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: github-login
  org: my-organization
  env: production
spec:
  strategy: github
  display_name: Sign in with GitHub
  show_as_button: true
  enabled_clients:
    - "developer-portal-client-id"
  social_options:
    client_id: "Iv1.your-github-app-id"
    client_secret: "your-github-client-secret"
    scopes:
      - read:user
      - user:email
      - read:org
```

### Facebook Login

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: facebook-social
  org: my-organization
  env: production
spec:
  strategy: facebook
  display_name: Continue with Facebook
  show_as_button: true
  enabled_clients:
    - "consumer-app-client-id"
  social_options:
    client_id: "your-facebook-app-id"
    client_secret: "your-facebook-app-secret"
    scopes:
      - email
      - public_profile
```

### Microsoft Account

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: microsoft-personal
  org: my-organization
  env: production
spec:
  strategy: microsoft-account
  display_name: Sign in with Microsoft
  show_as_button: true
  enabled_clients:
    - "web-app-client-id"
  social_options:
    client_id: "your-azure-app-id"
    client_secret: "your-azure-client-secret"
    scopes:
      - openid
      - profile
      - email
```

## Enterprise Connections

### SAML with Okta

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: okta-enterprise-sso
  org: my-organization
  env: production
spec:
  strategy: samlp
  display_name: Company SSO
  is_domain_connection: true
  realms:
    - company.com
    - company.io
  enabled_clients:
    - "internal-app-client-id"
  saml_options:
    sign_in_endpoint: "https://company.okta.com/app/app-id/sso/saml"
    signing_cert: |
      -----BEGIN CERTIFICATE-----
      MIICmTCCAYGgAwIBAgIJAKc...
      -----END CERTIFICATE-----
    sign_out_endpoint: "https://company.okta.com/app/app-id/slo/saml"
    entity_id: "http://www.okta.com/exk123abc"
    protocol_binding: "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
    sign_request: true
    signature_algorithm: rsa-sha256
    digest_algorithm: sha256
    attribute_mappings:
      email: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
      given_name: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname"
      family_name: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"
```

### OIDC with Keycloak

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: keycloak-oidc
  org: my-organization
  env: production
spec:
  strategy: oidc
  display_name: Login with Keycloak
  is_domain_connection: true
  enabled_clients:
    - "internal-app-client-id"
  oidc_options:
    issuer: "https://keycloak.company.com/realms/main"
    client_id: "auth0-integration"
    client_secret: "keycloak-client-secret"
    scopes:
      - openid
      - profile
      - email
      - groups
    type: front_channel
    attribute_mappings:
      email: email
      name: name
      given_name: given_name
      family_name: family_name
```

### Azure AD / Entra ID

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: azure-ad-enterprise
  org: my-organization
  env: production
spec:
  strategy: waad
  display_name: Microsoft Work Account
  is_domain_connection: true
  realms:
    - contoso.com
  enabled_clients:
    - "corporate-app-client-id"
  azure_ad_options:
    client_id: "your-azure-app-id-guid"
    client_secret: "your-azure-client-secret"
    domain: "contoso.onmicrosoft.com"
    tenant_id: "your-tenant-guid"
    use_common_endpoint: false
    max_groups_to_retrieve: 100
    should_trust_email_verified: true
    api_enable_users: false
```

### Azure AD Multi-Tenant

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: azure-ad-multi-tenant
  org: my-organization
  env: production
spec:
  strategy: waad
  display_name: Sign in with Work Account
  enabled_clients:
    - "saas-app-client-id"
  azure_ad_options:
    client_id: "your-multi-tenant-app-id"
    client_secret: "your-azure-client-secret"
    domain: "common"
    use_common_endpoint: true
    should_trust_email_verified: true
```

## Multi-Environment Configuration

### Development

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: user-database
  org: my-organization
  env: development
spec:
  strategy: auth0
  display_name: Development Login
  database_options:
    password_policy: low
    brute_force_protection: false
    disable_signup: false
```

### Staging

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: user-database
  org: my-organization
  env: staging
spec:
  strategy: auth0
  display_name: Staging Login
  database_options:
    password_policy: good
    brute_force_protection: true
    mfa_enabled: false
```

### Production

```yaml
apiVersion: auth0.project-planton.org/v1
kind: Auth0Connection
metadata:
  name: user-database
  org: my-organization
  env: production
spec:
  strategy: auth0
  display_name: Sign Up
  database_options:
    password_policy: excellent
    brute_force_protection: true
    password_history_size: 10
    password_no_personal_info: true
    password_dictionary: true
    mfa_enabled: true
```

## Deployment Commands

### Deploy a connection

```bash
project-planton apply --manifest auth0-connection.yaml \
  --auth0-provider-config auth0-creds.yaml
```

### Preview changes

```bash
project-planton plan --manifest auth0-connection.yaml \
  --auth0-provider-config auth0-creds.yaml
```

### Destroy a connection

```bash
project-planton destroy --manifest auth0-connection.yaml \
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
- `create:connections`
- `read:connections`
- `update:connections`
- `delete:connections`

