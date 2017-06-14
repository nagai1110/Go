// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

func join(sep string, values ...string) string {
	var s string
	for i, value := range values {
		if i != 0 {
			s += sep
		}

		s += value
	}

	return s
}

func main() {
	fmt.Println(join(","))
	fmt.Println(join(",", "a"))
	fmt.Println(join(",", "a", "b"))
	fmt.Println(join(":", "abc", "def", "ghi"))
}
