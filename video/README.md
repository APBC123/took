```text
启动rpc服务前需要启动etcd

cd video/video/rpc

go run video.go -f etc/video.yaml

goctl rpc protoc video.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```
储存桶使用腾讯云COS

腾讯云COS后台地址：https://console.cloud.tencent.com/cos/bucket

腾讯云COS帮助文档：https://cloud.tencent.com/document/product/436/31215
