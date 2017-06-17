// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		words []int
		want  string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{1}"},
		{[]int{1, 2}, "{1 2}"},
		{[]int{1, 2, 3}, "{1 2 3}"},
	}

	for _, test := range tests {
		var x IntSet
		for _, w := range test.words {
			x.Add(w)
		}

		if x.String() != test.want {
			t.Errorf("Add() = %s, want %s", x.String(), test.want)
		}
	}
}

func TestHas(t *testing.T) {
	var tests = []struct {
		words []int
		word  int
		want  bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1, 2}, 3, false},
		{[]int{1, 2, 3}, 3, true},
	}

	for _, test := range tests {
		x := makeIntSet(test.words)

		b := x.Has(test.word)
		if b != test.want {
			t.Errorf("Has() = %t, want %t", b, test.want)
		}
	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		words1 []int
		words2 []int
		want   string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 2, 3}, []int{2, 3, 4}, "{1 2 3 4}"},
	}

	for _, test := range tests {
		x1 := makeIntSet(test.words1)
		x2 := makeIntSet(test.words2)

		x1.UnionWith(&x2)
		if x1.String() != test.want {
			t.Errorf("UnionWith() = %s, want %s", x1.String(), test.want)
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
