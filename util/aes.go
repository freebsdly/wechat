package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// 使用AES的cbc模式加密明文
func AesCBCEncrypt(key, plaintext []byte) (ciphertext []byte, err error) {
	//秘钥长度需要是AES-256(32bytes)
	if len(key) != 32 {
		err = fmt.Errorf("aes key must be 32 bytes")
		return
	}

	//原文必须填充至blocksize的整数倍，填充方法可以参见https://tools.ietf.org/html/rfc5246#section-6.2.3.2
	//块大小在aes.BlockSize中定义
	if len(plaintext)%aes.BlockSize != 0 {
		plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	}

	//生成加密用的block
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// 对IV有随机性要求，但没有保密性要求，所以常见的做法是将IV包含在加密文本当中
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	//随机一个block大小作为IV
	//采用不同的IV时相同的秘钥将会产生不同的密文，可以理解为一次加密的session
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 谨记密文需要认证(i.e. by using crypto/hmac)
	return ciphertext, nil
}

// 使用AES的CBC模式解密
func AesCBCDecrypt(key, ciphertext []byte) (plaintext []byte, err error) {
	if len(ciphertext) < aes.BlockSize {
		err = fmt.Errorf("ciphertext too short")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size")
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks可以原地更新
	mode.CryptBlocks(ciphertext, ciphertext)

	return PKCS7UnPadding(ciphertext, aes.BlockSize), nil
}

// 将明文填充至blocksize的整数倍
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 将解密后的明文反补位
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
