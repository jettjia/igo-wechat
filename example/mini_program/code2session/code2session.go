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
	rsp := client.Code2Session("071ovw000ZNbiO19hu3005steJ2ovw0c")

	fmt.Println("auth: ", util.PrintJson(rsp))
}
