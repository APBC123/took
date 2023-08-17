package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-queue/kq"
	"took/chat/rpc/InitRedis"
	"took/chat/rpc/internal/config"
	"took/video/models"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config         config.Config
	Engine         *xorm.Engine
	RDB            *redis.Client
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Engine:         models.Init(c.Mysql.DataSource),
		RDB:            InitRedis.InitRedis(c),
		KqPusherClient: kq.NewPusher(c.Kq.Brokers, c.Kq.Topic),
	}
}
