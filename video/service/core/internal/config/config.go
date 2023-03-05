package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	VideoRpc zrpc.RpcClientConf
	Mysql    struct {
		DataSource string
	}
}
