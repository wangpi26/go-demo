package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 10 //  c
	fmt.Println("你好")
	fmt.Println("发送成功")
}
