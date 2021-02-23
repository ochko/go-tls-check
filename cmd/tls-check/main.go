package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ochko/go-tls-check/validator"
)

const LogFormat = "{\"status\":\"%s\",\"certificateCheckHost\":\"%s\",\"expirationDays\":%d,\"msg\":\"%v\"}"

func main() {
	var (
		alertWindowStr string
		connTimeoutStr string
	)

	flag.StringVar(&alertWindowStr, "w", "72h", "Allowd time before certificate expiration.")
	flag.StringVar(&connTimeoutStr, "t", "10s", "Connection timeout.")
	flag.Parse()

	hostnames := flag.Args()
	alertWindow, errParse1 := time.ParseDuration(alertWindowStr)
	connTimeout, errParse2 := time.ParseDuration(connTimeoutStr)

	if len(hostnames) == 0 || errParse1 != nil || errParse2 != nil {
		fmt.Println("Usage:\ntls-check [options] hostname1 hostname2 ...\n  options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	status := 0
	l := log.New(os.Stdout, "", 0)

	for _, name := range hostnames {
		exp, err := validator.Check(name, alertWindow, connTimeout)
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
