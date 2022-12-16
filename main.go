package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func main() {
	//爬虫
	collector := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.MaxDepth(2),                    //爬取页面深度,最多为两层
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"))

	// Find and visit all links
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("发起请求之前调用", r.URL)
	})

	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("请求期间发生错误,则调用:", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("收到响应后调用:", response.Body)
	})

	// url：请求具体的地址
	err := collector.Visit("请求具体的地址")
	if err != nil {
		fmt.Println("具体错误:", err)
	}

	collector.Visit("https://www.jianshu.com/p/690eb9bebe3c")
}
