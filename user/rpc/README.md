生成UserRPC服务代码：
```
cd user/rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

启动userRPC：

```
cd user/rpc
go run user.go -f etc/user.yaml
```

> 注意先开启 etcd、mysql