package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type CheckEncryptedDataRsp struct {
	MiniRspError
	Vaild      bool `json:"vaild"`
	CreateTime int  `json:"create_time"`
}

// CheckEncryptedData 检查加密信息
// encryptedMsgHash 加密数据的sha256
func (c *Client) CheckEncryptedData(encryptedMsgHash string) CheckEncryptedDataRsp {
	apiUrl := MiniApiProfix + "/wxa/business/checkencryptedmsg?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["encrypted_msg_hash"] = encryptedMsgHash
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp CheckEncryptedDataRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
