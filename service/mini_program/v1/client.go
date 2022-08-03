package v1

import "github.com/jettjia/wechat/sdk/mini_program"

type Client struct {
	mini_program.Config
	AuthToken string
}

func NewMiniClient(conf mini_program.Config) *Client {
	return &Client{Config: conf}
}

func NewMiniClientWithAuth(conf mini_program.Config, authToken string) *Client {
	return &Client{Config: conf, AuthToken: authToken}
}
