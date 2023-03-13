package svc

import (
	"github.com/go-redis/redis/v8"
	"took/user/model"
	"took/user/rpc/InitRedis"
	"took/user/rpc/internal/config"

	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.NewEngine(c.Mysql.DataSource),
		RDB:    InitRedis.InitRedis(c),
	}
}
