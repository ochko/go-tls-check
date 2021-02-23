package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ochko/go-tls-check/validator"
)

const LogFormat = "{\"status\":\"%s\",\"certificateCheckHost\":\"%s\",\"expirationDays\":%d,\"msg\":\"%v\"}"

func main() {
	status := 0
	hostnames := os.Args[1:]

	if len(hostnames) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: tls-cert-check hostname1 hostname2 ...\n")
		os.Exit(1)
	}

	l := log.New(os.Stdout, "", 0)

	for _, name := range hostnames {
		exp, err := validator.Check(name)
		expirationDays := int64(exp.Hours() / 24)

		if err != nil {
			status = 1
			l.Printf(LogFormat, "ng", name, expirationDays, err)
		} else {
			l.Printf(LogFormat, "ok", name, expirationDays, "valid certificate")
		}
	}

	os.Exit(status)
}
