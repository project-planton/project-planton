import React, { useState } from 'react';
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Copy, Check, Github, ArrowDown } from 'lucide-react';

export default function Hero() {
  const [copied, setCopied] = useState(false);
  
  const installCommand = 'brew install project-planton/tap/project-planton';

  const copyToClipboard = () => {
    navigator.clipboard.writeText(installCommand);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const scrollToExample = () => {
    const element = document.getElementById('examples');
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  };

  return (
    <div className="relative overflow-hidden">
      {/* Background gradient */}
      <div className="absolute inset-0 bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950" />
      <div className="absolute inset-0 opacity-50" style={{
        backgroundImage: `url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23334155' fill-opacity='0.1'%3E%3Ccircle cx='30' cy='30' r='1'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E")`
      }} />
      
      <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-32">
        <div className="text-center max-w-4xl mx-auto">
          {/* Badges */}
          <div className="flex flex-wrap justify-center gap-2 mb-8">
            <Badge variant="outline" className="border-slate-600 text-slate-300 bg-slate-800/50">
              Apache-2.0
            </Badge>
            <Badge variant="outline" className="border-slate-600 text-slate-300 bg-slate-800/50">
              KRM/Protobuf/Buf
            </Badge>
            <Badge variant="outline" className="border-slate-600 text-slate-300 bg-slate-800/50">
              Pulumi/OpenTofu
            </Badge>
            <Badge variant="outline" className="border-slate-600 text-slate-300 bg-slate-800/50">
              CLI-first
            </Badge>
            <Badge variant="outline" className="border-slate-600 text-slate-300 bg-slate-800/50">
              CI/CD-ready
            </Badge>
          </div>

          {/* Main heading */}
          <h1 className="text-5xl sm:text-6xl lg:text-7xl font-bold mb-8 leading-tight">
            <span className="bg-gradient-to-r from-white via-slate-200 to-slate-400 bg-clip-text text-transparent">
              Open‑Source Multi‑Cloud
            </span>
            <br />
            <span className="bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400 bg-clip-text text-transparent">
              Infrastructure Framework
            </span>
          </h1>

          {/* Subtitle */}
          <p className="text-xl sm:text-2xl text-slate-300 mb-12 leading-relaxed max-w-3xl mx-auto">
            Author KRM‑style YAML once, validate with Protobuf + Buf ProtoValidate, then execute with Pulumi or OpenTofu. 
            Consistent APIs across AWS, GCP, Azure, and Kubernetes—no provider‑specific yak‑shaving.
          </p>

          {/* CTAs */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-8">
            <Button
              size="lg"
              onClick={copyToClipboard}
              className="bg-blue-600 hover:bg-blue-700 text-white font-mono text-sm px-8 py-3 h-auto"
            >
              {copied ? (
                <>
                  <Check className="w-4 h-4 mr-2" />
                  Copied!
                </>
              ) : (
                <>
                  <Copy className="w-4 h-4 mr-2" />
                  {installCommand}
                </>
              )}
            </Button>
            
            <Button
              size="lg"
              variant="outline"
              onClick={scrollToExample}
              className="border-slate-600 bg-slate-800/50 text-slate-300 hover:text-white hover:bg-slate-700 px-8 py-3 h-auto"
            >
              Try an Example
              <ArrowDown className="w-4 h-4 ml-2" />
            </Button>
          </div>

          {/* Tertiary links */}
          <div className="flex flex-wrap justify-center gap-6 text-slate-400">
            <a 
              href="https://github.com/project-planton/project-planton" 
              target="_blank" 
              rel="noopener noreferrer"
              className="flex items-center gap-2 hover:text-white transition-colors"
            >
              <Github className="w-4 h-4" />
              View on GitHub
            </a>
          </div>

          {/* Tagline */}
          <div className="mt-16 pt-16 border-t border-slate-800">
            <p className="text-3xl font-bold text-center">
              <span className="text-slate-400">Define once.</span>
              <span className="text-blue-400 ml-4">Deploy anywhere.</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}