# Examples

## Minimal manifest (YAML)
```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: my-cdn
spec:
  enabled: true
  priceClass: PRICE_CLASS_100
  origins:
    - id: origin-1
      domainName: bucket.s3.amazonaws.com
  defaultOriginId: origin-1
  defaultRootObject: index.html
```

## CLI flows
- Validate:
```bash
project-planton validate --manifest ./manifest.yaml
```

- Pulumi deploy:
```bash
project-planton pulumi up --manifest ./manifest.yaml --stack org/project/stack
```

- Terraform deploy:
```bash
project-planton terraform apply --manifest ./manifest.yaml --stack org/project/stack
```

Note: Provider credentials are supplied via stack input, not in the spec.
