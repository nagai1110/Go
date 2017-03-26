// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"

	"github.com/nagai1110/Go/ch02/ex05/popcount"
)

func main() {
	fmt.Printf("PopCount(0) : %d\n", popcount.PopCount(0))
	fmt.Printf("PopCount(1) : %d\n", popcount.PopCount(1))
	fmt.Printf("PopCount(255) : %d\n", popcount.PopCount(255))
	fmt.Printf("PopCount(256) : %d\n", popcount.PopCount(256))

	fmt.Printf("PopCountByLSBClearance(0) : %d\n", popcount.PopCountByLSBClearance(0))
	fmt.Printf("PopCountByLSBClearance(1) : %d\n", popcount.PopCountByLSBClearance(1))
	fmt.Printf("PopCountByLSBClearance(255) : %d\n", popcount.PopCountByLSBClearance(255))
	fmt.Printf("PopCountByLSBClearance(256) : %d\n", popcount.PopCountByLSBClearance(256))
}
