package main

import (
	"fingerprintClient/config"
	"fmt"
	"time"
)

//main
func main() {
	// 初始化
	config.Init()

	// 常驻进程
	for {
		// todo 生成文件列表，任何上报
		fmt.Println(time.Now().Unix())
		// sleep 1分钟
		time.Sleep(time.Minute)
	}
}

