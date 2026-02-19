import CodeBlock from '@/components/CodeBlock';

export default function CLIScanPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo scan</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Scan the project for hardcoded secrets, API keys, tokens, and sensitive values.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo scan [path]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Walks the directory tree and inspects each file for secret-like patterns — AWS keys,
          GitHub tokens, generic API keys, database passwords, private key headers, and more.
          Binary files are detected via magic-byte inspection and automatically skipped to
          avoid false positives and wasted processing.
        </p>
        <p className="text-sm text-arch-text mb-4">
          The scanner uses a tightened allowlist: variable names containing{' '}
          <code className="text-arch-pink">secret</code> or{' '}
          <code className="text-arch-pink">password</code> are <em>never</em> suppressed —
          only placeholder-value patterns (e.g. <code className="text-arch-pink">your-key-here</code>,{' '}
          <code className="text-arch-pink">xxxx</code>, <code className="text-arch-pink">changeme</code>) are.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Binary File Detection</h2>
        <p className="text-sm text-arch-text mb-4">
          Before scanning any file, Exo reads the first 512 bytes and calls{' '}
          <code className="text-arch-pink">isBinaryFile()</code> which checks for null bytes
          and known magic-byte signatures:
        </p>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Magic Bytes</th><th>Type Skipped</th></tr></thead>
          <tbody>
            <tr><td><code>0x7f 0x45 0x4c 0x46</code></td><td>ELF executables (Linux binaries)</td></tr>
            <tr><td><code>0x4d 0x5a</code></td><td>PE executables (Windows .exe)</td></tr>
            <tr><td><code>0xff 0xd8 0xff</code></td><td>JPEG images</td></tr>
            <tr><td><code>0x89 0x50 0x4e 0x47</code></td><td>PNG images</td></tr>
            <tr><td><code>0x25 0x50 0x44 0x46</code></td><td>PDF documents</td></tr>
            <tr><td><code>0x50 0x4b 0x03 0x04</code></td><td>ZIP archives (also .jar, .docx, .xlsx)</td></tr>
            <tr><td><code>0x1f 0x8b</code></td><td>Gzip archives</td></tr>
            <tr><td>Null byte (<code>0x00</code>)</td><td>Any binary file with embedded null bytes</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Detected Patterns</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Pattern</th><th>Example Match</th></tr></thead>
          <tbody>
            <tr><td>AWS Access Key</td><td><code>AKIA...</code> (20-char key)</td></tr>
            <tr><td>AWS Secret Key</td><td><code>aws_secret_access_key = ...</code></td></tr>
            <tr><td>GitHub Token</td><td><code>ghp_...</code>, <code>github_token = ...</code></td></tr>
            <tr><td>Generic API Key</td><td><code>api_key = "..."</code>, <code>apikey: ...</code></td></tr>
            <tr><td>Database password</td><td><code>db_password = "..."</code></td></tr>
            <tr><td>Private key header</td><td><code>-----BEGIN RSA PRIVATE KEY-----</code></td></tr>
            <tr><td>Bearer token</td><td><code>Authorization: Bearer ey...</code></td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example Output</h2>
        <CodeBlock language="bash" code={`$ exo scan
Scanning for hardcoded secrets...

  [!] config/db.go:14
      db_password = "supersecret123"

  [!] .env:3
      GITHUB_TOKEN=ghp_abc123XYZ

  [i] main_linux: skipped (binary)
  [i] assets/logo.png: skipped (binary)

Scan complete. 2 issue(s) found.`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Best Practices</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Run <code className="text-arch-pink">exo scan</code> before every commit or in your CI pipeline</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Add secrets to <code className="text-arch-pink">.env</code> (gitignored) and never hardcode them in source files</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Use environment variable references like <code className="text-arch-pink">os.Getenv("DB_PASSWORD")</code> instead</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Pair with <code className="text-arch-pink">exo gen pre-commit</code> to add scan as a pre-commit hook automatically</span></div>
        </div>
      </section>
    </article>
  );
}
