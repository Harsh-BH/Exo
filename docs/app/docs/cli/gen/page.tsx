import CodeBlock from '@/components/CodeBlock';

export default function CLIGenPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo gen</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Generate individual infrastructure components from templates.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo gen <type> [flags]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Unlike <code className="text-arch-pink">exo init</code> which generates everything at once,{' '}
          <code className="text-arch-pink">exo gen</code> lets you generate specific components individually.
          Requires <code className="text-arch-pink">.exo.yaml</code> to exist (run <code className="text-arch-pink">exo init</code> first).
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Subcommands</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Subcommand</th><th>Flags</th><th>Output</th></tr></thead>
          <tbody>
            <tr><td><code>docker</code></td><td>—</td><td>Dockerfile (language-specific)</td></tr>
            <tr><td><code>infra</code></td><td><code>--provider</code></td><td>infra/&lt;provider&gt;/ (Terraform)</td></tr>
            <tr><td><code>k8s</code></td><td>—</td><td>k8s/ directory</td></tr>
            <tr><td><code>ci</code></td><td>—</td><td>.github/workflows/ or .gitlab-ci.yml</td></tr>
            <tr><td><code>helm</code></td><td><code>--name</code></td><td>charts/&lt;name&gt;/</td></tr>
            <tr><td><code>monitoring</code></td><td>—</td><td>monitoring/ directory</td></tr>
            <tr><td><code>db</code></td><td><code>--db</code></td><td>docker-compose.&lt;db&gt;.yml</td></tr>
            <tr><td><code>makefile</code></td><td>—</td><td>Makefile</td></tr>
            <tr><td><code>env</code></td><td>—</td><td>.env.example</td></tr>
            <tr><td><code>gitignore</code></td><td>—</td><td>.gitignore</td></tr>
            <tr><td><code>grafana</code></td><td>—</td><td>grafana_dashboard.json</td></tr>
            <tr><td><code>alerts</code></td><td>—</td><td>alerts.yml</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <CodeBlock language="bash" code={`# Generate a Dockerfile based on detected language
exo gen docker

# Generate Terraform for AWS
exo gen infra --provider aws

# Generate Helm chart with custom name
exo gen helm --name my-api

# Generate PostgreSQL docker-compose
exo gen db --db postgres

# Generate all Kubernetes manifests
exo gen k8s

# Generate monitoring stack
exo gen monitoring

# Generate CI pipeline (uses CI setting from .exo.yaml)
exo gen ci`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Flags</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Flag</th><th>Applies To</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>--provider</code></td><td><code>infra</code></td><td>Cloud provider: aws, gcp, azure</td></tr>
            <tr><td><code>--name</code></td><td><code>helm</code></td><td>Helm chart release name</td></tr>
            <tr><td><code>--db</code></td><td><code>db</code></td><td>Database type: postgres, mysql, mongo, redis</td></tr>
          </tbody>
        </table>
      </section>
    </article>
  );
}
