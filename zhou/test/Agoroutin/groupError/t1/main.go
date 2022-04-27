package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	fetchUrlDemo2()

}
func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.csdn.com",
	}
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				if url == "http://www.csdn.com" {
					fmt.Println(resp.Status)
					fmt.Println(resp.Proto)
					fmt.Println(resp.Cookies())
					h := resp.Header
					for a := range h {
						// fmt.Printf("header[%v]: %v\n", a, h[a])
						// fmt.Printf("header[%v][0]: %v\n", a, h[a][0])
						s := strings.Split(h[a][0], "=")
						if len(s) == 2 {
							i, _ := strconv.Atoi(s[1])
							hrs := i / 3600
							fmt.Printf("s[1]: %v\n", hrs)
						}
					}

					fmt.Println(resp.Location())
					fmt.Println(resp.Request.Method)
					// fmt.Println(resp.Body)
					b, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("body : " + string(b))
				}
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := g.Wait(); err != nil {
		// 处理可能出现的错误
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}
