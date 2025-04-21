package notify

import (
	"github.com/gen2brain/beeep"
)

func Notify(msg string) error {
	// return beeep.Notify("TLSwatcher Alert", msg, "assert/warning.png")
	return beeep.Alert("TLSwatcher Alert", msg, "assert/warning.png")
}
