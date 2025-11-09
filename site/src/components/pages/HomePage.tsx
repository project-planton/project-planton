"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import Link from "next/link";
import Hero from "@/components/sections/Hero";
import FeatureCards from "@/components/sections/FeatureCards";
import HowItWorks from "@/components/sections/HowItWorks";
import Quickstart from "@/components/sections/Quickstart";
import ExampleGallery from "@/components/sections/ExampleGallery";
import CLIReference from "@/components/sections/CLIReference";
import CICDSection from "@/components/sections/CICDSection";
import CompareSection from "@/components/sections/CompareSection";
import FAQ from "@/components/sections/FAQ";
import Footer from "@/components/sections/Footer";
import { GitHubStarBadge } from "@/components/ui/GitHubStarBadge";

interface NavItem {
  id: string;
  label: string;
  isExternal?: boolean;
  href?: string;
}

export default function HomePage() {
  const [activeSection, setActiveSection] = useState<string>("hero");

  const navItems: NavItem[] = [
    { id: "hero", label: "Home" },
    { id: "why", label: "Why" },
    { id: "how", label: "How it works" },
    { id: "quickstart", label: "Quickstart" },
    { id: "examples", label: "Examples" },
    { id: "cli", label: "CLI" },
    { id: "cicd", label: "CI/CD" },
    { id: "compare", label: "Compare" },
    { id: "faq", label: "FAQ" },
  ];

  useEffect(() => {
    const handleScroll = () => {
      const sections = navItems.map((item) => document.getElementById(item.id));
      const scrollPosition = window.scrollY + 100;
      for (let i = sections.length - 1; i >= 0; i--) {
        const section = sections[i];
        if (section && section.offsetTop <= scrollPosition) {
          setActiveSection(navItems[i].id);
          break;
        }
      }
    };
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, [navItems]);

  const scrollToSection = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: "smooth" });
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 text-white overflow-x-hidden">
      <nav className="fixed top-0 w-full bg-slate-950/95 backdrop-blur-sm border-b border-slate-800 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center gap-3">
              <Image src="/icon.png" alt="ProjectPlanton logo" width={36} height={36} className="h-9 w-auto object-contain" priority />
              <Image src="/logo-text.svg" alt="ProjectPlanton" width={160} height={40} className="h-10 w-auto object-contain" priority />
            </div>
            <div className="flex items-center gap-4 sm:gap-6">
              <Link
                href="/docs"
                className="text-sm font-medium text-slate-300 hover:text-white transition-colors"
              >
                Docs
              </Link>
              <GitHubStarBadge repo="project-planton/project-planton" />
            </div>
          </div>
        </div>
      </nav>

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


