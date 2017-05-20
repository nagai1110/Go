// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"os"
)

func main() {
	sha1 := sha256.Sum256([]byte(os.Args[1]))
	sha2 := sha256.Sum256([]byte(os.Args[2]))

	fmt.Printf("Different bit count is %v\n", countBitsDiff(sha1, sha2))
}

func countBitsDiff(bytes1, bytes2 [32]byte) int {
	len := int(math.Min(float64(len(bytes1)), float64(len(bytes2))))

	count := 0
	for i := 0; i < len; i++ {
		b := bytes1[i] ^ bytes2[i]
		count += popCount(int(b))
	}

	return count
}

func popCount(x int) int {
	pc := 0
	for x != 0 {
		pc++
		x = x & (x - 1)
	}

	return pc
}
