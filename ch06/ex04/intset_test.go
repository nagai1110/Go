// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package intset

import (
	"testing"
)

func TestElems(t *testing.T) {
	var tests = []struct {
		words []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
	}

	for _, test := range tests {
		x := makeIntSet(test.words)

		for i, word := range x.Elems() {
			if word != test.words[i] {
				t.Errorf("Elems() = %v, want %v", x.Elems(), test.words)
				break
			}
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
