package validator

import (
	"encoding/json"
	"errors"
	"testing"
	"time"
)

func TestOk(t *testing.T) {
	c := Cert{"example.com", time.Hour * 24, nil}
	var m map[string]string
	_ = json.Unmarshal([]byte(c.String()), &m)

	expected := "example.com"
	actual := m["host"]
	if actual != expected {
		t.Errorf("Cert.String() json 'host' field should be %s, but got %s", expected, actual)
	}

	expected = "24h0m0s"
	actual = m["expiration"]
	if actual != expected {
		t.Errorf("Cert.String() json 'expiration' field should be %s, but got %s", expected, actual)
	}

	expected = "ok"
	actual = m["status"]
	if actual != expected {
		t.Errorf("Cert.String() json 'status' field should be %s, but got %s", expected, actual)
	}
}

func TestNg(t *testing.T) {
	err := errors.New("expires at 2020-01-01 10PM JST")
	c := Cert{"example.com", time.Hour * (-48), err}
	var m map[string]string
	_ = json.Unmarshal([]byte(c.String()), &m)

	expected := "example.com"
	actual := m["host"]
	if actual != expected {
		t.Errorf("Cert.String() json 'host' field should be %s, but got %s", expected, actual)
	}

	expected = "-48h0m0s"
	actual = m["expiration"]
	if actual != expected {
		t.Errorf("Cert.String() json 'expiration' field should be %s, but got %s", expected, actual)
	}

	expected = "ng"
	actual = m["status"]
	if actual != expected {
		t.Errorf("Cert.String() json 'status' field should be %s, but got %s", expected, actual)
	}

	expected = "expires at 2020-01-01 10PM JST"
	actual = m["msg"]
	if actual != expected {
		t.Errorf("Cert.String() json 'msg' field should be %s, but got %s", expected, actual)
	}

}
