# Examples for AwsRdsInstance Pulumi Module

## Minimal Postgres instance in private subnets

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: db-postgres-dev
spec:
  subnetIds:
    - valueFrom:
        kind: AwsVpc
        name: network
        fieldPath: status.outputs.private_subnets.[*].id
    - valueFrom:
        kind: AwsVpc
        name: network
        fieldPath: status.outputs.private_subnets.[*].id
  securityGroupIds:
    - valueFrom:
        kind: AwsSecurityGroup
        name: db-sg
        fieldPath: status.outputs.security_group_id
  engine: postgres
  engineVersion: "14.10"
  instanceClass: db.t3.micro
  allocatedStorageGb: 20
  storageEncrypted: true
  username: master
  password: ${secrets-group/db/MASTER_PASSWORD}
  port: 5432
  publiclyAccessible: false
  multiAz: false
```

## Publicly-accessible MySQL using an existing subnet group

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: db-mysql-public
spec:
  dbSubnetGroupName:
    value: app-db-subnet-group
  securityGroupIds:
    - value: sg-0123456789abcdef0
  engine: mysql
  engineVersion: "8.0.35"
  instanceClass: db.m6g.large
  allocatedStorageGb: 100
  username: admin
  password: ${secrets-group/db/ADMIN_PASSWORD}
  port: 3306
  publiclyAccessible: true
  multiAz: true
```

## CLI flows (run from this pulumi directory)

Preview:
```bash
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

Apply:
```bash
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes
```

Refresh:
```bash
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

Destroy:
```bash
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes
```
