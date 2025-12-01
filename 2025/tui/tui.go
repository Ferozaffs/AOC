package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFD700")).
			Margin(1, 0, 1, 2)

	normalItemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("#AAAAAA"))

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("#00D7FF")).
				Bold(true).
				Background(lipgloss.Color("#303030"))

	helpStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("#666666"))
)

type tuiModel struct {
	cursor   int
	commands []*cobra.Command
	selected *cobra.Command
}

func model(root *cobra.Command) tuiModel {
	var days []*cobra.Command
	for _, c := range root.Commands() {
		if c.Name() == "help" || c.Name() == "completion" || c.Name() == "dayX" {
			continue
		}

		days = append(days, c)
	}
	return tuiModel{commands: days}
}

func (m tuiModel) Init() tea.Cmd { return nil }

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.commands[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m tuiModel) View() string {
	s := titleStyle.Render("Select a day") + "\n\n"

	for i, c := range m.commands {
		line := fmt.Sprintf("%s — %s", c.Name(), c.Short)
		if i == m.cursor {
			s += selectedItemStyle.Render("> " + line)
		} else {
			s += normalItemStyle.Render("  " + line)
		}
		s += "\n"
	}

	s += "\n" + helpStyle.Render("↑/↓ or j/k to navigate • Enter to run • q to quit")
	return s
}

func RunTUI(root *cobra.Command) error {
	p := tea.NewProgram(model(root))
	finalModel, err := p.Run()
	if err != nil {
		return err
	}

	m, ok := finalModel.(tuiModel)
	if !ok {
		return fmt.Errorf("failed to cast final model")
	}

	if m.selected == nil {
		return nil // user quit
	}

	root.SetArgs([]string{m.selected.Name()})
	return root.Execute()
}
