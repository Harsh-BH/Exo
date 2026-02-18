import CodeBlock from '@/components/CodeBlock';

export default function MonitoringOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Monitoring</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Prometheus + Grafana monitoring stack with pre-built dashboards and alert rules.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen monitoring`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generated Files</h2>
        <CodeBlock language="text" code={`monitoring/
├── prometheus.yml                    # Scrape configuration
└── docker-compose.monitoring.yml     # Prometheus + Grafana stack
grafana_dashboard.json                # Pre-built dashboard
alerts.yml                            # Prometheus alert rules`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Prometheus Config</h2>
        <CodeBlock language="yaml" filename="monitoring/prometheus.yml" showLineNumbers code={`global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'my-service'
    static_configs:
      - targets: ['host.docker.internal:8080']`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Docker Compose</h2>
        <CodeBlock language="yaml" filename="monitoring/docker-compose.monitoring.yml" showLineNumbers code={`version: '3.8'
services:
  prometheus:
    image: prom/prometheus:latest
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Alert Rules</h2>
        <CodeBlock language="yaml" filename="alerts.yml" showLineNumbers code={`groups:
  - name: my-service-alerts
    rules:
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.1
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High error rate detected"

      - alert: HighLatency
        expr: histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) > 0.5
        for: 5m
        labels:
          severity: warning`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Usage</h2>
        <CodeBlock language="bash" code={`# Start monitoring stack
docker-compose -f monitoring/docker-compose.monitoring.yml up -d

# Access
# Prometheus: http://localhost:9090
# Grafana:    http://localhost:3000 (admin/admin)`} />
      </section>
    </article>
  );
}
