package v1

import (
	"encoding/json"
	"github.com/jettjia/wechat/util"
)

type SendUniformMessageRsp struct {
	MiniRspError
}

type WeappTemplateMsg struct {
	TemplateId      int    `json:"template_id"`      // 小程序模板ID
	Page            string `json:"page"`             // 小程序页面路径
	FormId          int    `json:"form_id"`          // 小程序模板消息formid
	EmphasisKeyword string `json:"emphasis_keyword"` // 小程序模板放大关键词
	Data            string `json:"data"`             // 小程序模板数据
}

type MpTemplateMsg struct {
	Appid       string `json:"appid"`       // 公众号appid，要求与小程序有绑定且同主体
	TemplateId  string `json:"template_id"` // 公众号模板id
	Url         string `json:"url"`         // 公众号模板消息所要跳转的url
	Miniprogram string `json:"miniprogram"` // 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Data        string `json:"data"`        // 公众号模板消息的数据
}

// SendUniformMessage 下发小程序和公众号统一的服务消息
// touser 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
//

func (c *Client) SendUniformMessage(touser string, miniMsg WeappTemplateMsg, mpMsg MpTemplateMsg) SendUniformMessageRsp {
	apiUrl := MiniApiProfix + "/cgi-bin/message/wxopen/template/uniform_send?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	reqMap["touser"] = touser
	reqMap["weapp_template_msg"] = miniMsg
	reqMap["mp_template_msg"] = mpMsg
	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp SendUniformMessageRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.MiniRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.MiniRspError)
	}

	return rsp
}
