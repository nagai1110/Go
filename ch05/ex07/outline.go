// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if hasChild(n) {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrString(n))
		} else {
			fmt.Printf("%*s<%s%s />", depth*2, "", n.Data, attrString(n))
		}

		depth++
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--

		if hasChild(n) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func hasChild(n *html.Node) bool {
	return n.FirstChild != nil
}

func attrString(n *html.Node) string {
	s := ""

	for _, a := range n.Attr {
		s += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
	}

	return s
}
