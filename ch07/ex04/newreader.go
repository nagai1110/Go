// Copyright © 2017 Takaya.Nagai All Rights Res
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type StringReader string

func (sr *StringReader) Read(p []byte) (n int, err error) {
	n = len(*sr)
	copy(p, []byte(*sr))
	err = io.EOF
	return
}

func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}

func htmlString() string {
	return `<html>
<body>

<h1>hello</h1>

</body>
</html>"`
}

func main() {
	reader := NewReader(htmlString())

	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "newreader: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(elements []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		elements = append(elements, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = visit(elements, c)
	}

	return elements
}
