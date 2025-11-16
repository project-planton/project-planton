# Civo Load Balancer

Provision and manage Civo Load Balancers through Project Planton's declarative infrastructure-as-code approach.

## Overview

`CivoLoadBalancer` provides a Kubernetes-native way to deploy production-grade load balancers on Civo Cloud. It abstracts the complexity of configuring forwarding rules, health checks, and backend instances while enforcing best practices like private networking and reserved IPs.

### Key Features

- **Multiple Protocol Support**: HTTP, HTTPS (TLS passthrough), and TCP
- **Flexible Backend Selection**: Attach instances by explicit IDs or dynamically via tags
- **Advanced Health Checks**: HTTP/HTTPS path-based or TCP port checks
- **Sticky Sessions**: Optional session affinity using IP hashing
- **Reserved IP Integration**: Stable public endpoints that survive infrastructure changes
- **Multiple Forwarding Rules**: Route different ports/protocols to various backend ports

## Quick Start

### Basic Load Balancer

Deploy a simple HTTP load balancer with explicit instance attachment:

```yaml
apiVersion: code.project-planton.io/v1
kind: CivoLoadBalancer
metadata:
  name: web-lb
spec:
  loadBalancerName: web-lb
  region: lon1
  network:
    value: "net-12345678-abcd-1234-abcd-1234567890ab"
  forwardingRules:
    - entryPort: 80
      entryProtocol: http
      targetPort: 8080
      targetProtocol: http
  instanceIds:
    - value: "inst-web-1"
    - value: "inst-web-2"
```

### Production Load Balancer with HA

Deploy a production-ready load balancer with HTTPS, health checks, and reserved IP:

```yaml
apiVersion: code.project-planton.io/v1
kind: CivoLoadBalancer
metadata:
  name: prod-lb
  env: production
spec:
  loadBalancerName: prod-api-lb
  region: lon1
  network:
    kind: CivoVpc
    ref: prod-network
  forwardingRules:
    - entryPort: 80
      entryProtocol: http
      targetPort: 8080
      targetProtocol: http
    - entryPort: 443
      entryProtocol: https
      targetPort: 8443
      targetProtocol: https
  healthCheck:
    port: 8080
    protocol: http
    path: /healthz
  instanceIds:
    - kind: CivoComputeInstance
      ref: web-1
    - kind: CivoComputeInstance
      ref: web-2
  reservedIpId:
    kind: CivoIpAddress
    ref: prod-lb-ip
```

## Configuration Reference

### Load Balancer Name (`loadBalancerName`)

**Type**: `string` (required)  
**Constraints**: 1-64 characters, lowercase alphanumeric and hyphens only

A unique name for the load balancer within your Civo account.

**Example**:
```yaml
loadBalancerName: my-web-lb
```

### Region (`region`)

**Type**: `enum` (required)

The Civo region where the load balancer will be created.

**Available Regions**: `lon1`, `nyc1`, `fra1`, `phx1`

**Example**:
```yaml
region: lon1
```

### Network (`network`)

**Type**: `StringValueOrRef` (required)

The private network where the load balancer will be placed. This is critical for security.

**Literal Value**:
```yaml
network:
  value: "net-12345678"
```

**Reference to CivoVpc Resource**:
```yaml
network:
  kind: CivoVpc
  ref: my-network
```

### Forwarding Rules (`forwardingRules`)

**Type**: `list[CivoLoadBalancerForwardingRule]` (required, min: 1)

Defines how traffic is routed from the load balancer to backend instances.

**Fields per rule**:
- `entryPort` (1-65535): Port on the load balancer that receives traffic
- `entryProtocol`: Protocol for incoming traffic (`http`, `https`, `tcp`)
- `targetPort` (1-65535): Port on backend instances
- `targetProtocol`: Protocol for backend traffic (`http`, `https`, `tcp`)

**Examples**:

HTTP to HTTP:
```yaml
forwardingRules:
  - entryPort: 80
    entryProtocol: http
    targetPort: 8080
    targetProtocol: http
```

HTTPS Passthrough:
```yaml
forwardingRules:
  - entryPort: 443
    entryProtocol: https
    targetPort: 8443
    targetProtocol: https
```

TCP Database:
```yaml
forwardingRules:
  - entryPort: 3306
    entryProtocol: tcp
    targetPort: 3306
    targetProtocol: tcp
```

### Health Check (`healthCheck`)

**Type**: `CivoLoadBalancerHealthCheck` (optional)

Defines how the load balancer monitors backend instance health.

**Fields**:
- `port` (1-65535, required): Port to probe on instances
- `protocol` (required): `http`, `https`, or `tcp`
- `path` (optional): HTTP/HTTPS request path (e.g., `/healthz`)

**Example**:
```yaml
healthCheck:
  port: 8080
  protocol: http
  path: /healthz
```

For TCP services (no path needed):
```yaml
healthCheck:
  port: 3306
  protocol: tcp
```

### Instance Selection (Mutually Exclusive)

You must specify EITHER `instanceIds` OR `instance_tag`, not both.

#### Instance IDs (`instanceIds`)

**Type**: `list[StringValueOrRef]` (optional)

Explicitly list instances to attach to the load balancer.

**Literal Values**:
```yaml
instanceIds:
  - value: "inst-12345"
  - value: "inst-67890"
```

**References to CivoComputeInstance Resources**:
```yaml
instanceIds:
  - kind: CivoComputeInstance
    ref: web-1
  - kind: CivoComputeInstance
    ref: web-2
```

#### Instance Tag (`instance_tag`)

**Type**: `string` (optional, max 255 chars)

All instances with this tag automatically join the load balancer pool. Ideal for autoscaling.

**Example**:
```yaml
instance_tag: web-server
```

### Reserved IP (`reservedIpId`)

**Type**: `StringValueOrRef` (optional)

Assign a reserved IP for a stable public endpoint that persists across load balancer recreation.

**Literal Value**:
```yaml
reservedIpId:
  value: "ip-12345678"
```

**Reference to CivoIpAddress Resource**:
```yaml
reservedIpId:
  kind: CivoIpAddress
  ref: prod-lb-ip
```

### Sticky Sessions (`enableStickySessions`)

**Type**: `bool` (optional, default: false)

Enables session affinity using source IP hashing. Requests from the same client IP route to the same backend instance.

**Example**:
```yaml
enableStickySessions: true
```

**Note**: Only enable for stateful applications. Prefer stateless design with external session storage (Redis, etc.).

## Best Practices

### 1. Always Use Private Networks

Deploy load balancers in a dedicated Civo network for security:

```yaml
network:
  kind: CivoVpc
  ref: private-network
```

### 2. Use Reserved IPs for Production

Allocate a reserved IP to prevent DNS changes during maintenance:

```yaml
reservedIpId:
  kind: CivoIpAddress
  ref: prod-lb-ip
```

### 3. Implement Real Health Checks

Create a meaningful health endpoint:

```yaml
healthCheck:
  port: 8080
  protocol: http
  path: /healthz  # Should return 200 only when truly healthy
```

### 4. Choose the Right Instance Attachment Strategy

- **Use instance IDs** for fixed, stable pools
- **Use tags** for dynamic, autoscaling pools

### 5. Multiple Forwarding Rules for Multi-Port Services

```yaml
forwardingRules:
  - entryPort: 80
    entryProtocol: http
    targetPort: 8080
    targetProtocol: http
  - entryPort: 443
    entryProtocol: https
    targetPort: 8443
    targetProtocol: https
```

## Troubleshooting

### Load Balancer Not Routing Traffic

**Symptoms**: Instances are attached but traffic doesn't reach them.

**Solutions**:
1. Verify instances are in the same network as the load balancer
2. Check health check status - instances must pass health checks
3. Verify forwarding rule ports match your application ports
4. Ensure firewall rules allow traffic on the target ports

### Health Checks Failing

**Symptoms**: Instances constantly marked as unhealthy.

**Solutions**:
1. Verify the health check path returns HTTP 200
2. Check that the health check port matches your application port
3. Ensure instances can receive traffic on the health check port
4. Test health endpoint manually: `curl http://<instance-ip>:<port>/healthz`

### Cannot Attach Instances

**Symptoms**: Error when specifying both `instanceIds` and `instance_tag`.

**Solution**: These are mutually exclusive. Use one or the other, not both.

## Additional Resources

- **Examples**: See [`examples.md`](examples.md) for more configuration scenarios
- **Research Documentation**: See [`docs/README.md`](docs/README.md) for architectural patterns
- **Pulumi Module**: See [`iac/pulumi/README.md`](iac/pulumi/README.md)
- **Terraform Module**: See [`iac/tf/README.md`](iac/tf/README.md)

## Support

For questions or issues:
- Project Planton: [GitHub Repository](https://github.com/project-planton/project-planton)
- Civo Documentation: [civo.com/docs](https://www.civo.com/docs)
- Civo Support: [civo.com/support](https://www.civo.com/support)

