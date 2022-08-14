package v1

type QyRspError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
