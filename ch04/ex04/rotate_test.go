// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		s    []int
		num  int
		want []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
		{[]int{6, 7, 8, 9, 10, 11}, 3, []int{9, 10, 11, 6, 7, 8}},
	}

	for _, test := range tests {
		rotate(test.s, test.num)
		for i := 0; i < len(test.s); i++ {
			if test.s[i] != test.want[i] {
				t.Errorf("rotate(%v, %d) = %v, want %v", test.s, test.num, test.s, test.want)
				break
			}
		}
	}
}
