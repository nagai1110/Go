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
		fmt.Fprintf(os.Stderr, "textnode: %v\n", err)
		os.Exit(1)
	}

	for _, text := range findTextnode(nil, doc) {
		fmt.Println(text)
	}
}

func findTextnode(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			texts = append(texts, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = findTextnode(texts, c)
	}
	return texts
}
