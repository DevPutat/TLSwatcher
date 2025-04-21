package notify

import (
	"fmt"

	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/gen2brain/beeep"
)

var notifyTemplate = "Срок сертификата для %s подходит к концу %s\n"
var noConnTemplate = "Соединение с %s не было выполнено\n"

func Notify(msg string) error {
	// return beeep.Notify("TLSwatcher Alert", msg, "assert/warning.png")
	return beeep.Alert("TLSwatcher Alert", msg, "assert/warning.png")
}

func DomainNotify(domain types.Domain) {
	if !domain.IsConnected {
		msg := fmt.Sprintf(noConnTemplate, domain.Url)
		Notify(msg)
		return
	}
	if domain.IsAttention() {
		msg := fmt.Sprintf(notifyTemplate, domain.Url, domain.Expire.Format(types.DateFormat))
		Notify(msg)
	}
}
