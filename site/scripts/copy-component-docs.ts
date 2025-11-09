#!/usr/bin/env node

import * as fs from 'fs';
import * as path from 'path';

/**
 * Build script to copy deployment component documentation from apis/ to site/public/docs/provider/
 * 
 * Scans: apis/project/planton/provider/{provider}/{component}/v1/docs/README.md
 * Outputs: site/public/docs/provider/{provider}/{component}.md
 * 
 * Generates frontmatter for each component and creates provider index pages.
 */

interface ComponentDoc {
  provider: string;
  component: string;
  sourcePath: string;
  content: string;
  title: string;
}

interface Stats {
  total: number;
  copied: number;
  skipped: number;
  providers: Set<string>;
}

/**
 * Generate human-readable title from component name
 * Examples:
 *   awsalb -> AWS ALB
 *   gcpgkecluster -> GCP GKE Cluster
 *   kuberneteshttpendpoint -> Kubernetes HTTP Endpoint
 */
function generateTitle(component: string, provider: string): string {
  // Remove provider prefix if component starts with it
  let name = component;
  if (name.toLowerCase().startsWith(provider.toLowerCase())) {
    name = name.substring(provider.length);
  }

  // Split camelCase/PascalCase into words
  // Handle common acronyms: ALB, EKS, GKE, VPC, DNS, etc.
  const acronyms = ['ALB', 'EKS', 'GKE', 'VPC', 'DNS', 'IAM', 'ACM', 'S3', 'EC2', 'ECS', 'RDS', 'CDN', 'HTTP', 'HTTPS', 'API', 'SDK', 'CLI', 'NAT', 'IP', 'SSL', 'TLS', 'WAF', 'KV', 'D1', 'R2'];
  
  // Insert spaces before uppercase letters
  let spaced = name.replace(/([A-Z])/g, ' $1').trim();
  
  // Uppercase known acronyms
  acronyms.forEach(acronym => {
    const regex = new RegExp(`\\b${acronym}\\b`, 'gi');
    spaced = spaced.replace(regex, acronym);
  });

  // Capitalize provider name
  const providerTitle = provider.toUpperCase();
  
  return `${providerTitle} ${spaced}`;
}

/**
 * Generate frontmatter for a component doc
 */
function generateFrontmatter(title: string, component: string, description?: string): string {
  const desc = description || `${title} deployment documentation`;
  return `---
title: "${title}"
description: "${desc}"
icon: "package"
order: 100
componentName: "${component}"
---`;
}

/**
 * Scan a provider directory for components with docs
 */
function scanProvider(providerPath: string, provider: string): ComponentDoc[] {
  const docs: ComponentDoc[] = [];
  
  if (!fs.existsSync(providerPath)) {
    return docs;
  }

  const items = fs.readdirSync(providerPath);
  
  for (const item of items) {
    const componentPath = path.join(providerPath, item);
    const stat = fs.statSync(componentPath);
    
    if (!stat.isDirectory()) {
      continue;
    }

    // Check for v1/docs/README.md
    const docPath = path.join(componentPath, 'v1', 'docs', 'README.md');
    
    if (fs.existsSync(docPath)) {
      const content = fs.readFileSync(docPath, 'utf-8');
      const title = generateTitle(item, provider);
      
      docs.push({
        provider,
        component: item,
        sourcePath: docPath,
        content,
        title,
      });
    }
  }
  
  return docs;
}

/**
 * Write component doc to site/public/docs/provider/{provider}/{component}.md
 */
function writeComponentDoc(
  doc: ComponentDoc,
  outputRoot: string
): void {
  const providerDir = path.join(outputRoot, doc.provider);
  
  // Ensure provider directory exists
  if (!fs.existsSync(providerDir)) {
    fs.mkdirSync(providerDir, { recursive: true });
  }
  
  // Generate output with frontmatter
  const frontmatter = generateFrontmatter(doc.title, doc.component);
  const output = `${frontmatter}\n\n${doc.content}`;
  
  // Write to {provider}/{component}.md
  const outputPath = path.join(providerDir, `${doc.component}.md`);
  fs.writeFileSync(outputPath, output, 'utf-8');
}

/**
 * Generate provider index page listing all components
 */
function generateProviderIndex(
  provider: string,
  docs: ComponentDoc[],
  outputRoot: string
): void {
  const providerDir = path.join(outputRoot, provider);
  
  if (!fs.existsSync(providerDir)) {
    fs.mkdirSync(providerDir, { recursive: true });
  }

  const providerTitle = provider.toUpperCase();
  
  // Sort docs alphabetically by component name
  const sortedDocs = [...docs].sort((a, b) => 
    a.component.localeCompare(b.component)
  );
  
  // Generate component list
  const componentList = sortedDocs
    .map(doc => `- [${doc.title}](/docs/provider/${provider}/${doc.component})`)
    .join('\n');
  
  const indexContent = `---
title: "${providerTitle}"
description: "Deploy ${providerTitle} resources using Project Planton"
icon: "cloud"
order: 10
---

# ${providerTitle}

The following ${providerTitle} resources can be deployed using Project Planton:

${componentList}
`;

  const indexPath = path.join(providerDir, 'index.md');
  fs.writeFileSync(indexPath, indexContent, 'utf-8');
}

/**
 * Generate main provider index page
 */
function generateMainIndex(providers: string[], outputRoot: string): void {
  // Sort providers alphabetically
  const sortedProviders = [...providers].sort();
  
  const providerList = sortedProviders
    .map(provider => {
      const title = provider.toUpperCase();
      return `- [${title}](/docs/provider/${provider})`;
    })
    .join('\n');
  
  const indexContent = `---
title: "Deployment Components by Provider"
description: "Browse deployment components organized by cloud provider"
icon: "package"
order: 50
---

# Deployment Components

Browse deployment components by provider:

${providerList}
`;

  const indexPath = path.join(outputRoot, 'index.md');
  fs.writeFileSync(indexPath, indexContent, 'utf-8');
}

/**
 * Main function to copy all component docs
 */
async function copyComponentDocs(): Promise<void> {
  console.log('🚀 Starting component documentation copy process...\n');
  
  // Paths
  const scriptDir = __dirname;
  const projectRoot = path.join(scriptDir, '../..');
  const apisRoot = path.join(projectRoot, 'apis/project/planton/provider');
  const siteDocsRoot = path.join(scriptDir, '../public/docs/provider');
  
  // Clear existing generated docs
  if (fs.existsSync(siteDocsRoot)) {
    console.log(`🗑️  Clearing existing docs at ${path.relative(projectRoot, siteDocsRoot)}`);
    fs.rmSync(siteDocsRoot, { recursive: true });
  }
  
  // Create output directory
  fs.mkdirSync(siteDocsRoot, { recursive: true });
  
  // Stats
  const stats: Stats = {
    total: 0,
    copied: 0,
    skipped: 0,
    providers: new Set(),
  };
  
  // Track docs by provider for index generation
  const docsByProvider: Map<string, ComponentDoc[]> = new Map();
  
  // Scan all providers
  if (!fs.existsSync(apisRoot)) {
    console.error(`❌ Error: APIs directory not found at ${apisRoot}`);
    process.exit(1);
  }
  
  const providers = fs.readdirSync(apisRoot).filter(item => {
    const itemPath = path.join(apisRoot, item);
    return fs.statSync(itemPath).isDirectory();
  });
  
  console.log(`📁 Scanning ${providers.length} providers...\n`);
  
  for (const provider of providers) {
    const providerPath = path.join(apisRoot, provider);
    const docs = scanProvider(providerPath, provider);
    
    if (docs.length > 0) {
      stats.providers.add(provider);
      docsByProvider.set(provider, docs);
      
      console.log(`📦 ${provider.toUpperCase()}: Found ${docs.length} components`);
      
      // Write each component doc
      for (const doc of docs) {
        try {
          writeComponentDoc(doc, siteDocsRoot);
          stats.copied++;
          console.log(`   ✓ ${doc.component}`);
        } catch (error) {
          console.error(`   ✗ ${doc.component}: ${error}`);
          stats.skipped++;
        }
      }
      
      // Generate provider index
      generateProviderIndex(provider, docs, siteDocsRoot);
      console.log(`   ✓ Generated index page\n`);
    }
  }
  
  // Generate main provider index
  if (stats.providers.size > 0) {
    generateMainIndex(Array.from(stats.providers), siteDocsRoot);
    console.log(`✓ Generated main provider index\n`);
  }
  
  // Summary
  console.log('📊 Summary:');
  console.log(`   Providers: ${stats.providers.size}`);
  console.log(`   Components copied: ${stats.copied}`);
  console.log(`   Components skipped: ${stats.skipped}`);
  console.log(`   Output: ${path.relative(projectRoot, siteDocsRoot)}`);
  console.log('\n✅ Component documentation copy complete!\n');
}

// Run the script
copyComponentDocs().catch(error => {
  console.error('❌ Error copying component docs:', error);
  process.exit(1);
});

