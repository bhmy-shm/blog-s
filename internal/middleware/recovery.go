package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/pkg/app"
	"goweb/pkg/email"
	"goweb/pkg/errcode"
	"time"
)


func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL:global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.PassWord,
		From: global.EmailSetting.From,
	})

	return func(context *gin.Context) {
		defer func() {
			if err := recover() ; err != nil {
				global.Logger.WithCallersFrames().ErrorF(context,"panic recover err: %v",err)
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d",time.Now().Unix()),
					fmt.Sprintf("错误信息：%v",err),
					)
				if err != nil {
					global.Logger.PanicF(context,"mail.Send err: %v",err)
				}
				app.NewResponse(context).ToErrorResponse(errcode.ServiceError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
