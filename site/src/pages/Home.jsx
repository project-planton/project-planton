import React, { useState, useEffect } from 'react';
import { 
  Code,
  Menu,
  X,
  Github
} from 'lucide-react';

import Hero from '../components/landing/Hero';
import FeatureCards from '../components/landing/FeatureCards';
import HowItWorks from '../components/landing/HowItWorks';
import Quickstart from '../components/landing/Quickstart';
import ExampleGallery from '../components/landing/ExampleGallery';
import CLIReference from '../components/landing/CLIReference';
import CICDSection from '../components/landing/CICDSection';
import CompareSection from '../components/landing/CompareSection';
import FAQ from '../components/landing/FAQ';
import Footer from '../components/landing/Footer';

export default function Home() {
  const [activeSection, setActiveSection] = useState('hero');
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  const navItems = [
    { id: 'hero', label: 'Home' },
    { id: 'why', label: 'Why' },
    { id: 'how', label: 'How it works' },
    { id: 'quickstart', label: 'Quickstart' },
    { id: 'examples', label: 'Examples' },
    { id: 'cli', label: 'CLI' },
    { id: 'cicd', label: 'CI/CD' },
    { id: 'compare', label: 'Compare' },
    { id: 'faq', label: 'FAQ' }
  ];

  useEffect(() => {
    const handleScroll = () => {
      const sections = navItems.map(item => document.getElementById(item.id));
      const scrollPosition = window.scrollY + 100;

      for (let i = sections.length - 1; i >= 0; i--) {
        const section = sections[i];
        if (section && section.offsetTop <= scrollPosition) {
          setActiveSection(navItems[i].id);
          break;
        }
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const scrollToSection = (id) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
      setMobileMenuOpen(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 text-white overflow-x-hidden">
      {/* Navigation */}
      <nav className="fixed top-0 w-full bg-slate-950/95 backdrop-blur-sm border-b border-slate-800 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            {/* Logo */}
            <div className="flex items-center gap-3">
              <img src="/icon.png" alt="ProjectPlanton logo" className="h-9 w-auto object-contain" />
              <img src="/logo-text.svg" alt="ProjectPlanton" className="h-10 w-auto object-contain" />
            </div>

            {/* Desktop Navigation */}
            <div className="hidden lg:flex items-center gap-8">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => scrollToSection(item.id)}
                  className={`text-sm font-medium transition-colors duration-200 ${
                    activeSection === item.id
                      ? 'text-[#7a4183]'
                      : 'text-slate-300 hover:text-white'
                  }`}
                >
                  {item.label}
                </button>
              ))}
              <div className="flex items-center gap-3">
                <a 
                  href="https://github.com/project-planton/project-planton" 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="text-slate-300 hover:text-white transition-colors"
                >
                  <Github className="w-5 h-5" />
                </a>
              </div>
            </div>

            {/* Mobile menu button */}
            <button
              className="lg:hidden p-2"
              onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
            >
              {mobileMenuOpen ? <X className="w-5 h-5" /> : <Menu className="w-5 h-5" />}
            </button>
          </div>
        </div>

        {/* Mobile Navigation */}
        {mobileMenuOpen && (
          <div className="lg:hidden bg-slate-900 border-t border-slate-800">
            <div className="px-4 py-6 space-y-4">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => scrollToSection(item.id)}
                  className={`block w-full text-left py-2 text-base font-medium transition-colors duration-200 ${
                    activeSection === item.id
                      ? 'text-[#7a4183]'
                      : 'text-slate-300 hover:text-white'
                  }`}
                >
                  {item.label}
                </button>
              ))}
              <div className="pt-4 border-t border-slate-800 flex gap-4">
                <a 
                  href="https://github.com/project-planton/project-planton" 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="flex items-center gap-2 text-slate-300 hover:text-white transition-colors"
                >
                  <Github className="w-5 h-5" />
                  GitHub
                </a>
              </div>
            </div>
          </div>
        )}
      </nav>

      {/* Main Content */}
      <main className="pt-16">
        <section id="hero">
          <Hero />
        </section>
        
        <section id="why" className="py-24">
          <FeatureCards />
        </section>

        <section id="how" className="py-24">
          <HowItWorks />
        </section>

        <section id="quickstart" className="py-24">
          <Quickstart />
        </section>

        <section id="examples" className="py-24">
          <ExampleGallery />
        </section>

        <section id="cli" className="py-24">
          <CLIReference />
        </section>

        <section id="cicd" className="py-24">
          <CICDSection />
        </section>

        <section id="compare" className="py-24">
          <CompareSection />
        </section>

        <section id="faq" className="py-24">
          <FAQ />
        </section>
      </main>

      <Footer />
    </div>
  );
}