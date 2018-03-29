package model

// 二维码创建接口请求数据结构体
// {
//		"expire_seconds": 604800,
//		"action_name": "QR_SCENE",
//		"action_info": {
//			"scene": {
//				"scene_id": 123
//			}
//		}
// }
type QRcodeRequestData struct {
	Expires    int32              `json:"expire_seconds,omitempty"`
	ActionName string             `json:"action_name"`
	ActionInfo *ActionInformation `json:"action_info"`
}

//
type ActionInformation struct {
	SceneInfo *SceneInformation `json:"scene"`
}

//
type SceneInformation struct {
	Id  interface{} `json:"scene_id,omitempty"`
	Str interface{} `json:"scene_str,omitempty"`
}

// 二维码创建接口http响应数据结构体
// {
//		"ticket":"asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf",
//		"expire_seconds":60,
//		"url":"http://weixin.qq.com/q/asdfasdfadsfasdfadsfasdfasdf"
// }
// 错误返回如下
// {"errcode":40066,"errmsg":"invalid url hint: [k7PaiA0848vr49!]"}
type QRcodeResponeData struct {
	ErrorCode    int    `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
	Ticket       string `json:"ticket,omitempty"`
	Expires      int    `json:"expire_seconds,omitempty"`
	Url          string `json:"url,omitempty"`
}
