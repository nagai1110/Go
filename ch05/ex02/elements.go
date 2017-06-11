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
		fmt.Fprintf(os.Stderr, "elements: %v\n", err)
		os.Exit(1)
	}

	for element, count := range countElements(nil, doc) {
		fmt.Printf("%s : %d\n", element, count)
	}
}

func countElements(elements map[string]int, n *html.Node) map[string]int {
	if elements == nil {
		elements = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = countElements(elements, c)
	}
	return elements
}
