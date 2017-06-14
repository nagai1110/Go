// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"math"
)

func max(vals ...int) int {
	max := math.MinInt64
	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}

func max2(v int, vals ...int) int {
	max := v
	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}

func min(vals ...int) int {
	min := math.MaxInt64
	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}

func min2(v int, vals ...int) int {
	min := v
	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}

func main() {
	fmt.Println(max())
	fmt.Println(min())
	fmt.Println(max(1, 2))
	fmt.Println(min(1, 2))
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(max2(1, 2))
	fmt.Println(min2(1, 2))
	fmt.Println(max2(1, 2, 3, 4))
	fmt.Println(min2(1, 2, 3, 4))
}
