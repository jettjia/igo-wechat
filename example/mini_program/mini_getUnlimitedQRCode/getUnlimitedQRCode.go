package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/jettjia/wechat/example/consts"
	"github.com/jettjia/wechat/sdk/mini_program"
	v1 "github.com/jettjia/wechat/service/mini_program/v1"
	"path"
)

// 获取不限制的小程序个人推广码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func main() {
	conf := mini_program.Config{
		AppID:      consts.CAppID,
		AppSecret:  consts.CAppSecret,
		IsDebugLog: true,
	}
	client := v1.NewMiniClientWithAuth(conf, consts.CAuthToken)
	scene := "g_user_id=100"    // 传递给小程序的参数
	page := "pages/index/index" // 接收者的 openid
	envVersion := "develop"     // 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版
	buffer := client.GetUnlimitedQRCode(scene, page, envVersion)

	//fmt.Println("data: ", util.PrintJson(buffer))

	// 生成图片保存
	filename := "promotion.png"
	filepath := path.Join("./", filename)
	err := gfile.PutContents(filepath, buffer)
	if err != nil {
		fmt.Println("保存推广图片失败")
	}
}
