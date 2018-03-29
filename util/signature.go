package util

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// 生成签名
func makeSignature(token, timestamp, nonce string) string {
	str := []string{token, timestamp, nonce}
	sort.Strings(str)
	return sha1String(strings.Join(str, ""))
}

// 生成消息签名
func makeMessageSignature(token, timestamp, nonce, encryptMsg string) string {
	str := []string{token, timestamp, nonce, encryptMsg}
	sort.Strings(str)
	return sha1String(strings.Join(str, ""))
}

// 判断签名
func VerifySignature(token, sig, timestamp, nonce string) bool {
	if makeSignature(token, timestamp, nonce) == sig {
		return true
	} else {
		return false
	}
}

// 判断消息签名
func VerifyMessageSignature(token, msgSig, timestamp, nonce, encryptMsg string) bool {
	if makeMessageSignature(token, timestamp, nonce, encryptMsg) == msgSig {
		return true
	} else {
		return false
	}
}

// 生成sha1签名
func sha1String(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}
