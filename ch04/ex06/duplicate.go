// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "1  2  3"

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", string(removeDuplicateSpace([]byte(s))))
}

func removeDuplicateSpace(slice []byte) []byte {
	out := slice[:0]

	for i := 0; i < len(slice)-1; i++ {
		if !unicode.IsSpace(rune(slice[i])) || !unicode.IsSpace(rune(slice[i+1])) {
			out = append(out, slice[i])
		}
	}

	out = append(out, slice[len(slice)-1])
	return out
}
