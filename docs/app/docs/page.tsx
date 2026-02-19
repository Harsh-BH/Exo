import Link from 'next/link';
import CodeBlock from '@/components/CodeBlock';

export default function DocsHome() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── DOCUMENTATION ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Introduction</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          EXO is a single CLI that detects your project stack and generates production-ready DevOps
          scaffolding — Dockerfiles, Terraform, CI/CD, Kubernetes YAML, Helm charts, and monitoring — in seconds.
        </p>
      </div>

      {/* The Problem */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-red">##</span> The Problem
        </h2>
        <p className="text-sm text-arch-text leading-relaxed mb-4">
          Every new microservice demands the same boilerplate: Dockerfiles, Terraform modules, CI/CD pipelines,
          Kubernetes manifests, monitoring configs. It&apos;s repetitive, error-prone, and burns hours that should
          go toward building your product.
        </p>
      </section>

      {/* The Solution */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-green">##</span> The Solution
        </h2>
        <p className="text-sm text-arch-text leading-relaxed mb-4">
          You write the application code. EXO handles the infrastructure. One command scaffolds everything
          you need for a production deployment.
        </p>
        <CodeBlock
          language="bash"
          code={`$ cd ~/my-service
$ exo init
# → Dockerfile, Terraform, CI/CD, K8s, Monitoring generated!`}
        />
      </section>

      {/* Features */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4 flex items-center gap-2">
          <span className="text-arch-cyan">##</span> Key Features
        </h2>
        <table className="docs-table">
          <thead>
            <tr>
              <th>Feature</th>
              <th>What It Does</th>
            </tr>
          </thead>
          <tbody>
            <tr><td className="text-arch-cyan font-semibold">Smart Stack Detection</td><td>Auto-detects Go, Node.js, Python, Java, Rust from source files</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Multi-Stage Dockerfiles</td><td>Generates optimized, language-specific Dockerfiles</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Multi-Cloud Terraform</td><td>Scaffolds IaC modules for AWS, GCP, and Azure</td></tr>
            <tr><td className="text-arch-cyan font-semibold">CI/CD Pipelines</td><td>Generates GitHub Actions and GitLab CI workflows</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Kubernetes Manifests</td><td>Produces Deployment, Service, and Ingress YAML</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Helm Charts</td><td>Full Helm chart with values, templates, and ingress</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Monitoring Stack</td><td>Sets up Prometheus + Grafana with dashboards and alerts</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Database Setup</td><td>PostgreSQL, MySQL, MongoDB, Redis docker-compose configs</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Interactive Wizard</td><td>Guided setup via beautiful Bubble Tea terminal UI</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Config Persistence</td><td>Saves config in .exo.yaml for repeatable runs</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Plugin System</td><td>Community plugins and remote template registries</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Secret Scanner</td><td>Magic-byte binary skip + tightened allowlist (never suppresses password/secret)</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Security Validation</td><td><code>exo validate --security</code> runs Trivy on generated Dockerfiles</td></tr>
            <tr><td className="text-arch-cyan font-semibold">SBOM Generation</td><td><code>exo gen sbom</code> — CycloneDX JSON, uses syft if installed</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Diff Preview</td><td><code>exo diff</code> shows a unified diff before you overwrite files</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Clean Command</td><td><code>exo clean</code> removes all generated assets in one step</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Git Bootstrap</td><td><code>exo init --from-git &lt;url&gt;</code> clones and auto-scaffolds any repo</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Output Dir</td><td><code>--output-dir</code> writes generated files to a custom directory</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Dry Run</td><td><code>--dry-run</code> previews file list without writing anything</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Braille Spinner</td><td>Animated progress indicator during long-running generation</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Extended Detection</td><td>Java (pom.xml / build.gradle), Rust (Cargo.toml), 10+ frameworks</td></tr>
            <tr><td className="text-arch-cyan font-semibold">Cross-Platform</td><td>Pre-built binaries for Linux, macOS (Intel &amp; ARM), Windows</td></tr>
          </tbody>
        </table>
      </section>

      {/* Next Steps */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4 flex items-center gap-2">
          <span className="text-arch-purple">##</span> Next Steps
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
          <Link href="/docs/installation" className="glow-card p-4 block">
            <div className="text-arch-cyan font-semibold text-sm mb-1">Installation →</div>
            <div className="text-xs text-arch-text-dim">Install EXO on your system</div>
          </Link>
          <Link href="/docs/quickstart" className="glow-card p-4 block">
            <div className="text-arch-green font-semibold text-sm mb-1">Quick Start →</div>
            <div className="text-xs text-arch-text-dim">Generate your first project in 60 seconds</div>
          </Link>
          <Link href="/docs/architecture" className="glow-card p-4 block">
            <div className="text-arch-purple font-semibold text-sm mb-1">Architecture →</div>
            <div className="text-xs text-arch-text-dim">Understand how EXO works under the hood</div>
          </Link>
          <Link href="/docs/cli/init" className="glow-card p-4 block">
            <div className="text-arch-orange font-semibold text-sm mb-1">CLI Reference →</div>
            <div className="text-xs text-arch-text-dim">Full command documentation</div>
          </Link>
        </div>
      </section>
    </article>
  );
}
