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
          <code className="text-arch-pink">exo gen</code> lets you add or regenerate specific components individually.
          It reads <code className="text-arch-pink">.exo.yaml</code> for defaults and falls back to auto-detecting
          your stack — so flags are optional in most cases.
        </p>
        <p className="text-sm text-arch-text mb-4">
          Use <code className="text-arch-pink">--dry-run</code> to preview exactly what would be written, and{' '}
          <code className="text-arch-pink">--force</code> to overwrite existing files. Combine with{' '}
          <code className="text-arch-pink">--output-dir</code> to send output to a custom directory.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generator Types</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Type</th><th>Output</th></tr></thead>
          <tbody>
            <tr><td><code>docker</code></td><td>Language-specific multi-stage Dockerfile</td></tr>
            <tr><td><code>docker-compose</code></td><td>Full docker-compose.yml (app + DB + monitoring)</td></tr>
            <tr><td><code>infra</code></td><td><code>infra/&lt;provider&gt;/</code> — Terraform main, variables, provider</td></tr>
            <tr><td><code>k8s</code></td><td><code>k8s/</code> — Deployment, Service, Ingress YAML</td></tr>
            <tr><td><code>helm</code></td><td><code>charts/&lt;name&gt;/</code> — full Helm chart with templates</td></tr>
            <tr><td><code>ci</code></td><td>Language-aware CI pipeline (GitHub Actions or GitLab CI)</td></tr>
            <tr><td><code>db</code></td><td><code>docker-compose.&lt;db&gt;.yml</code> — database service</td></tr>
            <tr><td><code>makefile</code></td><td><code>Makefile</code></td></tr>
            <tr><td><code>env</code></td><td><code>.env.example</code></td></tr>
            <tr><td><code>gitignore</code></td><td><code>.gitignore</code></td></tr>
            <tr><td><code>grafana</code></td><td><code>grafana_dashboard.json</code></td></tr>
            <tr><td><code>alerts</code></td><td><code>alerts.yml</code> — Prometheus alert rules</td></tr>
            <tr><td><code>readme</code></td><td><code>README.md</code> with badges, install and usage sections</td></tr>
            <tr><td><code>pre-commit</code></td><td><code>.pre-commit-config.yaml</code> — language-aware hooks</td></tr>
            <tr><td><code>devcontainer</code></td><td><code>.devcontainer/devcontainer.json</code></td></tr>
            <tr><td><code>renovate</code></td><td><code>renovate.json</code> — automated dependency updates</td></tr>
            <tr><td><code>license</code></td><td><code>LICENSE</code> — MIT, Apache 2.0, or GPL 3.0</td></tr>
            <tr><td><code>dependabot</code></td><td><code>.github/dependabot.yml</code></td></tr>
            <tr><td><code>sonarqube</code></td><td><code>sonar-project.properties</code></td></tr>
            <tr><td><code>sbom</code></td><td><code>sbom.cdx.json</code> — CycloneDX SBOM (uses syft if available)</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Flags</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Flag</th><th>Short</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>--lang</code></td><td><code>-l</code></td><td>Language override: go, node, python, java, rust</td></tr>
            <tr><td><code>--name</code></td><td><code>-n</code></td><td>App name (defaults to .exo.yaml or directory name)</td></tr>
            <tr><td><code>--provider</code></td><td><code>-p</code></td><td>Cloud provider override: aws, gcp, azure</td></tr>
            <tr><td><code>--db</code></td><td>—</td><td>Database override: postgres, mysql, mongo, redis</td></tr>
            <tr><td><code>--monitoring</code></td><td>—</td><td>Monitoring override: prometheus, none</td></tr>
            <tr><td><code>--dry-run</code></td><td>—</td><td>Preview what would be generated without writing files</td></tr>
            <tr><td><code>--force</code></td><td>—</td><td>Overwrite existing files without prompting</td></tr>
            <tr><td><code>--output-dir</code></td><td><code>-o</code></td><td>Write output into this directory instead of CWD</td></tr>
            <tr><td><code>--license-type</code></td><td>—</td><td>License type for <code>exo gen license</code>: mit, apache2, gpl3</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Basic usage (reads .exo.yaml automatically)</p>
            <CodeBlock language="bash" code={`exo gen docker
exo gen k8s
exo gen ci`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Preview before writing</p>
            <CodeBlock language="bash" code={`exo gen helm --dry-run
exo gen infra --provider aws --dry-run`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Overwrite existing files</p>
            <CodeBlock language="bash" code={`exo gen docker --force
exo gen k8s --force`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Output to a custom directory</p>
            <CodeBlock language="bash" code={`exo gen k8s --output-dir ./deploy/
exo gen helm -o ./charts-out/`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">New generators</p>
            <CodeBlock language="bash" code={`# Full local dev stack
exo gen docker-compose

# Project readme, license, pre-commit hooks
exo gen readme
exo gen license --license-type apache2
exo gen pre-commit

# Security and quality
exo gen sbom            # uses syft if installed, else CycloneDX skeleton
exo gen sonarqube
exo gen dependabot`} />
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Config resolution order</h2>
        <p className="text-sm text-arch-text-dim mb-3">Priority (highest wins):</p>
        <div className="space-y-1 text-sm">
          <div className="flex gap-2"><span className="text-arch-cyan">1.</span><span className="text-arch-text-dim">CLI flags (<code className="text-arch-pink">--lang</code>, <code className="text-arch-pink">--provider</code>, etc.)</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">2.</span><span className="text-arch-text-dim"><code className="text-arch-pink">.exo.yaml</code> in current directory</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">3.</span><span className="text-arch-text-dim">Auto-detector (scans go.mod, package.json, pyproject.toml, Cargo.toml, pom.xml…)</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">4.</span><span className="text-arch-text-dim">Built-in defaults (port 8080, language unknown)</span></div>
        </div>
      </section>
    </article>
  );
}
