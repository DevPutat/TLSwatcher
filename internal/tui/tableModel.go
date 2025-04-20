package tui

import (
	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var columns = []table.Column{
	{Title: "URL", Width: 50},
	{Title: "Дата истечения", Width: 20},
	{Title: "Статус", Width: 20},
}

func makeData(domains []types.Domain) []table.Row {
	rows := []table.Row{}
	for _, d := range domains {

		icon := ""
		if d.IsAttention() {
			icon = types.WarningIcon
		}
		rows = append(rows, table.Row{d.Url, d.Expire.Format(types.DateFormat), icon})
	}
	return rows
}

func makeTable(domains []types.Domain) table.Model {
	data := makeData(domains)
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(data),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}
