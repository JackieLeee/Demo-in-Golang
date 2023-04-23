package colly

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"testing"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/gocolly/colly/v2/queue"
	"github.com/gocolly/redisstorage"

	myrand "Demo-in-Golang/util/rand"
)

func TestSimpleColly(t *testing.T) {
	// 初始化Collector
	c := GetGlobalCollector()
	//c.SetProxyFunc(GetProxySwitcher()) // 设置代理
	c.AllowedDomains = []string{"go-colly.org"} // 设置允许的域名

	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "httpbin_test",
	}
	// 设置存储
	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}
	// 清空存储
	if clearErr := storage.Clear(); clearErr != nil {
		log.Fatal(clearErr)
	}

	// 关闭连接
	defer storage.Client.Close()
	q, _ := queue.New(2, storage)

	// colly扩展
	extensions.RandomUserAgent(c) // 设置随机User-Agent
	extensions.Referer(c)         // 设置Referer
	urlSet := sync.Map{}

	// c.OnRequest() 请求执行之前调用
	// c.OnResponse() 响应返回之后调用
	// c.OnHTML() 监听执行 selector
	// c.OnXML() 监听执行 selector
	// c.OnHTMLDetach()，取消监听，参数为 selector 字符串
	// c.OnXMLDetach()，取消监听，参数为 selector 字符串
	// c.OnScraped()，完成抓取后执行，完成所有工作后执行
	// c.OnError()，错误回调
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		urlStr := e.Attr("href")
		// 判断是否合法的url
		_, ok := urlSet.Load(urlStr)
		if ok {
			return
		}
		if u, parseErr := url.Parse(urlStr); parseErr == nil && u.Scheme != "" && u.Host != "" {
			//q.AddURL(e.Attr("href"))
			c.Visit(e.Attr("href"))
		}
		urlSet.Store(urlStr, struct{}{})
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", myrand.RandStr(rand.Intn(10)+10, myrand.RandMode{
			LowerLetters: true,
			UpperLetters: true,
		})) // 随机设置 User-Agent
		fmt.Println("visiting", r.URL)
	})
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Visited", r.Request.URL)
	//})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("visit %s got error: %s \n", r.Request.URL, err.Error())
	})

	c.Async = true // 设置异步
	urlSet.Store("http://go-colly.org/", struct{}{})
	q.AddURL("http://go-colly.org/")
	q.Run(c)
	//c.Visit("http://go-colly.org/")
	c.Wait() // 等待异步执行完成
	fmt.Println("end")
}

func TestDistributedColly(t *testing.T) {
	// example usage: curl -s 'http://127.0.0.1:7171/?url=http://go-colly.org/'
	addr := ":7171"

	http.HandleFunc("/", handlerDistributedColly)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handlerDistributedColly 分布式爬虫——执行层面
func handlerDistributedColly(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", URL)

	c := GetGlobalCollector()

	p := &pageInfo{Links: make(map[string]int)}
	// 收集所有链接
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			p.Links[link]++
		}
	})
	// 打印调用状态码
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		p.StatusCode = r.StatusCode
	})
	// 打印错误
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
		p.StatusCode = r.StatusCode
	})

	// 开始访问
	c.Visit(URL)

	// 序列化收集信息
	b, err := json.Marshal(p)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	// 返回序列化后的信息
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

// 多收集器
func TestMultipleColly(t *testing.T) {
	c := GetGlobalCollector()
	c2 := c.Clone()
	// 复制出来c和c2的配置是一样的，这里复制出来后可以基于c的配置进行自定义的修改
	c2.UserAgent = "colly2"

	// 不通collector之间的数据传递，这个Context只是colly实现的数据共享的结构，并非Go标准库中的Contex
	c.OnResponse(func(r *colly.Response) {
		r.Ctx.Put("Custom-header", r.Headers.Get("Custom-Header"))
		c2.Request("GET", "https://foo.com/", nil, r.Ctx, nil)
	})
}
