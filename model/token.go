package model

// 获取token接口http响应数据结构体
type AccessTokenResponeData struct {
	ErrorCode    int    `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
	Token        string `json:"access_token,omitempty"`
	Expires      int    `json:"expires_in,omitempty"`
}
