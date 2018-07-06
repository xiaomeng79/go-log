package logrus

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xiaomeng79/go-utils/curl"
	"fmt"
	"github.com/xiaomeng79/go-log"
)

//公共的列
func getCommonFileds(ctx context.Context,llog *Log) (*logrus.Entry){
	return llog.logger.WithFields(logrus.Fields{
		"project": log.GetProjectName(ctx),
		"trace_id":    log.GetTraceId(ctx),
		"parent_id":    log.GetParentId(ctx),
	})
}

//字符串的列
func getStringFileds(ctx context.Context,llog *Log) (*logrus.Entry) {
	entry := getCommonFileds(ctx,llog)
	entry.Data["logType"] = "inside"
	return entry
}

//字符串--start
func (l *Log)Debug(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Debug(s)
}
func (l *Log)Info(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Info(s)
}
func (l *Log)Warn(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Warn(s)
}
func (l *Log)Error(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Error(s)
}
func (l *Log)Panic(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Panic(s)
}
func (l *Log)Fatal(ctx context.Context,s string) {
	entry := getStringFileds(ctx,l)
	entry.Fatal(s)
}

//判断其他类型--start
func getOtherFileds(ctx context.Context,unknow interface{},llog *Log) (*logrus.Entry) {
	entry := getCommonFileds(ctx,llog)
	switch t := unknow.(type) {
	case *curl.Curl:
			entry.Data["logType"] = "http"
			entry.Data["method"] = t.Method
			entry.Data["url"] = t.Url
			entry.Data["statusCode"] = t.StatusCode
			entry.Data["delay"] = t.Delay
			entry.Data["request"] = fmt.Sprintf("%+v",t.Request)
			entry.Data["response"] = fmt.Sprintf("%+v",t.Response)
	case error:
		entry.Data["logType"] = "system"
		entry.Data["error"] = fmt.Sprintf("%+v",unknow)
	default:
		entry.Data["logType"] = "unknow"
		entry.Data["unknow"] = fmt.Sprintf("%+v",unknow)
	}
	return entry
}

func (l *Log)DebugO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Debug(s)
}
func (l *Log)InfoO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Info(s)
}
func (l *Log)WarnO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Warn(s)
}
func (l *Log)ErrorO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Error(s)
}
func (l *Log)PanicO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Panic(s)
}
func (l *Log)FatalO(ctx context.Context,other interface{},s string) {
	entry := getOtherFileds(ctx,other,l)
	entry.Fatal(s)
}

//判断其他类型--end
