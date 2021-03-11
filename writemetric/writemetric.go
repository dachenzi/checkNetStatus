package writemetric

import (
	"fmt"
	"github.com/dachenzi/checkNetStatus/config"

	"os"
)

func WriteData(cfg *config.AppConfig) bool {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("get hostname err：", err)
	}

	filename := fmt.Sprintf("%s.$$", cfg.MetricFile)
	fileObj, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("metric file open error: ", err)
		return false
	}
	defer fileObj.Close()

	// 构建 metric 格式
	_, err = fmt.Fprint(fileObj,
		"# HELP This is record used for kubelet err mes: use a closed network connection\n",
		fmt.Sprintf("# TYPE %s counter\n", cfg.MetricKey),
		fmt.Sprintf("%s {host='%s'} %v\n", cfg.MetricKey, hostname, cfg.RestartNum),
	)

	if err != nil {
		fmt.Println("write metric file err: ", err)
		return false
	}

	// rename
	if err := os.Rename(filename, cfg.MetricFile); err != nil {
		fmt.Println("rename metric file err: ", err)
		return false
	}
	return true
}
