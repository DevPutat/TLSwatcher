package tui

import (
	"fmt"
	"time"

	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//TODO:: UPDATE TABLE

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type mainPage struct {
	text     string
	datetime time.Time
	table    table.Model
}

func (m mainPage) Init() tea.Cmd {
	return nil
}

func (m mainPage) View() string {
	title := fmt.Sprintf(
		"Последний запрос %s\nu - обновить список\nr - выполнить запросы\n",
		m.datetime.Format(types.DateFormat),
	)
	return title + baseStyle.Render(m.table.View()) + "\n"
}

func (m mainPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			requestAndSave()
			m.datetime = h.Datetime
			m.table = makeTable(h.Domains)
		case "u":
			updateHistory()
			m.datetime = h.Datetime
			m.table = makeTable(h.Domains)
		case "tab":
			m.table.SetCursor((m.table.Cursor() + 1) % len(m.table.Rows()))
		case "shift+tab":
			length := len(m.table.Rows())
			index := m.table.Cursor() - 1
			if index == -1 {
				index = length - 1
			}
			m.table.SetCursor(index % length)
		case "enter":
			return m, tea.Batch(
				tea.Printf("GO to %s", m.table.SelectedRow()[0]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
