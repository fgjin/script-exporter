package logs

import "os"

type stdWriter struct {
	*os.File
}

func (s *stdWriter) Flush() {
	s.Sync()
}

func newStdWriter() LogWriter {
	return &stdWriter{
		os.Stderr,
	}
}

func init() {
	RegisterInitWriterFunc("std", newStdWriter)
}
