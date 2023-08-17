package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis struct {
		Addr string
	}
	Kq struct {
		service.ServiceConf
		Brokers    []string
		Group      string
		Topic      string
		Offset     string `json:",options=first|last,default=last"`
		Conns      int    `json:",default=1"`
		Consumers  int    `json:",default=8"`
		Processors int    `json:",default=8"`
		MinBytes   int    `json:",default=10240"`    // 10K
		MaxBytes   int    `json:",default=10485760"` // 10M
		Username   string `json:",optional"`
		Password   string `json:",optional"`
	}
}
