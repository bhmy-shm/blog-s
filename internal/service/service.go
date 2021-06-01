package service

import (
	"context"
	"goweb/global"
	"goweb/internal/Check"
	"goweb/internal/dao"
)

type Service struct{
	ctx context.Context
	dao *dao.Dao
	check *Check.Check
}

func New(ctx context.Context) Service{
	svc := Service{ctx:ctx}
	svc.dao = dao.New(global.DBEngine,global.RedisEngin)
	svc.check = Check.New(global.DBEngine,global.RedisEngin)
	return svc
}

