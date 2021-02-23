package validator

import (
	"strings"
	"testing"
	"time"
)

var (
	window, timeout time.Duration
	_               error
)

func TestCheckUnknownHost(t *testing.T) {
	window, _ = time.ParseDuration("72h")
	timeout, _ = time.ParseDuration("10s")

	host := "unknown.no-such-domain.com"
	exp, err := Check(host, window, timeout)
	if err == nil {
		t.Errorf("Check(%s) succeeded, want error", host)
	}
	if exp != 0 {
		t.Errorf("Check(%s) returned %v, want %v", host, exp, 0)
	}
}

func TestExpireSoon(t *testing.T) {
	window, _ = time.ParseDuration("876582h")
	timeout, _ = time.ParseDuration("10s")

	host := "example.com"
	exp, err := Check(host, window, timeout)
	if err == nil {
		t.Errorf("Check(%s) succeeded, want error", host)
	}
	if !strings.Contains(err.Error(), "expires at") {
		t.Errorf("Check(%s) want expires at", host)
	}
	if exp == 0 {
		t.Errorf("Check(%s) returned 0 want non zero", host)
	}
}

func TestValid(t *testing.T) {
	window, _ = time.ParseDuration("72h")
	timeout, _ = time.ParseDuration("10s")

	host := "example.com"
	exp, err := Check(host, window, timeout)
	if err != nil {
		t.Errorf("Check(%s) failed with %v, want success", host, err)
	}
	if exp == 0 {
		t.Errorf("Check(%s) returned 0 want non zero", host)
	}
}
