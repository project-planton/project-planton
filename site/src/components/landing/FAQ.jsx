import React, { useState } from 'react';
import { Card, CardContent } from "@/components/ui/card";
import { ChevronDown, ChevronRight } from 'lucide-react';

export default function FAQ() {
  const [openItems, setOpenItems] = useState(new Set());

  const toggleItem = (index) => {
    const newOpenItems = new Set(openItems);
    if (newOpenItems.has(index)) {
      newOpenItems.delete(index);
    } else {
      newOpenItems.add(index);
    }
    setOpenItems(newOpenItems);
  };

  const faqs = [
    {
      question: "Do I need Pulumi/Terraform knowledge?",
      answer: "Not to start. You work with declarative manifests; the CLI selects and executes Pulumi/OpenTofu modules for you."
    },
    {
      question: "Which providers are supported today?",
      answer: "AWS, GCP, Azure, Kubernetes (plus others like Cloudflare, DigitalOcean, Civo, Confluent, Snowflake appear in the enums and modules). Use the examples and apis/.../provider/* directories for concrete kinds."
    },
    {
      question: "How are validations enforced?",
      answer: "Each spec is a Protobuf message with Buf ProtoValidate rules (including CEL). project-planton validate runs the validations before provisioning."
    },
    {
      question: "Where is state stored?",
      answer: "OpenTofu supports local|s3|gcs|azurerm backends (via tofu init). Pulumi uses standard Pulumi backends; stacks are selected via <org>/<project>/<stack>."
    },
    {
      question: "How do credentials work?",
      answer: "Provide credential yamls (e.g., aws-credential.yaml, gcp-credential.yaml, kubernetes-cluster.yaml) or pass file paths via flags; the CLI injects env vars/inputs for the selected engine."
    },
    {
      question: "How do I bring custom modules?",
      answer: "Point --module-dir at your repo path. The CLI resolves the kind/provider and executes the appropriate submodule folder."
    },
    {
      question: "How do I run in CI?",
      answer: "Install the CLI + engine (Pulumi/OpenTofu), provide credentials (as files or secrets), then run the same commands; see the CI/CD example."
    },
    {
      question: "Licensing & governance?",
      answer: "Apache‑2.0 license; modules apply consistent labels (e.g., planton.org/*). Validations and defaults are embedded in the schema for predictable behavior."
    }
  ];

  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div className="text-center mb-16">
        <h2 className="text-4xl font-bold mb-6">
          <span className="bg-gradient-to-r from-white to-slate-300 bg-clip-text text-transparent">
            Frequently Asked Questions
          </span>
        </h2>
        <p className="text-xl text-slate-400">
          Everything you need to know about ProjectPlanton
        </p>
      </div>

      <div className="space-y-4">
        {faqs.map((faq, index) => {
          const isOpen = openItems.has(index);
          return (
            <Card key={index} className="bg-slate-900/30 border-slate-700">
              <CardContent className="p-0">
                <button
                  className="w-full text-left p-6 flex items-center justify-between hover:bg-slate-800/50 transition-colors"
                  onClick={() => toggleItem(index)}
                >
                  <h3 className="text-lg font-semibold text-white pr-4">{faq.question}</h3>
                  {isOpen ? (
                    <ChevronDown className="w-5 h-5 text-slate-400 flex-shrink-0" />
                  ) : (
                    <ChevronRight className="w-5 h-5 text-slate-400 flex-shrink-0" />
                  )}
                </button>
                {isOpen && (
                  <div className="px-6 pb-6">
                    <p className="text-slate-400 leading-relaxed">{faq.answer}</p>
                  </div>
                )}
              </CardContent>
            </Card>
          );
        })}
      </div>
    </div>
  );
}