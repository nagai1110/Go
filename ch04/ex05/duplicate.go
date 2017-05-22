// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

func main() {
	s := []string{"1", "1", "2", "1", "2", "2", "3"}

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", removeDuplicate(s))
}

func removeDuplicate(slice []string) []string {
	out := slice[:0]

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != slice[i+1] {
			out = append(out, slice[i])
		}
	}

	out = append(out, slice[len(slice)-1])
	return out
}
