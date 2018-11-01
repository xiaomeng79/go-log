package conf

import (
	"fmt"
	"strings"
)

//日志级别
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}

	return "unknown"
}

func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "panic", "dpanic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
}

var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
}

//默认参数
const (
	LogPath     string = "/var/log" //日志保存路径
	LogName     string = "output"   //日志保存的名称，不些随机生成
	LogLevel    string = "debug"    //日志记录级别
	MaxSize     int    = 100        //日志分割的尺寸 MB
	MaxAge      int    = 7          //分割日志保存的时间 day
	Stacktrace  string = "error"    //记录堆栈的级别
	IsStdOut    string = "yes"      //是否标准输出console输出
	ProjectName string = "test"     //项目名称

)
