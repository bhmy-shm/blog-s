package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"goweb/global"
	"goweb/pkg/Setting"
	"log"
	"os"
	"time"
)

const (
	STATE_OPEN  = 0
	STATE_CLOSE = 1
)

// gorm.Model 定义
type Model struct {
	ID        uint8 `json:"id,omitempty",gorm:"primary_key"`
	CreatedAt uint32 `json:"created_at,omitempty" gorm:"type:int(10);unsigned"`
	UpdatedAt uint32	`json:"updated_at,omitempty" gorm:"type:int(10);unsigned"`
	DeletedAt uint32 `json:"deleted_at,omitempty" gorm:"type:int(10);unsigned"`
	IsDel uint8 `json:"is_del,omitempty" gorm:"type:tinyint(3);unsigned"`
}

func NewDBEngine(s *Setting.DatabaseSettingS) (*gorm.DB,error){

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.UserName,s.Password,s.Host,s.DBName,s.Charset,s.ParseTime)

	fmt.Println("mysql，dsn= ",dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)

	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil,err
	}
	if global.ServerSetting.RunMode == "debug" {
	}

	sqlDB,err := db.DB()
	sqlDB.SetMaxIdleConns(s.MaxIdleConns)
	sqlDB.SetMaxOpenConns(s.MaxOpenConns)

	return db,nil
}

