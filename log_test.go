package log

import (
	"context"
	"github.com/xiaomeng79/go-log/conf"
	"github.com/xiaomeng79/go-log/plugins/logrus"
	"github.com/xiaomeng79/go-log/plugins/zaplog"
	"github.com/xiaomeng79/go-log/tracer"
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
	Debug("hello", context.Background())
	//自定义跟踪信息
	ctx := context.WithValue(context.Background(),tracer.LogTraceKey,"46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1")
	Infof("hello %s", "world", ctx)
	Infof("hello %s", "world")
	Infof("hello %s,%d", "world", 2018, context.Background())
	Errorf("hello %s,%d", "world", 2018)
	l2 := logrus.New(conf.WithLogPath("tmp"), conf.WithLogName("logrus"), conf.WithProjectName("logrus test"))
	SetLogger(l2)
	Debugf("this is logrus %s", "test", context.Background())
	time.Sleep(time.Second * 5)
}
