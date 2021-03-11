package config

import (
	"gopkg.in/ini.v1"
)

//var (
//	cfg = new(AppConfig)
//)

type AppConfig struct {
	KubeletConfig `ini:"kubelet"`
}

type KubeletConfig struct {
	LogFileName      string `ini:"log"`
	Keyword          string `ini:"keyword"`
	Intervals        int64  `ini:"intervals"`
	Occurrence       int    `ini:"occurrence"`
	IntervalStatus   int
	OccurrenceStatus int
	LastRestartTime  int64
	MetricFile       string `ini:"metricFile"`
	MetricKey        string `ini:"metricKey"`
	RestartNum       int
}

func Init(cfg *AppConfig, conf string) (err error) {
	if conf == "" {
		cfg.LogFileName = "/var/log/kubernetes/kubelet/kubelet.ERROR"
		cfg.Keyword = "use a closed network connection"
		cfg.Intervals = 120
		cfg.Occurrence = 5
	} else {
		if err = ini.MapTo(cfg, conf); err != nil {
			return err
		}
	}
	return
}
