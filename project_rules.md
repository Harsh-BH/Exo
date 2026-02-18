# Project Rules and Guidelines for EXO

## 1. Coding Standards (Go)
- **Formatting:** Always run `gofmt -s -w .` before committing.
- **Linting:** Use `golangci-lint` to ensure code quality.
- **Error Handling:** 
    - Don't ignore errors. Handle them or propagate them.
    - Use `fmt.Errorf("context: %w", err)` to wrap errors.
- **Comments:** Exported functions and types must have comments.
- **Complexity:** Keep functions small and focused.

## 2. Project Structure
- `cmd/`: Application entry points (Cobra commands).
- `internal/`: Private application and library code.
    - `detector/`: Logic to detect project types.
    - `prompt/`: UI logic (Bubble Tea/Survey).
    - `renderer/`: Template rendering logic.
- `templates/`: Raw template files (`.tmpl`).
- `pkg/`: Library code that's ok to use by external applications (if any).

## 3. Workflow & Contribution
- **Commits:** Use [Conventional Commits](https://www.conventionalcommits.org/) (e.g., `feat: add node detection`, `fix: template rendering`).
- **Branches:** Feature branches off `main`.
- **PRs:** Require passing tests and linting.

## 4. Architecture Decision Records (ADR)
- **CLI Framework:** Cobra (Standard, robust).
- **UI:** Bubble Tea (Modern, interactive) or Survey (Simple prompts). *Decision: Start with Survey for simplicity, upgrade to Bubble Tea for complex dashboards.*
- **Templates:** Go `text/template` (Native, fast).

## 5. Security
- **Templates:** Ensure templates don't hardcode secrets. Use placeholders.
- **Dependencies:** Regularly scan `go.mod` for vulnerabilities.

## 6. Testing
- **Unit Tests:** `go test ./...` should pass.
- **Table-Driven Tests:** Preferred for logic like detectors and renderers.
