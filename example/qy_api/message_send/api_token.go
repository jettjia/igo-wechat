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
	client := v1.NewQyClientWithAuth(conf, consts.QyAuthToken)
	var data v1.ReqMessageSend
	data.Toparty = "2"       // 接单的部门
	data.Agentid = "1000002" // 应用ID
	data.Content = "您有新的订单，请及时处理。订单号是：xxxxxxxxxxxxxx，订单金额：100，下单人：张三!"
	data.Safe = 1
	rsp := client.MessageSend(data)

	fmt.Println("rsp: ", util.PrintJson(rsp))
}
