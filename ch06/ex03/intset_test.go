// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		words1 []int
		words2 []int
		want   string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{}, "{}"},
		{[]int{1, 2}, []int{1}, "{1}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{1 2 3}"},
	}

	for _, test := range tests {
		x1 := makeIntSet(test.words1)
		x2 := makeIntSet(test.words2)

		x1.IntersectWith(&x2)
		if x1.String() != test.want {
			t.Errorf("IntersectWith() = %s, want %s", x1.String(), test.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		words1 []int
		words2 []int
		want   string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{1, 2}, []int{1}, "{2}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{}"},
	}

	for _, test := range tests {
		x1 := makeIntSet(test.words1)
		x2 := makeIntSet(test.words2)

		x1.DifferenceWith(&x2)
		if x1.String() != test.want {
			t.Errorf("DifferenceWith() = %s, want %s", x1.String(), test.want)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		words1 []int
		words2 []int
		want   string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{1, 2}, []int{1}, "{2}"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{1 2 3 4 5 6}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{4 5 6}"},
	}

	for _, test := range tests {
		x1 := makeIntSet(test.words1)
		x2 := makeIntSet(test.words2)

		x1.SymmetricDifference(&x2)
		if x1.String() != test.want {
			t.Errorf("SymmetricDifference() = %s, want %s", x1.String(), test.want)
		}
	}
}

func makeIntSet(words []int) IntSet {
	var s IntSet
	for _, w := range words {
		s.Add(w)
	}

	return s
}
