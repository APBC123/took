package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"took/user/model"
	"took/user/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      *model.UserModel
	FollowModel    *model.FollowModel
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(c.Mysql.DataSource),
		FollowModel:    model.NewFollowModel(c.Mysql.DataSource),
		KqPusherClient: kq.NewPusher(c.Kq.Brokers, c.Kq.Topic),
	}
}
