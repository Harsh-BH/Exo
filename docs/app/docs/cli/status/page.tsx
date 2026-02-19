import CodeBlock from '@/components/CodeBlock';

export default function CLIStatusPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo status</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Display a grouped tree view of all generated infrastructure assets and their status.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo status`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Reads <code className="text-arch-pink">.exo.yaml</code> and scans for the full set of
          expected generated assets. Output is a tree grouped into 7 categories with file sizes,
          modification times, and a summary progress bar at the bottom.
        </p>
        <p className="text-sm text-arch-text mb-4">
          Uses <span className="text-arch-green font-bold">✓</span> for files/directories present
          and <span className="text-arch-red font-bold">✗</span> for missing items.
          Tree connectors (<code className="text-arch-pink">├──</code> / <code className="text-arch-pink">└──</code>)
          give a clear visual hierarchy.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example Output</h2>
        <CodeBlock language="text" code={`$ exo status
EXO Project Status ─ my-service

  Containers
  ├── ✓  Dockerfile          1.2 KB  •  2h ago
  └── ✓  docker-compose      0.8 KB  •  2h ago

  CI/CD
  ├── ✓  GitHub Actions      dir     •  2h ago
  └── ✗  GitLab CI           missing

  Infrastructure
  ├── ✓  Terraform (AWS)     dir     •  2h ago
  └── ✗  Terraform (GCP)     missing

  Kubernetes
  ├── ✓  K8s Manifests       dir     •  2h ago
  └── ✓  Helm Chart          dir     •  2h ago

  Monitoring
  ├── ✓  Monitoring Stack    dir     •  2h ago
  ├── ✓  Grafana Dashboard   4.1 KB  •  2h ago
  └── ✓  Alert Rules         0.9 KB  •  2h ago

  Database
  └── ✓  PostgreSQL          0.5 KB  •  2h ago

  Project
  ├── ✓  .exo.yaml           0.1 KB  •  2h ago
  ├── ✓  Makefile            0.8 KB  •  2h ago
  ├── ✓  .env.example        0.2 KB  •  2h ago
  └── ✓  .gitignore          0.4 KB  •  2h ago

  ████████████████░░  14 / 16 assets present`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Asset Categories</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Category</th><th>Assets Checked</th></tr></thead>
          <tbody>
            <tr><td><strong>Containers</strong></td><td>Dockerfile, docker-compose.yml</td></tr>
            <tr><td><strong>CI/CD</strong></td><td>.github/workflows/, .gitlab-ci.yml</td></tr>
            <tr><td><strong>Infrastructure</strong></td><td>infra/aws/, infra/gcp/, infra/azure/</td></tr>
            <tr><td><strong>Kubernetes</strong></td><td>k8s/, charts/&lt;name&gt;/</td></tr>
            <tr><td><strong>Monitoring</strong></td><td>monitoring/, grafana_dashboard.json, alerts.yml</td></tr>
            <tr><td><strong>Database</strong></td><td>docker-compose.&lt;db&gt;.yml</td></tr>
            <tr><td><strong>Project</strong></td><td>.exo.yaml, Makefile, .env.example, .gitignore</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Tips</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Run after <code className="text-arch-pink">exo init</code> to verify all assets were generated</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Use alongside <code className="text-arch-pink">exo diff</code> to preview what would change before regenerating</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Combine with <code className="text-arch-pink">exo clean</code> to remove all assets when starting fresh</span></div>
        </div>
      </section>
    </article>
  );
}
