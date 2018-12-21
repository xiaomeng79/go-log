package tracer

import (
	"context"
	"errors"
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"strings"
)

//定义默认的
const (
	Default_Trace_Id  = "system"
	Default_Span_Id   = "0"
	Default_Parent_Id = "0"
	Log_Trace = "log_trace" //格式 traceid:spanid:parentid  46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1
)

//定义错误
var NoTracerInfo = errors.New("no trace info")

//获取traceId,如果没有就是系统内部日志
func GetTraceId(ctx context.Context) string {
	s, err := decodeTracer(ctx)
	if err != nil {
		return Default_Trace_Id
	}
	return s[0]

}

//获取ParentId
func GetParentId(ctx context.Context) string {
	s, err := decodeTracer(ctx)
	if err != nil {
		return Default_Parent_Id
	}
	return s[2]
}

//获取SpanId
func GetSpanId(ctx context.Context) string {
	s, err := decodeTracer(ctx)
	if err != nil {
		return Default_Span_Id
	}
	return s[1]
}

//解析trace中的信息
func decodeTracer(ctx context.Context) ([]string, error) {
	s := make([]string,4,4)
	if val,ok :=ctx.Value(Log_Trace).(string);ok {
		s = strings.Split(fmt.Sprintf("%v", val), ":")
	} else {
		span := opentracing.SpanFromContext(ctx)
		s = strings.Split(fmt.Sprintf("%v", span), ":")
	}
	if len(s) >= 3 {
		return s, nil
	}
	return []string{}, NoTracerInfo
}
