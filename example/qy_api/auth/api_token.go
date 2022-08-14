package main

import (
	"fmt"
	"github.com/jettjia/wechat/sdk/qy_api"
	v1 "github.com/jettjia/wechat/service/qy_api/v1"

	"github.com/jettjia/wechat/example/consts"
	"github.com/jettjia/wechat/util"
)

func main() {
	conf := qy_api.Config{
		AppID:      consts.QyAppID,
		AppSecret:  consts.QyAppSecret,
		IsDebugLog: true,
	}
	client := v1.NewQyClient(conf)
	rsp := client.GetAccessToken()

	fmt.Println("rsp: ", util.PrintJson(rsp))
}
