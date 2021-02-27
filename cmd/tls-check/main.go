package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ochko/go-tls-check/validator"
)

func main() {
	var (
		alertWindowStr string
		connTimeoutStr string
	)

	flag.StringVar(&alertWindowStr, "w", "72h", "Allowd time before certificate expiration.")
	flag.StringVar(&connTimeoutStr, "t", "10s", "Connection timeout.")
	flag.Parse()

	hostnames := flag.Args()
	if len(hostnames) == 0 {
		exit()
	}

	alertWindow, err := time.ParseDuration(alertWindowStr)
	if err != nil {
		exit()
	}

	connTimeout, err := time.ParseDuration(connTimeoutStr)
	if err != nil {
		exit()
	}

	status := 0
	l := log.New(os.Stdout, "", 0)

	for _, name := range hostnames {
		cert := validator.NewCert(name, alertWindow, connTimeout)
		l.Print(cert)
		if cert.Invalid() {
			status = 1
		}
	}

	os.Exit(status)
}

func exit() {
	fmt.Println("Usage:\ntls-check [options] hostname1 hostname2 ...\n  options:")
	flag.PrintDefaults()
	os.Exit(1)
}
