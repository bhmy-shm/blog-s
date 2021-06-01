package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao struct{
	db *gorm.DB
	redis *redis.Client
}

func New(engine *gorm.DB,redis *redis.Client) *Dao{
	return &Dao{db: engine,redis: redis}
}

