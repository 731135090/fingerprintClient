package config

import (
	"fingerprintClient/action"
	"fingerprintClient/service"
	"fingerprintClient/util/dir"
	"fmt"
	"github.com/spf13/viper"
)

var (
	RootDir        string
	ProjectBaseDir string
	ServerConfig   map[string]string
)

func Init() {
	// flag 命令行解析
	InitFlag()
	// 初始化日志
	initLogger()
	// 初始化配置
	InitConfig()
}

func initLogger() {
	// TODO 初始化日志
}

func InitConfig() {
	RootDir = dir.GetPwd()
	// 加载配置文件
	loadConfigFile()
	if RunLevel != RUN_LEVEL_RELEASE {
		fmt.Printf("RootPath: %s\n", RootDir)
		fmt.Printf("run project: %s；project base dir：%s\n", Project, ProjectBaseDir)
	}
	// 初始化Server
	service.New(ServerConfig)
	// 加载项目配置
	action.LoadProjectConfig()
}

func loadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/conf/%s/", RootDir, Project))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if RunLevel == "" {
		RunLevel = viper.GetString("run_level")
	}
	ProjectBaseDir = viper.GetString("project_base_dir")
	ServerConfig = viper.GetStringMapString("server")
}
