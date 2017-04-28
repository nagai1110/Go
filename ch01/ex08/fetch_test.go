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
		{"gopl.io", true},
	}

	for _, test := range tests {
		isSucceeded := fetch(test.url)
		if isSucceeded != test.want {
			t.Errorf("fetch(%s) = %t, want %t", test.url, isSucceeded, test.want)
		}
	}
}

func TestCreateUrl(t *testing.T) {
	var tests = []struct {
		url  string
		want string
	}{
		{"http://gopl.io", "http://gopl.io"},
		{"gopl.io", "http://gopl.io"},
	}

	for _, test := range tests {
		url := createUrl(test.url)
		if url != test.want {
			t.Errorf("createUrl(%s) = %s, want %s", test.url, url, test.want)
		}
	}
}
