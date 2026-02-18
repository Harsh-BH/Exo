import CodeBlock from '@/components/CodeBlock';

export default function CLIHistoryPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo history</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Show the generation history log — what was generated and when.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo history`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Displays a timestamped log of all generation events. Reads from{' '}
          <code className="text-arch-pink">.exo/history.log</code> in the project root. Useful for
          auditing what was generated and when.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example</h2>
        <CodeBlock language="bash" code={`$ exo history
Generation History

  2024-01-15 10:30:00  exo init (full project)
  2024-01-15 10:35:00  exo gen docker
  2024-01-16 14:20:00  exo add db --db redis
  2024-01-17 09:00:00  exo gen helm --name api`} />
      </section>
    </article>
  );
}
