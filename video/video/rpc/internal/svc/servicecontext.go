package svc

import (
	"github.com/go-redis/redis/v8"
	"took/video/models"
	"took/video/video/rpc/InitRedis"
	"took/video/video/rpc/internal/config"
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
		Engine: models.Init(c.Mysql.DataSource),
		RDB:    InitRedis.InitRedis(c),
	}
}
