import CodeBlock from '@/components/CodeBlock';

export default function TerraformOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Terraform</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          Cloud provider infrastructure modules for AWS, GCP, and Azure.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen infra --provider aws    # or gcp, azure`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generated Structure</h2>
        <CodeBlock language="text" code={`infra/
└── aws/            # (or gcp/ or azure/)
    ├── main.tf         # Core resources (VPC, EKS/GKE/AKS)
    ├── variables.tf    # Input variables
    └── provider.tf     # Provider configuration`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> AWS — main.tf</h2>
        <CodeBlock language="hcl" filename="infra/aws/main.tf" showLineNumbers code={`resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "\${var.project_name}-vpc"
  }
}

resource "aws_subnet" "public" {
  count                   = 2
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, count.index)
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = var.project_name
  cluster_version = "1.27"
  vpc_id          = aws_vpc.main.id
  subnet_ids      = aws_subnet.public[*].id
}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> GCP — main.tf</h2>
        <CodeBlock language="hcl" filename="infra/gcp/main.tf" showLineNumbers code={`resource "google_container_cluster" "primary" {
  name     = var.project_name
  location = var.region

  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "\${var.project_name}-node-pool"
  location   = var.region
  cluster    = google_container_cluster.primary.name
  node_count = var.node_count

  node_config {
    machine_type = "e2-medium"
  }
}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Azure — main.tf</h2>
        <CodeBlock language="hcl" filename="infra/azure/main.tf" showLineNumbers code={`resource "azurerm_resource_group" "main" {
  name     = var.project_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "main" {
  name                = var.project_name
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  dns_prefix          = var.project_name

  default_node_pool {
    name       = "default"
    node_count = var.node_count
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }
}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Variables</h2>
        <p className="text-sm text-arch-text-dim mb-3">All providers share similar variables:</p>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Variable</th><th>Default</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>project_name</code></td><td>from .exo.yaml</td><td>Used for resource naming</td></tr>
            <tr><td><code>region</code></td><td>provider-specific</td><td>Deployment region</td></tr>
            <tr><td><code>node_count</code></td><td><code>2</code></td><td>Number of K8s worker nodes</td></tr>
            <tr><td><code>vpc_cidr</code></td><td><code>10.0.0.0/16</code></td><td>VPC CIDR (AWS only)</td></tr>
          </tbody>
        </table>
      </section>
    </article>
  );
}
