package validator

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"
)

const AlertWindow = time.Hour * 24 * 14 // 14 days
const DialerTimeout = time.Second * 10
const TimeFormat = "2006-01-02 15PM MST"

func Check(name string) (exp time.Duration, err error) {
	dialer := &net.Dialer{Timeout: DialerTimeout}
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

		if exp < AlertWindow {
			err = errors.New(fmt.Sprintf("expires at %v", c.NotAfter.Format(TimeFormat)))
			break
		}
	}
	return
}
