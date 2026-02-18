<div align="center">

<pre>
 ███████╗██╗  ██╗ ██████╗
 ██╔════╝╚██╗██╔╝██╔═══██╗
 █████╗   ╚███╔╝ ██║   ██║
 ██╔══╝   ██╔██╗ ██║   ██║
 ███████╗██╔╝ ██╗╚██████╔╝
 ╚══════╝╚═╝  ╚═╝ ╚═════╝
</pre>

# EXO
### The Cloud-Native Bootstrap CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/Harsh-BH/Exo)](https://goreportcard.com/report/github.com/Harsh-BH/Exo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Harsh-BH/Exo/go.yml?branch=main)](https://github.com/Harsh-BH/Exo/actions)

<p align="center">
  <img src="https://via.placeholder.com/800x400.png?text=EXO+CLI+Demo+GIF" alt="EXO Demo" width="800">
</p>

Example: `exo init` → Detects Stack → Generates DevOps Assets → Ready to Deploy.

</div>

---

## Why EXO?

Starting a new microservice usually means copy-pasting Dockerfiles, wrestling with Kubernetes YAML, and debugging CI/CD pipelines. **EXO** automates this zero-to-one phase.

It analyzes your source code and generates **production-ready** infrastructure code in seconds.

## Features

| Feature | Description |
| :--- | :--- |
| **Smart Detection** | Automatically identifies Go, Node.js, and Python projects. |
| **Containerization** | Generates optimized, multi-stage `Dockerfiles`. |
| **Infrastructure** | Creates Terraform modules for AWS (VPC, EKS). |
| **CI/CD** | Sets up GitHub Actions workflows for build and test. |
| **Interactive UI** | Beautiful terminal UI for easy configuration. |
| **Cross-Platform** | Binaries for Linux, macOS, and Windows. |

## Installation

Download the latest release from the [Releases](https://github.com/Harsh-BH/Exo/releases) page or build from source:

### Build from Source
```bash
git clone https://github.com/Harsh-BH/Exo.git
cd Exo
./scripts/build.sh
```
The binary will be available in `bin/` (e.g., `bin/exo-linux-amd64`).

### Quick Install (Planned)
```bash
curl -sL https://get.exo.sh | bash
```

## Quick Start

1.  **Navigate to your project folder:**
    ```bash
    cd ~/my-awesome-app
    ```

2.  **Initialize EXO:**
    ```bash
    exo init
    ```
    *EXO will detect your language and ask you a few questions.*

3.  **Generate Specific Assets:**
    ```bash
    # Generate a Dockerfile
    exo gen docker

    # Generate CI/CD Pipeline
    exo gen ci

    # Generate AWS Infrastructure
    exo gen infra
    ```

## Commands

- `exo init`: Interactive setup wizard.
- `exo gen docker [--lang=go|node|python]`: Generate Dockerfile.
- `exo gen infra`: Generate Terraform for AWS.
- `exo gen ci`: Generate GitHub Actions workflow.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'feat: add some amazing feature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

---

<div align="center">
  <sub>Built with Love by Harsh-BH using Cobra & Bubble Tea.</sub>
</div>