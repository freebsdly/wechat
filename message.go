package wechat

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
	"wechat/model"
	"wechat/util"
)

// 解密密文消息
func DecryptCipherMessage(token, msgSig, timestamp, nonce string, aesKey, postData []byte) (plainMsg *model.PlainMessage, err error) {
	// 微信post过来的数据是xml格式
	var cmsg = &model.CipherMessage{}
	err = xml.Unmarshal(postData, cmsg)
	if err != nil {
		err = fmt.Errorf("xml decode cipher message failed. %s\n", err)
		return
	}

	// 检查消息签名
	if !util.VerifyMessageSignature(token, msgSig, timestamp, nonce, cmsg.CipherText) {
		err = fmt.Errorf("checksum cipher message signature  failed\n")
		return
	}

	// 密文经过base64编码，需要base64解码
	ciphertext, err := base64.StdEncoding.DecodeString(cmsg.CipherText)
	if err != nil {
		err = fmt.Errorf("base64 decode ciphertext failed. %s\n", err)
		return
	}

	// 解密
	plaintext, err := util.AesCBCDecrypt(aesKey, ciphertext)
	if err != nil {
		err = fmt.Errorf("AesCBC decrypt aes ciphertext failed. %s\n", err)
		return
	}

	return ParsePlainMessage(plaintext)
}

// 解析明文消息
func ParsePlainMessage(content []byte) (msg *model.PlainMessage, err error) {
	if len(content) < MessageLenghtByteSize {
		err = fmt.Errorf("message too short")
		return
	}

	msgLen := util.BytesToInt(content[:MessageLenghtByteSize])
	if len(content) < msgLen+MessageLenghtByteSize {
		err = fmt.Errorf("message too short")
		return
	}

	msg = &model.PlainMessage{PlainText: make([]byte, 0)}
	msg.Lenght = msgLen
	msg.PlainText = content[MessageLenghtByteSize : MessageLenghtByteSize+msgLen]
	msg.AppID = fmt.Sprintf("%s", string(content[MessageLenghtByteSize+msgLen:]))

	return msg, nil

}

// 解析消息类型
func ParseMessageType(content []byte) (msgtype string, err error) {
	cmsg := new(model.CommonMessage)
	err = xml.Unmarshal(content, cmsg)
	if err != nil {
		return
	}

	msgtype = strings.ToLower(cmsg.Type)
	return
}
