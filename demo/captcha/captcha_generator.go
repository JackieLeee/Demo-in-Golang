package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mojocn/base64Captcha"

	myredis "Demo-in-Golang/util/redis"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 11:47
 * @Description 验证码生成
 */

const (
	CaptchaHeight = 60
	CaptchaWidth  = 240

	CaptchaExpireTime = time.Duration(10) * time.Minute

	CaptchaPrefix = "captcha:%s"
)

type RedisStore struct {
}

func (r RedisStore) Set(id string, value string) error {
	key := fmt.Sprintf(CaptchaPrefix, id)
	if err := myredis.Client.Set(key, value, CaptchaExpireTime).Err(); err != nil {
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := fmt.Sprintf(CaptchaPrefix, id)
	res, err := myredis.Client.Get(key).Result()
	if err != nil {
		return ""
	}
	if clear {
		if err := myredis.Client.Del(key).Err(); err != nil {
			return ""
		}
	}
	return res
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}

var (
	// store = base64Captcha.DefaultMemStore
	store = RedisStore{}
)

func getDefaultMathDriver() base64Captcha.Driver {
	driver := &base64Captcha.DriverMath{
		Height: CaptchaHeight,
		Width:  CaptchaWidth,
	}
	return driver
}

func GenerateCaptcha() (string, string, error) {
	driver := getDefaultMathDriver()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, base64, err := captcha.Generate()
	if err != nil {
		return "", "", err
	}
	return id, base64, nil
}

func VerifyCaptcha(id, value string) bool {
	return store.Verify(id, value, true)
}

func main() {
	myredis.InitializeRedisInstance()

	// 获取验证码
	http.HandleFunc("/api/getCaptcha", func(writer http.ResponseWriter, request *http.Request) {
		id, base64, err := GenerateCaptcha()
		if err != nil {
			return
		}
		body := map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]string{
				"base64": base64,
				"id":     id,
			},
		}
		json.NewEncoder(writer).Encode(body)
	})

	// 检验验证码
	http.HandleFunc("/api/verifyCaptcha", func(writer http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		type Param struct {
			Id     string `json:"id"`
			Answer string `json:"answer"`
		}
		var param Param
		if err := decoder.Decode(&param); err != nil {
			return
		}
		body := map[string]interface{}{"code": 1, "msg": "failed"}
		if VerifyCaptcha(param.Id, param.Answer) {
			body = map[string]interface{}{"code": 0, "msg": "ok"}
		}
		json.NewEncoder(writer).Encode(body)
	})

	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
