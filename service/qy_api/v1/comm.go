package v1

import "context"

var (
	Ctx         = context.WithValue(context.TODO(), "token", "")
	QyApiProfix = "https://qyapi.weixin.qq.com"
)
