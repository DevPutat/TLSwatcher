package types

import (
	"fmt"
	"time"
)

var ConfigFilePath = "domains.ini"
var HistoryFilePath = "history.json"
var LogFilePath = "tlswatcher.log"
var TimeToAttention = time.Hour * 24 * 3
var WarningIcon = "⚠️Warning"
var DateFormat = "02.01.2006"
var DateTimeFormat = "02.01.2006 15:04:05"

type ErrorLog struct {
	Package string
	Err     error
}

func (e ErrorLog) String() string {
	now := time.Now()
	return fmt.Sprintf("%s:[%s]:: %v", now.Format(DateTimeFormat), e.Package, e.Err)
}

type History struct {
	Domains  []Domain  `json:"domains"`
	Datetime time.Time `json:"datetime"`
}

type Domain struct {
	Url         string    `json:"url"`
	Expire      time.Time `json:"expire"`
	IsConnected bool      `json:"is_connected"`
}

func (d Domain) String() string {
	return d.Url
}

func (d Domain) Report() string {
	return fmt.Sprintf("%s: %s", d.Url, d.Expire.Format(DateFormat))
}

func (d Domain) IsAttention() bool {
	return time.Now().Add(TimeToAttention).After(d.Expire)
}
