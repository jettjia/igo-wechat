package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type AuthRsp struct {
	MiniRspError
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 小程序token
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func (c *Client) GetAccessToken() AuthRsp {
	apiUrl := MiniApiProfix + "/cgi-bin/token?grant_type=client_credential&appid=" + c.AppID + "&secret=" + c.AppSecret
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp AuthRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
