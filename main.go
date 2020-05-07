package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello, and welcome to Goils")
	letter := "Dear {{.Name}}"

	type Recipient struct {
		Name string
	}

	recipient := Recipient{Name: "Bob"}
	t := template.Must(template.New("letter").Parse(letter))
	err := t.Execute(os.Stdout, recipient)
	if err != nil {
		fmt.Println("err:", err)
	}
}
