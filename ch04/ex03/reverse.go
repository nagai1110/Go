// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	fmt.Printf("%v\n", a)
	reverse(&a)
	fmt.Printf("%v\n", a)
}

func reverse(array *[6]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
