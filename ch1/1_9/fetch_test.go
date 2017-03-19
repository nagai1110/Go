// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	var tests = []struct {
		url  string
		want string
	}{
		{"http://gopl.io", "status code: 200"},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		fetch(test.url)

		got := out.(*bytes.Buffer).String()
		if !strings.Contains(got, test.want) {
			t.Errorf("fetch(%s) = %s, Not contains %s", test.url, got, test.want)
		}
	}
}
