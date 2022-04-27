package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("helloworld")
	str := "I am ziop, How are you?"
	fmt.Println(str)

	// strings.Contains
	fmt.Println("Test Contains: { ", str, " } { am }", strings.Contains(str, "am"))
	fmt.Println("Test Contains: { ", str, " } { Am }", strings.Contains(str, "Am"))

	// JOIN
	s := []string{"I", "am", "fine"}
	fmt.Println(strings.Join(s, " "))

	// Repeat
	fmt.Println(strings.Repeat("哈", 10))
}
