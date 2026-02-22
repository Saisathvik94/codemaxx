package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Saisathvik94/codemaxx/internal/config"
	"github.com/Saisathvik94/codemaxx/internal/ui/colors"
)

type item struct {
	name    string
	current bool
}

func (i item) Title() string {
	if i.current {
		return "✔ " + i.name + " (current)"
	}
	return "  " + i.name
}

func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.name }

type model struct {
	list     list.Model
	quitting bool
}


func NewModelSelector(providers []string) model {
	cfg, _ := config.Load()

	items := []list.Item{}
	for _, p := range providers {
		items = append(items, item{
			name:    p,
			current: p == cfg.DefaultProvider,
		})
	}

	l := list.New(items, list.NewDefaultDelegate(), 40, 12)
	l.Title = "Select Default AI Provider"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return model{list: l}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			selected := m.list.SelectedItem().(item)

			err := config.SetDefaultProvider(selected.name)
			if err != nil {
				fmt.Println(colors.SuccessStyle.Render(
					fmt.Sprintln("Failed to update config: %w", err),
				))
				return m, tea.Quit
			}

			fmt.Println(colors.SuccessStyle.Render(
				fmt.Sprintf("✔ Default provider updated to: %s", selected.name),
			))

			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	header := colors.TitleStyle.Render(m.list.Title)
	footer := colors.FooterStyle.Render("\n↑/↓ navigate • enter select • q quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		m.list.View(),
		footer,
	)
}