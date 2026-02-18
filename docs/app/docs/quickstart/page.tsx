import CodeBlock from '@/components/CodeBlock';

export default function QuickStartPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GETTING STARTED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Quick Start</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Generate a complete production infrastructure for your project in under 60 seconds.
        </p>
      </div>

      {/* Step 1 */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-cyan mr-2">01.</span> Navigate to Your Project
        </h2>
        <CodeBlock language="bash" code={`cd ~/my-service`} />
      </section>

      {/* Step 2 */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-cyan mr-2">02.</span> Run the Interactive Wizard
        </h2>
        <p className="text-sm text-arch-text mb-4">
          The wizard walks you through selecting your language, cloud provider, CI/CD platform,
          monitoring stack, and database — then generates everything in one pass.
        </p>
        <CodeBlock language="bash" code={`exo init`} />
        <div className="mt-4 p-4 bg-arch-surface border border-arch-border rounded-md text-sm">
          <p className="text-arch-text-bright font-semibold mb-2">The wizard will ask you:</p>
          <div className="space-y-1 text-arch-text-dim text-xs">
            <p><span className="text-arch-pink font-bold mr-2">1.</span> Project name</p>
            <p><span className="text-arch-pink font-bold mr-2">2.</span> Language — Go, Node.js, or Python</p>
            <p><span className="text-arch-pink font-bold mr-2">3.</span> Cloud provider — AWS, GCP, Azure, or None</p>
            <p><span className="text-arch-pink font-bold mr-2">4.</span> CI/CD — GitHub Actions, GitLab CI, or None</p>
            <p><span className="text-arch-pink font-bold mr-2">5.</span> Monitoring — Prometheus + Grafana, or None</p>
            <p><span className="text-arch-pink font-bold mr-2">6.</span> Database — PostgreSQL, MySQL, MongoDB, Redis, or None</p>
            <p><span className="text-arch-pink font-bold mr-2">7.</span> Confirmation — Review and generate</p>
          </div>
        </div>
      </section>

      {/* Step 3 */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-cyan mr-2">03.</span> Or Generate Individual Components
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Instead of the full wizard, you can generate specific components individually:
        </p>
        <CodeBlock language="bash" code={`exo gen docker                  # Dockerfile
exo gen infra --provider aws    # Terraform modules
exo gen k8s                     # Kubernetes manifests
exo gen ci                      # CI/CD pipeline
exo gen helm --name myapp       # Helm chart
exo gen db --db postgres        # Database docker-compose
exo gen makefile                # Makefile
exo gen env                     # .env.example
exo gen gitignore               # .gitignore
exo gen grafana                 # Grafana dashboard
exo gen alerts                  # Prometheus alerting rules`} />
      </section>

      {/* Step 4 */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-cyan mr-2">04.</span> Check What&apos;s Been Generated
        </h2>
        <CodeBlock language="bash" code={`$ exo status
EXO Project Status

  ✓  Dockerfile                 1.2 KB  •  just now
  ✓  Terraform (AWS)            dir     •  just now
  ✓  GitHub Actions             dir     •  just now
  ✓  K8s Manifests              dir     •  just now
  ✓  Monitoring                 dir     •  just now
  ✓  DB (PostgreSQL)            0.5 KB  •  just now
  ✓  EXO Config                 0.1 KB  •  just now`} />
      </section>

      {/* Non-interactive */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-cyan mr-2">05.</span> Non-Interactive Mode (CI/CD)
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Skip the wizard entirely using flags — perfect for CI/CD pipelines:
        </p>
        <CodeBlock language="bash" code={`exo init \\
  --non-interactive \\
  --name=my-service \\
  --lang=go \\
  --provider=aws \\
  --ci=github-actions \\
  --monitoring=prometheus \\
  --db=postgres`} />
      </section>

      {/* Generated Output */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-green">##</span> Generated Output Structure
        </h2>
        <CodeBlock language="text" filename="Project Tree" showLineNumbers code={`my-service/
├── Dockerfile                          # Optimized multi-stage build
├── .exo.yaml                           # EXO configuration
├── Makefile                            # Build, test, run, deploy commands
├── .env.example                        # Environment template
├── .gitignore                          # Language-aware gitignore
├── .github/workflows/
│   └── go.yml                          # GitHub Actions CI pipeline
├── infra/
│   └── aws/
│       ├── main.tf                     # VPC + EKS resources
│       ├── variables.tf                # Input variables
│       └── provider.tf                 # Provider configuration
├── k8s/
│   ├── deployment.yaml                 # Kubernetes Deployment
│   ├── service.yaml                    # Kubernetes Service
│   └── ingress.yaml                    # Kubernetes Ingress
├── charts/
│   └── my-service/
│       ├── Chart.yaml                  # Helm chart metadata
│       ├── values.yaml                 # Helm values
│       └── templates/
│           ├── deployment.yaml         # Helm deployment template
│           ├── service.yaml            # Helm service template
│           └── ingress.yaml            # Helm ingress template
├── monitoring/
│   ├── prometheus.yml                  # Prometheus scrape config
│   └── docker-compose.monitoring.yml   # Prometheus + Grafana stack
├── grafana_dashboard.json              # Pre-built Grafana dashboard
├── alerts.yml                          # Prometheus alerting rules
└── docker-compose.postgres.yml         # Database setup`} />
      </section>
    </article>
  );
}
