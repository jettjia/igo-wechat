package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type Code2SessionRsp struct {
	MiniRspError
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

// Code2Session 小程序登录
// jsCode 登录时获取的 code，可通过wx.login获取
func (c *Client) Code2Session(jsCode string) Code2SessionRsp {
	apiUrl := MiniApiProfix + "/sns/jscode2session?appid=" + c.AppID + "&secret=" + c.AppSecret + "&js_code=" + jsCode + "&grant_type=authorization_code"
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp Code2SessionRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
