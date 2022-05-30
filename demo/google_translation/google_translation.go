package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"regexp"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 10:26
 * @Description
 */

type Language string

const (
	TranslationEngineUrl = "https://translate.googleapis.com/translate_a/single"
	ResultPattern        = `\[{3}\"(.*?)\"`

	LanguageCN Language = "zh-CN"
	LanguageEN Language = "en"
)

var (
	ResultRegexp = regexp.MustCompile(ResultPattern)
)

func Translate(text string, sourceLanguage, targetLanguage Language) (string, error) {
	var client http.Client
	url := fmt.Sprintf("%s?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		TranslationEngineUrl, sourceLanguage, targetLanguage, url2.QueryEscape(text))
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 返回的json反序列化比较麻烦, 直接正则提取
	res := ResultRegexp.FindSubmatch(bs)
	if len(res) < 2 {
		return "", nil
	}
	return string(res[1]), nil
}

func main() {
	res, err := Translate("工具", LanguageCN, LanguageEN)
	if err != nil {
		return
	}
	fmt.Println(res)
}
