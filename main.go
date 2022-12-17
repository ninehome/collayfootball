package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {

	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "50,400"),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	//taskCtx, cancel = context.WithTimeout(taskCtx, 10*time.Second)
	//defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// listen network event
	listenForNetworkEvent(taskCtx)

	chromedp.Run(taskCtx,
		network.Enable(),
		chromedp.Navigate(`https://test.f66b88sport.com/pc/index.html#/`),
		chromedp.WaitVisible(`body`, chromedp.BySearch),
		chromedp.Sleep(6*time.Second),
	)

	time.Sleep(30 * time.Second)

}

// 监听
func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {

		case *network.EventResponseReceived:
			resp := ev.Response

			if len(resp.Headers) != 0 {
				// log.Printf("received headers: %s", resp.Headers)

				//if resp.URL == "https://sportapi.fastball2.com/v1/match/getList" {
				//	fmt.Println("received headers: %s", resp.URL)
				//	fmt.Println(resp)
				//}

				if strings.Index(resp.URL, ".ts") != -1 {
					log.Printf("received headers: %s", resp.URL)
					fmt.Println("received headers: %s", resp.URL)
				}
			}

		}
		// other needed network Event
	})
}
