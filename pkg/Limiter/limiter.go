package Limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//实现令牌桶的策略
type LimiterIface interface {
	Key(c *gin.Context) string 	//获取对应的限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket,bool)	//获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface //新增多个令牌桶
}

//存储令牌桶与键值对的对应关系，通过map来记录令牌桶
type Limiter struct{
	limiterBuckets map[string]*ratelimit.Bucket
}

//令牌桶的规则属性
type LimiterBucketRule struct{
	Key string	//自定义一个键值对名称
	Capacity int64	//令牌桶的容量
	Quantum int64	//每次到达间隔时间后所放的具体令牌数量
	FillInterval time.Duration	//间隔多久时间放N个令牌
}



