package main

import (
	"fmt"
	"io/ioutil"
)

type Link struct {
	Href string
	Text string
}

func main() {
	data, err := ioutil.ReadFile("ex01.html")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(data))
}
