package log

import "context"

//使用string是为了减少使用Spintf
type ILog interface {
	//普通日志
	Debug(context.Context,string)
	Info(context.Context,string)
	Warn(context.Context,string)
	Error(context.Context,string)
	Panic(context.Context,string)
	Fatal(context.Context,string)
	//其他日志 如：HTTP RPC日志
	DebugO(context.Context,interface{},string)
	InfoO(context.Context,interface{},string)
	WarnO(context.Context,interface{},string)
	ErrorO(context.Context,interface{},string)
	PanicO(context.Context,interface{},string)
	FatalO(context.Context,interface{},string)
}
