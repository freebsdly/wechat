package wechat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/freebsdly/wechat/model"
)

// 创建新的二维码请求数据
// 微信官方文档描述：
// 按时间长短分为：临时、永久二维码
// 按二维码参数类型分为: 整型、字符串型参数二维码
// 组合下有4中类型
func NewQRcodeRequestData(aname string, expires int32, scenevalue interface{}) (qrcreqdata *model.QRcodeRequestData, err error) {
	qrcreqdata = &model.QRcodeRequestData{}
	name := strings.ToUpper(aname)

	switch name {
	case QRSceneIntType, QRLimitSceneIntType:
		if name == QRSceneIntType {
			qrcreqdata.Expires = expires
		}

		switch scenevalue.(type) {
		case int, int16, int32:
			break
		default:
			err = fmt.Errorf("scenevalue type is not intger")
			return
		}

		qrcreqdata.ActionInfo = &model.ActionInformation{
			SceneInfo: &model.SceneInformation{
				Id: scenevalue,
			},
		}
		break
	case QRSceneStringType, QRLimitSceneStringType:
		if name == QRSceneStringType {
			qrcreqdata.Expires = expires
		}

		switch scenevalue.(type) {
		case string:
			break
		default:
			err = fmt.Errorf("scenevalue type is not string")
			return
		}

		qrcreqdata.ActionInfo = &model.ActionInformation{
			SceneInfo: &model.SceneInformation{
				Str: scenevalue,
			},
		}
		break
	default:
		err = fmt.Errorf("Action Name %s invaildable", aname)
		return
	}

	qrcreqdata.ActionName = name

	return
}

// 创建带参数的二维码
func CreateQRcode(token string, data *model.QRcodeRequestData) (ticket *model.QRcodeResponeData, err error) {
	reqline, err := url.Parse(WeChatQRcodeCreateAPIUrl)
	if err != nil {
		return
	}
	v := reqline.Query()
	v.Add("access_token", token)
	reqline.RawQuery = v.Encode()

	var transport *http.Transport
	if strings.HasPrefix(reqline.String(), "https") {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	reqdata, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", reqline.String(), bytes.NewReader(reqdata))
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("create qrcode failed")
		return
	}

	ticket = new(model.QRcodeResponeData)
	err = json.Unmarshal(respdata, ticket)
	if err != nil {
		return
	}

	if ticket.ErrorCode != 0 {
		err = fmt.Errorf("%d, %s", ticket.ErrorCode, ticket.ErrorMessage)
		return
	}

	return ticket, nil
}
