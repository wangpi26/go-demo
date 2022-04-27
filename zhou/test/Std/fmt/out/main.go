package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./ziop.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "李政"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s\n", name)
	fileObj.Close()
	err1 := errors.New("big Error")
	fmt.Printf("%s\n", fmt.Errorf("this is a error, %w", err1).Error())
}
