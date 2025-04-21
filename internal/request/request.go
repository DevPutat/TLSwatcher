package request

import (
	"crypto/tls"

	"github.com/DevPutat/TLSwatcher/internal/logs"
	"github.com/DevPutat/TLSwatcher/internal/types"
)

func Request(domain types.Domain) types.Domain {
	domain.IsConnected = false
	conn, err := tls.Dial("tcp", domain.Url+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err == nil {
		domain.IsConnected = true
		defer conn.Close()
		certs := conn.ConnectionState().PeerCertificates
		if len(certs) != 0 {
			cert := certs[0]
			domain.Expire = cert.NotAfter
		}
	} else {
		logs.Add(types.ErrorLog{
			Package: "request",
			Err:     err,
		})
	}
	return domain
}
