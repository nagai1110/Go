// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		adds []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 1, 2, 3}, 3},
	}

	for _, test := range tests {
		var i IntSet
		for _, w := range test.adds {
			i.Add(w)
		}

		if i.Len() != test.want {
			t.Errorf("Len() = %d, want %d", i.Len(), test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		adds    []int
		removes []int
		want    string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{}"},
		{[]int{1, 2}, []int{1}, "{2}"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{1 2 3}"},
	}

	for _, test := range tests {
		var x IntSet
		for _, w := range test.adds {
			x.Add(w)
		}

		for _, w := range test.removes {
			x.Remove(w)
		}

		if x.String() != test.want {
			t.Errorf("words = %s, want %s", x.String(), test.want)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		adds []int
		want string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{}"},
		{[]int{1, 2}, "{}"},
		{[]int{1, 2, 3}, "{}"},
	}

	for _, test := range tests {
		var x IntSet
		for _, w := range test.adds {
			x.Add(w)
		}

		x.Clear()
		if x.String() != test.want {
			t.Errorf("words = %s, want %s", x.String(), test.want)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		adds []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
	}

	for _, test := range tests {
		var x IntSet
		for _, w := range test.adds {
			x.Add(w)
		}

		copy := x.Copy()
		if x.String() != copy.String() {
			t.Errorf("words = %s, want %s", x.String(), copy.String())
		}
	}
}
