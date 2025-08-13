# Examples

## Minimal manifest (YAML)
```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  org: acme
  env: dev
  name: demo-cloudfront
  id: cf-001
spec:
  aliases: []
  price_class: PRICE_CLASS_100
  origins:
    - id: origin-1
      domain_name: example.com
  default_cache_behavior:
    origin_id: origin-1
    viewer_protocol_policy: REDIRECT_TO_HTTPS
    compress: true
    allowed_methods: GET_HEAD
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

Note: credentials are resolved via stack input; they are not part of the spec.
