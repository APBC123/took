package svc

import (
	"took/user/model"
	"took/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserModel *model.UserModel
	FollowModel *model.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(c.Mysql.DataSource),
		FollowModel: model.NewFollowModel(c.Mysql.DataSource),
	}
}
