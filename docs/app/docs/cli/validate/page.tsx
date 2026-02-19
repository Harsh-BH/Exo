import CodeBlock from '@/components/CodeBlock';

export default function CLIValidatePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo validate</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Validate the .exo.yaml config and optionally run a Trivy security scan on the generated Dockerfile.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo validate [--security]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Without flags, reads <code className="text-arch-pink">.exo.yaml</code> and checks all field
          values against the set of allowed options. Reports errors for invalid languages, providers,
          CI platforms, monitoring options, or database types.
        </p>
        <p className="text-sm text-arch-text mb-4">
          With <code className="text-arch-pink">--security</code>, additionally invokes{' '}
          <a href="https://github.com/aquasecurity/trivy" className="text-arch-cyan underline" target="_blank" rel="noreferrer">Trivy</a>{' '}
          (if installed) to scan <code className="text-arch-pink">Dockerfile</code> for known CVEs,
          outdated base images, and misconfiguration issues.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Flags</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Flag</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>--security</code></td><td>Run Trivy security scan on Dockerfile after config validation</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Validation Rules</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Field</th><th>Valid Values</th></tr></thead>
          <tbody>
            <tr><td><code>name</code></td><td>Non-empty string</td></tr>
            <tr><td><code>language</code></td><td><code>go</code>, <code>node</code>, <code>python</code>, <code>java</code>, <code>rust</code></td></tr>
            <tr><td><code>provider</code></td><td><code>aws</code>, <code>gcp</code>, <code>azure</code>, <code>none</code></td></tr>
            <tr><td><code>ci</code></td><td><code>github-actions</code>, <code>gitlab-ci</code>, <code>none</code></td></tr>
            <tr><td><code>monitoring</code></td><td><code>prometheus</code>, <code>none</code></td></tr>
            <tr><td><code>database</code></td><td><code>postgres</code>, <code>mysql</code>, <code>mongo</code>, <code>redis</code>, <code>none</code></td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Config validation only</p>
            <CodeBlock language="bash" code={`$ exo validate
Validating .exo.yaml...

  ✓ name: my-service
  ✓ language: go
  ✓ provider: aws
  ✓ ci: github-actions
  ✓ monitoring: prometheus
  ✓ database: postgres

Configuration is valid.`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">With Trivy security scan</p>
            <CodeBlock language="bash" code={`$ exo validate --security
Validating .exo.yaml...  ✓ valid

Running Trivy security scan on Dockerfile...

  2024-01-15T10:23:45Z INFO  Scanning Dockerfile...
  Dockerfile (dockerfile)
  =======================
  Tests: 24 (SUCCESSES: 22, FAILURES: 2, EXCEPTIONS: 0)
  FAILURES: 2 (HIGH: 1, MEDIUM: 1)

  HIGH: Specify a tag in the FROM statement
  MEDIUM: Last USER command is root`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">When Trivy is not installed</p>
            <CodeBlock language="bash" code={`$ exo validate --security
Validating .exo.yaml...  ✓ valid

[warn] trivy not found in PATH — skipping security scan.
       Install: https://github.com/aquasecurity/trivy`} />
          </div>
        </div>
      </section>
    </article>
  );
}
