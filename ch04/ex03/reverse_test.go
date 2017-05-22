// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		array [6]int
		want  [6]int
	}{
		{[6]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
		{[6]int{6, 7, 8, 9, 10, 11}, [6]int{11, 10, 9, 8, 7, 6}},
	}

	for _, test := range tests {
		reverse(&test.array)
		if test.array != test.want {
			t.Errorf("reverse(%v) = %v, want %v", test.array, test.array, test.want)
		}
	}
}
