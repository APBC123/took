package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"took/server/service/core/internal/config"
	"took/video/models"
	"took/video/video/rpc/videoservice"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config   config.Config
	Engine   *xorm.Engine
	VideoRpc videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Engine:   models.Init(c.Mysql.DataSource),
		VideoRpc: videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}
