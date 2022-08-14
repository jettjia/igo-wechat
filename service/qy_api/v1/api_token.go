package v1

import (
	"encoding/json"
	"github.com/jettjia/wechat/util"
)

type AuthRsp struct {
	QyRspError
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 小程序token
// https://developer.work.weixin.qq.com/document/path/91039
func (c *Client) GetAccessToken() AuthRsp {
	apiUrl := QyApiProfix + "/cgi-bin/gettoken?corpid=" + c.AppID + "&corpsecret=" + c.AppSecret
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp AuthRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.QyRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.QyRspError)
	}

	return rsp
}
