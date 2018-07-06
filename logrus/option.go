package logrus


type Option func(*Options)

type Options struct {
	LogPath string //日志保存路径
	LogName string //日志保存的名称，不些随机生成
	RotationTime int //分割日志的时间 day
	MaxAge int //日志最大的时间 day
	IsStdOut bool //是否标准输出console输出
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

func WithRotationTime(rotationtime int) Option {
	return func(o *Options) {
		o.RotationTime = rotationtime
	}
}

func WithMaxAge(maxage int) Option {
	return func(o *Options) {
		o.MaxAge = maxage
	}
}

func WithIsStdOut(isstdout bool) Option {
	return func(o *Options) {
		o.IsStdOut = isstdout
	}
}
