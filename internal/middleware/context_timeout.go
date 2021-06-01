package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		ctime,cancel := context.WithTimeout(ctx.Request.Context(),t)
		defer cancel()
		//将新定义的超级时间控制写入到gin的context中
		ctx.Request = ctx.Request.WithContext(ctime)
		ctx.Next()
	}
}

