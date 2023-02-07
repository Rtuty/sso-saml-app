package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"sync"
)

type LoggerEx struct {
	sync.Mutex
	enabled bool
	Logger  *lumberjack.Logger
}

var _ io.WriteCloser = (*LoggerEx)(nil)

func (l *LoggerEx) Write(p []byte) (int, error) {
	if l.enabled {
		if n, err := l.Logger.Write(p); err != nil {
			return n, err
		}
	}
	return fmt.Print(string(p))
}

func (l *LoggerEx) Close() error {
	if l.enabled {
		return l.Logger.Close()
	}
	return nil
}

func (l *LoggerEx) SetEnabled(value bool) {
	l.Lock()
	defer l.Unlock()
	l.enabled = value
}
