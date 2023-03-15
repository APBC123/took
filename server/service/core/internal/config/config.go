package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	VideoRpc zrpc.RpcClientConf
	UserRpc  zrpc.RpcClientConf
	ChatRpc  zrpc.RpcClientConf
	Mysql    struct {
		DataSource string
	}
	CacheRedis struct {
		Addr string
	}
	JwtAuth struct {
		SecretKey string
		Duration  int64
	}
}
