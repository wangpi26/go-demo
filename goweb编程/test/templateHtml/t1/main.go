package main

import (
	"fmt"
	"html/template"
	"os"
)

type Persion struct {
	Name string
	Age  int
}

func main() {
	t := template.New("some template")
	t.Parse(`name = {{.Name}}
age = {{.Age}}\n`)
	persion := Persion{"Jac", 18}
	t.Execute(os.Stdout, persion)
	fmt.Println("vim-go")
}
