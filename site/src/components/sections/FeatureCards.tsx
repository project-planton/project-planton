import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Globe, Shield, Code, Zap, Lock, Puzzle } from "lucide-react";

const features = [
  {
    icon: Globe,
    title: "One model, many clouds",
    description:
      "A single CloudResourceKind enum and provider‑specific APIs (e.g., AwsStaticWebsite, GcpGkeCluster, AzureAksCluster, RedisKubernetes) with the same KRM‑style shape: apiVersion, kind, metadata, spec, status.",
  },
  {
    icon: Shield,
    title: "Validation & sane defaults",
    description:
      "Field‑level rules via Buf ProtoValidate; default/recommended_default options and message‑level CEL guardrails; manifests fail fast with clear errors before any provisioning.",
  },
  {
    icon: Code,
    title: "Modules under the hood",
    description:
      "Curated Pulumi and OpenTofu modules reside alongside each API (apis/.../iac/pulumi/module, apis/.../iac/tf), so you never write raw Terraform or Pulumi code.",
  },
  {
    icon: Zap,
    title: "Dev‑grade workflow",
    description:
      "pulumi preview/update/refresh/destroy and tofu init/plan/apply/destroy/refresh through one CLI; stack FQDN <org>/<project>/<stack>; reproducible workspaces under ~/.project-planton.",
  },
  {
    icon: Lock,
    title: "Security & governance",
    description:
      "Provider credentials are first‑class stack‑inputs (e.g., aws-credential.yaml, gcp-credential.yaml, kubernetes-cluster.yaml); modules apply consistent planton.org/* labels and support foreign‑key references across resources.",
  },
  {
    icon: Puzzle,
    title: "Extensibility",
    description:
      "Add kinds and modules in‑repo; APIs generate language stubs via Buf; modules are selected per kind/provider automatically; override any manifest value at the CLI with --set key=value.",
  },
];

export default function FeatureCards() {
  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            Why ProjectPlanton?
          </span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto">
          Built for DevOps engineers, platform teams, and infrastructure architects who need consistency across clouds
        </p>
      </div>
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {features.map((feature, index) => {
          const IconComponent = feature.icon;
          return (
            <Card
              key={index}
              className="bg-slate-900/50 border-slate-700 hover:bg-slate-900/70 transition-all duration-300 group"
            >
              <CardContent className="p-8">
                <div className="flex items-center gap-4 mb-4">
                  <div className="w-12 h-12 bg-gradient-to-r from-blue-500 to-cyan-400 rounded-xl flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                    <IconComponent className="w-6 h-6 text-white" />
                  </div>
                  <h3 className="text-xl font-bold text-white">{feature.title}</h3>
                </div>
                <p className="text-slate-400 leading-relaxed">{feature.description}</p>
              </CardContent>
            </Card>
          );
        })}
      </div>
    </div>
  );
}


