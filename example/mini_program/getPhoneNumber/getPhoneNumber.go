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
	rsp := client.GetPhoneNumber("5683b96a02a4f29997601d7362d5a4ee7cdc2cacbe083490b4eb9aeb36e14ffd")

	fmt.Println("auth: ", util.PrintJson(rsp))
}
