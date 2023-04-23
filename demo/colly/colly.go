package colly

import (
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

var (
	globalCollector *colly.Collector
	loggerWriter    io.Writer

	globalCollectorLock sync.Mutex
	loggerWriterLock    sync.Mutex
)

type pageInfo struct {
	StatusCode int
	Links      map[string]int
}

func GetGlobalCollector() *colly.Collector {
	if globalCollector == nil {
		globalCollectorLock.Lock()
		defer globalCollectorLock.Unlock()
		// 初始化 Collector
		c := colly.NewCollector(
			colly.UserAgent("xy"),   // 设置 User-Agent
			colly.AllowURLRevisit(), // 允许重复访问,
			colly.Debugger(
				&debug.LogDebugger{Output: GetGlobalLoggerWriter()}, // 设置日志输出
			),
			colly.MaxDepth(2), // 设置最大深度
		)
		// 设置http连接参数
		c.WithTransport(&http.Transport{
			//DisableKeepAlives: false, // 禁用keepAlive
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // 超时时间
				KeepAlive: 30 * time.Second, // keepAlive 超时时间
			}).DialContext,
			MaxIdleConns:          100,              // 最大空闲连接数
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时
			TLSHandshakeTimeout:   10 * time.Second, // TLS握手超时
			ExpectContinueTimeout: 1 * time.Second,  // 预期等待超时
		})
		globalCollector = c
	}
	return globalCollector
}

func GetGlobalLoggerWriter() io.Writer {
	if loggerWriter == nil {
		loggerWriterLock.Lock()
		defer loggerWriterLock.Unlock()
		// 初始化日志输出文件
		var err error
		loggerWriter, err = os.OpenFile("collector.log", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
	}
	return loggerWriter
}

func GetProxySwitcher() colly.ProxyFunc {
	//// 通过轮询方式实现代理切换
	//if p, proxyErr := proxy.RoundRobinProxySwitcher(
	//	"socks5://127.0.0.1:1337",
	//	"socks5://127.0.0.1:1338",
	//	"http://127.0.0.1:8080",
	//); proxyErr == nil {
	//	return p
	//}
	// 设置随机的代理
	var proxies = []*url.URL{
		{Host: "127.0.0.1:8080"},
		{Host: "127.0.0.1:8081"},
	}
	return func(_ *http.Request) (*url.URL, error) {
		return proxies[rand.Intn(len(proxies))], nil
	}
}
