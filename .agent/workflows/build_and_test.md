---
description: Build and test the EXO CLI project.
---
# Build and Test Workflow

This workflow builds the project binary and runs all tests.

1. Format the code:
    ```bash
    gofmt -s -w .
    ```

2. Run linter (if installed):
    ```bash
    # golangci-lint run
    echo "Skipping lint for now (install golangci-lint if needed)"
    ```

3. Run unit tests:
    ```bash
    go test -v ./...
    ```

4. Build the binary:
    ```bash
    go build -o exo main.go
    ```

5. Verify binary execution:
    ```bash
    ./exo --help
    ```
