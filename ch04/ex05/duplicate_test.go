// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestDuplicate(t *testing.T) {
	var tests = []struct {
		s    []string
		want []string
	}{
		{[]string{"1", "1", "1"}, []string{"1"}},
		{[]string{"1", "1", "2", "2", "3", "3"}, []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		out := removeDuplicate(test.s)
		for i := 0; i < len(out); i++ {
			if out[i] != test.want[i] {
				t.Errorf("removeDuplicate(%v) = %v, want %v", test.s, out, test.want)
				break
			}
		}
	}
}
