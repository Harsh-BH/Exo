<div align="center">

<pre>
 ███████╗██╗  ██╗ ██████╗
 ██╔════╝╚██╗██╔╝██╔═══██╗
 █████╗   ╚███╔╝ ██║   ██║
 ██╔══╝   ██╔██╗ ██║   ██║
 ███████╗██╔╝ ██╗╚██████╔╝
 ╚══════╝╚═╝  ╚═╝ ╚═════╝
</pre>

### Cloud-Native Bootstrap CLI  
**From Source Code → Production Infrastructure in Seconds**

[![Go Report Card](https://goreportcard.com/badge/github.com/Harsh-BH/Exo)](https://goreportcard.com/report/github.com/Harsh-BH/Exo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Harsh-BH/Exo/go.yml?branch=main)](https://github.com/Harsh-BH/Exo/actions)

</div>

---

## $ whoami

**EXO** is a cloud-native bootstrap engine for modern microservices.

It eliminates the repetitive DevOps overhead of:
- Writing Dockerfiles  
- Configuring CI/CD pipelines  
- Crafting Terraform modules  
- Setting up Kubernetes infrastructure  

You write application code.  
EXO generates the production scaffolding.

---

## $ why exo

Bootstrapping a service typically requires:

```
Dockerfile
Kubernetes YAML
Terraform modules
CI workflows
Environment configs
Cloud wiring
```

Most of it is repetitive.  
Most of it is error-prone.  

**EXO compresses this zero-to-deployment lifecycle into a single CLI workflow.**

---

## $ capabilities

| Module | What EXO Does |
|--------|---------------|
| Stack Detection | Automatically detects Go, Node.js, and Python projects |
| Container Engine | Generates optimized multi-stage Dockerfiles |
| Infrastructure | Provisions AWS-ready Terraform (VPC, EKS-ready structure) |
| CI/CD | Generates GitHub Actions pipelines |
| Interactive CLI | Guided configuration via terminal UI |
| Cross Platform | Linux, macOS, Windows builds |

---

## $ architecture

```
┌──────────────────┐
│   Source Code    │
└─────────┬────────┘
          │
          ▼
   Stack Detection Engine
          │
          ▼
  Asset Generation Layer
  ├── Dockerfile
  ├── Terraform
  ├── CI/CD
  └── Config Templates
          │
          ▼
   Production-Ready Repo
```

Design principles:
- Idempotent generation
- Minimal configuration
- Cloud-native defaults
- Extensible architecture

---

## $ install

### Option 1 — Build from Source

```bash
git clone https://github.com/Harsh-BH/Exo.git
cd Exo
./scripts/build.sh
```

Binary will be available in:

```
bin/exo-<os>-<arch>
```

---

### Option 2 — Quick Install (Planned)

```bash
curl -sL https://get.exo.sh | bash
```

---

## $ quickstart

Navigate to your service:

```bash
cd ~/my-service
```

Initialize EXO:

```bash
exo init
```

Generate components individually:

```bash
exo gen docker
exo gen ci
exo gen infra
```

Workflow example:

```
$ exo init
> Detecting stack...
> Found: Go
> Generating Dockerfile...
> Generating CI workflow...
> Generating Terraform modules...
> Setup complete.
```

---

## $ commands

| Command | Description |
|----------|-------------|
| `exo init` | Interactive setup wizard |
| `exo gen docker` | Generate Dockerfile |
| `exo gen infra` | Generate Terraform modules |
| `exo gen ci` | Generate GitHub Actions workflow |

---

## $ project structure (generated)

```
.
├── Dockerfile
├── .github/workflows/
│   └── ci.yml
├── infra/
│   ├── main.tf
│   ├── variables.tf
│   └── outputs.tf
└── exo.config.yaml
```

---

## $ roadmap

- GCP and Azure support  
- Helm chart generation  
- Kubernetes YAML generation  
- Multi-cloud templates  
- Plugin system for custom stacks  
- AI-assisted infrastructure recommendations  

---

## $ contributing

```
git fork
git checkout -b feature/your-feature
git commit -m "feat: add feature"
git push
```

Open a Pull Request.

EXO follows conventional commits and modular CLI architecture.

---

## $ license

MIT License — see `LICENSE` for details.

---

<div align="center">

Built by Harsh-BH  
Powered by Go, Cobra, and Bubble Tea  

</div>
