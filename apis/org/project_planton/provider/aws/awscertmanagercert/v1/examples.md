# AWS Cert Manager Cert Examples

## Create Using CLI

Create a YAML manifest using one of the examples below. After the YAML is created, apply it with ProjectPlanton:

```shell
project-planton pulumi up --manifest <yaml-path> --stack <stack-name>
```

Or, if using Terraform:

```shell
project-planton terraform apply --manifest <yaml-path> --stack <stack-name>
```

(You can also use the shorter form `planton apply -f <yaml-path>` if your environment is configured accordingly.)

---

## Basic Example (Single Domain)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCertManagerCert
metadata:
  name: single-cert-example
spec:
  primaryDomainName: example.com
  region: us-east-1
  route53HostedZoneId: Z123456ABCXYZ
```

This manifest requests an SSL certificate for `example.com` in `us-east-1` and uses Route53 hosted zone `Z123456ABCXYZ`
for DNS validation.

---

## Multiple Domains Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCertManagerCert
metadata:
  name: multi-cert-example
spec:
  primaryDomainName: myapp.io
  alternateDomainNames:
    - api.myapp.io
    - blog.myapp.io
  region: us-east-2
  route53HostedZoneId: Z987654XYZABC
  validationMethod: DNS
```

In this example, we issue a certificate for `myapp.io` as the primary domain and add two alternate domains:
`api.myapp.io` and `blog.myapp.io`. It uses `us-east-2` as the region and automatically creates DNS records in the
Route53 hosted zone `Z987654XYZABC`.

---

## Wildcard Domain Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCertManagerCert
metadata:
  name: wildcard-example
spec:
  primaryDomainName: '*.wildexample.com'
  region: us-east-1
  route53HostedZoneId: Z999999EXAMPLE
```

This manifests requests a wildcard certificate for `*.wildexample.com`, allowing you to secure any subdomain under
`wildexample.com` without creating separate certificates.

---

## Minimal Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCertManagerCert
metadata:
  name: minimal-cert
spec:
  primaryDomainName: barebones.io
  region: us-east-1
  route53HostedZoneId: Z1111111111111
```

Here, we only specify the required fields. The certificate will be created for `barebones.io`, with DNS records
automatically managed in the hosted zone. Alternate domain names are not specified.

---

## After Deploying

Once you’ve chosen one of these examples, apply it with ProjectPlanton using Pulumi or Terraform:

```shell
project-planton pulumi up --manifest aws-cert-manager-cert.yaml --stack myorg/dev
```

or

```shell
project-planton terraform apply --manifest aws-cert-manager-cert.yaml --stack myorg/dev
```

When the command completes successfully, your SSL certificate will be provisioned in ACM. You can confirm by checking
the AWS console or by using the AWS CLI:

```shell
aws acm list-certificates
```

You should see your new certificate and any domain validation records created within Route53.

---

Enjoy secure connections with ProjectPlanton!
