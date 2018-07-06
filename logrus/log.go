package logrus

import (
	"os"
	"path/filepath"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"bufio"
	"time"
)

type Log struct {
	logger *logrus.Logger
}
//初始化日志
func New(opts ...Option) *Log {
	log := &Log{}
	log.logger = logrus.New()

	o := &Options{
		LogPath:"tmp/log/",
		LogName:"output.log",
		RotationTime:1,
		MaxAge:7,
		IsStdOut:true,
	}
	for _,opt := range opts {
		opt(o)
	}
	//所有日志都输出到文件
	log.logger.Level = logrus.DebugLevel
	if !o.IsStdOut {
		log.logger.Out = setNull()//将日志写入空接口
	}

	err := os.MkdirAll(o.LogPath,0766)
	if err != nil {
		panic(err)
		return log
	}
	logpath := filepath.Join(o.LogPath,o.LogName)
	logpath,_ = filepath.Abs(logpath)
	writer, _ := rotatelogs.New(
		logpath+".%Y%m%d%H%M" + ".log",
		rotatelogs.WithLinkName(logpath),
		rotatelogs.WithMaxAge(time.Duration(o.MaxAge * 24) * time.Hour),//日志保留时间 天
		rotatelogs.WithRotationTime(time.Duration(o.RotationTime * 24)*time.Hour),//日志切割时间 天
	)


	log.logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: writer,
			logrus.FatalLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.WarnLevel: writer,
			logrus.InfoLevel: writer,
			logrus.DebugLevel: writer,
		},
		&logrus.JSONFormatter{},
	))
	return log
}

//设置一个空接口，将日志写入空接口

func setNull() *bufio.Writer{
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		panic(err)
		return nil
	}
	writer := bufio.NewWriter(src)
	return writer
}
