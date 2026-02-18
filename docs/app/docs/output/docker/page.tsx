import CodeBlock from '@/components/CodeBlock';

export default function DockerOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Docker</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Language-specific, production-ready Dockerfiles with multi-stage builds.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen docker`} />
        <p className="text-sm text-arch-text-dim mt-3">
          Generates a <code className="text-arch-pink">Dockerfile</code> in the project root. The template
          is selected based on the <code className="text-arch-pink">language</code> field in{' '}
          <code className="text-arch-pink">.exo.yaml</code>.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Go Dockerfile</h2>
        <CodeBlock language="dockerfile" filename="Dockerfile" showLineNumbers code={`# ── Build Stage ──
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-name

# ── Runtime Stage ──
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app-name .
EXPOSE 8080
CMD ["./app-name"]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Node.js Dockerfile</h2>
        <CodeBlock language="dockerfile" filename="Dockerfile" showLineNumbers code={`FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app/node_modules ./node_modules
COPY . .
EXPOSE 3000
CMD ["node", "index.js"]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Python Dockerfile</h2>
        <CodeBlock language="dockerfile" filename="Dockerfile" showLineNumbers code={`FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["python", "app.py"]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Features</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Multi-stage builds for smaller final images</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Alpine base images for minimal attack surface</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Dependency caching via layer ordering</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> CGO disabled for Go (static binary)</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Production-only npm install for Node</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> No-cache pip install for Python</div>
        </div>
      </section>
    </article>
  );
}
