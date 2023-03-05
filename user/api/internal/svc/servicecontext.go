package svc

import (
	"took/user/api/internal/config"
	"took/user/model"

	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.NewEngine(c.Mysql.DataSource),
	}
}
