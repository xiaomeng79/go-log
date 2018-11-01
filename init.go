package log

import (
	"github.com/xiaomeng79/go-log/plugins/zaplog"
)

//默认
var l ILog = zaplog.New()

//设置
func SetLogger(ll ILog) {
	l = ll
}

//普通日志
func Debug(msg string, args ...interface{}) {
	l.Debug(msg, args...)
}
func Info(msg string, args ...interface{}) {
	l.Info(msg, args...)
}
func Warn(msg string, args ...interface{}) {
	l.Warn(msg, args...)
}
func Error(msg string, args ...interface{}) {
	l.Error(msg, args...)
}
func Panic(msg string, args ...interface{}) {
	l.Panic(msg, args...)
}
func Fatal(msg string, args ...interface{}) {
	l.Fatal(msg, args...)
}

//其他日志 如：HTTP RPC日志
func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
