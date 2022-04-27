package main

import (
	"crypto"
	"crypto/hmac"
	_ "crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

var (
	Secret string
)

const (
	defaultHash = crypto.SHA256
)

var Logger *zap.Logger

func main() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		return
	}
	loadConf()

	// 客户端签名
	timestamp := time.Now().Unix()
	content := strconv.FormatInt(timestamp, 10)
	sign, err := GenerateSignature(content, Secret)
	if err != nil {
		Logger.Sugar().Warnf("generate sign failed, err: %s", err.Error())
		return
	}

	// 传输过程
	Logger.Sugar().Debugf("content: %s", content)
	Logger.Sugar().Debugf("sign: %s", sign)

	// 服务端验证签名
	if err = VerifySignature(content, sign, Secret); err != nil {
		Logger.Sugar().Warnf("verify signature failed, err: %s", err)
		return
	}
	Logger.Sugar().Debugf("verify signature success")
}

func loadConf() {
	apikeyConf, err := ini.Load("conf/apikey.conf")
	if err != nil {
		Logger.Sugar().Errorf("get apikey.conf failed, err: %s", err.Error())
		os.Exit(1)
	}
	Secret = apikeyConf.Section("apikey").Key("myapp").String()
}

func GenerateSignature(content, secret string) (string, error) {
	if !defaultHash.Available() {
		return "", errors.New("the requested hash function is unavailable")
	}

	hasher := hmac.New(defaultHash.New, []byte(secret))
	hasher.Write([]byte(content))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

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
