package config

import (
	"fmt"
)

var RootPath string

func Init() {
	// flag 命令行解析
	InitFlag()
	// 初始化配置
	InitConfig()
}

func InitConfig() {
	if RunLevel != RUN_LEVEL_RELEASE {
		fmt.Printf("RootPath: %s\n", RootDir)
		fmt.Printf("run project: %s；project base dir：%s\n", Project, ProjectBaseDir)
	}
}
