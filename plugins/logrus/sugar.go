package logrus

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xiaomeng79/go-log/tracer"
)


func getCtxFileds(llog *Log, args ...interface{}) (*logrus.Entry){
	//判断是否有context
	l := len(args)
	if l >0 {
		if ctx,ok := args[l-1].(context.Context);ok {
			return llog.logger.WithFields(logrus.Fields{
				"trace_id":    tracer.GetTraceId(ctx),
				"parent_id":    tracer.GetParentId(ctx),
				"span_id":    tracer.GetSpanId(ctx),
				"project": llog.ProjectName,
			})
		}
	}
	return llog.logger.WithField("project",llog.ProjectName)
}


//字符串--start
func (l *Log)Debug(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Debug(s)
}
func (l *Log)Info(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Info(s)
}
func (l *Log)Warn(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Warn(s)
}
func (l *Log)Error(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Error(s)
}
func (l *Log)Panic(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Panic(s)
}
func (l *Log)Fatal(s string,args ...interface{}) {
	entry := getCtxFileds(l,args...)
	entry.Fatal(s)
}

//判断其他类型--start
func getOtherFileds(llog *Log, format string, args ...interface{}) (string, *logrus.Entry) {
	l := len(args)
	if l >0 {
		if ctx,ok := args[l-1].(context.Context);ok {
			return fmt.Sprintf(format,args[:l-1]...), llog.logger.WithFields(logrus.Fields{
				"trace_id":    tracer.GetTraceId(ctx),
				"parent_id":    tracer.GetParentId(ctx),
				"span_id":    tracer.GetSpanId(ctx),
				"project": llog.ProjectName,
			})
		}
	}
	return format,llog.logger.WithField("project",llog.ProjectName)
}


func (l *Log)Debugf(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Debug(s)
}
func (l *Log)Infof(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Info(s)
}
func (l *Log)Warnf(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Warn(s)
}
func (l *Log)Errorf(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Error(s)
}
func (l *Log)Panicf(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Panic(s)
}
func (l *Log)Fatalf(format string, args ...interface{}) {
	s,entry := getOtherFileds(l,format,args...)
	entry.Fatal(s)
}

