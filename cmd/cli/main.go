package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/DevPutat/TLSwatcher/internal/config"
	"github.com/DevPutat/TLSwatcher/internal/logs"
	"github.com/DevPutat/TLSwatcher/internal/notify"
	"github.com/DevPutat/TLSwatcher/internal/report"
	"github.com/DevPutat/TLSwatcher/internal/request"
	"github.com/DevPutat/TLSwatcher/internal/types"
)

func main() {
	flagEdit := flag.Bool("edit", false, "заполнить список доменов")
	flagNotify := flag.Bool("notify", false, "без вывода - только уведомления")

	flag.Parse()
	logs.CreateLogFile()

	if *flagEdit {
		err := config.InputDomains(types.ConfigFilePath)
		if err != nil {
			panic(err)
		}

		fmt.Println("Сохранено!")
		fmt.Println("Выполнить проверку? (y/n)")
		var answer string
		_, err = fmt.Scanln(&answer)
		if err == nil && strings.ToLower(answer) != "y" {
			return
		}
	}

	domains, err := config.Domains(types.ConfigFilePath)
	if err != nil {
		panic(err)
	}
	for i, domain := range domains {
		domains[i] = request.Request(domain)
		if *flagNotify {
			notify.DomainNotify(domains[i])
		}
	}
	if !*flagNotify {
		fmt.Println(report.TextReports(domains))
	}

}
