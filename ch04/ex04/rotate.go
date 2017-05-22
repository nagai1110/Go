// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	fmt.Printf("%v\n", s)
	rotate(s, 2)
	fmt.Printf("%v\n", s)
}

func rotate(s []int, num int) {
	length := len(s)
	for i := 0; i < length-num; i++ {
		j := (i + num) % length
		s[i], s[j] = s[j], s[i]
	}
}
