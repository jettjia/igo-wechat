package v1

import (
	"encoding/json"
	"github.com/jettjia/wechat/util"
)

type GetCategoryRsp struct {
	MiniRspError
	Data []DataInfo `json:"data"` // 类目列表
}

type DataInfo struct {
	Id   int    `json:"id"`   // 类目id，查询公共库模版时需要
	Name string `json:"name"` // 类目的中文名
}

// GetCategory 获取小程序账号的类目
func (c *Client) GetCategory() GetCategoryRsp {
	apiUrl := MiniApiProfix + "/wxaapi/newtmpl/getcategory?access_token=" + c.AuthToken
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp GetCategoryRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
