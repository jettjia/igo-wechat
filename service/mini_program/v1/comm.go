package v1

import "context"

var (
	Ctx           = context.WithValue(context.TODO(), "token", "")
	MiniApiProfix = "https://api.weixin.qq.com"
)
