package types

import "time"

type Domain struct {
	Url       string
	Expire    time.Time
	Connected bool
	Code      int
}

func (d Domain) String() string {
	return d.Url
}
