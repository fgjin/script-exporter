package logs

import "io"

type LogWriter interface {
	Flush()
	io.Writer
}
type InitLogWriterFunc func() LogWriter

var writerAdapter = make(map[string]InitLogWriterFunc)

func RegisterInitWriterFunc(adapterName string, writerFunc InitLogWriterFunc) {
	writerAdapter[adapterName] = writerFunc
}
