import React, { useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Copy, Check } from 'lucide-react';

export default function Quickstart() {
  const [copiedSteps, setCopiedSteps] = useState({});

  const copyCode = (stepId, code) => {
    navigator.clipboard.writeText(code);
    setCopiedSteps(prev => ({ ...prev, [stepId]: true }));
    setTimeout(() => {
      setCopiedSteps(prev => ({ ...prev, [stepId]: false }));
    }, 2000);
  };

  const steps = [
    {
      id: 'install',
      title: '1. Install',
      code: 'brew install project-planton/tap/project-planton'
    },
    {
      id: 'validate',
      title: '2. Validate a manifest',
      code: 'project-planton validate manifest.yaml'
    },
    {
      id: 'deploy',
      title: '3. Deploy',
      code: `# Pulumi (example)
project-planton pulumi up \\
  --manifest manifest.yaml \\
  --stack myorg/myproject/dev \\
  --module-dir . \\
  --set metadata.labels.env=dev

# OpenTofu (example)
project-planton tofu init \\
  --manifest manifest.yaml \\
  --backend-type s3 \\
  --backend-config bucket=my-tf-state-bucket \\
  --backend-config dynamodb_table=my-tf-locks \\
  --backend-config region=us-west-2 \\
  --backend-config key=stacks/myproject/dev.tfstate

project-planton tofu plan --manifest manifest.yaml
project-planton tofu apply --manifest manifest.yaml --auto-approve`
    },
    {
      id: 'destroy',
      title: '4. Destroy (optional)',
      code: `project-planton pulumi destroy --manifest manifest.yaml --stack myorg/myproject/dev
# or
project-planton tofu destroy --manifest manifest.yaml --auto-approve`
    }
  ];

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            Get Started in Minutes
          </span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto">
          Four simple steps to deploy your first multi-cloud infrastructure
        </p>
      </div>

      <div className="grid gap-6 max-w-4xl mx-auto">
        {steps.map((step) => (
          <Card key={step.id} className="bg-slate-900/30 border-slate-700">
            <CardHeader className="pb-4">
              <CardTitle className="text-xl text-white flex items-center justify-between">
                {step.title}
                <Button
                  size="sm"
                  variant="outline"
                  onClick={() => copyCode(step.id, step.code)}
                  className="border-slate-600 bg-slate-800/50 text-slate-300 hover:text-white hover:bg-slate-700"
                >
                  {copiedSteps[step.id] ? (
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
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="bg-slate-950 rounded-lg p-4 font-mono text-sm overflow-x-auto">
                <pre className="text-slate-300 whitespace-pre-wrap">
                  <code>{step.code}</code>
                </pre>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      <Card className="mt-8 bg-blue-950/30 border-blue-800/50">
        <CardContent className="p-6">
          <h3 className="text-lg font-bold text-white mb-4">Prerequisites & Notes</h3>
          <ul className="text-slate-300 space-y-2">
            <li>• Provide credentials via files in an input dir (e.g., aws-credential.yaml, gcp-credential.yaml, kubernetes-cluster.yaml) or explicit flags</li>
            <li>• Pulumi/OpenTofu CLIs must be installed separately</li>
            <li>• OpenTofu supports backends: local, s3, gcs, azurerm</li>
            <li>• Stack FQDN format is &lt;org&gt;/&lt;project&gt;/&lt;stack&gt;</li>
          </ul>
        </CardContent>
      </Card>
    </div>
  );
}