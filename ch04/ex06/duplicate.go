// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "1  2  3"

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", string(removeDuplicateSpace([]byte(s))))
}

// TODO:日本語などのマルチバイト文字できてない
func removeDuplicateSpace(slice []byte) []byte {
	out := make([]byte, len(slice))

	index := 0
	for i := 0; i < len(slice); {
		r1, size := utf8.DecodeRune(slice[i:])
		if len(slice) <= i+size {
			out[index] = byte(r1)
			break
		}

		r2, _ := utf8.DecodeRune(slice[i+size:])
		if !unicode.IsSpace(r1) || !unicode.IsSpace(r2) {
			out[index] = byte(r1)
			index++
		}

		i += size
	}

	return out
}
