package main
import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct{
	width int
	height int
	logs []string
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

func initialModel() model {
	var logs []string

	for i := 0; i < 100; i++ {
		line := fmt.Sprintf("line %d: user login ok", i)
		logs = append(logs, line)
	}

	return model{logs: logs}
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	headerStyle := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#FFFFFF")).Width(m.width)
	header := headerStyle.Render("cwtail")

	footerStyle := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#767676")).Width(m.width)
	footer := footerStyle.Render("q quit")

	bodyHeight := m.height - lipgloss.Height(header) - lipgloss.Height(footer)
	bodyStyle := lipgloss.NewStyle().Height(bodyHeight)

	end := min(bodyHeight, len(m.logs))
	visible := m.logs[0:end]

	body := bodyStyle.Render(strings.Join(visible, "\n"))

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := p.Run()

	if err != nil {
		fmt.Println(err)
	}
}



