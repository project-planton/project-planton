import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Check, X, Minus } from 'lucide-react';

export default function CompareSection() {
  const comparisons = [
    {
      tool: 'ProjectPlanton',
      execution: 'On-demand CLI',
      syntax: 'KRM-style YAML',
      validation: 'Protobuf + Buf ProtoValidate',
      engines: 'Pulumi + OpenTofu',
      advantages: ['Declarative manifests', 'Strong validation', 'Multi-engine support']
    },
    {
      tool: 'Crossplane',
      execution: 'In-cluster controllers',
      syntax: 'KRM YAML',
      validation: 'OpenAPI schemas',
      engines: 'Custom providers',
      advantages: ['Kubernetes-native', 'GitOps ready', 'Continuous reconciliation']
    },
    {
      tool: 'Terraform HCL',
      execution: 'CLI-driven',
      syntax: 'HCL (imperative)',
      validation: 'Variable validation',
      engines: 'Terraform only',
      advantages: ['Mature ecosystem', 'Provider coverage', 'State management']
    },
    {
      tool: 'Pulumi Code',
      execution: 'CLI + Programs',
      syntax: 'TypeScript/Python/Go',
      validation: 'Language-specific',
      engines: 'Pulumi only',
      advantages: ['Programming languages', 'IDE support', 'Testing frameworks']
    }
  ];

  const features = [
    { name: 'Declarative YAML', projectplanton: true, crossplane: true, terraform: false, pulumi: false },
    { name: 'Multi-cloud consistency', projectplanton: true, crossplane: true, terraform: false, pulumi: false },
    { name: 'Strong validation', projectplanton: true, crossplane: false, terraform: false, pulumi: false },
    { name: 'Multiple backends', projectplanton: true, crossplane: false, terraform: true, pulumi: true },
    { name: 'CLI workflow', projectplanton: true, crossplane: false, terraform: true, pulumi: true },
    { name: 'Kubernetes native', projectplanton: false, crossplane: true, terraform: false, pulumi: false },
    { name: 'Programming languages', projectplanton: false, crossplane: false, terraform: false, pulumi: true },
    { name: 'Curated modules', projectplanton: true, crossplane: false, terraform: false, pulumi: false }
  ];

  const FeatureIcon = ({ enabled }) => {
    if (enabled === true) return <Check className="w-4 h-4 text-green-400" />;
    if (enabled === false) return <X className="w-4 h-4 text-red-400" />;
    return <Minus className="w-4 h-4 text-slate-500" />;
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            How We Compare
          </span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto">
          Short, factual comparison with other infrastructure tools
        </p>
      </div>

      {/* Quick comparison cards */}
      <div className="grid lg:grid-cols-3 gap-6 mb-12">
        <Card className="bg-slate-900/30 border-slate-700">
          <CardHeader>
            <CardTitle className="text-lg text-white">vs Crossplane</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-slate-400">
              <strong className="text-blue-400">Crossplane:</strong> In‑cluster controllers
            </p>
            <p className="text-slate-400 mt-2">
              <strong className="text-emerald-400">ProjectPlanton:</strong> Runs on‑demand via CLI with Pulumi/OpenTofu modules
            </p>
          </CardContent>
        </Card>

        <Card className="bg-slate-900/30 border-slate-700">
          <CardHeader>
            <CardTitle className="text-lg text-white">vs Terraform HCL</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-slate-400">
              <strong className="text-blue-400">Terraform HCL:</strong> Imperative DSL
            </p>
            <p className="text-slate-400 mt-2">
              <strong className="text-emerald-400">ProjectPlanton:</strong> Declarative KRM‑style YAML with strict proto validations
            </p>
          </CardContent>
        </Card>

        <Card className="bg-slate-900/30 border-slate-700">
          <CardHeader>
            <CardTitle className="text-lg text-white">vs Pulumi Code</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-slate-400">
              <strong className="text-blue-400">Pulumi code:</strong> App‑like programs
            </p>
            <p className="text-slate-400 mt-2">
              <strong className="text-emerald-400">ProjectPlanton:</strong> Declarative manifests while Pulumi executes curated modules
            </p>
          </CardContent>
        </Card>
      </div>

      {/* Feature comparison table */}
      <Card className="bg-slate-900/30 border-slate-700">
        <CardHeader>
          <CardTitle className="text-xl text-white">Feature Comparison</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-x-auto">
            <Table>
              <TableHeader>
                <TableRow className="border-slate-800">
                  <TableHead className="text-slate-300">Feature</TableHead>
                  <TableHead className="text-slate-300">ProjectPlanton</TableHead>
                  <TableHead className="text-slate-300">Crossplane</TableHead>
                  <TableHead className="text-slate-300">Terraform</TableHead>
                  <TableHead className="text-slate-300">Pulumi</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {features.map((feature, index) => (
                  <TableRow key={index} className="border-slate-800">
                    <TableCell className="font-medium text-white">{feature.name}</TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <FeatureIcon enabled={feature.projectplanton} />
                        {feature.projectplanton && <Badge className="bg-emerald-900 text-emerald-200">Yes</Badge>}
                      </div>
                    </TableCell>
                    <TableCell>
                      <FeatureIcon enabled={feature.crossplane} />
                    </TableCell>
                    <TableCell>
                      <FeatureIcon enabled={feature.terraform} />
                    </TableCell>
                    <TableCell>
                      <FeatureIcon enabled={feature.pulumi} />
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}