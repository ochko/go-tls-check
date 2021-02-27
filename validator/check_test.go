package validator

import (
	"strings"
	"testing"
	"time"
)

func TestCheckUnknownHost(t *testing.T) {
	host := "unknown.no-such-domain.com"
	exp, err := Check(host, time.Hour*24, time.Second*3)
	if err == nil {
		t.Errorf("Check(%s) succeeded, want error", host)
	}
	if exp != 0 {
		t.Errorf("Check(%s) returned %v, want %v", host, exp, 0)
	}
}

func TestExpireSoon(t *testing.T) {
	host := "example.com"
	exp, err := Check(host, time.Hour*876582, time.Second*10)
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
	host := "example.com"
	exp, err := Check(host, time.Hour*24, time.Second*10)
	if err != nil {
		t.Errorf("Check(%s) failed with %v, want success", host, err)
	}
	if exp == 0 {
		t.Errorf("Check(%s) returned 0 want non zero", host)
	}
}
