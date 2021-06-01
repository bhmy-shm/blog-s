package Limiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct{
	*Limiter
}

func NewLimiter() LimiterIface{
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

//获取对应的限流器的键值对名称
//将路由中?
func (this MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri,"?")
	fmt.Println("index=",index)
	if index == -1 {
		return uri
	}
	fmt.Println("uri=",uri[:index])
	return uri[:index]
}

//获取令牌桶,
func (this MethodLimiter) GetBucket(key string) (*ratelimit.Bucket,bool) {
	bucket,ok := this.limiterBuckets[key]
	return bucket,ok
}

//新增多个令牌桶
func (this MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _,rule := range rules{
		if _,ok := this.limiterBuckets[rule.Key] ; !ok {
			this.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval,rule.Capacity,rule.Quantum)
		}
	}
	return this
}

