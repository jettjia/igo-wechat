package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type GetPaidUnionidRsp struct {
	MiniRspError
	Unionid string `json:"unionid"`
}

// GetPaidUnionid 支付后获取 Unionid
// openid 用户的openid
// transactionId 微信支付订单号
func (c *Client) GetPaidUnionid(openid string, transactionId string) GetPaidUnionidRsp {
	apiUrl := MiniApiProfix + "/wxa/getpaidunionid?access_token=" + c.AuthToken + "&openid=" + openid + "&transaction_id=" + transactionId
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp GetPaidUnionidRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
