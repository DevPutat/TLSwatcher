package request

import (
	"crypto/tls"

	"github.com/DevPutat/TLSwatcher/internal/types"
)

func Request(domain types.Domain, domainCh chan types.Domain) {
	conn, err := tls.Dial("tcp", domain.Url+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err == nil {
		defer conn.Close()
		certs := conn.ConnectionState().PeerCertificates
		if len(certs) != 0 {
			cert := certs[0]
			domain.Expire = cert.NotAfter
		}
	}
	domainCh <- domain
}
