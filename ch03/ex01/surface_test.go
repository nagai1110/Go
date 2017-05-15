// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestCorner(t *testing.T) {
	var tests = []struct {
		i  int
		j  int
		ok bool
	}{
		{50, 50, false},
		{50, 49, true},
		{49, 50, true},
	}

	for _, test := range tests {
		_, _, ok := corner(test.i, test.j)
		if ok != test.ok {
			t.Errorf("corner(%v, %v) = %v, want %v", test.i, test.j, ok, test.ok)
		}
	}
}
