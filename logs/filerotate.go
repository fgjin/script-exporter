package logs

import (
	"script-exporter/global"

	"github.com/natefinch/lumberjack"
)

type fileRotateWriter struct {
	*lumberjack.Logger
}

func (f *fileRotateWriter) Flush() {
	f.Close()
}

func newFileRotateWriter() LogWriter {
	writer := getWriter()
	return &fileRotateWriter{
		writer,
	}
}

func getWriter() *lumberjack.Logger {
	logRotate := &lumberjack.Logger{
		Filename:   global.LogPath,
		MaxSize:    global.MaxSize, // MB
		MaxBackups: global.MaxBackups,
		MaxAge:     global.MaxAge,   //days
		Compress:   global.Compress, // disabled by default
	}
	return logRotate
}
func init() {
	RegisterInitWriterFunc("fileRotate", newFileRotateWriter)
}
