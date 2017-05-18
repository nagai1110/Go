// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"+12", "+12"},
		{"-123", "-123"},
		{"1234", "1,234"},
		{"+12345", "+12,345"},
		{"-123456", "-123,456"},
		{"123.456", "123.456"},
		{"+123.456", "+123.456"},
		{"-123.456", "-123.456"},
		{"123456.789012", "123,456.789012"},
		{"+123456.789012", "+123,456.789012"},
		{"-123456.789012", "-123,456.789012"},
	}

	for _, test := range tests {
		s := commaFloat(test.s)
		if s != test.want {
			t.Errorf("commaFloat(%s) = %s, want %s", test.s, s, test.want)
		}
	}
}
