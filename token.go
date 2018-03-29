package wechat

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"wechat/model"
)

// 获取access_token
func GetAccessToken(appID, appSecret string) (token *model.AccessTokenResponeData, err error) {
	reqline, err := url.Parse(WeChatTokenAPIUrl)
	if err != nil {
		return
	}
	v := reqline.Query()
	v.Add("grant_type", "client_credential")
	v.Add("appid", appID)
	v.Add("secret", appSecret)
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

	req, err := http.NewRequest("GET", reqline.String(), nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("get access token failed")
		return
	}

	msg := new(model.AccessTokenResponeData)
	err = json.Unmarshal(data, msg)
	if err != nil {
		return
	}

	if msg.ErrorCode != 0 {
		err = fmt.Errorf("%d,%s", msg.ErrorCode, msg.ErrorMessage)
		return
	}
	return msg, nil
}
