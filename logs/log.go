package logs

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Log struct {
	*logrus.Entry
	// *logrus.Logger
	LogWriter
}

func (l *Log) Flush() {
	l.LogWriter.Flush()
}

type LogConf struct {
	Level       logrus.Level
	AdapterName string
}

func InitLog(conf LogConf) *Log {
	adapterName := "std"
	if conf.AdapterName != "" {
		adapterName = conf.AdapterName
	}
	writer, ok := writerAdapter[adapterName]
	if !ok {
		adapterName = "std"
		writer, _ = writerAdapter[adapterName]
	}
	log := &Log{
		logrus.NewEntry(logrus.New()),
		writer(),
	}
	// 设置 logrus 的 io.Writer
	if conf.AdapterName == "std" {
		log.Logger.SetOutput(log.LogWriter)
	} else {
		// 同时输出到标准输出与文件
		log.Logger.SetOutput(io.MultiWriter(log.LogWriter, os.Stdout))
	}

	if conf.Level != 0 {
		log.Logger.SetLevel(conf.Level)
	}
	// log.Logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", ForceColors: true})
	log.Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	//打印文件、行号和主调函数
	log.Logger.SetReportCaller(true)
	return log
}
