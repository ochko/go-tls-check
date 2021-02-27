package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ochko/go-tls-check/validator"
)

var (
	alertWindowStr string
	connTimeoutStr string
	quietMode      bool
)

func init() {
	flag.StringVar(&alertWindowStr, "w", "72h", "Allowd time before certificate expiration.")
	flag.StringVar(&connTimeoutStr, "t", "10s", "Connection timeout.")
	flag.BoolVar(&quietMode, "q", false, "Quiet mode, no output unless validation fails.")
}

func main() {
	flag.Parse()

	hostnames := flag.Args()
	if len(hostnames) == 0 {
		exit("No hostname is given.")
	}

	alertWindow, err := time.ParseDuration(alertWindowStr)
	if err != nil {
		exit("Invalid value for option -w.")
	}

	connTimeout, err := time.ParseDuration(connTimeoutStr)
	if err != nil {
		exit("Invalid value for option -t.")
	}

	exitCode := 0
	l := log.New(os.Stdout, "", 0)

	for _, name := range hostnames {
		cert := validator.NewCert(name, alertWindow, connTimeout)
		if !quietMode {
			l.Print(cert)
		}
		if cert.Invalid() {
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}

func exit(msg string) {
	fmt.Println(msg, "Usage:")
	fmt.Println("tls-check [options] hostname1 hostname2 ...\n  options:")
	flag.PrintDefaults()
	os.Exit(2)
}
