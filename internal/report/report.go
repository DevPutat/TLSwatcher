package report

import (
	"fmt"

	"github.com/DevPutat/TLSwatcher/internal/types"
)

var headerReport = `
===============================================================================
                     TLSwatcher - Отчет по SSL-сертификатам                    
===============================================================================
Домены просканированы. 
Истекший срок, а так же срок скоро истекающий отмечены иконкой

%-30s | %-20s | %-10s
===============================================================================
`
var rowTemplate = "%-30s | %-20s | %-10s\n"
var warningIcon = "⚠️Warning"

func TextReports(domains []types.Domain) string {
	res := fmt.Sprintf(headerReport, "URL", "Дата истечения", "STATUS")
	for _, domain := range domains {
		icon := ""
		if domain.IsAttention() {
			icon = warningIcon
		}
		res += fmt.Sprintf(
			rowTemplate,
			domain.Url,
			domain.Expire.Format("09.07.2017"),
			icon,
		)
	}
	return res
}
