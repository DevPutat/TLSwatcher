package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/DevPutat/TLSwatcher/internal/config"
	"github.com/DevPutat/TLSwatcher/internal/report"
	"github.com/DevPutat/TLSwatcher/internal/request"
	"github.com/DevPutat/TLSwatcher/internal/types"
)

func main() {
	flagEdit := flag.Bool("edit", false, "заполнить список доменов")

	flag.Parse()

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
	}
	fmt.Println(report.TextReports(domains))
}
