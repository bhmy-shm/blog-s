package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RCookie() gin.HandlerFunc{
	return func(context *gin.Context) {

		if cookie,err := context.Cookie("15662166152") ;err == nil{
			fmt.Println("cookie:",cookie)
			context.Next()
		}else{
			fmt.Println("cookie err",err)
		}
		context.Abort()
	}
}