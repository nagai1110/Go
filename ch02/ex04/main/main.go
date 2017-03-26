// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"

	"github.com/nagai1110/Go/ch02/ex04/popcount"
)

func main() {
	fmt.Printf("PopCount(0) : %d\n", popcount.PopCount(0))
	fmt.Printf("PopCount(1) : %d\n", popcount.PopCount(1))
	fmt.Printf("PopCount(255) : %d\n", popcount.PopCount(255))
	fmt.Printf("PopCount(256) : %d\n", popcount.PopCount(256))

	fmt.Printf("PopCountByBitShift(0) : %d\n", popcount.PopCountByBitShift(0))
	fmt.Printf("PopCountByBitShift(1) : %d\n", popcount.PopCountByBitShift(1))
	fmt.Printf("PopCountByBitShift(255) : %d\n", popcount.PopCountByBitShift(255))
	fmt.Printf("PopCountByBitShift(256) : %d\n", popcount.PopCountByBitShift(256))
}
