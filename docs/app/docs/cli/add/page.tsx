import CodeBlock from '@/components/CodeBlock';

export default function CLIAddPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo add</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Add new infrastructure components to an existing Exo project.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo add <type> [flags]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Adds a component to your existing project. Updates <code className="text-arch-pink">.exo.yaml</code> and
          generates the corresponding files. Unlike <code className="text-arch-pink">exo gen</code>,
          this also modifies your config so <code className="text-arch-pink">exo status</code> tracks the new component.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Subcommands</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Subcommand</th><th>Flags</th><th>Effect</th></tr></thead>
          <tbody>
            <tr><td><code>db</code></td><td><code>--db</code></td><td>Adds database config + generates docker-compose</td></tr>
            <tr><td><code>ci</code></td><td><code>--ci</code></td><td>Adds CI config + generates pipeline</td></tr>
            <tr><td><code>monitoring</code></td><td>—</td><td>Adds monitoring config + generates prometheus/grafana</td></tr>
            <tr><td><code>infra</code></td><td><code>--provider</code></td><td>Adds provider config + generates terraform</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <CodeBlock language="bash" code={`# Add PostgreSQL to existing project
exo add db --db postgres

# Add GitHub Actions CI
exo add ci --ci github-actions

# Add monitoring stack
exo add monitoring

# Add Azure infrastructure
exo add infra --provider azure`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Workflow</h2>
        <div className="p-4 bg-arch-surface border border-arch-border rounded-md text-xs text-arch-text-dim font-mono">
          <pre>{`exo add db --db redis
  │
  ├─▸ Read .exo.yaml
  ├─▸ Update database: "redis"
  ├─▸ Write .exo.yaml
  ├─▸ Render db/redis.tmpl → docker-compose.redis.yml
  └─▸ Done ✓`}</pre>
        </div>
      </section>
    </article>
  );
}
