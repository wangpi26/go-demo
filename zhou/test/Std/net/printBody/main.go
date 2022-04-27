package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://101.43.250.216:8099"
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("http request get failed,err:%v", err)
	}
	b, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Printf("read body failed,%v", err2)
	}
	fmt.Println(string(b))
	r.Body.Close()
}
