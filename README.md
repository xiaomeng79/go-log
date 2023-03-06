# go-log

## 目录

- [go-log](#go-log)
  - [目录](#目录)
  - [简介](#简介)
  - [版本](#版本)
  - [获取](#获取)
  - [快速使用](#快速使用)
    - [使用opentracing定义链路跟踪信息](#使用opentracing定义链路跟踪信息)
    - [兼容istio链路跟踪，通过context包记录跟踪的信息](#兼容istio链路跟踪通过context包记录跟踪的信息)
    - [自定义参数(默认)](#自定义参数默认)
    - [自定义，初始化zap,记录日志(推荐)](#自定义初始化zap记录日志推荐)
    - [快速开始，初始化logrous,记录日志(其他插件自己拓展)](#快速开始初始化logrous记录日志其他插件自己拓展)


[![Build Status](https://travis-ci.org/xiaomeng79/go-log.svg?branch=master)](https://travis-ci.org/xiaomeng79/go-log) [![codecov](https://codecov.io/gh/xiaomeng79/go-log/branch/master/graph/badge.svg)](https://codecov.io/gh/xiaomeng79/go-log)

## 简介
go封装的日志 logrus zap,并且增加了分布式日志追踪，日志格式化成json,日志大小切割

## 版本

v1.0
v2.0

## 获取

```bash
go get -u github.com/xiaomeng79/go-log
```

## 快速使用

### 使用opentracing定义链路跟踪信息

[opentracing-go](https://github.com/opentracing/opentracing-go)

```go
//引入包
import "github.com/xiaomeng79/go-log"

//默认使用zap 插件
log.Info("test")
//输出日志:{"level":"info","@timestamp":"2018-11-01T15:04:47.079+0800","caller":"go-log/log_test.go:21","msg":"test","project":"zap test"}

//使用context包来记录分布式跟踪日志,配合OpenTracing链路跟踪实现代码分析
// 使用链路跟踪
sp := opentracing.StartSpan("operation_name")
defer sp.Finish()
ctx := opentracing.ContextWithSpan(context.Background(), sp)
// 使用ctx来将跟踪信息记录到日志
log.Info("test",ctx)
//输出日志:{"level":"info","@timestamp":"2018-11-01T15:04:47.079+0800","caller":"go-log/log_test.go:21","msg":"hello world","project":"zap test","trace_id":"3ece70e8f602a46d","parent_id":"5e57855e4f15604c","span_id":"791c3d0180bb66ad"}
```

### 兼容istio链路跟踪，通过context包记录跟踪的信息

```go
// 如果需要自定义跟踪信息,格式如下(将istio的请求头中的X-B3-的跟踪信息取出来,拼成以下格式):
// 格式 traceid:spanid:parentid:1  46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1

ctx := context.WithValue(ctx,tracer.LogTraceKey,"46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1")

// 只需要将带有链路信息的context放到日志方法的末尾即可，必须是最后一个参数
// 如：
//不带格式化参数：   log.Info("test",ctx)
// 待格式化参数的：  log.Debugf("this is zap test %s","test",ctx)
```

### 自定义参数(默认)

```go
LogPath string = "/var/log" //日志保存路径
LogName string = "output" //日志保存的名称，不些随机生成
LogLevel string = "debug"  //日志记录级别
MaxSize int = 100 //日志分割的尺寸 MB
MaxAge int = 7 //分割日志保存的时间 day
Stacktrace string = "error" //记录堆栈的级别
IsStdOut string  = "yes"//是否标准输出console输出 yes 输出 no 不输出
ProjectName string = "test" //项目名称
```

### 自定义，初始化zap,记录日志(推荐)

```go
//初始化zap
//引入包
import (
    gLog "github.com/xiaomeng79/go-log"
    "github.com/xiaomeng79/go-log/conf"
    "github.com/xiaomeng79/go-log/plugins/zaplog"
)
//初始化
gLog.SetLogger(zaplog.New(
    conf.WithProjectName("zap test"),
    conf.WithLogPath("tmp"),
    conf.WithLogLevel("info"),
    ))

//使用
gLog.Debugf("this is zap test %s","test",ctx)
```

### 快速开始，初始化logrous,记录日志(其他插件自己拓展)

```go
//初始化zap
//引入包
import (
    gLog "github.com/xiaomeng79/go-log"
    "github.com/xiaomeng79/go-log/conf"
    "github.com/xiaomeng79/go-log/plugins/logrus"
)
//初始化
gLog.SetLogger(logrus.New(
conf.WithProjectName("logrus test"),
// ... 不写执行默认参数
))

//使用
gLog.Debugf("this is logrus test %s","test",ctx)
```
