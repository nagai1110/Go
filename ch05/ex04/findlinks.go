// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for data, links := range visit(nil, doc) {
		fmt.Printf("[%s]\n", data)
		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println()
	}
}

func visit(links map[string][]string, n *html.Node) map[string][]string {
	if links == nil {
		links = make(map[string][]string)
	}

	if n.Type == html.ElementNode {
		var link string

		switch n.Data {
		case "a":
			link = getAttr(n, "href")
		case "img":
			link = getAttr(n, "src")
		case "link":
			link = getAttr(n, "href")
		case "script":
			link = getAttr(n, "src")
		}

		if link != "" {
			links[n.Data] = append(links[n.Data], link)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}
