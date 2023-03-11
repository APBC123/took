```text
启动rpc服务前需要启动etcd

go run chat.go -f etc/chat.yaml

goctl rpc protoc chat.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```