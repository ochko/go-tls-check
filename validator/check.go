package validator

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"
)

const TimeFormat = "2006-01-02 15PM MST"

func NewCert(name string, alertWindow time.Duration, connTimeout time.Duration) Cert {
	exp, err := Check(name, alertWindow, connTimeout)
	return Cert{name: name, exp: exp, err: err}
}

func Check(name string, alertWindow time.Duration, connTimeout time.Duration) (exp time.Duration, err error) {
	dialer := &net.Dialer{Timeout: connTimeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", name+":443", nil)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.VerifyHostname(name)
	if err != nil {
		return
	}

	now := time.Now()
	for _, c := range conn.ConnectionState().PeerCertificates {
		if now.Before(c.NotBefore) {
			err = errors.New(fmt.Sprintf("not usable until %v", c.NotBefore.Format(TimeFormat)))
			break
		}

		sub := c.NotAfter.Sub(now)
		if exp == 0 || exp > sub {
			exp = sub
		}

		if exp < alertWindow {
			err = errors.New(fmt.Sprintf("expires at %v", c.NotAfter.Format(TimeFormat)))
			break
		}
	}
	return
}
