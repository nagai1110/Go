// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"$abc", "ABC"},
		{"abc", "abc"},
		{"$abc def $ghi", "ABC def GHI"},
		{"abc def ghi", "abc def ghi"},
	}

	for _, test := range tests {
		s := expand(test.s, strings.ToUpper)
		if s != test.want {
			t.Errorf("expand(%s) = %s, want %s", test.s, s, test.want)
		}
	}
}
