package report

import (
	"fmt"

	"github.com/DevPutat/TLSwatcher/internal/notify"
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
var notifyTemplate = "Срок сертификата для %s подходит к концу %s\n"

func TextReports(domains []types.Domain) string {
	res := fmt.Sprintf(headerReport, "URL", "Дата истечения", "STATUS")
	for _, domain := range domains {
		d := domain.Expire.Format(types.DateFormat)
		icon := ""
		if domain.IsAttention() {
			icon = types.WarningIcon
			msg := fmt.Sprintf(notifyTemplate, domain.Url, d)
			notify.Notify(msg)
		}
		res += fmt.Sprintf(
			rowTemplate,
			domain.Url,
			d,
			icon,
		)
	}
	return res
}
