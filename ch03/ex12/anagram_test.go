// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"1", "12", false},
		{"21", "12", true},
		{"abc", "aabc", false},
		{"acb", "bca", true},
		{"あいう", "あいうえ", false},
		{"あいう", "ういあ", true},
	}

	for _, test := range tests {
		b := isAnagram(test.s1, test.s2)
		if b != test.want {
			t.Errorf("isAnagram(%s, %s) = %t, want %t", test.s1, test.s2, b, test.want)
		}
	}
}
