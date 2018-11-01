package logrus

import (
	"bufio"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/xiaomeng79/go-log/conf"
	"github.com/xiaomeng79/go-log/fileout"
	"os"
)

type Log struct {
	logger      *logrus.Logger
	ProjectName string
}

//初始化日志
func New(opts ...conf.Option) *Log {
	log := &Log{}
	log.logger = logrus.New()

	o := &conf.Options{
		LogPath:     conf.LogPath,
		LogName:     conf.LogName,
		LogLevel:    conf.LogLevel,
		MaxSize:     conf.MaxSize,
		MaxAge:      conf.MaxAge,
		IsStdOut:    conf.IsStdOut,
		ProjectName: conf.ProjectName,
	}
	for _, opt := range opts {
		opt(o)
	}
	//设置项目名称
	log.ProjectName = o.ProjectName
	//所有日志都输出到文件
	lev, err := logrus.ParseLevel(o.LogLevel)
	if err != nil {
		panic(err.Error())
	}
	log.logger.Level = lev
	if o.IsStdOut != "yes" {
		log.logger.Out = setNull() //将日志写入空接口
	}
	writer := fileout.NewRollingFile(o.LogPath, o.LogName, o.MaxSize, o.MaxAge)

	log.logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: writer,
			logrus.FatalLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.WarnLevel:  writer,
			logrus.InfoLevel:  writer,
			logrus.DebugLevel: writer,
		},
		&logrus.JSONFormatter{},
	))
	return log
}

//设置一个空接口，将日志写入空接口

func setNull() *bufio.Writer {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	return writer
}
