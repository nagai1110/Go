// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestDouble(t *testing.T) {
	var tests = []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, test := range tests {
		r := double(test.x)
		if r != test.want {
			t.Errorf("double(%d) = %d, want %d", test.x, r, test.want)
		}
	}
}
