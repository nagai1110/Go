// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var tests = []struct {
		values []int
		want   string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{1}"},
		{[]int{1, 2}, "{1 2}"},
		{[]int{1, 2, 3}, "{1 2 3}"},
		{[]int{1, 2, 3, 1, 2, 3}, "{1 2 3}"},
	}

	for _, test := range tests {
		var x IntSet

		x.AddAll(test.values...)
		if x.String() != test.want {
			t.Errorf("AddAll() = %s, want %s", x.String(), test.want)
		}
	}
}
