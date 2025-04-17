package request

import (
	"fmt"
	"testing"
	"time"

	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	domain := types.Domain{
		Url: "google.com",
	}

	domainCh := make(chan types.Domain, 1)

	go Request(domain, domainCh)
	select {
	case res := <-domainCh:
		fmt.Println(res.Expire)
		assert.False(t, res.Expire.IsZero(), "не-аписано время истечения сертификата")
	case <-time.After(5 * time.Second):
		// Если прошло более 5 секунд, тест завершается с ошибкой
		t.Fatal("Тест не завершился в течение 5 секунд")
	}
}
