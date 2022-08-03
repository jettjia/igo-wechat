package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type GetuserencryptkeyRsp struct {
	MiniRspError
	KeyInfoList []EncryptKeyInfo `json:"key_info_list"`
}

type EncryptKeyInfo struct {
	EncryptKey string `json:"encrypt_key"`
	Version    int    `json:"version"`
	ExpireIn   int    `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int    `json:"create_time"`
}

// Getuserencryptkey 用于获取用户encryptKey。 会获取用户最近3次的key，每个 key 的存活时间为3600s
// openid 用户的openid
// signature 用 sessionkey 对空字符串签名得到的结果。session_key可通过code2Session接口获得。
// sig_method 签名方法，只支持 hmac_sha256
func (c *Client) Getuserencryptkey(openid string, signature string, sigMethod string) GetuserencryptkeyRsp {
	apiUrl := MiniApiProfix + "/wxa/business/getuserencryptkey?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["openid"] = openid
	reqMap["signature"] = signature
	reqMap["sig_method"] = sigMethod
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp GetuserencryptkeyRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
