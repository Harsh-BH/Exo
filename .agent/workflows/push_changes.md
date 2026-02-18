---
description: Safely commit and push changes with tests and conventional commit messages.
---
# Push Changes Workflow

This workflow ensures code quality before pushing.

1. Check current status:
    ```bash
    git status
    ```

2. Run tests to ensure safety:
    ```bash
    go test ./...
    ```
    *If tests fail, STOP and fix them.*

3. Format code:
    ```bash
    gofmt -s -w .
    ```

4. Stage all changes:
    ```bash
    git add .
    ```

5. Commit changes:
    - Determine the type (feat, fix, docs, style, refactor, test, chore).
    - Write a short, imperative summary.
    - Example: `git commit -m "feat: add project detector logic"`
    ```bash
    # Replace with your actual message
    git commit -m "your_message_here"
    ```

6. Push to remote:
    ```bash
    git push origin main
    ```
