package validator

import (
	"fmt"
	"time"
)

type Cert struct {
	name string
	exp  time.Duration
	err  error
}

const logFormat = `{` +
	`"status":"%s",` +
	`"certificateCheckHost":"%s",` +
	`"expirationDays":%d,` +
	`"msg":"%s"}`

func (c Cert) String() string {
	days := int64(c.exp.Hours() / 24)
	if c.err != nil {
		return fmt.Sprintf(logFormat, "ng", c.name, days, c.err.Error())
	} else {
		return fmt.Sprintf(logFormat, "ok", c.name, days, "valid certificate")
	}
}

func (c Cert) Invalid() bool {
	return c.err != nil
}
