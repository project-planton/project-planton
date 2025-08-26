import React from 'react';
import { Card, CardContent } from "@/components/ui/card";
import { ArrowRight, FileText, Search, Package, Play, CheckCircle } from 'lucide-react';

export default function HowItWorks() {
  const steps = [
    {
      icon: FileText,
      title: "Parse",
      description: "apiVersion/kind/metadata/spec → Protobuf object"
    },
    {
      icon: Search,
      title: "Validate",
      description: "Buf ProtoValidate/CEL on spec"
    },
    {
      icon: Package,
      title: "Build",
      description: "stack‑input = { provisioner, pulumi|terraform, target, providerCredential }"
    },
    {
      icon: Play,
      title: "Plan/Preview",
      description: "pulumi preview or tofu plan"
    },
    {
      icon: CheckCircle,
      title: "Apply",
      description: "pulumi update or tofu apply (backends: local|s3|gcs|azurerm)"
    }
  ];

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            How It Works
          </span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto mb-12">
          From YAML manifest to deployed infrastructure in five clear steps
        </p>
      </div>

      {/* Workflow Diagram */}
      <div className="mb-16">
        <div className="flex flex-col lg:flex-row items-center justify-center gap-4 lg:gap-8">
          {steps.map((step, index) => {
            const Icon = step.icon;
            return (
              <div key={index} className="flex flex-col lg:flex-row items-center gap-4">
                <Card className="bg-slate-900/50 border-slate-700 w-full lg:w-64">
                  <CardContent className="p-6 text-center">
                    <div className="w-12 h-12 bg-gradient-to-r from-blue-500 to-cyan-400 rounded-xl flex items-center justify-center mx-auto mb-4">
                      <Icon className="w-6 h-6 text-white" />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">{step.title}</h3>
                    <p className="text-sm text-slate-400">{step.description}</p>
                  </CardContent>
                </Card>
                {index < steps.length - 1 && (
                  <ArrowRight className="w-6 h-6 text-slate-500 hidden lg:block" />
                )}
              </div>
            );
          })}
        </div>
      </div>

      {/* Key Steps Details */}
      <div className="grid md:grid-cols-2 gap-8">
        <Card className="bg-slate-900/30 border-slate-700">
          <CardContent className="p-8">
            <h3 className="text-2xl font-bold text-white mb-6">Input: YAML Manifest</h3>
            <div className="bg-slate-900 rounded-lg p-4 font-mono text-sm">
              <div className="text-slate-400"># example: aws-static-website.yaml</div>
              <div className="text-blue-400">apiVersion: <span className="text-white">aws.project-planton.org/v1</span></div>
              <div className="text-blue-400">kind: <span className="text-white">AwsStaticWebsite</span></div>
              <div className="text-blue-400">metadata:</div>
              <div className="text-blue-400 ml-4">name: <span className="text-white">my-site</span></div>
              <div className="text-blue-400">spec:</div>
              <div className="text-blue-400 ml-4">enableCdn: <span className="text-emerald-400">true</span></div>
            </div>
          </CardContent>
        </Card>

        <Card className="bg-slate-900/30 border-slate-700">
          <CardContent className="p-8">
            <h3 className="text-2xl font-bold text-white mb-6">Output: Status/Results</h3>
            <div className="bg-slate-900 rounded-lg p-4 font-mono text-sm">
              <div className="text-slate-400"># modules export status.outputs</div>
              <div className="text-blue-400">status:</div>
              <div className="text-blue-400 ml-4">outputs:</div>
              <div className="text-blue-400 ml-8">hostname: <span className="text-white">my-site.s3-website.amazonaws.com</span></div>
              <div className="text-blue-400 ml-8">cloudfront_url: <span className="text-white">d1234.cloudfront.net</span></div>
              <div className="text-blue-400 ml-8">bucket_name: <span className="text-white">my-site-bucket-xyz</span></div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}