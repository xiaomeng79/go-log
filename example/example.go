package example

import (
	"github.com/xiaomeng79/go-log/zaplog"
	"context"
	"errors"
	"github.com/xiaomeng79/go-utils/curl"
	"github.com/xiaomeng79/go-log"
)

//初始化日志,可以再这里初始化不同日志引擎的日志 、、 zap logrous

//zap

func LogInit() {
	//设置日志里面记录的项目名称
	log.SetProjectName("test")
	//初始化zap
	l:= zaplog.New(
		zaplog.WithLogPath("tmp/log/"),
		zaplog.WithLogName("test"),
		zaplog.WithMaxAge(7),
		zaplog.WithMaxSize(100),
		zaplog.WithIsStdOut(true),
	)
	//设置日志引擎为刚初始化的
	log.SetLogger(l)
}

func ZapTest() {
	//设置项目名称
	LogInit()
	//普通info日志
	log.Info(context.Background(),"test")
	//错误日志，打印错误栈信息
	log.Error(context.Background(),"inside error")
	//模拟新建一个错误日志类型,打印警告信息
	err := errors.New("this is a test error")
	log.WarnO(context.Background(),err,"inside error")
	//模拟一个http请求,打印请求的一些信息,包括请求头，请求体，延迟，请求类型，方法，响应等
	cb := &curl.CurlBuilder{}
	c :=cb.SetMethod("GET").SetUrl("https://www.baidu.com/").SetHeader("Content-Type","application/json").Build()
	//执行请求
	c.Do()
	log.InfoO(context.Background(),c,"请求百度")
}


//logrus引擎
//var LogrousLog = logrus.New(
//	logrus.WithLogPath("tmp/log/"),
//	logrus.WithLogName("test111"),
//	logrus.WithIsStdOut(false),
//)



