package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Saisathvik94/codemaxx/internal/keys"
	"github.com/Saisathvik94/codemaxx/internal/ui/colors"
)

type keyItem struct {
	name   string
	status string
}

func (i keyItem) Title() string       { return fmt.Sprintf("%-12s %s", i.name, i.status) }
func (i keyItem) Description() string { return "" }
func (i keyItem) FilterValue() string { return i.name }

type keyModel struct {
	list      list.Model
	input     textinput.Model
	mode      string // "list" or "input"
	selected  string
}


func SetNewKey(providers []string) keyModel {

	items := []list.Item{}
	for _, p := range providers {
		status := "✖ not set"
		if key, _ := keys.GetKey(p); key != "" {
			status = "✔ configured"
		}

		items = append(items, keyItem{
			name:   p,
			status: status,
		})
	}

	l := list.New(items, list.NewDefaultDelegate(), 45, 12)
	l.Title = "Manage Your API Keys"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	ti := textinput.New()
	ti.Placeholder = "Enter API key"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 40
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'

	return keyModel{
		list:  l,
		input: ti,
		mode:  "list",
	}
}

func (m keyModel) Init() tea.Cmd {
	return nil
}

func (m keyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.mode {

	case "list":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "enter":
				selected := m.list.SelectedItem().(keyItem)
				m.selected = selected.name
				m.mode = "input"
				m.input.SetValue("")
				m.input.Focus()
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	case "input":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "enter":
				err := keys.SetKey(m.selected, m.input.Value())
				if err != nil {
					m.input.SetValue("") 
					m.input.Placeholder = err.Error() // Show error as placeholder
					return m, nil
				}
				fmt.Println("✔ Key updated for", m.selected)
				return m, tea.Quit

			case "esc":
				m.mode = "list"
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m keyModel) View() string {

	if m.mode == "input" {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			colors.TitleStyle.Render("Set API Key for "+m.selected),
			"",
			m.input.View(),
			"",
			"Press Enter to save • Esc to cancel",
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		colors.TitleStyle.Render(m.list.Title),
		"",
		m.list.View(),
		"",
		"↑/↓ navigate • enter edit • q quit",
	)
}