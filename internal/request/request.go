package request

import "github.com/DevPutat/TLSwatcher/internal/types"

func Request(d types.Domain, dCh chan types.Domain) {
	dCh <- d
}
