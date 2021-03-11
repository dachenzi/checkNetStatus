package monitor

import (
	"fmt"
	"github.com/dachenzi/checkNetStatus/config"
	"github.com/dachenzi/checkNetStatus/service"
	"github.com/dachenzi/checkNetStatus/writemetric"
	"github.com/hpcloud/tail"
	"strings"
	"time"
)

var (
	logReader *tail.Tail
)

// open log file
func Init(c *config.AppConfig) (err error) {
	iniConf := tail.Config{
		ReOpen:    true,
		MustExist: true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	logReader, err = tail.TailFile(c.LogFileName, iniConf)
	if err != nil {
		return err
	}
	return
}

// check partition and occurrence
func checkOccurrence(c *config.AppConfig, data string) bool {
	if strings.Contains(data, c.Keyword) {
		c.OccurrenceStatus++
		if c.OccurrenceStatus == c.Occurrence {
			c.OccurrenceStatus = 0 // 重置间隔次数
			return true
		}
	}
	return false
}

// check Intervals
func checkIntervals(c *config.AppConfig) bool {
	now := time.Now().Unix()
	if now-c.LastRestartTime >= c.Intervals {
		c.LastRestartTime = now // 记录重启时间
		return true
	}
	return false
}

func Start(c *config.AppConfig) {
	for {
		select {
		case t := <-logReader.Lines:
			fmt.Println("日志来了:", t.Text)
			if checkOccurrence(c, t.Text) {
				if checkIntervals(c) {
					fmt.Printf("当前次数：%v, 执行重启", c.Occurrence)
					retCode, err := service.RestartKubeletService()
					if err != nil {
						fmt.Println("执行kubelet重启异常: ", err)
					}
					if retCode != 0 {
						fmt.Println("重启kubelet服务异常: ", err)
					}
					// 记录符合node_exporter 收集的信息文件
					c.RestartNum++
					fmt.Printf("当前重启次数 %v,写入文件", c.RestartNum)
					writemetric.WriteData(c)
				}
			}
		default:
		}
	}
}
