package report

import (
	"strings"
	"testing"
	"time"

	"github.com/DevPutat/TLSwatcher/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestTextReports(t *testing.T) {
	now := time.Now()
	validTime := now.Add(types.TimeToAttention).Add(time.Minute) // Время, которое не требует внимания
	domains := []types.Domain{
		{
			Url:    "valid.test",
			Expire: validTime,
		},
		{
			Url:    "invalid.test",
			Expire: now, // Истекший срок действия
		},
	}

	text := TextReports(domains)
	count := strings.Count(text, "Warning")

	assert.Equal(t, 1, count, "слово 'Warning' должно встречаться ровно один раз")

	assert.Contains(t, text, "Warning", "строка с 'invalid.test' должна содержать предупреждение")
}
