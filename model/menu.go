package model

//
const (
	ClickButtonType           = "click"
	ViewButtonType            = "view"
	MiniProgramButtonType     = "miniprogram"
	LocationSelectButtonType  = "location_select"
	ScanCodePushButtonType    = "scancode_push"
	ScanCodeMessageButtonType = "scancode_waitmsg"
	PicSysPhotoButtonType     = "pic_sysphoto"
	PicPhotoOrAlbumButtonType = "pic_photo_or_album"
	PicWeChatButtonType       = "pic_weixin"
	MediaIdButtonType         = "media_id"
	ViewLimitedButtonType     = "view_limited"
)

//参数	是否必须	说明
//button	是	一级菜单数组，个数应为1~3个
//sub_button	否	二级菜单数组，个数应为1~5个
//type	是	菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
//name	是	菜单标题，不超过16个字节，子菜单不超过60个字节
//key	click等点击类型必须	菜单KEY值，用于消息接口推送，不超过128字节
//url	view、miniprogram类型必须	网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。
//media_id	media_id类型和view_limited类型必须	调用新增永久素材接口返回的合法media_id
//appid	miniprogram类型必须	小程序的appid（仅认证公众号可配置）
//pagepath	miniprogram类型必须	小程序的页面路径

// 菜单
type Menu struct {
	Button []interface{} `json:"button"`
}

func (p *Menu) AddButton(b interface{}) {
	p.Button = append(p.Button, b)
}

//
func NewMenu() *Menu {
	return &Menu{
		Button: make([]interface{}, 0),
	}
}

// 子菜单
type SubButton struct {
	Name    string        `json:"name"`
	Buttons []interface{} `json:"sub_button"`
}

//
func (p *SubButton) AddButton(b interface{}) {
	p.Buttons = append(p.Buttons, b)
}

//
func NewSubButton(name string) *SubButton {
	return &SubButton{
		Name:    name,
		Buttons: make([]interface{}, 0),
	}
}

// 点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），
// 并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；
type ClickButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

//
func NewClickButton(name, key string) *ClickButton {
	return &ClickButton{
		Type: ClickButtonType,
		Name: name,
		Key:  key,
	}
}

// 跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，
// 可与网页授权获取用户基本信息接口结合，获得用户基本信息。
// appid和pagepath是为小程序准备的
type ViewButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

//
func NewViewButton(name, url string) *ViewButton {
	return &ViewButton{
		Name: name,
		Url:  url,
		Type: ViewButtonType,
	}
}

// 跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，
// 可与网页授权获取用户基本信息接口结合，获得用户基本信息。
// appid和pagepath是为小程序准备的
type MiniProgramButton struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}

//
func NewMiniProgramButton(name, url, appid, pp string) *MiniProgramButton {
	return &MiniProgramButton{
		Name:     name,
		Url:      url,
		Type:     MiniProgramButtonType,
		AppID:    appid,
		PagePath: pp,
	}
}

// scancode_push：扫码推事件用户点击按钮后，微信客户端将调起扫一扫工具，
// 完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
// scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框用户点击按钮后，
// 微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，
// 同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
type ScanCodeButton struct {
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Key        string        `json:"key"`
	SubButtons []interface{} `json:"sub_button"`
}

//
func NewScanCodeButton(name, t, key string) *ScanCodeButton {
	return &ScanCodeButton{
		Name:       name,
		Type:       t,
		Key:        key,
		SubButtons: make([]interface{}, 0),
	}
}

// pic_sysphoto：弹出系统拍照发图用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，
// 会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
// pic_photo_or_album：弹出拍照或者相册发图用户点击按钮后，
// 微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
// pic_weixin：弹出微信相册发图器用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，
// 将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
type PictureButton struct {
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Key        string        `json:"key"`
	SubButtons []interface{} `json:"sub_button"`
}

//
func NewPictureButton(name, t, key string) *PictureButton {
	return &PictureButton{
		Name:       name,
		Type:       t,
		Key:        key,
		SubButtons: make([]interface{}, 0),
	}
}

// location_select：弹出地理位置选择器用户点击按钮后，微信客户端将调起地理位置选择工具，
// 完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
type LocationSelectButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

//
func NewLocationSelectButton(name, key string) *LocationSelectButton {
	return &LocationSelectButton{
		Type: LocationSelectButtonType,
		Name: name,
		Key:  key,
	}
}

// media_id：下发消息（除文本消息）用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，
// 永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
type MediaIdButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

func NewMediaIdButton(name, id string) *MediaIdButton {
	return &MediaIdButton{
		Type:    MediaIdButtonType,
		Name:    name,
		MediaID: id,
	}
}

// view_limited：跳转图文消息URL用户点击view_limited类型按钮后，
// 微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。
// 请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
type ViewLimitedButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

func NewViewLimitedButton(name, id string) *ViewLimitedButton {
	return &ViewLimitedButton{
		Type:    ViewLimitedButtonType,
		Name:    name,
		MediaID: id,
	}
}
