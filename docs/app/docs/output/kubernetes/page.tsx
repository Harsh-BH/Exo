import CodeBlock from '@/components/CodeBlock';

export default function KubernetesOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Kubernetes</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Production-ready Kubernetes manifests — Deployment, Service, and Ingress.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen k8s`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Deployment</h2>
        <CodeBlock language="yaml" filename="k8s/deployment.yaml" showLineNumbers code={`apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-service
  labels:
    app: my-service
spec:
  replicas: 2
  selector:
    matchLabels:
      app: my-service
  template:
    metadata:
      labels:
        app: my-service
    spec:
      containers:
        - name: my-service
          image: my-service:latest
          ports:
            - containerPort: 8080
          resources:
            requests:
              memory: "64Mi"
              cpu: "250m"
            limits:
              memory: "128Mi"
              cpu: "500m"`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Service</h2>
        <CodeBlock language="yaml" filename="k8s/service.yaml" showLineNumbers code={`apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: my-service
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: ClusterIP`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Ingress</h2>
        <CodeBlock language="yaml" filename="k8s/ingress.yaml" showLineNumbers code={`apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-service
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
    - host: my-service.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: my-service
                port:
                  number: 80`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Features</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Resource requests and limits pre-configured</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> 2 replicas for basic high availability</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> ClusterIP service type (internal by default)</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Nginx Ingress annotations included</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Project name injected from .exo.yaml</div>
        </div>
      </section>
    </article>
  );
}
