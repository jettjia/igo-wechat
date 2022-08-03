package v1

type MiniRspError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
