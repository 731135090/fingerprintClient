package config

import (
	"fmt"
	"os"
)

var RootPath string

func Init() {
	// 初始化配置
	InitConfig()
	// flag 命令行解析
	InitFlag()
}

func InitConfig() {
	RootPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if RunLevel != RUN_LEVEL_RELEASE {
		fmt.Printf("RootPath: %s\n", RootPath)
	}
}
