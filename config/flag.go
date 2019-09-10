package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	file    string
	test    bool
	daemon  bool
	help    bool
	version bool

	Project  string
	RunLevel string
)

func init() {
	flag.StringVar(&file, "f", "", "Location of client config file")
	flag.StringVar(&Project, "p", "test", "Run project")
	flag.BoolVar(&test, "t", false, "Test config file")
	flag.StringVar(&RunLevel, "l", "", "Enable debug mode")
	flag.BoolVar(&daemon, "d", false, "Run client in background")
	flag.BoolVar(&help, "h", false, "Print Usage")
	flag.BoolVar(&version, "v", false, "Print version information and quit")
	flag.Usage = usage
}

func InitFlag() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		ShowVersion()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `fpClient
Usage: fpClient [-hv] [-ft filename] [-d daemon]

Options:
`)
	flag.PrintDefaults()
}
