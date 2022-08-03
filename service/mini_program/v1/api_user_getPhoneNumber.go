package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type GetPhoneNumberRsp struct {
	MiniRspError
	PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string    `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string    `json:"countryCode"`     // 区号
	Watermark       Watermark `json:"watermark"`       // 数据水印
}

type Watermark struct {
	Timestamp int    `json:"timestamp"` // 用户获取手机号操作的时间戳
	Apid      string `json:"appid"`     // 小程序appid
}

// GetPhoneNumber 将 code 换取用户手机号。 说明，每个 code 只能使用一次，code的有效期为5min
func (c *Client) GetPhoneNumber(code string) GetPhoneNumberRsp {
	apiUrl := MiniApiProfix + "/wxa/business/getuserphonenumber?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["code"] = code
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp GetPhoneNumberRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
