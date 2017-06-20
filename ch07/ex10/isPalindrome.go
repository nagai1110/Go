// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"sort"
)

type runes []rune

func (s runes) Len() int           { return len(s) }
func (s runes) Less(i, j int) bool { return s[i] != s[j] }
func (s runes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	palindrome(runes([]rune("TestseT")))
	palindrome(runes([]rune("hogehoge")))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

func palindrome(r runes) {
	if IsPalindrome(r) {
		fmt.Printf("[%s] is palindrome\n", string(r))
	} else {
		fmt.Printf("[%s] is not palindrome\n", string(r))
	}
}
