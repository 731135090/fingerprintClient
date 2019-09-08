package config

import (
	"fingerprintClient/util/dir"
	"flag"
	"fmt"
	"github.com/spf13/viper"
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
	RunDir   string
	RootDir  string

	ProjectBaseDir string
)

func init() {
	flag.StringVar(&file, "f", "./config.yaml", "Location of client config file")
	flag.StringVar(&Project, "p", "test", "Run project")
	flag.BoolVar(&test, "t", false, "Test config file")
	flag.StringVar(&RunLevel, "l", "", "Enable debug mode")
	flag.BoolVar(&daemon, "d", false, "Run client in background")
	flag.BoolVar(&help, "h", false, "Print Usage")
	flag.BoolVar(&version, "v", false, "Print version information and quit")
	flag.Usage = usage

	RunDir = dir.GetPwd()
	RootDir = dir.GetParentDir(RunDir, 1)
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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/conf/%s/", RootDir, Project))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}


	if RunLevel == "" {
		RunLevel = viper.GetString("run_level")
	}
	ProjectBaseDir = viper.GetString("project_base_dir")
}

func usage() {
	fmt.Fprintf(os.Stderr, `fpClient
Usage: fpClient [-hv] [-f filename] [-d daemon] [-D debug]

Options:
`)
	flag.PrintDefaults()
}
