package log

import (
	"context"
	"github.com/xiaomeng79/go-log/conf"
	"github.com/xiaomeng79/go-log/plugins/logrus"
	"github.com/xiaomeng79/go-log/plugins/zaplog"
	"testing"
	"time"
)

func TestSetLogger(t *testing.T) {
	//设置为当前目录下 //设置级别
	SetLogger(zaplog.New(
		conf.WithProjectName("zap test"),
		conf.WithLogPath("tmp"),
		conf.WithLogLevel("info"),
		))
	Debug("this is zap")
	Debug("hello",context.Background())
	Infof("hello %s","world",context.Background())
	l2 := logrus.New(conf.WithLogPath("tmp"),conf.WithLogName("logrus"),conf.WithProjectName("logrus test"))
	SetLogger(l2)
	Debugf("this is logrus %s","test",context.Background())
	time.Sleep(time.Second * 5)
}
