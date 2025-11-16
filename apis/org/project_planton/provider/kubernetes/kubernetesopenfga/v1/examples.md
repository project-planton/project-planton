# Multiple Examples for `OpenFgaKubernetes` API-Resource

## Example with Minimal Configuration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: basic-openfga
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
      limits:
        cpu: 200m
        memory: 512Mi
  datastore:
    engine: postgres
    uri: postgres://user:password@db-host:5432/openfga
```

---

## Example with Ingress Enabled

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-service
spec:
  container:
    replicas: 3
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
  ingress:
    enabled: true
    hostname: openfga.example.com
  datastore:
    engine: postgres
    uri: postgres://user:password@db-host:5432/openfga
```

---

## Example with MySQL Datastore

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-mysql
spec:
  container:
    replicas: 2
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
  datastore:
    engine: mysql
    uri: mysql://user:password@mysql-db:3306/openfga
```

---

## Example with High Availability Configuration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-high-availability
spec:
  container:
    replicas: 5
    resources:
      requests:
        cpu: 500m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 4Gi
  ingress:
    enabled: true
    hostname: openfga-ha.example.com
  datastore:
    engine: postgres
    uri: postgres://user:securepassword@ha-db-host:5432/openfga
```

---

## Example with Production-Grade Resources

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-production
spec:
  container:
    replicas: 3
    resources:
      requests:
        cpu: 1000m
        memory: 2Gi
      limits:
        cpu: 4000m
        memory: 8Gi
  ingress:
    enabled: true
    hostname: openfga-prod.company.com
  datastore:
    engine: postgres
    uri: postgres://openfga_user:securepassword@prod-db-host:5432/openfga_prod?sslmode=require
```
