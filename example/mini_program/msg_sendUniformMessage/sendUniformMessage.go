package main

import (
	"fmt"

	"github.com/jettjia/wechat/example/consts"
	"github.com/jettjia/wechat/sdk/mini_program"
	v1 "github.com/jettjia/wechat/service/mini_program/v1"
	"github.com/jettjia/wechat/util"
)

func main() {
	conf := mini_program.Config{
		AppID:      consts.CAppID,
		AppSecret:  consts.CAppSecret,
		IsDebugLog: true,
	}
	client := v1.NewMiniClientWithAuth(conf, consts.CAuthToken)
	miniMsg := v1.WeappTemplateMsg{}
	mpMsg := v1.MpTemplateMsg{}
	rsp := client.SendUniformMessage("openId", miniMsg, mpMsg)

	fmt.Println("auth: ", util.PrintJson(rsp))
}
