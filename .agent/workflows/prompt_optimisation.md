---
description: How to improve and optimise the interactive Bubble Tea prompt in EXO
---

# Prompt Optimisation Workflow

This workflow guides you through improving the interactive CLI prompt experience in `internal/prompt/prompt.go`.

## When to use
- The prompt feels slow or unresponsive
- You want to add new input fields (e.g., cloud provider, language selection)
- You want to improve the visual styling with lipgloss
- You want to add validation to user inputs

---

## Steps

1. **Review the current prompt model**
    ```bash
    cat internal/prompt/prompt.go
    ```

2. **Identify what to improve** — choose one or more:
    - Add a new field (e.g., project language, cloud provider)
    - Add input validation (e.g., reject empty names)
    - Add lipgloss styling (colors, borders)
    - Add multi-step prompts (wizard-style)

3. **Add a new input field** (example: adding a language selector)
    - Add a new `textinput.Model` or `list.Model` to the `model` struct
    - Update `Init()` to blink the new input
    - Update `Update()` to handle tab/enter navigation between fields
    - Update `View()` to render the new field
    - Update `ProjectData` struct to include the new field
    - Update `Run()` to return the new field value

4. **Add input validation**
    - In `Update()`, on `tea.KeyEnter`, check if the value is empty
    - If invalid, set an `errMsg string` on the model and display it in `View()`
    - Only call `tea.Quit` when input is valid

5. **Add lipgloss styling**
    ```go
    import "github.com/charmbracelet/lipgloss"

    var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
    var errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
    ```
    - Wrap labels and error messages with these styles in `View()`

6. **Run and test the prompt interactively**
    ```bash
    go run main.go init
    ```
    - Verify the prompt renders correctly
    - Test edge cases: empty input, very long names, special characters

7. **Run unit tests**
    ```bash
    go test ./internal/prompt/...
    ```

8. **Push changes**
    - Follow the `/push_changes` workflow
