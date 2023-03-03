package define

import "os"

// 腾讯云对象储存
var TencentSecretID = os.Getenv("TencentSecretID")
var TencentSecretKey = os.Getenv("TencentSecretKey")
var CosBucket = "https://2290312980-1316376654.cos.ap-nanjing.myqcloud.com"
