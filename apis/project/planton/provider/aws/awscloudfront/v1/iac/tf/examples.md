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

- Tofu plan/apply:
```bash
project-planton tofu init --manifest ./manifest.yaml
project-planton tofu plan --manifest ./manifest.yaml
project-planton tofu apply --manifest ./manifest.yaml --auto-approve
```

Note: Provider credentials are supplied via stack input, not in the spec.


