package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"goweb/global"
	"goweb/pkg/MyLogger"
	"log"
	"time"
)

/*
	日志管理中间件
*/

var (
	logger *MyLogger.Logger
)


type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int,error){
	if n,err := w.body.Write(p) ; err != nil{
		return n,err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc{
	return func(context *gin.Context) {
		bodyWrite := &AccessLogWriter{ body:bytes.NewBufferString(""),ResponseWriter:context.Writer}
		context.Writer = bodyWrite
		host := context.Request.Host

		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()

		fields := MyLogger.Fields{
			"request":context.Request.PostForm.Encode(),
			"response":bodyWrite.body.String(),
		}

		//创建正常访问日志
		logger = MyLogger.NewLogger(&lumberjack.Logger{
			Filename: global.AppSetting.AccessLogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,  //设置日志文件允许的最大占用空间 600MB
			MaxAge:    10,   //日志文件的最大声明周期 10天
			LocalTime: true, //日志文件名的时间格式为本地时间
		},"",log.LstdFlags).WithCaller(2)	//找两层栈调用进行记录


		logger.WithFields(fields).InfoF(context,"access log: method: %s,status_code: %d,host: %s,begin_time: %d,end_time: %d",
			context.Request.Method,
			bodyWrite.Status(),
			host,
			beginTime,
			endTime,
		)
	}
}