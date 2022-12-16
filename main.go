package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func main() {
	DouBanBook()
}

//	func main() {
//		//爬虫
//		collector := colly.NewCollector(
//			colly.Debugger(&debug.LogDebugger{}), // 开启debug
//			colly.MaxDepth(2),                    //爬取页面深度,最多为两层
//			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"))
//
//		// Find and visit all links
//		collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
//			e.Request.Visit(e.Attr("href"))
//		})
//
//		collector.OnRequest(func(r *colly.Request) {
//			fmt.Println("发起请求之前调用", r.URL)
//		})
//
//		// 请求期间发生错误,则调用
//		collector.OnError(func(response *colly.Response, err error) {
//			fmt.Println("请求期间发生错误,则调用:", err)
//		})
//		// 收到响应后调用
//		collector.OnResponse(func(response *colly.Response) {
//			fmt.Println("收到响应后调用:", response.Body)
//		})
//
//		// url：请求具体的地址
//		err := collector.Visit("请求具体的地址")
//		if err != nil {
//			fmt.Println("具体错误:", err)
//		}
//
//		collector.Visit("https://www.jianshu.com/p/690eb9bebe3c")
//	}
//
// 豆瓣书榜单
func DouBanBook() error {
	// 创建 Collector 对象
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("回调函数OnRequest: 在请求之前调用")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("回调函数OnError: 请求错误", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("回调函数OnResponse: 收到响应后调用")
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("ul[class='subject-list']", func(element *colly.HTMLElement) {
		// 遍历li
		element.ForEach("li", func(i int, el *colly.HTMLElement) {
			// 获取封面图片
			coverImg := el.ChildAttr("div[class='pic'] > a[class='nbg'] > img", "src")
			// 获取书名
			bookName := el.ChildText("div[class='info'] > h2")
			// 获取发版信息，并从中解析出作者名称
			authorInfo := el.ChildText("div[class='info'] > div[class='pub']")
			split := strings.Split(authorInfo, "/")
			author := split[0]
			fmt.Printf("封面: %v 书名:%v 作者:%v\n", coverImg, trimSpace(bookName), author)
		})
	})
	// 发起请求
	return collector.Visit("https://book.douban.com/tag/小说")
}

// 删除字符串中的空格信息
func trimSpace(str string) string {
	// 替换所有的空格
	str = strings.ReplaceAll(str, " ", "")
	// 替换所有的换行
	return strings.ReplaceAll(str, "\n", "")
}
