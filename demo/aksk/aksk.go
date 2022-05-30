package main

import (
	"crypto"
	"crypto/hmac"
	_ "crypto/sha256"
	"encoding/hex"
	"errors"
)

/**
 * @Author  Flagship
 * @Date  2022/4/27 16:41
 * @Description aksk签名认证，对时间戳进行对称加密，校验身份是否正确
 */

var (
	Secret string
)

const (
	// defaultHash 加密方式
	defaultHash = crypto.SHA256
)

// GenerateSignature 生成签名
func GenerateSignature(content, secret string) (string, error) {
	if !defaultHash.Available() {
		return "", errors.New("the requested hash function is unavailable")
	}

	hasher := hmac.New(defaultHash.New, []byte(secret))
	hasher.Write([]byte(content))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// VerifySignature 校验签名
func VerifySignature(content, sign, secret string) error {
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}
	if !defaultHash.Available() {
		return errors.New("the requested hash function is unavailable")
	}

	hasher := hmac.New(defaultHash.New, []byte(secret))
	hasher.Write([]byte(content))

	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return errors.New("signature is invalid")
	}

	return nil
}
