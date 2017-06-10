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

	for _, text := range textnode(nil, doc) {
		fmt.Println(text)
	}
}

func textnode(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			texts = append(texts, n.Data)
		}
	}

	texts = textnode(texts, n.FirstChild)
	return textnode(texts, n.NextSibling)
}
