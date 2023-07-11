package main

import (
	"testing"
	"time"

	rand "Demo-in-Golang/util/rand"
)

/**
 * @Author  Flagship
 * @Date  2022/4/27 16:03
 * @Description
 */

func TestToken(t *testing.T) {
	token, err := GenerateToken("flagship")
	if err != nil {
		t.Errorf("generate token failed, err: %s", err.Error())
		return
	}
	t.Logf("token[%s]", token)
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("parse token tailed, err: %s", err.Error())
		return
	}
	t.Logf("user_id[%s]", claims.UID)
	if claims.ExpiresAt.Before(time.Now()) {
		t.Errorf("this token is expired, err: %s", err.Error())
		return
	}
}

func TestMain(m *testing.M) {
	AccessSecret = rand.RandStr(
		64, rand.RandMode{
			Numbers:      true,
			LowerLetters: true,
			UpperLetters: true,
		},
	)
	AccessExpire = 7

	m.Run()
}
