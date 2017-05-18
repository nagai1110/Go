// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	offset := n % 3

	var buf bytes.Buffer
	buf.WriteString(s[:offset])

	for i := offset; i < n; i++ {
		if i != 0 && i%3 == offset {
			buf.WriteRune(',')
		}

		buf.WriteByte(s[i])
	}

	return buf.String()
}
