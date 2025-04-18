package request

import (
	"testing"

	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	domain := types.Domain{
		Url: "google.com",
	}

	res := Request(domain)
	assert.False(t, res.Expire.IsZero(), "не-написано время истечения сертификата")
}
