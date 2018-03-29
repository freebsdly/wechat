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
	"wechat/model"
)

// 获取订阅用户基本信息
func GetUserInformation(token string, data *model.UserInfoRequestData) (userinfo *model.UserInformation, err error) {
	reqline, err := url.Parse(WeChatUserInfoAPIUrl)
	if err != nil {
		return
	}
	v := reqline.Query()
	v.Add("access_token", token)
	v.Add("openid", data.OpenId)
	v.Add("lang", data.Language)
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

	respdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("call getuserinfo http api return status code %d\n", resp.StatusCode)
		return
	}

	userinfo = new(model.UserInformation)
	err = json.Unmarshal(respdata, userinfo)
	if err != nil {
		return
	}

	if userinfo.ErrorCode != 0 {
		err = fmt.Errorf("%d, %s", userinfo.ErrorCode, userinfo.ErrorMessage)
		return
	}

	return userinfo, nil
}

// 批量获取用户基本信息
func BatchGetUserInformation(token string, data *model.BatchUserInfoRequestData) (userlist *model.UserInformationList, err error) {
	reqline, err := url.Parse(WechatUserInfoBatchAPIUrl)
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
		err = fmt.Errorf("call batchgetuserinfo http api return status code %d\n", resp.StatusCode)
		return
	}

	userlist = model.NewUserInformationList()
	err = json.Unmarshal(respdata, userlist)
	if err != nil {
		return
	}

	if userlist.ErrorCode != 0 {
		err = fmt.Errorf("%d, %s", userlist.ErrorCode, userlist.ErrorMessage)
		return
	}

	return userlist, nil
}
