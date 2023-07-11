package captchagenerator

import (
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"

	redis "Demo-in-Golang/util/redis"
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
	if err := redis.Client.Set(key, value, CaptchaExpireTime).Err(); err != nil {
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := fmt.Sprintf(CaptchaPrefix, id)
	res, err := redis.Client.Get(key).Result()
	if err != nil {
		return ""
	}
	if clear {
		if err := redis.Client.Del(key).Err(); err != nil {
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
