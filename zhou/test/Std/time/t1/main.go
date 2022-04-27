package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeDemo()
}

// timeDemo 时间对象的年月日时分秒
func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now.Local())

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
	t := time.Date(2022, 12, 21, 12, 21, 12, 2100, time.Local)
	fmt.Fprint(os.Stdout, t, "\n ")
	c := time.Tick(time.Second)
	fmt.Printf("%T\n", c)
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		}
	}
}
