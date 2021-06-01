package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goweb/pkg/MyJWT"
)
//必须是登录后的用户+写入cookie
func MustLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if token := context.GetHeader("token") ; len(token) == 0 {
			context.String(401,"缺少token参数")
			context.Abort()
		}else{
			use := getJWTUser(token)
			//将用户名和token写成cookie
			path := context.Request.URL.String()
			context.SetCookie(use,token,1000,path,"localhost",false,true)
		}
	}
}

//通过jwt解码拿到注册jwt的用户名
func getJWTUser(token string) string {
	//如果有token则到redis中拿到数据
	claims,err := MyJWT.ParseToken(token)
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
	return claims.AppUser
}
