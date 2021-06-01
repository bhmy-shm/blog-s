package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"goweb/pkg/MyLogger"
	"goweb/pkg/Setting"
	"sync"
)

var (
	//Viper 配置项
	ServerSetting *Setting.ServerSettingS
	AppSetting	*Setting.AppSettingS
	DataBaseSetting *Setting.DatabaseSettingS
	RabbitMQSetting *Setting.RabbitMQS
	JWTSetting *Setting.JWTSettingS
	RedisSetting *Setting.RedisSettingS
	EmailSetting *Setting.EmailSettings

	//日志记录
	Logger *MyLogger.Logger

	//第三方连接
	DBEngine *gorm.DB
	//RQEngine *amqp.Connection
	RedisEngin *redis.Client

	//全局GMap
	GMap sync.Map

	//日志追踪
	Tracer opentracing.Tracer
)