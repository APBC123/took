package svc

import (
	"took/server/service/core/internal/config"
	"took/video/models"
	"took/video/video/rpc/videoservice"
	"took/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config   config.Config
	Engine   *xorm.Engine
	VideoRpc videoservice.VideoService
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Engine:   models.Init(c.Mysql.DataSource),
		VideoRpc: videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
