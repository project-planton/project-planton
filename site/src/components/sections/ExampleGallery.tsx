"use client";
import React, { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Copy, Check } from "lucide-react";

export default function ExampleGallery() {
  const [copiedExamples, setCopiedExamples] = useState<Record<string, boolean>>({});

  const copyCode = (exampleId: string, code: string) => {
    navigator.clipboard.writeText(code);
    setCopiedExamples((prev) => ({ ...prev, [exampleId]: true }));
    setTimeout(() => {
      setCopiedExamples((prev) => ({ ...prev, [exampleId]: false }));
    }, 2000);
  };

  const examples = {
    aws: [
      {
        id: "aws-rds",
        filename: "aws-rds-instance.yaml",
        title: "RDS PostgreSQL Instance",
        manifest: `apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: payments-db
spec:
  subnet_ids:
    - valueFrom:
        kind: AwsVpc
        name: main-vpc
        fieldPath: status.outputs.private_subnets.[0].id
    - valueFrom:
        kind: AwsVpc
        name: main-vpc
        fieldPath: status.outputs.private_subnets.[1].id
  engine: postgres
  engine_version: "15.4"
  instance_class: db.t3.medium
  allocated_storage_gb: 100
  storage_encrypted: true
  username: dbadmin
  password: <secret>
  port: 5432
  multi_az: true`,
        deploy: `project-planton validate aws-rds-instance.yaml
project-planton pulumi up --manifest aws-rds-instance.yaml --stack myorg/database/prod`,
      },
      {
        id: "aws-ec2",
        filename: "aws-ec2.yaml",
        title: "EC2 Instance with References",
        manifest: `apiVersion: aws.project-planton.org/v1
kind: AwsEc2Instance
metadata:
  name: app-vm
spec:
  subnet_id:
    valueFrom:
      kind: AwsSubnet
      name: my-vpc-subnet
      fieldPath: status.outputs.subnet_id
  security_group_ids:
    - valueFrom:
        kind: AwsSecurityGroup
        name: app-sg
        fieldPath: status.outputs.security_group_id
  ami_id: ami-0123456789abcdef0
  instance_type: t3.micro
  connection_method: SSM`,
        deploy: `project-planton tofu plan --manifest aws-ec2.yaml
project-planton tofu apply --manifest aws-ec2.yaml --auto-approve`,
      },
    ],
    gcp: [
      {
        id: "gcp-gke",
        filename: "gcp-gke-cluster.yaml",
        title: "GKE Cluster with Autoscaling",
        manifest: `apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: main-gke
spec:
  cluster_project_id: <project-id>
  region: us-central1
  zone: us-central1-a
  is_workload_logs_enabled: false
  cluster_autoscaling_config:
    is_enabled: true
    cpu_min_cores: 4
    cpu_max_cores: 32
    memory_min_gb: 16
    memory_max_gb: 128
  node_pools:
    - name: general-pool
      machine_type: n2-standard-8
      min_node_count: 1
      max_node_count: 5
      is_spot_enabled: false`,
        deploy: `project-planton pulumi preview --manifest gcp-gke-cluster.yaml --stack myorg/platform/prod
project-planton pulumi up --manifest gcp-gke-cluster.yaml --stack myorg/platform/prod`,
      },
      {
        id: "gcp-cloudrun",
        filename: "gcp-cloud-run.yaml",
        title: "Cloud Run Service",
        manifest: `apiVersion: gcp.project-planton.org/v1
kind: GcpCloudRun
metadata:
  name: hello-run
spec:
  project_id: <project-id>
  region: us-central1
  service:
    name: hello
    image: us-docker.pkg.dev/cloudrun/container/hello
    allow_unauthenticated: true`,
        deploy: `project-planton tofu apply --manifest gcp-cloud-run.yaml --auto-approve`,
      },
    ],
    azure: [
      {
        id: "azure-aks",
        filename: "azure-aks.yaml",
        title: "AKS Cluster",
        manifest: `apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: ops-aks
spec:
  subscription_id: <subscription>
  resource_group_name: rg-ops
  region: eastus
  node_pools:
    - name: system
      vm_size: Standard_DS2_v2
      min_node_count: 1
      max_node_count: 3`,
        deploy: `project-planton pulumi up --manifest azure-aks.yaml --stack myorg/azure/dev`,
      },
      {
        id: "azure-acr",
        filename: "azure-acr.yaml",
        title: "Container Registry",
        manifest: `apiVersion: azure.project-planton.org/v1
kind: AzureContainerRegistry
metadata:
  name: app-registry
spec:
  subscription_id: <subscription>
  resource_group_name: rg-ops
  region: eastus
  sku: Basic`,
        deploy: `project-planton tofu apply --manifest azure-acr.yaml --auto-approve`,
      },
    ],
    kubernetes: [
      {
        id: "redis-k8s",
        filename: "redis-kubernetes.yaml",
        title: "Redis on Kubernetes",
        manifest: `apiVersion: kubernetes.project-planton.org/v1
kind: RedisKubernetes
metadata:
  name: session-cache
spec:
  container:
    replicas: 1
    resources:
      limits:
        cpu: "1000m"
        memory: "2Gi"
      requests:
        cpu: "50m"
        memory: "100Mi"
    persistence_enabled: true
    disk_size: "5Gi"`,
        deploy: `project-planton pulumi preview --manifest redis-kubernetes.yaml --stack myorg/cache/prod
project-planton pulumi up --manifest redis-kubernetes.yaml --stack myorg/cache/prod`,
      },
      {
        id: "postgres-k8s",
        filename: "postgres-kubernetes.yaml",
        title: "PostgreSQL on Kubernetes",
        manifest: `apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: app-db
spec:
  container:
    replicas: 3
    resources:
      limits:
        cpu: "2000m"
        memory: "4Gi"
      requests:
        cpu: "100m"
        memory: "256Mi"
    disk_size: "100Gi"
  ingress:
    enabled: true
    hostname: db.example.com`,
        deploy: `project-planton tofu plan --manifest postgres-kubernetes.yaml
project-planton tofu apply --manifest postgres-kubernetes.yaml --auto-approve`,
      },
    ],
  } as const;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            Example Gallery
          </span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto">
          Real manifests from the repo (trimmed for clarity). Replace placeholders like &lt;project-id&gt;.
        </p>
      </div>

      <Tabs defaultValue="aws" className="w-full">
        <TabsList className="grid grid-cols-4 w-full max-w-md mx-auto mb-8 bg-slate-800 border-slate-700">
          <TabsTrigger value="aws" className="data-[state=active]:bg-slate-700">AWS</TabsTrigger>
          <TabsTrigger value="gcp" className="data-[state=active]:bg-slate-700">GCP</TabsTrigger>
          <TabsTrigger value="azure" className="data-[state=active]:bg-slate-700">Azure</TabsTrigger>
          <TabsTrigger value="kubernetes" className="data-[state=active]:bg-slate-700">Kubernetes</TabsTrigger>
        </TabsList>

        {Object.entries(examples).map(([provider, providerExamples]) => (
          <TabsContent key={provider} value={provider} className="space-y-8">
            <div className="grid lg:grid-cols-2 gap-8">
              {providerExamples.map((example) => (
                <Card key={example.id} className="bg-slate-900/30 border-slate-700">
                  <CardHeader className="pb-4">
                    <div className="flex items-center justify-between">
                      <div>
                        <CardTitle className="text-lg text-white">{example.title}</CardTitle>
                        <Badge variant="outline" className="border-slate-600 text-slate-400 mt-2">
                          {example.filename}
                        </Badge>
                      </div>
                      <Button
                        size="sm"
                        variant="outline"
                        onClick={() => copyCode(example.id, example.manifest)}
                        className="border-slate-600 bg-slate-800/50 text-slate-300 hover:text-white hover:bg-slate-700"
                      >
                        {copiedExamples[example.id] ? (
                          <>
                            <Check className="w-3 h-3 mr-1" />
                            Copied
                          </>
                        ) : (
                          <>
                            <Copy className="w-3 h-3 mr-1" />
                            Copy
                          </>
                        )}
                      </Button>
                    </div>
                  </CardHeader>
                  <CardContent className="space-y-4">
                    <div className="bg-slate-950 rounded-lg p-4 font-mono text-xs overflow-x-auto max-w-full">
                      <pre className="text-slate-300 whitespace-pre-wrap break-words w-full">
                        <code className="block">{example.manifest}</code>
                      </pre>
                    </div>
                    <div>
                      <h4 className="text-sm font-bold text-white mb-2">Deploy:</h4>
                      <div className="bg-slate-900 rounded-lg p-3 font-mono text-xs overflow-x-auto max-w-full">
                        <pre className="text-green-400 whitespace-pre-wrap break-words w-full">
                          <code className="block">{example.deploy}</code>
                        </pre>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </TabsContent>
        ))}
      </Tabs>
    </div>
  );
}


