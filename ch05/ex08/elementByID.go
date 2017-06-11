// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, _ := http.Get(os.Args[1])
	defer resp.Body.Close()

	doc, _ := html.Parse(resp.Body)

	for _, id := range os.Args[2:] {
		n := ElementByID(doc, id)
		if n != nil {
			fmt.Printf("%v\n", n)
		}
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var foundNode *html.Node

	startElement := func(n *html.Node) bool {
		if foundNode != nil {
			return false
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				foundNode = n
				return false
			}
		}

		return true
	}

	endElement := func(n *html.Node) bool {
		return true
	}

	forEachNode(doc, startElement, endElement)

	return foundNode
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(n) {
		return
	}
}
