```text
video\video\rpc

go run video.go -f etc/video.yaml

goctl rpc protoc video.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```