package v1

import (
	"github.com/jettjia/wechat/util"
)

type GetUnlimitedQRCodeRsp struct {
	MiniRspError
	//ContentType string `json:"contentType"`
	Buffer []byte `json:"buffer"`
}

// GetUnlimitedQRCode 获取小程序二维码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func (c *Client) GetUnlimitedQRCode(scene string, page string, envVersion string) string {
	apiUrl := MiniApiProfix + "/wxa/getwxacodeunlimit?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["scene"] = scene
	reqMap["page"] = page
	reqMap["env_version"] = envVersion
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	return str
}
