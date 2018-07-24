package log

import (
	"context"
	"errors"
	opentracing "github.com/opentracing/opentracing-go"
	"strings"
	"fmt"
)

//定义默认的
const (
	Default_Project_Name = "srv"
	Default_Trace_Id     = "system"
	Default_Span_Id    = "0"
	Default_Parent_Id    = "0"
)

//定义错误
var NoTracerInfo = errors.New("no trace info")

//项目名称
var logProjectName string
//设置项目名称
func SetProjectName(projectName string)  {
	logProjectName = projectName
}

//获取项目名称
func GetProjectName(ctx context.Context) string {
	if len(logProjectName) == 0 {
		return Default_Project_Name
	}
	return logProjectName
}

//获取traceId,如果没有就是系统内部日志
func GetTraceId(ctx context.Context) string {
	s,err := decodeTracer(ctx)
	if err != nil {
		return Default_Trace_Id
	}
	return s[0]
}

//获取ParentId
func GetParentId(ctx context.Context) string {
	s,err := decodeTracer(ctx)
	if err != nil {
		return Default_Parent_Id
	}
	return s[2]
}

//获取SpanId
func GetSpanId(ctx context.Context) string {
	s,err := decodeTracer(ctx)
	if err != nil {
		return Default_Span_Id
	}
	return s[1]
}

//解析trace中的信息
func decodeTracer(ctx context.Context) ([]string,error) {
	span := opentracing.SpanFromContext(ctx)
	s := strings.Split(fmt.Sprintf("%v",span),":")
	if len(s) >=3 {
		return s,nil
	}
	return []string{},NoTracerInfo
}
