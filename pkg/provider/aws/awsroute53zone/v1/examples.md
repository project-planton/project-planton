# AWS Route53 Zone Examples

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

## Basic Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRoute53Zone
metadata:
  name: example.com
spec:
  records:
    - recordType: A
      name: example.com.
      values:
        - "192.0.2.1"
      ttlSeconds: 300
```

This sets up a Route53 hosted zone for `example.com` and creates a simple A record pointing to the IP address
`192.0.2.1`.

---

## Multiple Records Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRoute53Zone
metadata:
  name: example-multi.com
spec:
  records:
    - recordType: A
      name: www.example-multi.com.
      values:
        - "192.0.2.10"
      ttlSeconds: 300
    - recordType: CNAME
      name: api.example-multi.com.
      values:
        - "www.example-multi.com."
      ttlSeconds: 300
```

In this manifest, the hosted zone for `example-multi.com` has:
• An A record to route `www.example-multi.com` to `192.0.2.10`.
• A CNAME record `api.example-multi.com` that points to `www.example-multi.com.`.

---

## MX Record Example

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRoute53Zone
metadata:
  name: mail-example.com
spec:
  records:
    - recordType: MX
      name: mail-example.com.
      values:
        - "10 mailserver1.example.org."
        - "20 mailserver2.example.org."
      ttlSeconds: 3600
```

Here, the zone `mail-example.com` includes an MX record with two mail servers:
• Priority 10: `mailserver1.example.org.`
• Priority 20: `mailserver2.example.org.`

---

## TXT Record with Multiple Values

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRoute53Zone
metadata:
  name: txt-example.com
spec:
  records:
    - recordType: TXT
      name: txt-example.com.
      values:
        - "\"google-site-verification=abc123\""
        - "\"v=spf1 include:_spf.example.com ~all\""
      ttlSeconds: 300
```

This example adds a TXT record with multiple strings, which can be used for verification tokens or SPF records.

---

## Minimal Example (No Records)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRoute53Zone
metadata:
  name: minimal.com
spec:
  records: [ ]
```

This creates a hosted zone named `minimal.com` with no DNS records. It’s useful when you simply want to provision the
zone first and add records later.

---

## After Deploying

Once you’ve chosen one of these examples, apply it with ProjectPlanton using Pulumi or Terraform:

```shell
project-planton pulumi up --manifest aws-route53-zone.yaml --stack myorg/dev
```

or

```shell
project-planton terraform apply --manifest aws-route53-zone.yaml --stack myorg/dev
```

When the command completes successfully, your Route53 zone will be created in AWS. You can confirm by checking the AWS
console or by using the AWS CLI:

```shell
aws route53 list-hosted-zones
```

You should see your new hosted zone, along with any records that you defined in the manifest.

---

Happy routing with ProjectPlanton!
