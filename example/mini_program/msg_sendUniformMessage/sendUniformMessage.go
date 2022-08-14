package main

import (
	"fmt"

	"github.com/jettjia/wechat/example/consts"
	"github.com/jettjia/wechat/sdk/mini_program"
	v1 "github.com/jettjia/wechat/service/mini_program/v1"
	"github.com/jettjia/wechat/util"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/uniform-message/sendUniformMessage.html
func main() {
	conf := mini_program.Config{
		AppID:      consts.CAppID,
		AppSecret:  consts.CAppSecret,
		IsDebugLog: true,
	}
	client := v1.NewMiniClientWithAuth(conf, consts.CAuthToken)
	miniMsg := v1.WeappTemplateMsg{
		TemplateId: "", // 小程序模板ID
		Page:       "page/page/index",
		FormId:     "", //小程序模板消息formid
		Data:       "",
	}
	mpMsg := v1.MpTemplateMsg{}
	touser := "odgb84qgDlZldSppKekCu6olrJMM" // 接收者的 openid
	rsp := client.SendUniformMessage(touser, miniMsg, mpMsg)

	fmt.Println("auth: ", util.PrintJson(rsp))
}
