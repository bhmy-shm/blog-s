package Check

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"goweb/internal/dao"
	"goweb/internal/model/UserModel"
	"sync"
)

type Check struct{
	db *gorm.DB
	redis *redis.Client
	dao *dao.Dao
	users *UserModel.User
	lock  sync.Mutex
}

func New(engine *gorm.DB,redis *redis.Client) *Check {
	check := Check{db: engine,redis:redis}
	check.dao = dao.New(engine,redis)
	return &check
}

