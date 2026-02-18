import CodeBlock from '@/components/CodeBlock';

export default function CLIDocsCommandPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo docs</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Open the Exo documentation in your default browser.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo docs`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text">
          Opens the Exo documentation website in your default browser. Supports macOS{' '}
          (<code className="text-arch-pink">open</code>), Linux (<code className="text-arch-pink">xdg-open</code>),
          and Windows (<code className="text-arch-pink">start</code>).
        </p>
      </section>
    </article>
  );
}
