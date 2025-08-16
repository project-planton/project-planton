# Terraform examples for AwsDynamodb

## Minimal manifest (YAML)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: orders
  labels:
    app: shop
spec:
  billingMode: BILLING_MODE_PAY_PER_REQUEST
  attributeDefinitions:
    - name: pk
      type: ATTRIBUTE_TYPE_S
  keySchema:
    - attributeName: pk
      keyType: KEY_TYPE_HASH
```

## PROVISIONED with GSI and TTL

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: orders
spec:
  billingMode: BILLING_MODE_PROVISIONED
  provisionedThroughput:
    readCapacityUnits: 10
    writeCapacityUnits: 10
  attributeDefinitions:
    - name: pk
      type: ATTRIBUTE_TYPE_S
    - name: sk
      type: ATTRIBUTE_TYPE_S
    - name: status
      type: ATTRIBUTE_TYPE_S
  keySchema:
    - attributeName: pk
      keyType: KEY_TYPE_HASH
    - attributeName: sk
      keyType: KEY_TYPE_RANGE
  globalSecondaryIndexes:
    - name: status-index
      keySchema:
        - attributeName: status
          keyType: KEY_TYPE_HASH
      projection:
        type: PROJECTION_TYPE_KEYS_ONLY
      provisionedThroughput:
        readCapacityUnits: 5
        writeCapacityUnits: 5
  ttl:
    enabled: true
    attributeName: expiresAt
  streamEnabled: true
  streamViewType: STREAM_VIEW_TYPE_NEW_AND_OLD_IMAGES
```

## CLI flows (tofu)

Init:
```bash
project-planton tofu init --manifest ../hack/manifest.yaml | cat
```

Plan:
```bash
project-planton tofu plan --manifest ../hack/manifest.yaml | cat
```

Apply:
```bash
project-planton tofu apply --manifest ../hack/manifest.yaml --auto-approve | cat
```

Destroy:
```bash
project-planton tofu destroy --manifest ../hack/manifest.yaml --auto-approve | cat
```

> Note: Provider credentials are supplied via stack input, not in the spec.
