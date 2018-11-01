package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
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

	r := strings.NewReader(string(data))
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(doc.Type)
	if doc.Type == html.ElementNode && doc.Data == "a" {
		fmt.Println([]*html.Node{doc})
	}

	fmt.Println(doc.FirstChild.NextSibling)

	nodes := linkNodes(doc)

	for _, node := range nodes {
		fmt.Println(node)
	}
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
