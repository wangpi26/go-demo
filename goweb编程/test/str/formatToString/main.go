package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'G', 12, 64)
	c := strconv.FormatInt(1234, 10)
	f := strconv.FormatInt(5, 2)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e, f)
}
