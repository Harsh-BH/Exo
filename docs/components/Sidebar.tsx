'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navSections = [
  {
    title: 'GETTING STARTED',
    items: [
      { label: 'Introduction', href: '/docs' },
      { label: 'Installation', href: '/docs/installation' },
      { label: 'Quick Start', href: '/docs/quickstart' },
    ],
  },
  {
    title: 'CORE CONCEPTS',
    items: [
      { label: 'Architecture', href: '/docs/architecture' },
      { label: 'Configuration', href: '/docs/configuration' },
      { label: 'Stack Detection', href: '/docs/detection' },
      { label: 'Template Engine', href: '/docs/templates' },
    ],
  },
  {
    title: 'CLI REFERENCE',
    items: [
      { label: 'exo init', href: '/docs/cli/init' },
      { label: 'exo gen', href: '/docs/cli/gen' },
      { label: 'exo add', href: '/docs/cli/add' },
      { label: 'exo status', href: '/docs/cli/status' },
      { label: 'exo validate', href: '/docs/cli/validate' },
      { label: 'exo scan', href: '/docs/cli/scan' },
      { label: 'exo lint', href: '/docs/cli/lint' },
      { label: 'exo plugin', href: '/docs/cli/plugin' },
      { label: 'exo template', href: '/docs/cli/template' },
      { label: 'exo history', href: '/docs/cli/history' },
      { label: 'exo docs', href: '/docs/cli/docs' },
      { label: 'exo update', href: '/docs/cli/update' },
      { label: 'exo completion', href: '/docs/cli/completion' },
    ],
  },
  {
    title: 'GENERATED OUTPUT',
    items: [
      { label: 'Dockerfiles', href: '/docs/output/docker' },
      { label: 'Terraform', href: '/docs/output/terraform' },
      { label: 'Kubernetes', href: '/docs/output/kubernetes' },
      { label: 'Helm Charts', href: '/docs/output/helm' },
      { label: 'CI/CD Pipelines', href: '/docs/output/cicd' },
      { label: 'Monitoring', href: '/docs/output/monitoring' },
      { label: 'Databases', href: '/docs/output/databases' },
    ],
  },
  {
    title: 'ADVANCED',
    items: [
      { label: 'Plugins', href: '/docs/advanced/plugins' },
      { label: 'Custom Templates', href: '/docs/advanced/custom-templates' },
      { label: 'CI/CD Integration', href: '/docs/advanced/ci-integration' },
      { label: 'Contributing', href: '/docs/contributing' },
    ],
  },
];

export default function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="w-64 shrink-0 border-r border-arch-border bg-arch-bg overflow-y-auto h-[calc(100vh-56px)] sticky top-14 hidden lg:block">
      <nav className="py-6 px-2">
        {navSections.map((section) => (
          <div key={section.title} className="mb-6">
            <h3 className="px-3 mb-2 text-[10px] font-bold tracking-widest text-arch-text-dim uppercase">
              {section.title}
            </h3>
            {section.items.map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className={`sidebar-link ${pathname === item.href ? 'active' : ''}`}
              >
                <span className="text-arch-text-dim mr-2">›</span>
                {item.label}
              </Link>
            ))}
          </div>
        ))}
      </nav>
    </aside>
  );
}
