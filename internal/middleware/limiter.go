package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/pkg/Limiter"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
)

func RateLimit(l Limiter.LimiterIface) gin.HandlerFunc{
	return func(context *gin.Context) {
		//1.从上下文中拿到URI并创建key
		key := l.Key(context)
		//2.写入map进行记录
		if bucket,ok := l.GetBucket(key) ; ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(context)
				response.ToErrorResponse(errcode.TooManyRequests)
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
