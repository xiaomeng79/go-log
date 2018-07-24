package log

import (
	"context"
	"github.com/xiaomeng79/go-log/zaplog"
	"github.com/xiaomeng79/go-log/tracer"
)

//默认
var l ILog = zaplog.New()
//设置
func SetLogger(l ILog) {
	l = l
}
//设置项目名称
func SetProjectName(s string) {
	tracer.SetProjectName(s)
}

//普通日志
func Debug(ctx context.Context, s string){
	l.Debug(ctx,s)
}
func Info(ctx context.Context, s string){
	l.Info(ctx,s)
}
func Warn(ctx context.Context, s string){
	l.Warn(ctx,s)
}
func Error(ctx context.Context, s string) {
	l.Error(ctx,s)
}
func Panic(ctx context.Context, s string) {
	l.Panic(ctx,s)
}
func Fatal(ctx context.Context, s string) {
	l.Fatal(ctx,s)
}
//其他日志 如：HTTP RPC日志
func DebugO(ctx context.Context,other interface{},s string) {
	l.DebugO(ctx,other,s)
}
func InfoO(ctx context.Context,other interface{},s string) {
	l.InfoO(ctx,other,s)
}
func WarnO(ctx context.Context,other interface{},s string) {
	l.WarnO(ctx,other,s)
}
func ErrorO(ctx context.Context,other interface{},s string) {
	l.ErrorO(ctx,other,s)
}
func PanicO(ctx context.Context,other interface{},s string) {
	l.PanicO(ctx,other,s)
}
func FatalO(ctx context.Context,other interface{},s string) {
	l.FatalO(ctx,other,s)
}
