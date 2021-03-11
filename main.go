package main

import (
	"flag"
	"fmt"
	"github.com/dachenzi/checkNetStatus/config"
	"github.com/dachenzi/checkNetStatus/monitor"
)

var (
	cfg = new(config.AppConfig)
)

func main() {
	// 0、解析命令行参数
	var cmdConf string
	flag.StringVar(&cmdConf, "conf", "", "--conf ./config.ini")
	flag.Parse()

	// 1、读取配置文件确认kubelet文件位置
	err := config.Init(cfg, cmdConf)
	if err != nil {
		fmt.Println("配置文件初始化失败: ", err)
		return
	}
	fmt.Println(cfg)

	// 2、根据配置文件初始化追踪文件
	err = monitor.Init(cfg)
	if err != nil {
		fmt.Println("读取日志文件失败: ", err)
		return
	}

	// 3、启动日志追踪，发现异常日志，重启kubelet，继续监控
	monitor.Start(cfg)
}
