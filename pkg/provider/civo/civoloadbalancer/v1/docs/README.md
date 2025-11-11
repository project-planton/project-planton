# Deploying Civo Load Balancers: Simplicity Meets Production Readiness

## The Developer-First Cloud Paradox

For years, cloud load balancers followed a familiar pattern: start simple, then layer on so much enterprise complexity that deploying one requires navigating a maze of policies, target groups, listener rules, and cross-zone configurations. AWS's Application Load Balancer, while powerful, exemplifies this—complete with path-based routing, host-based routing, WebSocket support, and WAF integration. But here's the question: what if 80% of workloads only need 20% of those features?

Civo's Load Balancer represents a different philosophy: **purposeful simplicity**. It's a regional load balancer that distributes traffic across compute instances or Kubernetes clusters using straightforward forwarding rules, health checks, and session affinity. No complex routing tables. No overwhelming feature sets. Just the essentials, delivered with predictable flat-rate pricing (~$10/month for 10,000 concurrent connections with unlimited bandwidth).

This simplicity isn't a limitation—it's a strategic choice. Civo LBs handle HTTP, HTTPS (via TLS passthrough), and TCP traffic with robust health checking and multiple load balancing algorithms (round-robin, least-connections). They integrate natively with Civo's Kubernetes clusters through a cloud controller manager, automatically provisioning load balancers when you create a Service of type LoadBalancer. For the vast majority of web applications, API services, and TCP-based workloads, this covers everything you need.

This document explores how to deploy Civo Load Balancers, from manual console clicks to production-grade Infrastructure as Code. We'll examine the maturity spectrum of deployment approaches, compare the leading IaC tools, and explain why Project Planton abstracts these details into a clean, multi-cloud API that focuses on the 20% of configuration that 80% of users actually need.

## The Deployment Maturity Spectrum

### Level 0: Manual Console Management

**The Approach:** Civo's web dashboard provides a point-and-click interface for creating load balancers. Navigate to "Networking > Load Balancers," fill in a form (name, region, network, algorithm, instance pool, health check configuration), and click create. The load balancer provisions in seconds with an assigned public IP and DNS name (`*.lb.civo.com`).

**What It Solves:** Perfect for learning, prototyping, or one-off services. The UI visualizes your configuration clearly: which instances are attached, what ports are being forwarded, whether health checks are passing. For teams just getting started with Civo or validating a concept, the console is the fastest path from idea to working load balancer.

**What It Doesn't Solve:** 

- **Zero Reproducibility:** Every environment (dev, staging, prod) requires manual recreation. No versioning, no audit trail.
- **Configuration Drift:** Two engineers creating "identical" load balancers will make subtly different choices. Six months later, no one remembers why staging uses round-robin while prod uses least-connections.
- **No Integration:** Manual processes don't fit into CI/CD pipelines. You can't programmatically update a load balancer when deploying new instances.
- **Human Error:** Forgetting to configure the health check path, selecting the wrong network, or mixing up source/target ports happens more often than we'd like to admit.

**Verdict:** Use the console for exploration and debugging. For anything that needs to exist beyond next week, move to code.

### Level 1: CLI Scripting and Automation

**The Approach:** The Civo CLI (`civo`) wraps the API in a command-line interface. You can script load balancer creation with commands like:

```bash
civo loadbalancer create my-lb \
  --network_ID default \
  --algorithm round_robin \
  --create-firewall "80,443" \
  --backend ip=10.0.0.10,source-port=80,target-port=80,protocol=http
```

The CLI supports all core operations: create, list, update, remove. You can embed these commands in deployment scripts, CI/CD jobs, or Ansible playbooks.

**What It Solves:** 

- **Automation Begins:** You can script the creation of consistent load balancers across environments. Store scripts in version control for basic reproducibility.
- **Pipeline Integration:** Shell scripts can call `civo` commands during deployment, enabling semi-automated infrastructure updates.
- **Speed:** For experienced engineers, typing a CLI command is often faster than clicking through a UI.

**What It Doesn't Solve:**

- **Imperative, Not Declarative:** Scripts describe *how* to create resources, not *what* the end state should be. If the script runs twice, you get errors or duplicates. There's no concept of convergence toward a desired state.
- **No State Management:** The CLI doesn't track what exists. You must manually check if a load balancer already exists before creating it, then handle updates vs. creates differently.
- **Limited Tag Support:** The CLI expects IP addresses for backends rather than instance tags, making dynamic instance management awkward.
- **Fragile Error Handling:** If a step fails midway through a complex script, rolling back or recovering requires custom logic.

**Verdict:** CLI scripting is a stepping stone. It's useful for one-off automation tasks or when you need quick programmatic control, but it lacks the robustness required for production infrastructure management. Think of it as glue code, not a foundation.

### Level 2: Infrastructure as Code (Terraform and Pulumi)

**The Approach:** Treat your load balancer as declarative code using tools like Terraform or Pulumi. Define the desired state—network, instances, forwarding rules, health checks—and let the tool figure out how to create, update, or delete resources to match that state.

**Terraform Example:**

```hcl
resource "civo_loadbalancer" "web_lb" {
  name         = "web-lb"
  network_id   = var.network_id
  algorithm    = "round_robin"
  
  backend {
    instance_id = civo_instance.web1.id
    protocol    = "http"
    port        = 80
  }
  
  health_check_path = "/healthz"
  enable_sticky_sessions = false
}
```

**Pulumi Example (TypeScript):**

```typescript
const lb = new civo.LoadBalancer("webLb", {
    name: "web-lb",
    networkId: network.id,
    algorithm: "round_robin",
    backends: [
        { instanceId: web1.id, protocol: "http", port: 80 },
        { instanceId: web2.id, protocol: "http", port: 80 }
    ],
    healthCheckPath: "/healthz",
    enableStickySessions: false
});
```

Both tools maintain state files that track what exists, enabling smart updates. Change the health check path in your code, run `terraform apply` or `pulumi up`, and only the necessary API calls are made.

**What It Solves:**

- **Declarative State:** You describe *what* you want, not *how* to build it. The tool reconciles reality with your code.
- **Idempotency:** Run the same code 100 times, get the same result. No duplicates, no errors from "already exists."
- **Versioning and Collaboration:** Infrastructure lives in Git alongside application code. Pull requests, code reviews, and rollbacks become standard practice.
- **Dependency Management:** Terraform and Pulumi understand relationships. They create instances before attaching them to the load balancer. They can reference reserved IPs, networks, and firewalls by resource handles rather than hardcoded IDs.
- **Multi-Environment Patterns:** Use workspaces, stack parameters, or separate state files to maintain identical infrastructure across dev/staging/prod with different values.

**What It Doesn't Solve (Compared to Higher Abstraction):**

- **Cloud-Specific Syntax:** Terraform HCL for Civo looks different from Terraform HCL for AWS. You're still writing provider-specific code.
- **Boilerplate:** Every project needs provider setup, variable definitions, output declarations. Small teams copy-paste patterns, leading to subtle inconsistencies.
- **Limited Validation:** While Pulumi offers type checking and Terraform has some validation, neither deeply understands *business logic*—like ensuring a reserved IP is actually assigned in production.

**Verdict:** IaC is the minimum bar for production infrastructure. It brings repeatability, version control, and collaboration. Terraform is mature, widely adopted, and works well for teams already invested in HCL. Pulumi offers flexibility for teams that prefer real programming languages and need complex conditional logic. Both are viable; choose based on your team's existing skills and tooling.

### Level 3: Multi-Cloud Abstraction (Project Planton)

**The Approach:** Instead of writing cloud-specific Terraform or Pulumi, define infrastructure using high-level, cloud-agnostic protobuf APIs. Project Planton translates these abstract definitions into provider-specific IaC under the hood (using Terraform or Pulumi modules).

**Example: CivoLoadBalancerSpec (Protobuf/YAML):**

```yaml
apiVersion: civo.project-planton.org/v1
kind: CivoLoadBalancer
metadata:
  name: web-lb
spec:
  loadBalancerName: web-lb
  region: LON1
  network:
    kind: CivoVpc
    ref: my-network
  forwardingRules:
    - entryPort: 80
      entryProtocol: http
      targetPort: 80
      targetProtocol: http
  healthCheck:
    port: 80
    protocol: http
    path: /healthz
  instanceIds:
    - kind: CivoComputeInstance
      ref: web1
    - kind: CivoComputeInstance
      ref: web2
  enableStickySessions: false
```

**What It Solves:**

- **Cloud Portability:** The same conceptual model works across clouds. Define a load balancer on Civo with this syntax; define one on AWS with a structurally similar spec. The platform handles the translation.
- **80/20 API Design:** Project Planton exposes only the fields that matter for most use cases: region, network, forwarding rules, health checks, instance selectors, sticky sessions. Esoteric settings (connection limits, proxy protocol, custom timeouts) are handled by sensible defaults or advanced annotations.
- **Strong Typing and Validation:** Protobuf schemas enforce field types, required fields, and constraints at design time. You can't accidentally create a load balancer with an invalid port range or a missing health check.
- **Declarative Dependencies:** Use foreign key references (`kind: CivoVpc, ref: my-network`) instead of hardcoded IDs. The platform resolves these dynamically, ensuring instances, networks, and IPs are created in the correct order.
- **Best Practices Baked In:** Reserved IPs for stable endpoints, private networks for security, tag-based instance selection for autoscaling—these patterns are naturally expressed in the API rather than left to engineer judgment.
- **Focus on Intent, Not Mechanics:** Engineers declare *what* infrastructure they need, not *how* to provision it. This reduces cognitive load and allows infrastructure teams to evolve the underlying implementation (switch from Terraform to Pulumi, upgrade provider versions) without changing user-facing specs.

**What It Doesn't Solve:**

- **Escape Hatches for Edge Cases:** If you need a feature Civo supports but Planton's API doesn't expose (like custom connection timeouts), you might need to drop down to raw Terraform or submit a feature request.
- **Learning Curve:** Teams already fluent in Terraform HCL need to learn a new abstraction. The payoff is portability and simplicity, but there's an upfront investment.

**Verdict:** For organizations building multi-cloud platforms or managing infrastructure at scale, Project Planton's abstraction layer is the production-grade solution. It delivers the benefits of IaC (versioning, collaboration, idempotency) while raising the level of abstraction to intent-based infrastructure. You're not just coding infrastructure; you're defining the *shape* of your architecture in a way that's portable, validated, and maintainable.

## IaC Tool Deep Dive: Terraform vs. Pulumi for Civo

Both Terraform and Pulumi can manage Civo Load Balancers via the official Civo provider. Understanding their trade-offs helps choose the right foundation—whether you're using them directly or relying on Project Planton's abstraction (which can leverage either under the hood).

### Provider Maturity and Coverage

**Terraform:** The Civo Terraform provider is officially maintained by Civo and covers all major resources: instances, networks, firewalls, DNS, Kubernetes clusters, and load balancers. The load balancer resource supports:

- Multiple backends (by instance ID)
- Forwarding rules with separate entry/target ports and protocols
- Health check configuration (port, path, protocol)
- Sticky sessions (via `enable_sticky_sessions` or algorithm `ip_hash`)
- Network and firewall attachment
- Reserved IP assignment

The provider is stable and production-ready. However, note that the load balancer resource underwent refactoring in 2023 to support multiple forwarding rules (previously limited to one port per LB). Current versions fully support multi-port configurations.

**Pulumi:** Pulumi's Civo provider is a Terraform bridge—it wraps the Terraform provider in Pulumi's SDK. This means feature parity is guaranteed: anything Terraform can do, Pulumi can do. You get the same resource coverage with the added benefit of using real programming languages (TypeScript, Python, Go, C#).

### Configuration Syntax and Developer Experience

**Terraform (HCL):**

```hcl
resource "civo_reserved_ip" "lb_ip" {
  name   = "web-lb-ip"
  region = "LON1"
}

resource "civo_loadbalancer" "web_lb" {
  name              = "web-lb"
  network_id        = var.network_id
  reserved_ip_id    = civo_reserved_ip.lb_ip.id
  algorithm         = "round_robin"
  
  backend {
    instance_id = civo_instance.web1.id
    protocol    = "http"
    port        = 80
  }
  
  health_check_path = "/healthz"
}
```

HCL is declarative and concise for straightforward scenarios. Terraform's strength is its maturity, extensive community examples, and robust state management. The syntax is purpose-built for infrastructure, which makes it readable even for those unfamiliar with general-purpose programming.

**Pulumi (TypeScript):**

```typescript
const reservedIp = new civo.ReservedIp("lb-ip", {
    name: "web-lb-ip",
    region: "LON1"
});

const lb = new civo.LoadBalancer("webLb", {
    name: "web-lb",
    networkId: network.id,
    reservedIpId: reservedIp.id,
    algorithm: "round_robin",
    backends: instances.map(inst => ({
        instanceId: inst.id,
        protocol: "http",
        port: 80
    })),
    healthCheckPath: "/healthz"
});
```

Pulumi lets you use loops, conditionals, functions, and libraries from your language ecosystem. If you need to dynamically attach instances based on a naming convention or environment variable, Pulumi makes this trivial. The trade-off is increased complexity for simple cases—sometimes you just want declarative HCL without thinking about imperative logic.

### State Management and Multi-Environment Patterns

**Terraform:** Uses workspace-based or directory-based state separation. For dev/staging/prod, you either:

- Create separate workspaces (`terraform workspace select prod`)
- Maintain separate directories with distinct state files

Both patterns work. Terraform Cloud and Terraform Enterprise offer remote state backends with locking, versioning, and collaboration features.

**Pulumi:** Uses "stacks" as first-class multi-environment primitives. Each stack has its own config file and state. You define stack-specific variables in `Pulumi.<stack>.yaml` and reference them in code:

```typescript
const config = new pulumi.Config();
const region = config.require("region"); // Different per stack
```

Pulumi's state backend is remote by default (Pulumi Cloud, self-hosted, or S3/Azure/GCS). This eliminates a common Terraform pitfall: forgetting to configure remote state and losing it on a local machine.

### Best Practices for Reserved IPs and Immutable Infrastructure

**Always Use Reserved IPs in Production:** Both tools support reserved IP management. A load balancer without a reserved IP gets an ephemeral public IP that changes if the LB is destroyed and recreated. With a reserved IP:

- DNS records remain stable across infrastructure rebuilds
- You can perform blue-green deployments by switching reserved IPs between load balancers
- Planned maintenance doesn't require DNS propagation delays

**Terraform Pattern:**

```hcl
resource "civo_reserved_ip" "web_lb_ip" {
  name   = "web-lb-ip"
  region = var.region
}

resource "civo_loadbalancer" "web_lb" {
  reserved_ip_id = civo_reserved_ip.web_lb_ip.id
  # ... rest of config
}
```

If Terraform needs to replace the LB (due to certain config changes), it will recreate it with the same reserved IP, ensuring zero DNS downtime.

**Pulumi Pattern:**

```typescript
const reservedIp = new civo.ReservedIp("web-lb-ip", { 
    name: "web-lb-ip" 
});

const lb = new civo.LoadBalancer("webLb", {
    reservedIpId: reservedIp.id,
    // ... rest of config
});
```

Same guarantee: the reserved IP is a separate resource. Even if the load balancer resource is replaced, the IP persists and reassociates.

### Feature Gaps and Advanced Settings

Both tools expose most Civo LB features, but a few advanced settings might be missing or require workarounds:

- **Proxy Protocol:** Civo's Kubernetes Cloud Controller Manager supports enabling proxy protocol via annotations. In Terraform/Pulumi, check if `enable_proxy_protocol` exists as a field. If not, you may need to handle it via annotations on Kubernetes Services or accept the default.
- **Client/Server Timeouts:** Civo's API supports custom timeouts for long-lived connections. These may not be exposed in the Terraform provider's schema. If you need to tweak them, you might use the Civo API directly or file a provider enhancement request.
- **Max Concurrent Connections:** Default is 10,000. Scaling beyond requires contacting Civo support to increase the limit. This operational setting isn't typically in IaC specs.

For most production workloads, the current provider coverage is sufficient. Edge cases can often be solved by combining IaC with API calls or relying on sensible platform defaults.

### The Verdict: Terraform or Pulumi?

**Choose Terraform if:**

- Your team already uses Terraform for other clouds
- You value ecosystem maturity, extensive community modules, and a proven track record
- Declarative HCL is intuitive for your infrastructure engineers
- You prefer a single tool across all providers

**Choose Pulumi if:**

- Your team prefers writing infrastructure in familiar programming languages
- You need complex conditional logic, loops, or integration with external systems (APIs, databases)
- You want type safety and IDE autocompletion for infrastructure code
- You're building a platform that generates infrastructure dynamically

**Neither choice is wrong.** Both will reliably manage Civo Load Balancers. Pulumi offers more flexibility; Terraform offers more standardization. Project Planton can use either as its backend, allowing you to benefit from the abstraction layer regardless of which tool powers it.

## What Project Planton Supports and Why

Project Planton's `CivoLoadBalancer` API is built on the 80/20 principle: expose the 20% of configuration that 80% of users need, with sensible defaults for the rest. This reduces complexity without sacrificing production readiness.

### Core Supported Features

1. **Region and Network Placement**
   - Specify the Civo region (e.g., `LON1`, `NYC1`)
   - Reference the network (VPC) by kind and resource name, not hardcoded IDs
   - Ensures load balancers are placed in private networks for security

2. **Flexible Backend Selection**
   - **By Instance IDs:** Explicitly list compute instances to attach (supports foreign key references to `CivoComputeInstance` resources)
   - **By Tag:** Specify a tag; all instances with that tag dynamically join the pool
   - These are mutually exclusive, preventing configuration errors

3. **Forwarding Rules**
   - Define multiple rules for different ports and protocols
   - Each rule specifies: entry port, entry protocol, target port, target protocol
   - Supports HTTP, HTTPS (TLS passthrough), and TCP
   - Example: Listen on 80 (HTTP) → forward to 80 (HTTP); listen on 443 (TCP) → forward to 443 (TCP)

4. **Health Checks**
   - Configure protocol (HTTP, HTTPS, TCP), port, and path (for HTTP/S)
   - Example: `GET http://<instance>:80/healthz` every interval
   - Instances failing health checks are removed from rotation automatically

5. **Sticky Sessions**
   - Optional boolean flag to enable session affinity (disabled by default)
   - Uses source IP hashing for consistent client-to-instance routing
   - Recommended only for stateful applications; prefer stateless design

6. **Reserved IP Integration**
   - Reference a reserved IP resource for stable public endpoints
   - Prevents IP changes during load balancer recreation
   - Critical for production services with DNS records

### Design Decisions

**TLS Passthrough, Not Termination:** Civo Load Balancers natively support TLS passthrough—encrypted traffic flows end-to-end from clients to backend instances or Kubernetes ingress controllers. The load balancer does not decrypt traffic. This design choice:

- Preserves end-to-end encryption
- Allows certificate management via Let's Encrypt on instances or cert-manager in Kubernetes
- Avoids the complexity of uploading certificates to the LB

While some clouds (DigitalOcean, AWS) offer LB-based TLS termination, Civo's approach aligns with cloud-native best practices: manage certificates where your applications run, not in the network layer.

**No Firewall Configuration in LB Spec:** Firewalls are managed as separate resources. The load balancer references a network, and you attach firewalls to that network or create dedicated firewall rules. This separation of concerns keeps the LB spec focused on traffic distribution, not access control.

**Default Algorithm: Round-Robin:** Most workloads distribute load evenly across stateless instances, making round-robin the sensible default. For workloads with varying instance capacity or long-lived connections, you can specify `least_connections`. Advanced algorithms (random, custom weighted) are not exposed—they serve niche use cases better handled outside the core API.

**No Connection Limits in Spec:** The default 10,000 concurrent connections suffice for most services. Exceeding this requires contacting Civo to scale the LB (charged per 10k block). Since this is an operational scaling decision, not a design-time configuration, it's omitted from the API.

### Integration with Project Planton's Multi-Cloud Model

The same conceptual load balancer spec works across clouds:

- **On Civo:** Deploys a Civo Load Balancer
- **On AWS:** Could deploy an Application Load Balancer or Network Load Balancer (depending on protocol)
- **On GCP:** Could deploy a Regional TCP/UDP Load Balancer or HTTP(S) Load Balancer

Project Planton abstracts cloud-specific quirks. For example:

- **Forwarding Rules:** Mapped to Civo's backend definitions, AWS ALB listener rules, or GCP forwarding rules
- **Health Checks:** Translated to provider-specific formats (ALB target group health checks, GCP health check resources)
- **Reserved IPs:** Civo reserved IPs, AWS Elastic IPs, GCP static IPs

This abstraction allows teams to define infrastructure once and deploy anywhere, with confidence that production best practices (reserved IPs, private networks, health checks) are enforced by the platform.

## Production Deployment Checklist

When deploying Civo Load Balancers to production, ensure you've addressed these essentials:

### 1. Health Checks Are Properly Configured

- **Implement a real health endpoint:** `/healthz` should return 200 OK only when your application is truly healthy (database reachable, dependencies available)
- **Use HTTP health checks for HTTP services:** More accurate than TCP checks (which only verify the port is open)
- **Match the health check port to your service:** If your app serves on 8080, health check on 8080, not 80

### 2. Use Reserved IPs for Stability

- **Allocate a reserved IP per production load balancer:** Prevents IP changes during maintenance or infrastructure updates
- **Point DNS records at the reserved IP:** Use A records or CNAME to `*.lb.civo.com` hostnames
- **Tag reserved IPs clearly:** Name them after the service (e.g., `api-prod-lb-ip`) for easy identification

### 3. Place Load Balancers in Private Networks

- **Create dedicated VPCs for environments:** Separate dev, staging, and prod networks
- **Attach load balancers to the correct network:** Ensure instances and LB are on the same network
- **Restrict firewall rules:** Open only necessary ports (80, 443) and consider allowlisting source IPs for non-public services

### 4. Choose the Right Instance Attachment Strategy

- **Use instance IDs for stable, fixed pools:** When you know exactly which instances should receive traffic
- **Use tags for autoscaling or dynamic pools:** Tag instances with `role:webserver`, and the LB automatically includes them
- **Never mix both:** Either select by tag or by ID list, not both

### 5. Think Twice Before Enabling Sticky Sessions

- **Default: stateless applications with sticky sessions off:** Let the LB freely distribute traffic for optimal balance
- **Enable only when necessary:** For apps with in-memory sessions that can't be externalized to Redis or similar
- **Understand the trade-offs:** Sticky sessions can cause uneven load and reduce failover transparency

### 6. Plan for Multi-Region Availability

- Civo LBs are regional. For multi-region HA, use **DNS-based load balancing** or failover:
  - Deploy identical stacks in `LON1` and `NYC1`
  - Use geo-DNS to route users to the nearest region
  - Implement health-based failover to redirect traffic if one region goes down

### 7. Monitor and Alert

- **External uptime monitoring:** Use services like Pingdom or UptimeRobot to probe the LB's public IP
- **Backend instance monitoring:** Watch instance CPU, memory, and application metrics
- **Health check status:** If all instances fail health checks simultaneously, investigate immediately
- **Connection limits:** Monitor concurrent connections if you approach the 10k default; plan to scale if needed

### 8. Automate Certificate Management

- **For Kubernetes:** Deploy cert-manager to automatically provision and renew Let's Encrypt certificates
- **For VMs:** Use certbot with DNS validation or HTTP-01 challenges behind the LB
- **Wildcard certificates:** Consider a single wildcard cert for `*.example.com` distributed to all instances

## Conclusion: Infrastructure as Intent, Not Implementation

The evolution from manual console clicks to multi-cloud abstraction reflects a broader shift in how we think about infrastructure. We started by describing *how* to create resources—click here, run this command, execute these API calls. We've moved toward declaring *what* we need—a load balancer that distributes traffic across instances with health checks and sticky sessions.

Civo Load Balancers embody this philosophy at the product level: they strip away unnecessary complexity and deliver the 80% use case with clarity and performance. Project Planton takes it further: you define the *intent* (distribute traffic, check health, use a stable IP), and the platform handles the implementation details (whether that's Terraform modules, Pulumi programs, or API calls).

For production infrastructure, this is the path forward. Not because abstraction is trendy, but because it lets engineers focus on architecture instead of syntax, on resilience instead of boilerplate, and on delivering value instead of wrestling with cloud provider quirks.

Use the console to learn. Use the CLI to script. Use Terraform or Pulumi for production IaC. And when you're ready to scale across clouds and teams, use Project Planton to codify your infrastructure as portable, validated, intent-based specifications.

Your load balancer doesn't care whether you created it with a mouse click or a protobuf spec. But your team will care when they can deploy to three clouds with one YAML file, enforce best practices with schema validation, and sleep soundly knowing reserved IPs and health checks are guaranteed, not optional.

That's the promise of treating infrastructure as intent: you declare what you need, and the platform ensures it happens—correctly, repeatably, and without the toil.

