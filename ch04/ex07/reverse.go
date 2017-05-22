// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "あいうえお"
	b := []byte(s)

	fmt.Printf("%v\n", string(b))
	fmt.Printf("%v\n", string(reverse(b)))
}

func reverse(b []byte) []byte {
	runes := bytes.Runes(b)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
