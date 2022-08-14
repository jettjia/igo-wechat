package v1

import (
	"github.com/jettjia/wechat/sdk/qy_api"
)

type Client struct {
	qy_api.Config
	AuthToken string
}

func NewQyClient(conf qy_api.Config) *Client {
	return &Client{Config: conf}
}

func NewQyClientWithAuth(conf qy_api.Config, authToken string) *Client {
	return &Client{Config: conf, AuthToken: authToken}
}
