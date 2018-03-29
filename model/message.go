package model

// 微信服务器post过来的加密消息结构体
type CipherMessage struct {
	ToUserName string `xml:"ToUserName"`
	CipherText string `xml:"Encrypt"`
}

// 解密后的消息格式 |4 bytes 消息长度|消息内容|公众号appid|
type PlainMessage struct {
	Lenght    int
	PlainText []byte
	AppID     string
}

// 公共消息
type CommonMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	Type         string `xml:"MsgType"`
}

// 订阅事件消息
//<xml>
//	<ToUserName><![CDATA[asdfadsfadsf]]></ToUserName>
//	<FromUserName><![CDATA[sdfasdfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164501</CreateTime>
//	<MsgType><![CDATA[event]]></MsgType>
//	<Event><![CDATA[subscribe]]></Event>
//	<EventKey><![CDATA[qrscene_123456]]></EventKey>
//	<Ticket><![CDATA[sdfadfasdfsafdsdafsdfasdfasdfasdf]]></Ticket>
//</xml>
// 退订事件消息
//<xml>
//	<ToUserName><![CDATA[sdfsdfsdf]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164439</CreateTime>
//	<MsgType><![CDATA[event]]></MsgType>
//	<Event><![CDATA[unsubscribe]]></Event>
//	<EventKey><![CDATA[]]></EventKey>
//</xml>
// 扫描事件消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522163356</CreateTime>
//	<MsgType><![CDATA[event]]></MsgType>
//	<Event><![CDATA[SCAN]]></Event>
//	<EventKey><![CDATA[123456]]></EventKey>
//	<Ticket><![CDATA[adfadsfasdfasdfadsfasdfasdfasdfadsf]]></Ticket>
//</xml>
// 从报文来看，scene_id和scene_str类型的参数一样，在报文中都是字符串类型
type EventMessage struct {
	*CommonMessage
	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket,omitempty"`
}

// 用户上报位置信息
// 开启用户上报地理位置，且用户同意，用户扫描二维码或进入公众号会话时会上报
// 如果用户不同意，则不会上报
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522225667</CreateTime>
//	<MsgType><![CDATA[event]]></MsgType>
//	<Event><![CDATA[LOCATION]]></Event>
//	<Latitude>34567.214600</Latitude>
//	<Longitude>12345.389801</Longitude>
//	<Precision>20.000000</Precision>
//</xml>
type LocationEventMessage struct {
	*CommonMessage
	Event     string  `xml:"Event"`
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision int     `xml:"Precision"`
}

// 文本消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522163762</CreateTime>
//	<MsgType><![CDATA[text]]></MsgType>
//	<Content><![CDATA[哦了]]></Content>
//	<MsgId>6537643577378888972</MsgId>
//</xml>
// 表情消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522163859</CreateTime>
//	<MsgType><![CDATA[text]]></MsgType>
//	<Content><![CDATA[/::~]]></Content>
//	<MsgId>6537643993990716696</MsgId>
//</xml>
// 自定义图片表情发送后报文格式是文本消息格式，如下：
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522166704</CreateTime>
//	<MsgType><![CDATA[text]]></MsgType>
//	<Content><![CDATA[【收到不支持的消息类型，暂无法显示】]]></Content>
//	<MsgId>6537656213172674037</MsgId>
//</xml>
type TextMessage struct {
	*CommonMessage
	Content string `xml:"Content"`
	Id      int64  `xml:"MsgId,omitempty"`
}

// 声音消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522163960</CreateTime>
//	<MsgType><![CDATA[voice]]></MsgType>
//	<MediaId><![CDATA[sadfasdfasdfasdfasdfadsfasdfadsfadsfadsf]]></MediaId>
//	<Format><![CDATA[amr]]></Format>
//	<MsgId>6537644427782413596</MsgId>
//	<Recognition><![CDATA[]]></Recognition>
//</xml>
type VoiceMessage struct {
	*CommonMessage
	MediaId     string `xml:"MediaId"`
	Format      string `xml:"Format"`
	Id          int64  `xml:"MsgId,omitempty"`
	Recognition string `xml:"Recognition"`
}

// 链接消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164050</CreateTime>
//	<MsgType><![CDATA[link]]></MsgType>
//	<Title><![CDATA[今天我们来聊一个很高级的话题：如何设计一个大规模远程命令执行系统]]></Title>
//	<Description><![CDATA[通过构建CCS系统，我们解决了命令在大量服务器上规模执行的问题。]]></Description>
//	<Url><![CDATA[http://mp.weixin.qq.com/s?__biz=MzA4Nzg5Nzc5OA==&mid=2651670207&idx=1&sn=5eef25a77c28aff2e9124cd0b5a1cc52&chksm=8bcb8516bcbc0c0058dd589fe4aac37581a6a20ba5b1703a4ab9b3caf7711f1d87f6a15446af&scene=0#rd]]></Url>
//	<MsgId>6537644814329470245</MsgId>
//</xml>
type LinkMessage struct {
	*CommonMessage
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	Id          int64  `xml:"MsgId,omitempty"`
	Url         string `xml:"Url"`
}

// 图片消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164143</CreateTime>
//	<MsgType><![CDATA[image]]></MsgType>
//	<PicUrl><![CDATA[http://mmbiz.qpic.cn/mmbiz_jpg/asdfasdfasdfasdfasdfasdfasdfasdf/0]]></PicUrl>
//	<MsgId>6537645213761428779</MsgId>
//	<MediaId><![CDATA[asdfasdfasdfasdfadsfadsfasdfadsfasdf]]></MediaId>
//</xml>
type ImageMessage struct {
	*CommonMessage
	MediaId string `xml:"MediaId"`
	Id      int64  `xml:"MsgId,omitempty"`
	Url     string `xml:"PicUrl"`
}

// 文件消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164211</CreateTime>
//	<MsgType><![CDATA[file]]></MsgType>
//	<Title><![CDATA[asdfasdfasdfasdfasdf.pdf]]></Title>
//	<Description><![CDATA[225.5 KB]]></Description>
//	<FileKey><![CDATA[asdfasdfasdfasdfasdfasdfasdfasdf]]></FileKey>
//	<FileMd5><![CDATA[asdfasdfasdfasdfasdfasdfasdf]]></FileMd5>
//	<FileTotalLen>230895</FileTotalLen>
//	<MsgId>6537645505819204914</MsgId>
//</xml>
type FileMessage struct {
	*CommonMessage
	Title           string `xml:"Title"`
	Id              int64  `xml:"MsgId,omitempty"`
	Description     string `xml:"Description"`
	FileKey         string `xml:"FileKey"`
	FileMd5         string `xml:"FileMd5"`
	FileTotalLenght int64  `xml:"FileTotalLen"`
}

// 视频消息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522164336</CreateTime>
//	<MsgType><![CDATA[video]]></MsgType>
//	<MediaId><![CDATA[asdfasdfasdfadsfadsfasdfasdfasdfasdfasdfasdf]]></MediaId>
//	<ThumbMediaId><![CDATA[asdfasdfasdfasdfasdfasdfasdfasdfasdfasdf]]></ThumbMediaId>
//	<MsgId>6537646042690116922</MsgId>
//</xml>
type VideoMessage struct {
	*CommonMessage
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	Id           int64  `xml:"MsgId,omitempty"`
}

// 位置信息
//<xml>
//	<ToUserName><![CDATA[qweqweqwrqwe]]></ToUserName>
//	<FromUserName><![CDATA[adsfasdfasdfasdf]]></FromUserName>
//	<CreateTime>1522223489</CreateTime>
//	<MsgType><![CDATA[location]]></MsgType>
//	<Location_X>123.45678</Location_X>
//	<Location_Y>345.67890</Location_Y>
//	<Scale>160</Scale>
//	<Label><![CDATA[asdfasdfasdfadsfasdf(asdfasdfasdf)]]></Label>
//	<MsgId>6537900102890582647</MsgId>
//</xml>
type LocationMessage struct {
	*CommonMessage
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
	Id        int64   `xml:"MsgId"`
}
