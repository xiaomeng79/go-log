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
	DefaultTraceIdName  = "trace_id"
	DefaultSpanIdName   = "span_id"
	DefaultParentIdName = "parent_id"
	// 默认分隔符
	DefaultSeparator = ":"
)

// 定义日志中Trace的key
var LogTraceKey struct{} //格式 traceid:spanid:parentid  46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1

//定义错误
var NoTracerInfo = errors.New("no trace info")


// 解析trace中的信息
func decodeTracer(ctx context.Context) ([]string, error) {
	s := make([]string,0,4)
	if val,ok :=ctx.Value(LogTraceKey).(string);ok {
		s = strings.Split(val, DefaultSeparator)
	} else {
		span := opentracing.SpanFromContext(ctx)
		s = strings.Split(fmt.Sprintf("%v", span), DefaultSeparator)
	}
	if len(s) >= 3 {
		return s, nil
	}
	return []string{}, NoTracerInfo
}

// 定义日志中链路跟踪的信息
func GetTraceInfo(ctx context.Context) map[string]string {
	s,err := decodeTracer(ctx)
	trace := make(map[string]string)
	if err != nil {
		return trace
	}
	trace[DefaultTraceIdName] = s[0]
	trace[DefaultSpanIdName] = s[1]
	trace[DefaultParentIdName] = s[2]
	return trace
}
