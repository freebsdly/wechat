package model

// 用户信息
type UserInformation struct {
	ErrorCode      int    `json:"errcode,omitempty"`
	ErrorMessage   string `json:"errmsg,omitempty"`
	Subscribe      int    `json:"subscribe"`
	OpenId         string `json:"openid,omitempty"`
	NickName       string `json:"nickname,omitempty"`
	Sex            int    `json:"sex,omitempty"`
	City           string `json:"city,omitempty"`
	Country        string `json:"country,omitempty"`
	Province       string `json:"province,omitempty"`
	Language       string `json:"language,omitempty"`
	HeadImageUrl   string `json:"headimgurl,omitempty"`
	SubscribeTime  int64  `json:"subscribe_time,omitempty"`
	UnionId        string `json:"unionid,omitempty"`
	Remark         string `json:"remark,omitempty"`
	GroupId        int    `json:"groupid,omitempty"`
	TagIdList      []int  `json:"tagid_list,omitempty"`
	SubscribeScene string `json:"subscribe_scene,omitempty"`
	QRScene        int32  `json:"qr_scene,omitempty"`
	QRSceneString  string `json:"qr_scene_str,omitempty"`
}

//
type UserInformationList struct {
	ErrorCode    int                `json:"errcode,omitempty"`
	ErrorMessage string             `json:"errmsg,omitempty"`
	List         []*UserInformation `json:"user_info_list,omitempty"`
}

//
func NewUserInformationList() *UserInformationList {
	return &UserInformationList{
		List: make([]*UserInformation, 0),
	}
}

//
type BatchUserInfoRequestData struct {
	UserList []*UserInfoRequestData `json:"user_list,omitempty"`
}

//
type UserInfoRequestData struct {
	OpenId   string `json:"openid"`
	Language string `json:"lang"`
}

// 生成获取用户信息的请求数据
func NewUserInfoRequestData(openid, lang string) *UserInfoRequestData {
	return &UserInfoRequestData{
		OpenId:   openid,
		Language: lang,
	}
}

//  生成批量获取用户信息的请求数据
func NewBatchUserInfoRequestData(openids []string, lang string) *BatchUserInfoRequestData {
	num := len(openids)
	datalist := make([]*UserInfoRequestData, num)
	for i := 0; i < num; i++ {
		datalist[i] = NewUserInfoRequestData(openids[i], lang)
	}

	return &BatchUserInfoRequestData{
		UserList: datalist,
	}
}
