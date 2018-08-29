package zaplog

import "go.uber.org/zap/zapcore"

type Option func(*Options)

type Options struct {
	LogPath string //日志保存路径
	LogName string //日志保存的名称，不些随机生成
	LogLevel zapcore.Level //日志记录级别
	MaxSize int //日志分割的尺寸 MB
	MaxAge int //分割日志保存的时间 day
	Stacktrace zapcore.Level //记录堆栈的级别
	IsStdOut string //是否标准输出console输出
}

func WithLogPath(logpath string) Option {
	return func(o *Options) {
		o.LogPath = logpath
	}
}

func WithLogName(logname string) Option {
	return func(o *Options) {
		o.LogName = logname
	}
}

func WithLogLevel(loglevel zapcore.Level) Option {
	return func(o *Options) {
		o.LogLevel = loglevel
	}
}

func WithMaxSize(maxsize int)  Option{
	return func(o *Options) {
		o.MaxAge = maxsize
	}
}

func WithMaxAge(maxage int) Option {
	return func(o *Options) {
		o.MaxAge = maxage
	}
}

func WithStacktrace(stacktrace zapcore.Level) Option {
	return func(o *Options) {
		o.Stacktrace = stacktrace
	}
}

func WithIsStdOut(isstdout string) Option {
	return func(o *Options) {
		o.IsStdOut = isstdout
	}
}
