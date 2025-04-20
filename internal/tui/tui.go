package tui

import (
	"fmt"
	"os"

	"github.com/DevPutat/TLSwatcher/internal/config"
	"github.com/DevPutat/TLSwatcher/internal/history"
	"github.com/DevPutat/TLSwatcher/internal/request"
	"github.com/DevPutat/TLSwatcher/internal/types"
	tea "github.com/charmbracelet/bubbletea"
)

var h types.History

func Run() {
	h = history.Read(types.HistoryFilePath)
	t := makeTable(h.Domains)
	mPage := mainPage{text: "TLSwatcher", table: t, datetime: h.Datetime}
	if _, err := tea.NewProgram(mPage).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func requestAndSave() {
	domains := h.Domains
	for i, domain := range domains {
		domains[i] = request.Request(domain)
	}
	history.Write(types.HistoryFilePath, domains)
	h = history.Read(types.HistoryFilePath)
}

func updateHistory() {
	domains, err := config.Domains(types.ConfigFilePath)
	if err == nil {
		history.Write(types.HistoryFilePath, domains)
	}
	h = history.Read(types.HistoryFilePath)
}
