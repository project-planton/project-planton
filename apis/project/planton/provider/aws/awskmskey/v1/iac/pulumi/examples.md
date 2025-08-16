```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsKmsKey
metadata:
  name: example
spec: {}
```

CLI:

```bash
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes
```


