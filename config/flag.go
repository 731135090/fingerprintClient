package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	file    string
	test    bool
	Debug   bool
	daemon  bool
	help    bool
	version bool
)

func init() {
	flag.StringVar(&file, "f", "./config.yaml", "Location of client config file")
	flag.BoolVar(&test, "t", false, "Test config file")
	flag.BoolVar(&Debug, "D", false, "Enable debug mode")
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
	// 检查是否有配置文件
	checkConfigFile()
}

func checkConfigFile() {
	absConfigFile, err := filepath.Abs(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if _, err = os.Stat(absConfigFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// todo 解析配置文件
}

func usage() {
	fmt.Fprintf(os.Stderr, `fpClient
Usage: fpClient [-hv] [-f filename] [-d daemon] [-D debug]

Options:
`)
	flag.PrintDefaults()
}
