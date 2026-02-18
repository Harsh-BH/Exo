package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ─── Styles ───────────────────────────────────────────────────────────────────

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	stepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)

	highlightStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))

	summaryKeyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	progressFilledStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	progressEmptyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("237"))

	logoStyle = lipgloss.NewStyle().
			Bold(true)
)

const AsciiLogo = `
  ███████╗██╗  ██╗ ██████╗ 
  ██╔════╝╚██╗██╔╝██╔═══██╗
  █████╗   ╚███╔╝ ██║   ██║
  ██╔══╝   ██╔██╗ ██║   ██║
  ███████╗██╔╝ ██╗╚██████╔╝
  ╚══════╝╚═╝  ╚═╝ ╚═════╝ 
`

// GetLogo returns the styled ASCII logo with a colorful gradient.
func GetLogo() string {
	lines := strings.Split(strings.TrimSpace(AsciiLogo), "\n")
	var styled strings.Builder
	// Vibrant gradient colors (cyan to magenta)
	colors := []string{"#00f5d4", "#00bbf9", "#9b5de5", "#f15bb5", "#fee440", "#ff9f1c"}
	for i, line := range lines {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[i%len(colors)])).Bold(true)
		styled.WriteString(style.Render(line) + "\n")
	}
	return styled.String()
}

// ─── Data ─────────────────────────────────────────────────────────────────────

// ProjectData holds all configuration collected from the wizard.
type ProjectData struct {
	Name       string
	Language   string
	Provider   string
	CI         string
	Monitoring string
}

// ─── List Item ────────────────────────────────────────────────────────────────

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// ─── Steps ────────────────────────────────────────────────────────────────────

const (
	stepName = iota
	stepLanguage
	stepProvider
	stepCI
	stepMonitoring
	stepConfirm
	stepDone
)

const totalSteps = stepDone // 6

// ─── Model ────────────────────────────────────────────────────────────────────

type model struct {
	step      int
	nameInput textinput.Model
	lists     [4]list.Model
	data      ProjectData
	errMsg    string
	confirmed bool
}

func newList(title string, items []list.Item) list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("86")).
		BorderLeftForeground(lipgloss.Color("86"))
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("243")).
		BorderLeftForeground(lipgloss.Color("86"))

	l := list.New(items, delegate, 44, 10)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	return l
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-service"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30

	langList := newList("Select Language", []list.Item{
		item{"go", "Go (Golang)"},
		item{"node", "Node.js"},
		item{"python", "Python"},
	})
	providerList := newList("Select Cloud Provider", []list.Item{
		item{"aws", "Amazon Web Services (EKS)"},
		item{"gcp", "Google Cloud Platform (GKE)"},
		item{"azure", "Microsoft Azure (AKS)"},
		item{"none", "Skip infrastructure generation"},
	})
	ciList := newList("Select CI/CD Tool", []list.Item{
		item{"github-actions", "GitHub Actions"},
		item{"gitlab-ci", "GitLab CI"},
		item{"none", "Skip CI/CD generation"},
	})
	monitoringList := newList("Select Monitoring Stack", []list.Item{
		item{"prometheus", "Prometheus + Grafana"},
		item{"none", "Skip monitoring setup"},
	})

	return model{
		step:      stepName,
		nameInput: ti,
		lists:     [4]list.Model{langList, providerList, ciList, monitoringList},
	}
}

// ─── Progress Bar ─────────────────────────────────────────────────────────────

func progressBar(current, total int) string {
	filled := current
	empty := total - current
	bar := progressFilledStyle.Render(strings.Repeat("█", filled))
	bar += progressEmptyStyle.Render(strings.Repeat("░", empty))
	return fmt.Sprintf("[%s] %d/%d", bar, current, total)
}

// ─── Init ─────────────────────────────────────────────────────────────────────

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// ─── Update ───────────────────────────────────────────────────────────────────

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			return m.advance()
		case tea.KeyEsc:
			if m.step > stepName {
				m.step--
				m.errMsg = ""
				return m, nil
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	if m.step == stepName {
		m.nameInput, cmd = m.nameInput.Update(msg)
	} else if m.step >= stepLanguage && m.step <= stepMonitoring {
		listIdx := m.step - stepLanguage
		m.lists[listIdx], cmd = m.lists[listIdx].Update(msg)
	}
	return m, cmd
}

func (m model) advance() (tea.Model, tea.Cmd) {
	m.errMsg = ""
	switch m.step {
	case stepName:
		name := strings.TrimSpace(m.nameInput.Value())
		if name == "" {
			m.errMsg = "Project name cannot be empty."
			return m, nil
		}
		m.data.Name = name
	case stepLanguage:
		if sel, ok := m.lists[0].SelectedItem().(item); ok {
			m.data.Language = sel.title
		}
	case stepProvider:
		if sel, ok := m.lists[1].SelectedItem().(item); ok {
			m.data.Provider = sel.title
		}
	case stepCI:
		if sel, ok := m.lists[2].SelectedItem().(item); ok {
			m.data.CI = sel.title
		}
	case stepMonitoring:
		if sel, ok := m.lists[3].SelectedItem().(item); ok {
			m.data.Monitoring = sel.title
		}
	case stepConfirm:
		m.confirmed = true
		m.step = stepDone
		return m, tea.Quit
	}
	m.step++
	return m, nil
}

// ─── View ─────────────────────────────────────────────────────────────────────

func (m model) View() string {
	var b strings.Builder

	// Logo (#12)
	b.WriteString(GetLogo() + "\n")

	// Header with progress bar
	b.WriteString(titleStyle.Render("EXO Setup Wizard") + "\n")
	b.WriteString(progressBar(m.step, totalSteps) + "\n\n")

	switch m.step {
	case stepName:
		b.WriteString("What is the name of your project?\n\n")
		b.WriteString(m.nameInput.View() + "\n")
	case stepLanguage, stepProvider, stepCI, stepMonitoring:
		listIdx := m.step - stepLanguage
		b.WriteString(m.lists[listIdx].View() + "\n")
	case stepConfirm:
		b.WriteString(titleStyle.Render("Summary — Review your choices") + "\n\n")
		b.WriteString(summaryKeyStyle.Render("  Project:    ") + m.data.Name + "\n")
		b.WriteString(summaryKeyStyle.Render("  Language:   ") + m.data.Language + "\n")
		b.WriteString(summaryKeyStyle.Render("  Provider:   ") + m.data.Provider + "\n")
		b.WriteString(summaryKeyStyle.Render("  CI/CD:      ") + m.data.CI + "\n")
		b.WriteString(summaryKeyStyle.Render("  Monitoring: ") + m.data.Monitoring + "\n\n")
		b.WriteString(highlightStyle.Render("Press Enter to generate all assets.") + "\n")
	}

	if m.errMsg != "" {
		b.WriteString("\n" + errorStyle.Render("✗ "+m.errMsg) + "\n")
	}

	// Navigation hints (#3)
	if m.step == stepName {
		b.WriteString(stepStyle.Render("\n↵ Enter to continue  •  Ctrl+C to quit") + "\n")
	} else if m.step != stepConfirm {
		b.WriteString(stepStyle.Render("\n↵ Enter to continue  •  Esc to go back  •  Ctrl+C to quit") + "\n")
	} else {
		b.WriteString(stepStyle.Render("\n↵ Enter to confirm  •  Esc to go back  •  Ctrl+C to quit") + "\n")
	}

	return b.String()
}

// ─── Public API ───────────────────────────────────────────────────────────────

// Run starts the interactive wizard and returns the collected project data.
func Run() (*ProjectData, error) {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return nil, err
	}
	if finalModel, ok := m.(model); ok {
		if !finalModel.confirmed {
			return nil, fmt.Errorf("setup cancelled")
		}
		return &finalModel.data, nil
	}
	return nil, fmt.Errorf("could not retrieve model")
}
