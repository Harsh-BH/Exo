import CodeBlock from '@/components/CodeBlock';

export default function HelmOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Helm Charts</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Complete Helm chart with Chart.yaml, values.yaml, and templated K8s manifests.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen helm --name my-api`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generated Structure</h2>
        <CodeBlock language="text" code={`charts/
└── my-api/
    ├── Chart.yaml
    ├── values.yaml
    └── templates/
        ├── deployment.yaml
        ├── service.yaml
        └── ingress.yaml`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Chart.yaml</h2>
        <CodeBlock language="yaml" filename="charts/my-api/Chart.yaml" showLineNumbers code={`apiVersion: v2
name: my-api
description: A Helm chart for my-api
type: application
version: 0.1.0
appVersion: "1.0.0"`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> values.yaml</h2>
        <CodeBlock language="yaml" filename="charts/my-api/values.yaml" showLineNumbers code={`replicaCount: 2

image:
  repository: my-api
  pullPolicy: IfNotPresent
  tag: "latest"

service:
  type: ClusterIP
  port: 80

ingress:
  enabled: true
  className: nginx
  hosts:
    - host: my-api.local
      paths:
        - path: /
          pathType: Prefix

resources:
  limits:
    cpu: 500m
    memory: 128Mi
  requests:
    cpu: 250m
    memory: 64Mi`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Usage</h2>
        <CodeBlock language="bash" code={`# Install the chart
helm install my-api ./charts/my-api

# Upgrade with custom values
helm upgrade my-api ./charts/my-api --set replicaCount=3

# Template (dry-run)
helm template my-api ./charts/my-api`} />
      </section>
    </article>
  );
}
