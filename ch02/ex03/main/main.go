// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"

	"github.com/nagai1110/Go/ch02/ex03/popcount"
)

func main() {
	fmt.Printf("PopCount(0) : %d\n", popcount.PopCount(0))
	fmt.Printf("PopCount(1) : %d\n", popcount.PopCount(1))
	fmt.Printf("PopCount(255) : %d\n", popcount.PopCount(255))
	fmt.Printf("PopCount(256) : %d\n", popcount.PopCount(256))

	fmt.Printf("PopCountByLoop(0) : %d\n", popcount.PopCountByLoop(0))
	fmt.Printf("PopCountByLoop(1) : %d\n", popcount.PopCountByLoop(1))
	fmt.Printf("PopCountByLoop(255) : %d\n", popcount.PopCountByLoop(255))
	fmt.Printf("PopCountByLoop(256) : %d\n", popcount.PopCountByLoop(256))
}
