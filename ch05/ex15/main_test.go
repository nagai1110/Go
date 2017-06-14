// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		values []int
		want   int
	}{
		{[]int{}, math.MinInt64},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
	}

	for _, test := range tests {
		v := max(test.values...)
		if v != test.want {
			t.Errorf("max(%v) = %d, want %d", test.values, v, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		values []int
		want   int
	}{
		{[]int{}, math.MaxInt64},
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
	}

	for _, test := range tests {
		v := min(test.values...)
		if v != test.want {
			t.Errorf("min(%v) = %d, want %d", test.values, v, test.want)
		}
	}
}

func TestMax2(t *testing.T) {
	var tests = []struct {
		value  int
		values []int
		want   int
	}{
		{1, []int{}, 1},
		{1, []int{2}, 2},
		{3, []int{2, 1}, 3},
	}

	for _, test := range tests {
		v := max2(test.value, test.values...)
		if v != test.want {
			t.Errorf("max2(%d, %v) = %d, want %d", test.value, test.values, v, test.want)
		}
	}
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		value  int
		values []int
		want   int
	}{
		{1, []int{}, 1},
		{1, []int{2}, 1},
		{3, []int{2, 1}, 1},
	}

	for _, test := range tests {
		v := min2(test.value, test.values...)
		if v != test.want {
			t.Errorf("min2(%d, %v) = %d, want %d", test.value, test.values, v, test.want)
		}
	}
}
