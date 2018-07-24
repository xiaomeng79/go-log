package zaplog

import (
	"context"
	"go.uber.org/zap"
	"github.com/xiaomeng79/go-utils/curl"
	"fmt"
	"github.com/xiaomeng79/go-log/tracer"
)

//公共的列
func getCommonFileds(ctx context.Context) ([]zap.Field){
	return []zap.Field{
		zap.String("project", tracer.GetProjectName(ctx)),
		zap.String("trace_id", tracer.GetTraceId(ctx)),
		zap.String("parent_id", tracer.GetParentId(ctx)),
		zap.String("span_id", tracer.GetSpanId(ctx)),
	}
}
//字符串的列
func getStringFileds(ctx context.Context) ([]zap.Field) {
	cf := getCommonFileds(ctx)
	cf = append(cf,[]zap.Field{
		zap.String("logType","inside"),
	}...)
	return cf
}
//字符串--end

//
func (l *Log)Debug(ctx context.Context,s string) {
	l.logger.Debug(s,getStringFileds(ctx)...)
}
func (l *Log)Info(ctx context.Context,s string) {
	l.logger.Info(s,getStringFileds(ctx)...)
}
func (l *Log)Warn(ctx context.Context,s string) {
	l.logger.Warn(s,getStringFileds(ctx)...)
}
func (l *Log)Error(ctx context.Context,s string) {
	l.logger.Error(s,getStringFileds(ctx)...)
}
func (l *Log)Panic(ctx context.Context,s string) {
	l.logger.Panic(s,getStringFileds(ctx)...)
}
func (l *Log)Fatal(ctx context.Context,s string) {
	l.logger.Fatal(s,getStringFileds(ctx)...)
}
//

//判断其他类型--start
func getOtherFileds(ctx context.Context,unknow interface{}) ([]zap.Field) {
	cf := getCommonFileds(ctx)
	switch t := unknow.(type) {
	case *curl.Curl:
		cf = append(cf,[]zap.Field{
			zap.String("logType","http"),
			zap.String("method",t.Method),
			zap.String("url",t.Url),
			zap.Int("statusCode",t.StatusCode),
			zap.Int64("delay",t.Delay),
			zap.String("request",fmt.Sprintf("%+v",t.Request)),
			zap.String("response",fmt.Sprintf("%+v",t.Response)),
		}...)
	case error:
		cf = append(cf,[]zap.Field{
			zap.String("logType","system"),
			zap.String("error",fmt.Sprintf("%+v",unknow)),
		}...)
	default:
		cf = append(cf,[]zap.Field{
			zap.String("logType","unknow"),
			zap.String("unknow",fmt.Sprintf("%+v",unknow)),
		}...)
	}
	return cf
}

func (l *Log)DebugO(ctx context.Context,other interface{},s string) {
	l.logger.Debug(s,getOtherFileds(ctx,other)...)
}
func (l *Log)InfoO(ctx context.Context,other interface{},s string) {
	l.logger.Info(s,getOtherFileds(ctx,other)...)
}
func (l *Log)WarnO(ctx context.Context,other interface{},s string) {
	l.logger.Warn(s,getOtherFileds(ctx,other)...)
}
func (l *Log)ErrorO(ctx context.Context,other interface{},s string) {
	l.logger.Error(s,getOtherFileds(ctx,other)...)
}
func (l *Log)PanicO(ctx context.Context,other interface{},s string) {
	l.logger.Panic(s,getOtherFileds(ctx,other)...)
}
func (l *Log)FatalO(ctx context.Context,other interface{},s string) {
	l.logger.Fatal(s,getOtherFileds(ctx,other)...)
}

//判断其他类型--end



