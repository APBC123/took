```text
启动rpc服务前需要启动etcd

go run video.go -f etc/video.yaml

goctl rpc protoc video.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```