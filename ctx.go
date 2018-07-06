package log

import "context"

const (
	Log_Project_Name  = "Project_Name"
	Log_Trace_Id  = "Log-Trace-Id"
	Log_Parent_Id = "Log-Parent-Id"
)

//定义默认的
const (
	Default_Project_Name = "srv"
	Default_Trace_Id = "system"
	Default_Parent_Id = "0"
)

//获取项目名称
func GetProjectName(ctx context.Context) string {
	if v, ok := ctx.Value(Log_Project_Name).(string); ok {
		if len(v) !=0 {
			return v
		}
	}
	return Default_Project_Name

}

//获取traceId,如果没有就是系统内部日志
func GetTraceId(ctx context.Context) string {
	if v, ok := ctx.Value(Log_Trace_Id).(string); ok {
		if len(v) !=0 {
			return v
		}
	}
	return Default_Trace_Id

}
//获取上级Id
func GetParentId(ctx context.Context) string {
	if v, ok := ctx.Value(Log_Parent_Id).(string); ok {
		if len(v) !=0 {
			return v
		}
	}
	return Default_Parent_Id
}