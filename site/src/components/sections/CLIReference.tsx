import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Terminal, Flag, Command } from "lucide-react";

export default function CLIReference() {
  const primaryCommands = [
    { command: "project-planton pulumi preview", description: "Preview Pulumi changes" },
    { command: "project-planton pulumi update", description: "Apply Pulumi changes" },
    { command: "project-planton pulumi refresh", description: "Refresh Pulumi state" },
    { command: "project-planton pulumi destroy", description: "Destroy Pulumi resources" },
    { command: "project-planton tofu init", description: "Initialize OpenTofu backend" },
    { command: "project-planton tofu plan", description: "Create OpenTofu execution plan" },
    { command: "project-planton tofu apply", description: "Apply OpenTofu changes" },
    { command: "project-planton tofu destroy", description: "Destroy OpenTofu resources" },
    { command: "project-planton tofu refresh", description: "Refresh OpenTofu state" },
    { command: "project-planton validate <manifest.yaml>", description: "Validate manifest with Buf ProtoValidate" },
    { command: "project-planton version", description: "Show CLI version" },
  ];

  const coreFlags = [
    { flag: "--manifest <path>", description: "path to the deployment‑component manifest file", required: true },
    { flag: "--stack <org/project/stack>", description: "Pulumi stack FQDN (required for pulumi commands)", required: true },
    { flag: "--module-dir <dir>", description: "directory containing the Pulumi/Tofu module (defaults to current dir)", required: false },
    { flag: "--input-dir <dir>", description: "directory containing target.yaml and/or credential yamls", required: false },
    { flag: "--set key=value", description: "override manifest fields (deep dot‑paths supported)", required: false },
    { flag: "--auto-approve", description: "skip interactive approval (tofu apply/destroy)", required: false },
    { flag: "--destroy", description: "create a destroy plan (tofu plan)", required: false },
    { flag: "--backend-type {local|s3|gcs|azurerm}", description: "choose backend (tofu init)", required: false },
    { flag: "--backend-config key=value", description: "repeatable backend arguments (tofu init)", required: false },
  ];

  const credentialFlags = [
    { flag: "--aws-credential <file>", description: "AWS credential file" },
    { flag: "--gcp-credential <file>", description: "GCP credential file" },
    { flag: "--azure-credential <file>", description: "Azure credential file" },
    { flag: "--kubernetes-cluster <file>", description: "Kubernetes cluster config file" },
    { flag: "--confluent-credential <file>", description: "Confluent credential file" },
    { flag: "--mongodb-atlas-credential <file>", description: "MongoDB Atlas credential file" },
    { flag: "--snowflake-credential <file>", description: "Snowflake credential file" },
  ];

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">CLI Reference</span>
        </h2>
        <p className="text-xl text-slate-400 max-w-3xl mx-auto">Complete command reference for the ProjectPlanton CLI</p>
      </div>

      <div className="grid lg:grid-cols-2 gap-8 mb-8">
        <Card className="bg-slate-900/30 border-slate-700">
          <CardHeader className="pb-4">
            <CardTitle className="text-xl text-white flex items-center gap-2">
              <Terminal className="w-5 h-5" />
              Primary Commands
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3">
              {primaryCommands.map((cmd, index) => (
                <div key={index} className="border-b border-slate-800 pb-3 last:border-b-0 last:pb-0">
                  <code className="text-blue-400 text-sm">{cmd.command}</code>
                  <p className="text-slate-400 text-sm mt-1">{cmd.description}</p>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>

        <Card className="bg-slate-900/30 border-slate-700">
          <CardHeader className="pb-4">
            <CardTitle className="text-xl text-white flex items-center gap-2">
              <Command className="w-5 h-5" />
              Validation Example
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="bg-slate-950 rounded-lg p-4 font-mono text-sm">
              <div className="text-slate-400 mb-2"># Validate a manifest before deployment</div>
              <div className="text-blue-400">project-planton validate manifest.yaml</div>
              <div className="text-slate-400 mt-4 mb-2"># Example output:</div>
              <div className="text-green-400">✓ Manifest validation passed</div>
              <div className="text-green-400">✓ Buf ProtoValidate rules: OK</div>
              <div className="text-green-400">✓ CEL expressions: OK</div>
              <div className="text-green-400">✓ Required fields: OK</div>
            </div>
          </CardContent>
        </Card>
      </div>

      <Card className="bg-slate-900/30 border-slate-700 mb-8">
        <CardHeader className="pb-4">
          <CardTitle className="text-xl text-white flex items-center gap-2">
            <Flag className="w-5 h-5" />
            Core Flags
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-x-auto">
            <Table>
              <TableHeader>
                <TableRow className="border-slate-800">
                  <TableHead className="text-slate-300">Flag</TableHead>
                  <TableHead className="text-slate-300">Description</TableHead>
                  <TableHead className="text-slate-300">Required</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {coreFlags.map((flag, index) => (
                  <TableRow key={index} className="border-slate-800">
                    <TableCell className="font-mono text-blue-400">{flag.flag}</TableCell>
                    <TableCell className="text-slate-300">{flag.description}</TableCell>
                    <TableCell>
                      {flag.required ? (
                        <Badge className="bg-red-900 text-red-200">Required</Badge>
                      ) : (
                        <Badge variant="outline" className="border-slate-600 text-slate-400">Optional</Badge>
                      )}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>

      <Card className="bg-slate-900/30 border-slate-700">
        <CardHeader className="pb-4">
          <CardTitle className="text-xl text-white">Provider Credential Flags</CardTitle>
          <p className="text-slate-400 text-sm">Supported by stackinputcredentials</p>
        </CardHeader>
        <CardContent>
          <div className="grid md:grid-cols-2 gap-4">
            {credentialFlags.map((flag, index) => (
              <div key={index} className="border border-slate-800 rounded-lg p-4">
                <code className="text-blue-400 text-sm">{flag.flag}</code>
                <p className="text-slate-400 text-sm mt-1">{flag.description}</p>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}


