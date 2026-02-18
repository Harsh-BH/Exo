<div align="center">

<pre>
 ███████╗██╗  ██╗ ██████╗
 ██╔════╝╚██╗██╔╝██╔═══██╗
 █████╗   ╚███╔╝ ██║   ██║
 ██╔══╝   ██╔██╗ ██║   ██║
 ███████╗██╔╝ ██╗╚██████╔╝
 ╚══════╝╚═╝  ╚═╝ ╚═════╝
</pre>

### **Ship Infrastructure, Not YAML**
**The Cloud-Native Bootstrap CLI — From Source Code to Production Infrastructure in Seconds**

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/Harsh-BH/Exo)](https://goreportcard.com/report/github.com/Harsh-BH/Exo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Harsh-BH/Exo/go.yml?branch=main)](https://github.com/Harsh-BH/Exo/actions)
[![Cobra](https://img.shields.io/badge/CLI-Cobra-blue)](https://cobra.dev/)
[![Bubble Tea](https://img.shields.io/badge/TUI-Bubble%20Tea-ff69b4)](https://github.com/charmbracelet/bubbletea)
[![Terraform](https://img.shields.io/badge/IaC-Terraform-623CE4?logo=terraform&logoColor=white)](https://www.terraform.io/)
[![Docker](https://img.shields.io/badge/Container-Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Orchestration-Kubernetes-326CE5?logo=kubernetes&logoColor=white)](https://kubernetes.io/)

</div>

---

## The Problem

Every new microservice demands the same boilerplate: Dockerfiles, Terraform modules, CI/CD pipelines, Kubernetes manifests, monitoring configs. It's repetitive, error-prone, and burns hours that should go toward building your product.

## The Solution

**EXO** is a single CLI that detects your project stack and generates production-ready DevOps scaffolding — Dockerfiles, Terraform, CI/CD, Kubernetes YAML, and monitoring — in seconds. You write the application code. EXO handles the infrastructure.

---

## Key Features

| Feature | What It Does | Why It Matters |
|---------|-------------|----------------|
| **Smart Stack Detection** | Auto-detects Go, Node.js, and Python projects from source files | Zero manual configuration needed |
| **Multi-Stage Dockerfiles** | Generates optimized, language-specific Dockerfiles | Smaller images, faster builds, production-ready defaults |
| **Multi-Cloud Terraform** | Scaffolds IaC modules for **AWS**, **GCP**, and **Azure** | One tool for any cloud — VPC, networking, and compute-ready |
| **CI/CD Pipelines** | Generates **GitHub Actions** and **GitLab CI** workflows | Push-to-deploy from day one |
| **Kubernetes Manifests** | Produces Deployment, Service, and Ingress YAML | Container orchestration without the YAML headaches |
| **Monitoring Stack** | Sets up **Prometheus** + **Grafana** with Docker Compose | Observability built in from the start |
| **Interactive Wizard** | Guided setup via a beautiful terminal UI (Bubble Tea) | Intuitive, developer-friendly experience |
| **Config Persistence** | Saves project config in `.exo.yaml` for repeatable runs | Idempotent, version-controllable infrastructure |
| **Cross-Platform** | Pre-built binaries for Linux, macOS (Intel & ARM), and Windows | Works everywhere your team does |

---

## Architecture

```
┌──────────────────────────────────────────────────────┐
│                    Source Code                        │
└──────────────────────┬───────────────────────────────┘
                       │
                       ▼
          ┌────────────────────────┐
          │  Stack Detection       │  ← go.mod · package.json · requirements.txt
          └────────────┬───────────┘
                       │
                       ▼
          ┌────────────────────────┐
          │  Interactive Wizard    │  ← Bubble Tea TUI
          │  (Language · Cloud ·   │
          │   CI · Monitoring)     │
          └────────────┬───────────┘
                       │
                       ▼
          ┌────────────────────────┐
          │  Template Engine       │  ← Go text/template
          └────────────┬───────────┘
                       │
        ┌──────────────┼──────────────────┐
        │              │                  │
        ▼              ▼                  ▼
  ┌───────────┐  ┌───────────┐  ┌──────────────┐
  │ Dockerfile │  │ Terraform │  │  CI/CD       │
  │ K8s YAML   │  │ (AWS/GCP/ │  │  Monitoring  │
  │            │  │  Azure)   │  │  Config      │
  └───────────┘  └───────────┘  └──────────────┘
        │              │                  │
        └──────────────┼──────────────────┘
                       │
                       ▼
          ┌────────────────────────┐
          │  Production-Ready Repo │  ← .exo.yaml persisted
          └────────────────────────┘
```

**Design Principles:**
- **Idempotent generation** — safe to re-run at any time
- **Minimal configuration** — smart defaults, override when needed
- **Cloud-native defaults** — best practices baked in
- **Extensible architecture** — modular templates, easy to add new stacks

---

## Tech Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| **Language** | Go 1.25 | Fast, compiled, cross-platform CLI |
| **CLI Framework** | [Cobra](https://cobra.dev/) | Industry-standard command routing and flag parsing |
| **Terminal UI** | [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Bubbles](https://github.com/charmbracelet/bubbles) | Interactive, state-driven wizard experience |
| **Styling** | [Lipgloss](https://github.com/charmbracelet/lipgloss) | Rich terminal colors and formatting |
| **Config** | YAML (gopkg.in/yaml.v3) | Human-readable project configuration |
| **Templates** | Go `text/template` | Fast, native template rendering engine |
| **Infra Targets** | Terraform, Docker, Kubernetes, Prometheus | Full-spectrum DevOps artifact generation |

---

## Installation

### Option 1 — Build from Source

```bash
git clone https://github.com/Harsh-BH/Exo.git
cd Exo
./scripts/build.sh
```

Binaries are output to the `bin/` directory:

```
bin/exo-linux-amd64
bin/exo-darwin-amd64       # macOS Intel
bin/exo-darwin-arm64       # macOS Apple Silicon
bin/exo-windows-amd64.exe
```

Move the appropriate binary to your `PATH`:

```bash
sudo mv bin/exo-linux-amd64 /usr/local/bin/exo
```

### Option 2 — Quick Install (Coming Soon)

```bash
curl -sL https://get.exo.sh | bash
```

> _A Homebrew formula and release binaries via GitHub Releases are planned._

---

## Quick Start

**1. Navigate to your project:**

```bash
cd ~/my-service
```

**2. Run the interactive wizard:**

```bash
exo init
```

The wizard walks you through selecting your language, cloud provider, CI/CD platform, and monitoring — then generates everything in one pass.

**3. Or generate individual components:**

```bash
exo gen docker                  # Dockerfile
exo gen infra --provider aws    # Terraform modules
exo gen k8s                     # Kubernetes manifests
exo gen ci                      # CI/CD pipeline
```

**4. Check what's been generated:**

```bash
exo status
```

---

## CLI Reference

| Command | Description | Flags |
|---------|-------------|-------|
| `exo init` | Launch interactive setup wizard | — |
| `exo gen docker` | Generate a multi-stage Dockerfile | `--name`, `--lang` (go/node/python) |
| `exo gen infra` | Generate Terraform modules | `--name`, `--provider` (aws/gcp/azure) |
| `exo gen k8s` | Generate Kubernetes manifests | `--name` |
| `exo gen ci` | Generate CI/CD pipeline | — |
| `exo status` | Show generated artifact status | — |
| `exo upgrade` | Re-run wizard with existing config pre-filled | — |
| `exo version` | Print version, build date, and Go/OS info | — |

---

## Configuration

EXO persists project settings in `.exo.yaml` at the project root. This file is created by `exo init` and read by `exo upgrade` and `exo status`.

```yaml
name: my-service        # Project name
language: go            # go | node | python
provider: aws           # aws | gcp | azure | none
ci: github-actions      # github-actions | gitlab-ci
monitoring: prometheus   # prometheus | none
```

This file should be committed to version control so your team shares the same infrastructure configuration.

---

## Generated Output Structure

After running `exo init`, your project will contain:

```
my-service/
├── Dockerfile                          # Optimized multi-stage build
├── .exo.yaml                           # EXO configuration
├── .github/workflows/
│   └── go.yml                          # GitHub Actions CI pipeline
├── .gitlab-ci.yml                      # GitLab CI pipeline (if selected)
├── infra/
│   └── aws/                            # (or gcp/ or azure/)
│       ├── main.tf                     # Core infrastructure resources
│       ├── variables.tf                # Input variables
│       └── provider.tf                 # Provider configuration
├── k8s/
│   ├── deployment.yaml                 # Kubernetes Deployment
│   ├── service.yaml                    # Kubernetes Service
│   └── ingress.yaml                    # Kubernetes Ingress
└── monitoring/
    ├── prometheus.yml                  # Prometheus scrape config
    └── docker-compose.monitoring.yml   # Prometheus + Grafana stack
```

---

## Project Structure (Source)

```
Exo/
├── main.go                     # CLI entry point
├── go.mod                      # Go module definition
├── cmd/exo/                    # Cobra command definitions
│   ├── root.go                 #   Root command and banner
│   ├── init.go                 #   Interactive wizard
│   ├── gen.go                  #   Generation dispatcher (docker/infra/k8s/ci)
│   ├── status.go               #   Artifact status checker
│   ├── upgrade.go              #   Config-aware re-run
│   └── version.go              #   Version info
├── internal/                   # Private application code
│   ├── config/                 #   .exo.yaml read/write
│   ├── detector/               #   Stack detection (Go/Node/Python)
│   ├── prompt/                 #   Bubble Tea interactive wizard
│   └── renderer/               #   Template rendering engine
├── templates/                  # Go text/template files
│   ├── docker/                 #   dockerfile.tmpl, node.tmpl, python.tmpl
│   ├── terraform/              #   aws/, gcp/, azure/ modules
│   ├── ci/                     #   github-actions.tmpl, gitlab-ci.tmpl
│   ├── k8s/                    #   deployment, service, ingress templates
│   └── monitoring/             #   prometheus.tmpl, docker-compose.monitoring.tmpl
├── infra/                      # Sample generated Terraform output
├── k8s/                        # Sample generated K8s manifests
├── scripts/
│   └── build.sh                # Cross-platform build script
└── bin/                        # Compiled binaries
```

---

## Deployment

### As a CLI Tool

EXO is a standalone binary — no runtime dependencies. Build once, distribute to your team:

```bash
# Build for all platforms
./scripts/build.sh

# Or build for your current platform only
go build -o exo main.go
```

### In CI/CD Pipelines

EXO can be integrated into your CI/CD to enforce consistent infrastructure scaffolding:

```yaml
# Example: GitHub Actions step
- name: Generate infrastructure
  run: |
    ./exo gen docker --name ${{ github.event.repository.name }} --lang go
    ./exo gen infra --provider aws
    ./exo gen k8s
```

---

## Screenshots

<!-- TODO: Add terminal screenshots or GIFs demonstrating:
  - `exo init` interactive wizard
  - `exo gen docker` output
  - `exo status` dashboard
  - Generated Dockerfile / Terraform / K8s YAML
-->

> _Screenshots and demo GIFs coming soon._

---

## Roadmap

- [x] Go, Node.js, Python stack detection
- [x] Multi-stage Dockerfile generation
- [x] Terraform scaffolding for AWS, GCP, Azure
- [x] GitHub Actions and GitLab CI pipeline generation
- [x] Kubernetes Deployment, Service, Ingress generation
- [x] Prometheus + Grafana monitoring setup
- [x] Interactive Bubble Tea wizard
- [x] `.exo.yaml` config persistence and upgrade flow
- [x] Cross-platform binary builds
- [ ] Helm chart generation
- [ ] Custom template plugin system
- [ ] `exo destroy` — clean up generated artifacts
- [ ] Embedded templates via `go:embed`
- [ ] Homebrew and Scoop package distribution
- [ ] Release binaries via GitHub Releases
- [ ] AI-assisted infrastructure recommendations
- [ ] Multi-service (monorepo) support
- [ ] Secrets management integration (Vault, AWS Secrets Manager)

---

## Contributing

Contributions are welcome! EXO follows a modular architecture that makes it easy to add new templates, stacks, and cloud providers.

### Getting Started

```bash
# Fork and clone the repository
git clone https://github.com/<your-username>/Exo.git
cd Exo

# Run tests
go test ./...

# Format code
gofmt -s -w .

# Build
go build -o exo main.go
```

### Guidelines

- **Commits**: Use [Conventional Commits](https://www.conventionalcommits.org/) — `feat:`, `fix:`, `docs:`, `refactor:`
- **Branches**: Create feature branches from `main`
- **Tests**: Table-driven tests preferred; run `go test ./...` before submitting
- **Linting**: Run `golangci-lint run` to check code quality
- **Templates**: New templates go in `templates/<category>/` with `.tmpl` extension

### Adding a New Stack or Cloud Provider

1. Add a detection rule in `internal/detector/detect.go`
2. Create template files in `templates/`
3. Wire up the generation logic in `cmd/exo/gen.go`
4. Update the interactive wizard in `internal/prompt/prompt.go`
5. Add tests and update documentation

---

## License

MIT License — see [`LICENSE`](LICENSE) for details.

Copyright © 2026 [Harsh Bhatt](https://github.com/Harsh-BH)

---

## Suggested Enhancements

> _Areas identified for future improvement based on codebase analysis:_

| Area | Suggestion |
|------|-----------|
| **Template Embedding** | Use `go:embed` to bundle templates into the binary, removing the need to resolve paths at runtime |
| **Validation** | Add post-generation validation (e.g., `terraform validate`, `docker build --check`) |
| **Dry Run Mode** | Add `--dry-run` flag to preview generated output without writing files |
| **Verbose/Quiet Modes** | Add `--verbose` and `--quiet` flags for CI-friendly and debug output |
| **Test Coverage** | Expand tests for `init.go`, `status.go`, and `upgrade.go` commands |
| **Telemetry** | Optional anonymous usage analytics to guide feature prioritization |
| **Config Versioning** | Version the `.exo.yaml` schema to support backward-compatible upgrades |
| **Interactive Docs** | Add `exo docs` command to open documentation in the browser |

---

<div align="center">

**Built by [Harsh Bhatt](https://github.com/Harsh-BH)**

Powered by **Go** · **Cobra** · **Bubble Tea** · **Terraform** · **Docker** · **Kubernetes**

[⬆ Back to Top](#)

</div>
