package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("directory = " + dir)
	os.Mkdir("ziop", 0766)
	// file, _ := os.Create("ziop/ziop.txt")
	// file.WriteString("hello ziop")
	file, _ := os.OpenFile("ziop.txt", os.O_RDWR, 0766)
	file.WriteString("jkdfaklsfls后会有期   ")
}
