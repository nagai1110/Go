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

	nodes := ElementsByTagName(doc, os.Args[2:]...)
	for _, n := range nodes {
		fmt.Printf("%v\n", n)
	}
}

func ElementsByTagName(doc *html.Node, tagNames ...string) []*html.Node {
	var foundNodes []*html.Node

	if len(tagNames) == 0 {
		return nil
	}

	startElement := func(n *html.Node) bool {
		for _, tagName := range tagNames {
			if n.Type == html.ElementNode && n.Data == tagName {
				foundNodes = append(foundNodes, n)
			}
		}

		return true
	}

	endElement := func(n *html.Node) bool {
		return true
	}

	forEachNode(doc, startElement, endElement)

	return foundNodes
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
