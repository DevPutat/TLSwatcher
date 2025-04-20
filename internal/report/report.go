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

func TextReports(domains []types.Domain) string {
	res := fmt.Sprintf(headerReport, "URL", "Дата истечения", "STATUS")
	for _, domain := range domains {
		icon := ""
		if domain.IsAttention() {
			icon = types.WarningIcon
		}
		res += fmt.Sprintf(
			rowTemplate,
			domain.Url,
			domain.Expire.Format(types.DateFormat),
			icon,
		)
	}
	return res
}
