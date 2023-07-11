package captchagenerator

import (
	"encoding/json"
	"net/http"
	"testing"

	redis "Demo-in-Golang/util/redis"
)

func TestCaptchaGenerator(t *testing.T) {
	redis.InitializeRedisInstance()

	// 获取验证码
	http.HandleFunc(
		"/api/getCaptcha", func(writer http.ResponseWriter, request *http.Request) {
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
			err = json.NewEncoder(writer).Encode(body)
			if err != nil {
				t.Fatal(err)
			}
		},
	)

	// 检验验证码
	http.HandleFunc(
		"/api/verifyCaptcha", func(writer http.ResponseWriter, request *http.Request) {
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
			err := json.NewEncoder(writer).Encode(body)
			if err != nil {
				t.Fatal(err)
			}
		},
	)

	if err := http.ListenAndServe(":8777", nil); err != nil {
		t.Fatal(err)
	}
}
