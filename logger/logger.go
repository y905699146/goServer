package logger

import (
	"fmt"
	"log"
)

const (
	KDebug = iota
	KInfo
	KWarn
	KError
	KFatal
)

var logPrefix = []string{
	"[DBG] ",
	"[INF] ",
	"[WRN] ",
	"[ERR] ",
	"[FAL] ",
}

type ILogger interface {
	Output(level int, calldepth int, f string) error
}

type defaultLogger struct {
}

func (d *defaultLogger) Output(level int, calldepth int, f string) error {
	text := logPrefix[level] + " " + f
	return log.Output(calldepth, text)
}

var myDefaultLogger defaultLogger
var myLogger ILogger = &myDefaultLogger

func SetLogger(logger ILogger) {
	myLogger = logger
}

func _log(level int, calldepth int, f string) {
	myLogger.Output(level, calldepth, f)
}

func Debug_MSG(f string, v ...interface{}) {
	_log(KDebug, 2, fmt.Sprintf(f, v...))
}

func Error_MSG(f string, v ...interface{}) {
	_log(KError, 2, fmt.Sprintf(f, v...))
}

func Info_MSG(f string, v ...interface{}) {
	_log(KInfo, 2, fmt.Sprintf(f, v...))
}

func Fatal_MSG(f string, v ...interface{}) {
	_log(KFatal, 2, fmt.Sprintf(f, v...))
}

func Warn_MSG(f string, v ...interface{}) {
	_log(KWarn, 2, fmt.Sprintf(f, v...))
}
