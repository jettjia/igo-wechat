package v1

import (
	"encoding/json"
	"github.com/jettjia/wechat/util"
	"strconv"
)

type GetPubTemplateTitleListRsp struct {
	MiniRspError
	Count int            `json:"count"`
	Data  []TempDataInfo `json:"data"`
}

type TempDataInfo struct {
	Tid        int    `json:"tid"`        // 模版标题 id
	Title      string `json:"title"`      // 模版标题
	Type       int    `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
	CategoryId int    `json:"categoryId"` // 模版所属类目 id
}

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
// ids 类目 id，多个用逗号隔开
// start 用于分页，表示从 start 开始。从 0 开始计数
// limit 用于分页，用于分页，表示拉取 limit 条记录。最大为 30
func (c *Client) GetPubTemplateTitleList(ids string, start int, limit int) GetPubTemplateTitleListRsp {
	apiUrl := MiniApiProfix + "/wxaapi/newtmpl/getpubtemplatetitles?access_token=" + c.AuthToken + "&ids=" + ids + "&start=" + strconv.Itoa(start) + "&limit=" + strconv.Itoa(limit)
	str, err := util.HttpClient(Ctx, apiUrl, "GET", nil, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp GetPubTemplateTitleListRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
