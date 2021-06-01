package MyLogger

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

/* --------------------------------- 日志格式化 ------------------------------------- */

//将所有要输出的日志信息按照json格式化组合在一起
func (l *Logger) JSONFormat(level Level,message string) map[string]interface{}{
	data := make(Fields,len(l.fields)+4)

	data["level"] = level.String()
	data["time"] = time.Now().Local().String()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.fields) > 0 {
		for k,v := range l.fields {
			if _,ok := data[k] ; !ok {
				data[k] = v
			}
		}
	}
	return data
}
//然后根据不同的日志级别，进行输出
func(l *Logger) Output(level Level,message string) {
	//输出前要将组合的日志信息map,转变成字符串
	body,_ := json.Marshal(l.JSONFormat(level,message))
	content := string(body)

	//然后根据不同的日志级别，选择不同的log.Print / log.Fatal 输出方式
	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}


/* -------------------------------- 日志分级别输出 ------------------------------------- */

//根据先前定义的日志分级，编写对应的日志输出的外部方法，继续写入如下代码：
func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug,fmt.Sprint(v...))
}

func (l *Logger) DebugF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) InfoF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) WarnF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) ErrorF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) FatalF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) PanicF(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}
