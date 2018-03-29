package wechat

import (
	"fmt"
)

// 定义wechat的api地址
const (
	WeChatBaseAPIUrl   = "https://api.weixin.qq.com/cgi-bin"
	WeChatTokenAPIUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	WeChatQRcodeAPIUrl = "https://api.weixin.qq.com/cgi-bin/qrcode"
	WeChatUserAPIUrl   = "https://api.weixin.qq.com/cgi-bin/user"
)

var (
	WeChatQRcodeCreateAPIUrl  = fmt.Sprintf("%s/create", WeChatQRcodeAPIUrl)
	WeChatUserInfoAPIUrl      = fmt.Sprintf("%s/info", WeChatUserAPIUrl)
	WechatUserInfoBatchAPIUrl = fmt.Sprintf("%s/batchget", WeChatUserInfoAPIUrl)
)

// 定义语言
const (
	LanguageZHCN    = "zh_CN"
	LanguageZHTW    = "zh_TW"
	LanguageEnglish = "en"
)

// 定义加密通信下，解密后的明文消息长度字段的长度
const (
	MessageLenghtByteSize = 4
)

// 定义二维码类型即相关参数
const (
	QRSceneIntType         = "QR_SCENE"
	QRSceneStringType      = "QR_STR_SCENE"
	QRLimitSceneIntType    = "QR_LIMIT_SCENE"
	QRLimitSceneStringType = "QR_LIMIT_STR_SCENE"

	QRMaxExpires           = 2592000
	QRMaxSceneIdLenght     = 32
	QRMaxSceneLimitId      = 1000000
	QRMaxSceneStringLenght = 64
)

// 定义消息类型
const (
	MessageEventType = "event"
	MessageTextType  = "text"
	MessageVoiceType = "voice"
	MessageVideoType = "video"
	MessageLinkType  = "link"
	MessageImageType = "image"
	MessageFileType  = "file"

	MessageEncryptType = "aes"
)

// 定义事件类型
const (
	EventScanType     = "SCAN"
	EventLocationType = "LOCATION"
)
