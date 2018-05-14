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

// 创建自定义菜单
func CreateCustomMenu(token string, data *model.Menu) (err error) {
	reqline, err := url.Parse(WeChatMenuCreateAPIUrl)
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
		err = fmt.Errorf("create menu failed, statuescode %d", resp.StatusCode)
		return
	}

	r := new(model.CommonResponeData)
	err = json.Unmarshal(respdata, r)
	if err != nil {
		return
	}

	if r.ErrorCode != 0 {
		err = fmt.Errorf("%d, %s", r.ErrorCode, r.ErrorMessage)
		return
	}

	return nil
}

// 查询自定义菜单
//func GetCustomMenu(token string) {
//}
