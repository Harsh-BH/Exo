import CodeBlock from '@/components/CodeBlock';

export default function CLIInitPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo init</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Initialize a new Exo project with an interactive wizard or non-interactive flags.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-green mr-2">##</span> Synopsis
        </h2>
        <CodeBlock language="bash" code={`exo init [flags]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-green mr-2">##</span> Description
        </h2>
        <p className="text-sm text-arch-text mb-4">
          The <code className="text-arch-pink">init</code> command is the primary entry point for bootstrapping a project.
          It performs three steps:
        </p>
        <div className="space-y-2 text-sm">
          <div className="flex gap-2"><span className="text-arch-cyan">1.</span><span className="text-arch-text-dim"><strong className="text-arch-text-bright">Detect</strong> — Scans the current directory for language markers (go.mod, package.json, etc.)</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">2.</span><span className="text-arch-text-dim"><strong className="text-arch-text-bright">Prompt</strong> — Launches a 7-step Bubble Tea wizard for project configuration</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">3.</span><span className="text-arch-text-dim"><strong className="text-arch-text-bright">Render</strong> — Generates all selected infrastructure files via templates</span></div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-green mr-2">##</span> Flags
        </h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Flag</th><th>Type</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>--non-interactive</code></td><td>bool</td><td>Skip the wizard, use flags for config</td></tr>
            <tr><td><code>--name</code></td><td>string</td><td>Project name</td></tr>
            <tr><td><code>--lang</code></td><td>string</td><td>Language: go, node, python, java, rust</td></tr>
            <tr><td><code>--provider</code></td><td>string</td><td>Cloud: aws, gcp, azure, none</td></tr>
            <tr><td><code>--ci</code></td><td>string</td><td>CI: github-actions, gitlab-ci, none</td></tr>
            <tr><td><code>--monitoring</code></td><td>string</td><td>Monitoring: prometheus, none</td></tr>
            <tr><td><code>--db</code></td><td>string</td><td>Database: postgres, mysql, mongo, redis, none</td></tr>
            <tr><td><code>--from-git</code></td><td>string</td><td>Clone a git URL and auto-detect + scaffold the cloned repo</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-green mr-2">##</span> Examples
        </h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Interactive mode</p>
            <CodeBlock language="bash" code={`exo init`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Non-interactive (CI/CD)</p>
            <CodeBlock language="bash" code={`exo init --non-interactive \\
  --name=api-gateway \\
  --lang=go \\
  --provider=aws \\
  --ci=github-actions \\
  --monitoring=prometheus \\
  --db=postgres`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Bootstrap from an existing git repository</p>
            <CodeBlock language="bash" code={`# Clone repo, auto-detect stack, launch wizard pre-filled
exo init --from-git https://github.com/org/my-service

# Fully non-interactive from git
exo init --from-git https://github.com/org/my-service \\
  --non-interactive --name=my-service --provider=aws`} />
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3">
          <span className="text-arch-green mr-2">##</span> Generated Files
        </h2>
        <p className="text-sm text-arch-text-dim mb-4">
          Depending on selections, init generates: <code className="text-arch-pink">.exo.yaml</code>,{' '}
          <code className="text-arch-pink">Dockerfile</code>,{' '}
          <code className="text-arch-pink">infra/</code>,{' '}
          <code className="text-arch-pink">k8s/</code>,{' '}
          <code className="text-arch-pink">charts/</code>,{' '}
          <code className="text-arch-pink">.github/workflows/</code>,{' '}
          <code className="text-arch-pink">monitoring/</code>,{' '}
          <code className="text-arch-pink">docker-compose.*.yml</code>,{' '}
          <code className="text-arch-pink">Makefile</code>,{' '}
          <code className="text-arch-pink">.env.example</code>,{' '}
          <code className="text-arch-pink">.gitignore</code>,{' '}
          <code className="text-arch-pink">grafana_dashboard.json</code>, and{' '}
          <code className="text-arch-pink">alerts.yml</code>.
        </p>
      </section>
    </article>
  );
}
