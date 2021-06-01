package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goweb/pkg/MyJWT"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
)

func JWT() gin.HandlerFunc{
	return func(c *gin.Context){
		var token string
		var err error
		//ecode = errcode.Success

		//从http头中获取token
		if s,exist := c.GetQuery("token") ;exist{
			token = s
		}else {
			token = c.GetHeader("token")
		}

		//判断token是否为空，再判断和redis中拿到的token一样不
		if token == ""{
			err = fmt.Errorf("token 为空")
			//ecode = errcode.InvalidParams
		}else {
			_,err = MyJWT.ParseToken(token)
			if err != nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = fmt.Errorf("token 过期")
					//ecode = errcode.UnauthorizedTokenTimeOut
				default:
					err = fmt.Errorf("token 错误")
					//ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		//if ecode != errcode.Success{
		//	response := app.NewResponse(c)
		//	response.ToErrorResponse(ecode)
		//	c.Abort()
		//	return
		//}
		if err != nil {
			response := app.NewResponse(c)
			response.ToErrorResponse(errcode.UnauthorizedTokenError)
			c.Abort()
			return
		}
		c.Next()
	}
}