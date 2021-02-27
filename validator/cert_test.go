package validator

import (
	"errors"
	"testing"
	"time"
)

func TestOk(t *testing.T) {
	c := Cert{"example.com", time.Hour * 24, nil}
	expected := `{"status":"ok","certificateCheckHost":"example.com","expirationDays":1,"msg":"valid certificate"}`
	actual := c.String()
	if actual != expected {
		t.Errorf("Cert.String() should return %s, but got %s", expected, actual)
	}
}

func TestNg(t *testing.T) {
	err := errors.New("expires at 2020-01-01 10PM JST")
	c := Cert{"example.com", time.Hour * (-48), err}
	expected := `{"status":"ng","certificateCheckHost":"example.com","expirationDays":-2,"msg":"expires at 2020-01-01 10PM JST"}`
	actual := c.String()
	if actual != expected {
		t.Errorf("Cert.String()\nexpected: %s\n but got: %s", expected, actual)
	}
}
