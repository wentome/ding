// alert
package ding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strconv"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Mannager struct {
	Url          string
	Access_token string
	Secret       string
}

type Ding interface {
	SendSignMsg(message string) (string, error)
}

func NewDing(url string, access_token string, secret string) Ding {
	m := new(Mannager)
	m.Url = url
	m.Access_token = access_token
	m.Secret = secret
	return m
}

func (m *Mannager) SendSignMsg(message string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	hash := hmac.New(sha256.New, []byte(m.Secret))
	hash.Write([]byte(timestamp + "\n" + m.Secret)) // 写入加密数据
	// StdEncoding有 ”+“  URLEncoding有”-“ 结果是不一样的
	hmac_code := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	sign := url.QueryEscape(hmac_code)

	sendMessage := fmt.Sprintf(`{"msgtype":"text","text":{"content":"%s"}}`, message)
	url := fmt.Sprintf("%s?access_token=%s&timestamp=%s&sign=%s", m.Url, m.Access_token, timestamp, sign)
	resp, err := http.Post(url, "application/json", strings.NewReader(sendMessage))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}


