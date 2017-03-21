// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestFetch(t *testing.T) {
	var tests = []struct {
		url  string
		want bool
	}{
		{"http://gopl.io", true},
		{"http://bad.gopl.io", false},
	}

	for _, test := range tests {
		isSucceeded := fetch(test.url)
		if isSucceeded != test.want {
			t.Errorf("fetch(%s) = %t, want %t", test.url, isSucceeded, test.want)
		}
	}
}
