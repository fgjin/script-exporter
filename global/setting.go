package global

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	Registry    = prometheus.NewRegistry()
	Respath     string
	LogPath     = "promlog/prom.log"
	AdapterName = "fileRotate" //std|fileRotate
	Level       = logrus.InfoLevel
	Port        = ":8080"
	Compress    = true
)

const (
	MaxSize    = 5
	MaxBackups = 3
	MaxAge     = 3
	Cost       = 10
)
