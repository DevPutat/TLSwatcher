package types

import (
	"fmt"
	"time"
)

var TimeToAttention = time.Hour * 24 * 3

type Domain struct {
	Url         string
	Expire      time.Time
	IsConnected bool
}

func (d Domain) String() string {
	return d.Url
}

func (d Domain) Report() string {
	return fmt.Sprintf("%s: %s", d.Url, d.Expire.Format("09.07.2017"))
}

func (d Domain) IsAttention() bool {
	return time.Now().Add(TimeToAttention).After(d.Expire)
}
