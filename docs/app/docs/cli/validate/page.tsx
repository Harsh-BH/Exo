import CodeBlock from '@/components/CodeBlock';

export default function CLIValidatePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo validate</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Validate the .exo.yaml configuration file against the expected schema.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo validate`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Reads <code className="text-arch-pink">.exo.yaml</code> and checks all field values against
          the set of allowed options. Reports errors for invalid languages, providers, CI platforms,
          monitoring options, or database types.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Validation Rules</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Field</th><th>Valid Values</th></tr></thead>
          <tbody>
            <tr><td><code>name</code></td><td>Non-empty string</td></tr>
            <tr><td><code>language</code></td><td><code>go</code>, <code>node</code>, <code>python</code></td></tr>
            <tr><td><code>provider</code></td><td><code>aws</code>, <code>gcp</code>, <code>azure</code>, <code>none</code></td></tr>
            <tr><td><code>ci</code></td><td><code>github-actions</code>, <code>gitlab-ci</code>, <code>none</code></td></tr>
            <tr><td><code>monitoring</code></td><td><code>prometheus</code>, <code>none</code></td></tr>
            <tr><td><code>database</code></td><td><code>postgres</code>, <code>mysql</code>, <code>mongo</code>, <code>redis</code>, <code>none</code></td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example</h2>
        <CodeBlock language="bash" code={`$ exo validate
Validating .exo.yaml...

  ✓ name: my-service
  ✓ language: go
  ✓ provider: aws
  ✓ ci: github-actions
  ✓ monitoring: prometheus
  ✓ database: postgres

Configuration is valid.`} />
      </section>
    </article>
  );
}
