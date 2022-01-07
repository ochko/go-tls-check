package validator

import (
	"encoding/json"
	"fmt"
	"time"
)

type Cert struct {
	name string
	exp  time.Duration
	err  error
}

func (c Cert) String() string {
	s := map[string]string{
		"host":       c.name,
		"expiration": c.exp.String(),
		"status":     "ok",
	}

	if c.err != nil {
		s["status"] = "ng"
		s["msg"] = c.err.Error()
	}

	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("%v", s)
	}
	return string(b)
}

func (c Cert) Invalid() bool {
	return c.err != nil
}
