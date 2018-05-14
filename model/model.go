// model包定义了wechat库中各类方法函数使用的结构体
package model

//
type CommonResponeData struct {
	ErrorCode    int    `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
}
