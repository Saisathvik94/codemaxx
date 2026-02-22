package colors

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor = lipgloss.Color("#7C3AED")
	SuccessColor = lipgloss.Color("#22C55E")
	ErrorColor   = lipgloss.Color("#EF4444")
	MutedColor   = lipgloss.Color("#6B7280")

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor)

	FooterStyle = lipgloss.NewStyle().
			Foreground(MutedColor)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)
)
