package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

type ClearQuotaRsp struct {
	MiniRspError
}

// ClearQuota 重置API调用次数
func (c *Client) ClearQuota() ClearQuotaRsp {
	apiUrl := MiniApiProfix + "/cgi-bin/clear_quota?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["appid"] = c.AppID
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp ClearQuotaRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
