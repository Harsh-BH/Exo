---
description: Instructions for adding a new template (e.g., for a new language or tool).
---
# Add New Template Workflow

This workflow guides you through adding a new template (e.g., Rust support, Azure Terraform).

1. Create the template file:
    - Create a new file in `templates/<category>/<name>.tmpl`.
    - Example: `templates/docker/rust.tmpl`.
    - Use Go template syntax (e.g., `{{ .AppName }}`).

2. Update the Detector (if applicable):
    - Identify the file(s) that indicate this language/tool (e.g., `Cargo.toml`).
    - Update `internal/detector/detect.go` to recognize this file.
    - Add a new constant or enum case for the new type.

3. Update the Renderer:
    - Update `internal/renderer/render.go` to include the new template path.
    - Ensure the data struct passed to the template has the necessary fields.

4. Add a Test Case:
    - Add a test case in `internal/detector/detect_test.go` with a sample file structure.
    - Add a test case in `internal/renderer/render_test.go` to verify output generation.

5. Verify:
    - Run the `build_and_test` workflow.
