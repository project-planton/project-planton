import React from 'react';
import { Github, Scale } from 'lucide-react';

export default function Footer() {
  const links = [
    { name: 'GitHub', href: 'https://github.com/project-planton/project-planton', external: true },
    { name: 'Providers', href: '#', external: false },
    { name: 'CLI', href: '#cli', external: false },
    { name: 'Examples', href: '#examples', external: false },
  ];

  return (
    <footer className="bg-slate-900 border-t border-slate-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
        <div className="grid md:grid-cols-4 gap-8">
          {/* Brand */}
          <div className="md:col-span-2">
            <div className="flex items-center gap-3 mb-4">
              <div className="w-8 h-8 bg-gradient-to-r from-blue-500 to-cyan-400 rounded-lg flex items-center justify-center">
                <Github className="w-5 h-5 text-white" />
              </div>
              <span className="text-xl font-bold text-white">ProjectPlanton</span>
            </div>
            <p className="text-slate-400 max-w-md mb-6">
              Kubernetes‑style manifests for multi‑cloud infrastructure. 
              Define once. Deploy anywhere.
            </p>
            <div className="flex items-center gap-2 text-sm text-slate-500">
              <Scale className="w-4 h-4" />
              <span>Apache-2.0 License</span>
            </div>
          </div>

          {/* Links */}
          <div>
            <h3 className="font-semibold text-white mb-4">Resources</h3>
            <div className="space-y-3">
              {links.map((link) => (
                <a
                  key={link.name}
                  href={link.href}
                  className="block text-slate-400 hover:text-white transition-colors"
                  target={link.external ? '_blank' : undefined}
                  rel={link.external ? 'noopener noreferrer' : undefined}
                >
                  {link.name}
                </a>
              ))}
            </div>
          </div>

          {/* Community */}
          <div>
            <h3 className="font-semibold text-white mb-4">Community</h3>
            <div className="space-y-3">
              <a
                href="https://github.com/project-planton/project-planton"
                target="_blank"
                rel="noopener noreferrer"
                className="block text-slate-400 hover:text-white transition-colors"
              >
                <div className="flex items-center gap-2">
                  <Github className="w-4 h-4" />
                  GitHub
                </div>
              </a>
              <a
                href="#"
                className="block text-slate-400 hover:text-white transition-colors"
              >
                Discussions
              </a>
              <a
                href="#"
                className="block text-slate-400 hover:text-white transition-colors"
              >
                Issues
              </a>
            </div>
          </div>
        </div>

        <div className="border-t border-slate-800 mt-12 pt-8 text-center">
          <p className="text-slate-500">
            © 2024 ProjectPlanton. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
}