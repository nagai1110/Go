// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	b := isAnagram(os.Args[1], os.Args[2])
	if b {
		fmt.Printf("%s, %s are anagram", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s, %s are not anagram", os.Args[1], os.Args[2])
	}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	slice1 := strings.Split(s1, "")
	slice2 := strings.Split(s2, "")
	sort.Strings(slice1)
	sort.Strings(slice2)

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
