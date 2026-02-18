'use client';

import Link from 'next/link';
import { useState, useEffect } from 'react';
import TerminalTyping from '@/components/TerminalTyping';
import Navbar from '@/components/Navbar';

const terminalLines = [
  {
    prompt: '❯',
    command: 'exo init',
    output: [
      '<span class="text-[#9b5de5]">EXO Setup Wizard</span>',
      '<span class="text-[#6c6c8a]">[████████░░] 4/7</span>',
      '',
      '  <span class="text-[#f15bb5] font-bold">Project:</span>    my-service',
      '  <span class="text-[#f15bb5] font-bold">Language:</span>   go',
      '  <span class="text-[#f15bb5] font-bold">Provider:</span>   aws',
      '  <span class="text-[#f15bb5] font-bold">CI/CD:</span>      github-actions',
      '  <span class="text-[#f15bb5] font-bold">Monitoring:</span> prometheus',
      '  <span class="text-[#f15bb5] font-bold">Database:</span>   postgres',
    ],
  },
  {
    prompt: '❯',
    command: 'exo status',
    output: [
      '<span class="text-[#f15bb5] font-bold">EXO Project Status</span>',
      '',
      '  <span class="text-[#00ff41] font-bold">✓</span>  Dockerfile                 <span class="text-[#6c6c8a]">1.2 KB  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  Terraform (AWS)            <span class="text-[#6c6c8a]">dir  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  GitHub Actions             <span class="text-[#6c6c8a]">dir  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  K8s Manifests              <span class="text-[#6c6c8a]">dir  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  Monitoring                 <span class="text-[#6c6c8a]">dir  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  DB (PostgreSQL)            <span class="text-[#6c6c8a]">0.5 KB  •  just now</span>',
      '  <span class="text-[#00ff41] font-bold">✓</span>  EXO Config                 <span class="text-[#6c6c8a]">0.1 KB  •  just now</span>',
    ],
  },
];

const features = [
  {
    icon: '⚡',
    title: 'Smart Detection',
    desc: 'Auto-detects Go, Node.js, Python from source files. Zero config.',
    color: 'text-arch-cyan',
  },
  {
    icon: '🐳',
    title: 'Multi-Stage Docker',
    desc: 'Optimized, language-specific Dockerfiles. Smaller images, faster builds.',
    color: 'text-arch-blue',
  },
  {
    icon: '☁️',
    title: 'Multi-Cloud Terraform',
    desc: 'AWS, GCP, Azure — VPC, EKS/GKE/AKS with one command.',
    color: 'text-arch-purple',
  },
  {
    icon: '🔄',
    title: 'CI/CD Pipelines',
    desc: 'GitHub Actions & GitLab CI workflows generated automatically.',
    color: 'text-arch-green',
  },
  {
    icon: '☸️',
    title: 'Kubernetes Ready',
    desc: 'Deployment, Service, Ingress + Helm charts. No YAML headaches.',
    color: 'text-arch-orange',
  },
  {
    icon: '📊',
    title: 'Monitoring Stack',
    desc: 'Prometheus + Grafana with dashboards and alert rules built in.',
    color: 'text-arch-pink',
  },
  {
    icon: '🗄️',
    title: 'Database Setup',
    desc: 'PostgreSQL, MySQL, MongoDB, Redis docker-compose configs.',
    color: 'text-arch-yellow',
  },
  {
    icon: '🔌',
    title: 'Plugin System',
    desc: 'Community plugins and remote template registries via Git.',
    color: 'text-arch-red',
  },
];

const asciiLogo = `  ███████╗██╗  ██╗ ██████╗ 
  ██╔════╝╚██╗██╔╝██╔═══██╗
  █████╗   ╚███╔╝ ██║   ██║
  ██╔══╝   ██╔██╗ ██║   ██║
  ███████╗██╔╝ ██╗╚██████╔╝
  ╚══════╝╚═╝  ╚═╝ ╚═════╝`;

const logoColors = ['#00f5d4', '#00bbf9', '#9b5de5', '#f15bb5', '#fee440', '#ff9f1c'];

export default function Home() {
  const [showContent, setShowContent] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => setShowContent(true), 500);
    return () => clearTimeout(timer);
  }, []);

  return (
    <main className="min-h-screen">
      <Navbar />

      {/* ── Hero Section ───────────────────────────────────────────────── */}
      <section className="relative pt-14 min-h-screen flex flex-col items-center justify-center px-4 grid-bg overflow-hidden">
        {/* Radial gradient overlay */}
        <div className="absolute inset-0 bg-gradient-to-b from-transparent via-transparent to-arch-bg pointer-events-none" />
        <div className="absolute top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-arch-cyan/5 rounded-full blur-[120px] pointer-events-none" />

        <div className={`relative z-10 max-w-4xl w-full transition-all duration-1000 ${showContent ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'}`}>
          {/* ASCII Logo */}
          <div className="text-center mb-6">
            <pre className="inline-block text-left text-[10px] sm:text-xs md:text-sm leading-tight select-none">
              {asciiLogo.split('\n').map((line, i) => (
                <span key={i} style={{ color: logoColors[i % logoColors.length] }} className="block font-bold">
                  {line}
                </span>
              ))}
            </pre>
          </div>

          {/* Tagline */}
          <div className="text-center mb-8">
            <h1 className="text-xl sm:text-2xl md:text-3xl font-bold mb-3">
              <span className="text-arch-text-bright">Ship Infrastructure,</span>{' '}
              <span className="text-arch-cyan crt-glow">Not YAML</span>
            </h1>
            <p className="text-arch-text-dim text-sm sm:text-base max-w-2xl mx-auto">
              The Cloud-Native Bootstrap CLI — From source code to production infrastructure in seconds.
            </p>
          </div>

          {/* Terminal Demo */}
          <div className="terminal-window max-w-3xl mx-auto mb-8">
            <div className="terminal-titlebar">
              <div className="terminal-dot terminal-dot-red" />
              <div className="terminal-dot terminal-dot-yellow" />
              <div className="terminal-dot terminal-dot-green" />
              <span className="text-xs text-arch-text-dim ml-3 font-mono">
                harsh@arch ~ /my-service
              </span>
            </div>
            <TerminalTyping lines={terminalLines} speed={35} />
          </div>

          {/* CTA Buttons */}
          <div className="flex flex-col sm:flex-row items-center justify-center gap-3 mb-8">
            <Link
              href="/docs/quickstart"
              className="group px-6 py-2.5 bg-arch-cyan/10 border border-arch-cyan/40 text-arch-cyan rounded-md text-sm font-semibold hover:bg-arch-cyan/20 hover:border-arch-cyan transition-all"
            >
              <span className="mr-2 text-arch-green">❯</span>
              Quick Start
              <span className="ml-2 group-hover:translate-x-1 inline-block transition-transform">→</span>
            </Link>
            <Link
              href="/docs"
              className="px-6 py-2.5 border border-arch-border text-arch-text-dim rounded-md text-sm hover:border-arch-border-bright hover:text-arch-text transition-all"
            >
              Read the Docs
            </Link>
            <a
              href="https://github.com/Harsh-BH/Exo"
              target="_blank"
              rel="noopener noreferrer"
              className="px-6 py-2.5 border border-arch-border text-arch-text-dim rounded-md text-sm hover:border-arch-border-bright hover:text-arch-text transition-all"
            >
              GitHub ↗
            </a>
          </div>

          {/* Install one-liner */}
          <div className="text-center">
            <div className="inline-flex items-center gap-2 px-4 py-2 bg-arch-surface border border-arch-border rounded-md text-sm">
              <span className="text-arch-green">$</span>
              <code className="text-arch-text">curl -sSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash</code>
            </div>
          </div>
        </div>
      </section>

      {/* ── Features Section ───────────────────────────────────────────── */}
      <section className="relative py-20 px-4">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-12">
            <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-2">── CAPABILITIES ──</p>
            <h2 className="text-2xl font-bold text-arch-text-bright mb-3">
              Everything You Need to Ship
            </h2>
            <p className="text-arch-text-dim text-sm max-w-xl mx-auto">
              One CLI generates production-ready DevOps scaffolding for your entire stack.
            </p>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
            {features.map((f, i) => (
              <div
                key={i}
                className="glow-card p-5 animate-fade-in"
                style={{ animationDelay: `${i * 100}ms` }}
              >
                <div className={`text-2xl mb-3 ${f.color}`}>{f.icon}</div>
                <h3 className="text-sm font-bold text-arch-text-bright mb-2">{f.title}</h3>
                <p className="text-xs text-arch-text-dim leading-relaxed">{f.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* ── Architecture Section ───────────────────────────────────────── */}
      <section className="py-20 px-4 border-t border-arch-border">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-10">
            <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-2">── ARCHITECTURE ──</p>
            <h2 className="text-2xl font-bold text-arch-text-bright">How EXO Works</h2>
          </div>
          <div className="terminal-window">
            <div className="terminal-titlebar">
              <div className="terminal-dot terminal-dot-red" />
              <div className="terminal-dot terminal-dot-yellow" />
              <div className="terminal-dot terminal-dot-green" />
              <span className="text-xs text-arch-text-dim ml-3">architecture.txt</span>
            </div>
            <pre className="terminal-body text-xs leading-relaxed text-arch-text-dim">
{`  ┌──────────────────────────────────────────────────────┐
  │                    Source Code                        │
  └──────────────────────┬───────────────────────────────┘
                         │
                         ▼
            ┌────────────────────────┐
            │  `}<span className="text-arch-cyan">Stack Detection</span>{`       │  ← go.mod · package.json · requirements.txt
            └────────────┬───────────┘
                         │
                         ▼
            ┌────────────────────────┐
            │  `}<span className="text-arch-purple">Interactive Wizard</span>{`    │  ← Bubble Tea TUI
            └────────────┬───────────┘
                         │
                         ▼
            ┌────────────────────────┐
            │  `}<span className="text-arch-blue">Template Engine</span>{`       │  ← Go text/template + embed.FS
            └────────────┬───────────┘
                         │
          ┌──────────────┼──────────────────┐
          │              │                  │
          ▼              ▼                  ▼
    ┌───────────┐  ┌───────────┐  ┌──────────────┐
    │ `}<span className="text-arch-green">Dockerfile</span>{` │  │ `}<span className="text-arch-orange">Terraform</span>{` │  │  `}<span className="text-arch-pink">CI/CD</span>{`       │
    │ `}<span className="text-arch-green">K8s YAML</span>{`   │  │ `}<span className="text-arch-orange">(AWS/GCP/</span>{` │  │  `}<span className="text-arch-pink">Monitoring</span>{`  │
    │ `}<span className="text-arch-green">Helm</span>{`       │  │ `}<span className="text-arch-orange"> Azure)</span>{`  │  │  `}<span className="text-arch-pink">Database</span>{`    │
    └───────────┘  └───────────┘  └──────────────┘
          │              │                  │
          └──────────────┼──────────────────┘
                         │
                         ▼
            ┌────────────────────────┐
            │  `}<span className="text-arch-yellow">Production-Ready Repo</span>{` │  ← .exo.yaml persisted
            └────────────────────────┘`}
            </pre>
          </div>
        </div>
      </section>

      {/* ── Tech Stack Section ─────────────────────────────────────────── */}
      <section className="py-20 px-4 border-t border-arch-border">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-10">
            <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-2">── TECH STACK ──</p>
            <h2 className="text-2xl font-bold text-arch-text-bright">Built With</h2>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
            {[
              { name: 'Go 1.25', desc: 'CLI Language', color: 'border-arch-cyan' },
              { name: 'Cobra', desc: 'CLI Framework', color: 'border-arch-blue' },
              { name: 'Bubble Tea', desc: 'Terminal UI', color: 'border-arch-pink' },
              { name: 'Lipgloss', desc: 'Terminal Style', color: 'border-arch-purple' },
              { name: 'Terraform', desc: 'IaC Target', color: 'border-arch-orange' },
              { name: 'Docker', desc: 'Containers', color: 'border-arch-blue' },
              { name: 'Kubernetes', desc: 'Orchestration', color: 'border-arch-cyan' },
              { name: 'Prometheus', desc: 'Monitoring', color: 'border-arch-red' },
            ].map((tech, i) => (
              <div key={i} className={`p-4 bg-arch-surface border-l-2 ${tech.color} rounded-r-md`}>
                <div className="text-sm font-bold text-arch-text-bright">{tech.name}</div>
                <div className="text-xs text-arch-text-dim">{tech.desc}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* ── Footer ─────────────────────────────────────────────────────── */}
      <footer className="border-t border-arch-border py-10 px-4">
        <div className="max-w-4xl mx-auto text-center">
          <p className="text-arch-text-dim text-xs mb-2">
            Built by{' '}
            <a href="https://github.com/Harsh-BH" className="text-arch-cyan hover:underline">
              Harsh Bhatt
            </a>
          </p>
          <p className="text-arch-text-dim text-[10px]">
            Powered by Go · Cobra · Bubble Tea · Terraform · Docker · Kubernetes
          </p>
          <p className="text-arch-text-dim text-[10px] mt-1">MIT License © 2026</p>
        </div>
      </footer>
    </main>
  );
}
