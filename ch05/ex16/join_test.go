// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestJoin(t *testing.T) {
	var tests = []struct {
		sep    string
		values []string
		want   string
	}{
		{",", []string{}, ""},
		{",", []string{"a"}, "a"},
		{",", []string{"a", "b"}, "a,b"},
		{":", []string{"abc", "def", "ghi"}, "abc:def:ghi"},
	}

	for _, test := range tests {
		v := join(test.sep, test.values...)
		if v != test.want {
			t.Errorf("join(%s, %v) = %s, want %s", test.sep, test.values, v, test.want)
		}
	}
}
