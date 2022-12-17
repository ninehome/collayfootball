package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	// 创建 collector
	collector := colly.NewCollector()

	// 设置UA
	//collector.UserAgent = USER_AGENT

	// 事件监听
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("url:", r.URL.String())
		// url: https://www.baidu.com/
	})

	// 解析元素
	collector.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		// 百度一下，你就知道
	})

	// 访问网页
	collector.Visit("https://www.huxiu.com/")
}
