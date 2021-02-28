package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ochko/go-tls-check/validator"
)

var (
	fs          *flag.FlagSet
	quiet       bool
	alertWindow time.Duration
	connTimeout time.Duration
)

func init() {
	fs = flag.NewFlagSet("tls-check", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] hostname [hostname2 ...]\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.DurationVar(&alertWindow, "w", time.Hour*72, "Allowd time before certificate expiration.")
	fs.DurationVar(&connTimeout, "t", time.Second*10, "Connection timeout.")
	fs.BoolVar(&quiet, "q", false, "Quiet mode, no output unless validation fails.")
	fs.Parse(os.Args[1:])
	if len(fs.Args()) == 0 {
		fs.Usage()
		os.Exit(2)
	}
}

func main() {
	exitCode := 0

	for _, name := range fs.Args() {
		cert := validator.NewCert(name, alertWindow, connTimeout)
		if !quiet {
			fmt.Println(cert)
		}
		if cert.Invalid() {
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
