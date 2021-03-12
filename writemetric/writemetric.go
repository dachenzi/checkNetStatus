package writemetric

import (
	"fmt"
	"github.com/dachenzi/checkNetStatus/config"

	"os"
)

// InitMetricFile ...
func InitMetricFile(cfg *config.AppConfig) {
	// check metric file exist
	hostname, _ := os.Hostname()
	_, err := os.Stat(cfg.MetricFile)
	if err != nil {
		if os.IsNotExist(err) {
			fileObj, err := os.OpenFile(cfg.MetricFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("init metric file err ", err)
			}
			// 构建 metric 格式
			_, err = fmt.Fprint(fileObj,
				"# HELP This is record used for kubelet err mes: use a closed network connection\n",
				fmt.Sprintf("# TYPE %s counter\n", cfg.MetricKey),
				fmt.Sprintf("%s {host='%s'} %v\n", cfg.MetricKey, hostname, 0),
			)
		}
	}
}

// WriteMetricData ...
func WriteMetricData(cfg *config.AppConfig) bool {
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
