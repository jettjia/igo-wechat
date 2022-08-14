package v1

import (
	"encoding/json"

	"github.com/jettjia/wechat/util"
)

// RspMessageSendRsp 微信返回
type RspMessageSendRsp struct {
	QyRspError
}

// ReqMessageSend 发送的消息
type ReqMessageSend struct {
	Touser  string `json:"touser"`  // 指定接收消息的成员，成员ID列表（多个接收者用‘|’分隔，最多支持1000个）
	Toparty string `json:"toparty"` // 指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。
	Agentid string `json:"agentid"` // 企业应用的id，整型。
	Content string `json:"content"` // 消息内容，最长不超过2048个字节
	Safe    int    `json:"safe"`    // 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
}

// MessageSend 发送应用消息
// https://developer.work.weixin.qq.com/document/path/90236
func (c *Client) MessageSend(data ReqMessageSend) RspMessageSendRsp {
	apiUrl := QyApiProfix + "/cgi-bin/message/send?access_token=" + c.AuthToken
	reqMap := make(map[string]interface{})
	sendData := make(map[string]interface{})
	if data.Touser != "" {
		reqMap["touser"] = data.Touser
	}
	if data.Toparty != "" {
		reqMap["toparty"] = data.Toparty
	}
	reqMap["agentid"] = data.Agentid
	reqMap["msgtype"] = "text"
	sendData["content"] = data.Content
	reqMap["text"] = sendData
	reqMap["safe"] = data.Safe

	str, err := util.HttpClient(Ctx, apiUrl, "POST", reqMap, nil, c.IsDebugLog)
	if err != nil {
		util.NewLogger().Errorln(err)
	}

	var rsp RspMessageSendRsp
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		util.NewLogger().Errorln(err)
	}
	if rsp.QyRspError.Errcode != 0 {
		util.NewLogger().Errorln(rsp.QyRspError)
	}

	return rsp
}
