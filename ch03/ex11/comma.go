// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", commaFloat(os.Args[i]))
	}
}

func commaFloat(s string) string {
	var buf bytes.Buffer

	f := strings.Split(s, ".")

	// 整数部
	integer := f[0]
	if strings.ContainsAny(f[0], "+-") {
		buf.WriteByte(integer[0])
		integer = integer[1:]
	}

	buf.WriteString(commma(integer))

	if len(f) < 2 {
		return buf.String()
	}

	// 小数部
	buf.WriteRune('.')
	buf.WriteString(f[1])

	return buf.String()
}

func commma(s string) string {
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
