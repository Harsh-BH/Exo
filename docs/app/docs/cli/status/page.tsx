import CodeBlock from '@/components/CodeBlock';

export default function CLIStatusPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo status</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Display an overview of all generated infrastructure files and their status.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo status`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Reads <code className="text-arch-pink">.exo.yaml</code> and scans for generated files.
          Shows a formatted table with file names, sizes, modification times, and presence status.
          Uses <span className="text-arch-green">✓</span> for files found and{' '}
          <span className="text-arch-red">✗</span> for missing files.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example Output</h2>
        <CodeBlock language="text" code={`$ exo status
EXO Project Status

  ✓  Dockerfile                 1.2 KB  •  2 hours ago
  ✓  .exo.yaml                  0.1 KB  •  2 hours ago
  ✓  Terraform (AWS)            dir     •  2 hours ago
  ✓  GitHub Actions             dir     •  2 hours ago
  ✓  K8s Manifests              dir     •  2 hours ago
  ✓  Helm Chart                 dir     •  2 hours ago
  ✓  Monitoring                 dir     •  2 hours ago
  ✓  DB (PostgreSQL)            0.5 KB  •  2 hours ago
  ✓  Makefile                   0.8 KB  •  2 hours ago
  ✓  .env.example               0.2 KB  •  2 hours ago
  ✓  .gitignore                 0.4 KB  •  2 hours ago
  ✓  Grafana Dashboard          4.1 KB  •  2 hours ago
  ✓  Alert Rules                0.9 KB  •  2 hours ago`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Checked Paths</h2>
        <p className="text-sm text-arch-text-dim">
          The command checks these paths based on your config: Dockerfile, .exo.yaml, infra/&lt;provider&gt;/,
          .github/workflows/ or .gitlab-ci.yml, k8s/, charts/&lt;name&gt;/, monitoring/,
          docker-compose.&lt;db&gt;.yml, Makefile, .env.example, .gitignore, grafana_dashboard.json, alerts.yml.
        </p>
      </section>
    </article>
  );
}
