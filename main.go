package main
import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct{
	width int
	height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#FFFFFF")).Width(m.width)
	header := headerStyle.Render("cwtail")

	footerStyle := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#767676")).Width(m.width)
	footer := footerStyle.Render("q quit")

	bodyHeight := m.height - lipgloss.Height(header) - lipgloss.Height(footer)
	bodyStyle := lipgloss.NewStyle().Height(bodyHeight)
	body := bodyStyle.Render("body")

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	_, err := p.Run()

	if err != nil {
		fmt.Println(err)
	}
}



