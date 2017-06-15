// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

func double(x int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 2 * x
		}
	}()

	panic(0)
}

func main() {
	fmt.Println(double(1))
	fmt.Println(double(2))
	fmt.Println(double(3))
}
