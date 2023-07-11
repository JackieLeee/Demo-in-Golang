package main

import (
	"strconv"
	"testing"
	"time"

	myrand "Demo-in-Golang/util/rand"
)

/**
 * @Author  JackieLee
 * @Date  2022/4/27 16:41
 * @Description
 */

func TestAkSk(t *testing.T) {
	// 客户端签名
	timestamp := time.Now().Unix()
	content := strconv.FormatInt(timestamp, 10)
	sign, err := GenerateSignature(content, Secret)
	if err != nil {
		t.Errorf("generate sign failed, err: %s", err.Error())
		return
	}

	// 传输过程
	t.Logf("content: %s", content)
	t.Logf("sign: %s", sign)

	// 服务端验证签名
	if err = VerifySignature(content, sign, Secret); err != nil {
		t.Errorf("verify signature failed, err: %s", err)
		return
	}
	t.Logf("verify signature success")
}

func TestMain(m *testing.M) {
	Secret = myrand.RandStr(
		64, myrand.RandMode{
			Numbers:      true,
			LowerLetters: true,
			UpperLetters: true,
		},
	)

	m.Run()
}
